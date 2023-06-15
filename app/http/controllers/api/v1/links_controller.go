package v1

import (
	"GOHUB/app/models/link"
	"GOHUB/pkg/response"

	"github.com/gin-gonic/gin"
)

type LinksController struct {
	BaseAPIController
}

func (ctrl *LinksController) Index(c *gin.Context) {

	response.Data(c, link.AllCached())
}
