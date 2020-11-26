package error_value

type errString string

func (error errString) Error() string {
	return string(error)
}

func New(error string) error {
	return errString(error)
}
