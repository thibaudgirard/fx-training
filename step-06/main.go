package main

import (
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(NewServer),
		fx.Invoke(StartServer),
		fx.Supply(&Config{Port: 8080}),
		fx.Provide(NewRouter),
		fx.Provide(NewStatusHandler),
	).Run()
}
