package nhsnumbervalidator

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type NHSValidator struct {
}

// NewValidator creates a new NHSValidator object
func NewValidator() (*NHSValidator, error) {
	return &NHSValidator{}, nil
}

// ValidateNumber validates the nhs number for a patient using the modulus 11 algorithm. If the validation fails,
// an error is returned
func (v *NHSValidator) ValidateNumber(num int) error {
	numLength := len(strconv.Itoa(num))
	originalNum := num
	if numLength != 10 {
		return fmt.Errorf("expected a digit of length 10, got digit of length - %d", numLength)
	}

	checkDigit := num % 10
	num /= 10
	// Multiply each digit by the appropriate weighting factor
	checksum := getMod11CheckDigit(num)

	if checksum != checkDigit {
		return fmt.Errorf("the nhs number provided - %d did not meet nhs number validation criteria", originalNum)
	}
	return nil
}

// GenerateValidNumber generates a 10-digit number which passes the nhs number validation criteria
func (v *NHSValidator) GenerateValidNumber() int {
	rand.Seed(time.Now().UnixNano())

	// Generate the first 9 digits randomly
	num := rand.Intn(900000000) + 100000000
	originalNum := num

	checkDigit := getMod11CheckDigit(num)

	return originalNum*10 + checkDigit
}

func getMod11CheckDigit(num int) int {
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
	checkDigit := 11 - remainder
	return checkDigit
}
