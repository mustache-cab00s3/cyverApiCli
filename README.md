# Cyver API CLI

A comprehensive command-line interface tool for interacting with the Cyver API. This CLI provides easy access to manage clients, projects, findings, users, teams, and continuous projects through a well-structured command hierarchy.

## Features

- **Multi-role Support**: Separate command groups for client and pentester operations
- **Flexible Output Formats**: Support for table, JSON, short, and custom output formats
- **Configurable Verbosity**: Multiple verbosity levels for detailed logging (including raw request/response with `-vvv`)
- **Token Management**: Automatic token refresh and secure storage
- **API Version Support**: Support for multiple API versions (v2.2)
- **Interactive Configuration**: Guided setup and configuration management
- **Named profiles (AWS-style)**: Multiple API endpoints and credentials in one YAML file (`profiles`, `current_profile`), like `[default]` and `[profile dev-user]` in `~/.aws/config`; selectable via `--profile`, `CYVER_PROFILE`, with legacy YAML `instances` / `current_instance` (and `CYVER_INSTANCE`) still read for defaults
- **HTTP User-Agent**: Outbound requests use a desktop Chrome-style `User-Agent` with a **CyverCliTool** product token (for example `… Safari/537.36 CyverCliTool/1.0`) so traffic can be distinguished from real browsers while keeping a conventional browser prefix
- **Direct HTTP (`customUrl`)**: Optional `GET`/`POST`/… against a path under the configured base URL or a full URL, using the same auth and `User-Agent` as the main client
- **Comprehensive Error Handling**: Structured error management with retry mechanisms
- **Input Validation**: Robust validation with user-friendly error messages
- **Retry Logic**: Automatic retry for transient failures with exponential backoff
- **Findings Import**: Import findings from JSON or Markdown (Obsidian frontmatter) files with field mapping and enum conversion
- **Evidence Management**: Create evidence attached to findings; import evidence from JSON or Markdown files
- **Pentester Labels**: List and filter labels by type (Finding, Client, Project, Assets, All)
- **Finding templates (web app)**: `templates` commands for finding-library export, listing, downloads, and saves via non-documented portal endpoints. Full detail, warnings, and examples: **[internal/api/services/README.md](internal/api/services/README.md)**

## Installation

### From Source
```bash
go install github.com/yourusername/cyverApiCli@latest
```

### From Repository
```bash
git clone https://github.com/yourusername/cyverApiCli.git
cd cyverApiCli
go mod download
go build
```

### Prerequisites
- Go 1.23.0 or later
- Git

## Quick Start

1. **Initialize Configuration**:
```bash
cyverApiCli config init
```

2. **Authenticate**:
```bash
cyverApiCli apiAuth getToken --username your-email@example.com
```

3. **List Available Commands**:
```bash
cyverApiCli --help
```

## Command Tree

