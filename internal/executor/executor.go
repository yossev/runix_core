package executor

import (
	"bytes"
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func ExecuteCode(code, language string) (string, error) {
	filePath, err := createTempFile(code, language)
	if err != nil {
		return "", fmt.Errorf("failed to create temporary file %s", err)
	}
	// Clean the temp file after execution
	defer func() {
		exec.Command("rm", "-f", filePath).Run()
	}()

	cmd := getCommand(language, filePath)
	if cmd == nil {
		return "", fmt.Errorf("Invalid language %s", language)
	}

	// Run the CMD using bash
	var out, stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("execution failed: %v, stderr: %s", err, stderr.String())
	}

	return out.String(), nil
}

func getCommand(language string, filepath string) *exec.Cmd {
	switch language {
	case "python":
		return exec.Command("python3", filepath)
	case "bash":
		return exec.Command("bash", filepath)
	default:
		return nil
	}
}

func createTempFile(code, language string) (string, error) {
	var extension string
	switch language {
	case "python":
		extension = ".py"
	case "bash":
		extension = ".sh"
	default:
		return "", fmt.Errorf("unsupported language: %s", language)
	}
	filePath := fmt.Sprintf("/tmp/runix_code_%s%s", generateRandomString(), extension)
	err := writeFile(filePath, code)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateRandomString() string {
	b := make([]rune, 5)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func writeFile(filePath, code string) error {
	cmd := exec.Command("bash", "-c", fmt.Sprintf("echo '%s' > %s", escapeSingleQuotes(code), filePath))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}
	return nil
}

// escapeSingleQuotes escapes single quotes in the code for safe Bash execution
func escapeSingleQuotes(code string) string {
	return strings.ReplaceAll(code, "'", "'\"'\"'")
}
