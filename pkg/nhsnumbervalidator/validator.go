package nhsnumbervalidator

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type NHSValidator struct {
	nhsNumLength int
}

// NewNHSValidator creates a new NHSValidator object
func NewNHSValidator(nhsNumLength int) (*NHSValidator, error) {
	return &NHSValidator{
		nhsNumLength: nhsNumLength,
	}, nil
}

// ValidateNumber validates the nhs number for a patient using the modulus 11 algorithm. If the validation fails,
// an error is returned
func (v *NHSValidator) ValidateNumber(num int) error {
	numLength := len(strconv.Itoa(num))
	originalNum := num
	if numLength != v.nhsNumLength {
		return fmt.Errorf("expected a digit of length %d, got digit of length - %d", v.nhsNumLength, numLength)
	}

	checkDigit := num % 10
	num /= 10
	// Multiply each digit by the appropriate weighting factor
	checksum := getMod11CheckDigit(num)

	if checksum != checkDigit {
		return fmt.Errorf("the nhs number provided - %d did not meet nhs number validation criteria", originalNum)
	}

	log.Printf("[info] successfully validated nhs number - %d", originalNum)
	return nil
}

// GenerateValidNumber generates a 10-digit number which passes the nhs number validation criteria
func (v *NHSValidator) GenerateValidNumber() int {
	rand.Seed(time.Now().UnixNano())
	var generatedNHSNumber int

	for {
		// Generate the first 9 digits randomly
		num := rand.Intn(900000000) + 100000000
		originalNum := num

		checkDigit := getMod11CheckDigit(num)
		if checkDigit >= 10 {
			continue
		}
		generatedNHSNumber = originalNum*10 + checkDigit
		break
	}

	log.Printf("[info] successfully generated nhs number - %d", generatedNHSNumber)

	return generatedNHSNumber
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
	if checkDigit == 11 {
		return 0
	}

	return checkDigit
}