```
cyverApiCli
├── apiAuth                    # API Authentication
│   └── getToken              # Get authentication token
├── customUrl                 # Direct HTTP request (configured base URL or full URL)
├── client                    # Client Operations
│   ├── get-client-info       # Get client information
│   ├── update-client-info    # Update client information
│   ├── list-clients          # List all clients
│   ├── get-projects          # Get client projects
│   ├── get-project-by-id     # Get specific project by ID
│   ├── get-project-request-forms # Get project request forms
│   ├── request-project       # Request a new project
│   ├── get-continuous-projects # Get continuous projects
│   ├── get-continuous-project-by-id # Get specific continuous project
│   ├── get-continuous-project-request-forms # Get continuous project forms
│   ├── request-continuous-project # Request continuous project
│   ├── get-findings          # Get findings
│   ├── get-finding-by-id     # Get specific finding by ID
│   ├── set-finding-status    # Update finding status
│   ├── get-assets            # Get assets
│   ├── create-asset          # Create new asset
│   ├── delete-asset          # Delete asset
│   ├── update-asset          # Update asset
│   ├── get-users             # Get users
│   └── create-user           # Create new user
├── config                    # Configuration Management
│   ├── init                  # Initialize CLI configuration
│   ├── view                  # View current configuration
│   ├── refresh-token         # Refresh access token
│   ├── re-auth               # Re-authenticate
│   ├── update                # Patch config values (flags)
│   └── profile               # Named profiles (alias: instance)
│       ├── list              # List profile names
│       └── use               # Set default profile (current_profile)
├── pentester                 # Pentester Operations
│   ├── clients               # Client Management (Pentester View)
│   │   ├── list              # List pentester clients
│   │   ├── get               # Get specific client
│   │   ├── create            # Create new client
│   │   ├── update            # Update client
│   │   ├── delete            # Delete client
│   │   ├── get-assets        # Get client assets
│   │   ├── create-asset      # Create client asset
│   │   └── update-asset      # Update client asset
│   ├── findings              # Findings Management
│   │   ├── list              # List findings
│   │   ├── get               # Get specific finding
│   │   ├── create            # Create new finding
│   │   ├── update            # Update finding
│   │   ├── delete            # Delete finding
│   │   ├── import            # Import findings (JSON or Markdown)
│   │   ├── custom-import     # Import findings from a folder template (manifest + evidence files)
│   │   └── evidence          # Evidence attached to findings
│   │       ├── create        # Create evidence for a finding
│   │       └── import        # Import evidence from file (JSON or Markdown)
│   ├── labels                # Labels Management
│   │   └── list              # List labels (by type, with pagination)
│   ├── projects              # Projects Management
│   │   ├── list              # List projects
│   │   ├── create            # Create new project
│   │   ├── get               # Get specific project
│   │   ├── delete            # Delete project
│   │   ├── update-status     # Update project status
│   │   ├── set-assets        # Set project assets
│   │   ├── set-users         # Set project users
│   │   ├── set-teams         # Set project teams
│   │   ├── get-checklists    # Get project checklists
│   │   ├── get-compliance-norms # Get compliance norms
│   │   ├── get-report-versions # Get report versions
│   │   ├── get-report        # Get project report
│   │   └── upload-file       # Upload project file
│   ├── users                 # Users Management
│   │   ├── list              # List users
│   │   ├── create            # Create new user
│   │   ├── get               # Get specific user
│   │   ├── update            # Update user
│   │   └── delete            # Delete user
│   ├── teams                 # Teams Management
│   │   ├── list              # List teams
│   │   ├── create            # Create new team
│   │   ├── get               # Get specific team
│   │   ├── update            # Update team
│   │   └── delete            # Delete team
│   └── continuous-projects   # Continuous Projects Management
│       ├── list              # List continuous projects
│       ├── create            # Create continuous project
│       ├── get               # Get specific continuous project
│       ├── delete            # Delete continuous project
│       ├── update-status     # Update continuous project status
│       ├── set-assets        # Set continuous project assets
│       ├── set-users         # Set continuous project users
│       ├── set-teams         # Set continuous project teams
│       ├── get-runs          # Get continuous project runs
│       ├── complete-run      # Complete continuous project run
│       ├── get-report-versions # Get report versions
│       ├── get-report        # Get continuous project report
│       └── upload-file       # Upload file to continuous project
├── templates                 # Finding libraries / templates (non-supported web endpoints — see internal/api/services/README.md)
│   ├── export                # Export templates metadata; optional --download
│   ├── export-json (exj)     # Download all templates in a library as JSON array
│   ├── libraries             # List finding libraries
│   ├── download              # Download temp file from export token
│   ├── save-template         # Create or update a finding template (JSON)
│   └── save-library          # Create or update a finding library (JSON)
└── help                      # Help about any command
```

## Finding templates and non-supported services

The **`templates`** command group uses HTTP paths exposed by the web application, not the versioned public API. Those calls are implemented under `internal/api/services` and may change without notice. For endpoint list, response models, header policy, verbosity logging, and copy-paste CLI examples, see **[internal/api/services/README.md](internal/api/services/README.md)**.

## Configuration

### Initial Setup
The CLI uses a YAML configuration file located at `~/.cyverApiCli.yaml`. Initialize it with:

```bash
cyverApiCli config init
```

If the file **already exists**, you are prompted to **O**verwrite the whole file, **A**dd or replace a single named profile (other settings stay), or **C**ancel. Use **A** to add another environment without losing your current config.

### Configuration Options
- **API Base URL**: The base URL for the Cyver API
- **API Version**: Supported versions (v2.2)
- **Token Management**: Automatic token refresh and storage
- **Output Format**: Default output format (table, JSON, custom)

### Viewing Configuration
```bash
cyverApiCli config view
```

When you use named profiles, `config view` also shows the default profile, the active profile for the current run, and the list of configured names.

### Named profiles (AWS-style, multi-environment)

This mirrors the idea of **`~/.aws/config`**: a **`[default]`** profile and additional **`[profile name]`** sections. In YAML you use a **`profiles`** map and **`current_profile`** (preferred). The legacy **`instances`** map and **`current_instance`** key are still **read** for backward compatibility.

**YAML shape (example):**

```yaml
current_profile: default

profiles:
  default:
    api:
      version: latest
      base_url: https://api.cyver.io
      api_key: ""
    auth:
      email: you@example.com
    token:
      access_token: "..."
      refresh_token: "..."
      expireInSeconds: 3600
      refresh_expires_in: ...
      token_created_at: "..."
      refresh_token_created_at: "..."
  dev-user:
    api:
      version: latest
      base_url: https://other.example.com
      api_key: ""
    auth:
      email: you@example.com
    token: { ... }
```

