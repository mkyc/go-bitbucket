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

type ProjectBranchingModelBranchTypes struct {
	// The kind of branch.
	Kind string `json:"kind"`
	// The prefix for this branch type. A branch with this prefix will be classified as per `kind`. The prefix must be a valid prefix for a branch and must always exist. It cannot be blank, empty or `null`.
	Prefix string `json:"prefix"`
}
