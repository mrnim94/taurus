package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os/exec"
	"taurus/log"
	"taurus/model"
	"taurus/model/request"
)

type AwsHandler struct {
}

func (a AwsHandler) HandlerUpdateAutoScalingGroupAWS(c echo.Context) error {
	req := request.ReqUpdateAutoScalingGroup{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	out, err := exec.Command("aws",
		"autoscaling",
		"update-auto-scaling-group",
		"--auto-scaling-group-name", req.Name,
		"--min-size", req.Min,
		"--max-size", req.Max,
		"--desired-capacity", req.Desired,
		"--profile", req.Profile).Output()

	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "thành công",
		Data:       string(out),
	})
}

func (a AwsHandler) HandlerGetAutoScalingGroupAWS(c echo.Context) error {
	req := request.ReqAutoScalingGroup{}

	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	//https://stackoverflow.com/questions/1877045/how-do-you-get-the-output-of-a-system-command-in-go
	out, err := exec.Command("aws", "autoscaling", "describe-auto-scaling-groups",
		"--auto-scaling-group-names", req.Name,
		"--query", "AutoScalingGroups[*].{MinSize:MinSize,MaxSize:MaxSize,DesiredCapacity:DesiredCapacity}",
		"--profile", req.Profile,
		"--output", "json").Output()
	//cmd.Stderr = os.Stderr

	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "thành công",
		Data:       string(out),
	})
}
