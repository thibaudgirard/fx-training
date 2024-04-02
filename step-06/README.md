# Step 05: Decouple registration

Our router constructor has a direct dependency on the `StatusHandler`. This is not ideal as it makes the router constructor less reusable. We can use the `fx` library to decouple the registration of the `StatusHandler` and the `Router`.

## Documentation

FX lets you annotate the dependencies to provide more contextual information.

Here is an example to tell FX to provide a concrete implementation for an interface.

```go
fx.Provide(
    fx.Annotate(NewConcreteTypeConstructor, fx.As(new(InterfaceType))),
),
```

## Tasks

1. Make `StatusHandler` implement the `RouteHandler` interface (see `route_handler.go`).
2. Annotate the `StatusHandler` constructor to provide a `RouteHandler` implementation.
3. Make the `NewRouter` constructor accept a `RouteHandler` dependency instead of a `StatusHandler`.
4. Make the router load the injected `RouteHandler` pattern.

If you have done everything correctly, your application should still work without any difference.

If so, you shall now proceed to the [next step](../step-07/README.md).