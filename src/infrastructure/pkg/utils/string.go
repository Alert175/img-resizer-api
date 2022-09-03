package utils

import "regexp"

func ConvertToPathFormat(arg string) (string, error) {
	reg, err := regexp.Compile(`[^\w]`)
	if err != nil {
		return "", err
	}
	return reg.ReplaceAllString(arg, ""), nil
}