- **Flat configs** with only top-level `api`, `auth`, and `token` (no `profiles` / `instances`) behave as before.
- **Which profile is used:** **`--profile` / `-p`** overrides for one command. Otherwise the default comes from **`current_profile`** in the file, or **`CYVER_PROFILE`**, with legacy **`current_instance`** / **`CYVER_INSTANCE`** still honored when **`current_profile`** is unset.

The CLI merges the selected profile into the top-level `api`, `auth`, and `token` keys each run. Token and auth updates are stored under the active profile and mirrored to the top-level keys for that session.

**Commands:**

```bash
cyverApiCli config profile list
cyverApiCli config profile use dev-user
# "config instance …" still works as an alias for "config profile …"
```

**Examples:**

```bash
cyverApiCli -p dev-user pentester projects list
export CYVER_PROFILE=dev-user
cyverApiCli pentester projects list
```

**Update values without re-init:**

```bash
cyverApiCli config update --base-url https://api.cyver.io
cyverApiCli config update --for-profile dev-user --base-url https://other.example.com
```

Only flags you pass are written. Useful flags:

- **Profile scope:** `--for-profile <name>`
- **API:** `--api-version`, `--base-url`, `--api-key` (use `--api-key ""` to clear a stored key)
- **Auth:** `--auth-email` (stored email for token re-authentication)
- **Global (not per profile):** `--proxy-url`, `--proxy-user`, `--proxy-password`, `--log-level`, `--log-file`, `--output-format`, `--output-color`, `--client-timeout` (seconds, 1–300)

Run `cyverApiCli config update --help` for the full flag list.

### Outbound HTTP (User-Agent)

All normal API calls and the **`customUrl`** direct HTTP command send a shared **User-Agent** string: a typical **Chrome on Windows** token sequence, followed by **`CyverCliTool/1.0`**, defined in `internal/api/useragent.go`. Servers or proxies can key off `CyverCliTool` for analytics or allowlists without dropping browser-like compatibility.

## Authentication

### Password login and stored tokens

**Password-based** flows (`apiAuth getToken`, `config init` when no API key is set, `config re-auth`) call **`/api/TokenAuth/Authenticate`** (and related **`SendTwoFactorAuthCode`** / **`RefreshToken`**) **without** sending a stored access token from the config file. The CLI still merges profile tokens into viper for normal API calls, but login and refresh must not attach an unrelated JWT (for example from another profile or legacy top-level `token.*` keys) to a different **base URL**, which would otherwise often produce **403** or a generic **“Invalid username or password”** even when the password is correct.

### Getting a Token
```bash
cyverApiCli apiAuth getToken --username your-email@example.com
```

### Token Refresh
```bash
cyverApiCli config refresh-token
```

### Re-authentication
```bash
cyverApiCli config re-auth
```

## Usage Examples

### Client Operations
```bash
# List all clients
cyverApiCli client list-clients

# Get specific project
cyverApiCli client get-project-by-id -p 550e8400-e29b-41d4-a716-446655440000

# Create a new asset
cyverApiCli client create-asset --body '{"name": "Web Server", "type": "server"}'

# Get findings with custom output
cyverApiCli client get-findings -o custom -C 6
```

### Pentester Operations
```bash
# List all projects (pentester view)
cyverApiCli pentester projects list

# Create a new finding (title, type, severity, status)
cyverApiCli pentester findings create -P <uuid> -t "SQL Injection" -s high -y Vulnerability -u draft

# On success, the create command prints the finding URL

# Import findings from JSON
cyverApiCli pentester findings import findings.json -P <uuid>

# Import findings from Markdown/Obsidian frontmatter
cyverApiCli pentester findings import finding.md -P <uuid> -y markdown

# Custom folder import (manifest + evidence files)
cyverApiCli pentester findings custom-import --project-id <uuid> ./import_templates/Finding1

# List labels (default: all types)
cyverApiCli pentester labels list
cyverApiCli pentester labels list --type finding --max-results 20 --filter "web"

# Create evidence for a finding
cyverApiCli pentester findings evidence create -f <uuid> -t "Request capture" -l /login -I 192.168.1.1

# Import evidence from JSON or Markdown
cyverApiCli pentester findings evidence import evidence.json -f <uuid>
cyverApiCli pentester findings evidence import evidence.md -f <uuid> -y markdown

# Update a finding (title, status, severity, recommendation, CWE list, etc.)
cyverApiCli pentester findings update <finding-uuid> -t "New title" -s fixed -r "Apply patch" -C "CWE-89,CWE-564"

# Update project status
cyverApiCli pentester projects update-status -p 550e8400-e29b-41d4-a716-446655440000 --status "in-progress"

# Get team information
cyverApiCli pentester teams get --team-id 6ba7b810-9dad-11d1-80b4-00c04fd430c8
```

## Pentester Findings

### Create Finding

Creates a new finding with the given attributes. The `--title` value is sent to the API as `name`.

**Required:** `--project-id` (`-P`), `--title` (`-t`)

