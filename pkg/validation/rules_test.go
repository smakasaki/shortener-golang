package validation

import (
	"testing"
)

func TestValidateRequired(t *testing.T) {
	tests := []struct {
		name      string
		fieldName string
		value     string
		wantErr   bool
	}{
		{"Field is required", "Name", "", true},
		{"Field is not required", "Name", "John", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := ValidateRequired(tt.fieldName, tt.value)
			err := rule.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateRequired() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateLength(t *testing.T) {
	tests := []struct {
		name      string
		fieldName string
		value     string
		min       int
		max       int
		wantErr   bool
	}{
		{"Too short", "Name", "Jo", 3, 50, true},
		{"Valid length", "Name", "John", 3, 50, false},
		{"Too long", "Name", "ThisNameIsWayTooLongForValidation", 3, 20, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := ValidateLength(tt.fieldName, tt.value, tt.min, tt.max)
			err := rule.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateLength() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name      string
		fieldName string
		value     string
		wantErr   bool
	}{
		{"Valid email", "Email", "test@example.com", false},
		{"Invalid email without domain", "Email", "test@", true},
		{"Invalid email without @", "Email", "testexample.com", true},
		{"Invalid email with special characters", "Email", "test@ex!ample.com", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := ValidateEmail(tt.fieldName, tt.value)
			err := rule.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
