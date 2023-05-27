package fileutil

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

var appendAllLinesMutex sync.Mutex
var appendAllTextMutex sync.Mutex
var copyMutex sync.Mutex
var copyOverWriteMutex sync.Mutex
var moveMutex sync.Mutex
var readAllLinesMutex sync.Mutex
var readAllTextMutex sync.Mutex
var writeAllBytesMutex sync.Mutex
var writeAllLinesMutex sync.Mutex
var writeAllTextMutex sync.Mutex

func AppendAllLines(fp string, lines []string) error {
	var lineEnding string
	if runtime.GOOS == "windows" {
		lineEnding = "\r\n"
	} else {
		lineEnding = "\n"
	}
	err := os.MkdirAll(filepath.Dir(fp), os.ModePerm)
	if err != nil {
		return err
	}
	appendAllLinesMutex.Lock()
	defer appendAllLinesMutex.Unlock()
	file, err := os.OpenFile(fp, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, line := range lines {
		_, err := file.WriteString(line + lineEnding)
		if err != nil {
			return err
		}
	}
	return nil
}
func AppendAllText(fp string, txt string) error {
	err := os.MkdirAll(filepath.Dir(fp), os.ModePerm)
	if err != nil {
		return err
	}
	appendAllTextMutex.Lock()
	defer appendAllTextMutex.Unlock()
	file, err := os.OpenFile(fp, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(txt)
	if err != nil {
		return err
	}
	return nil
}
func Copy(fp string, dst string) error {
	copyMutex.Lock()
	defer copyMutex.Unlock()

	srcFile, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	err = os.MkdirAll(filepath.Dir(dst), os.ModePerm)
	if err != nil {
		return err
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	err = dstFile.Sync()
	if err != nil {
		return err
	}

	return nil
}

func CopyOverWrite(fp string, dst string) error {
	copyOverWriteMutex.Lock()
	defer copyOverWriteMutex.Unlock()
	err := os.Remove(dst)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	err = Copy(fp, dst)
	if err != nil {
		return err
	}
	return nil
}
func Exists(filepath string) bool {
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	} else if err != nil {
		return false
	}
	return true
}
func Exist2(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
func Delete(fp string) error {
	err := os.Remove(fp)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

func Move(sourcePath, destinationPath string) error {
	moveMutex.Lock()
	defer moveMutex.Unlock()

	source, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer source.Close()

	err = os.MkdirAll(filepath.Dir(destinationPath), os.ModePerm)
	if err != nil {
		return err
	}

	destination, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}

	err = os.Remove(sourcePath)
	if err != nil {
		return err
	}

	return nil
}
func ReadAllBytes(fp string) ([]byte, error) {
	file, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func ReadAllLines(fp string) ([]string, error) {
	readAllLinesMutex.Lock()
	defer readAllLinesMutex.Unlock()

	file, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func ReadAllText(fp string) string {
	readAllTextMutex.Lock()
	defer readAllTextMutex.Unlock()

	file, err := os.Open(fp)
	if err != nil {
		return ""
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return ""
	}

	return string(bytes)
}
func WriteAllBytes(fp string, data []byte) error {
	err := os.MkdirAll(filepath.Dir(fp), os.ModePerm)
	if err != nil {
		return err
	}
	writeAllBytesMutex.Lock()
	defer writeAllBytesMutex.Unlock()
	err = os.WriteFile(fp, data, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
func WriteAllLines(fp string, lines []string) error {
	err := os.MkdirAll(filepath.Dir(fp), os.ModePerm)
	if err != nil {
		return err
	}

	var lineEnding string
	if runtime.GOOS == "windows" {
		lineEnding = "\r\n"
	} else {
		lineEnding = "\n"
	}
	content := strings.Join(lines, lineEnding)
	writeAllLinesMutex.Lock()
	defer writeAllLinesMutex.Unlock()
	err = os.WriteFile(fp, []byte(content), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
func WriteAllText(fp string, txt string) error {
	err := os.MkdirAll(filepath.Dir(fp), os.ModePerm)
	if err != nil {
		return err
	}
	writeAllTextMutex.Lock()
	defer writeAllTextMutex.Unlock()
	err = os.WriteFile(fp, []byte(txt), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
