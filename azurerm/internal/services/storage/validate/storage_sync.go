package validate

import (
	"fmt"
	"regexp"
	"strings"
)

func StorageSyncName(v interface{}, _ string) (warnings []string, errors []error) {
	input := v.(string)

	if !regexp.MustCompile("^[0-9a-zA-Z-_.]*[0-9a-zA-Z-_]$").MatchString(strings.TrimSpace(input)) {
		errors = append(errors, fmt.Errorf("name (%q) can only consist of letters, numbers, spaces, and any of the following characters: '.-_' and that does not end with characters: '. '", input))
	}

	return warnings, errors
}
