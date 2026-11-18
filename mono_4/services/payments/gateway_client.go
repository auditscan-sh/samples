package main

import (
	"io"
	"net/http"
)

// CWE-918 sink reached from ReceiptHandler.
func FetchReceipt(target string) (string, error) {
	response, err := http.Get(target)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	return string(body), err
}

// CWE-918: fixed metadata endpoint.
func FetchMetadata() (string, error) {
	return FetchReceipt("http://169.254.169.254/latest/meta-data/iam/security-credentials/")
}

// CWE-295: TLS verification disabled.
func FetchInsecure(target string) (string, error) {
	transport := &http.Transport{TLSClientConfig: nil}
	client := &http.Client{Transport: transport}
	response, err := client.Get(target)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	return string(body), nil
}
