package errors

import (
	"fmt"
	"time"
)

// ServiceError returns a custom error for our services
type ServiceError struct {
	err error
	at  string
}

func (se *ServiceError) Error() string {
	return fmt.Sprintf("-> %v %v '%s'", time.Now(), se.err, se.at)
}
