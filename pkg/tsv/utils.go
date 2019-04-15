package tsv

import "strings"

// Escape replaces control characters by a representation
func Escape(token string) string {
	token = strings.ReplaceAll(token, "\\", "\\\\")
	token = strings.ReplaceAll(token, "\t", "\\t")
	token = strings.ReplaceAll(token, "\r", "\\r")
	token = strings.ReplaceAll(token, "\n", "\\n")
	return token
}

// Unescapes replaces representations by control characters
func Unescape(token string) string {
	token = strings.ReplaceAll(token, "\\n", "\n")
	token = strings.ReplaceAll(token, "\\r", "\r")
	token = strings.ReplaceAll(token, "\\t", "\t")
	token = strings.ReplaceAll(token, "\\\\", "\\")
	return token
}

// normalizeLine removes newlines and carriage returns from the line
func normalizeLine(line string) string {
	n := len(line)
	if n >= 2 && line[n-2] == '\r' && line[n-1] == '\n' {
		line = line[:n-2]
	} else if n >= 1 && line[n-1] == '\n' || line[n-1] == '\r' {
		line = line[:n-1]
	}
	return line
}
