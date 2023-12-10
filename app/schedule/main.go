package schedule

import (
	"github.com/fwidjaya20/symphonic-skeleton/bootstrap"
	"github.com/fwidjaya20/symphonic/facades"
)

func RunServer() {
	bootstrap.Boot()

	facades.App().GetSchedule().Run()

	select {}
}
