package utils

import "strings"

func ContainsAsSubstring(slice []string, element string) bool {
	if slice != nil && len(slice) > 0 {
		for _, e := range slice {
			if strings.Contains(element, e) {
				return true
			}
		}
	}
	return false
}

func ContainsEitherSubstring(slice []string, elements []string) bool {
	if slice != nil && len(slice) > 0 && elements != nil && len(elements) > 0 {
		for _, e := range slice {
			for _, substring := range elements {
				if strings.Contains(e, substring) {
					return true
				}
			}
		}
	}
	return false
}
