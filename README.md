# Learning Go

Repo for hands-on learning of Go basics. Companion course: [Basics of Go](https://frontendmasters.com/courses/go-basics/) by Max Firtman.

Note: There are readmes in most folders, so read those for more specific notes.

Note also: This repo isn't representative of a typical Go project. I valued having everything in a single repo while learning.

## Generally interesting things about Go

Compiled language quirk. When you run `go run .`, it compiles a new binary. This has interesting downstream effects, like if you have a firewall blocking apps from making outgoing HTTP requests, you'll need to re-allow the app on every run.

Filenames are not used much in Go. Package names, defined as first line in Go files, are what's important. e.g. standard library's `json.Unmarshal` function is defined in a file called decode.go.

Working with JSON. Use online tools to convert JSON responses into Go structs. You need this because to export types, you must use TitleCase for key names, and that doesn't usually match how JSON keys are. e.g.:

```json
{
  "timestamp": "168484000199395"
}
```

```go
type AutoGenerated struct {
  Timestamp string `json:"timestamp"`
}
```

Mostly synchronous. `http.Get(...)` is a sync operation. Most stuff are sync in Go by default, unless you use goroutines with `go func()()`, which you will need for performance reasons. To do that, and ensure that the main process doesn't exit before the calls have received responses, you can use `sync.WaitGroup`:

```go
func main() {
	currencies := []string{"BTC", "BCH", "ETH"}
	var wg sync.WaitGroup
	for _, currency := range currencies {
		wg.Add(1)
		go func(currencyCode string) {
			getCurrencyData(currencyCode)
			wg.Done()
		}(currency)
	}
	wg.Wait()
}
```

Array vs slice. Array has fixed size and cannot change once initialised. Slice has dynamic length, but under the hood is actually connecting multiple Arrays. Slices are much more commonly used, but Arrays exist for optimisation purposes.

`[3]string` is an array of size 3, string type.

`[]string` is a slice of dynamic size, string type.

## Testing in Go

Built in test runner `go test`.

Uses "Table driven test" design pattern, which involves setting up a table and having combination of inputs as rows.

Name files xxx_test.go and it will be recognised as a test. Compiler will ignore all test files for build.

The built in way of testing involves throwing errors using the `testing.T` function. No errors thrown during test = passing test.

You don't need any additional frameworks to write Go tests. They exist but they are not necessary.

## WebAssembly

Go can be compiled to WebAssembly and be run on browsers.

Go can also be compiled to JS and then be run on browsers, but not many people do that, cos why would you?
