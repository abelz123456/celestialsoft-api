package helpers

import "regexp"

func IsValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regExp := regexp.MustCompile(pattern)
	return regExp.MatchString(email)
}
