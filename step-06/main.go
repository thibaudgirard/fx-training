package main

import (
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(NewServer),
		fx.Invoke(StartServer),
		fx.Provide(NewRouter),
		fx.Provide(AsRouteHandler(NewStatusHandler)),
		fx.Provide(NewLogger),
	).Run()
}

func AsRouteHandler(h any) any {
	return fx.Annotate(
		h,
		fx.As(new(RouteHandler)),
		// Hum, something is missing here...
	)
}
