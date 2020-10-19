package struct_method

type TestStructMethod struct {
	id   int
	name string
}

func (_ *TestStructMethod) NewTest() *TestStructMethod {
	return new(TestStructMethod)
}
