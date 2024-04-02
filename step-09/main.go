package main

import (
	"dojo-fx/step-09/handler"
	"dojo-fx/step-09/http"
	"dojo-fx/step-09/log"
	"dojo-fx/step-09/repository"
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
