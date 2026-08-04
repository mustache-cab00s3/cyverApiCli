# Known Issues — v2.2 API Implementation

Identified during systematic audit against `Full_api.json` on April 17, 2026.
These are deferred bugs not addressed in the naming/comment pass. Each issue
is labelled with a severity: **Critical**, **High**, or **Medium**.

---

## Enum Value Mismatches

### KI-001 — `LabelTypeEnum` — All values wrong [Critical]
**File:** `models.go`
**Status:** Fixed August 3, 2026 — values now match Full_api.json (`Finding=0`, `Client=1`, `Project=2`, `Assets=3`, `All=4`).

---

### KI-002 — `ApiRolesEnum` — All Pentester values off-by-one; two values missing [Critical]
**File:** `models.go`
**Status:** Fixed August 3, 2026 — includes `Client_General=3`, `Client=4`, shifted pentester roles, and `Pentester_Team_Manager=11`.

---

### KI-003 — `FormFieldTypeEnum` — Entirely wrong values and names [Critical]
**File:** `models.go`
**Status:** Fixed August 3, 2026 — `Text=1`, `Multitext=2`, `Dropdown=3`, `Multiselect=4`.

---

### KI-004 — `ContinuousProjectVulnerabilityTypeEnum` — Missing value 3 [High]
**File:** `models.go`

`TenableWebAppScan = 3` is defined in the spec but absent from Go.

**Fix:** Add `ContinuousProjectVulnerabilityTypeEnum_TenableWebAppScan = 3`.

---

### KI-005 — `ImportFileTypeEnum` — Missing value 27 [Medium]
**File:** `models.go`
**Status:** Fixed August 3, 2026 — `ImportFileTypeEnum_Horizon = 27`.

---

## Struct Field Mismatches

### KI-006 — `CreateOrUpdateFindingRequest` — Severely incomplete [Critical]
**File:** `models.go`
**Status:** Fixed August 3, 2026 — field set aligned to Full_api.json; `severity` uses `FindingCriticalityEnum`; `findingEvidenceList` uses `FindingEvidenceDto`; `complianceStatus` uses `FindingPciComplianceEnum`.

---

### KI-007 — `CreateClientRequest` — Wrong shape [Critical]
**File:** `models.go`

Go struct has fields that do not exist in the spec, and is missing all spec-defined optional fields.

| Status | Field | Notes |
|---|---|---|
| Missing | `status` | Required by spec (`ClientStatusEnum`) |
| Missing | `clientNumber` | nullable string |
| Missing | `accountManagerId` | nullable uuid |
| Missing | `labelIdList` | array of uuid |
| Missing | `teamIdList` | array of uuid |
| Missing | `clientInformation` | nullable object |
| Extra (not in spec) | `Description` | — |
| Extra (not in spec) | `Email` | — |
| Extra (not in spec) | `Phone` | — |
| Extra (not in spec) | `Address` | — |

---

### KI-008 — `CreateProjectRequestV2` — Missing fields and wrong field name [High]
**File:** `models.go`

| Status | Field | Notes |
|---|---|---|
| Missing | `code` | nullable string |
| Missing | `status` | nullable string |
| Missing | `projectTemplateId` | nullable uuid |
| Wrong name | `labelIds` → should be `labelIdList` | json tag mismatch |
| Extra (not in spec) | `Description` | — |

---

### KI-009 — `AssetDto` and `CreateOrUpdateAssetRequest` — Wrong JSON tag and missing fields [High]
**File:** `models.go`

`OS *string json:"os"` — spec field name is `operatingSystem`. The JSON tag `"os"` will
serialize/deserialize incorrectly.

Both structs are also missing: `documentationLink`, `identifier`, `category`, `group`.
`CreateOrUpdateAssetRequest` additionally missing: `testingFrequency`, `labelIdList`.

---

### KI-010 — `CreateOrUpdateFindingEvidenceRequest` — Missing `assetId` / `evidenceFiles` [High]
**File:** `models.go`
**Status:** Fixed August 3, 2026 — includes `assetId` and `evidenceFiles []string` (file tokens).

---

### KI-011 — `PlanningDateDtoV2` — Entirely wrong fields [High]
**File:** `models.go`
**Status:** Fixed August 3, 2026 — `status`, `startDate`, `endDate`.

