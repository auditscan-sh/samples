package main

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// CWE-918 sink: the target is fully caller-controlled.
func ProxyFetch(target string) (string, error) {
	response, err := http.Get(target)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	return string(body), err
}

// CWE-295: certificate verification disabled and redirects followed.
func FetchInsecure(target string) (string, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	response, err := client.Get(target)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	return string(body), nil
}

// Safe false-positive bait: scheme + host allowlist before the request.
func SafeFetch(target string) (string, error) {
	parsed, err := url.Parse(target)
	if err != nil {
		return "", err
	}
	allowed := map[string]bool{"api.northwind-commerce.example": true, "cdn.northwind-commerce.example": true}
	if parsed.Scheme != "https" || !allowed[parsed.Host] {
		return "", nil
	}
	response, err := http.Get(parsed.String())
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	return string(body), nil
}

func ProxyPost(target string, payload string) (string, error) {
	response, err := http.Post(target, "application/json", strings.NewReader(payload))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	return string(body), nil
}

// Reaches the cloud metadata endpoint on purpose.
func MetadataFetch() (string, error) {
	return ProxyFetch("http://169.254.169.254/latest/meta-data/iam/security-credentials/")
}
