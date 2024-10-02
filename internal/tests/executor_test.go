package executor_test

import (
	"os/exec"
	"reflect"
	"testing"

	"runix/internal/executor"
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

func TestGetCommand(t *testing.T) {
	tests := []struct {
		language string
		code     string
		expected *exec.Cmd
	}{
		{
			language: "python",
			code:     "print('Hello, World!')",
			expected: exec.Command("docker", "run", "--rm", "--network", "none", "--memory", "100m", "--memory-swap", "100m", "--cpus", "0.5", "--pids-limit", "50", "runix-executor", "python", "-c", "print('Hello, World!')"),
		},
		{
			language: "bash",
			code:     "echo 'Hello, World!'",
			expected: exec.Command("docker", "run", "--rm", "--network", "none", "--memory", "100m", "--memory-swap", "100m", "--cpus", "0.5", "--pids-limit", "50", "runix-executor", "bash", "-c", "echo 'Hello, World!'"),
		},
		{
			language: "javascript",
			code:     "console.log('Hello, World!')",
			expected: exec.Command("docker", "run", "--rm", "--network", "none", "--memory", "100m", "--memory-swap", "100m", "--cpus", "0.5", "--pids-limit", "50", "runix-executor", "node", "-e", "console.log('Hello, World!')"),
		},
	}

	for _, test := range tests {
		t.Run(test.language, func(t *testing.T) {
			cmd := executor.GetCommand(test.language, test.code)
			if cmd == nil {
				t.Fatalf("Expected command, got nil")
			}

			if !reflect.DeepEqual(cmd.Args, test.expected.Args) {
				t.Errorf("Expected args %v, got %v", test.expected.Args, cmd.Args)
			}
		})
	}
}
