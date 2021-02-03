module demo/proto_demo/client

go 1.15

require (
	demo/proto_demo/protobuf v1.0.0
	github.com/pkg/errors v0.9.1
)

replace demo/proto_demo/protobuf => ../protobuf
