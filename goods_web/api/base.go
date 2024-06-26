package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"shop_api/goods_web/global"
	"strings"
)

func RemoveTopStruct(fields map[string]string) map[string]string {
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
		"msg": RemoveTopStruct(errs.Translate(global.Trans)),
	})
	return
}
