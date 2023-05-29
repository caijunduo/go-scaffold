package common_api

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Paginate struct {
	Page     int `form:"page" json:"page"`
	PageSize int `form:"page_size" json:"page_size"`
}

func (req *Paginate) Validate(c *gin.Context) error {
	_ = c.ShouldBindQuery(req)
	_ = c.ShouldBindJSON(req)
	if err := validation.Validate(&req.Page,
		validation.Required.Error("页数不能小于1"),
		validation.Min(1).Error("页数不能小于1")); err != nil {
		return err
	}
	if err := validation.Validate(&req.PageSize,
		validation.Required.Error("每页数量不在1-50之间"),
		validation.Length(1, 50).Error("每页数量不在1-50之间")); err != nil {
		return err
	}
	return nil
}
