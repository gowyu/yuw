package controllers

import (
	"github.com/gin-gonic/gin"
	E "github.com/gowyu/yuw/exceptions"
	s "github.com/gowyu/yuw/mvc/services"
	"net/http"
)

type (
	Controllers struct {
		Srv *s.Services
	}
)

func New() *Controllers {
	return &Controllers {
		Srv: s.New(),
	}
}

func (c *Controllers) To(ctx *gin.Context, res *s.PoT) {
	switch res.Type {
	case s.ToJSON:
		ctx.JSON(
			res.Code,
			res.Response,
		)

		return

	case s.ToHTML:
		ctx.HTML(
			res.Code,
			res.HTML,
			res.Response,
		)

		return

	default:
		ctx.AbortWithError(
			http.StatusNotFound,
			E.Err("yuw^error",E.ErrPosition()),
		)

		return
	}
}

