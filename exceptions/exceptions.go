package exceptions

import "fmt"

type DataValidationError struct {
	Reason	string `json:"reason"`
}

func (e *DataValidationError) Error() string {
	return fmt.Sprintf("Data validation error: %s", e.Reason)
}