/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ReplicationTrigger struct {
	// The replication policy trigger type. The valid values are manual, event_based and scheduled.
	Type_           string                      `json:"type,omitempty"`
	TriggerSettings *ReplicationTriggerSettings `json:"trigger_settings,omitempty"`
}