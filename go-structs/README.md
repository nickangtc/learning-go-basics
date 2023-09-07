# Structures (structs)

## Embeddings

Embeddings is kind of like inheritance but not really.

You embed another struct by putting name of struct within another struct without specifying a name/type pair. e.g. `Course` instead of `Course Course`:

```go
type Workshop struct {
  Course // this is embedding Course struct
  Date         time.Time
  MaxAttendees int
}
```

See example: https://github.com/firtman/go-fundamentals/blob/a4e261eee4b961e6ab0ff3f4689ac1ae3a61366c/femgo/data/workshop.go#L9C4-L9C4
