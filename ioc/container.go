package ioc

import (
	"sync"
)

type Container struct {
}

var (
	instanceIoC Container
	syncIoC     sync.Once
)

func Injector() Container {
	syncIoC.Do(func() {
		instanceIoC = newContainer()
	})

	return instanceIoC
}

func newContainer() Container {
	return Container{}
}
