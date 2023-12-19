/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type StartReplicationExecution struct {
	// The ID of policy that the execution belongs to.
	PolicyId int64 `json:"policy_id,omitempty"`
}