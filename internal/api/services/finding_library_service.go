package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/yourusername/cyverApiCli/internal/api"
)

func (c *NonSupportedServiceClient) GetFindingTemplatesToExport(ctx context.Context, findingLibraryID string) ([]byte, int, error) {
	if strings.TrimSpace(findingLibraryID) == "" {
		return nil, 0, fmt.Errorf("findingLibraryId is required")
	}

	query := url.Values{}
	query.Set("findingLibraryId", findingLibraryID)
	return c.DoGET(ctx, "/api/services/app/FindingLibrary/GetFindingTemplatesToExport", query)
}

// GetFindingLibrariesRaw calls the non-supported FindingLibrary listing endpoint.
func (c *NonSupportedServiceClient) GetFindingLibrariesRaw(ctx context.Context, filter string, maxResultCount, skipCount int) ([]byte, int, error) {
	if maxResultCount <= 0 {
		maxResultCount = 10
	}
	if skipCount < 0 {
		skipCount = 0
	}

	query := url.Values{}
	query.Set("filter", filter)
	query.Set("maxResultCount", fmt.Sprintf("%d", maxResultCount))
	query.Set("skipCount", fmt.Sprintf("%d", skipCount))
	return c.DoGET(ctx, "/api/services/app/FindingLibrary/GetFindingLibraries", query)
}

// GetFindingLibraries returns a typed ABP response for finding library list requests.
func (c *NonSupportedServiceClient) GetFindingLibraries(ctx context.Context, filter string, maxResultCount, skipCount int) (*FindingLibrariesResponse, []byte, int, error) {
	body, statusCode, err := c.GetFindingLibrariesRaw(ctx, filter, maxResultCount, skipCount)
	if err != nil {
		return nil, body, statusCode, err
	}

	var out FindingLibrariesResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, body, statusCode, fmt.Errorf("decode GetFindingLibraries response: %w", err)
	}
	if !out.Success {
		if out.Error != nil && strings.TrimSpace(out.Error.Message) != "" {
			return &out, body, statusCode, fmt.Errorf("service response unsuccessful: %s", out.Error.Message)
		}
		return &out, body, statusCode, fmt.Errorf("service response unsuccessful")
	}

	return &out, body, statusCode, nil
}

// GetFindingTemplatesRaw calls the non-supported FindingTemplate list endpoint.
func (c *NonSupportedServiceClient) GetFindingTemplatesRaw(ctx context.Context, findingLibraryID, filter, status string, maxResultCount, skipCount int) ([]byte, int, error) {
	if strings.TrimSpace(findingLibraryID) == "" {
		return nil, 0, fmt.Errorf("findingLibraryId is required")
	}
	if maxResultCount <= 0 {
		maxResultCount = 10
	}
	if skipCount < 0 {
		skipCount = 0
	}

	query := url.Values{}
	query.Set("findingLibraryId", findingLibraryID)
	query.Set("filter", filter)
	query.Set("status", status)
	query.Set("maxResultCount", fmt.Sprintf("%d", maxResultCount))
	query.Set("skipCount", fmt.Sprintf("%d", skipCount))
	return c.DoGET(ctx, "/api/services/app/FindingLibrary/GetFindingTemplates", query)
}

// GetFindingTemplates returns a typed ABP response for finding template list requests.
func (c *NonSupportedServiceClient) GetFindingTemplates(ctx context.Context, findingLibraryID, filter, status string, maxResultCount, skipCount int) (*FindingTemplatesResponse, []byte, int, error) {
	body, statusCode, err := c.GetFindingTemplatesRaw(ctx, findingLibraryID, filter, status, maxResultCount, skipCount)
	if err != nil {
		return nil, body, statusCode, err
	}

	var out FindingTemplatesResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, body, statusCode, fmt.Errorf("decode GetFindingTemplates response: %w", err)
	}
	if !out.Success {
		if out.Error != nil && strings.TrimSpace(out.Error.Message) != "" {
			return &out, body, statusCode, fmt.Errorf("service response unsuccessful: %s", out.Error.Message)
		}
		return &out, body, statusCode, fmt.Errorf("service response unsuccessful")
	}

	return &out, body, statusCode, nil
}

// CreateOrUpdateFindingTemplateRaw calls the non-supported create/update template endpoint.
func (c *NonSupportedServiceClient) CreateOrUpdateFindingTemplateRaw(ctx context.Context, payload interface{}) ([]byte, int, error) {
	if payload == nil {
		return nil, 0, fmt.Errorf("payload is required")
	}
	return c.DoJSON(ctx, http.MethodPost, "/api/services/app/FindingLibrary/CreateOrUpdateFindingTemplate", nil, payload)
}

// CreateOrUpdateFindingTemplate calls the endpoint and decodes the ABP response envelope.
func (c *NonSupportedServiceClient) CreateOrUpdateFindingTemplate(ctx context.Context, payload interface{}) (*CreateOrUpdateFindingTemplateResponse, []byte, int, error) {
	body, statusCode, err := c.CreateOrUpdateFindingTemplateRaw(ctx, payload)
	if err != nil {
		return nil, body, statusCode, err
	}

	var out CreateOrUpdateFindingTemplateResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, body, statusCode, fmt.Errorf("decode CreateOrUpdateFindingTemplate response: %w", err)
	}
	if !out.Success {
		if out.Error != nil && strings.TrimSpace(out.Error.Message) != "" {
			return &out, body, statusCode, fmt.Errorf("service response unsuccessful: %s", out.Error.Message)
		}
		return &out, body, statusCode, fmt.Errorf("service response unsuccessful")
	}

	return &out, body, statusCode, nil
}

