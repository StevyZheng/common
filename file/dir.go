package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetCurrentDirectory() (path string, err error) {
	path, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return strings.Replace(path, "\\", "/", -1), err
}

func GetRootDir() (path string, err error) {
	path, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		path = fmt.Sprintf(".%s", string(os.PathSeparator))
	} else {
		path = fmt.Sprintf("%s%s", path, string(os.PathSeparator))
	}
	return path, err
}

func GetExecFilePath() (path string, err error) {
	path, err = exec.LookPath(os.Args[0])
	if err != nil {
		path = fmt.Sprintf(".%s", string(os.PathSeparator))
	} else {
		path, err = filepath.Abs(path)
	}
	return path, err
}

func ListFiles(path string) (files []string, err error) {
	fs, err := ioutil.ReadDir(path)
	for _, file := range fs {
		if !file.IsDir() {
			files = append(files, filepath.Join(path, file.Name()))
		}
	}
	return files, err
}
