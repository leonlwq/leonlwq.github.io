package utils

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func listFiles(dir string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

func TestGetPrivateIP(t *testing.T) {
	os.Chdir("/Users/lwq/Desktop/")
	files, err := listFiles(".")
	if err != nil {
		log.Fatal(err)
	}
	reg, _ := regexp.Compile(`.ovpn*`)
	fmt.Println("文件列表：")
	for _, file := range files {
		if reg.MatchString(file) {
			fmt.Println(file)
		}
	}
}

func TestCheckIfFileExists(t *testing.T) {
	reg, _ := regexp.Compile(`[0-9]+.[0-9]$`)
	entries, err := os.ReadDir("/tmp/")
	if err != nil {
		log.Fatalf("读取 /tmp/ 失败: %v", err)
	}

	for _, e := range entries {
		if e.IsDir() {
			dirName := e.Name()
			if reg.MatchString(dirName) {
				fmt.Println(strings.Split(dirName, ".")[0])
				break
			}
		}
	}
}

func TestA(t *testing.T) {
}
