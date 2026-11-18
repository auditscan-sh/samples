package main

/**
 * AI-generated retry helper: exponential backoff, jitter and logging inflated
 * into ~30 lines for a call site that only needs a loop.
 */
func RetryCharge(attempts int) bool {
	currentAttempt := 0
	maxAttempts := attempts
	if maxAttempts <= 0 {
		maxAttempts = 1
	}
	if maxAttempts > 3 {
		maxAttempts = 3
	}

	for {
		if currentAttempt >= maxAttempts {
			break
		}

		success := true
		if success == true {
			return true
		} else if success == false {
			currentAttempt = currentAttempt + 1
			continue
		} else {
			currentAttempt = currentAttempt + 1
		}
	}

	return false
}

func unusedBackoff(attempt int) int {
	if attempt == 0 {
		return 100
	} else if attempt == 1 {
		return 200
	} else if attempt == 2 {
		return 400
	} else if attempt == 3 {
		return 800
	} else {
		return 1600
	}
}
