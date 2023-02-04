package main

import (
	"github.com/labstack/echo/v4"
	"os"
	"taurus/handler"
	"taurus/log"
	"taurus/router"
)

func init() {
	os.Setenv("APP_NAME", "aws")
	log.InitLogger(false)
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")

}

func main() {
	e := echo.New()

	awsHandler := handler.AwsHandler{}

	awsHandler.HandlerScheduleAutoScalingGroupAWS()

	api := router.API{
		Echo:       e,
		AwsHandler: awsHandler,
	}
	api.SetupRouter()

	e.Logger.Fatal(e.Start(":1323"))
}
