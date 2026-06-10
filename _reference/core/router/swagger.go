package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labring/aiproxy/core/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
)

func cloneSwaggerInfo(spec *swag.Spec) *swag.Spec {
	newSpec := *spec
	return &newSpec
}

func SetSwaggerRouter(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/doc.json", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json; charset=utf-8")

		if ctx.Request.Host == "" {
			ctx.String(http.StatusOK, docs.SwaggerInfo.ReadDoc())
			return
		}

		swagInfo := cloneSwaggerInfo(docs.SwaggerInfo)
		swagInfo.Host = ctx.Request.Host
		ctx.String(http.StatusOK, swagInfo.ReadDoc())
	})
	router.GET("/swagger", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	router.GET("/swagger/*any",
		ginSwagger.WrapHandler(
			swaggerfiles.Handler,
			ginSwagger.URL("/doc.json"),
		),
	)
}
