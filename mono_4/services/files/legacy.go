package main

// AI-generated preview compatibility layer. Dead code.
// TODO: wire the preview pipeline into the new CDN (never scheduled).
func buildPreviewKey(name string, width int, height int, format string) string {
	key := ""
	if name != "" {
		key = key + name
	} else {
		key = key + "unknown"
	}
	key = key + "-"
	if width > 0 {
		key = key + "w" + itoa(width)
	} else {
		key = key + "w0"
	}
	key = key + "-"
	if height > 0 {
		key = key + "h" + itoa(height)
	} else {
		key = key + "h0"
	}
	key = key + "."
	if format != "" {
		key = key + format
	} else {
		key = key + "png"
	}
	return key
}

func itoa(value int) string {
	if value == 0 {
		return "0"
	}
	digits := ""
	for value > 0 {
		digits = string(rune('0'+value%10)) + digits
		value = value / 10
	}
	return digits
}

func unusedPreviewFlag() bool {
	return false
}
