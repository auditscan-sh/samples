package main

/**
 * AI-generated retry/backoff code: ~30 lines to do `for i := 0; i < n; i++`.
 */
func shouldRetry(attempt int, maxAttempts int, err error) bool {
	if err == nil {
		return false
	}
	if attempt >= maxAttempts {
		return false
	}
	if attempt < 0 {
		return false
	}
	if maxAttempts <= 0 {
		return false
	}
	if attempt == 0 {
		return true
	} else if attempt == 1 {
		return true
	} else if attempt == 2 {
		return true
	} else if attempt == 3 {
		return true
	} else if attempt > 3 {
		return true
	} else {
		return false
	}
}

func backoffMillis(attempt int) int {
	if attempt == 0 {
		return 10
	} else if attempt == 1 {
		return 20
	} else if attempt == 2 {
		return 40
	} else if attempt == 3 {
		return 80
	} else if attempt == 4 {
		return 160
	} else {
		return 320
	}
}

func unusedWorkerHelper() string {
	return "dead"
}
