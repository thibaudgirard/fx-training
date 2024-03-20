package main

import (
	"dojo-fx/step-08/handler"
	"dojo-fx/step-08/http"
	"dojo-fx/step-08/log"
	"dojo-fx/step-08/repository"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		handler.Module,
		repository.Module,
		http.Module,
		log.Module,
	).Run()
}
