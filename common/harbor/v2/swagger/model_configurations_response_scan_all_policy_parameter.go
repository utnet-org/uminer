/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// The parameters of the policy, the values are dependent on the type of the policy.
type ConfigurationsResponseScanAllPolicyParameter struct {
	// The offset in seconds of UTC 0 o'clock, only valid when the policy type is \"daily\"
	DailyTime int32 `json:"daily_time,omitempty"`
}