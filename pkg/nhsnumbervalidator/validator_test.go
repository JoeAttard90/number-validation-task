package nhsnumbervalidator

import (
	"testing"
)

func TestValidator_ValidateNHSNumber(t *testing.T) {
	validator := &NHSValidator{
		nhsNumLength: 10,
	}

	tests := []struct {
		name      string
		nhsNumber int
		wantErr   bool
	}{
		{
			name:      "Validation_Successful_Known_Value_1",
			nhsNumber: 5990128088,
			wantErr:   false,
		},
		{
			name:      "Validation_Successful_Known_Value_2",
			nhsNumber: 1275988113,
			wantErr:   false,
		},
		{
			name:      "Validation_Successful_Known_Value_3",
			nhsNumber: 4536026665,
			wantErr:   false,
		},
		{
			name:      "Validation_Failed_Known_Value_1",
			nhsNumber: 5990128087,
			wantErr:   true,
		},
		{
			name:      "Validation_Failed_Known_Value_2",
			nhsNumber: 4536016660,
			wantErr:   true,
		},
		{
			name:      "Validation_Successful_Generated_Value_1",
			nhsNumber: validator.GenerateValidNumber(),
			wantErr:   false,
		},
		{
			name:      "Validation_Successful_Generated_Value_2",
			nhsNumber: validator.GenerateValidNumber(),
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateNumber(tt.nhsNumber)
			if err != nil && !tt.wantErr {
				t.Errorf("ValidateNumber() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func BenchmarkValidator_ValidateNHSNumber(b *testing.B) {
	validator := &NHSValidator{
		nhsNumLength: 10,
	}
	num := validator.GenerateValidNumber()
	for n := 0; n < b.N; n++ {
		_ = validator.ValidateNumber(num)
	}
}

func BenchmarkValidator_GenerateValidNumber(b *testing.B) {
	validator := &NHSValidator{
		nhsNumLength: 10,
	}
	for n := 0; n < b.N; n++ {
		_ = validator.GenerateValidNumber()
	}
}
