package main

import (
	"dojo-fx/step-10/config"
	"dojo-fx/step-10/handler"
	"dojo-fx/step-10/http"
	"dojo-fx/step-10/log"
	"dojo-fx/step-10/repository"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		handler.Module,
		repository.Module,
		http.Module,
		log.Module,
		config.Module,
	).Run()
}
