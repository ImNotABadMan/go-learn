module demo/chrome_demo

go 1.15

require (
	demo/chrome_demo/click_demo v1.0.0
	demo/chrome_demo/gl_csv_import v1.0.0
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect
)

replace (
	demo/chrome_demo/click_demo => ./click_demo
	demo/chrome_demo/gl_csv_import => ./gl_csv_import
)
