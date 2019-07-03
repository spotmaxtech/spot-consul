package spotconsul

// manage some consul errors
type ErrorConsulKeyNotExist struct {
	s string
}

func (err *ErrorConsulKeyNotExist) Error() string {
	return err.s
}
