# Services (Non-Supported API Requests)

This directory contains service clients and request models for **non-documented web application endpoints**.

## Important

- These requests are **not part of the supported/public API contract**.
- Use them **only for specific instances** where the endpoint behavior is known and validated.
- Endpoint shapes and behavior can change without notice between environments/releases.
- Prefer documented versioned endpoints under `internal/api/versions` whenever possible.

## Current Capabilities

Implemented in `finding_library_service.go` and `models.go`:

- `GetFindingLibraries`
  - Calls `/api/services/app/FindingLibrary/GetFindingLibraries` with:
  - `filter`
  - `maxResultCount`
  - `skipCount`

- `GetFindingTemplates`
  - Calls `/api/services/app/FindingLibrary/GetFindingTemplates` with:
  - `findingLibraryId`
  - `filter`
  - `status`
  - `maxResultCount`
  - `skipCount`

- `GetFindingTemplatesToExport`
  - Calls `/api/services/app/FindingLibrary/GetFindingTemplatesToExport` with:
  - `findingLibraryId`

- `CreateOrUpdateFindingTemplate`
  - Calls `/api/services/app/FindingLibrary/CreateOrUpdateFindingTemplate`
  - Accepts a JSON payload object (create/update body)
  - Supports raw and typed ABP-envelope response handling

- `DownloadTempFile`
  - Calls `/File/DownloadTempFile` with:
  - `fileType`
  - `fileToken`
  - `fileName`
  - Returns file bytes and metadata (`content-type`, `content-length`, resolved filename)

- `GetZendeskCompose`
  - Calls external compose endpoint:
    - `https://ekr.zdassets.com/compose/{composeId}`
  - Returns response body and metadata (`status`, `content-type`, `content-length`)

## Models Added

- ABP-style response envelopes:
  - `FindingLibrariesResponse`
  - `FindingTemplatesResponse`
  - `FindingTemplatesExportResponse`
  - `CreateOrUpdateFindingTemplateResponse`
  - `ServiceErrorResponse`

- Finding library list models:
  - `FindingLibraryListResult`
  - `FindingLibraryDto` (`guid`, `name`, `description`, `status`, `findingLibraryTemplateStatus`, optional finding-fields template refs)

- Finding template list models:
  - `FindingTemplateListResult` (`totalCount`, `items`)
  - `FindingTemplateDto` (expanded template detail model from observed response payloads)
  - `FindingTemplateExternalURL`
  - `FindingTemplateCustomField`

- Export/download models:
  - `FindingTemplatesExportFile`
  - `TempFileDownloadRequest`
  - `TempFileDownloadResult`
  - `ExternalComposeResponse`

## Request / Response Logging

Services now support verbosity-driven HTTP logging similar to `v2_2`:

- `-v`: high-level request lifecycle logs
- `-vv`: structured request/response metadata logs
- `-vvv`: raw request and response dumps to stderr

Configured through `services.SetVerboseLevel(...)` from root command pre-run.

## Auth and Header Policy

Service requests now require:

- `Authorization: Bearer <token>` (required; requests fail early if missing)

Service requests include only essential default headers:

- `Authorization: Bearer <token>`
- `Accept`
- `Content-Type` (JSON requests)
- Shared browser-like `User-Agent` (`Chrome... CyverCliTool/1.0`)

## CLI Examples

Examples use the `templates` command group that calls this services layer.

### List finding libraries

```bash
cyverApiCli templates libraries --filter "" --max-result-count 10 --skip-count 0
```

### List finding templates in a library

Service support exists (`GetFindingTemplates`), but a dedicated `templates` subcommand is not yet wired.
You can currently call it via the generic request path if needed.

### Export templates and auto-download

```bash
cyverApiCli templates export --finding-library-id b75ad1c3-f76e-4225-b6f0-deecd36bcc89 --download --output ./Reviewed-Findings.xlsx
```

### Export templates as JSON array

```bash
# Full command
cyverApiCli templates export-json --finding-library-id b75ad1c3-f76e-4225-b6f0-deecd36bcc89 --output ./finding_templates.json

# Short alias command
cyverApiCli templates exj -l b75ad1c3-f76e-4225-b6f0-deecd36bcc89 -o ./finding_templates.json

# Alternative through export command
cyverApiCli templates export -l b75ad1c3-f76e-4225-b6f0-deecd36bcc89 --download --download-type json -o ./finding_templates.json
```

### Download from known token/name/type

```bash
cyverApiCli templates download --file-token <token> --file-name <name.xlsx> --file-type application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
```

### Save a finding template from file

```bash
cyverApiCli templates save-template --data-file ./finding-template.json
```

### Save a finding template with inline JSON

```bash
cyverApiCli templates save-template --data-json "{\"Guid\":\"...\",\"FindingLibraryId\":\"...\",\"Name\":\"Example\"}"
```

### Save a finding library from file

```bash
cyverApiCli templates save-library --data-file ./finding-library.json
```