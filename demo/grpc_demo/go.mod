module demo/grpc_demo

go 1.15

require (
	demo/grpc_demo/hello v1.0.0
)

replace (
	demo/grpc_demo/hello => ./hello
)

