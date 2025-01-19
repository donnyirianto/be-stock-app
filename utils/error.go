// modules/sales_wilayah/usecase/error.go

package utils

import (
	"go.uber.org/zap"
)

// LogError logs the error with the given branch, attempt number, message, and error details.
func LogError(branch string, attempt int, message string, err error) {
	GetLogger().Error(
		message,
		zap.String("cabang", branch),
		zap.Int("attempt", attempt),
		zap.Error(err),
	)
}