**Flags:**
- `--title` (`-t`) – Finding title (required; sent as `name` to API)
- `--description` (`-d`) – Finding description
- `--type` (`-y`) – `Vulnerability` or `Observation`
- `--severity` (`-s`) – `critical`, `high`, `medium`, `low`, `info` (API scale: info=0, low=1, medium=2, high=3, critical=4)
- `--status` (`-u`) – Default `draft`; other values: pending-fix, fixed, accepted, to-review, reviewed, mitigated, partial-fix, false-positive, raised, reopen, acknowledged, identified
- `--trigger-events` (`-r`) – Whether to trigger events (default: false)

On successful creation (201), the CLI prints the finding URL:  
`{Host}/App/Projects/Details/{project-id}/Finding/{finding-id}`

### Import Findings

Import findings from structured files. Supports **JSON** and **Markdown (Obsidian frontmatter)**.

**Required:** `--project-id` (`-P`), file path

**Flags:**
- `--file-type` (`-y`) – `json` (default) or `markdown` (also accepts `md`, `obsidian`)
- `--trigger-events` (`-T`) – Whether to trigger events (default: false)

**JSON format:** Array of objects or a single object. Fields map to the API; unknown fields are ignored.  
Supported fields include: `name`/`title`, `description`, `type`, `status`, `severity`, `code`, `complianceStatus`, `complianceComment`, `impact`, `impactDescription`, `likelihood`, `likelihoodDescription`, `recommendation`, `backgroundInformation`, `cweList`, `cveList`, `mitreAttackTacticsList`, `mitreAttackTechniquesList`, `assetIdList`, `labelIds`, and others from the CreateOrUpdateFindingRequest schema.

**Markdown (Obsidian) format:** Frontmatter between `---` delimiters.  
- One key-value per line: `key: value`
- Lists use indented lines starting with `-`

Example Markdown frontmatter:
```markdown
---
title: SQL Injection
description: Found in login form
type: Vulnerability
severity: high
status: draft
cweList:
  - CWE-89
  - CWE-564
---
```

Enum strings (`type`, `status`, `severity`) are converted to the API’s numeric values. Extra properties in the file are ignored.

### Custom Import (folder workflow)

Use `custom-import` when your findings are arranged in a folder with a manifest JSON and local evidence files.

**Command:** `cyverApiCli pentester findings custom-import --project-id <uuid> <folder>`

**Default folder inputs:**
- `findings_import.json` (required; manifest)
- `evidence_import.json` (optional extra evidence file for single-finding manifests)

**Batch modes:**
- `--each-subdir` - import each immediate child directory containing the manifest.
- `--batch` (`-b`) - recursively import every directory (including nested) containing the manifest.

**Behavior summary:**
- Creates findings via v2.2 API.
- Uploads evidence files via supported `POST /api/v2.2/pentester/projects/{id}/upload-file` and captures `fileToken`.
- Creates each evidence via supported `POST /api/v2.2/pentester/findings/{findingId}/evidences`, sending uploaded tokens in `evidenceFiles`.
- Updates the finding via supported `PUT /api/v2.2/pentester/findings/{id}` with `findingEvidenceList`.
- Does not use non-supported/legacy app-service endpoints (`CreateOrEditFindingInstance`, `/App/Projects/UploadFindingEvidenceFiles`).
- Writes a post-import snapshot file named `<finding-guid>.json` into the source folder for traceability.
- In batch mode, each discovered finding folder is processed using its own working path, so manifest inputs, evidence file resolution, and snapshot outputs stay local to that folder.

For a concrete template and field examples, see `import_templates/Finding1/README.md`.

### Update Finding

Update an existing finding by ID. **Before applying changes**, the CLI fetches the current finding and saves a JSON backup (unless `--no-backup`). Finding data and the change set are logged for audit. Only provided flags are sent; at least one update field is required. The API uses the CreateOrUpdateFindingRequest schema; `--title` is sent as `name`.

**Command:** `cyverApiCli pentester findings update [finding-id]`

**Required:** finding-id (positional), at least one update flag below

**Backup and logging:**
- By default, the current finding is **fetched** via GET, then a **backup** is written to `finding-backups/finding_<id>_<timestamp>.json` (override with `--backup-dir`).
- Logs include: fetch start, backup path, “Finding data before update” (name, code, projectId, status, severity), “Applying changes” (fields and full change body), and “Successfully updated finding” with field count.
- Use `--no-backup` to skip the GET and backup (not recommended for production).

**All update fields:**

