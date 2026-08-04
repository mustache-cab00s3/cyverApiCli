package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/yourusername/cyverApiCli/internal/api"
	"golang.org/x/net/html"
)

const (
	findingCreateOrEditInstancePath   = "/api/services/app/Finding/CreateOrEditFindingInstance"
	findingGetInstanceForEditPath     = "/App/Projects/CreateOrEditFindingInstanceModal"
	findingUploadEvidenceFilesPath    = "/App/Projects/UploadFindingEvidenceFiles"
)

// GetFindingInstanceForEditRaw performs POST /App/Projects/CreateOrEditFindingInstanceModal
// with form parameters: findingId, findingInstanceId, and projectId.
func (c *NonSupportedServiceClient) GetFindingInstanceForEditRaw(ctx context.Context, findingID, findingInstanceID, projectID string) ([]byte, int, error) {
	findingID = strings.TrimSpace(findingID)
	findingInstanceID = strings.TrimSpace(findingInstanceID)
	projectID = strings.TrimSpace(projectID)
	if findingID == "" {
		return nil, 0, fmt.Errorf("findingId is required")
	}
	if findingInstanceID == "" {
		return nil, 0, fmt.Errorf("findingInstanceId is required")
	}
	if projectID == "" {
		return nil, 0, fmt.Errorf("projectId is required")
	}
	form := url.Values{}
	form.Set("findingId", findingID)
	form.Set("findingInstanceId", findingInstanceID)
	form.Set("projectId", projectID)
	return c.DoForm(ctx, http.MethodPost, findingGetInstanceForEditPath, nil, form)
}

// GetFindingInstanceForEdit calls the modal endpoint and parses the HTML into the ABP-style response.
func (c *NonSupportedServiceClient) GetFindingInstanceForEdit(ctx context.Context, findingID, findingInstanceID, projectID string) (*GetFindingInstanceForEditResponse, []byte, int, error) {
	body, statusCode, err := c.GetFindingInstanceForEditRaw(ctx, findingID, findingInstanceID, projectID)
	if err != nil {
		return nil, body, statusCode, err
	}

	result, err := parseFindingInstanceModalHTML(body)
	if err != nil {
		return nil, body, statusCode, fmt.Errorf("parse CreateOrEditFindingInstanceModal response: %w", err)
	}
	resultJSON, err := json.Marshal(result)
	if err != nil {
		return nil, body, statusCode, fmt.Errorf("marshal parsed modal result: %w", err)
	}
	out := &GetFindingInstanceForEditResponse{
		Result:  resultJSON,
		Success: true,
		ABP:     true,
	}
	return out, body, statusCode, nil
}

func parseFindingInstanceModalHTML(rawHTML []byte) (map[string]interface{}, error) {
	root, err := html.Parse(strings.NewReader(string(rawHTML)))
	if err != nil {
		return nil, err
	}
	modalBody := findNode(root, func(n *html.Node) bool {
		return n.Type == html.ElementNode && n.Data == "div" && hasClass(n, "modal-body")
	})
	if modalBody == nil {
		return nil, fmt.Errorf("modal-body not found in html response")
	}

	out := map[string]interface{}{
		"Guid":              "",
		"Title":             "",
		"Asset":             "",
		"Location":          "",
		"Version":           "",
		"Ip":                "",
		"Hostname":          "",
		"Port":              "",
		"Protocol":          "",
		"IssueDetails":      "",
		"Reproduce":         "",
		"Evidence":          "",
		"IsVisibleInReport": "false",
		"EvidenceFiles":     []interface{}{},
		"NewEvidenceFiles":  []interface{}{},
	}

	if v, ok := getAttr(modalBody, "data-instance-id"); ok {
		out["Guid"] = strings.TrimSpace(v)
	}
	if v, ok := getAttr(modalBody, "data-existing-evidence-files"); ok && strings.TrimSpace(v) != "" {
		var evidenceFiles []interface{}
		if err := json.Unmarshal([]byte(v), &evidenceFiles); err == nil {
			out["EvidenceFiles"] = evidenceFiles
		}
	}

	fieldNames := []string{"Title", "Asset", "Location", "Version", "Ip", "Hostname", "Port", "Protocol", "IssueDetails", "Reproduce", "Evidence"}
	for _, name := range fieldNames {
		if v, ok := findFormFieldValue(modalBody, name); ok {
			out[name] = strings.TrimSpace(v)
		}
	}
	out["IsVisibleInReport"] = strings.ToLower(findCheckboxValue(modalBody, "IsVisibleInReport"))
	return out, nil
}

