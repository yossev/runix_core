package executor

import (
	"bytes"
	"fmt"
	"os/exec"
)

func ExecuteCode(code, language string) (string, error) {
	cmd := GetCommand(language, code)
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
func GetCommand(language, code string) *exec.Cmd {
	baseCmd := []string{
		"docker", "run", "--rm",
		"--network", "none",
		"--memory", "100m", // Limit memory to 100 MB
		"--memory-swap", "100m", // Disable swap
		"--cpus", "0.5", // Limit to 0.5 CPU cores
		"--pids-limit", "50", // Limit number of processes/threads
		"--ulimit", "nproc=50:50", // Limit the number of processes
		"--ulimit", "nofile=256:256", // Limit the number of open files, prevent with resource exhaustion attacks
	}

	switch language {
	case "python":
		return exec.Command(baseCmd[0], append(baseCmd[1:], "runix-executor", "timeout", "3s", "python", "-c", code)...)
	case "bash":
		return exec.Command(baseCmd[0], append(baseCmd[1:], "runix-executor", "timeout", "3s", "bash", "-c", code)...)
	case "javascript":
		return exec.Command(baseCmd[0], append(baseCmd[1:], "runix-executor", "timeout", "3s", "node", "-e", code)...)
	case "cpp":
		cppCode := fmt.Sprintf(`#include <iostream>
int main() {
%s
return 0;
}`, code)
		compileAndRun := fmt.Sprintf(`echo '%s' > /tmp/code.cpp && g++ -o /tmp/code /tmp/code.cpp && /tmp/code`, cppCode)
		return exec.Command(baseCmd[0], append(baseCmd[1:], "runix-executor", "timeout", "5s", "sh", "-c", compileAndRun)...)
	default:
		return nil
	}
}