---

### KI-012 — `TaskGroupTemplateDto` and `TaskTemplateDto` — Missing fields [Medium]
**File:** `models.go`

`TaskGroupTemplateDto` missing: `code` (nullable string).
`TaskTemplateDto` missing: `code` (nullable string), `externalUrl` (nullable string).

---

### KI-013 — `RefreshTokenResult` — Extra fields not in spec [Medium]
**File:** `models.go`

`RefreshToken` and `TokenType` are present in Go but do not exist in the spec schema.
They will never be populated by the API. Safe to remove.

---

### KI-014 — `FindingDtoPagedResultDtoAjaxResponse.TargetUrl` — Missing `omitempty` [Medium]
**File:** `models.go`
**Status:** Fixed August 3, 2026.

---

## Missing Structs

### KI-015 — `FormDataDto` — Type name used in spec but not defined in Go [Medium]
**File:** `models.go`

`RequestProjectFormDto.formData` references `FormDataDto` in the spec.
Go uses `RequestFormDataDto` instead (a type that does not appear in the spec).
If both types have the same field shapes this may work at runtime, but the name is wrong.

---

### KI-016 — `ClientInformationDto` — Referenced in spec but not defined [Medium]
**File:** `models.go`

`CreateClientRequest.clientInformation` references `ClientInformationDto`. No such struct exists.
Blocked by KI-007 (the parent struct is itself wrong).

---

## Functional Bugs in ops files

### KI-017 — `GetProjectByID` (now `ApiV22ClientProjectsByIdGet`) — Nil response pointer [Critical]
**File:** `client_ops.go` (fixed in renaming pass — verify fix is complete)

Original code passed `nil` as the response pointer to `DoRequest`, then declared the
response variable *after* the call. The function always returned a zero-value struct.
**Status:** Partially addressed during rename rewrite — confirm `DoRequest` now receives `&response`.

---

### KI-018 — `ApiV22PentesterFindingEvidencePost` — Uses non-spec legacy path [High]
**File:** `pentester_ops.go`
**Status:** Marked Deprecated August 3, 2026. Prefer `ApiV22PentesterFindingsByFindingIdEvidencesPost` (POST fixed). Kept for legacy CLI unsupported flows.

---

### KI-019 — Token auth paths missing version prefix [High]
**File:** `token_auth_ops.go`

`ApiTokenauthAuthenticatePost`, `GetUserId`, `ApiTokenauthRefreshtokenPost`, and
`ApiTokenauthSendtwofactorauthcodePost` all use bare `/api/TokenAuth/...` paths.
All other ops files use `/api/v2.2/` as the base. Confirm whether the token auth
endpoints intentionally bypass versioning (some APIs do this for auth routes).

---

### KI-020 — Missing functions: GET evidences list, UPDATE/DELETE client user [High]
**File:** `pentester_ops.go`, `client_ops.go`
**Status (pentester):** Fixed August 3, 2026 — added `ApiV22PentesterFindingsByIdEvidencesGet`.
Still missing (client scope, out of this pass):
- `PUT /api/v2.2/client/users/{id}` → `ApiV22ClientUsersByIdPut`
- `DELETE /api/v2.2/client/users/{id}` → `ApiV22ClientUsersByIdDelete`

Also: `message_ops.go` contains only a package declaration — no functions implemented.

---

### KI-021 — Evidence create used GET instead of POST [Critical]
**File:** `pentester_ops.go`
**Status:** Fixed August 3, 2026 — `ApiV22PentesterFindingsByFindingIdEvidencesPost` now uses HTTP POST.

---

### KI-022 — Upload-file / evidence DTO alignment [High]
**File:** `models.go`, `pentester_ops.go`
**Status:** Fixed August 3, 2026 —
- Added `FileInfoDto` / `FileInfoDtoAjaxResponse`
- `FindingEvidenceDto` includes `assetId` + `evidenceFiles []*FileInfoDto`
- Multipart uploads return typed `FileInfoDtoAjaxResponse`
- Added continuous-project multipart upload helper
- Added latest-report GETs for projects and continuous projects
- Finding GET now sends `includeEvidence` query

---

*End of known issues.*
