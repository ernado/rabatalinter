# rabatalinter

Linter for rabataio projects.

```bash
go get github.com/ernado/rabatalinter

rabatalinter ./...
```

## TODO

- [ ] RB0 - All receivers should be `m`
- [ ] RB1 - Alphabetical order of constants
- [ ] RB2 - `suite.Require().Error()` before `suite.Require.Is()`
- [ ] RB3 - Always `suite.Require()` on main goroutine