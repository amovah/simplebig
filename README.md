# simplebig

simplebig is a wrapper around `math/big` built-in package and main goal is convient to use `math/big`.

## Why?

### Short Version

trying to solve these with simple wrapper around `math/big`:

- no need to pass first parameter in most functions:

```go
numb := simpleint.New(12)
reuslt := numb.Add(simpleint.New(13))
fmt.Println(result) // 25
```

- return instead of modify original value, it's thread-safe:

```go
numb := simpleint.New(12)
reuslt := numb.Sub(simpleint.New(13))
fmt.Println(numb, result) // 12 -1
```

- `sql.Scanner` and `sql/driver.Valuer` interfaces are implemented, `Int` and `Float` can be used
  directly for databases.
- some helpers for convient like `Pow` or converting functions between
  `simplebig.Float` and
  `simplebig.Int` are added.

- method chaining:

```go
x := simpleint.New(2)
result := x.Add(simpleint.New(13)).Sub(simpleint.New(5)).Mul(simpleint.New(2)).Pow(simpleint.New(3))
fmt.Println(result, x) // 8000 2
```

- it's only a wrapper, so it does whatever `math/big` do and it's accurate as `math/big` is accurate.

## Example

```go
x := simpleint.New(2).Add(simpleint.New(-5)).Mul(simpleint.New(-1))

// get copy of underlying big.Int
x.BigInt()

fmt.Println(x, x.BigInt())

// safe to use in gorm models
type User struct {
    Name string
    Balance simepleint.Int
}
```

## Database Compatibile and Break Changes:

`simplebig.Int` implements two interface to compatibile types with using in databases, but this changes breaks one
feature and only this one from `math/big`. `math/big.Int` implements `fmt.Scanner` but
`simplebig.Int` implements `sql.Scanner`. due to conflict in names of these two
interfaces, I decided to choose one of them which is database compatibility. because there is
a workaround for `fmt.Scanner` but there is no workaround for database compatibility.

## Todos

- [ ] implement wrapper for ModInverse and ModSqrt on `simpleint.Int`
- [ ] implement wrapper for `math/big.Float`
- [ ] implement wrapper for `math/big.Rat`

## LICENSE

MIT
