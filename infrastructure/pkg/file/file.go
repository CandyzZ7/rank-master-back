package file

import (
	"bufio"
	"io"
	"os"
	"regexp"
)

// IsFileExists 判断文件是否存在
func IsFileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// SelfPath 获取当前执行文件的路径
func SelfPath() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	return exePath, nil
}

func GrepFile(patten string, filename string) ([]string, error) {
	re, err := regexp.Compile(patten)
	if err != nil {
		return nil, err
	}

	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	var lines []string
	reader := bufio.NewReader(fd)

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}
		if re.MatchString(line) {
			lines = append(lines, line)
		}
		if err == io.EOF {
			break
		}
	}
	return lines, nil
}
