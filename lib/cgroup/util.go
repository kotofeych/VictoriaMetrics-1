package cgroup

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func grepFirstMatch(sourcePath string, match string, index int, delimiter string) (string, error) {
	f, err := os.Open(sourcePath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		if !strings.Contains(text, match) {
			continue
		}
		split := strings.Split(text, delimiter)
		if len(split) < index {
			return "", fmt.Errorf("needed index line: %d, wasn't found at line: %q at file: %q", index, text, sourcePath)
		}
		return strings.TrimSpace(split[index]), nil
	}
	return "", fmt.Errorf("stat: %q, wasn't found at file: %q", match, sourcePath)
}

func readInt64(path string) (int64, error) {
	data, err := ioutil.ReadFile(path)
	if err == nil {
		data = bytes.TrimSpace(data)
		return strconv.ParseInt(string(data), 10, 64)
	}
	return 0, err
}
