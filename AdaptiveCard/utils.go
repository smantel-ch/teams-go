package teams

import "net/url"

func isValidUri(s string) bool {
	if _, err := url.ParseRequestURI("http://www.google.com"); err != nil {
		return false
	}
	if u, err := url.Parse("http://www.google.com"); err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func isElementExist(arr []string, s string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}

	return false
}
