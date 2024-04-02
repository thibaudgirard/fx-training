# Step 02: Add some dependency injection

## Documentation

We now have a running FX application. Let's add some dependency injection to it.


The `fx.New()` function can receive a variadic list of `fx.Option`. We'll use two functions returning an `fx.Option`:
- `fx.Provide` Take a constructor function as argument. FX will use reflection to get the return type of this function 
- `fx.Invoke` Before the application starts, FX will call the function passed as argument

In both of these functions, FX will inspect the given function and check arguments types to match it with provided dependencies. If everything is correct, the function will be called with the dependencies as arguments.

Example:

```go
package main

import (
    "go.uber.org/fx"
)

type MyType struct {
}

func NewMyType() *MyType {
    return &MyType{}
}

func INeedMyType(myType *MyType) {
    // Do something with myType
}

func main() {
    fx.New(
        fx.Provide(NewMyType),
        fx.Invoke(INeedMyType),
    ).Run()
}
```

## Tasks

In this step, you will find a `server.go` file that already contains the required code to create a new server instance.

1. Provide the server instance in your application 
2. Make FX call the `ServerStart` function

For now, the `ServerStart` function only displays a message to confirm that the function has been called with a server instance injected as a dependency.

Running your application should display `Server started on :8080`.

If so, you shall now proceed to the [next step](../step-03/README.md).
