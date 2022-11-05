package utils

import (
	"errors"
	"strings"
)

// Получить из названия файла: имя файла, и его расширение
func SeparateFileName(arg string) (string, string, error) {
	splitFileName := strings.Split(arg, ".")
	if len(arg) == 0 || len(splitFileName) < 2 {
		return "", "", errors.New("not valid arg")
	}
	fileName := []string{}
	fileExtention := ""

	for i, v := range splitFileName {
		if i != len(splitFileName)-1 {
			fileName = append(fileName, v)
		} else {
			fileExtention = v
		}
	}

	return strings.Join(fileName, ""), fileExtention, nil
}
