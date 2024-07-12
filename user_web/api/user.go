package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"shop_api/user_web/forms"
	"shop_api/user_web/global"
	"shop_api/user_web/global/response"
	"shop_api/user_web/middlewares"
	"shop_api/user_web/models"
	"shop_api/user_web/proto"
	"strconv"
	"strings"
	"time"
)

func removeTopStruct(fields map[string]string) map[string]string {
	resp := make(map[string]string)
	for field, value := range fields {
		resp[field[strings.Index(field, ".")+1:]] = value
	}
	return resp
}

func HandleGrpcErrorToHttp(err error, ctx *gin.Context) {
	// 将grpc错误转换为http错误
	if err != nil {
		if s, ok := status.FromError(err); ok {
			switch s.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, gin.H{
					"msg": s.Message(),
				})
			case codes.Internal:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				ctx.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				ctx.JSON(http.StatusServiceUnavailable, gin.H{
					"msg": "服务不可用",
				})

			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			}
		}

	}
	return
}
func HandleValidatorError(ctx *gin.Context, err error) {
	// 获取验证器错误
	zap.S().Errorw("[PasswordLogin] 绑定参数失败", "msg", err.Error())
	var errs validator.ValidationErrors
	ok := errors.As(err, &errs)
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})

	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": removeTopStruct(errs.Translate(global.Trans)),
	})
	return
}

func GetUserList(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	zap.S().Infof("[GetUserList] 访问用户 ：%d", currentUser.ID)

	pn := ctx.DefaultQuery("pn", "0")
	pSize := ctx.DefaultQuery("psize", "10")
	pnInt, _ := strconv.Atoi(pn)
	pSizeInt, _ := strconv.Atoi(pSize)

	userListResponse, err := global.UserSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    int32(pnInt),
		PSize: int32(pSizeInt),
	})
	if err != nil {
		zap.S().Errorw("[GetUserList]获取用户列表失败", "msg")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	result := make([]interface{}, 0)
	for _, value := range userListResponse.Data {

		user := response.UserResponse{
			Id:       value.Id,
			NickName: value.NickName,
			Birthday: time.Unix(value.Birthday, 0),
			Mobile:   value.Mobile,
			Gender:   value.Gender,
		}
		result = append(result, user)
	}
	ctx.JSON(http.StatusOK, result)
	zap.S().Debugf("获取用户列表页")
}

func PasswordLogin(ctx *gin.Context) {
	// 验证表单

	passwordLoginForm := forms.PasswordLoginForm{}
	if err := ctx.ShouldBind(&passwordLoginForm); err != nil {
		HandleValidatorError(ctx, err)

		return
	}
	if !store.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, true) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误",
		})
		return
	}

	// 登录
	var resp *proto.UserInfoResponse
	var err error
	if resp, err = global.UserSrvClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: passwordLoginForm.Mobile,
	}); err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusBadRequest, gin.H{
					"msg": "用户不存在",
				})
				return
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误,登陆失败",
				})
				return

			}
		}

	}
	var passResp *proto.CheckResponse
	if passResp, err = global.UserSrvClient.CheckPassword(context.Background(), &proto.PasswordCheckInfo{
		Password:          passwordLoginForm.Password,
		EncryptedPassword: resp.Password,
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"password": "登陆失败",
		})
	} else {
		if passResp.Success {
			// 生成token
			j := middlewares.NewJWT()
			claims := models.CustomClaims{
				ID:          int(resp.Id),
				NickName:    resp.NickName,
				AuthorityId: int(resp.Role),
				StandardClaims: jwt.StandardClaims{
					NotBefore: time.Now().Unix(),                       // 签名生效时间
					ExpiresAt: time.Now().Unix() + 60*60*24*7*60*10000, // 签名失效时间7天
					Issuer:    "andras",
				},
			}
			token, err := j.CreateToken(claims)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "生成token失败",
				})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"code":       200,
				"msg":        "登陆成功",
				"token":      token,
				"expired_at": time.Now().Unix() + 60*60*24*7,
				"data": gin.H{
					"user_id":   resp.Id,
					"user_name": resp.NickName,
				},
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "登陆失败",
			})
		}
	}
}

func Register(ctx *gin.Context) {
	// 用户注册
	registerForm := forms.RegisterForm{}
	if err := ctx.ShouldBind(&registerForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}

	//redis读验证码
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	})
	value, err := rdb.Get(context.Background(), registerForm.Mobile).Result()
	if errors.Is(err, redis.Nil) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": "验证码错误",
		})
		return
	} else {
		if value != registerForm.Code {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": "验证码错误",
			})
			return
		}
	}

	// 调用接口

	var resp *proto.UserInfoResponse
	resp, err = global.UserSrvClient.CreateUser(context.Background(), &proto.CreateUserInfo{
		NickName: registerForm.Mobile,
		Mobile:   registerForm.Mobile,
		Password: registerForm.Password,
	})
	if err != nil {
		zap.S().Errorf("[Register]注册用户失败:%s", err.Error())
		HandleGrpcErrorToHttp(err, ctx)
		return
	}

	// 生成token
	j := middlewares.NewJWT()
	claims := models.CustomClaims{
		ID:          int(resp.Id),
		NickName:    resp.NickName,
		AuthorityId: int(resp.Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),              // 签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*7, // 签名失效时间7天
			Issuer:    "andras",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":         resp.Id,
		"nick_name":  resp.NickName,
		"token":      token,
		"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
	})

}
func GetUserDetail(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	zap.S().Infof("访问用户: %d", currentUser.ID)

	rsp, err := global.UserSrvClient.GetUserById(context.Background(), &proto.IdRequest{
		Id: int32(currentUser.ID),
	})
	if err != nil {
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"name":     rsp.NickName,
		"birthday": time.Unix(int64(rsp.Birthday), 0).Format("2006-01-02"),
		"gender":   rsp.Gender,
		"mobile":   rsp.Mobile,
	})
}

func UpdateUser(ctx *gin.Context) {
	updateUserForm := forms.UpdateUserForm{}
	if err := ctx.ShouldBind(&updateUserForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}

	claims, _ := ctx.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	zap.S().Infof("访问用户: %d", currentUser.ID)

	//将前端传递过来的日期格式转换成int
	loc, _ := time.LoadLocation("Local")
	birthDay, _ := time.ParseInLocation("2006-01-02", updateUserForm.Birthday, loc)
	_, err := global.UserSrvClient.UpdateUser(context.Background(), &proto.UpdateUserInfo{
		Id:       int32(currentUser.ID),
		NickName: updateUserForm.Name,
		Gender:   updateUserForm.Gender,
		Birthday: int64(birthDay.Unix()),
	})
	if err != nil {
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}
