/*
 * Storage Scales
 *
 * 💪💪 It is an application that calculates the size of files under a directory and returns them in JSON format.
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package scales

import (
	"encoding/json"
	"log"
	"net/http"
)

// GetWeight Aggregate the size of the files under a directory
func GetWeight(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	stat, total := ToConvertBytes(3221225472)

	// response body
	body, err := json.Marshal(
		map[string]interface{}{
			"target": "/base_dir",
			"total":  total,
		})
	if err != nil {
		stat = http.StatusInternalServerError
		body = []byte("")
		log.Fatalln(err)
	}

	w.WriteHeader(stat)
	w.Write(body)
}
