package app

import (
	"go-project-structure-example/internal/utils/e"
	"net/http"

	"github.com/go-playground/validator"

	"github.com/gin-gonic/gin"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	v := validator.New()
	if errs := v.Struct(form); errs != nil {
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}
