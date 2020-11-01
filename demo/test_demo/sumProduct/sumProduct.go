package sumProduct

type Product struct {
	id int
	title string
	price float32
}

type ProductList []Product

func (productList ProductList) add(newProduct Product)  {
	productList = append(productList, newProduct)
}

func (product *Product) new(title string, price float32)  {
	titleInts := []rune(title)
	for _,v := range titleInts {
		product.id = product.id + int(v)
	}

	product.title = title

	product.price = price
}
