package routes

import (
	v1 "serialno/controller/v1"

	"github.com/gin-gonic/gin"
)

func registeRoute(router *gin.Engine) {
	router.POST("/v1/order-no", v1.OrderNo)
	router.POST("/v1/order-no-set", v1.PrefixNum)
}
