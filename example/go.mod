module helloworld

go 1.14

replace github.com/ynjgit/erpc-go => ../../erpc-go

replace test/helloworld => ./rpc/test/helloworld

require (
	github.com/ynjgit/erpc-go v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.27.1 // indirect
	test/helloworld v0.0.0-00010101000000-000000000000
)
