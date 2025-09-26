package v2_2

// This file contains Go structs generated from the models.json schema
// Generated structs for API v2.2 models
// Verified structs: 120+ out of ~150+ total structs (verified against JSON schemas on September 26, 2025)
// Updated with Misc.json structures on September 26, 2025

// --- Basic DTOs ---

// Verified against Misc.json schema on September 26, 2025
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

// Verified against Asset.json schema on September 26, 2025
type AssetEnvironmentEnum int32

const (
	AssetEnvironmentEnum_Unknown     AssetEnvironmentEnum = 0
	AssetEnvironmentEnum_Development AssetEnvironmentEnum = 1
	AssetEnvironmentEnum_Test        AssetEnvironmentEnum = 2
	AssetEnvironmentEnum_Production  AssetEnvironmentEnum = 3
)

// Verified against Asset.json schema on September 26, 2025
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

// Verified against Asset.json schema on September 26, 2025
type AssetMatchOptionEnum int32

const (
	AssetMatchOptionEnum_CreateAssetAndMatch                       AssetMatchOptionEnum = 0
	AssetMatchOptionEnum_OnlyTakeFindingsThatMatchWithClientAssets AssetMatchOptionEnum = 1
)

// Verified against Asset.json schema on September 26, 2025
type AssetPublicFacingEnum int32

const (
	AssetPublicFacingEnum_Unknown  AssetPublicFacingEnum = 0
	AssetPublicFacingEnum_Internal AssetPublicFacingEnum = 1
	AssetPublicFacingEnum_External AssetPublicFacingEnum = 2
)

// Verified against Asset.json schema on September 26, 2025
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

// Verified against Asset.json schema on September 26, 2025
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

// Verified against Asset.json schema on September 26, 2025
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

// Verified against Asset.json schema on September 26, 2025
type AssetDtoPagedResultDto struct {
	Items      []*AssetDto `json:"items,omitempty"`
	TotalCount int32       `json:"totalCount"`
}

// Verified against Asset.json schema on September 26, 2025
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

// Verified against Misc.json schema on September 26, 2025
type ClientStatusEnum int32

const (
	ClientStatusEnum_Active   ClientStatusEnum = 0
	ClientStatusEnum_Inactive ClientStatusEnum = 1
	ClientStatusEnum_New      ClientStatusEnum = 2
	ClientStatusEnum_SignUp   ClientStatusEnum = 3
)

// Verified against Misc.json schema on September 26, 2025
type ClientInformationDto struct {
	CompanyName *string     `json:"companyName,omitempty"`
	Website     *string     `json:"website,omitempty"`
	Address     *AddressDto `json:"address,omitempty"`
}

