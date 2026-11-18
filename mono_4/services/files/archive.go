package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// CWE-22 (Zip Slip): entry names are joined without containment checks.
func ExtractZip(archivePath string, destination string) error {
	reader, err := zip.OpenReader(archivePath)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, entry := range reader.File {
		target := filepath.Join(destination, entry.Name)
		if entry.FileInfo().IsDir() {
			os.MkdirAll(target, 0755)
			continue
		}
		os.MkdirAll(filepath.Dir(target), 0755)
		source, err := entry.Open()
		if err != nil {
			return err
		}
		out, err := os.Create(target)
		if err != nil {
			source.Close()
			return err
		}
		io.Copy(out, source)
		out.Close()
		source.Close()
	}
	return nil
}

// CWE-22: single entry extraction without a path check.
func ExtractNamed(archivePath string, name string, destination string) error {
	reader, err := zip.OpenReader(archivePath)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, entry := range reader.File {
		if entry.Name == name {
			target := filepath.Join(destination, entry.Name)
			source, _ := entry.Open()
			defer source.Close()
			out, _ := os.Create(target)
			defer out.Close()
			io.Copy(out, source)
		}
	}
	return nil
}