| Flag | Description | Values / format |
|------|-------------|-----------------|
| **Basic** | | |
| `--title` | Finding title (API: name) | string (max 250) |
| `--description` | Finding description | string (max 10000) |
| `--code` | Finding code | string (e.g. F-2025-3508, max 100) |
| `--type` | Finding type | vulnerability, nonconformity, observation, incident, risk — or 1, 2, 4, 8, 16 |
| `--status` | Finding status | draft(1), pending-fix(2), fixed(3), ready-retest(4), accepted(5), to-review(6), reviewed(7), mitigated(8), partial-fix(9), false-positive(10), raised(11), reopen(12), acknowledged(13), identified(14) — or 1–14 |
| `--severity` | Finding severity | info(0), low(1), medium(2), high(3), critical(4) — or 0–4 |
| **Compliance** | | |
| `--compliance-status` | PCI compliance status | pass(0), fail(1) — or 0, 1 |
| `--compliance-comment` | Compliance comment | string |
| **Risk / impact** | | |
| `--impact` | Impact score | integer 0–5 (omit flag to leave unchanged) |
| `--impact-description` | Impact description | string |
| `--likelihood` | Likelihood score | integer 0–5 (omit flag to leave unchanged) |
| `--likelihood-description` | Likelihood description | string |
| **Remediation** | | |
| `--recommendation` | Remediation recommendation | string |
| `--background-information` | Background information | string |
| **Assignment** | | |
| `--reviewer-id` | Reviewer user ID | UUID |
| `--project-task-id` | Project task ID | UUID |
| **CVSS** | | |
| `--cvss-json` | CVSS object | JSON, e.g. `{"cvss31Vector":"CVSS:3.1/...", "cvss31Score":5.8}` or cvss20/cvss30/cvss40 fields |
| **Lists (comma-separated)** | | |
| `--cwe-list` | CWE IDs | e.g. CWE-89,CWE-564 |
| `--cve-list` | CVE IDs | e.g. CVE-2021-44228 |
| `--mitre-tactics` | MITRE ATT&CK tactic IDs | e.g. TA0011 |
| `--mitre-techniques` | MITRE ATT&CK technique IDs | e.g. T1001 |
| `--mitre-mitigations` | MITRE ATT&CK mitigation IDs | e.g. M1031 |
| `--vulnerability-type-list` | Vulnerability types | e.g. BypassSomething |
| `--asset-id-list` | Asset UUIDs | comma-separated UUIDs |
| `--label-id-list` | Label UUIDs | comma-separated UUIDs |
| `--project-control-id-list` | Project control UUIDs | comma-separated UUIDs |
| **External URLs** | | |
| `--external-url-json` | External URLs | JSON array of `{title, link}`, e.g. `[{"title":"test","link":"https://example.com"}]` |
| **Backup / audit** | | |
| `--backup-dir` | Backup directory | default: `finding-backups` |
| `--no-backup` | Skip fetch and backup | flag (no value) |
| **Other** | | |
| `--trigger-events` | Trigger events on update | flag (default: false) |

**Examples:**
```bash
# Basic: title, status, severity
cyverApiCli pentester findings update cfce90c1-1ef1-4307-9796-f88ab4f694e8 -t "Updated title" -s fixed -S high

# Remediation and CWE, impact
cyverApiCli pentester findings update <uuid> -r "Apply patch" -C "CWE-89,CWE-564" --impact 3

# CVSS, external URLs, and project controls
cyverApiCli pentester findings update <uuid> -j "{\"cvss31Vector\":\"CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:C/C:N/I:L/A:N\", \"cvss31Score\":5.8}" -e "[{\"title\":\"Ref\",\"link\":\"https://example.com\"}]" -P "uuid1,uuid2"
```

## Pentester Findings Evidence

Create and import evidence attached to findings (CreateOrEditFindingInstance API).

### Create Evidence

**Command:** `cyverApiCli pentester findings evidence create`

**Required:** `--finding-id` (`-f`), `--title` (`-t`)

**Flags:**
- `--finding-id` (`-f`) – Finding ID to attach evidence to (required)
- `--title` (`-t`) – Evidence title (required)
- `--asset` (`-a`), `--location` (`-l`), `--version` (`-W`)
- `--ip` (`-I`), `--hostname` (`-H`), `--port` (`-P`), `--protocol` (`-r`)
- `--issue-details` (`-d`), `--reproduce` (`-R`), `--evidence` (`-e`)
- `--visible-in-report` (`-V`) – Include in report (default: true)

### Import Evidence

Import evidence from a structured file and attach all records to one finding.

**Command:** `cyverApiCli pentester findings evidence import [file-path]`

**Required:** `--finding-id` (`-f`), file path

**Flags:**
- `--finding-id` (`-f`) – Finding ID to attach evidence to (required)
- `--file-type` (`-y`) – `json` (default) or `markdown` (also `md`, `obsidian`)

**JSON:** Array of evidence objects or a single object. Each record must have `title`. Other supported fields: `asset`, `location`, `version`, `ip`, `hostname`, `port`, `protocol`, `issueDetails` / `issue-details`, `reproduce`, `evidence`, `visibleInReport` / `visible-in-report`. Unknown fields are ignored.

