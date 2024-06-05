package tools

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const PROJECT_NAME = "devenv"

func GetProjectPath() string {
	_, file, _, ok := runtime.Caller(0)
	if ok {
		absPath, err := filepath.Abs(file)
		if err == nil {
			projectPath := removeAfter(absPath)
			if projectPath != "" {
				return projectPath + PROJECT_NAME + "/"
			}
		}
	}
	return ""
}

func removeAfter(s string) string {
	parts := strings.Split(s, PROJECT_NAME)
	if len(parts) > 1 {
		return strings.Join(parts[:len(parts)-1], PROJECT_NAME)
	}
	return ""
}

func ReadFile(wd string, funcGetLine func(l string)) error {
	// read-only
	file, err := os.Open(wd)
	if err != nil {
		fmt.Println("Error opening read-only file:", err)
		return err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading line:", err)
			return err
		}
		funcGetLine(line)
	}
	return nil
}

func ReadAllContent(filePath string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home dir:", err)
		return "", err
	}
	file, err := os.Open(fmt.Sprintf("%s/%s", homeDir, filePath))
	if err != nil {
		fmt.Println("Error opening read-only file:", err)
		return "", err
	}
	defer file.Close()

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading line:", err)
		return "", err
	}
	return string(fileContent), nil
}

func AppendStringToFile(targetFile, s string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home dir:", err)
		return err
	}
	file, err := os.OpenFile(
		fmt.Sprintf("%s/%s", homeDir, targetFile),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()
	// Write the new content to the file
	_, err = file.WriteString(s)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}
	// Content appended to file successfully
	return nil
}

func HighlightLine(s string) {
	startTag := "\033[32m"
	endTag := "\033[0m"

	fmt.Println("  ✅ " + startTag + s + endTag)
}

func HighlightErrorLine(s string) {
	startTag := "\033[31m"
	endTag := "\033[0m"

	fmt.Println("  ❌ " + startTag + s + endTag)
}

func CopyDirectory(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		fmt.Println("Error Stat:", err)
		return err
	}

	if srcInfo.IsDir() {
		return copyDirectoryRecursive(src, dst)
	}

	return CopyFile(src, dst)
}

func copyDirectoryRecursive(src, dst string) error {
	err := os.MkdirAll(dst, 0755)
	if err != nil {
		fmt.Println("Error creating destination directory:", err)
		return err
	}
	files, err := ioutil.ReadDir(src)
	if err != nil {
		fmt.Println("Error reading source directory:", err)
		return err
	}
	for _, file := range files {
		srcPath := filepath.Join(src, file.Name())
		dstPath := filepath.Join(dst, file.Name())

		if file.IsDir() {
			err = copyDirectoryRecursive(srcPath, dstPath)
			if err != nil {
				return err
			}
		} else {
			err = CopyFile(srcPath, dstPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func CopyFile(src, dst string) error {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Println("Error reading source file:", err)
		return err
	}
	err = ioutil.WriteFile(dst, data, 0644)
	if err != nil {
		fmt.Println("Error writting to destination file:", err)
		return err
	}
	return nil
}
