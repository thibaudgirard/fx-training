# Step 03: Start & stop HTTP server

We now have an FX application with an HTTP server. 

It could be great to provide some configuration instead of having static strings.

## Documentation

### FX supply

The FX Supply provides instantiated values for dependency injection as if they had been provided using a constructor that simply returns them.

Example:
```go
package main

import (
    "go.uber.org/fx"
)

type Conf struct {
	AppName string
}

type App struct {
	Name string
}

func NewApp(conf *Conf) App {
	return App{Name: conf.AppName}
}

func main() {
      fx.New(
		  fx.Provide(NewApp),
          fx.Supply(&Conf{AppName: "MyApp"}),
      ).Run()
}
```

## Tasks

1. Add in `config.go` a new struct Config with a Port property.
2. In the `main` function, use `fx.Supply` to create an instance of our new config with a port number.
3. Update `server.go` to provide the configuration in the constructor. 

Your web server should be working like the last step. Try to change the port number, and you should see it
in the console log.

If you have the right output, you shall now proceed to the [next step](../step-04/README.md).
