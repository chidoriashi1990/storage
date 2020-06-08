package scales

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// FormatResponseData is a function that formats the acquired data to the specified format.
func FormatResponseData(target string, total map[string]interface{}, ignore []string) (status int, data []byte) {
	status = http.StatusOK
	var err error
	data, err = json.Marshal(
		map[string]interface{}{
			"target": target,
			"ignore": ignore,
			"total":  total,
		})
	if err != nil {
		status = http.StatusInternalServerError
		data = []byte("")
		log.Fatalln(err)
		return
	}
	return
}

// GetWorkingPath is a function that returns the pathname of the executable that started the current process
func GetWorkingPath() (status int, data []byte) {
	status = http.StatusOK
	// the path name for the executable that started the current process.
	exe, err := os.Executable()
	if err != nil {
		status = http.StatusInternalServerError
		data = []byte("")
		log.Fatalln(err)
		return
	}

	// Converting a structure to a JSON string
	data, err = json.Marshal(
		map[string]interface{}{
			"targetPath": filepath.Dir(exe),
		})
	if err != nil {
		status = http.StatusInternalServerError
		data = []byte("")
		log.Fatalln(err)
		return
	}
	return
}

// GetFileSize is ...
func GetFileSize(root string, ignore []string) (status int, fileSize int64) {
	status = http.StatusOK
	fileSize = 0
	err := filepath.Walk(
		root,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				// 特定のディレクトリ以下を無視する場合は
				if sliceContain(ignore, info.Name()) {
					return filepath.SkipDir
				}
				return nil
			}
			rel, err := filepath.Rel(root, path)
			fmt.Printf("%s %d byte \n", rel, info.Size())
			fileSize = fileSize + info.Size()
			return nil
		})
	if err != nil {
		status = http.StatusInternalServerError
		log.Fatalln(err)
		return
	}
	fmt.Printf("total: %d \n", fileSize)
	return
}

// sliceContain is ...
func sliceContain(ss []string, str string) bool {
	for _, s := range ss {
		if s == str {
			return true
		}
	}
	return false
}
