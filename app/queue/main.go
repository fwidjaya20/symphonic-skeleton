package main

import (
	"github.com/fwidjaya20/symphonic-skeleton/bootstrap"
	"github.com/fwidjaya20/symphonic-skeleton/bootstrap/event"
	"github.com/fwidjaya20/symphonic/facades"
)

func main() {
	bootstrap.Boot()

	facades.Event().Register(event.Kernel{}.Subscribers())

	select {}
}
