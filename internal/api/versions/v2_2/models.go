package v2_2

// This file contains Go structs generated from the models.json schema
// Generated structs for API v2.2 models

// --- Basic DTOs ---

type AddressDto struct {
	Street          *string `json:"street,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	City            *string `json:"city,omitempty"`
	Country         *string `json:"country,omitempty"`
	StateOrProvince *string `json:"stateOrProvince,omitempty"`
}

type AjaxResponse struct {
	TargetUrl           *string     `json:"targetUrl,omitempty"`
	Success             bool        `json:"success"`
	Error               *ErrorInfo  `json:"error,omitempty"`
	UnAuthorizedRequest bool        `json:"unAuthorizedRequest"`
	Abp                 bool        `json:"__abp"`
	Result              interface{} `json:"result,omitempty"`
}

// --- Enums ---

type ApiRolesEnum int32

const (
	ApiRolesEnum_Client_View_Only         ApiRolesEnum = 0
	ApiRolesEnum_Client_Finding_Only      ApiRolesEnum = 1
	ApiRolesEnum_Client_Project_Only      ApiRolesEnum = 2
	ApiRolesEnum_Client                   ApiRolesEnum = 3
	ApiRolesEnum_Pentester_View_Only      ApiRolesEnum = 4
	ApiRolesEnum_Pentester_Project_Only   ApiRolesEnum = 5
	ApiRolesEnum_Pentester_General        ApiRolesEnum = 6
	ApiRolesEnum_Pentester_ProjectManager ApiRolesEnum = 7
	ApiRolesEnum_Pentester_Manager        ApiRolesEnum = 8
	ApiRolesEnum_Pentester_Owner          ApiRolesEnum = 9
)

type AssetEnvironmentEnum int32

const (
	AssetEnvironmentEnum_Unknown     AssetEnvironmentEnum = 0
	AssetEnvironmentEnum_Development AssetEnvironmentEnum = 1
	AssetEnvironmentEnum_Test        AssetEnvironmentEnum = 2
	AssetEnvironmentEnum_Production  AssetEnvironmentEnum = 3
)

type AssetHostingTypeEnum int32

const (
	AssetHostingTypeEnum_PublicCloud_Azure       AssetHostingTypeEnum = 1
	AssetHostingTypeEnum_PublicCloud_AWS         AssetHostingTypeEnum = 2
	AssetHostingTypeEnum_PublicCloud_GoogleCloud AssetHostingTypeEnum = 3
	AssetHostingTypeEnum_PublicCloud_Other       AssetHostingTypeEnum = 4
	AssetHostingTypeEnum_OnPremiseInfrastructure AssetHostingTypeEnum = 5
	AssetHostingTypeEnum_PrivateCloud            AssetHostingTypeEnum = 6
	AssetHostingTypeEnum_Other                   AssetHostingTypeEnum = 7
)

type AssetMatchOptionEnum int32

const (
	AssetMatchOptionEnum_CreateAssetAndMatch                       AssetMatchOptionEnum = 0
	AssetMatchOptionEnum_OnlyTakeFindingsThatMatchWithClientAssets AssetMatchOptionEnum = 1
)

type AssetPublicFacingEnum int32

const (
	AssetPublicFacingEnum_Unknown  AssetPublicFacingEnum = 0
	AssetPublicFacingEnum_Internal AssetPublicFacingEnum = 1
	AssetPublicFacingEnum_External AssetPublicFacingEnum = 2
)

type AssetTypeEnum int32

const (
	AssetTypeEnum_WebApplication    AssetTypeEnum = 2
	AssetTypeEnum_MobileApplication AssetTypeEnum = 3
	AssetTypeEnum_API               AssetTypeEnum = 4
	AssetTypeEnum_Network           AssetTypeEnum = 5
	AssetTypeEnum_WifiNetwork       AssetTypeEnum = 6
	AssetTypeEnum_ActiveDirectory   AssetTypeEnum = 7
	AssetTypeEnum_PhysicalAsset     AssetTypeEnum = 8
	AssetTypeEnum_HardwareDevice    AssetTypeEnum = 9
	AssetTypeEnum_SmartContract     AssetTypeEnum = 10
	AssetTypeEnum_SourceCode        AssetTypeEnum = 11
	AssetTypeEnum_User              AssetTypeEnum = 12
	AssetTypeEnum_Other             AssetTypeEnum = 13
)

type AssetTypeHardwareEnum int32

const (
	AssetTypeHardwareEnum_Other       AssetTypeHardwareEnum = 0
	AssetTypeHardwareEnum_Workstation AssetTypeHardwareEnum = 1
	AssetTypeHardwareEnum_Desktop     AssetTypeHardwareEnum = 2
	AssetTypeHardwareEnum_Laptop      AssetTypeHardwareEnum = 3
	AssetTypeHardwareEnum_Printer     AssetTypeHardwareEnum = 4
	AssetTypeHardwareEnum_IotDevice   AssetTypeHardwareEnum = 5
	AssetTypeHardwareEnum_Cloud       AssetTypeHardwareEnum = 6
)

// --- Asset Models ---

type AssetDto struct {
	ID           string                 `json:"id"`
	Domain       *string                `json:"domain,omitempty"`
	IP           *string                `json:"ip,omitempty"`
	Title        *string                `json:"title,omitempty"`
	Description  *string                `json:"description,omitempty"`
	URL          *string                `json:"url,omitempty"`
	Port         *int32                 `json:"port,omitempty"`
	SSID         *string                `json:"ssid,omitempty"`
	Path         *string                `json:"path,omitempty"`
	Protocol     *string                `json:"protocol,omitempty"`
	Location     *string                `json:"location,omitempty"`
	OS           *string                `json:"os,omitempty"`
	Vendor       *string                `json:"vendor,omitempty"`
	Product      *string                `json:"product,omitempty"`
	Version      *string                `json:"version,omitempty"`
	Commit       *string                `json:"commit,omitempty"`
	LinesOfCode  *string                `json:"linesOfCode,omitempty"`
	Repository   *string                `json:"repository,omitempty"`
	Technology   *string                `json:"technology,omitempty"`
	DomainAD     *string                `json:"domainAD,omitempty"`
	PublicFacing *AssetPublicFacingEnum `json:"publicFacing,omitempty"`
	TypeHardware *AssetTypeHardwareEnum `json:"typeHardware,omitempty"`
	Type         *AssetTypeEnum         `json:"type,omitempty"`
	HostingType  *AssetHostingTypeEnum  `json:"hostingType,omitempty"`
	Environment  *AssetEnvironmentEnum  `json:"environment,omitempty"`
}

type AssetDtoPagedResultDto struct {
	Items      []*AssetDto `json:"items,omitempty"`
	TotalCount int32       `json:"totalCount"`
}

type AssetDtoPagedResultDtoAjaxResponse struct {
	TargetUrl           *string                 `json:"targetUrl,omitempty"`
	Success             bool                    `json:"success"`
	Error               *ErrorInfo              `json:"error,omitempty"`
	UnAuthorizedRequest bool                    `json:"unAuthorizedRequest"`
	Abp                 bool                    `json:"__abp"`
	Result              *AssetDtoPagedResultDto `json:"result,omitempty"`
}

// --- Authentication Models ---

type AuthenticateModel struct {
	UserNameOrEmailAddress       string  `json:"userNameOrEmailAddress"`
	Password                     string  `json:"password"`
	TwoFactorVerificationCode    *string `json:"twoFactorVerificationCode,omitempty"`
	RememberClient               bool    `json:"rememberClient"`
	TwoFactorRememberClientToken *string `json:"twoFactorRememberClientToken,omitempty"`
	SingleSignIn                 *bool   `json:"singleSignIn,omitempty"`
	ReturnUrl                    *string `json:"returnUrl,omitempty"`
	CaptchaResponse              *string `json:"captchaResponse,omitempty"`
}

type AuthenticateResultModel struct {
	AccessToken                   *string  `json:"accessToken,omitempty"`
	EncryptedAccessToken          *string  `json:"encryptedAccessToken,omitempty"`
	ExpireInSeconds               int32    `json:"expireInSeconds"`
	ShouldResetPassword           bool     `json:"shouldResetPassword"`
	PasswordResetCode             *string  `json:"passwordResetCode,omitempty"`
	UserId                        string   `json:"userId"`
	RequiresTwoFactorVerification bool     `json:"requiresTwoFactorVerification"`
	TwoFactorAuthProviders        []string `json:"twoFactorAuthProviders,omitempty"`
	TwoFactorRememberClientToken  *string  `json:"twoFactorRememberClientToken,omitempty"`
	ReturnUrl                     *string  `json:"returnUrl,omitempty"`
	RefreshToken                  *string  `json:"refreshToken,omitempty"`
	RefreshTokenExpireInSeconds   int32    `json:"refreshTokenExpireInSeconds"`
}

type AuthenticateResultModelAjaxResponse struct {
	TargetUrl           *string                  `json:"targetUrl,omitempty"`
	Success             bool                     `json:"success"`
	Error               *ErrorInfo               `json:"error,omitempty"`
	UnAuthorizedRequest bool                     `json:"unAuthorizedRequest"`
	Abp                 bool                     `json:"__abp"`
	Result              *AuthenticateResultModel `json:"result,omitempty"`
}

// --- Error Models ---

type ErrorInfo struct {
	Code             int32                 `json:"code"`
	Message          *string               `json:"message,omitempty"`
	Details          *string               `json:"details,omitempty"`
	ValidationErrors []ValidationErrorInfo `json:"validationErrors,omitempty"`
}

type ValidationErrorInfo struct {
	Message *string  `json:"message,omitempty"`
	Members []string `json:"members,omitempty"`
}

// --- Client Models ---

type ClientStatusEnum int32

const (
	ClientStatusEnum_Active   ClientStatusEnum = 1
	ClientStatusEnum_Inactive ClientStatusEnum = 2
)

type ClientDto struct {
	ID          string            `json:"id"`
	Name        *string           `json:"name,omitempty"`
	Description *string           `json:"description,omitempty"`
	Email       *string           `json:"email,omitempty"`
	Phone       *string           `json:"phone,omitempty"`
	Address     *AddressDto       `json:"address,omitempty"`
	Status      *ClientStatusEnum `json:"status,omitempty"`
}

type ClientDtoAjaxResponse struct {
	TargetUrl           *string    `json:"targetUrl,omitempty"`
	Success             bool       `json:"success"`
	Error               *ErrorInfo `json:"error,omitempty"`
	UnAuthorizedRequest bool       `json:"unAuthorizedRequest"`
	Abp                 bool       `json:"__abp"`
	Result              *ClientDto `json:"result,omitempty"`
}

type ClientDtoPagedResultDto struct {
	Items      []*ClientDto `json:"items,omitempty"`
	TotalCount int32        `json:"totalCount"`
}

type ClientDtoPagedResultDtoAjaxResponse struct {
	TargetUrl           *string                  `json:"targetUrl,omitempty"`
	Success             bool                     `json:"success"`
	Error               *ErrorInfo               `json:"error,omitempty"`
	UnAuthorizedRequest bool                     `json:"unAuthorizedRequest"`
	Abp                 bool                     `json:"__abp"`
	Result              *ClientDtoPagedResultDto `json:"result,omitempty"`
}

// --- Project Models ---

type ProjectStatusEnum int32

const (
	ProjectStatusEnum_Draft     ProjectStatusEnum = 1
	ProjectStatusEnum_Active    ProjectStatusEnum = 2
	ProjectStatusEnum_Completed ProjectStatusEnum = 3
	ProjectStatusEnum_Cancelled ProjectStatusEnum = 4
)

type ProjectDtoV2 struct {
	ID          string             `json:"id"`
	Name        *string            `json:"name,omitempty"`
	Description *string            `json:"description,omitempty"`
	Status      *ProjectStatusEnum `json:"status,omitempty"`
	ClientID    *string            `json:"clientId,omitempty"`
	LabelIDs    []string           `json:"labelIds,omitempty"`
	CreatedAt   *string            `json:"creationTime,omitempty"`
	UpdatedAt   *string            `json:"lastModificationTime,omitempty"`
	StartDate   *string            `json:"startDate,omitempty"`
	EndDate     *string            `json:"endDate,omitempty"`
	DueDate     *string            `json:"dueDate,omitempty"`
}

type ProjectDtoV2AjaxResponse struct {
	TargetUrl           *string       `json:"targetUrl,omitempty"`
	Success             bool          `json:"success"`
	Error               *ErrorInfo    `json:"error,omitempty"`
	UnAuthorizedRequest bool          `json:"unAuthorizedRequest"`
	Abp                 bool          `json:"__abp"`
	Result              *ProjectDtoV2 `json:"result,omitempty"`
	RawJSON             []byte        `json:"-"` // Raw JSON response for output handling
}

// GetRawJSON returns the raw JSON response for output handling
func (r *ProjectDtoV2AjaxResponse) GetRawJSON() interface{} {
	return r.RawJSON
}

type ProjectDtoV2PagedResultDto struct {
	Items      []*ProjectDtoV2 `json:"items,omitempty"`
	TotalCount int32           `json:"totalCount"`
}

type ProjectDtoV2PagedResultDtoAjaxResponse struct {
	TargetUrl           *string                     `json:"targetUrl,omitempty"`
	Success             bool                        `json:"success"`
	Error               *ErrorInfo                  `json:"error,omitempty"`
	UnAuthorizedRequest bool                        `json:"unAuthorizedRequest"`
	Abp                 bool                        `json:"__abp"`
	Result              *ProjectDtoV2PagedResultDto `json:"result,omitempty"`
}

// --- Finding Models ---

type FindingSeverityEnum int32

const (
	FindingSeverityEnum_Critical FindingSeverityEnum = 1
	FindingSeverityEnum_High     FindingSeverityEnum = 2
	FindingSeverityEnum_Medium   FindingSeverityEnum = 3
	FindingSeverityEnum_Low      FindingSeverityEnum = 4
	FindingSeverityEnum_Info     FindingSeverityEnum = 5
)

type FindingStatusEnum int32

const (
	FindingStatusEnum_Open     FindingStatusEnum = 1
	FindingStatusEnum_Closed   FindingStatusEnum = 2
	FindingStatusEnum_Resolved FindingStatusEnum = 3
	FindingStatusEnum_Accepted FindingStatusEnum = 4
)

type FindingTypeEnum int32

const (
	FindingTypeEnum_Vulnerability FindingTypeEnum = 1
	FindingTypeEnum_Nonconformity FindingTypeEnum = 2
	FindingTypeEnum_Observation   FindingTypeEnum = 4
	FindingTypeEnum_Incident      FindingTypeEnum = 8
	FindingTypeEnum_Risk          FindingTypeEnum = 16
)

type FindingCriticalityEnum int32

const (
	FindingCriticalityEnum_Info     FindingCriticalityEnum = 0
	FindingCriticalityEnum_Low      FindingCriticalityEnum = 1
	FindingCriticalityEnum_Medium   FindingCriticalityEnum = 2
	FindingCriticalityEnum_High     FindingCriticalityEnum = 3
	FindingCriticalityEnum_Critical FindingCriticalityEnum = 4
)

type FindingOccurrenceEnum int32

const (
	FindingOccurrenceEnum_New          FindingOccurrenceEnum = 1
	FindingOccurrenceEnum_Reoccurrence FindingOccurrenceEnum = 2
)

type RunStatusEnum int32

const (
	RunStatusEnum_Draft      RunStatusEnum = 0
	RunStatusEnum_Completed  RunStatusEnum = 1
	RunStatusEnum_Processing RunStatusEnum = 2
	RunStatusEnum_Failed     RunStatusEnum = 3
	RunStatusEnum_Scanning   RunStatusEnum = 4
)

type FindingPciComplianceEnum int32

const (
	FindingPciComplianceEnum_Pass FindingPciComplianceEnum = 0
	FindingPciComplianceEnum_Fail FindingPciComplianceEnum = 1
)

// --- Additional DTOs for FindingDto ---

type FindingCvssDto struct {
	Cvss40Vector *string  `json:"cvss40Vector,omitempty"`
	Cvss40Score  *float64 `json:"cvss40Score,omitempty"`
	Cvss31Vector *string  `json:"cvss31Vector,omitempty"`
	Cvss31Score  *float64 `json:"cvss31Score,omitempty"`
	Cvss30Vector *string  `json:"cvss30Vector,omitempty"`
	Cvss30Score  *float64 `json:"cvss30Score,omitempty"`
	Cvss20Vector *string  `json:"cvss20Vector,omitempty"`
	Cvss20Score  *float64 `json:"cvss20Score,omitempty"`
}

type ProjectTaskDto struct {
	ID          string          `json:"id"`
	Code        *string         `json:"code,omitempty"`
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Status      *string         `json:"status,omitempty"`
	ProjectID   *string         `json:"projectId,omitempty"`
	CreatedAt   *string         `json:"creationTime,omitempty"`
	UpdatedAt   *string         `json:"lastModificationTime,omitempty"`
	ParentTask  *ProjectTaskDto `json:"parentTask,omitempty"`
}

type ProjectControlDto struct {
	ID            string             `json:"id"`
	Code          *string            `json:"code,omitempty"`
	Name          *string            `json:"name,omitempty"`
	Description   *string            `json:"description,omitempty"`
	Status        *string            `json:"status,omitempty"`
	ProjectID     *string            `json:"projectId,omitempty"`
	CreatedAt     *string            `json:"creationTime,omitempty"`
	UpdatedAt     *string            `json:"lastModificationTime,omitempty"`
	ParentControl *ProjectControlDto `json:"parentControl,omitempty"`
}

type LabelTypeEnum int32

const (
	LabelTypeEnum_Project LabelTypeEnum = 0
	LabelTypeEnum_Finding LabelTypeEnum = 1
	LabelTypeEnum_Asset   LabelTypeEnum = 2
	LabelTypeEnum_Client  LabelTypeEnum = 3
)

type LabelDto struct {
	ID   string         `json:"id"`
	Text *string        `json:"text,omitempty"`
	Type *LabelTypeEnum `json:"type,omitempty"`
}

type FindingEvidenceDto struct {
	ID          *string `json:"id,omitempty"`
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	FilePath    *string `json:"filePath,omitempty"`
	FileSize    *int64  `json:"fileSize,omitempty"`
	MimeType    *string `json:"mimeType,omitempty"`
	CreatedAt   *string `json:"creationTime,omitempty"`
}

type FindingRunDto struct {
	RunID     string  `json:"runId"`
	RunNumber int32   `json:"runNumber"`
	RunName   *string `json:"runName,omitempty"`
	Status    *string `json:"status,omitempty"`
	CreatedAt *string `json:"creationTime,omitempty"`
}

type ExternalUrlDto struct {
	Title *string `json:"title,omitempty"`
	Link  *string `json:"link,omitempty"`
}

type FormFieldTypeEnum int32

const (
	FormFieldTypeEnum_Text     FormFieldTypeEnum = 0
	FormFieldTypeEnum_Number   FormFieldTypeEnum = 1
	FormFieldTypeEnum_Date     FormFieldTypeEnum = 2
	FormFieldTypeEnum_Select   FormFieldTypeEnum = 3
	FormFieldTypeEnum_TextArea FormFieldTypeEnum = 4
)

type CustomFindingFieldAPIDto struct {
	Field     *string            `json:"field,omitempty"`
	FieldType *FormFieldTypeEnum `json:"fieldType,omitempty"`
	Value     *string            `json:"value,omitempty"`
}

type FindingDto struct {
	ID                         string                      `json:"id"`
	ParentID                   *string                     `json:"parentId,omitempty"`
	Code                       *string                     `json:"code,omitempty"`
	Name                       *string                     `json:"name,omitempty"`
	Description                *string                     `json:"description,omitempty"`
	JiraIssueKey               *string                     `json:"jiraIssueKey,omitempty"`
	Impact                     *int32                      `json:"impact,omitempty"`
	ImpactDescription          *string                     `json:"impactDescription,omitempty"`
	Likelihood                 *int32                      `json:"likelihood,omitempty"`
	LikelihoodDescription      *string                     `json:"likelihoodDescription,omitempty"`
	Type                       *FindingTypeEnum            `json:"type,omitempty"`
	Status                     *FindingStatusEnum          `json:"status,omitempty"`
	Severity                   *FindingCriticalityEnum     `json:"severity,omitempty"`
	Occurrence                 *FindingOccurrenceEnum      `json:"occurrence,omitempty"`
	RunStatus                  *RunStatusEnum              `json:"runStatus,omitempty"`
	ComplianceStatus           *FindingPciComplianceEnum   `json:"complianceStatus,omitempty"`
	ComplianceComment          *string                     `json:"complianceComment,omitempty"`
	Recommendation             *string                     `json:"recommendation,omitempty"`
	BackgroundInformation      *string                     `json:"backgroundInformation,omitempty"`
	CVSS                       *FindingCvssDto             `json:"cvss,omitempty"`
	ProjectName                *string                     `json:"projectName,omitempty"`
	ProjectID                  *string                     `json:"projectId,omitempty"`
	ClientName                 *string                     `json:"clientName,omitempty"`
	ClientID                   *string                     `json:"clientId,omitempty"`
	ReporterName               *string                     `json:"reporterName,omitempty"`
	ReporterID                 *string                     `json:"reporterId,omitempty"`
	ReviewerName               *string                     `json:"reviewerName,omitempty"`
	ReviewerID                 *string                     `json:"reviewerId,omitempty"`
	ClientAssigneeName         *string                     `json:"clientAssigneeName,omitempty"`
	ClientAssigneeID           *string                     `json:"clientAssigneeId,omitempty"`
	CreatedAt                  *string                     `json:"creationTime,omitempty"`
	VisibleToClient            *string                     `json:"visibleToClient,omitempty"`
	ClosedOn                   *string                     `json:"closedOn,omitempty"`
	ReviewedOn                 *string                     `json:"reviewedOn,omitempty"`
	ProjectTask                *ProjectTaskDto             `json:"projectTask,omitempty"`
	VulnerabilityTypeList      []string                    `json:"vulnerabilityTypeList,omitempty"`
	CWEList                    []string                    `json:"cweList,omitempty"`
	CVEList                    []string                    `json:"cveList,omitempty"`
	MitreAttackTacticsList     []string                    `json:"mitreAttackTacticsList,omitempty"`
	MitreAttackTechniquesList  []string                    `json:"mitreAttackTechniquesList,omitempty"`
	MitreAttackMitigationsList []string                    `json:"mitreAttackMitigationsList,omitempty"`
	AssetIDList                []string                    `json:"assetIdList,omitempty"`
	ReoccurrenceIDList         []string                    `json:"reoccurrenceIdList,omitempty"`
	ProjectControlList         []*ProjectControlDto        `json:"projectControlList,omitempty"`
	LabelList                  []*LabelDto                 `json:"labelList,omitempty"`
	FindingEvidenceList        []*FindingEvidenceDto       `json:"findingEvidenceList,omitempty"`
	RunList                    []*FindingRunDto            `json:"runList,omitempty"`
	ExternalUrlList            []*ExternalUrlDto           `json:"externalUrlList,omitempty"`
	CustomFields               []*CustomFindingFieldAPIDto `json:"customFields,omitempty"`
	ExternalUrlJSON            *string                     `json:"externalUrlJson,omitempty"`
	CustomFieldsJSON           *string                     `json:"customFieldsJson,omitempty"`
}

type FindingDtoAjaxResponse struct {
	TargetUrl           *string     `json:"targetUrl,omitempty"`
	Success             bool        `json:"success"`
	Error               *ErrorInfo  `json:"error,omitempty"`
	UnAuthorizedRequest bool        `json:"unAuthorizedRequest"`
	Abp                 bool        `json:"__abp"`
	Result              *FindingDto `json:"result,omitempty"`
}

type FindingDtoPagedResultDto struct {
	Items      []*FindingDto `json:"items,omitempty"`
	TotalCount int32         `json:"totalCount"`
}

type FindingDtoPagedResultDtoAjaxResponse struct {
	TargetUrl           *string                   `json:"targetUrl"`
	Success             bool                      `json:"success"`
	Error               *ErrorInfo                `json:"error,omitempty"`
	UnAuthorizedRequest bool                      `json:"unAuthorizedRequest"`
	Abp                 bool                      `json:"__abp"`
	Result              *FindingDtoPagedResultDto `json:"result,omitempty"`
	RawJSON             interface{}               `json:"-"` // Raw JSON response for full output
}

// GetRawJSON returns the raw JSON response
func (r *FindingDtoPagedResultDtoAjaxResponse) GetRawJSON() interface{} {
	return r.RawJSON
}

// --- User Models ---

type UserDto struct {
	ID          string  `json:"id"`
	Name        *string `json:"name,omitempty"`
	Email       *string `json:"email,omitempty"`
	Description *string `json:"description,omitempty"`
}

type UserDtoAjaxResponse struct {
	TargetUrl           *string    `json:"targetUrl,omitempty"`
	Success             bool       `json:"success"`
	Error               *ErrorInfo `json:"error,omitempty"`
	UnAuthorizedRequest bool       `json:"unAuthorizedRequest"`
	Abp                 bool       `json:"__abp"`
	Result              *UserDto   `json:"result,omitempty"`
}

type UserDtoPagedResultDto struct {
	Items      []*UserDto `json:"items,omitempty"`
	TotalCount int32      `json:"totalCount"`
}

type UserDtoPagedResultDtoAjaxResponse struct {
	TargetUrl           *string                `json:"targetUrl,omitempty"`
	Success             bool                   `json:"success"`
	Error               *ErrorInfo             `json:"error,omitempty"`
	UnAuthorizedRequest bool                   `json:"unAuthorizedRequest"`
	Abp                 bool                   `json:"__abp"`
	Result              *UserDtoPagedResultDto `json:"result,omitempty"`
}

// --- Request Models ---

type CreateClientRequest struct {
	Name        string      `json:"name"`
	Description *string     `json:"description,omitempty"`
	Email       *string     `json:"email,omitempty"`
	Phone       *string     `json:"phone,omitempty"`
	Address     *AddressDto `json:"address,omitempty"`
}

type CreateProjectRequestV2 struct {
	Name        string   `json:"name"`
	Description *string  `json:"description,omitempty"`
	ClientID    string   `json:"clientId"`
	LabelIDs    []string `json:"labelIds,omitempty"`
}

type CreateOrUpdateAssetRequest struct {
	Domain       *string                `json:"domain,omitempty"`
	IP           *string                `json:"ip,omitempty"`
	Title        *string                `json:"title,omitempty"`
	Description  *string                `json:"description,omitempty"`
	URL          *string                `json:"url,omitempty"`
	Port         *int32                 `json:"port,omitempty"`
	SSID         *string                `json:"ssid,omitempty"`
	Path         *string                `json:"path,omitempty"`
	Protocol     *string                `json:"protocol,omitempty"`
	Location     *string                `json:"location,omitempty"`
	OS           *string                `json:"os,omitempty"`
	Vendor       *string                `json:"vendor,omitempty"`
	Product      *string                `json:"product,omitempty"`
	Version      *string                `json:"version,omitempty"`
	Commit       *string                `json:"commit,omitempty"`
	LinesOfCode  *string                `json:"linesOfCode,omitempty"`
	Repository   *string                `json:"repository,omitempty"`
	Technology   *string                `json:"technology,omitempty"`
	DomainAD     *string                `json:"domainAD,omitempty"`
	PublicFacing *AssetPublicFacingEnum `json:"publicFacing,omitempty"`
	TypeHardware *AssetTypeHardwareEnum `json:"typeHardware,omitempty"`
	Type         *AssetTypeEnum         `json:"type,omitempty"`
	HostingType  *AssetHostingTypeEnum  `json:"hostingType,omitempty"`
	Environment  *AssetEnvironmentEnum  `json:"environment,omitempty"`
}

type CreateOrUpdateFindingRequest struct {
	Name        string               `json:"name"`
	Description *string              `json:"description,omitempty"`
	Severity    *FindingSeverityEnum `json:"severity,omitempty"`
	Status      *FindingStatusEnum   `json:"status,omitempty"`
	ProjectID   string               `json:"projectId"`
}

// --- Response Models ---

type GuidAjaxResponse struct {
	TargetUrl           *string    `json:"targetUrl,omitempty"`
	Success             bool       `json:"success"`
	Error               *ErrorInfo `json:"error,omitempty"`
	UnAuthorizedRequest bool       `json:"unAuthorizedRequest"`
	Abp                 bool       `json:"__abp"`
	Result              *string    `json:"result,omitempty"`
}

type Int32AjaxResponse struct {
	TargetUrl           *string    `json:"targetUrl,omitempty"`
	Success             bool       `json:"success"`
	Error               *ErrorInfo `json:"error,omitempty"`
	UnAuthorizedRequest bool       `json:"unAuthorizedRequest"`
	Abp                 bool       `json:"__abp"`
	Result              *int32     `json:"result,omitempty"`
}

// --- Token Models ---

type RefreshTokenResult struct {
	AccessToken     *string `json:"accessToken,omitempty"`
	RefreshToken    *string `json:"refreshToken,omitempty"`
	ExpireInSeconds int32   `json:"expireInSeconds"`
	TokenType       *string `json:"tokenType,omitempty"`
}

type RefreshTokenResultAjaxResponse struct {
	TargetUrl           *string             `json:"targetUrl,omitempty"`
	Success             bool                `json:"success"`
	Error               *ErrorInfo          `json:"error,omitempty"`
	UnAuthorizedRequest bool                `json:"unAuthorizedRequest"`
	Abp                 bool                `json:"__abp"`
	Result              *RefreshTokenResult `json:"result,omitempty"`
}

// --- Two Factor Auth Models ---

type SendTwoFactorAuthCodeModel struct {
	UserId   string  `json:"userId"`
	Provider *string `json:"provider,omitempty"`
}

// --- Legacy Models (for backward compatibility) ---

// These models are kept for backward compatibility with existing code
type RequestProjectFormDto struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type RequestProjectFormDtoPagedResultDtoAjaxResponse struct {
	Success bool                    `json:"success"`
	Data    []RequestProjectFormDto `json:"data"`
	Error   string                  `json:"error"`
}

type ContinuousProjectDto struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type ContinuousProjectDtoPagedResultDtoAjaxResponse struct {
	Success bool                   `json:"success"`
	Data    []ContinuousProjectDto `json:"data"`
	Error   string                 `json:"error"`
}

type ContinuousProjectDtoAjaxResponse struct {
	Success bool                 `json:"success"`
	Data    ContinuousProjectDto `json:"data"`
	Error   string               `json:"error"`
}

type RequestProjectFormDtoAjaxResponse struct {
	Success bool                    `json:"success"`
	Data    []RequestProjectFormDto `json:"data"`
	Error   string                  `json:"error"`
}

// Legacy response models with different structure for backward compatibility
type UserDtoPagedResultDtoAjaxResponseLegacy struct {
	Success bool      `json:"success"`
	Data    []UserDto `json:"data"`
	Error   string    `json:"error"`
}

type AssetDtoPagedResultDtoAjaxResponseLegacy struct {
	Success bool       `json:"success"`
	Data    []AssetDto `json:"data"`
	Error   string     `json:"error"`
}

// --- Pentester Models ---

type PentesterInfoModel struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description"`
}

