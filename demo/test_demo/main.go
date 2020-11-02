package main

import (
	"demo/test_demo/sumProduct"
	"fmt"
)

func main() {
	// 产品列表
	productList := sumProduct.ProductList{}
	product := sumProduct.Product{}
	product.SetData("camera", 150.5)

	product1 := sumProduct.Product{}
	product1.SetData("computer", 2000.5)

	productList.Add(&product)
	productList.Add(&product1)

	fmt.Println(productList)
	fmt.Println("sum:", productList.Sum())
}
