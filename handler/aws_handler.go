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

func (a AwsHandler) HandlerUpdateAutoScalingGroupAWS(name string, min string, max string, desired string, profile string) {
	cmd := exec.Command("aws",
		"autoscaling",
		"update-auto-scaling-group",
		"--auto-scaling-group-name", name,
		"--min-size", min,
		"--max-size", max,
		"--desired-capacity", desired,
		"--profile", profile)
	cmd.Run()
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
