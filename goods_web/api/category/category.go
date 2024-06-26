package category

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"shop_api/goods_web/api"
	"shop_api/goods_web/forms"
	"shop_api/goods_web/global"
	"shop_api/goods_web/proto"
	"strconv"
)

// CategoryRouter.GET("", category.List)          // 商品类别列表页
// CategoryRouter.DELETE("/:id", category.Delete) // 删除分类
// CategoryRouter.GET("/:id", category.Detail)    // 获取分类详情
// CategoryRouter.POST("", category.New)          //新建分类
// CategoryRouter.PUT("/:id", category.Update)    //修改分类信息
func List(ctx *gin.Context) {
	resp, err := global.GoodsSrvClient.GetAllCategorysList(ctx, &emptypb.Empty{})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	data := make([]interface{}, 0)
	err = json.Unmarshal([]byte(resp.JsonData), &data)
	if err != nil {
		zap.S().Errorw("[List] 查询分类列表失败", "msg", err.Error())
	}
	ctx.JSON(http.StatusOK, data)
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	_, err = global.GoodsSrvClient.DeleteCategory(ctx, &proto.DeleteCategoryRequest{Id: int32(i)})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.Status(http.StatusOK)
}
func Detail(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	reMap := make(map[string]interface{})
	subCategorys := make([]interface{}, 0)
	if r, err := global.GoodsSrvClient.GetSubCategory(context.Background(), &proto.CategoryListRequest{
		Id: int32(i),
	}); err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	} else {
		//写文档 特别是数据多的时候很慢， 先开发后写文档
		for _, value := range r.SubCategorys {
			subCategorys = append(subCategorys, map[string]interface{}{
				"id":              value.Id,
				"name":            value.Name,
				"level":           value.Level,
				"parent_category": value.ParentCategory,
				"is_tab":          value.IsTab,
			})
		}
		reMap["id"] = r.Info.Id
		reMap["name"] = r.Info.Name
		reMap["level"] = r.Info.Level
		reMap["parent_category"] = r.Info.ParentCategory
		reMap["is_tab"] = r.Info.IsTab
		reMap["sub_categorys"] = subCategorys

		ctx.JSON(http.StatusOK, reMap)
	}
	return
}
func New(ctx *gin.Context) {
	categoryForm := forms.CategoryForm{}
	if err := ctx.ShouldBindJSON(&categoryForm); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}

	rsp, err := global.GoodsSrvClient.CreateCategory(context.Background(), &proto.CategoryInfoRequest{
		Name:           categoryForm.Name,
		IsTab:          *categoryForm.IsTab,
		Level:          categoryForm.Level,
		ParentCategory: categoryForm.ParentCategory,
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	request := make(map[string]interface{})
	request["id"] = rsp.Id
	request["name"] = rsp.Name
	request["parent"] = rsp.ParentCategory
	request["level"] = rsp.Level
	request["is_tab"] = rsp.IsTab

	ctx.JSON(http.StatusOK, request)

}

func Update(ctx *gin.Context) {
	categoryForm := forms.UpdateCategoryForm{}
	if err := ctx.ShouldBindJSON(&categoryForm); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}

	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	request := &proto.CategoryInfoRequest{
		Id:   int32(i),
		Name: categoryForm.Name,
	}
	if categoryForm.IsTab != nil {
		request.IsTab = *categoryForm.IsTab
	}
	_, err = global.GoodsSrvClient.UpdateCategory(context.Background(), request)
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	ctx.Status(http.StatusOK)
}
