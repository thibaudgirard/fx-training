package main

import (
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(NewServer),
		fx.Invoke(StartServer),
	).Run()
}
