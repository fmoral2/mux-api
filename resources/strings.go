package resources

import "strings"

// StringIsEmpty checks if a string is empty ” or if it only has empty spaces
func StringIsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// StringIsNotEmpty checks if a string has at least one non empty character
func StringIsNotEmpty(s string) bool {
	return !StringIsEmpty(s)
}

// StringPointerIsEmpty checks if a string pointer is empty
// or if it only has empty spaces
func StringPointerIsEmpty(s *string) bool {
	return s == nil || strings.TrimSpace(*s) == ""
}

// StringPointerIsNotEmpty checks if a string pointer has at least one non
func StringPointerIsNotEmpty(s *string) bool {
	return !StringPointerIsEmpty(s)
}
