/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type HasAdminRole struct {
	// 1-has admin, 0-not.
	HasAdminRole int32 `json:"has_admin_role,omitempty"`
}