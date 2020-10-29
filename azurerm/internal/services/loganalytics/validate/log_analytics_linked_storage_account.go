package validate

import (
	"fmt"
	"regexp"
)

func LogAnalyticsLinkedStorageAccountWorkspaceName(i interface{}, k string) (warnings []string, errors []error) {
	return logAnalyticsGenericName(i, k)
}

func LogAnalyticsLinkedStorageAccountName(i interface{}, k string) (warnings []string, errors []error) {
	return logAnalyticsGenericName(i, k)
}

func logAnalyticsGenericName(i interface{}, k string) (warnings []string, errors []error) {
	v, ok := i.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected type of %s to be string", k))
		return
	}
	if len(v) < 4 {
		errors = append(errors, fmt.Errorf("length should be greater than %d", 4))
		return
	}
	if len(v) > 63 {
		errors = append(errors, fmt.Errorf("length should be less than %d", 63))
		return
	}
	if !regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$`).MatchString(v) {
		errors = append(errors, fmt.Errorf("expected value of %s does not match regular expression, got %v", k, v))
		return
	}
	return
}
