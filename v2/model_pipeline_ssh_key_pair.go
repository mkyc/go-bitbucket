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

type PipelineSshKeyPair struct {
	// The SSH private key. This value will be empty when retrieving the SSH key pair.
	PrivateKey string `json:"private_key,omitempty"`
	// The SSH public key.
	PublicKey string `json:"public_key,omitempty"`
	Type_     string `json:"type"`
}
