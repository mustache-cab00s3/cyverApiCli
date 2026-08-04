package services

import "encoding/json"

// FindingTemplatesExportFile describes the downloadable file metadata returned by the non-supported endpoint.
type FindingTemplatesExportFile struct {
	FileName  string `json:"fileName"`
	FileType  string `json:"fileType"`
	FileToken string `json:"fileToken"`
}

// ServiceErrorResponse matches ABP-style error envelopes.
type ServiceErrorResponse struct {
	Code             int         `json:"code"`
	Message          string      `json:"message"`
	Details          interface{} `json:"details"`
	ValidationErrors interface{} `json:"validationErrors"`
}

// FindingTemplatesExportResponse is the full ABP-style response payload for GetFindingTemplatesToExport.
type FindingTemplatesExportResponse struct {
	Result              *FindingTemplatesExportFile `json:"result"`
	TargetURL           interface{}                 `json:"targetUrl"`
	Success             bool                        `json:"success"`
	Error               *ServiceErrorResponse       `json:"error"`
	UnAuthorizedRequest bool                        `json:"unAuthorizedRequest"`
	ABP                 bool                        `json:"__abp"`
}

// UploadFindingEvidenceFilesResponse is the ABP response for
// /App/Projects/UploadFindingEvidenceFiles.
type UploadFindingEvidenceFilesResponse struct {
	Result              []FindingTemplatesExportFile `json:"result"`
	TargetURL           interface{}                  `json:"targetUrl"`
	Success             bool                         `json:"success"`
	Error               *ServiceErrorResponse        `json:"error"`
	UnAuthorizedRequest bool                         `json:"unAuthorizedRequest"`
	ABP                 bool                         `json:"__abp"`
}

// TempFileDownloadRequest contains query parameters for /File/DownloadTempFile.
type TempFileDownloadRequest struct {
	FileType  string
	FileToken string
	FileName  string
}

// TempFileDownloadResult is the parsed output for downloaded file content.
type TempFileDownloadResult struct {
	FileName        string
	FileType        string
	ContentType     string
	ContentLength   int
	ContentEncoding string
	Body            []byte
}

// FindingLibraryDto matches observed GetFindingLibraries item payloads.
type FindingLibraryDto struct {
	GUID                         string  `json:"guid"`
	Name                         string  `json:"name"`
	Description                  string  `json:"description"`
	Status                       string  `json:"status"`
	FindingLibraryTemplateStatus string  `json:"findingLibraryTemplateStatus"`
	FindingFieldsTemplateGUID    *string `json:"findingFieldsTemplateGuid"`
	FindingFieldsTemplateName    *string `json:"findingFieldsTemplateName"`
}

// FindingLibraryListResult captures paginated list payloads returned by GetFindingLibraries.
type FindingLibraryListResult struct {
	Items      []FindingLibraryDto `json:"items"`
	TotalCount int                 `json:"totalCount"`
}

// FindingLibrariesResponse is the ABP envelope for GetFindingLibraries.
type FindingLibrariesResponse struct {
	Result              *FindingLibraryListResult `json:"result"`
	TargetURL           interface{}               `json:"targetUrl"`
	Success             bool                      `json:"success"`
	Error               *ServiceErrorResponse     `json:"error"`
	UnAuthorizedRequest bool                      `json:"unAuthorizedRequest"`
	ABP                 bool                      `json:"__abp"`
}

// FindingTemplateExternalURL captures title/link pairs under externalUrl.
type FindingTemplateExternalURL struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

// FindingTemplateCustomField is intentionally generic because custom field shape varies by instance.
type FindingTemplateCustomField map[string]interface{}

// FindingTemplateDto captures observed fields from GetFindingTemplates items.
// Non-documented payloads can vary by instance; unknown fields are ignored.
type FindingTemplateDto struct {
	GUID                          string                       `json:"guid"`
	Name                          string                       `json:"name"`
	Title                         string                       `json:"title"`
	Status                        string                       `json:"status"`
	ExternalURL                   []FindingTemplateExternalURL `json:"externalUrl"`
	Type                          int                          `json:"type"`
	Description                   string                       `json:"description"`
	Criticality                   int                          `json:"criticality"`
	Impact                        int                          `json:"impact"`
	ImpactDescription             string                       `json:"impactDescription"`
	Likelihood                    int                          `json:"likelihood"`
	LikelihoodDescription         string                       `json:"likelihoodDescription"`
	CVSS40Vector                  *string                      `json:"cvsS40Vector"`
	CVSS40Score                   *float64                     `json:"cvsS40Score"`
	CVSS31Vector                  string                       `json:"cvsS31Vector"`
	CVSS31Score                   *float64                     `json:"cvsS31Score"`
	CVSS30Vector                  string                       `json:"cvsS30Vector"`
	CVSS30Score                   *float64                     `json:"cvsS30Score"`
	CVSS20Vector                  string                       `json:"cvsS20Vector"`
	CVSS20Score                   *float64                     `json:"cvsS20Score"`
	Recommendation                string                       `json:"recommendation"`
	BackgroundInformation         string                       `json:"backgroundInformation"`
	VulnerabilityTypes            interface{}                  `json:"vulnerabilityTypes"`
	FindingCWEs                   []string                     `json:"findingCWEs"`
	FindingCVEs                   []string                     `json:"findingCVEs"`
	FindingMitreAttackTactics     []string                     `json:"findingMitreAttackTactics"`
	FindingMitreAttackTechniques  []string                     `json:"findingMitreAttackTechniques"`
	FindingMitreAttackMitigations []string                     `json:"findingMitreAttackMitigations"`
	ExploitIDs                    []string                     `json:"exploitIds"`
	FindingLibraryGUID            string                       `json:"findingLibraryGuid"`
	FindingLibraryName            string                       `json:"findingLibraryName"`
	ControlTemplatesIDs           []string                     `json:"controlTemplatesIds"`
	MergeTitleRegex1              string                       `json:"mergeTitleRegex1"`
	MergeTitleRegex2              *string                      `json:"mergeTitleRegex2"`
	MergeTitleRegex3              *string                      `json:"mergeTitleRegex3"`
	MergeTitleRegex4              *string                      `json:"mergeTitleRegex4"`
	MergeTitleRegex5              *string                      `json:"mergeTitleRegex5"`
	Labels                        string                       `json:"labels"`
	ComplianceStatus              *string                      `json:"complianceStatus"`
	ComplianceComment             *string                      `json:"complianceComment"`
	CustomFields                  []FindingTemplateCustomField `json:"customFields"`
}

