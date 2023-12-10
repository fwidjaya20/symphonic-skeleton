package config

import (
	"github.com/fwidjaya20/symphonic/facades"
)

func init() {
	config := facades.Config()

	config.Add("grpc", map[string]any{
		"host": config.Get("GRPC_HOST", "localhost"),
		"port": config.Get("GRPC_PORT", "9000"),
	})
}
