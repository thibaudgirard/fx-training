package main

import (
	"fmt"
	"go.uber.org/fx"
)

func main() {
	fmt.Println("FX version", fx.Version)
}