**Markdown (Obsidian):** One evidence per file. Frontmatter between `---` with `key: value`; same field names as above.

Example Markdown:
```markdown
---
title: Login request capture
location: /login
ip: 192.168.1.1
hostname: target.example.com
port: "443"
protocol: HTTPS
issue-details: Credentials in cleartext
reproduce: "1. Intercept traffic\n2. Submit form"
evidence: Base64 request dump...
visible-in-report: true
---
```

## Pentester Labels

List labels with optional filters and pagination.

**Command:** `cyverApiCli pentester labels list`

**Flags:**
- `--type` – Label type: `finding`/`0`, `client`/`1`, `project`/`2`, `assets`/`3`, `all`/`4` (default: `all`, i.e. 4)
- `--max-results` – Max results (default: 10)
- `--skip-count` – Number to skip (default: 0)
- `--filter` – Text filter
- `--output` – `json`, `table`, or `custom` (default: table)
- `--max-columns` – For custom table (default: 4)

## Output Formats

The CLI supports multiple output formats:

- **table**: Human-readable table format (default)
- **json**: Complete JSON response
- **short**: Simplified JSON with ID and name only
- **custom**: Interactive field selection for table output

### Output Format Examples
```bash
# Table output (default)
cyverApiCli client get-projects

# Full JSON output
cyverApiCli client get-projects --output json

# Custom table with specific columns
cyverApiCli client get-projects --output custom --max-columns 4
```

## Verbosity Levels

Control the amount of logging output:

- **-v**: Show basic request information (method, URL)
- **-vv**: Show request headers and basic response info
- **-vvv**: Show full request/response details including body

### Verbosity Examples
```bash
# Basic verbosity
cyverApiCli -v client get-projects

# High verbosity for debugging
cyverApiCli -vvv pentester findings list
```

## Global Flags

- `-c, --config`: Specify config file path
- `-p, --profile`: Use a named profile for this run (overrides `CYVER_PROFILE` / `current_profile`)
- `-v, --verbose`: Increase verbosity level (can be used multiple times)
- `-h, --help`: Show help information

**Shorthand collisions:** `-p`, `-c`, and `-v` are reserved at the root for profile, config, and verbosity. Subcommands avoid reusing those letters for their own shorthands (for example pentester/client findings use `-P` for `--project-id`). Cobra only supports single-character shorthands, not multi-letter ones.

**Environment:** `CYVER_PROFILE` is bound to `current_profile`. `CYVER_INSTANCE` is bound to legacy `current_instance`. Precedence for the active profile: **`--profile`**, then `current_profile` (including `CYVER_PROFILE`), then `current_instance` (including `CYVER_INSTANCE`).

## Recent Updates and Changes

### Configuration (AWS-style profiles)

- **`profiles` map** in `~/.cyverApiCli.yaml`: each profile name holds `api`, and optionally `auth` and `token` (like `[default]` / `[profile dev-user]` in AWS config).
- **`current_profile`**: default profile; set with `cyverApiCli config profile use <name>`.
- **Legacy `instances` / `current_instance`**: still read; new writes prefer `profiles` / `current_profile`.
- **`config profile list`** / **`config profile use`** (alias **`config instance`**).
- **Global `--profile` / `-p`**.
- **`CYVER_PROFILE`** / **`CYVER_INSTANCE`** for default profile selection.
- **`config update`**: patch YAML via flags without re-running `config init` (`--for-profile`, `--base-url`, `--api-key`, `--api-version`, `--auth-email`, proxy, logging, output, client timeout). See the **Update values without re-init** subsection under **Named profiles** and `config update --help`.

### Findings

- **Create**
  - `--title` is sent as `name` in the API request body.
  - Added `--type` with values `Vulnerability` or `Observation`.
  - `--status` defaults to `draft`; supports all API status values.
  - `--severity` limited to `critical`, `high`, `medium`, `low`, `info`. API uses 0–4 (info=0, low=1, medium=2, high=3, critical=4).
  - On 201 success, prints the finding URL using the `result` field from the response.

- **Import**
  - New structured file import for creating findings (replaces simple file-path upload).
  - **File types:** `--file-type json` (default) or `--file-type markdown` (or `md` / `obsidian`).
  - **JSON:** Accepts an array of finding objects or a single object. Unknown fields are ignored; known fields are mapped to the API (e.g. `title` → `name`). Enums (`type`, `status`, `severity`) can be strings or numbers.
  - **Markdown:** Parses Obsidian-style frontmatter between `---` delimiters. Supports `key: value` and indented list items with `-`. One finding per file.
  - Requires `--project-id` (`-P`). Each finding is created via the pentester findings create API; import reports success/failure counts and prints finding URLs for created items.

