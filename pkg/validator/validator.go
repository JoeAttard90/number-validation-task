package validator

type Validator interface {
	// ValidateNumber validates that an integer passes a given validation criteria
	ValidateNumber(num int) error
	// GenerateValidNumber generates an integer that would pass the validation criteria
	GenerateValidNumber() int
}
