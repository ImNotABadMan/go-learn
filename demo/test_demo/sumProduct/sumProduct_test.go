package sumProduct

import (
	"fmt"
	"testing"
)

// 功能测试
func TestProduct_SetData(t *testing.T) {
	product := Product{}
	if product.title != "" {
		t.Error("product.title is not empty")
	}

	product.SetData("test product", 100.1)

	if product.title != "test product" {
		t.Error("product title is not 'test product'")
	}

	if product.price != 100.1 {
		t.Error("product price != 100.1")
	}
}

func TestFail(t *testing.T) {
	// fail还是会继续执行
	//t.Fail()
	// failNow不会继续执行
	//t.FailNow()
	t.Log("test fail")
}

func TestProductList_Sum(t *testing.T) {
	productList := ProductList{}
	sum := productList.Sum()
	if sum == 0 {
		t.Log("sum is 0")
	}
}

// 性能测试
func BenchmarkProduct_SetData(b *testing.B) {
	product := Product{}
	for i := 0; i < b.N; i++ {
		product.SetData("test product", 100.1)
	}
	fmt.Println(b.N)
}