- **Update**
  - Update finding uses `name` (not `title`) in the request body when changing the title.
  - **Backup and logging:** Before applying changes, the CLI fetches the current finding and saves a JSON backup to `--backup-dir` (default: `finding-backups`). Finding data (name, code, projectId, status, severity) and the full change set are logged. Use `--no-backup` to skip fetch and backup.
  - **Update fields:** All supported flags are documented in the **Update Finding** section (table and examples): basic (title, description, code, type, status, severity), compliance, risk/impact, remediation, assignment, CVSS, comma-separated lists (cwe, cve, MITRE, vulnerability-type, asset-id, label-id, project-control-id), external-url-json, backup options, and trigger-events. At least one update field is required.

### Pentester Findings Evidence

- **Create:** `pentester findings evidence create` – Create evidence attached to a finding. Required: `--finding-id` (`-f`), `--title` (`-t`). Optional: `--asset` (`-a`), `--location` (`-l`), `--version` (`-W`), `--ip` (`-I`), `--hostname` (`-H`), `--port` (`-P`), `--protocol` (`-r`), `--issue-details` (`-d`), `--reproduce` (`-R`), `--evidence` (`-e`), `--visible-in-report` (`-V`, default true). Uses the CreateOrEditFindingInstance app service endpoint.
- **Import:** `pentester findings evidence import [file-path]` – Import evidence from JSON or Markdown. Requires `--finding-id` (`-f`). `--file-type` (`-y`) accepts `json` (default) or `markdown`. JSON: array or single object; each record needs `title`. Markdown: one evidence per file via Obsidian frontmatter. Reports success/failure counts.

### Pentester Labels

- New `pentester labels` command group with `list` subcommand.
- List labels with `--type`, `--max-results`, `--skip-count`, `--filter`.
- Type values: `finding` (0), `client` (1), `project` (2), `assets` (3), `all` (4). Default is `all`.
- Output formats: `json`, `table`, `custom`.

### General

- **Verbosity:** With `-vvv`, raw HTTP request and response (including body) are printed to stderr. Avoid `-vvv` when typing secrets; bodies and debug lines may contain sensitive data.
- **Auth:** MFA / 2FA runs only when the API returns `requiresTwoFactorVerification: true` on the authenticate result (password-only sandboxes are not treated as MFA). The send-code step uses `twoFactorAuthProviders` from the API when present, otherwise Google Authenticator. Re-authentication paths treat a missing authenticate `result` as an error instead of assuming MFA.
- **HTTP User-Agent:** Chrome-style desktop string plus **`CyverCliTool/1.0`** on outbound API requests for tracking.
- **TokenAuth HTTP headers:** `Authenticate`, `SendTwoFactorAuthCode`, and `RefreshToken` do **not** receive a `Authorization: Bearer` header from stored config, so switching profiles or hosts does not accidentally send another environment’s JWT on login.
- **Project details:** `pentester projects get` correctly shows project data from the API’s `Result` field.

## Development

### Prerequisites
- Go 1.23.0 or later
- Git

### Building from Source
```bash
# Clone the repository
git clone https://github.com/yourusername/cyverApiCli.git
cd cyverApiCli

# Install dependencies
go mod download

# Build the project
go build

# Run tests
go test ./...
```

### Project Structure
```
cyverApiCli/
├── cmd/                    # Command definitions
│   ├── client/           # Client-specific commands
│   ├── pentester/         # Pentester-specific commands
│   ├── shared/            # Shared utilities
│   ├── error_handler.go   # Error handling utilities
│   ├── config.go          # Configuration commands
│   └── root.go            # Root command
├── internal/              # Internal packages
│   ├── api/              # API client implementations
│   │   ├── services/   # Non-documented web-app endpoints (`README.md` documents behavior & `templates` CLI)
│   │   └── versions/     # API version implementations
│   ├── config/           # Configuration management
│   └── errors/           # Error handling system
├── logger/               # Logging utilities
├── output/               # Output formatting
├── docs/                 # Documentation
└── main.go              # Application entry point
```

### Error Handling System
The project includes a comprehensive error handling system located in `internal/errors/`:

- **`errors.go`**: Core error types and utilities
- **`retry.go`**: Retry mechanisms with exponential backoff
- **`validation.go`**: Input validation utilities

### Key Dependencies
- **Cobra**: CLI framework
- **Viper**: Configuration management
- **Zerolog**: Structured logging
- **go-pretty**: Table formatting

## Troubleshooting

### Common Issues

1. **Authentication Errors**
   - Ensure your credentials are correct
   - Check if 2FA is enabled and provide the code
   - Verify your account has the necessary permissions
   - If password login to a **new base URL** or profile failed before an update: an old access token in the file could have been sent on `TokenAuth/Authenticate`; current builds omit that. On an older binary, clear or ignore conflicting top-level `token.*` keys or upgrade the CLI

