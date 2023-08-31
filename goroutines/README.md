# Goroutines

Just by adding `go` keyword in front of ANY function call, go takes care of creating another process to execute it.

Under the hood it's more complex, but can think of it as multi-threading. Real multi-threading may depend on presence of multiple CPU cores.

Great use of `time.Sleep` to teach simultaneous goroutines by Maximiliano Firtman.

## Channels

You can use `chan` and `make` to pass messages around goroutines.
