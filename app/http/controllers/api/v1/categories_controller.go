package v1

import (
	"GOHUB/app/models/category"
	"GOHUB/app/requests"
	"GOHUB/pkg/response"

	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	BaseAPIController
}

func (ctrl *CategoriesController) Store(c *gin.Context) {

	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}

	categoryModel := category.Category{
		Name:        request.Name,
		Description: request.Description,
	}
	categoryModel.Create()
	if categoryModel.ID > 0 {
		response.Created(c, categoryModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

// Update更新分类接口
// 控制器
func (ctrl *CategoriesController) Update(c *gin.Context) {

	// 验证 url 参数 id 是否正确
	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}

	// 表单验证
	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}

	// 保存数据
	categoryModel.Name = request.Name
	categoryModel.Description = request.Description
	rowsAffected := categoryModel.Save()

	if rowsAffected > 0 {
		response.Data(c, categoryModel)
	} else {
		response.Abort500(c)
	}
}

func (*CategoriesController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}

	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}
	data, pager := category.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

//删除分类接口

func (ctrl *CategoriesController) Delete(c *gin.Context) {

	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}

	rowsAffected := categoryModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}