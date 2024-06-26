/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type LdapFailedImportUsers struct {
	// the uid can't add to system.
	LdapUid string `json:"ldap_uid,omitempty"`
	// fail reason.
	Error_ string `json:"error,omitempty"`
}
