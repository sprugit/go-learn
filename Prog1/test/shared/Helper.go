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

func LoadFileContents(resourcePath []string) (string, error) {
	f, err := os.Open(FindTestResource(resourcePath))
	if err != nil {
		return "", err
	}
	var content string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		content += scanner.Text()
	}
	return content, nil
}

func AssertMatchFileContents(resourcePath []string, result string) (bool, error) {
	contents, err := LoadFileContents(resourcePath)
	if err != nil {
		return false, err
	}
	return result == contents, nil
}
