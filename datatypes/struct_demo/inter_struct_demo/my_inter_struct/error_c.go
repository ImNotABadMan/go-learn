package my_inter_struct

type MyError struct {
	prefix string
}

func (err *MyError) SetPointerPrefix(prefix string) *MyError {
	err.prefix = prefix
	return err
}

func (err MyError) SetStructPrefix(prefix string) MyError {
	err.prefix = prefix
	return err
}

func (err MyError) Error() string {
	if len(err.prefix) == 0 {
		err.prefix = "prefix"
	}

	return err.prefix + " my error"
}
