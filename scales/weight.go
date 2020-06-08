package scales

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// GetWorkingPath is a function that returns the pathname of the executable that started the current process
func GetWorkingPath() (status int, data []byte) {
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