// --- Project Checklist Models ---

type ProjectChecklistDto struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type ProjectChecklistDtoListResultDtoAjaxResponse struct {
	Success bool                  `json:"success"`
	Data    []ProjectChecklistDto `json:"data"`
	Error   string                `json:"error"`
}

// --- Project Compliance Models ---

type ProjectComplianceNormDto struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type ProjectComplianceNormDtoListResultDtoAjaxResponse struct {
	Success bool                       `json:"success"`
	Data    []ProjectComplianceNormDto `json:"data"`
	Error   string                     `json:"error"`
}

// --- Report Models ---

type ReportVersionDto struct {
	ID        string `json:"id"`
	Version   string `json:"version"`
	CreatedAt string `json:"createdAt"`
	Status    string `json:"status"`
}

type ReportVersionDtoListResultDtoAjaxResponse struct {
	Success bool               `json:"success"`
	Data    []ReportVersionDto `json:"data"`
	Error   string             `json:"error"`
}

type ReportDto struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	Version     string `json:"version"`
	CreatedAt   string `json:"createdAt"`
	PublishedAt string `json:"publishedAt"`
}

type ReportDtoAjaxResponse struct {
	Success bool      `json:"success"`
	Data    ReportDto `json:"data"`
	Error   string    `json:"error"`
}

