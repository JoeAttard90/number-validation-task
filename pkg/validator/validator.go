package validator

type Validator interface {
	ValidateNumber(num int) error
	GenerateValidNumber() int
}
