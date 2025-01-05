package gologram

import "strings"

/**
 * @author  papajuan
 * @date    1/5/2025
 **/

// escapeJSON escapes special characters in JSON strings.
func escapeJSON(s string) string {
	var builder strings.Builder
	for _, c := range s {
		switch c {
		case '"':
			builder.WriteString(`\"`)
		case '\\':
			builder.WriteString(`\\`)
		case '\n':
			builder.WriteString(`\n`)
		case '\r':
			builder.WriteString(`\r`)
		case '\t':
			builder.WriteString(`\t`)
		default:
			builder.WriteRune(c)
		}
	}
	return builder.String()
}
