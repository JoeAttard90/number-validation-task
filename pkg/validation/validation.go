package validation

type Validator struct {
}

// NewValidator creates a new Validator object
func NewValidator() (*Validator, error) {
	return &Validator{}, nil
}

// ValidateNHSNumber validates the nhs number for a patient using the modulus 11 algorithm. If the validation fails,
// and error is returned
func (v *Validator) ValidateNHSNumber(nhsNumber int) error {
	panic("implement me")
}

func (v *Validator) GenerateValidNHSNumber() int {
	panic("implement me")
}
