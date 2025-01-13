# go-dep-examples

## app1/cmd/add

This module dependa on lib1.

```
$ go run ./cmd/add 5 8
lib1/add.go init
app1/add_string.go init
lib1/add_string_test_helper.go init
app1/cmd/add/main.go init
13
```

## app1/cmd/sub

This module dependa on lib2 which depends on lib1.

```
$ go run ./cmd/sub 5 8
lib1/add.go init
app1/add_string.go init
lib1/add_string_test_helper.go init
lib2/sub.go init
app1/cmd/sub/main.go init
-3
```

## lib2 module is not loaded from app1/cmd/add

The dependency to lib2 is defined in app1/go.mod but lib2 is not used
from app1/cmd/add. So app1/cmd/sub outputs `lib2/sub.go init` but
app1/cmd/add does not output it.

## WARNING

`lib1/add_string_test_helper.go` is loaded from app1/cmd/add and app1/cmd/sub .
It is used in lib1/add_string_test.go but its name does not match `*_test.go` so
it is treated as a code not for test.