// FindingTemplateListResult captures paginated list payloads returned by GetFindingTemplates.
type FindingTemplateListResult struct {
	Items      []FindingTemplateDto `json:"items"`
	TotalCount int                  `json:"totalCount"`
}

// FindingTemplatesResponse is the ABP envelope for GetFindingTemplates.
type FindingTemplatesResponse struct {
	Result              *FindingTemplateListResult `json:"result"`
	TargetURL           interface{}                `json:"targetUrl"`
	Success             bool                       `json:"success"`
	Error               *ServiceErrorResponse      `json:"error"`
	UnAuthorizedRequest bool                       `json:"unAuthorizedRequest"`
	ABP                 bool                       `json:"__abp"`
}

// CreateOrUpdateFindingTemplateResponse is the ABP envelope returned by
// /api/services/app/FindingLibrary/CreateOrUpdateFindingTemplate.
// Result shape can vary by instance/version, so it is kept generic.
type CreateOrUpdateFindingTemplateResponse struct {
	Result              interface{}           `json:"result"`
	TargetURL           interface{}           `json:"targetUrl"`
	Success             bool                  `json:"success"`
	Error               *ServiceErrorResponse `json:"error"`
	UnAuthorizedRequest bool                  `json:"unAuthorizedRequest"`
	ABP                 bool                  `json:"__abp"`
}

// CreateOrUpdateFindingLibraryResponse is the ABP envelope returned by
// /api/services/app/FindingLibrary/CreateOrUpdateFindingLibrary.
// Result shape can vary by instance/version, so it is kept generic.
type CreateOrUpdateFindingLibraryResponse struct {
	Result              interface{}           `json:"result"`
	TargetURL           interface{}           `json:"targetUrl"`
	Success             bool                  `json:"success"`
	Error               *ServiceErrorResponse `json:"error"`
	UnAuthorizedRequest bool                  `json:"unAuthorizedRequest"`
	ABP                 bool                  `json:"__abp"`
}

// CreateOrEditFindingInstanceResponse is the ABP envelope returned by
// /api/services/app/Finding/CreateOrEditFindingInstance.
type CreateOrEditFindingInstanceResponse = CreateOrUpdateFindingTemplateResponse

// GetFindingInstanceForEditResponse is the ABP envelope returned by
// /api/services/app/Finding/GetFindingInstanceForEdit
// (query: findingId, findingInstanceId, projectId when required).
// Result is the same logical shape as [CreateOrEditFindingInstanceRequest] (PascalCase keys in JSON).
type GetFindingInstanceForEditResponse struct {
	Result              json.RawMessage       `json:"result"`
	TargetURL           interface{}           `json:"targetUrl"`
	Success             bool                  `json:"success"`
	Error               *ServiceErrorResponse `json:"error"`
	UnAuthorizedRequest bool                  `json:"unAuthorizedRequest"`
	ABP                 bool                  `json:"__abp"`
}

// FindingInstanceNewEvidenceFile is one uploaded file in NewEvidenceFiles (camelCase JSON keys).
// It matches FindingTemplatesExportFile; the upload flow returns fileToken values used here.
type FindingInstanceNewEvidenceFile = FindingTemplatesExportFile

// CreateOrEditFindingInstanceRequest is the JSON body for CreateOrEditFindingInstance.
// Pass findingId as a query parameter, not in this struct. Use Guid "00000000-0000-0000-0000-000000000000"
// for a new instance. IsVisibleInReport is sent as the strings "true" or "false", not a JSON boolean.
// Initialize EvidenceFiles and NewEvidenceFiles to non-nil slices when you need JSON [] instead of null.
type CreateOrEditFindingInstanceRequest struct {
	Guid              string                           `json:"Guid"`
	Title             string                           `json:"Title"`
	Asset             string                           `json:"Asset"`
	Location          string                           `json:"Location"`
	Version           string                           `json:"Version"`
	IP                string                           `json:"Ip"`
	Hostname          string                           `json:"Hostname"`
	Port              string                           `json:"Port"`
	Protocol          string                           `json:"Protocol"`
	IssueDetails      string                           `json:"IssueDetails"`
	Reproduce         string                           `json:"Reproduce"`
	Evidence          string                           `json:"Evidence"`
	IsVisibleInReport string                           `json:"IsVisibleInReport"`
	EvidenceFiles     []json.RawMessage                `json:"EvidenceFiles"`
	NewEvidenceFiles  []FindingInstanceNewEvidenceFile `json:"NewEvidenceFiles"`
}