2. **Configuration Issues**
   - Run `cyverApiCli config init` to recreate configuration
   - Check file permissions on `~/.cyverApiCli.yaml`
   - Verify API base URL is correct
   - If you use named profiles, confirm `profiles.<name>` (or legacy `instances.<name>`) exists and `current_profile` / `current_instance` is set (or pass `-p` / `CYVER_PROFILE`). Run `cyverApiCli config profile list` to see names

3. **Token Expiration**
   - Use `cyverApiCli config refresh-token` to refresh
   - Re-authenticate with `cyverApiCli config re-auth`

4. **Output Format Issues**
   - Use `--output json` for complete JSON responses
   - Use `--output custom` for interactive field selection
   - Adjust `--max-columns` for table formatting

### Getting Help
```bash
# General help
cyverApiCli --help

# Command-specific help
cyverApiCli client --help
cyverApiCli pentester projects --help

# Verbose output for debugging
cyverApiCli -vvv [command]
```

## Error Handling

The CLI includes a comprehensive error handling system with:

### Error Types and Codes
- **Configuration Errors**: `CONFIG_MISSING`, `CONFIG_INVALID`, `CONFIG_FILE_NOT_FOUND`
- **Authentication Errors**: `AUTH_FAILED`, `TOKEN_EXPIRED`, `TOKEN_INVALID`, `CREDENTIALS_INVALID`
- **API Errors**: `API_UNAUTHORIZED`, `API_FORBIDDEN`, `API_NOT_FOUND`, `API_RATE_LIMITED`, `API_SERVER_ERROR`
- **Validation Errors**: `VALIDATION_FAILED`, `INVALID_INPUT`, `MISSING_REQUIRED`
- **Internal Errors**: `INTERNAL_ERROR`, `NOT_IMPLEMENTED`, `UNEXPECTED_TYPE`

### Error Severity Levels
- **Low**: Warnings that don't prevent execution
- **Medium**: Errors that prevent command execution
- **High**: Critical errors that may affect system stability
- **Critical**: Fatal errors that require immediate attention

### Key Features
- **Structured Error Types**: Standardized error codes and severity levels
- **Automatic Retry**: Built-in retry logic for transient failures with exponential backoff
- **Input Validation**: Comprehensive validation with clear error messages
- **User-Friendly Messages**: Clear, actionable error messages for users
- **Logging Integration**: Detailed logging for debugging and monitoring
- **Exit Code Mapping**: Appropriate exit codes based on error severity

### Error Handling Examples

#### Basic Error Handling
```go
// Validate input parameters
if maxResultCount < 0 {
    shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeValidationFailed, "max-results must be non-negative", nil))
    return
}

// Handle API errors
projects, err := client.ClientOps.GetProjects(status, maxResultCount, skipCount, filter)
if err != nil {
    shared.HandleError(cmd, err)
    return
}
```

#### Input Validation
```go
// Validate required parameters
if bodyJSON == "" {
    shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeValidationFailed, "body is required", nil))
    return
}

// Validate JSON parsing
var body interface{}
if err := json.Unmarshal([]byte(bodyJSON), &body); err != nil {
    shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeValidationFailed, "invalid JSON body", err))
    return
}
```

#### Output Format Validation
```go
validFormats := []string{"json", "short", "table"}
isValidFormat := false
for _, format := range validFormats {
    if outputFormat == format {
        isValidFormat = true
        break
    }
}

if !isValidFormat {
    shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeValidationFailed, 
        fmt.Sprintf("Invalid output format '%s'. Valid options are: %s", outputFormat, strings.Join(validFormats, ", ")), nil))
    return
}
```

#### Error Types and Codes

The CLI uses structured error codes for better error handling:

**Validation Errors**:
- `VALIDATION_FAILED`: Input validation failed
- `INVALID_INPUT`: Invalid input format
- `MISSING_REQUIRED`: Required parameter missing

**API Errors**:
- `API_UNAUTHORIZED`: Authentication failed
- `API_FORBIDDEN`: Access denied
- `API_NOT_FOUND`: Resource not found
- `API_RATE_LIMITED`: Rate limit exceeded

**Configuration Errors**:
- `CONFIG_INVALID`: Invalid configuration
- `CONFIG_MISSING`: Configuration missing

**Internal Errors**:
- `INTERNAL_ERROR`: Unexpected internal error
- `NOT_IMPLEMENTED`: Feature not implemented
- `UNEXPECTED_TYPE`: Unexpected type error

#### Error Severity Levels

- **Low**: Warnings that don't prevent execution
- **Medium**: Errors that prevent command execution
- **High**: Critical errors that may affect system stability
- **Critical**: Fatal errors that require immediate attention

#### Best Practices

1. **Always use `shared.HandleError`** for consistent error handling
2. **Provide context** in error messages for better user experience
3. **Use appropriate error codes** for different error types
4. **Handle errors immediately** after they occur
5. **Validate input** before making API calls
6. **Test error scenarios** to ensure proper error handling

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

MIT License - see LICENSE file for details 