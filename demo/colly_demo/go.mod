module demo/colly_demo

go 1.15

require (
	demo/colly_demo/gl_colly v1.0.0
	github.com/gocolly/colly/v2 v2.1.0
)

replace demo/colly_demo/gl_colly => ./gl_colly
