package errors

import (
	"fmt"
	"reflect"
	"strings"
)

// ValidationError represents a validation error
type ValidationError struct {
	*CyverError
	Field string      `json:"field"`
	Value interface{} `json:"value,omitempty"`
	Rule  string      `json:"rule"`
}

// NewValidationError creates a new validation error
func NewValidationError(field, rule string, value interface{}, message string) *ValidationError {
	return &ValidationError{
		CyverError: NewCyverError(ErrCodeValidationFailed, message, nil),
		Field:      field,
		Value:      value,
		Rule:       rule,
	}
}

// ValidationRule represents a validation rule
type ValidationRule struct {
	Field    string
	Required bool
	MinLen   int
	MaxLen   int
	Pattern  string
	Custom   func(interface{}) error
}

// Validator provides validation functionality
type Validator struct {
	rules []ValidationRule
}

// NewValidator creates a new validator
func NewValidator() *Validator {
	return &Validator{
		rules: make([]ValidationRule, 0),
	}
}

// AddRule adds a validation rule
func (v *Validator) AddRule(rule ValidationRule) *Validator {
	v.rules = append(v.rules, rule)
	return v
}

// Validate validates a struct against the defined rules
func (v *Validator) Validate(data interface{}) *ErrorCollection {
	collection := &ErrorCollection{}

	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		collection.Add(NewCyverError(ErrCodeUnexpectedType, "validation target must be a struct", nil))
		return collection
	}

	for _, rule := range v.rules {
		field := val.FieldByName(rule.Field)
		if !field.IsValid() {
			collection.Add(NewCyverError(ErrCodeValidationFailed, fmt.Sprintf("field '%s' not found in struct", rule.Field), nil))
			continue
		}

		// Check if field is required
		if rule.Required && isEmpty(field) {
			collection.Add(NewCyverError(ErrCodeValidationFailed, fmt.Sprintf("field '%s' is required", rule.Field), nil))
			continue
		}

		// Skip validation if field is empty and not required
		if isEmpty(field) && !rule.Required {
			continue
		}

		// Validate string fields
		if field.Kind() == reflect.String {
			str := field.String()

			if rule.MinLen > 0 && len(str) < rule.MinLen {
				collection.Add(NewCyverError(ErrCodeValidationFailed, fmt.Sprintf("field '%s' must be at least %d characters", rule.Field, rule.MinLen), nil))
			}

			if rule.MaxLen > 0 && len(str) > rule.MaxLen {
				collection.Add(NewCyverError(ErrCodeValidationFailed, fmt.Sprintf("field '%s' must be at most %d characters", rule.Field, rule.MaxLen), nil))
			}

			if rule.Pattern != "" && !matchesPattern(str, rule.Pattern) {
				collection.Add(NewCyverError(ErrCodeValidationFailed, fmt.Sprintf("field '%s' does not match required pattern", rule.Field), nil))
			}
		}

		// Custom validation
		if rule.Custom != nil {
			if err := rule.Custom(field.Interface()); err != nil {
				collection.Add(NewCyverError(ErrCodeValidationFailed, err.Error(), nil))
			}
		}
	}

	return collection
}

// isEmpty checks if a reflect.Value is empty
func isEmpty(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Slice, reflect.Array, reflect.Map:
		return v.Len() == 0
	default:
		return false
	}
}

// matchesPattern checks if a string matches a pattern (simple implementation)
func matchesPattern(str, pattern string) bool {
	// This is a simple pattern matching implementation
	// In a real application, you might want to use regex
	return strings.Contains(str, pattern)
}

// Common validation rules
var (
	// RequiredString creates a required string validation rule
	RequiredString = func(field string) ValidationRule {
		return ValidationRule{
			Field:    field,
			Required: true,
		}
	}

	// OptionalString creates an optional string validation rule
	OptionalString = func(field string) ValidationRule {
		return ValidationRule{
			Field:    field,
			Required: false,
		}
	}

	// StringWithLength creates a string validation rule with length constraints
	StringWithLength = func(field string, minLen, maxLen int) ValidationRule {
		return ValidationRule{
			Field:    field,
			Required: true,
			MinLen:   minLen,
			MaxLen:   maxLen,
		}
	}

	// Email creates an email validation rule
	Email = func(field string) ValidationRule {
		return ValidationRule{
			Field:    field,
			Required: true,
			Pattern:  "@",
			Custom: func(value interface{}) error {
				str, ok := value.(string)
				if !ok {
					return fmt.Errorf("email must be a string")
				}
				if !strings.Contains(str, "@") || !strings.Contains(str, ".") {
					return fmt.Errorf("invalid email format")
				}
				return nil
			},
		}
	}

	// URL creates a URL validation rule
	URL = func(field string) ValidationRule {
		return ValidationRule{
			Field:    field,
			Required: true,
			Custom: func(value interface{}) error {
				str, ok := value.(string)
				if !ok {
					return fmt.Errorf("URL must be a string")
				}
				if !strings.HasPrefix(str, "http://") && !strings.HasPrefix(str, "https://") {
					return fmt.Errorf("URL must start with http:// or https://")
				}
				return nil
			},
		}
	}
)

// ValidateStruct validates a struct using reflection and common rules
func ValidateStruct(data interface{}) *ErrorCollection {
	validator := NewValidator()

	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		collection := &ErrorCollection{}
		collection.Add(NewCyverError(ErrCodeUnexpectedType, "validation target must be a struct", nil))
		return collection
	}

	// Add validation rules based on struct tags
	for i := 0; i < val.NumField(); i++ {
		fieldType := val.Type().Field(i)

		// Check for validation tags
		if tag := fieldType.Tag.Get("validate"); tag != "" {
			rules := strings.Split(tag, ",")
			rule := ValidationRule{
				Field: fieldType.Name,
			}

			for _, r := range rules {
				switch r {
				case "required":
					rule.Required = true
				case "email":
					rule.Custom = func(value interface{}) error {
						str, ok := value.(string)
						if !ok {
							return fmt.Errorf("email must be a string")
						}
						if !strings.Contains(str, "@") || !strings.Contains(str, ".") {
							return fmt.Errorf("invalid email format")
						}
						return nil
					}
				case "url":
					rule.Custom = func(value interface{}) error {
						str, ok := value.(string)
						if !ok {
							return fmt.Errorf("URL must be a string")
						}
						if !strings.HasPrefix(str, "http://") && !strings.HasPrefix(str, "https://") {
							return fmt.Errorf("URL must start with http:// or https://")
						}
						return nil
					}
				}
			}

			validator.AddRule(rule)
		}
	}

	return validator.Validate(data)
}
