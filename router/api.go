package router

import (
	"github.com/labstack/echo/v4"
	"taurus/handler"
)

type API struct {
	Echo       *echo.Echo
	AwsHandler handler.AwsHandler
}

func (api *API) SetupRouter() {
	api.Echo.GET("/", handler.HandlerXxx)

	user := api.Echo.Group("/aws")
	user.POST("/get-autoscaling-group", api.AwsHandler.HandlerGetAutoScalingGroupAWS)
	user.POST("/update-autoscaling-group", api.AwsHandler.HandlerUpdateAutoScalingGroupAWS)
}
