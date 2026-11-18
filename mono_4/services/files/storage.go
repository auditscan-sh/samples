package main

import (
	"os"
	"path/filepath"
	"strings"
)

// CWE-22: write outside the upload directory.
func SaveFile(name string, content []byte) (string, error) {
	target := filepath.Join(uploadDir, name)
	return target, os.WriteFile(target, content, 0644)
}

// CWE-22: read outside the upload directory.
func ReadFile(name string) ([]byte, error) {
	return os.ReadFile(filepath.Join(uploadDir, name))
}

// CWE-22/CWE-73: delete outside the upload directory.
func DeleteFile(name string) error {
	return os.Remove(filepath.Join(uploadDir, name))
}

// Safe false-positive bait: resolve + prefix check.
func ReadInsideUploadDir(name string) ([]byte, error) {
	root, err := filepath.Abs(uploadDir)
	if err != nil {
		return nil, err
	}
	target, err := filepath.Abs(filepath.Join(root, name))
	if err != nil {
		return nil, err
	}
	if !strings.HasPrefix(target, root+string(os.PathSeparator)) {
		return nil, os.ErrPermission
	}
	return os.ReadFile(target)
}

// CWE-59: follows a symlink outside the upload root.
func OpenFollowingLinks(name string) (*os.File, error) {
	return os.Open(filepath.Join(uploadDir, name))
}

// Intentional resource leak for the resource-closed invariant.
func LeakHandle(name string) (int, error) {
	file, err := os.Open(filepath.Join(uploadDir, name))
	if err != nil {
		return 0, err
	}
	info, err := file.Stat()
	if err != nil {
		return 0, err
	}
	return int(info.Size()), nil
}
