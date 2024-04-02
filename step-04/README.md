# Step 03: Add a router and a first handler

We now want to be able to handle a request on a `/status` route and respond with the `OK` string.

## Documentation

In the `router.go` file, you'll find a new function to create a new Chi router.

This router can be given to our server as the main router to handle requests:

```go
http.Server{Addr: ":8080", Handler: chiRouter}
```

You can provide a handler for a specific route using the `Method` function on the router:

```go
chiRouter.Method(http.MethodGet, "/status", statusHandler)
```

## Tasks

1. Provide the `NewRouter` function int the FX app
2. Use the router in the `NewServer` function as an HTTP server handler
3. Create a new `StatusHandler` struct implementing the `http.Handler` interface, it should respond with the `OK` string
4. Create a `NewStatusHandler` function and provide it to the FX app
5. Update the `NewRouter` function to use the `StatusHandler` for the `GET /status` route

Run the application, the following command should return `OK`:

```sh
curl http://localhost:8080/status
```

Any other route should return a `404 Not Found` error.

If you get the right responses, you shall now proceed to the [next step](../step-05/README.md).