// --- File Type Enum ---

type FileTypeEnum string

const (
	FileTypeReport     FileTypeEnum = "report"
	FileTypeAttachment FileTypeEnum = "attachment"
)

// --- Update Request Models ---

type UpdateProjectStatusRequestV2 struct {
	Status string `json:"status"`
}

// --- Reduced Auth Model ---

type ReducedAuthenticateModel struct {
	UserNameOrEmailAddress string `json:"userNameOrEmailAddress"`
	Password               string `json:"password"`
}

// --- Additional Models for Existing Code Compatibility ---

// Run Models
type RunDto struct {
	ID          string  `json:"id"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Status      *string `json:"status,omitempty"`
	ProjectID   *string `json:"projectId,omitempty"`
	CreatedAt   *string `json:"creationTime,omitempty"`
	UpdatedAt   *string `json:"lastModificationTime,omitempty"`
}

type RunDtoListResultDto struct {
	Items []*RunDto `json:"items,omitempty"`
}

type RunDtoListResultDtoAjaxResponse struct {
	TargetUrl           *string              `json:"targetUrl,omitempty"`
	Success             bool                 `json:"success"`
	Error               *ErrorInfo           `json:"error,omitempty"`
	UnAuthorizedRequest bool                 `json:"unAuthorizedRequest"`
	Abp                 bool                 `json:"__abp"`
	Result              *RunDtoListResultDto `json:"result,omitempty"`
}

// Project Template Models
type ProjectTemplateDto struct {
	ID          string  `json:"id"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Status      *string `json:"status,omitempty"`
}

