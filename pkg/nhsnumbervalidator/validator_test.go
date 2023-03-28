package nhsnumbervalidator

import (
	"testing"
)

func TestValidator_ValidateNHSNumber(t *testing.T) {
	validator := &NHSValidator{}

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
			if err := validator.ValidateNumber(tt.nhsNumber); (err != nil) != tt.wantErr {
				t.Errorf("ValidateNumber() error = %v, wantErr %validator", err, tt.wantErr)
			}
		})
	}
}
