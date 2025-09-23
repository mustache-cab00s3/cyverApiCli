package v2_2

import (
	"fmt"
	"net/http"
	"net/url"
)

// ClientOps implements the ClientInterface for V2.2
type ClientOps struct {
	*Client
}

// ------------------ Projects ------------------

func (c *ClientOps) GetProjects(status string, maxResultCount, skipCount int, filter string) (*ProjectDtoV2PagedResultDtoAjaxResponse, error) {
	getLogger().Debug("Starting GetProjects request", "status", status, "maxResultCount", maxResultCount, "skipCount", skipCount, "filter", filter)

	q := url.Values{}
	path := fmt.Sprintf("/api/v2.2/client/projects?%s", q.Encode())
	if status != "" {
		q.Set("Status", status)
	}
	q.Set("MaxResultCount", fmt.Sprint(maxResultCount))
	q.Set("SkipCount", fmt.Sprint(skipCount))
	if filter != "" {
		q.Set("Filter", filter)
	}

	getLogger().Info("Making API request", "method", http.MethodGet, "path", path, "queryParams", q.Encode())
	var response ProjectDtoV2PagedResultDtoAjaxResponse
	_, err := c.DoRequest(http.MethodGet, path, nil, &response)
	if err != nil {
		getLogger().Error("Failed to get projects", "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully retrieved projects")
	return &response, nil
}

func (c *ClientOps) GetProjectByID(id string) (*ProjectDtoV2AjaxResponse, error) {
	getLogger().Debug("Starting GetProjectByID request", "projectID", id)

	path := fmt.Sprintf("/api/v2.2/client/projects/%s", id)

	getLogger().Info("Making API request", "method", http.MethodGet, "path", path)
	_, err := c.DoRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		getLogger().Error("Failed to get project by ID", "projectID", id, "error", err)
		return nil, err
	}

	var response ProjectDtoV2AjaxResponse
	getLogger().Debug("Successfully retrieved project", "projectID", id)
	return &response, nil
}

