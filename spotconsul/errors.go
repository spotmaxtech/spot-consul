package spotconsul

import "github.com/pkg/errors"

// manage some consul errors
// type ErrorConsulKeyNotExist struct {
// 	s string
// }
//
// func (err *ErrorConsulKeyNotExist) Error() string {
// 	return err.s
// }

var ErrorConsulKeyNotExist = errors.New("key not exist")