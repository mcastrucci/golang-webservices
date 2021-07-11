/*
	File that will contain utils used by the application
*/

package main

// contains checks if an URL from a string is present in a slice
func ContainsUrl(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
