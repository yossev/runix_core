package executor_test

import (
	"testing"

	"github.com/yossev/runix_core/internal/executor"
)

func TestExecuteCode(t *testing.T) {
	code := "print('Hello, World!')"
	result, err := executor.ExecuteCode(code, "python")
	if err != nil {
		t.Errorf("Execution failed: %v", err)
	}
	if result != "Hello, World!\n" {
		t.Errorf("Expected 'Hello, World!' but got %s", result)
	}
}
