package validation

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Validator struct {
}

// NewValidator creates a new Validator object
func NewValidator() (*Validator, error) {
	return &Validator{}, nil
}

// ValidateNHSNumber validates the nhs number for a patient using the modulus 11 algorithm. If the validation fails,
// an error is returned
func (v *Validator) ValidateNHSNumber(nhsNumber int) error {
	numLength := len(strconv.Itoa(nhsNumber))
	originalNum := nhsNumber
	if numLength != 10 {
		return fmt.Errorf("expected a digit of length 10, got digit of length - %d", numLength)
	}

	var sum int
	checkDigit := nhsNumber % 10
	nhsNumber /= 10
	// Multiply each digit by the appropriate weighting factor
	for i := 1; i < 10; i++ {
		digit := nhsNumber % 10
		weight := i + 1
		sum += digit * weight
		if nhsNumber >= 10 {
			nhsNumber /= 10
		}
	}
	checksum := 11 - (sum % 11)
	if checksum != checkDigit {
		return fmt.Errorf("the nhs number provided - %d did not meeet the validation criteria", originalNum)
	}
	return nil
}

// GenerateValidNHSNumber generates a 10-digit number which passes the nhs number validation criteria
func (v *Validator) GenerateValidNHSNumber() int {
	rand.Seed(time.Now().UnixNano())

	// Generate the first 9 digits randomly
	num := rand.Intn(900000000) + 100000000
	originalNum := num

	// Compute the last digit such that the number passes the modulus 11 algorithm
	var sum int
	for i := 1; i < 10; i++ {
		digit := num % 10
		weight := i + 1
		sum += digit * weight
		if num >= 10 {
			num /= 10
		}
	}
	remainder := sum % 11
	checksum := 11 - remainder

	return originalNum*10 + checksum
}