func findFormFieldValue(root *html.Node, fieldName string) (string, bool) {
	n := findNode(root, func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return false
		}
		if n.Data != "input" && n.Data != "textarea" && n.Data != "select" {
			return false
		}
		v, ok := getAttr(n, "name")
		return ok && v == fieldName
	})
	if n == nil {
		return "", false
	}
	switch n.Data {
	case "input":
		v, _ := getAttr(n, "value")
		return v, true
	case "textarea":
		return strings.TrimSpace(nodeText(n)), true
	case "select":
		selected := findNode(n, func(o *html.Node) bool {
			if o.Type != html.ElementNode || o.Data != "option" {
				return false
			}
			_, ok := getAttr(o, "selected")
			return ok
		})
		if selected != nil {
			v, _ := getAttr(selected, "value")
			return v, true
		}
		firstOption := findNode(n, func(o *html.Node) bool {
			return o.Type == html.ElementNode && o.Data == "option"
		})
		if firstOption != nil {
			v, _ := getAttr(firstOption, "value")
			return v, true
		}
	}
	return "", true
}

func findCheckboxValue(root *html.Node, fieldName string) string {
	n := findNode(root, func(n *html.Node) bool {
		if n.Type != html.ElementNode || n.Data != "input" {
			return false
		}
		t, _ := getAttr(n, "type")
		if !strings.EqualFold(t, "checkbox") {
			return false
		}
		v, ok := getAttr(n, "name")
		return ok && v == fieldName
	})
	if n == nil {
		return "false"
	}
	_, checked := getAttr(n, "checked")
	if checked {
		v, ok := getAttr(n, "value")
		if ok && strings.TrimSpace(v) != "" {
			return v
		}
		return "true"
	}
	return "false"
}

func findNode(root *html.Node, match func(*html.Node) bool) *html.Node {
	if root == nil {
		return nil
	}
	if match(root) {
		return root
	}
	for c := root.FirstChild; c != nil; c = c.NextSibling {
		if found := findNode(c, match); found != nil {
			return found
		}
	}
	return nil
}

func hasClass(n *html.Node, className string) bool {
	classes, ok := getAttr(n, "class")
	if !ok {
		return false
	}
	for _, c := range strings.Fields(classes) {
		if c == className {
			return true
		}
	}
	return false
}

func getAttr(n *html.Node, key string) (string, bool) {
	for _, a := range n.Attr {
		if a.Key == key {
			return a.Val, true
		}
	}
	return "", false
}

func nodeText(n *html.Node) string {
	var b strings.Builder
	var walk func(*html.Node)
	walk = func(cur *html.Node) {
		if cur.Type == html.TextNode {
			b.WriteString(cur.Data)
		}
		for c := cur.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(n)
	return b.String()
}

// CreateOrEditFindingInstanceRaw POSTs to /api/services/app/Finding/CreateOrEditFindingInstance
// with findingId as a query parameter. Use [CreateOrEditFindingInstanceRequest] for a typed body.
//
// Request JSON body (PascalCase top-level keys), for example:
//
//	Guid, Title, Asset, Location, Version, Ip, Hostname, Port, Protocol,
//	IssueDetails, Reproduce, Evidence, IsVisibleInReport,
//	EvidenceFiles (array), NewEvidenceFiles (array).
//
// EvidenceFiles is typically []. Each element of NewEvidenceFiles uses camelCase:
// fileName, fileType, fileToken (tokens come from the project file-upload flow).
// IsVisibleInReport is often sent as the string "true" or "false", not a JSON boolean.
func (c *NonSupportedServiceClient) CreateOrEditFindingInstanceRaw(ctx context.Context, findingID string, payload interface{}) ([]byte, int, error) {
	if strings.TrimSpace(findingID) == "" {
		return nil, 0, fmt.Errorf("findingId is required")
	}
	if payload == nil {
		return nil, 0, fmt.Errorf("payload is required")
	}
	q := url.Values{}
	q.Set("findingId", findingID)
	return c.DoJSON(ctx, http.MethodPost, findingCreateOrEditInstancePath, q, payload)
}

// CreateOrEditFindingInstance calls the endpoint and decodes the ABP response envelope.
func (c *NonSupportedServiceClient) CreateOrEditFindingInstance(ctx context.Context, findingID string, payload interface{}) (*CreateOrEditFindingInstanceResponse, []byte, int, error) {
	body, statusCode, err := c.CreateOrEditFindingInstanceRaw(ctx, findingID, payload)
	if err != nil {
		return nil, body, statusCode, err
	}

	var out CreateOrEditFindingInstanceResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, body, statusCode, fmt.Errorf("decode CreateOrEditFindingInstance response: %w", err)
	}
	if !out.Success {
		if out.Error != nil && strings.TrimSpace(out.Error.Message) != "" {
			return &out, body, statusCode, fmt.Errorf("service response unsuccessful: %s", out.Error.Message)
		}
		return &out, body, statusCode, fmt.Errorf("service response unsuccessful")
	}

	return &out, body, statusCode, nil
}

