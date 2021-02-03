module demo/proto_demo/server

go 1.15

require (
	github.com/pkg/errors v0.9.1
	demo/proto_demo/protobuf v1.0.0
)

replace demo/proto_demo/protobuf => ../protobuf
