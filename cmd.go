package cmdify

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// RunCommand runs a shell command and returns the output or an error.
func RunCompinedCommand(fmtCmd string) (string, error) {
    var out bytes.Buffer
    var stderr bytes.Buffer


	shells := DetectShells()

	if len(shells) < 0 {
		return "", fmt.Errorf("no shell detected")
	}

    cmd := exec.Command(shells[0], "-c", fmtCmd)
    cmd.Stdout = &out
    cmd.Stderr = &stderr

    err := cmd.Run()
    if err != nil {
        return "", fmt.Errorf("command failed: %s, stderr: %s", err, stderr.String())
    }

    return out.String(), nil
}


func RunCommand(cmdName string, args ...string) (string, error) {
    var out bytes.Buffer
    var stderr bytes.Buffer



    cmd := exec.Command(cmdName, args...)
    cmd.Stdout = &out
    cmd.Stderr = &stderr

    err := cmd.Run()
    if err != nil {
        return "", fmt.Errorf("command failed: %s, stderr: %s", err, stderr.String())
    }

    return out.String(), nil
}


// Ls runs the 'ls' command and returns a slice of filenames.
func Ls(args ...string) ([]string, error) {
    output, err := RunCommand("ls", args...)
    if err != nil {
        return nil, err
    }
    
    // Split the output by newlines to get file/directory names
    files := strings.Split(strings.TrimSpace(output), "\n")
    return files, nil
}

// Pwd runs the 'pwd' command and returns the current directory path.
func Pwd() (string, error) {
    output, err := RunCommand("pwd")
    if err != nil {
        return "", err
    }

    // Remove any trailing newline
    return strings.TrimSpace(output), nil
}

// Host runs the 'host' command to resolve a domain name and returns the resolved addresses.
func Host(domain string) ([]string, error) {
    output, err := RunCommand(fmt.Sprintf("host %s", domain))
    if err != nil {
        return nil, err
    }
    
    // Split the output by newlines to get each resolution result
    results := strings.Split(strings.TrimSpace(output), "\n")
    return results, nil
}

// Mkdir runs the 'mkdir' command to create a directory and returns an error if it fails.
func Mkdir(dirName string) error {
    _, err := RunCommand("mkdir", dirName)
    return err
}

// Touch runs the 'touch' command to create a file and returns an error if it fails.
func Touch(fileName string) error {
    _, err := RunCommand("touch", fileName)
    return err
}


// DetectShells checks environment variables to detect active shells like bash, zsh, and sh.
func DetectShells() []string {
    shells := []string{}

    // Check the SHELL environment variable (typically contains the user's default shell)
    shellPath := os.Getenv("SHELL")
    if shellPath != "" {
        parts := strings.Split(shellPath, "/")
        if len(parts) > 0 {
            shells = append(shells, parts[len(parts)-1])
        }
    }

    // Check if BASH is active via BASH_VERSION
    if os.Getenv("BASH_VERSION") != "" {
        shells = append(shells, "bash")
    }

    // Check if ZSH is active via ZSH_VERSION
    if os.Getenv("ZSH_VERSION") != "" {
        shells = append(shells, "zsh")
    }

    // Optionally check other shells based on their environment variables
    if os.Getenv("SHLVL") != "" {
        shells = append(shells, "sh")
    }

    return shells
}
