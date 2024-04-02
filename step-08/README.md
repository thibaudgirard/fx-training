# Step 07: Modules

You can now try to split your code in multiple packages and use fx module to organize your code.

## Documentation

`fx.Module("name', ...)` lets you provide multiples services. A module still needs to be registered in the main app.

```go
package http

import "go.uber.org/fx"

var Module = fx.Module(
	"http", 
	fx.Provide(NewServer),
	fx.Provide(NewRouter),
	// ...
)
```

```go
package main

import (
	"go.uber.org/fx"
	"http"
)

func main() {
    fx.New(
        http.Module,
    ).Run()
}
```

## Tasks

1. Reorganize your code in multiple packages using FX modules to handle dependencies injection for each of them.

There is no perfect answer to this last step. But you can check the [next step](../step-09/README.md) for a possible solution.