// CreateOrUpdateFindingLibraryRaw calls the non-supported create/update library endpoint.
func (c *NonSupportedServiceClient) CreateOrUpdateFindingLibraryRaw(ctx context.Context, payload interface{}) ([]byte, int, error) {
	if payload == nil {
		return nil, 0, fmt.Errorf("payload is required")
	}
	return c.DoJSON(ctx, http.MethodPost, "/api/services/app/FindingLibrary/CreateOrUpdateFindingLibrary", nil, payload)
}

// CreateOrUpdateFindingLibrary calls the endpoint and decodes the ABP response envelope.
func (c *NonSupportedServiceClient) CreateOrUpdateFindingLibrary(ctx context.Context, payload interface{}) (*CreateOrUpdateFindingLibraryResponse, []byte, int, error) {
	body, statusCode, err := c.CreateOrUpdateFindingLibraryRaw(ctx, payload)
	if err != nil {
		return nil, body, statusCode, err
	}

	var out CreateOrUpdateFindingLibraryResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, body, statusCode, fmt.Errorf("decode CreateOrUpdateFindingLibrary response: %w", err)
	}
	if !out.Success {
		if out.Error != nil && strings.TrimSpace(out.Error.Message) != "" {
			return &out, body, statusCode, fmt.Errorf("service response unsuccessful: %s", out.Error.Message)
		}
		return &out, body, statusCode, fmt.Errorf("service response unsuccessful")
	}

	return &out, body, statusCode, nil
}

// GetFindingTemplatesToExportModel returns a typed ABP response and keeps raw bytes for diagnostics.
func (c *NonSupportedServiceClient) GetFindingTemplatesToExportModel(ctx context.Context, findingLibraryID string) (*FindingTemplatesExportResponse, []byte, int, error) {
	body, statusCode, err := c.GetFindingTemplatesToExport(ctx, findingLibraryID)
	if err != nil {
		return nil, body, statusCode, err
	}

	var out FindingTemplatesExportResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, body, statusCode, fmt.Errorf("decode GetFindingTemplatesToExport response: %w", err)
	}
	if !out.Success {
		if out.Error != nil && strings.TrimSpace(out.Error.Message) != "" {
			return &out, body, statusCode, fmt.Errorf("service response unsuccessful: %s", out.Error.Message)
		}
		return &out, body, statusCode, fmt.Errorf("service response unsuccessful")
	}
	if out.Result == nil || strings.TrimSpace(out.Result.FileToken) == "" {
		return &out, body, statusCode, fmt.Errorf("service response missing result.fileToken")
	}

	return &out, body, statusCode, nil
}

var contentDispositionFilenameRE = regexp.MustCompile(`(?i)filename="?([^";]+)"?`)

// DownloadTempFile downloads the temporary file content from a non-supported web endpoint.
func (c *NonSupportedServiceClient) DownloadTempFile(ctx context.Context, reqData TempFileDownloadRequest) (*TempFileDownloadResult, int, error) {
	bearer, err := c.requireBearerToken()
	if err != nil {
		return nil, 0, err
	}
	if strings.TrimSpace(reqData.FileToken) == "" {
		return nil, 0, fmt.Errorf("fileToken is required")
	}
	if strings.TrimSpace(reqData.FileName) == "" {
		return nil, 0, fmt.Errorf("fileName is required")
	}
	if strings.TrimSpace(reqData.FileType) == "" {
		return nil, 0, fmt.Errorf("fileType is required")
	}

	query := url.Values{}
	query.Set("fileType", reqData.FileType)
	query.Set("fileToken", reqData.FileToken)
	query.Set("fileName", reqData.FileName)

	reqURL := c.BaseURL + "/File/DownloadTempFile?" + query.Encode()
	getLogger().Info("Making non-supported service request", "method", http.MethodGet, "url", reqURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, 0, fmt.Errorf("create download request: %w", err)
	}

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("User-Agent", api.ChromeUserAgent)
	req.Header.Set("Authorization", "Bearer "+bearer)
	logServiceRequest(req)

	start := time.Now()
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("execute download request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("read download response body: %w", err)
	}
	logServiceResponse(resp, body, time.Since(start))
	if resp.StatusCode >= 400 {
		return nil, resp.StatusCode, fmt.Errorf("download request failed with status %d: %s", resp.StatusCode, string(body))
	}

	fileName := reqData.FileName
	if cd := resp.Header.Get("Content-Disposition"); cd != "" {
		if m := contentDispositionFilenameRE.FindStringSubmatch(cd); len(m) > 1 && strings.TrimSpace(m[1]) != "" {
			fileName = strings.TrimSpace(m[1])
		}
	}
	fileName = filepath.Base(fileName)

	return &TempFileDownloadResult{
		FileName:        fileName,
		FileType:        reqData.FileType,
		ContentType:     resp.Header.Get("Content-Type"),
		ContentLength:   len(body),
		ContentEncoding: resp.Header.Get("Content-Encoding"),
		Body:            body,
	}, resp.StatusCode, nil
}
