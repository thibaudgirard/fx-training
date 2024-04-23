# Step 00: Here it begins

## Check your environment

Before going to the very first step of this journey, we need to make sure your system can run the code we will be writing.

First, check if you have Go installed with the right version `1.22.X`. Then, run the following command:

```bash
make run-00
```

It should display the FX version installed.

## What are we going to do?

We're going to create a simple application that handles and responds to HTTP requests.

But our real goal is to learn how to use dependency injection with FX.

## Project structure

Each step has its own directory where you'll find a `README.md` file with a documentation part providing useful information to perform the tasks.

You'll also find, at least, a `main.go` file that contains the main application code.

To run a step, you can use the following command:

```bash
make run-XX
```

Where `XX` is the step number.

You can also run the following command to watch each step and rebuild the application automatically:

```bash
make watch-XX
```

The watch command requires [`entr`](https://github.com/clibs/entr) to be installed on your system.

If you're struggling with a step, you can find the solution in the next step directory.

## Ready ?

You shall now proceed to the [next step](../step-01/README.md).
