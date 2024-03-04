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

type PipelineCache struct {
	// The UUID identifying the pipeline cache.
	Uuid string `json:"uuid,omitempty"`
	// The UUID of the pipeline that created the cache.
	PipelineUuid string `json:"pipeline_uuid,omitempty"`
	// The uuid of the step that created the cache.
	StepUuid string `json:"step_uuid,omitempty"`
	// The name of the cache.
	Name string `json:"name,omitempty"`
	// The key hash of the cache version.
	KeyHash string `json:"key_hash,omitempty"`
	// The path where the cache contents were retrieved from.
	Path string `json:"path,omitempty"`
	// The size of the file containing the archive of the cache.
	FileSizeBytes int32 `json:"file_size_bytes,omitempty"`
	// The timestamp when the cache was created.
	CreatedOn time.Time `json:"created_on,omitempty"`
	Type_     string    `json:"type"`
}
