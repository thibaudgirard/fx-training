package http

import "go.uber.org/fx"

var Module = fx.Module(
	"http",
	fx.Provide(NewServer),
	fx.Invoke(StartServer),
	fx.Provide(fx.Annotate(NewRouter, fx.ParamTags(`group:"routeHandlers"`))),
)
