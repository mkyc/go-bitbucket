/*
 * Bitbucket API
 *
 * Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.
 *
 * API version: 2.0
 * Contact: support@bitbucket.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PipelineSshPublicKey struct {
	// The type of the public key.
	KeyType string `json:"key_type,omitempty"`
	// The base64 encoded public key.
	Key string `json:"key,omitempty"`
	// The MD5 fingerprint of the public key.
	Md5Fingerprint string `json:"md5_fingerprint,omitempty"`
	// The SHA-256 fingerprint of the public key.
	Sha256Fingerprint string `json:"sha256_fingerprint,omitempty"`
	Type_             string `json:"type"`
}
