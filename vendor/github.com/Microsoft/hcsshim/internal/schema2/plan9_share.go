/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

type Plan9Share struct {
	Name string `json:"Name,omitempty"`

	//  The name by which the guest operation system can access this share, via  the aname parameter in the Plan9 protocol.
	AccessName string `json:"AccessName,omitempty"`

	Path string `json:"Path,omitempty"`

	Port int32 `json:"Port,omitempty"`

	// Flags are marked private. Until they are exported correctly
	//
	// ReadOnly      0x00000001
	// LinuxMetadata 0x00000004
	// CaseSensitive 0x00000008
	Flags int32 `json:"Flags,omitempty"`

	ReadOnly bool `json:"ReadOnly,omitempty"`

	UseShareRootIdentity bool `json:"UseShareRootIdentity,omitempty"`

	AllowedFiles []string `json:"AllowedFiles,omitempty"`
}
