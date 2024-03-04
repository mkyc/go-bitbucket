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

import (
	"time"
)

type Tag struct {
	// The message associated with the tag, if available.
	Message string `json:"message,omitempty"`
	// The date that the tag was created, if available
	Date   time.Time `json:"date,omitempty"`
	Tagger *Author   `json:"tagger,omitempty"`
	Type_  string    `json:"type"`
	Links  *RefLinks `json:"links,omitempty"`
	// The name of the ref.
	Name   string  `json:"name,omitempty"`
	Target *Commit `json:"target,omitempty"`
}
