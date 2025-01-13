module app1

go 1.23.2

replace github.com/akm/go-dep-examples/lib1 => ../lib1

require (
	github.com/akm/go-dep-examples/lib1 v0.0.0-00010101000000-000000000000 // indirect
	github.com/akm/go-dep-examples/lib2 v0.0.0-00010101000000-000000000000 // indirect
)

replace github.com/akm/go-dep-examples/lib2 => ../lib2
