/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type AccessLog struct {
	// The ID of the log entry.
	LogId int32 `json:"log_id,omitempty"`
	// Username of the user in this log entry.
	Username string `json:"username,omitempty"`
	// Name of the repository in this log entry.
	RepoName string `json:"repo_name,omitempty"`
	// Tag of the repository in this log entry.
	RepoTag string `json:"repo_tag,omitempty"`
	// The operation against the repository in this log entry.
	Operation string `json:"operation,omitempty"`
	// The time when this operation is triggered.
	OpTime string `json:"op_time,omitempty"`
}
