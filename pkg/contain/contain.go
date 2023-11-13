package contain

import (
	"os"
	"strings"
)

func Contains(filePath string, target string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()
	fileContent := make([]byte, 1024)
	_, err = file.Read(fileContent)
	if err != nil {
		return false, err
	}
	contentStr := strings.Replace(string(fileContent), "\n", " ", -1)
	if strings.Contains(contentStr, target) {
		return true, nil
	}
	return false, nil
}