type ProjectTemplateDtoPagedResultDto struct {
	Items      []*ProjectTemplateDto `json:"items,omitempty"`
	TotalCount int32                 `json:"totalCount"`
}

type ProjectTemplateDtoPagedResultDtoAjaxResponse struct {
	TargetUrl           *string                           `json:"targetUrl,omitempty"`
	Success             bool                              `json:"success"`
	Error               *ErrorInfo                        `json:"error,omitempty"`
	UnAuthorizedRequest bool                              `json:"unAuthorizedRequest"`
	Abp                 bool                              `json:"__abp"`
	Result              *ProjectTemplateDtoPagedResultDto `json:"result,omitempty"`
}

// Report Template Models
type ReportTemplateDto struct {
	ID          string  `json:"id"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Status      *string `json:"status,omitempty"`
}

type ReportTemplateDtoPagedResultDto struct {
	Items      []*ReportTemplateDto `json:"items,omitempty"`
	TotalCount int32                `json:"totalCount"`
}

type ReportTemplateDtoPagedResultDtoAjaxResponse struct {
	TargetUrl           *string                          `json:"targetUrl,omitempty"`
	Success             bool                             `json:"success"`
	Error               *ErrorInfo                       `json:"error,omitempty"`
	UnAuthorizedRequest bool                             `json:"unAuthorizedRequest"`
	Abp                 bool                             `json:"__abp"`
	Result              *ReportTemplateDtoPagedResultDto `json:"result,omitempty"`
}

// Checklist Status Enum
type ChecklistStatusEnum int32

const (
	ChecklistStatusEnum_Draft     ChecklistStatusEnum = 1
	ChecklistStatusEnum_Published ChecklistStatusEnum = 2
)

// Task Group Template Models
type TaskGroupTemplateDto struct {
	ID          string             `json:"id"`
	Name        *string            `json:"name,omitempty"`
	Description *string            `json:"description,omitempty"`
	TaskList    []*TaskTemplateDto `json:"taskList,omitempty"`
}

type TaskTemplateDto struct {
	ID          string  `json:"id"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// Checklist Template Models
type ChecklistTemplateDto struct {
	ID            string                  `json:"id"`
	Status        *ChecklistStatusEnum    `json:"status,omitempty"`
	Name          *string                 `json:"name,omitempty"`
	ExternalUrl   *string                 `json:"externalUrl,omitempty"`
	Description   *string                 `json:"description,omitempty"`
	TaskGroupList []*TaskGroupTemplateDto `json:"taskGroupList,omitempty"`
}

type ChecklistTemplateDtoPagedResultDto struct {
	Items      []*ChecklistTemplateDto `json:"items,omitempty"`
	TotalCount int32                   `json:"totalCount"`
}

type ChecklistTemplateDtoPagedResultDtoAjaxResponse struct {
	TargetUrl           *string                             `json:"targetUrl,omitempty"`
	Success             bool                                `json:"success"`
	Error               *ErrorInfo                          `json:"error,omitempty"`
	UnAuthorizedRequest bool                                `json:"unAuthorizedRequest"`
	Abp                 bool                                `json:"__abp"`
	Result              *ChecklistTemplateDtoPagedResultDto `json:"result,omitempty"`
}

// Compliance Norm Template Models
type ComplianceNormTemplateDto struct {
	ID          string  `json:"id"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Status      *string `json:"status,omitempty"`
}

type ComplianceNormTemplateDtoPagedResultDto struct {
	Items      []*ComplianceNormTemplateDto `json:"items,omitempty"`
	TotalCount int32                        `json:"totalCount"`
}

type ComplianceNormTemplateDtoPagedResultDtoAjaxResponse struct {
	TargetUrl           *string                                  `json:"targetUrl,omitempty"`
	Success             bool                                     `json:"success"`
	Error               *ErrorInfo                               `json:"error,omitempty"`
	UnAuthorizedRequest bool                                     `json:"unAuthorizedRequest"`
	Abp                 bool                                     `json:"__abp"`
	Result              *ComplianceNormTemplateDtoPagedResultDto `json:"result,omitempty"`
}

// Label Models (LabelDto is already defined above)

type LabelDtoPagedResultDto struct {
	Items      []*LabelDto `json:"items,omitempty"`
	TotalCount int32       `json:"totalCount"`
}

type LabelDtoPagedResultDtoAjaxResponse struct {
	TargetUrl           *string                 `json:"targetUrl,omitempty"`
	Success             bool                    `json:"success"`
	Error               *ErrorInfo              `json:"error,omitempty"`
	UnAuthorizedRequest bool                    `json:"unAuthorizedRequest"`
	Abp                 bool                    `json:"__abp"`
	Result              *LabelDtoPagedResultDto `json:"result,omitempty"`
}

// Vulnerability Type Models
type VulnerabilityTypeDto struct {
	ID          string  `json:"id"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type VulnerabilityTypeDtoPagedResultDto struct {
	Items      []*VulnerabilityTypeDto `json:"items,omitempty"`
	TotalCount int32                   `json:"totalCount"`
}

type VulnerabilityTypeDtoPagedResultDtoAjaxResponse struct {
	TargetUrl           *string                             `json:"targetUrl,omitempty"`
	Success             bool                                `json:"success"`
	Error               *ErrorInfo                          `json:"error,omitempty"`
	UnAuthorizedRequest bool                                `json:"unAuthorizedRequest"`
	Abp                 bool                                `json:"__abp"`
	Result              *VulnerabilityTypeDtoPagedResultDto `json:"result,omitempty"`
}

// Team Models
type TeamDto struct {
	ID          string  `json:"id"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Type        *string `json:"type,omitempty"`
}

type TeamDtoAjaxResponse struct {
	TargetUrl           *string    `json:"targetUrl,omitempty"`
	Success             bool       `json:"success"`
	Error               *ErrorInfo `json:"error,omitempty"`
	UnAuthorizedRequest bool       `json:"unAuthorizedRequest"`
	Abp                 bool       `json:"__abp"`
	Result              *TeamDto   `json:"result,omitempty"`
}

type TeamDtoPagedResultDto struct {
	Items      []*TeamDto `json:"items,omitempty"`
	TotalCount int32      `json:"totalCount"`
}

type TeamDtoPagedResultDtoAjaxResponse struct {
	TargetUrl           *string                `json:"targetUrl,omitempty"`
	Success             bool                   `json:"success"`
	Error               *ErrorInfo             `json:"error,omitempty"`
	UnAuthorizedRequest bool                   `json:"unAuthorizedRequest"`
	Abp                 bool                   `json:"__abp"`
	Result              *TeamDtoPagedResultDto `json:"result,omitempty"`
}
