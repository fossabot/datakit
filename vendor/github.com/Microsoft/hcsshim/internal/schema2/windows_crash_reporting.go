/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

type WindowsCrashReporting struct {
	DumpFileName string `json:"DumpFileName,omitempty"`

	MaxDumpSize int64 `json:"MaxDumpSize,omitempty"`
}
