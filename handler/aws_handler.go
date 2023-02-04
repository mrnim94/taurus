package handler

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v3"
	"net/http"
	"os"
	"os/exec"
	"taurus/log"
	"taurus/model"
	"taurus/model/request"
	"time"
)

type AwsHandler struct {
	Echo *echo.Context
}

const Profile = "--profile"

func (a *AwsHandler) HandlerScheduleAutoScalingGroupAWS() error {

	var cfg model.Schedule
	loadConfigFile(&cfg)

	s := gocron.NewScheduler(time.UTC)

	for i, autoscaling := range cfg.Autoscalings {
		autoscaling := autoscaling
		log.Info("Setup Schedule ", i, " ==> ", autoscaling.Schedule)
		s.Cron(autoscaling.Schedule).Do(func() {
			out, err := exec.Command("aws",
				"autoscaling",
				"update-auto-scaling-group",
				"--auto-scaling-group-name", autoscaling.GroupName,
				"--min-size", autoscaling.Config.Min,
				"--max-size", autoscaling.Config.Max,
				"--desired-capacity", autoscaling.Config.Desired,
				Profile, autoscaling.Profile).Output()
			if err != nil {
				log.Error(err)
			}
			log.Info(out)
		})
	}

	s.StartAsync()
	return nil
}

func loadConfigFile(cfg *model.Schedule) {
	f, err := os.ReadFile("config_file/autoscaling_on_off.yaml")
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		fmt.Println(err)
	}
}

func (a *AwsHandler) HandlerUpdateAutoScalingGroupAWS(c echo.Context) error {
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
		Profile, req.Profile).Output()

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
		Message:    "Updateing AutoScalingGroup-AWS is Successful",
		Data:       string(out),
	})
}

func (a *AwsHandler) HandlerGetAutoScalingGroupAWS(c echo.Context) error {
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
		Profile, req.Profile,
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
