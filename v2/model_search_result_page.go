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

type SearchResultPage struct {
	Size             int64                    `json:"size,omitempty"`
	Page             int32                    `json:"page,omitempty"`
	Pagelen          int32                    `json:"pagelen,omitempty"`
	QuerySubstituted bool                     `json:"query_substituted,omitempty"`
	Next             string                   `json:"next,omitempty"`
	Previous         string                   `json:"previous,omitempty"`
	Values           []SearchCodeSearchResult `json:"values,omitempty"`
}
