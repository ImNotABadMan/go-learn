module demo/test_demo

go 1.15

require (
	demo/test_demo/sumProduct v1.0.0
)

replace (
	demo/test_demo/sumProduct => ./sumProduct
)
