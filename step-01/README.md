# Step 01: Create an empty FX application

## Documentation

To create a new application, use `fx.New()` which returns an `*fx.App`. 
Then, you can call `Run()` on the newly created application to start it.

## Tasks
1. Create a new FX application
2. Run the application

Running your application should display the following output:

```bash
[Fx] PROVIDE    fx.Lifecycle <= go.uber.org/fx.New.func1()
[Fx] PROVIDE    fx.Shutdowner <= go.uber.org/fx.(*App).shutdowner-fm()
[Fx] PROVIDE    fx.DotGraph <= go.uber.org/fx.(*App).dotGraph-fm()
[Fx] RUNNING
```

Nothing happens until you press ctrl+c. As stated in the [documentation](https://uber-go.github.io/fx/get-started/minimal.html):

> Fx is primarily intended for long-running server applications; these applications typically receive a signal from the deployment system when it's time to shut down.

You can also see that even if you just created an empty application, FX is already providing some useful features like `fx.Lifecycle`, `fx.Shutdowner` and `fx.DotGraph`.

 - `fx.Lifecycle` is used to manage the lifecycle of the application. We'll talk more about it later.
 - `fx.Shutdowner` is used to shutdown the application gracefully.
 - `fx.DotGraph` is used to generate a graph of the application's dependencies. It's provided as a string that can be given to a [graphviz tool](https://dreampuf.github.io/GraphvizOnline/) to generate a visual representation of the application's dependency graph.

All good? You shall now proceed to the [next step](../step-02/README.md).
