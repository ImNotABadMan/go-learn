module go-traning/error_demos

go 1.15

require (
	go-traning/error_demos/error_pointer v1.0.0
	go-traning/error_demos/error_value v1.0.0
)

replace (
	go-traning/error_demos/error_pointer => ./error_pointer
	go-traning/error_demos/error_value => ./error_value
)
