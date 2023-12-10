package main

import (
	"log"
	"os"

	"github.com/fwidjaya20/symphonic-skeleton/app/grpc"
	"github.com/fwidjaya20/symphonic-skeleton/app/http"
	"github.com/fwidjaya20/symphonic-skeleton/app/schedule"
	"github.com/fwidjaya20/symphonic-skeleton/bootstrap"
)

func main() {
	bootstrap.Boot()

	runner, ok := os.LookupEnv("SERVER_TO_RUN")

	if !ok {
		log.Fatal("env SERVER_TO_RUN not configured")
		return
	}

	switch runner {
	case "grpc":
		grpc.RunServer()
		return
	case "http":
		http.RunServer()
		return
	case "schedule":
		schedule.RunServer()
		return
	}
}
