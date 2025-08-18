package utils

import (
	"fmt"
	"os"
	"strings"
)

func LoadFilesByPathAndExtension(dir string, ext string) (files []string, err error) {
	dir = strings.TrimRight(dir, "/")
	files = []string{}
	items, err := os.ReadDir(dir)
	if err != nil {
		return files, nil
	}
	for _, item := range items {
		if !strings.Contains(item.Name(), ext) {
			continue
		}
		path := fmt.Sprintf("%s/%s", dir, item.Name())
		files = append(files, path)
	}
	return
}

func MakeDir(path string) (err error) {
	err = os.MkdirAll(path, os.ModePerm)
	return
}
