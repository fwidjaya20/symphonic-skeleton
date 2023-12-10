package schedule

import (
	"github.com/fwidjaya20/symphonic/contracts/foundation"
	"github.com/fwidjaya20/symphonic/facades"
)

type TaskSchedulerServiceProvider struct {
}

func (sp *TaskSchedulerServiceProvider) Boot(_ foundation.Application) {}

func (sp *TaskSchedulerServiceProvider) Register(_ foundation.Application) {
	kernel := Kernel{}

	facades.Schedule().Register(kernel.Schedule())
}