// Verified against Misc.json schema on September 26, 2025
type ClientDto struct {
	ID                string                `json:"id"`
	Name              *string               `json:"name,omitempty"`
	ClientNumber      *string               `json:"clientNumber,omitempty"`
	TeamsAssigned     *string               `json:"teamsAssigned,omitempty"`
	Status            *ClientStatusEnum     `json:"status,omitempty"`
	ClientInformation *ClientInformationDto `json:"clientInformation,omitempty"`
	AssetIDList       []string              `json:"assetIdList,omitempty"`
	ProjectIDList     []string              `json:"projectIdList,omitempty"`
	LabelList         []*LabelDto           `json:"labelList,omitempty"`
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

// Verified against Project.json schema on September 26, 2025
type ProjectFrequencyEnum int32

const (
	ProjectFrequencyEnum_OnceAMonth   ProjectFrequencyEnum = 1
	ProjectFrequencyEnum_OnceAQuarter ProjectFrequencyEnum = 2
	ProjectFrequencyEnum_TwiceAYear   ProjectFrequencyEnum = 3
	ProjectFrequencyEnum_OnceAYear    ProjectFrequencyEnum = 4
	ProjectFrequencyEnum_OneTimeOnly  ProjectFrequencyEnum = 5
	ProjectFrequencyEnum_Other        ProjectFrequencyEnum = 6
	ProjectFrequencyEnum_IDontKnowYet ProjectFrequencyEnum = 7
	ProjectFrequencyEnum_Daily        ProjectFrequencyEnum = 8
	ProjectFrequencyEnum_Weekly       ProjectFrequencyEnum = 9
	ProjectFrequencyEnum_Monthly      ProjectFrequencyEnum = 10
	ProjectFrequencyEnum_Realtime     ProjectFrequencyEnum = 11
)

// Verified against Project.json schema on September 26, 2025
type ProjectObjectivesEnum int32

const (
	ProjectObjectivesEnum_RegulatoryCompliance  ProjectObjectivesEnum = 1
	ProjectObjectivesEnum_ProductLaunch         ProjectObjectivesEnum = 2
	ProjectObjectivesEnum_PassAVendorAssessment ProjectObjectivesEnum = 3
	ProjectObjectivesEnum_FindVulnerabilities   ProjectObjectivesEnum = 4
	ProjectObjectivesEnum_Other                 ProjectObjectivesEnum = 5
)

// Verified against Project.json schema on September 26, 2025
type ProjectStartEnum int32

const (
	ProjectStartEnum_RightAway       ProjectStartEnum = 1
	ProjectStartEnum_TimeIsFlexible  ProjectStartEnum = 2
	ProjectStartEnum_OneMonth        ProjectStartEnum = 3
	ProjectStartEnum_TwoMonthsOrMore ProjectStartEnum = 4
)

// Verified against Project.json schema on September 26, 2025
type ProjectTaskStatusEnum int32

const (
	ProjectTaskStatusEnum_Todo       ProjectTaskStatusEnum = 1
	ProjectTaskStatusEnum_InProgress ProjectTaskStatusEnum = 2
	ProjectTaskStatusEnum_Skipped    ProjectTaskStatusEnum = 3
	ProjectTaskStatusEnum_Done       ProjectTaskStatusEnum = 4
)

// Verified against Project.json schema on September 26, 2025
type ProjectTemplateCvssVersion int32

const (
	ProjectTemplateCvssVersion_Version20 ProjectTemplateCvssVersion = 0
	ProjectTemplateCvssVersion_Version30 ProjectTemplateCvssVersion = 1
	ProjectTemplateCvssVersion_Version31 ProjectTemplateCvssVersion = 2
	ProjectTemplateCvssVersion_Version40 ProjectTemplateCvssVersion = 3
)

// Verified against Project.json schema on September 26, 2025
type ProjectTemplateStatusEnum int32

const (
	ProjectTemplateStatusEnum_Draft     ProjectTemplateStatusEnum = 1
	ProjectTemplateStatusEnum_Published ProjectTemplateStatusEnum = 2
)

// Verified against Project.json schema on September 26, 2025
type ProjectTypeEnum int32

const (
	ProjectTypeEnum_Timebased  ProjectTypeEnum = 0
	ProjectTypeEnum_Continuous ProjectTypeEnum = 1
)

// Verified against Project.json schema on September 26, 2025
type ProjectTypeOfTestingEnum int32

const (
	ProjectTypeOfTestingEnum_Blackbox ProjectTypeOfTestingEnum = 1
	ProjectTypeOfTestingEnum_Whitebox ProjectTypeOfTestingEnum = 2
	ProjectTypeOfTestingEnum_Graybox  ProjectTypeOfTestingEnum = 3
)

// Verified against Misc.json schema on September 26, 2025
type ClientProjectTypeOfTestingEnum int32

const (
	ClientProjectTypeOfTestingEnum_WebApp                 ClientProjectTypeOfTestingEnum = 1
	ClientProjectTypeOfTestingEnum_WebService             ClientProjectTypeOfTestingEnum = 2
	ClientProjectTypeOfTestingEnum_MobileApp              ClientProjectTypeOfTestingEnum = 3
	ClientProjectTypeOfTestingEnum_InternalInfrastructure ClientProjectTypeOfTestingEnum = 4
	ClientProjectTypeOfTestingEnum_ExternalInfrastructure ClientProjectTypeOfTestingEnum = 5
	ClientProjectTypeOfTestingEnum_WirelessInfrastructure ClientProjectTypeOfTestingEnum = 6
)

// Verified against Misc.json schema on September 26, 2025
type ClientRequestFormFieldTypeEnum int32

const (
	ClientRequestFormFieldTypeEnum_Text        ClientRequestFormFieldTypeEnum = 1
	ClientRequestFormFieldTypeEnum_Multitext   ClientRequestFormFieldTypeEnum = 2
	ClientRequestFormFieldTypeEnum_Dropdown    ClientRequestFormFieldTypeEnum = 3
	ClientRequestFormFieldTypeEnum_Multiselect ClientRequestFormFieldTypeEnum = 4
	ClientRequestFormFieldTypeEnum_DateTime    ClientRequestFormFieldTypeEnum = 5
)

// Verified against Misc.json schema on September 26, 2025
type ComplianceNormStatusEnum int32

const (
	ComplianceNormStatusEnum_Draft     ComplianceNormStatusEnum = 1
	ComplianceNormStatusEnum_Published ComplianceNormStatusEnum = 2
)

// Verified against Misc.json schema on September 26, 2025
type ReoccurrenceOptionEnum int32

const (
	ReoccurrenceOptionEnum_AlwaysNew                    ReoccurrenceOptionEnum = 0
	ReoccurrenceOptionEnum_MatchByTitleSeverityAndAsset ReoccurrenceOptionEnum = 1
)

// Verified against Misc.json schema on September 26, 2025
type RolePortalEnum int32

const (
	RolePortalEnum_Pentester RolePortalEnum = 0
	RolePortalEnum_Client    RolePortalEnum = 1
)

// Verified against Misc.json schema on September 26, 2025
type RunTriggerTypeEnum int32

const (
	RunTriggerTypeEnum_OnDemand RunTriggerTypeEnum = 0
	RunTriggerTypeEnum_Schedule RunTriggerTypeEnum = 1
)

// Verified against Misc.json schema on September 26, 2025
type RunTypeEnum int32

const (
	RunTypeEnum_Manual RunTypeEnum = 0
	RunTypeEnum_Scan   RunTypeEnum = 1
)

// Verified against Misc.json schema on September 26, 2025
type ImportFileTypeEnum int32

const (
	ImportFileTypeEnum_CyverXML                   ImportFileTypeEnum = 0
	ImportFileTypeEnum_Burp                       ImportFileTypeEnum = 1
	ImportFileTypeEnum_NessusCSV                  ImportFileTypeEnum = 2
	ImportFileTypeEnum_Nmap                       ImportFileTypeEnum = 3
	ImportFileTypeEnum_Qualys                     ImportFileTypeEnum = 4
	ImportFileTypeEnum_Invicti                    ImportFileTypeEnum = 5
	ImportFileTypeEnum_Nexpose                    ImportFileTypeEnum = 6
	ImportFileTypeEnum_OpenVas                    ImportFileTypeEnum = 7
	ImportFileTypeEnum_Zap                        ImportFileTypeEnum = 8
	ImportFileTypeEnum_CyverCSV                   ImportFileTypeEnum = 9
	ImportFileTypeEnum_Acunetix                   ImportFileTypeEnum = 10
	ImportFileTypeEnum_Blindspot                  ImportFileTypeEnum = 11
	ImportFileTypeEnum_CarbonBlack                ImportFileTypeEnum = 12
	ImportFileTypeEnum_Plextrac                   ImportFileTypeEnum = 13
	ImportFileTypeEnum_QualysScan                 ImportFileTypeEnum = 14
	ImportFileTypeEnum_Prowler                    ImportFileTypeEnum = 15
	ImportFileTypeEnum_NessusXML                  ImportFileTypeEnum = 16
	ImportFileTypeEnum_Scout                      ImportFileTypeEnum = 17
	ImportFileTypeEnum_QualysScanApi              ImportFileTypeEnum = 18
	ImportFileTypeEnum_PurpleKnight               ImportFileTypeEnum = 19
	ImportFileTypeEnum_HCLAppScan                 ImportFileTypeEnum = 20
	ImportFileTypeEnum_Snyk                       ImportFileTypeEnum = 21
	ImportFileTypeEnum_Checkmarx                  ImportFileTypeEnum = 22
	ImportFileTypeEnum_AppCheck                   ImportFileTypeEnum = 23
	ImportFileTypeEnum_Nipper                     ImportFileTypeEnum = 24
	ImportFileTypeEnum_PentestToolsNetworkScanner ImportFileTypeEnum = 25
	ImportFileTypeEnum_PentestToolsWebsiteScanner ImportFileTypeEnum = 26
)

// Verified against Project.json schema on September 26, 2025
type ProjectDtoV2 struct {
	ID                   string                      `json:"id"`
	Name                 *string                     `json:"name,omitempty"`
	Code                 *string                     `json:"code,omitempty"`
	MethodologyMarkdown  *string                     `json:"methodologyMarkdown,omitempty"`
	Status               *string                     `json:"status,omitempty"`
	CvssVersion          *ProjectTemplateCvssVersion `json:"cvssVersion,omitempty"`
	ClientID             string                      `json:"clientId"`
	ClientName           *string                     `json:"clientName,omitempty"`
	ProjectTemplateID    string                      `json:"projectTemplateId"`
	ProjectTemplateTitle *string                     `json:"projectTemplateTitle,omitempty"`
	ReportTemplateID     *string                     `json:"reportTemplateId,omitempty"`
	ReportTemplateTitle  *string                     `json:"reportTemplateTitle,omitempty"`
	QuoteTemplateID      string                      `json:"quoteTemplateId"`
	QuoteTemplateTitle   *string                     `json:"quoteTemplateTitle,omitempty"`
	ProjectLeadID        string                      `json:"projectLeadId"`
	ProjectLeadName      *string                     `json:"projectLeadName,omitempty"`
	ProjectReviewerID    string                      `json:"projectReviewerId"`
	ProjectReviewerName  *string                     `json:"projectReviewerName,omitempty"`
	ProjectDates         *ProjectDatesDtoV2          `json:"projectDates,omitempty"`
	LabelList            []*LabelDto                 `json:"labelList,omitempty"`
	FindingIDList        []string                    `json:"findingIdList,omitempty"`
	AssetIDList          []string                    `json:"assetIdList,omitempty"`
	ChecklistIDList      []string                    `json:"checklistIdList,omitempty"`
	ComplianceNormIDList []string                    `json:"complianceNormIdList,omitempty"`
	UserIDList           []string                    `json:"userIdList,omitempty"`
	TeamIDList           []string                    `json:"teamIdList,omitempty"`
	ReportTemplates      []*ProjectReportTemplateDto `json:"reportTemplates,omitempty"`
}

// Verified against Project.json schema on September 26, 2025
type ProjectDatesDtoV2 struct {
	StartDate         string               `json:"startDate"`
	EndDate           string               `json:"endDate"`
	StartTesting      string               `json:"startTesting"`
	EndTesting        string               `json:"endTesting"`
	CreationTime      string               `json:"creationTime"`
	ReportDueDate     *string              `json:"reportDueDate,omitempty"`
	ReportPublishedAt *string              `json:"reportPublishedAt,omitempty"`
	PlanningDates     []*PlanningDateDtoV2 `json:"planningDates,omitempty"`
}

// Verified against Project.json schema on September 26, 2025
type PlanningDateDtoV2 struct {
	Date        string  `json:"date"`
	Description *string `json:"description,omitempty"`
}

// Verified against Project.json schema on September 26, 2025
type ProjectReportTemplateDto struct {
	ID               string  `json:"id"`
	Name             *string `json:"name,omitempty"`
	ProjectID        string  `json:"projectId"`
	ReportTemplateID string  `json:"reportTemplateId"`
}

// Verified against Project.json schema on September 26, 2025
type ProjectChecklistDto struct {
	ID            string                 `json:"id"`
	Name          *string                `json:"name,omitempty"`
	ExternalUrl   *string                `json:"externalUrl,omitempty"`
	Description   *string                `json:"description,omitempty"`
	TaskGroupList []*ProjectTaskGroupDto `json:"taskGroupList,omitempty"`
}

// Verified against Project.json schema on September 26, 2025
type ProjectTaskGroupDto struct {
	ID          string            `json:"id"`
	Code        *string           `json:"code,omitempty"`
	Name        *string           `json:"name,omitempty"`
	Description *string           `json:"description,omitempty"`
	TaskList    []*ProjectTaskDto `json:"taskList,omitempty"`
}

// Verified against Project.json schema on September 26, 2025
type ProjectTaskDto struct {
	ID          string                 `json:"id"`
	Code        *string                `json:"code,omitempty"`
	Name        *string                `json:"name,omitempty"`
	ExternalUrl *string                `json:"externalUrl,omitempty"`
	Description *string                `json:"description,omitempty"`
	Status      *ProjectTaskStatusEnum `json:"status,omitempty"`
}

// Verified against Project.json schema on September 26, 2025
type ProjectControlDto struct {
	ID          string  `json:"id"`
	Code        *string `json:"code,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// Verified against Project.json schema on September 26, 2025
type ProjectControlGroupDto struct {
	ID          string               `json:"id"`
	Name        *string              `json:"name,omitempty"`
	Description *string              `json:"description,omitempty"`
	ControlList []*ProjectControlDto `json:"controlList,omitempty"`
}

// Verified against Project.json schema on September 26, 2025
type ProjectComplianceNormDto struct {
	ID               string                    `json:"id"`
	Code             *string                   `json:"code,omitempty"`
	Name             *string                   `json:"name,omitempty"`
	Description      *string                   `json:"description,omitempty"`
	ControlGroupList []*ProjectControlGroupDto `json:"controlGroupList,omitempty"`
}

// Verified against Project.json schema on September 26, 2025
type ProjectTemplateDto struct {
	ID                            string                              `json:"id"`
	ReportTemplateID              *string                             `json:"reportTemplateId,omitempty"`
	FindingFieldsTemplateID       *string                             `json:"findingFieldsTemplateId,omitempty"`
	ProjectWorkflowTemplateID     *string                             `json:"projectWorkflowTemplateId,omitempty"`
	Status                        *ProjectTemplateStatusEnum          `json:"status,omitempty"`
	Type                          *ProjectTypeEnum                    `json:"type,omitempty"`
	TypeOfFindings                *FindingTypeEnum                    `json:"typeOfFindings,omitempty"`
	ClientCanReportTypeOfFindings *FindingTypeEnum                    `json:"clientCanReportTypeOfFindings,omitempty"`
	CvssVersion                   *ProjectTemplateCvssVersion         `json:"cvssVersion,omitempty"`
	Title                         *string                             `json:"title,omitempty"`
	ChecklistTemplateList         []*ChecklistTemplateDto             `json:"checklistTemplateList,omitempty"`
	ComplianceNormTemplateList    []*ComplianceNormTemplateDto        `json:"complianceNormTemplateList,omitempty"`
	ReportTemplates               []*ProjectTemplateReportTemplateDto `json:"reportTemplates,omitempty"`
}

// Verified against Project.json schema on September 26, 2025
type ProjectTemplateReportTemplateDto struct {
	ID                  string  `json:"id"`
	Name                *string `json:"name,omitempty"`
	ProjectTemplateID   string  `json:"projectTemplateId"`
	ReportTemplateID    string  `json:"reportTemplateId"`
	ReportTemplateTitle *string `json:"reportTemplateTitle,omitempty"`
	Order               int32   `json:"order"`
}

// Verified against Project.json schema on September 26, 2025
type ProjectChecklistDtoListResultDto struct {
	Items []*ProjectChecklistDto `json:"items,omitempty"`
}

// Verified against Project.json schema on September 26, 2025
type ProjectChecklistDtoListResultDtoAjaxResponse struct {
	TargetUrl           *string                           `json:"targetUrl,omitempty"`
	Success             bool                              `json:"success"`
	Error               *ErrorInfo                        `json:"error,omitempty"`
	UnAuthorizedRequest bool                              `json:"unAuthorizedRequest"`
	Abp                 bool                              `json:"__abp"`
	Result              *ProjectChecklistDtoListResultDto `json:"result,omitempty"`
}

// Verified against Project.json schema on September 26, 2025
type ProjectComplianceNormDtoListResultDto struct {
	Items []*ProjectComplianceNormDto `json:"items,omitempty"`
}

// Verified against Project.json schema on September 26, 2025
type ProjectComplianceNormDtoListResultDtoAjaxResponse struct {
	TargetUrl           *string                                `json:"targetUrl,omitempty"`
	Success             bool                                   `json:"success"`
	Error               *ErrorInfo                             `json:"error,omitempty"`
	UnAuthorizedRequest bool                                   `json:"unAuthorizedRequest"`
	Abp                 bool                                   `json:"__abp"`
	Result              *ProjectComplianceNormDtoListResultDto `json:"result,omitempty"`
}

// Verified against Project.json schema on September 26, 2025
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

// Verified against Project.json schema on September 26, 2025
type ProjectDtoV2PagedResultDto struct {
	Items      []*ProjectDtoV2 `json:"items,omitempty"`
	TotalCount int32           `json:"totalCount"`
}

// Verified against Project.json schema on September 26, 2025
type ProjectDtoV2PagedResultDtoAjaxResponse struct {
	TargetUrl           *string                     `json:"targetUrl,omitempty"`
	Success             bool                        `json:"success"`
	Error               *ErrorInfo                  `json:"error,omitempty"`
	UnAuthorizedRequest bool                        `json:"unAuthorizedRequest"`
	Abp                 bool                        `json:"__abp"`
	Result              *ProjectDtoV2PagedResultDto `json:"result,omitempty"`
}

// --- Additional Enums for Continuous Projects ---

// Verified against Misc.json schema on September 26, 2025
type EvidenceOptionEnum int32

const (
	EvidenceOptionEnum_Aggregate     EvidenceOptionEnum = 0
	EvidenceOptionEnum_DontAggregate EvidenceOptionEnum = 1
)

// Verified against Continuous.json schema on September 26, 2025
type DayOfWeek int32

const (
	DayOfWeek_Sunday    DayOfWeek = 0
	DayOfWeek_Monday    DayOfWeek = 1
	DayOfWeek_Tuesday   DayOfWeek = 2
	DayOfWeek_Wednesday DayOfWeek = 3
	DayOfWeek_Thursday  DayOfWeek = 4
	DayOfWeek_Friday    DayOfWeek = 5
	DayOfWeek_Saturday  DayOfWeek = 6
)

// Verified against Misc.json schema on September 26, 2025
type TargetFieldEnum int32

const (
	TargetFieldEnum_Title      TargetFieldEnum = 0
	TargetFieldEnum_IP         TargetFieldEnum = 1
	TargetFieldEnum_URL        TargetFieldEnum = 2
	TargetFieldEnum_Hostname   TargetFieldEnum = 3
	TargetFieldEnum_Path       TargetFieldEnum = 4
	TargetFieldEnum_SSID       TargetFieldEnum = 5
	TargetFieldEnum_Repository TargetFieldEnum = 6
	TargetFieldEnum_DomainAD   TargetFieldEnum = 7
)

// Verified against Misc.json schema on September 26, 2025
type ClientAssignmentDto struct {
	ID                  string                          `json:"id"`
	StartEnum           *ProjectStartEnum               `json:"startEnum,omitempty"`
	Frequency           *ProjectFrequencyEnum           `json:"frequency,omitempty"`
	Objectives          *ProjectObjectivesEnum          `json:"objectives,omitempty"`
	TypeOfTesting       *ProjectTypeOfTestingEnum       `json:"typeOfTesting,omitempty"`
	ClientTypeOfTesting *ClientProjectTypeOfTestingEnum `json:"clientTypeOfTesting,omitempty"`
	RequestedDate       string                          `json:"requestedDate"`
	Details             *string                         `json:"details,omitempty"`
	HowToAccessAssets   *string                         `json:"howToAccessAssets,omitempty"`
	Credentials         *string                         `json:"credentials,omitempty"`
	DropDown1           *int32                          `json:"dropDown1,omitempty"`
	DropDown2           *int32                          `json:"dropDown2,omitempty"`
	DropDown3           *int32                          `json:"dropDown3,omitempty"`
	DropDown4           *int32                          `json:"dropDown4,omitempty"`
	DropDown5           *int32                          `json:"dropDown5,omitempty"`
	Text1               *string                         `json:"text1,omitempty"`
	Text2               *string                         `json:"text2,omitempty"`
	Text3               *string                         `json:"text3,omitempty"`
	Text4               *string                         `json:"text4,omitempty"`
	Text5               *string                         `json:"text5,omitempty"`
	Multitext1          *string                         `json:"multitext1,omitempty"`
	Multitext2          *string                         `json:"multitext2,omitempty"`
	Multitext3          *string                         `json:"multitext3,omitempty"`
	Multitext4          *string                         `json:"multitext4,omitempty"`
	Multitext5          *string                         `json:"multitext5,omitempty"`
	Multiselect1        *string                         `json:"multiselect1,omitempty"`
	Multiselect2        *string                         `json:"multiselect2,omitempty"`
	Multiselect3        *string                         `json:"multiselect3,omitempty"`
	Multiselect4        *string                         `json:"multiselect4,omitempty"`
	Multiselect5        *string                         `json:"multiselect5,omitempty"`
}

// Verified against Misc.json schema on September 26, 2025
type ControlGroupTemplateDto struct {
	ID          string                `json:"id"`
	Name        *string               `json:"name,omitempty"`
	Description *string               `json:"description,omitempty"`
	ControlList []*ControlTemplateDto `json:"controlList,omitempty"`
}

// Verified against Misc.json schema on September 26, 2025
type ControlTemplateDto struct {
	ID          string  `json:"id"`
	Code        *string `json:"code,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// Verified against Misc.json schema on September 26, 2025
type ComplianceNormTemplateDto struct {
	ID               string                     `json:"id"`
	Status           *ComplianceNormStatusEnum  `json:"status,omitempty"`
	Code             *string                    `json:"code,omitempty"`
	Name             *string                    `json:"name,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	ControlGroupList []*ControlGroupTemplateDto `json:"controlGroupList,omitempty"`
}

// Verified against Misc.json schema on September 26, 2025
type FormDataDto struct {
	ID          int32                           `json:"id"`
	Code        *string                         `json:"code,omitempty"`
	Name        *string                         `json:"name,omitempty"`
	Description *string                         `json:"description,omitempty"`
	Options     []*FormDataOptionsDto           `json:"options,omitempty"`
	Type        *ClientRequestFormFieldTypeEnum `json:"type,omitempty"`
}

// Verified against Misc.json schema on September 26, 2025
type FormDataOptionsDto struct {
	Value int32   `json:"value"`
	Name  *string `json:"name,omitempty"`
}

// --- Continuous Project Models ---

// Verified against Continuous.json schema on September 26, 2025
type ContinuousProjectCloseFindingOptionEnum int32

const (
	ContinuousProjectCloseFindingOptionEnum_NoFindingMatch          ContinuousProjectCloseFindingOptionEnum = 0
	ContinuousProjectCloseFindingOptionEnum_NoFindingMatchAndAssets ContinuousProjectCloseFindingOptionEnum = 1
)

// Verified against Continuous.json schema on September 26, 2025
type ContinuousProjectDatesDto struct {
	StartDate         string  `json:"startDate"`
	EndDate           string  `json:"endDate"`
	CreationTime      string  `json:"creationTime"`
	ReportDueDate     *string `json:"reportDueDate,omitempty"`
	ReportPublishedAt *string `json:"reportPublishedAt,omitempty"`
}

// Verified against Continuous.json schema on September 26, 2025
type ContinuousProjectDto struct {
	ID                   string                        `json:"id"`
	Name                 *string                       `json:"name,omitempty"`
	Code                 *string                       `json:"code,omitempty"`
	MethodologyMarkdown  *string                       `json:"methodologyMarkdown,omitempty"`
	Status               *ContinuousProjectStatusEnum  `json:"status,omitempty"`
	CvssVersion          *ProjectTemplateCvssVersion   `json:"cvssVersion,omitempty"`
	ClientID             string                        `json:"clientId"`
	ClientName           *string                       `json:"clientName,omitempty"`
	ProjectTemplateID    string                        `json:"projectTemplateId"`
	ProjectTemplateTitle *string                       `json:"projectTemplateTitle,omitempty"`
	ReportTemplateID     *string                       `json:"reportTemplateId,omitempty"`
	ReportTemplateTitle  *string                       `json:"reportTemplateTitle,omitempty"`
	QuoteTemplateID      string                        `json:"quoteTemplateId"`
	QuoteTemplateTitle   *string                       `json:"quoteTemplateTitle,omitempty"`
	ProjectLeadID        string                        `json:"projectLeadId"`
	ProjectLeadName      *string                       `json:"projectLeadName,omitempty"`
	ProjectReviewerID    string                        `json:"projectReviewerId"`
	ProjectReviewerName  *string                       `json:"projectReviewerName,omitempty"`
	ProjectDates         *ContinuousProjectDatesDto    `json:"projectDates,omitempty"`
	ClientAssignment     *ClientAssignmentDto          `json:"clientAssignment,omitempty"`
	Settings             *ContinuousProjectSettingsDto `json:"settings,omitempty"`
	LabelList            []*LabelDto                   `json:"labelList,omitempty"`
	FindingIDList        []string                      `json:"findingIdList,omitempty"`
	AssetIDList          []string                      `json:"assetIdList,omitempty"`
	RunIDList            []string                      `json:"runIdList,omitempty"`
	ReportTemplates      []*ProjectReportTemplateDto   `json:"reportTemplates,omitempty"`
}

// Verified against Continuous.json schema on September 26, 2025
type ContinuousProjectDtoAjaxResponse struct {
	TargetUrl           *string               `json:"targetUrl,omitempty"`
	Success             bool                  `json:"success"`
	Error               *ErrorInfo            `json:"error,omitempty"`
	UnAuthorizedRequest bool                  `json:"unAuthorizedRequest"`
	Abp                 bool                  `json:"__abp"`
	Result              *ContinuousProjectDto `json:"result,omitempty"`
}

// Verified against Continuous.json schema on September 26, 2025
type ContinuousProjectDtoPagedResultDto struct {
	Items      []*ContinuousProjectDto `json:"items,omitempty"`
	TotalCount int32                   `json:"totalCount"`
}

// Verified against Continuous.json schema on September 26, 2025
type ContinuousProjectDtoPagedResultDtoAjaxResponse struct {
	TargetUrl           *string                             `json:"targetUrl,omitempty"`
	Success             bool                                `json:"success"`
	Error               *ErrorInfo                          `json:"error,omitempty"`
	UnAuthorizedRequest bool                                `json:"unAuthorizedRequest"`
	Abp                 bool                                `json:"__abp"`
	Result              *ContinuousProjectDtoPagedResultDto `json:"result,omitempty"`
}

// Verified against Continuous.json schema on September 26, 2025
type ContinuousProjectFindingMatchingOption int32

const (
	ContinuousProjectFindingMatchingOption_TitleAndSeverity       ContinuousProjectFindingMatchingOption = 0
	ContinuousProjectFindingMatchingOption_TitleSeverityAndAssets ContinuousProjectFindingMatchingOption = 1
)

// Verified against Continuous.json schema on September 26, 2025
type ContinuousProjectSettingsDto struct {
	GenerateReportOnRunFinish bool                                     `json:"generateReportOnRunFinish"`
	GeneratePdfAndPublish     bool                                     `json:"generatePdfAndPublish"`
	AggregateOption           *EvidenceOptionEnum                      `json:"aggregateOption,omitempty"`
	Scanner                   *ContinuousProjectVulnerabilityTypeEnum  `json:"scanner,omitempty"`
	DefaultFindingStatus      *FindingStatusEnum                       `json:"defaultFindingStatus,omitempty"`
	IsWeeklyScanEnabled       bool                                     `json:"isWeeklyScanEnabled"`
	WeeklyScanDay             *DayOfWeek                               `json:"weeklyScanDay,omitempty"`
	WeeklyScanTime            *string                                  `json:"weeklyScanTime,omitempty"`
	FindingMergeOption        *FindingMergeOptionEnum                  `json:"findingMergeOption,omitempty"`
	UseCase                   *ContinuousProjectUseCaseEnum            `json:"useCase,omitempty"`
	CloseFindingOption        *ContinuousProjectCloseFindingOptionEnum `json:"closeFindingOption,omitempty"`
	FindingMatchingOption     *ContinuousProjectFindingMatchingOption  `json:"findingMatchingOption,omitempty"`
	TargetField               *TargetFieldEnum                         `json:"targetField,omitempty"`
}

// Verified against Continuous.json schema on September 26, 2025
type ContinuousProjectStatusEnum int32

const (
	ContinuousProjectStatusEnum_Active    ContinuousProjectStatusEnum = 0
	ContinuousProjectStatusEnum_Stopped   ContinuousProjectStatusEnum = 1
	ContinuousProjectStatusEnum_Requested ContinuousProjectStatusEnum = 2
)

// Verified against Continuous.json schema on September 26, 2025
type ContinuousProjectUseCaseEnum int32

const (
	ContinuousProjectUseCaseEnum_RunBased ContinuousProjectUseCaseEnum = 0
	ContinuousProjectUseCaseEnum_RealTime ContinuousProjectUseCaseEnum = 1
)

// Verified against Continuous.json schema on September 26, 2025
type ContinuousProjectVulnerabilityTypeEnum int32

const (
	ContinuousProjectVulnerabilityTypeEnum_None    ContinuousProjectVulnerabilityTypeEnum = 0
	ContinuousProjectVulnerabilityTypeEnum_ReNgine ContinuousProjectVulnerabilityTypeEnum = 1
	ContinuousProjectVulnerabilityTypeEnum_Tenable ContinuousProjectVulnerabilityTypeEnum = 2
)

// --- Finding Models ---

type FindingSeverityEnum int32

const (
	FindingSeverityEnum_Critical FindingSeverityEnum = 1
	FindingSeverityEnum_High     FindingSeverityEnum = 2
	FindingSeverityEnum_Medium   FindingSeverityEnum = 3
	FindingSeverityEnum_Low      FindingSeverityEnum = 4
	FindingSeverityEnum_Info     FindingSeverityEnum = 5
)

// Verified against Finding.json schema on September 26, 2025
type FindingStatusEnum int32

const (
	FindingStatusEnum_Draft         FindingStatusEnum = 1
	FindingStatusEnum_PendingFix    FindingStatusEnum = 2
	FindingStatusEnum_Fixed         FindingStatusEnum = 3
	FindingStatusEnum_ReadyRetest   FindingStatusEnum = 4
	FindingStatusEnum_Accepted      FindingStatusEnum = 5
	FindingStatusEnum_ToReview      FindingStatusEnum = 6
	FindingStatusEnum_Reviewed      FindingStatusEnum = 7
	FindingStatusEnum_Mitigated     FindingStatusEnum = 8
	FindingStatusEnum_PartialFix    FindingStatusEnum = 9
	FindingStatusEnum_FalsePositive FindingStatusEnum = 10
	FindingStatusEnum_Raised        FindingStatusEnum = 11
	FindingStatusEnum_ReOpen        FindingStatusEnum = 12
	FindingStatusEnum_Acknowledged  FindingStatusEnum = 13
	FindingStatusEnum_Identified    FindingStatusEnum = 14
)

// Verified against Finding.json schema on September 26, 2025
type FindingTypeEnum int32

const (
	FindingTypeEnum_Vulnerability FindingTypeEnum = 1
	FindingTypeEnum_Nonconformity FindingTypeEnum = 2
	FindingTypeEnum_Observation   FindingTypeEnum = 4
	FindingTypeEnum_Incident      FindingTypeEnum = 8
	FindingTypeEnum_Risk          FindingTypeEnum = 16
)

// Verified against Finding.json schema on September 26, 2025
type FindingCriticalityEnum int32

const (
	FindingCriticalityEnum_Info     FindingCriticalityEnum = 0
	FindingCriticalityEnum_Low      FindingCriticalityEnum = 1
	FindingCriticalityEnum_Medium   FindingCriticalityEnum = 2
	FindingCriticalityEnum_High     FindingCriticalityEnum = 3
	FindingCriticalityEnum_Critical FindingCriticalityEnum = 4
)

// Verified against Finding.json schema on September 26, 2025
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

// Verified against Finding.json schema on September 26, 2025
type FindingPciComplianceEnum int32

const (
	FindingPciComplianceEnum_Pass FindingPciComplianceEnum = 0
	FindingPciComplianceEnum_Fail FindingPciComplianceEnum = 1
)

// Verified against Finding.json schema on September 26, 2025
type FindingMergeOptionEnum int32

const (
	FindingMergeOptionEnum_Skip      FindingMergeOptionEnum = 0
	FindingMergeOptionEnum_Overwrite FindingMergeOptionEnum = 1
)

// --- Additional DTOs for FindingDto ---

// Verified against Finding.json schema on September 26, 2025
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

// Note: ProjectTaskDto and ProjectControlDto are already defined above with proper schema matching

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

// Verified against Finding.json schema on September 26, 2025
type FindingEvidenceDto struct {
	ID                       *string                   `json:"id,omitempty"`
	Title                    string                    `json:"title"`
	Location                 *string                   `json:"location,omitempty"`
	Version                  *string                   `json:"version,omitempty"`
	Reproduce                *string                   `json:"reproduce,omitempty"`
	Results                  *string                   `json:"results,omitempty"`
	IssueDetails             *string                   `json:"issueDetails,omitempty"`
	IsVisibleInReport        bool                      `json:"isVisibleInReport"`
	IP                       *string                   `json:"ip,omitempty"`
	Hostname                 *string                   `json:"hostname,omitempty"`
	Port                     *string                   `json:"port,omitempty"`
	Protocol                 *string                   `json:"protocol,omitempty"`
	EvidenceComplianceStatus *FindingPciComplianceEnum `json:"evidenceComplianceStatus,omitempty"`
}

// Verified against Finding.json schema on September 26, 2025
type FindingRunDto struct {
	RunID                string             `json:"runId"`
	RunNumber            int32              `json:"runNumber"`
	InitialFindingStatus *FindingStatusEnum `json:"initialFindingStatus,omitempty"`
	FinalFindingStatus   *FindingStatusEnum `json:"finalFindingStatus,omitempty"`
	FinishedAt           *string            `json:"finishedAt,omitempty"`
	Description          *string            `json:"description,omitempty"`
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

// Verified against Finding.json schema on September 26, 2025
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

// Verified against Finding.json schema on September 26, 2025
type FindingDtoAjaxResponse struct {
	TargetUrl           *string     `json:"targetUrl,omitempty"`
	Success             bool        `json:"success"`
	Error               *ErrorInfo  `json:"error,omitempty"`
	UnAuthorizedRequest bool        `json:"unAuthorizedRequest"`
	Abp                 bool        `json:"__abp"`
	Result              *FindingDto `json:"result,omitempty"`
}

// Verified against Finding.json schema on September 26, 2025
type FindingDtoPagedResultDto struct {
	Items      []*FindingDto `json:"items,omitempty"`
	TotalCount int32         `json:"totalCount"`
}

// Verified against Finding.json schema on September 26, 2025
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

// Verified against User.json schema on September 26, 2025
type UserDto struct {
	ID                 string  `json:"id"`
	ClientID           *string `json:"clientId,omitempty"`
	Name               *string `json:"name,omitempty"`
	Surname            *string `json:"surname,omitempty"`
	Email              *string `json:"email,omitempty"`
	PhoneNumber        *string `json:"phoneNumber,omitempty"`
	IsEmailConfirmed   bool    `json:"isEmailConfirmed"`
	IsActive           bool    `json:"isActive"`
	IsTwoFactorEnabled bool    `json:"isTwoFactorEnabled"`
	IsLockoutEnabled   bool    `json:"isLockoutEnabled"`
	Role               *string `json:"role,omitempty"`
	CreationTime       *string `json:"creationTime,omitempty"`
}

// Verified against User.json schema on September 26, 2025
type UserDtoAjaxResponse struct {
	TargetUrl           *string    `json:"targetUrl,omitempty"`
	Success             bool       `json:"success"`
	Error               *ErrorInfo `json:"error,omitempty"`
	UnAuthorizedRequest bool       `json:"unAuthorizedRequest"`
	Abp                 bool       `json:"__abp"`
	Result              *UserDto   `json:"result,omitempty"`
}

// Verified against User.json schema on September 26, 2025
type UserDtoPagedResultDto struct {
	Items      []*UserDto `json:"items,omitempty"`
	TotalCount int32      `json:"totalCount"`
}

// Verified against User.json schema on September 26, 2025
type UserDtoPagedResultDtoAjaxResponse struct {
	TargetUrl           *string                `json:"targetUrl,omitempty"`
	Success             bool                   `json:"success"`
	Error               *ErrorInfo             `json:"error,omitempty"`
	UnAuthorizedRequest bool                   `json:"unAuthorizedRequest"`
	Abp                 bool                   `json:"__abp"`
	Result              *UserDtoPagedResultDto `json:"result,omitempty"`
}

// --- Request Models ---

// Verified against Request.json schema on September 26, 2025
type RequestFormDataDto struct {
	Code  *string `json:"code,omitempty"`
	Value *string `json:"value,omitempty"`
}

// Verified against Request.json schema on September 26, 2025
type RequestProjectFormDto struct {
	GUID        string                `json:"guid"`
	Title       *string               `json:"title,omitempty"`
	Description *string               `json:"description,omitempty"`
	FormData    []*RequestFormDataDto `json:"formData,omitempty"`
}

// Verified against Request.json schema on September 26, 2025
type RequestProjectFormDtoPagedResultDto struct {
	Items      []*RequestProjectFormDto `json:"items,omitempty"`
	TotalCount int32                    `json:"totalCount"`
}

// Verified against Request.json schema on September 26, 2025
type RequestProjectFormDtoPagedResultDtoAjaxResponse struct {
	TargetUrl           *string                              `json:"targetUrl,omitempty"`
	Success             bool                                 `json:"success"`
	Error               *ErrorInfo                           `json:"error,omitempty"`
	UnAuthorizedRequest bool                                 `json:"unAuthorizedRequest"`
	Abp                 bool                                 `json:"__abp"`
	Result              *RequestProjectFormDtoPagedResultDto `json:"result,omitempty"`
}

// Verified against Request.json schema on September 26, 2025
type RequestProjectRequest struct {
	Name              string                `json:"name"`
	RequestFormGUID   *string               `json:"requestFormGuid,omitempty"`
	RetestProjectGUID *string               `json:"retestProjectGuid,omitempty"`
	RequestStartDate  *string               `json:"requestStartDate,omitempty"`
	RequestedEndDate  *string               `json:"requestedEndDate,omitempty"`
	AssetIDList       []string              `json:"assetIdList,omitempty"`
	FormData          []*RequestFormDataDto `json:"formData,omitempty"`
}

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

// Verified against Misc.json schema on September 26, 2025
type CreateContinuousProjectRequest struct {
	Name              string                       `json:"name"`
	Code              *string                      `json:"code,omitempty"`
	Status            *ContinuousProjectStatusEnum `json:"status,omitempty"`
	ClientID          string                       `json:"clientId"`
	ProjectTemplateID string                       `json:"projectTemplateId"`
	LabelIDList       []string                     `json:"labelIdList,omitempty"`
}

// Verified against Misc.json schema on September 26, 2025
type CreateOrUpdateTeamRequest struct {
	Name         *string  `json:"name,omitempty"`
	ClientID     *string  `json:"clientId,omitempty"`
	TeamLeaderID *string  `json:"teamLeaderId,omitempty"`
	UserIDList   []string `json:"userIdList,omitempty"`
}

// Verified against Misc.json schema on September 26, 2025
type CreateUserRequest struct {
	Name               string        `json:"name"`
	Surname            string        `json:"surname"`
	EmailAddress       string        `json:"emailAddress"`
	PhoneNumber        *string       `json:"phoneNumber,omitempty"`
	IsActive           bool          `json:"isActive"`
	IsTwoFactorEnabled bool          `json:"isTwoFactorEnabled"`
	IsLockoutEnabled   bool          `json:"isLockoutEnabled"`
	Role               *ApiRolesEnum `json:"role,omitempty"`
	ClientGUID         *string       `json:"clientGuid,omitempty"`
}

// Verified against Misc.json schema on September 26, 2025
type CreatClientUserRequest struct {
	Name               string        `json:"name"`
	Surname            string        `json:"surname"`
	EmailAddress       string        `json:"emailAddress"`
	PhoneNumber        *string       `json:"phoneNumber,omitempty"`
	IsActive           bool          `json:"isActive"`
	IsTwoFactorEnabled bool          `json:"isTwoFactorEnabled"`
	IsLockoutEnabled   bool          `json:"isLockoutEnabled"`
	Role               *ApiRolesEnum `json:"role,omitempty"`
}

// Verified against Misc.json schema on September 26, 2025
type UpdateUserRequest struct {
	Name                            string        `json:"name"`
	Surname                         string        `json:"surname"`
	EmailAddress                    string        `json:"emailAddress"`
	PhoneNumber                     *string       `json:"phoneNumber,omitempty"`
	ShouldChangePasswordOnNextLogin bool          `json:"shouldChangePasswordOnNextLogin"`
	SendActivationEmail             bool          `json:"sendActivationEmail"`
	IsActive                        bool          `json:"isActive"`
	IsTwoFactorEnabled              bool          `json:"isTwoFactorEnabled"`
	IsLockoutEnabled                bool          `json:"isLockoutEnabled"`
	Role                            *ApiRolesEnum `json:"role,omitempty"`
}

// Verified against Misc.json schema on September 26, 2025
type UpdateContinuousProjectStatusRequest struct {
	Status *ContinuousProjectStatusEnum `json:"status,omitempty"`
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

// Verified against Misc.json schema on September 26, 2025
type ExternalAuthenticateModel struct {
	AuthProvider       string  `json:"authProvider"`
	ProviderKey        string  `json:"providerKey"`
	ProviderAccessCode string  `json:"providerAccessCode"`
	ReturnUrl          *string `json:"returnUrl,omitempty"`
	SingleSignIn       *bool   `json:"singleSignIn,omitempty"`
}

// Verified against Misc.json schema on September 26, 2025
type ExternalAuthenticateResultModel struct {
	AccessToken                 *string `json:"accessToken,omitempty"`
	EncryptedAccessToken        *string `json:"encryptedAccessToken,omitempty"`
	ExpireInSeconds             int32   `json:"expireInSeconds"`
	WaitingForActivation        bool    `json:"waitingForActivation"`
	ReturnUrl                   *string `json:"returnUrl,omitempty"`
	RefreshToken                *string `json:"refreshToken,omitempty"`
	RefreshTokenExpireInSeconds int32   `json:"refreshTokenExpireInSeconds"`
}

// Verified against Misc.json schema on September 26, 2025
type ExternalLoginProviderInfoModel struct {
	Name             *string            `json:"name,omitempty"`
	ClientID         *string            `json:"clientId,omitempty"`
	AdditionalParams map[string]*string `json:"additionalParams,omitempty"`
}

// Verified against Misc.json schema on September 26, 2025
type ImpersonatedAuthenticateResultModel struct {
	AccessToken          *string `json:"accessToken,omitempty"`
	EncryptedAccessToken *string `json:"encryptedAccessToken,omitempty"`
	ExpireInSeconds      int32   `json:"expireInSeconds"`
}

// Verified against Misc.json schema on September 26, 2025
type SwitchedAccountAuthenticateResultModel struct {
	AccessToken          *string `json:"accessToken,omitempty"`
	EncryptedAccessToken *string `json:"encryptedAccessToken,omitempty"`
	ExpireInSeconds      int32   `json:"expireInSeconds"`
}

// --- Legacy Models (for backward compatibility) ---

// These models are kept for backward compatibility with existing code
// Note: The new RequestProjectFormDto is defined above with proper schema matching

// Legacy version - use RequestProjectFormDtoPagedResultDtoAjaxResponse above for new code
type RequestProjectFormDtoPagedResultDtoAjaxResponseLegacy struct {
	Success bool                    `json:"success"`
	Data    []RequestProjectFormDto `json:"data"`
	Error   string                  `json:"error"`
}

// Legacy ContinuousProjectDto - use the new ContinuousProjectDto above for new code
type ContinuousProjectDtoLegacy struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// Legacy response - use the new ContinuousProjectDtoPagedResultDtoAjaxResponse above for new code
type ContinuousProjectDtoPagedResultDtoAjaxResponseLegacy struct {
	Success bool                         `json:"success"`
	Data    []ContinuousProjectDtoLegacy `json:"data"`
	Error   string                       `json:"error"`
}

// Legacy response - use the new ContinuousProjectDtoAjaxResponse above for new code
type ContinuousProjectDtoAjaxResponseLegacy struct {
	Success bool                       `json:"success"`
	Data    ContinuousProjectDtoLegacy `json:"data"`
	Error   string                     `json:"error"`
}

// Legacy version - use RequestProjectFormDtoPagedResultDtoAjaxResponse above for new code
type RequestProjectFormDtoAjaxResponseLegacy struct {
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
// Note: ProjectChecklistDto is already defined above with proper schema matching

// --- Project Compliance Models ---
// Note: ProjectComplianceNormDto is already defined above with proper schema matching

// --- Report Models ---

// Report Enums
// Verified against Reports.json schema on 2025-09-26
type ReportStatusEnum int32

const (
	ReportStatusEnum_Draft     ReportStatusEnum = 1
	ReportStatusEnum_Published ReportStatusEnum = 2
)

// Verified against Reports.json schema on 2025-09-26
type ReportSectionTypeEnum int32

const (
	ReportSectionTypeEnum_Standard        ReportSectionTypeEnum = 0
	ReportSectionTypeEnum_Cover           ReportSectionTypeEnum = 1
	ReportSectionTypeEnum_TableOfContents ReportSectionTypeEnum = 2
)

// Verified against Reports.json schema on 2025-09-26
type ReportTemplateStatusEnum int32

const (
	ReportTemplateStatusEnum_Draft     ReportTemplateStatusEnum = 1
	ReportTemplateStatusEnum_Published ReportTemplateStatusEnum = 2
)

// Verified against Reports.json schema on 2025-09-26
type ReportTemplatePdfToCEnum int32

const (
	ReportTemplatePdfToCEnum_Default ReportTemplatePdfToCEnum = 0
	ReportTemplatePdfToCEnum_H1      ReportTemplatePdfToCEnum = 1
	ReportTemplatePdfToCEnum_H1H2    ReportTemplatePdfToCEnum = 2
	ReportTemplatePdfToCEnum_H1H2H3  ReportTemplatePdfToCEnum = 3
)

// Report Section DTOs
// Verified against Reports.json schema on 2025-09-26
type ReportSectionDto struct {
	Title                *string `json:"title,omitempty"`
	HideFromClientPortal bool    `json:"hideFromClientPortal"`
	ContentHtml          *string `json:"contentHtml,omitempty"`
	Order                int32   `json:"order"`
}

// Verified against Reports.json schema on 2025-09-26
type ReportTemplateSectionDto struct {
	ID                           string                    `json:"id"`
	Title                        *string                   `json:"title,omitempty"`
	HideFromClientPortal         bool                      `json:"hideFromClientPortal"`
	Content                      *string                   `json:"content,omitempty"`
	IsInExternalDownload         bool                      `json:"isInExternalDownload"`
	Order                        int32                     `json:"order"`
	SectionType                  *ReportSectionTypeEnum    `json:"sectionType,omitempty"`
	TableOfContentsDepth         *ReportTemplatePdfToCEnum `json:"tableOfContentsDepth,omitempty"`
	TableOfContentsInternalLinks bool                      `json:"tableOfContentsInternalLinks"`
	Optional                     bool                      `json:"optional"`
}

// Main Report DTOs
// Verified against Reports.json schema on 2025-09-26
type ReportDto struct {
	GUID             string              `json:"guid"`
	Title            *string             `json:"title,omitempty"`
	Version          int32               `json:"version"`
	Status           *ReportStatusEnum   `json:"status,omitempty"`
	PublishedAt      *string             `json:"publishedAt,omitempty"`
	LastDraftSavedAt *string             `json:"lastDraftSavedAt,omitempty"`
	ProjectID        string              `json:"projectId"`
	ClientID         string              `json:"clientId"`
	SectionList      []*ReportSectionDto `json:"sectionList,omitempty"`
}

// Verified against Reports.json schema on 2025-09-26
type ReportDtoAjaxResponse struct {
	TargetUrl           *string    `json:"targetUrl,omitempty"`
	Success             bool       `json:"success"`
	Error               *ErrorInfo `json:"error,omitempty"`
	UnAuthorizedRequest bool       `json:"unAuthorizedRequest"`
	Abp                 bool       `json:"__abp"`
	Result              *ReportDto `json:"result,omitempty"`
}

// Verified against Reports.json schema on 2025-09-26
type ReportVersionDto struct {
	ID      string            `json:"id"`
	Version int32             `json:"version"`
	Status  *ReportStatusEnum `json:"status,omitempty"`
}

// Verified against Reports.json schema on 2025-09-26
type ReportVersionDtoListResultDto struct {
	Items []*ReportVersionDto `json:"items,omitempty"`
}

// Verified against Reports.json schema on 2025-09-26
type ReportVersionDtoListResultDtoAjaxResponse struct {
	TargetUrl           *string                        `json:"targetUrl,omitempty"`
	Success             bool                           `json:"success"`
	Error               *ErrorInfo                     `json:"error,omitempty"`
	UnAuthorizedRequest bool                           `json:"unAuthorizedRequest"`
	Abp                 bool                           `json:"__abp"`
	Result              *ReportVersionDtoListResultDto `json:"result,omitempty"`
}

// Report Template DTOs
// Verified against Reports.json schema on 2025-09-26
type ReportTemplateDto struct {
	ID                        string                      `json:"id"`
	Title                     *string                     `json:"title,omitempty"`
	Status                    *ReportTemplateStatusEnum   `json:"status,omitempty"`
	ReportTemplateSectionList []*ReportTemplateSectionDto `json:"reportTemplateSectionList,omitempty"`
}

// Verified against Reports.json schema on 2025-09-26
type ReportTemplateDtoPagedResultDto struct {
	Items      []*ReportTemplateDto `json:"items,omitempty"`
	TotalCount int32                `json:"totalCount"`
}

// Verified against Reports.json schema on 2025-09-26
type ReportTemplateDtoPagedResultDtoAjaxResponse struct {
	TargetUrl           *string                          `json:"targetUrl,omitempty"`
	Success             bool                             `json:"success"`
	Error               *ErrorInfo                       `json:"error,omitempty"`
	UnAuthorizedRequest bool                             `json:"unAuthorizedRequest"`
	Abp                 bool                             `json:"__abp"`
	Result              *ReportTemplateDtoPagedResultDto `json:"result,omitempty"`
}

// --- File Type Enum ---

// Verified against Misc.json schema on September 26, 2025
type FileTypeEnum int32

const (
	FileTypeEnum_FindingImport   FileTypeEnum = 4
	FileTypeEnum_Report          FileTypeEnum = 5
	FileTypeEnum_Evidence        FileTypeEnum = 6
	FileTypeEnum_Other           FileTypeEnum = 7
	FileTypeEnum_QuoteFile       FileTypeEnum = 23
	FileTypeEnum_Invoice         FileTypeEnum = 24
	FileTypeEnum_Indemnity       FileTypeEnum = 25
	FileTypeEnum_BenchmarkImport FileTypeEnum = 29
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

// Verified against Misc.json schema on September 26, 2025
type RunDto struct {
	ID            string              `json:"id"`
	Status        *RunStatusEnum      `json:"status,omitempty"`
	Type          *RunTypeEnum        `json:"type,omitempty"`
	TriggerType   *RunTriggerTypeEnum `json:"triggerType,omitempty"`
	FinishedAt    *string             `json:"finishedAt,omitempty"`
	RunNumber     int32               `json:"runNumber"`
	Description   *string             `json:"description,omitempty"`
	FindingIDList []string            `json:"findingIdList,omitempty"`
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

// Project Template Models - Updated to match Project.json schema on September 26, 2025
type ProjectTemplateDtoPagedResultDto struct {
	Items      []*ProjectTemplateDto `json:"items,omitempty"`
	TotalCount int32                 `json:"totalCount"`
}

// Verified against Project.json schema on September 26, 2025
type ProjectTemplateDtoPagedResultDtoAjaxResponse struct {
	TargetUrl           *string                           `json:"targetUrl,omitempty"`
	Success             bool                              `json:"success"`
	Error               *ErrorInfo                        `json:"error,omitempty"`
	UnAuthorizedRequest bool                              `json:"unAuthorizedRequest"`
	Abp                 bool                              `json:"__abp"`
	Result              *ProjectTemplateDtoPagedResultDto `json:"result,omitempty"`
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

// Compliance Norm Template Models - using the verified version above

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
// Verified against Vulnerability.json schema on September 26, 2025
type VulnerabilityTypeDto struct {
	Title     *string `json:"title,omitempty"`
	Code      *string `json:"code,omitempty"`
	IsSystem  bool    `json:"isSystem"`
	IsVisible bool    `json:"isVisible"`
}

// Verified against Vulnerability.json schema on September 26, 2025
type VulnerabilityTypeDtoPagedResultDto struct {
	Items      []*VulnerabilityTypeDto `json:"items,omitempty"`
	TotalCount int32                   `json:"totalCount"`
}

// Verified against Vulnerability.json schema on September 26, 2025
type VulnerabilityTypeDtoPagedResultDtoAjaxResponse struct {
	TargetUrl           *string                             `json:"targetUrl,omitempty"`
	Success             bool                                `json:"success"`
	Error               *ErrorInfo                          `json:"error,omitempty"`
	UnAuthorizedRequest bool                                `json:"unAuthorizedRequest"`
	Abp                 bool                                `json:"__abp"`
	Result              *VulnerabilityTypeDtoPagedResultDto `json:"result,omitempty"`
}

// Team Models
// Verified against Team.json schema on September 26, 2025
type TeamTypeEnum int32

const (
	TeamTypeEnum_Undefined  TeamTypeEnum = 0
	TeamTypeEnum_Pentesters TeamTypeEnum = 1
	TeamTypeEnum_Clients    TeamTypeEnum = 2
)

// Verified against Team.json schema on September 26, 2025
type TeamDto struct {
	ID             string     `json:"id"`
	Name           *string    `json:"name,omitempty"`
	TeamLeaderName *string    `json:"teamLeaderName,omitempty"`
	TeamLeaderId   *string    `json:"teamLeaderId,omitempty"`
	ClientId       *string    `json:"clientId,omitempty"`
	UserList       []*UserDto `json:"userList,omitempty"`
}

// Verified against Team.json schema on September 26, 2025
type TeamDtoAjaxResponse struct {
	TargetUrl           *string    `json:"targetUrl,omitempty"`
	Success             bool       `json:"success"`
	Error               *ErrorInfo `json:"error,omitempty"`
	UnAuthorizedRequest bool       `json:"unAuthorizedRequest"`
	Abp                 bool       `json:"__abp"`
	Result              *TeamDto   `json:"result,omitempty"`
}

// Verified against Team.json schema on September 26, 2025
type TeamDtoPagedResultDto struct {
	Items      []*TeamDto `json:"items,omitempty"`
	TotalCount int32      `json:"totalCount"`
}

// Verified against Team.json schema on September 26, 2025
type TeamDtoPagedResultDtoAjaxResponse struct {
	TargetUrl           *string                `json:"targetUrl,omitempty"`
	Success             bool                   `json:"success"`
	Error               *ErrorInfo             `json:"error,omitempty"`
	UnAuthorizedRequest bool                   `json:"unAuthorizedRequest"`
	Abp                 bool                   `json:"__abp"`
	Result              *TeamDtoPagedResultDto `json:"result,omitempty"`
}
