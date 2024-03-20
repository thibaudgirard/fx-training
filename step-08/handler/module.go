package handler

import "go.uber.org/fx"

func AsRouteHandler(h any) any {
	return fx.Annotate(
		h,
		fx.As(new(Handler)),
		fx.ResultTags(`group:"routeHandlers"`),
	)
}

var Module = fx.Module(
	"http handlers",
	fx.Provide(AsRouteHandler(NewHelloHandler)),
	fx.Provide(AsRouteHandler(NewStatusHandler)),
)