func (c *ClientOps) GetProjectRequestForms(maxResultCount, skipCount int, filter string) (*RequestProjectFormDtoPagedResultDtoAjaxResponse, error) {
	getLogger().Debug("Starting GetProjectRequestForms request", "maxResultCount", maxResultCount, "skipCount", skipCount, "filter", filter)

	q := url.Values{}
	q.Set("MaxResultCount", fmt.Sprint(maxResultCount))
	q.Set("SkipCount", fmt.Sprint(skipCount))
	if filter != "" {
		q.Set("Filter", filter)
	}
	path := fmt.Sprintf("/api/v2.2/client/projects/request-forms?%s", q.Encode())

	getLogger().Info("Making API request", "method", http.MethodGet, "path", path, "queryParams", q.Encode())
	var response RequestProjectFormDtoPagedResultDtoAjaxResponse
	_, err := c.DoRequest(http.MethodGet, path, q, &response)
	if err != nil {
		getLogger().Error("Failed to get project request forms", "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully retrieved project request forms")
	return &response, nil
}

func (c *ClientOps) RequestProject(triggerEvents bool, body interface{}) (*GuidAjaxResponse, error) {
	getLogger().Debug("Starting RequestProject request", "triggerEvents", triggerEvents)

	q := url.Values{}
	q.Set("triggerEvents", fmt.Sprint(triggerEvents))
	path := fmt.Sprintf("/api/v2.2/client/projects/request?%s", q.Encode())

	getLogger().Info("Making API request", "method", http.MethodPost, "path", path, "queryParams", q.Encode())
	var response GuidAjaxResponse
	_, err := c.DoRequest(http.MethodPost, path, q, &response)
	if err != nil {
		getLogger().Error("Failed to request project", "triggerEvents", triggerEvents, "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully requested project", "triggerEvents", triggerEvents)
	return &response, nil
}

// ------------------ Continuous Projects ------------------

func (c *ClientOps) GetContinuousProjects(status string, maxResultCount, skipCount int, filter string) (*ContinuousProjectDtoPagedResultDtoAjaxResponse, error) {
	getLogger().Debug("Starting GetContinuousProjects request", "status", status, "maxResultCount", maxResultCount, "skipCount", skipCount, "filter", filter)

	q := url.Values{}
	if status != "" {
		q.Set("Status", status)
	}
	q.Set("MaxResultCount", fmt.Sprint(maxResultCount))
	q.Set("SkipCount", fmt.Sprint(skipCount))
	if filter != "" {
		q.Set("Filter", filter)
	}
	path := fmt.Sprintf("/api/v2.2/client/continuous-projects?%s", q.Encode())

	getLogger().Info("Making API request", "method", http.MethodGet, "path", path, "queryParams", q.Encode())
	var response ContinuousProjectDtoPagedResultDtoAjaxResponse
	_, err := c.DoRequest(http.MethodGet, path, nil, &response)
	if err != nil {
		getLogger().Error("Failed to get continuous projects", "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully retrieved continuous projects")
	return &response, nil
}

func (c *ClientOps) GetContinuousProjectByID(id string) (*ContinuousProjectDtoAjaxResponse, error) {
	getLogger().Debug("Starting GetContinuousProjectByID request", "projectID", id)

	path := fmt.Sprintf("/api/v2.2/client/continuous-projects/%s", id)

	getLogger().Info("Making API request", "method", http.MethodGet, "path", path)
	var response ContinuousProjectDtoAjaxResponse
	_, err := c.DoRequest(http.MethodGet, path, nil, &response)
	if err != nil {
		getLogger().Error("Failed to get continuous project by ID", "projectID", id, "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully retrieved continuous project", "projectID", id)
	return &response, nil
}

func (c *ClientOps) GetContinuousProjectRequestForms(maxResultCount, skipCount int, filter string) (*RequestProjectFormDtoPagedResultDtoAjaxResponse, error) {
	getLogger().Debug("Starting GetContinuousProjectRequestForms request", "maxResultCount", maxResultCount, "skipCount", skipCount, "filter", filter)

	q := url.Values{}
	q.Set("MaxResultCount", fmt.Sprint(maxResultCount))
	q.Set("SkipCount", fmt.Sprint(skipCount))
	if filter != "" {
		q.Set("Filter", filter)
	}
	path := fmt.Sprintf("/api/v2.2/client/continuous-projects/request-forms?%s", q.Encode())

	getLogger().Info("Making API request", "method", http.MethodGet, "path", path, "queryParams", q.Encode())
	var response RequestProjectFormDtoPagedResultDtoAjaxResponse
	_, err := c.DoRequest(http.MethodGet, path, q, &response)
	if err != nil {
		getLogger().Error("Failed to get continuous project request forms", "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully retrieved continuous project request forms")
	return &response, nil
}

func (c *ClientOps) RequestContinuousProject(triggerEvents bool, body interface{}) (*GuidAjaxResponse, error) {
	getLogger().Debug("Starting RequestContinuousProject request", "triggerEvents", triggerEvents)

	q := url.Values{}
	q.Set("triggerEvents", fmt.Sprint(triggerEvents))
	path := fmt.Sprintf("/api/v2.2/client/continuous-projects/request?%s", q.Encode())

	getLogger().Info("Making API request", "method", http.MethodPost, "path", path, "queryParams", q.Encode())
	var response GuidAjaxResponse
	_, err := c.DoRequest(http.MethodPost, path, q, &response)
	if err != nil {
		getLogger().Error("Failed to request continuous project", "triggerEvents", triggerEvents, "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully requested continuous project", "triggerEvents", triggerEvents)
	return &response, nil
}

// ------------------ Findings ------------------

func (c *ClientOps) GetFindings(projectId string, maxResultCount, skipCount int) (*FindingDtoPagedResultDtoAjaxResponse, error) {
	getLogger().Debug("Starting GetFindings request", "projectId", projectId, "maxResultCount", maxResultCount, "skipCount", skipCount)

	q := url.Values{}
	if projectId != "" {
		q.Set("ProjectId", projectId)
	}
	q.Set("MaxResultCount", fmt.Sprint(maxResultCount))
	q.Set("SkipCount", fmt.Sprint(skipCount))
	path := fmt.Sprintf("/api/v2.2/client/findings?%s", q.Encode())

	getLogger().Info("Making API request", "method", http.MethodGet, "path", path, "queryParams", q.Encode())
	var response FindingDtoPagedResultDtoAjaxResponse
	_, err := c.DoRequest(http.MethodGet, path, q, &response)
	if err != nil {
		getLogger().Error("Failed to get findings", "projectId", projectId, "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully retrieved findings", "projectId", projectId)
	return &response, nil
}

func (c *ClientOps) GetFindingByID(id string, includeEvidence bool) (*FindingDtoAjaxResponse, error) {
	getLogger().Debug("Starting GetFindingByID request", "findingID", id, "includeEvidence", includeEvidence)

	q := url.Values{}
	q.Set("includeEvidence", fmt.Sprint(includeEvidence))
	path := fmt.Sprintf("/api/v2.2/client/findings/%s", id)

	getLogger().Info("Making API request", "method", http.MethodGet, "path", path, "queryParams", q.Encode())
	var response FindingDtoAjaxResponse
	_, err := c.DoRequest(http.MethodGet, path, q, &response)
	if err != nil {
		getLogger().Error("Failed to get finding by ID", "findingID", id, "includeEvidence", includeEvidence, "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully retrieved finding", "findingID", id, "includeEvidence", includeEvidence)
	return &response, nil
}

func (c *ClientOps) SetFindingStatus(id string, triggerEvents int, statusBody interface{}) error {
	getLogger().Debug("Starting SetFindingStatus request", "findingID", id, "triggerEvents", triggerEvents)

	q := url.Values{}
	q.Set("triggerEvents", fmt.Sprint(triggerEvents))
	path := fmt.Sprintf("/api/v2.2/client/findings/%s", id)

	getLogger().Info("Making API request", "method", http.MethodPost, "path", path, "queryParams", q.Encode())
	var response GuidAjaxResponse
	_, err := c.DoRequest(http.MethodPost, path, q, &response)
	if err != nil {
		getLogger().Error("Failed to set finding status", "findingID", id, "triggerEvents", triggerEvents, "error", err)
		return err
	}

	getLogger().Debug("Successfully set finding status", "findingID", id, "triggerEvents", triggerEvents)
	return nil
}

// ------------------ Assets ------------------

func (c *ClientOps) GetAssets(maxResultCount, skipCount int, filter string) (*AssetDtoPagedResultDtoAjaxResponse, error) {
	getLogger().Debug("Starting GetAssets request", "maxResultCount", maxResultCount, "skipCount", skipCount, "filter", filter)

	q := url.Values{}
	q.Set("MaxResultCount", fmt.Sprint(maxResultCount))
	q.Set("SkipCount", fmt.Sprint(skipCount))
	if filter != "" {
		q.Set("Filter", filter)
	}
	path := fmt.Sprintf("/api/v2.2/client/assets?%s", q.Encode())

	getLogger().Info("Making API request", "method", http.MethodGet, "path", path, "queryParams", q.Encode())
	var response AssetDtoPagedResultDtoAjaxResponse
	_, err := c.DoRequest(http.MethodGet, path, q, &response)
	if err != nil {
		getLogger().Error("Failed to get assets", "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully retrieved assets")
	return &response, nil
}

func (c *ClientOps) CreateAsset(body interface{}) (*GuidAjaxResponse, error) {
	getLogger().Debug("Starting CreateAsset request")

	path := "/api/v2.2/client/assets"

	getLogger().Info("Making API request", "method", http.MethodPost, "path", path)
	var response GuidAjaxResponse
	_, err := c.DoRequest(http.MethodPost, path, nil, &response)
	if err != nil {
		getLogger().Error("Failed to create asset", "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully created asset")
	return &response, nil
}

func (c *ClientOps) DeleteAsset(id string) (*GuidAjaxResponse, error) {
	getLogger().Debug("Starting DeleteAsset request", "assetID", id)

	path := fmt.Sprintf("/api/v2.2/client/assets/%s", id)

	getLogger().Info("Making API request", "method", http.MethodDelete, "path", path)
	var response GuidAjaxResponse
	_, err := c.DoRequest(http.MethodDelete, path, nil, &response)
	if err != nil {
		getLogger().Error("Failed to delete asset", "assetID", id, "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully deleted asset", "assetID", id)
	return &response, nil
}

func (c *ClientOps) UpdateAsset(id string, body interface{}) error {
	getLogger().Debug("Starting UpdateAsset request", "assetID", id)

	path := fmt.Sprintf("/api/v2.2/client/assets/%s", id)

	getLogger().Info("Making API request", "method", http.MethodPut, "path", path)
	var response GuidAjaxResponse
	_, err := c.DoRequest(http.MethodPut, path, nil, &response)
	if err != nil {
		getLogger().Error("Failed to update asset", "assetID", id, "error", err)
		return err
	}

	getLogger().Debug("Successfully updated asset", "assetID", id)
	return nil
}

// ------------------ Users ------------------

func (c *ClientOps) GetUsers(maxResultCount, skipCount int, filter string) (*UserDtoPagedResultDtoAjaxResponse, error) {
	getLogger().Debug("Starting GetUsers request", "maxResultCount", maxResultCount, "skipCount", skipCount, "filter", filter)

	q := url.Values{}
	q.Set("MaxResultCount", fmt.Sprint(maxResultCount))
	q.Set("SkipCount", fmt.Sprint(skipCount))
	if filter != "" {
		q.Set("Filter", filter)
	}
	path := fmt.Sprintf("/api/v2.2/client/users?%s", q.Encode())

	getLogger().Info("Making API request", "method", http.MethodGet, "path", path, "queryParams", q.Encode())
	var response UserDtoPagedResultDtoAjaxResponse
	_, err := c.DoRequest(http.MethodGet, path, q, &response)
	if err != nil {
		getLogger().Error("Failed to get users", "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully retrieved users")
	return &response, nil
}

func (c *ClientOps) CreateUser(body interface{}) (*GuidAjaxResponse, error) {
	getLogger().Debug("Starting CreateUser request")

	path := "/api/v2.2/client/users"

	getLogger().Info("Making API request", "method", http.MethodPost, "path", path)
	var response GuidAjaxResponse
	_, err := c.DoRequest(http.MethodPost, path, nil, &response)
	if err != nil {
		getLogger().Error("Failed to create user", "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully created user")
	return &response, nil
}
