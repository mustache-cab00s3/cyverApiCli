package errors

import (
	"context"
	"fmt"
	"math"
	"time"
)

// RetryConfig represents configuration for retry behavior
type RetryConfig struct {
	MaxAttempts int           `json:"max_attempts"`
	BaseDelay   time.Duration `json:"base_delay"`
	MaxDelay    time.Duration `json:"max_delay"`
	Multiplier  float64       `json:"multiplier"`
	Jitter      bool          `json:"jitter"`
}

// DefaultRetryConfig returns a default retry configuration
func DefaultRetryConfig() *RetryConfig {
	return &RetryConfig{
		MaxAttempts: 3,
		BaseDelay:   time.Second,
		MaxDelay:    30 * time.Second,
		Multiplier:  2.0,
		Jitter:      true,
	}
}

// RetryableFunc represents a function that can be retried
type RetryableFunc func() error

// Retry executes a function with retry logic
func Retry(ctx context.Context, fn RetryableFunc, config *RetryConfig) error {
	if config == nil {
		config = DefaultRetryConfig()
	}

	var lastErr error
	attempt := 0

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("retry cancelled: %w", ctx.Err())
		default:
		}

		attempt++
		err := fn()

		if err == nil {
			return nil
		}

		lastErr = err

		// Check if error is retryable
		if !IsRetryable(err) {
			return err
		}

		// Check if we've exceeded max attempts
		if attempt >= config.MaxAttempts {
			return fmt.Errorf("max retry attempts (%d) exceeded, last error: %w", config.MaxAttempts, lastErr)
		}

		// Calculate delay with exponential backoff
		delay := calculateDelay(attempt, config)

		// Wait before retry
		select {
		case <-ctx.Done():
			return fmt.Errorf("retry cancelled: %w", ctx.Err())
		case <-time.After(delay):
			// Continue to next attempt
		}
	}
}

// RetryWithResult executes a function with retry logic and returns a result
func RetryWithResult[T any](ctx context.Context, fn func() (T, error), config *RetryConfig) (T, error) {
	var zero T

	if config == nil {
		config = DefaultRetryConfig()
	}

	var lastErr error
	attempt := 0

	for {
		select {
		case <-ctx.Done():
			return zero, fmt.Errorf("retry cancelled: %w", ctx.Err())
		default:
		}

		attempt++
		result, err := fn()

		if err == nil {
			return result, nil
		}

		lastErr = err

		// Check if error is retryable
		if !IsRetryable(err) {
			return zero, err
		}

		// Check if we've exceeded max attempts
		if attempt >= config.MaxAttempts {
			return zero, fmt.Errorf("max retry attempts (%d) exceeded, last error: %w", config.MaxAttempts, lastErr)
		}

		// Calculate delay with exponential backoff
		delay := calculateDelay(attempt, config)

		// Wait before retry
		select {
		case <-ctx.Done():
			return zero, fmt.Errorf("retry cancelled: %w", ctx.Err())
		case <-time.After(delay):
			// Continue to next attempt
		}
	}
}

// calculateDelay calculates the delay for the next retry attempt
func calculateDelay(attempt int, config *RetryConfig) time.Duration {
	// Calculate exponential backoff: baseDelay * (multiplier ^ (attempt - 1))
	delay := float64(config.BaseDelay) * math.Pow(config.Multiplier, float64(attempt-1))

	// Cap at max delay
	if delay > float64(config.MaxDelay) {
		delay = float64(config.MaxDelay)
	}

	// Add jitter if enabled (±25% random variation)
	if config.Jitter {
		jitter := delay * 0.25
		delay = delay - jitter + (jitter * 2 * (float64(time.Now().UnixNano()) / float64(time.Second)))
	}

	return time.Duration(delay)
}

// RetryConfigBuilder provides a fluent interface for building retry configurations
type RetryConfigBuilder struct {
	config *RetryConfig
}

// NewRetryConfigBuilder creates a new retry configuration builder
func NewRetryConfigBuilder() *RetryConfigBuilder {
	return &RetryConfigBuilder{
		config: DefaultRetryConfig(),
	}
}

// WithMaxAttempts sets the maximum number of retry attempts
func (b *RetryConfigBuilder) WithMaxAttempts(attempts int) *RetryConfigBuilder {
	b.config.MaxAttempts = attempts
	return b
}

// WithBaseDelay sets the base delay between retries
func (b *RetryConfigBuilder) WithBaseDelay(delay time.Duration) *RetryConfigBuilder {
	b.config.BaseDelay = delay
	return b
}

// WithMaxDelay sets the maximum delay between retries
func (b *RetryConfigBuilder) WithMaxDelay(delay time.Duration) *RetryConfigBuilder {
	b.config.MaxDelay = delay
	return b
}

// WithMultiplier sets the exponential backoff multiplier
func (b *RetryConfigBuilder) WithMultiplier(multiplier float64) *RetryConfigBuilder {
	b.config.Multiplier = multiplier
	return b
}

// WithJitter enables or disables jitter
func (b *RetryConfigBuilder) WithJitter(enabled bool) *RetryConfigBuilder {
	b.config.Jitter = enabled
	return b
}

// Build returns the configured RetryConfig
func (b *RetryConfigBuilder) Build() *RetryConfig {
	return b.config
}

// RetryableError represents an error that can be retried
type RetryableError struct {
	*CyverError
	RetryAfter time.Duration `json:"retry_after,omitempty"`
}

// NewRetryableError creates a new retryable error
func NewRetryableError(code ErrorCode, message string, retryAfter time.Duration, err error) *RetryableError {
	return &RetryableError{
		CyverError: NewCyverError(code, message, err),
		RetryAfter: retryAfter,
	}
}

// GetRetryAfter returns the suggested retry delay
func (e *RetryableError) GetRetryAfter() time.Duration {
	return e.RetryAfter
}