// UploadFindingEvidenceFilesRaw uploads one file to /App/Projects/UploadFindingEvidenceFiles.
func (c *NonSupportedServiceClient) UploadFindingEvidenceFilesRaw(ctx context.Context, filePath string) ([]byte, int, error) {
	bearer, err := c.requireBearerToken()
	if err != nil {
		return nil, 0, err
	}
	filePath = strings.TrimSpace(filePath)
	if filePath == "" {
		return nil, 0, fmt.Errorf("file path is required")
	}
	f, err := os.Open(filePath)
	if err != nil {
		return nil, 0, fmt.Errorf("open upload file: %w", err)
	}
	defer f.Close()

	var bodyBuf bytes.Buffer
	writer := multipart.NewWriter(&bodyBuf)
	fileName := filepath.Base(filePath)
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, fileName))
	h.Set("Content-Type", detectUploadFileContentType(fileName))
	part, err := writer.CreatePart(h)
	if err != nil {
		return nil, 0, fmt.Errorf("create multipart file part: %w", err)
	}
	if _, err := io.Copy(part, f); err != nil {
		return nil, 0, fmt.Errorf("write multipart file data: %w", err)
	}
	if err := writer.Close(); err != nil {
		return nil, 0, fmt.Errorf("finalize multipart body: %w", err)
	}

	reqURL := c.BaseURL + findingUploadEvidenceFilesPath
	getLogger().Info("Making non-supported service request", "method", http.MethodPost, "url", reqURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, &bodyBuf)
	if err != nil {
		return nil, 0, fmt.Errorf("create upload request: %w", err)
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("User-Agent", api.ChromeUserAgent)
	req.Header.Set("Authorization", "Bearer "+bearer)
	logServiceRequest(req)

	start := time.Now()
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("execute upload request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("read upload response body: %w", err)
	}
	logServiceResponse(resp, respBody, time.Since(start))
	if resp.StatusCode >= 400 {
		return respBody, resp.StatusCode, fmt.Errorf("service request failed with status %d", resp.StatusCode)
	}
	return respBody, resp.StatusCode, nil
}

// UploadFindingEvidenceFiles uploads one file and returns parsed file metadata (including fileToken).
func (c *NonSupportedServiceClient) UploadFindingEvidenceFiles(ctx context.Context, filePath string) ([]FindingTemplatesExportFile, []byte, int, error) {
	body, statusCode, err := c.UploadFindingEvidenceFilesRaw(ctx, filePath)
	if err != nil {
		return nil, body, statusCode, err
	}
	var out UploadFindingEvidenceFilesResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, body, statusCode, fmt.Errorf("decode UploadFindingEvidenceFiles response: %w", err)
	}
	if !out.Success {
		if out.Error != nil && strings.TrimSpace(out.Error.Message) != "" {
			return out.Result, body, statusCode, fmt.Errorf("service response unsuccessful: %s", out.Error.Message)
		}
		return out.Result, body, statusCode, fmt.Errorf("service response unsuccessful")
	}
	if len(out.Result) == 0 {
		return nil, body, statusCode, fmt.Errorf("service response missing result items")
	}
	return out.Result, body, statusCode, nil
}

func detectUploadFileContentType(fileName string) string {
	ext := strings.ToLower(filepath.Ext(strings.TrimSpace(fileName)))
	if ext == "" {
		return "application/octet-stream"
	}
	if ct := strings.TrimSpace(mime.TypeByExtension(ext)); ct != "" {
		if base, _, ok := strings.Cut(ct, ";"); ok {
			ct = strings.TrimSpace(base)
		}
		if ct != "" {
			return ct
		}
	}
	switch ext {
	case ".txt", ".log", ".md", ".csv":
		return "text/plain"
	case ".json":
		return "application/json"
	case ".yml", ".yaml":
		return "application/x-yaml"
	case ".png":
		return "image/png"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".gif":
		return "image/gif"
	case ".webp":
		return "image/webp"
	default:
		return "application/octet-stream"
	}
}
