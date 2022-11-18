# Calling Dagger from `go test`

**Do NOT try this at home....or anywhere else :)**

This is a crazy experiment showing that [Dagger](https://dagger.io/) can be called from a simple Go test.
To avoid an endless cycle of Dagger running dagger again though, the appropriate test is guarded so it needs to be called by name:

```shell
go test -run TestDagger -v .
```
