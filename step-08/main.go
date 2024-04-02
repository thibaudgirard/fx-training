package main

import (
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(NewServer),
		fx.Provide(AsRouteHandler(NewHelloHandler)),
		fx.Provide(AsRouteHandler(NewStatusHandler)),
		fx.Provide(fx.Annotate(NewRouter, fx.ParamTags(`group:"routeHandlers"`))),
		fx.Provide(NewLogger),
		fx.Invoke(StartServer),
	).Run()
}

func AsRouteHandler(h any) any {
	return fx.Annotate(
		h,
		fx.As(new(RouteHandler)),
		fx.ResultTags(`group:"routeHandlers"`),
	)
}
