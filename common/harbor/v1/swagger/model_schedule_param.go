/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ScheduleParam struct {
	// The schedule type. The valid values are daily and weekly.
	Type_ string `json:"type,omitempty"`
	// Optional, only used when the type is weedly. The valid values are 1-7.
	Weekday int32 `json:"weekday,omitempty"`
	// The time offset with the UTC 00:00 in seconds.
	Offtime int64 `json:"offtime,omitempty"`
}
