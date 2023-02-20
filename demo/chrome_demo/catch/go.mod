module demo/chrome_demo/catch

go 1.15


require (
	"demo/chrome_demo/catch/code" v1.0.0
)

replace (
	demo/chrome_demo/catch/code => ./code
)