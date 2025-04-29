# rabatalinter

Linter for rabataio projects.

```bash
go get github.com/ernado/rabatalinter
```

```bash
rabatalinter ./...
```

## TODO

- [x] RB0 - All receivers should be in `m`, `suite` list
- [ ] RB1 - Alphabetical order of constants
- [ ] RB2 - `suite.Require().Error()` before `suite.Require.Is()`
- [ ] RB3 - Always `suite.Require()` on main goroutine