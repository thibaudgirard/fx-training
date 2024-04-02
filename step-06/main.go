package main

import (
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(NewServer),
		fx.Invoke(StartServer),
		fx.Provide(NewRouter),
		fx.Provide(NewStatusHandler),
		fx.Provide(NewLogger),
	).Run()
}
