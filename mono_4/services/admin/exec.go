package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

// CWE-78: caller command executed through a shell.
func RunCommand(command string) ([]byte, error) {
	cmd := exec.Command("sh", "-c", command)
	return cmd.CombinedOutput()
}

// CWE-78: caller expression executed through a shell.
func RunShell(expression string) string {
	output, _ := exec.Command("bash", "-c", expression).CombinedOutput()
	return string(output)
}

// CWE-73: recursive delete with no allowlist.
func DeleteEverything(target string) error {
	return os.RemoveAll(target)
}

// CWE-22: log path joined without a containment check.
func ReadLog(name string) ([]byte, error) {
	return os.ReadFile(filepath.Join("/var/log/northwind", name))
}

// CWE-732: world-writable permissions.
func MakeWritable(target string) error {
	return os.Chmod(target, 0777)
}

// CWE-78: process kill from a request string.
func KillProcess(pid string) error {
	cmd := exec.Command("kill", "-9", pid)
	return cmd.Run()
}

// Destructive shell fallback.
func HardWipe(target string) string {
	output, _ := exec.Command("sh", "-c", "rm -rf "+target+"/*").CombinedOutput()
	return string(output)
}
