module second/main

go 1.15


require (
	second/learn v1.0.0
)

replace (
	second/learn => ../learn
)