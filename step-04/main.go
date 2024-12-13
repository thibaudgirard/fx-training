package main

import (
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(NewServer),
		fx.Invoke(ServerStart),
		fx.Supply(&Config{Port: 8080}),
	).Run()
}
