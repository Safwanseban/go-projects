package utils

import (
	"fmt"

	"github.com/google/uuid"
)

func ManageResult(result []map[string]any, file, status string, err any) []map[string]any {
	if err != nil {
		err = err.(error).Error()
	}
	result = append(result, map[string]any{
		"file":   file,
		"status": status,
		"error":  err,
	})
	return result
}

// func ManageErrorResult(result []map[string]any, options ...any) {

// 	result = append(result, map[string]any{
// 		"file"
// 	})
// }
func NewUUIDFile(extenstion string) string {

	return fmt.Sprintf("%s.%s", uuid.NewString(), extenstion)
}
