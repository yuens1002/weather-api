package utils

// Trim takes out the matching char from string
func Trim(t string, c string) string {
	if len(t) > 0 && t[0] == c[0] {
		t = t[1:]
	}
	if len(t) > 0 && t[len(t)-1] == c[0] {
		t = t[:len(t)-1]
	}
	return t
}