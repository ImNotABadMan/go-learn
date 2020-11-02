package sumProduct

type Product struct {
	id    int
	title string
	price float32
}

type ProductList []Product

// 接收者为指针类型
func (productList *ProductList) Add(newProduct *Product) {
	// 是对 "指针类型的值" 操作
	*productList = append(*productList, *newProduct)
}

func (product *Product) SetData(title string, price float32) {
	titleInts := []rune(title)
	for _, v := range titleInts {
		product.id = product.id + int(v)
	}

	product.title = title

	product.price = price
}

func (productList *ProductList) Sum() float32 {
	var sum float32
	for _, product := range *productList {
		sum += product.price
	}

	return sum
}
