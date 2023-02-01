package schedule

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"time"
)

type GoCron struct {
	GR           *gocron.Scheduler
	ZoneName     string
	SecondsOfUTC int
}

func (gc *GoCron) GetGoCron() {
	gc.GR = gocron.NewScheduler(time.FixedZone(gc.ZoneName, gc.SecondsOfUTC))
	fmt.Println("Scheduler is beginning")
}
