/*
 * Storage Scales
 *
 * 💪💪 It is an application that calculates the size of files under a directory and returns them in JSON format.
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package scales

// WeightResponse /weight response struct
type WeightResponse struct {
	Ignore []string `json:"ignore,omitempty"`
	Target string   `json:"target,omitempty"`
	Total  *Size    `json:"total,omitempty"`
}
