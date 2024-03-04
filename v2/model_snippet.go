/*
 * Bitbucket API
 *
 * Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.
 *
 * API version: 2.0
 * Contact: support@bitbucket.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package bitbucket

import (
	"time"
)

type Snippet struct {
	Id    int32  `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	// The DVCS used to store the snippet.
	Scm       string    `json:"scm,omitempty"`
	CreatedOn time.Time `json:"created_on,omitempty"`
	UpdatedOn time.Time `json:"updated_on,omitempty"`
	Owner     *Account  `json:"owner,omitempty"`
	Creator   *Account  `json:"creator,omitempty"`
	IsPrivate bool      `json:"is_private,omitempty"`
	Type_     string    `json:"type"`
}
