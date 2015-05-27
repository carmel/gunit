package generate

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GatherAssertions(directory string) (map[string]map[int]string, error) {
	listing, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	assertions := make(map[string]map[int]string)

	for _, fileInfo := range SelectGoTestFiles(listing) {
		filename := fileInfo.Name()
		assertions[filename], err = assertionsInFile(filepath.Join(directory, filename))
		if err != nil {
			return nil, err
		}
	}

	return assertions, nil
}

func assertionsInFile(path string) (map[int]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	assertions := make(map[int]string)

	scanner := bufio.NewScanner(file)
	for line := 1; scanner.Scan(); line++ {
		text := strings.TrimSpace(scanner.Text())
		if strings.Contains(text, ".So(") || strings.Contains(text, ".Ok(") {
			assertions[line] = text
		}
	}

	return assertions, nil
}
