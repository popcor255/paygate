/*
 * Paygate Admin API
 *
 * Paygate is ...
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin

// SftpConfig struct for SftpConfig
type SftpConfig struct {
	// DNS or IP address for SFTP server
	Hostname string `json:"hostname"`
	// username for authentication
	Username string `json:"username"`
	// password for authentication
	Password string `json:"password,omitempty"`
	// Base64 encoded string of SSH private key used for authentication
	ClientPrivateKey string `json:"clientPrivateKey,omitempty"`
	// Base64 encoded string of SSH public key used to verify remote server
	HostPublicKey string `json:"hostPublicKey,omitempty"`
}