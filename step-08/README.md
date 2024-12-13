# Step 07:  Decoupled multiple injections

In a more realistic scenario, we would be able to inject multiple HTTP handlers in our router without having to manually inject each of them.

## Documentation

### FX Tag

We've already seen that we can annotate a dependency in FX using `fx.Annotate`. Now we will use the same function to add a _tag_ annotation.

We will use the tag `group` which allows to create a named collection.

```go
// A new custom type
type MyType struct {
    // ...
}

// A constructor for MyType
func NewMyType() *MyType {
    // ...
}

// Another constructor for MyType
func NewMyTypeBis() *MyType {
    // ...
}

// A function that needs a collection of MyType
func INeedAMyTypeCollection(_ []MyType) {
    // ...
}

func main() {
    fx.New(
        fx.Provide(
            // Put the result in the "mytypes" group/collection
            fx.Annotate(NewMyType), fx.ResultTags(`group:"mytypes"`)
        ),
        fx.Provide(
            // Put the result in the "mytypes" group/collection
            fx.Annotate(NewMyTypeBis), fx.ResultTags(`group:"mytypes"`)
        ),
        fx.Invoke(
            // Use the "mytypes" group/collection as a parameter
            fx.Annotate(INeedAMyTypeCollection, fx.ParamTags(`group:"mytypes"`)),
        ),
    ).Run()
}
```

 - ``fx.ResultTags(`group:"mytypes"`)`` tells FX that the provided result is part of a group named `mytypes`. 
 - ``fx.ParamTags(`group:"mytypes"`)`` tells FX that the parameter is the group formed by all results tagged with `group:"mytypes"`.

## Tasks

1. In `main.go`, fix the `AsRouteHandler` function implementation: it should return a tagged fx service to be injected in a `routeHandlers` group/collection.
2. Use this function to provide and tag the `HelloHandler` and `StatusHandler` in `main.go`.
3. Make the `NewRouter` function accept a collection of `RouteHandler`.
4. Make sure FX will inject the tagged collection in `NewRouter`.

Now run your application. You should be able to call the `/hello` and `/status` routes.

If so, you shall now proceed to the [next step](../step-09/README.md).