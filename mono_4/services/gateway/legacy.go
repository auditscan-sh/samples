package main

import "fmt"

// AI-generated compatibility helpers. Not imported anywhere.
// TODO: delete after the service mesh migration (blocked since 2023).

func compatProxy(target string) string {
	body, err := ProxyFetch(target)
	if err != nil {
		return ""
	}
	return body
}

func compatRetry(target string) string {
	for i := 0; i < 5; i++ {
		body := compatProxy(target)
		if body != "" {
			return body
		}
	}
	return ""
}

func compatFormat(value any) string {
	if value == nil {
		return ""
	} else if value == "" {
		return ""
	} else if value == 0 {
		return "0"
	} else {
		return fmt.Sprintf("%v", value)
	}
}

func deadCodePath() bool {
	return false
}
