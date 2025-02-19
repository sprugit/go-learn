package shared

import (
	"bufio"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

const ProjectRoot = "Prog1"

func FindProjectBase() string {
	cwd, _ := os.Getwd()
	dirs := strings.Split(cwd, string(os.PathSeparator))
	return strings.Join(dirs[:slices.Index(dirs, ProjectRoot)+1],
		string(os.PathSeparator))
}

func FindTestResource(resourcePath []string) string {
	return filepath.Join(FindProjectBase(),
		strings.Join(resourcePath, string(os.PathSeparator)))
}

func LoadFileContents(resourcePath []string) ([]string, error) {
	f, err := os.Open(FindTestResource(resourcePath))
	if err != nil {
		return []string{}, err
	}
	var content = make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		content = append(content, strings.Trim(scanner.Text(), " "))
	}
	return content, nil
}

func AssertMatchFileContents(resourcePath []string, result string) (bool, error) {
	expectedContents, err := LoadFileContents(resourcePath)
	if err != nil {
		return false, err
	}
	resultContents := strings.Split(result, "\n")
	for index, expectedLine := range expectedContents {
		resultLine := strings.Trim(resultContents[index], "\n")
		if expectedLine != resultLine {
			return false, nil
		}
	}
	return true, nil
}
