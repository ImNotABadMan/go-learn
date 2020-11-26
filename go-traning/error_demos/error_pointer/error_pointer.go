package error_pointer

type errString struct {
	err string
}

func (errString errString) Error() string {
	return errString.err
}

func NewErrorPointer(error string) error {
	return &errString{error}
}
