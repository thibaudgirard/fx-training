# Step 06: 

In a more realistic scenario, we would be able to injection multiple HTTP handler in our router without having to manually inject each of them.

## Documentation

### FX Tag

We've already seen that we can annotate de dependency in FX using `fx.Annotate`. Now we will use the same function to add a _tag_ annotation.

We will use the tag `group` which allows to create collection.

```go
type MyType struct {
    // ...
}

func NewMyType() *MyType {
    // ...
}

func NewMyTypeBis() *MyType {
    // ...
}


func INeedAMyTypeCollection(_ []MyType) {
    // ...
}

func main() {
    fx.New(
        fx.Provide(
            fx.Annotate(NewMyType), fx.ResultTags(`group:"mytypes"`)
        ),
        fx.Provide(
            fx.Annotate(NewMyTypeBis), fx.ResultTags(`group:"mytypes"`)
        ),
        fx.Invoke(
            fx.Annotate(INeedAMyTypeCollection, fx.ParamTags(`group:"mutypes"`)),
        ),
    ).Run()
}
```

 - `fx.ResultTags(`group:"mytypes"`)` tells FX that the provided result is part of a group named `mytypes`. 
 - `fx.ParamTags(`group:"mytypes"`)` tells FX that the parameter is the group formed by all results tagged with `group:"mytypes"`.

## Tasks

1. In `main.go`, fix the `AsRouteHandler` function implementation: it should return a tagged fx service to be injected in a route handler collection.
2. Use this function to provide and tag the `HelloHandler` and `StatusHandler` in `main.go`.
3. Make the `NewRouter` function accept a collection of `RouteHandler`.
4. Make sure FX will inject the tagged collection in `NewRouter`.

Now run your application. you should be able to call the `/hello` and `/status` routes.

If so, you shall now proceed to the [final step](../step-final/README.md).