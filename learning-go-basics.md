## Go CLI

go.mod is like package.json

```sh
go mod init someurl/looking/thing
# => go.mod file
```

## Vars

var
const

```go
var x int
var name string
const y = 2
```

var initialisation shortcut

```go
// works only in function, not global scope
text := "Bye!"
```

If you don't use `:=` then you'll be initialising a variable > set to `nil` > then assign value.

Implicitly typed.

Don't say "global variable" but "package variable" instead.

## Packages

Compiler sees packages, not files. Combines files that are part of the same package into one big package.

Package variables are accessible in any file that is part of the same package. So, fewer imports.

Current package can access all functions and variables by default. Public methods are methods meant to be accessed outside of the current package. Public methods in a package must use TitleCase:

```go
package main

// this would be exported
func SendMoney() {
  println("sending money")
}
```

... and can be used in another package:

```go
package banking

import "main"

func main() {
  // invoking imported methods
  banking.SendMoney()
}
```

Imports work per file, so if 2 files under the same package need to use the same import, they both need to do `import "fmt"` for example.

## Function arguments

Arguments are passed by value by default:

```go
func birthday(age int) {
  age += 1
}

func main() {
  var age = 22
  birthday(age)
  fmt.Println(age) //=> 22, not 23
}
```

But you can override that and pass by pointer:

```go
func birthday(age *int) {
  *age += 1
}

func main() {
  var age = 22
  birthday(&age)
  fmt.Println(age) //=> 22, not 23
}
```

## panic

Is NOT for try catch or error management.

## defer

Just something to run something after a method is finished executing.

```go
func main() {
  defer print("too old to be true")

  print("ageing....")
}

main() //=> "ageing....too old to be true"
```

Example use case: things you don't wanna forget to do later, like closing database connection after opening.

If there's a `panic` call in the method, it will honour the `defer` anyway and execute it.
