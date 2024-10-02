package executor

import (
	"bytes"
	"fmt"
	"os/exec"
)

func ExecuteCode(code, language string) (string, error) {
	cmd := getCommand(language, code)
	if cmd == nil {
		return "", fmt.Errorf("Invalid language %s", language)
	}

	var outputBuffer bytes.Buffer
	cmd.Stdout = &outputBuffer
	cmd.Stderr = &outputBuffer

	err := cmd.Run()
	output := outputBuffer.String()

	if len(output) > 65533 {
		output = output[:65533] + "\n... (output truncated)"
	}

	if err != nil {
		return "", fmt.Errorf("execution failed: %v, output: %s", err, output)
	}

	return output, nil
}

// Init a Docker Image for each execution, clear it after the execution process
func getCommand(language, code string) *exec.Cmd {
	baseCmd := []string{
		"docker", "run", "--rm",
		"--network", "none",
		"--memory", "100m", // Limit memory to 100 MB
		"--memory-swap", "100m", // Disable swap
		"--cpus", "0.5", // Limit to 0.5 CPU cores
		"--pids-limit", "50", // Limit number of processes/threads
		"timeout", "3s",
	}

	switch language {
	case "python":
		return exec.Command(baseCmd[0], append(baseCmd[1:], "runix-executor", "python", "-c", code)...)
	case "bash":
		return exec.Command(baseCmd[0], append(baseCmd[1:], "runix-executor", "bash", "-c", code)...)
	case "javascript":
		return exec.Command(baseCmd[0], append(baseCmd[1:], "runix-executor", "node", "-e", code)...)
	default:
		return nil
	}
}
