package main

import (
	"io"
	"net/http"
	"os/exec"
)

// CWE-918 sink reached from PreviewHandler.
func FetchPreview(target string) (string, error) {
	response, err := http.Get(target)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	return string(body), err
}

// CWE-78: file path interpolated into a shell command.
func GenerateThumbnail(file string) (string, error) {
	command := exec.Command("sh", "-c", "convert "+file+" -resize 100x100 /srv/uploads/small.png")
	output, err := command.CombinedOutput()
	return string(output), err
}

// CWE-78: ffmpeg with an unquoted, caller-controlled format.
func ConvertVideo(file string, format string) (string, error) {
	command := exec.Command("sh", "-c", "ffmpeg -i "+file+" -f "+format+" /srv/uploads/out."+format)
	output, err := command.CombinedOutput()
	return string(output), err
}

// CWE-918: fixed metadata endpoint.
func FetchMetadata() (string, error) {
	return FetchPreview("http://169.254.169.254/latest/meta-data/")
}
