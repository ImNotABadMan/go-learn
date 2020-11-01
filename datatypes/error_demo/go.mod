module datatypes/error_demo

go 1.15

require (
	datatypes/struct_demo/inter_struct_demo/my_inter_struct v1.0.0
)

replace (
	datatypes/struct_demo/inter_struct_demo/my_inter_struct => ../struct_demo/inter_struct_demo/my_inter_struct
)

