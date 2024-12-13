package config

import "go.uber.org/fx"

var Module = fx.Module(
	"http handlers",
	fx.Supply(&Config{
		Port: 8080,
	}),
)
