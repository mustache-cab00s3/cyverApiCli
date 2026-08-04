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

// Get a paginated list of Projects based on a filter
// Verified against Full_api.json on April 17, 2026
func (c *ClientOps) ApiV22ClientProjectsGet(status string, maxResultCount, skipCount int, filter string) (*ProjectDtoV2PagedResultDtoAjaxResponse, error) {
	getLogger().Debug("Starting ApiV22ClientProjectsGet request", "status", status, "maxResultCount", maxResultCount, "skipCount", skipCount, "filter", filter)

	q := url.Values{}
	if status != "" {
		q.Set("Status", status)
	}
	q.Set("MaxResultCount", fmt.Sprint(maxResultCount))
	q.Set("SkipCount", fmt.Sprint(skipCount))
	if filter != "" {
		q.Set("Filter", filter)
	}
	path := fmt.Sprintf("/api/v2.2/client/projects?%s", q.Encode())

	getLogger().Info("Making API request", "method", http.MethodGet, "path", path, "queryParams", q.Encode())
	var response ProjectDtoV2PagedResultDtoAjaxResponse
	_, err := c.DoRequest(http.MethodGet, path, nil, &response)
	if err != nil {
		getLogger().Error("Failed ApiV22ClientProjectsGet", "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully completed ApiV22ClientProjectsGet")
	return &response, nil
}

// Get a Project by ID
// Verified against Full_api.json on April 17, 2026
func (c *ClientOps) ApiV22ClientProjectsByIdGet(id string) (*ProjectDtoV2AjaxResponse, error) {
	getLogger().Debug("Starting ApiV22ClientProjectsByIdGet request", "projectID", id)

	var response ProjectDtoV2AjaxResponse
	path := fmt.Sprintf("/api/v2.2/client/projects/%s", id)

	getLogger().Info("Making API request", "method", http.MethodGet, "path", path)
	_, err := c.DoRequest(http.MethodGet, path, nil, &response)
	if err != nil {
		getLogger().Error("Failed ApiV22ClientProjectsByIdGet", "projectID", id, "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully completed ApiV22ClientProjectsByIdGet", "projectID", id)
	return &response, nil
}

// Get a paginated list of Project request forms
// Verified against Full_api.json on April 17, 2026
func (c *ClientOps) ApiV22ClientProjectsRequestFormsGet(maxResultCount, skipCount int, filter string) (*RequestProjectFormDtoPagedResultDtoAjaxResponse, error) {
	getLogger().Debug("Starting ApiV22ClientProjectsRequestFormsGet request", "maxResultCount", maxResultCount, "skipCount", skipCount, "filter", filter)

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
		getLogger().Error("Failed ApiV22ClientProjectsRequestFormsGet", "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully completed ApiV22ClientProjectsRequestFormsGet")
	return &response, nil
}

// Submit a Project request
// Verified against Full_api.json on April 17, 2026
func (c *ClientOps) ApiV22ClientProjectsRequestPost(triggerEvents bool, body interface{}) (*GuidAjaxResponse, error) {
	getLogger().Debug("Starting ApiV22ClientProjectsRequestPost request", "triggerEvents", triggerEvents)

	q := url.Values{}
	q.Set("triggerEvents", fmt.Sprint(triggerEvents))
	path := fmt.Sprintf("/api/v2.2/client/projects/request?%s", q.Encode())

	getLogger().Info("Making API request", "method", http.MethodPost, "path", path)
	var response GuidAjaxResponse
	_, err := c.DoRequest(http.MethodPost, path, body, &response)
	if err != nil {
		getLogger().Error("Failed ApiV22ClientProjectsRequestPost", "triggerEvents", triggerEvents, "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully completed ApiV22ClientProjectsRequestPost", "triggerEvents", triggerEvents)
	return &response, nil
}

// ------------------ Continuous Projects ------------------

// Get a paginated list of Continuous Projects based on a filter
// Verified against Full_api.json on April 17, 2026
func (c *ClientOps) ApiV22ClientContinuousProjectsGet(status string, maxResultCount, skipCount int, filter string) (*ContinuousProjectDtoPagedResultDtoAjaxResponse, error) {
	getLogger().Debug("Starting ApiV22ClientContinuousProjectsGet request", "status", status, "maxResultCount", maxResultCount, "skipCount", skipCount, "filter", filter)

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
		getLogger().Error("Failed ApiV22ClientContinuousProjectsGet", "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully completed ApiV22ClientContinuousProjectsGet")
	return &response, nil
}

// Get a Continuous Project by ID
// Verified against Full_api.json on April 17, 2026
func (c *ClientOps) ApiV22ClientContinuousProjectsByIdGet(id string) (*ContinuousProjectDtoAjaxResponse, error) {
	getLogger().Debug("Starting ApiV22ClientContinuousProjectsByIdGet request", "projectID", id)

	var response ContinuousProjectDtoAjaxResponse
	path := fmt.Sprintf("/api/v2.2/client/continuous-projects/%s", id)

	getLogger().Info("Making API request", "method", http.MethodGet, "path", path)
	_, err := c.DoRequest(http.MethodGet, path, nil, &response)
	if err != nil {
		getLogger().Error("Failed ApiV22ClientContinuousProjectsByIdGet", "projectID", id, "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully completed ApiV22ClientContinuousProjectsByIdGet", "projectID", id)
	return &response, nil
}

// Get a paginated list of Continuous Project request forms
// Verified against Full_api.json on April 17, 2026
func (c *ClientOps) ApiV22ClientContinuousProjectsRequestFormsGet(maxResultCount, skipCount int, filter string) (*RequestProjectFormDtoPagedResultDtoAjaxResponse, error) {
	getLogger().Debug("Starting ApiV22ClientContinuousProjectsRequestFormsGet request", "maxResultCount", maxResultCount, "skipCount", skipCount, "filter", filter)

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
		getLogger().Error("Failed ApiV22ClientContinuousProjectsRequestFormsGet", "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully completed ApiV22ClientContinuousProjectsRequestFormsGet")
	return &response, nil
}

// Submit a Continuous Project request
// Verified against Full_api.json on April 17, 2026
func (c *ClientOps) ApiV22ClientContinuousProjectsRequestPost(triggerEvents bool, body interface{}) (*GuidAjaxResponse, error) {
	getLogger().Debug("Starting ApiV22ClientContinuousProjectsRequestPost request", "triggerEvents", triggerEvents)

	q := url.Values{}
	q.Set("triggerEvents", fmt.Sprint(triggerEvents))
	path := fmt.Sprintf("/api/v2.2/client/continuous-projects/request?%s", q.Encode())

	getLogger().Info("Making API request", "method", http.MethodPost, "path", path)
	var response GuidAjaxResponse
	_, err := c.DoRequest(http.MethodPost, path, body, &response)
	if err != nil {
		getLogger().Error("Failed ApiV22ClientContinuousProjectsRequestPost", "triggerEvents", triggerEvents, "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully completed ApiV22ClientContinuousProjectsRequestPost", "triggerEvents", triggerEvents)
	return &response, nil
}

// ------------------ Findings ------------------

// Get a paginated list of Findings based on a filter
// Verified against Full_api.json on April 17, 2026
func (c *ClientOps) ApiV22ClientFindingsGet(projectId string, maxResultCount, skipCount int) (*FindingDtoPagedResultDtoAjaxResponse, error) {
	getLogger().Debug("Starting ApiV22ClientFindingsGet request", "projectId", projectId, "maxResultCount", maxResultCount, "skipCount", skipCount)

	q := url.Values{}
	if projectId != "" {
		q.Set("ProjectId", projectId)
	}
	q.Set("MaxResultCount", fmt.Sprint(maxResultCount))
	q.Set("SkipCount", fmt.Sprint(skipCount))
	path := fmt.Sprintf("/api/v2.2/client/findings?%s", q.Encode())

	getLogger().Info("Making API request", "method", http.MethodGet, "path", path, "queryParams", q.Encode())
	var response FindingDtoPagedResultDtoAjaxResponse
	_, err := c.DoRequest(http.MethodGet, path, nil, &response)
	if err != nil {
		getLogger().Error("Failed ApiV22ClientFindingsGet", "projectId", projectId, "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully completed ApiV22ClientFindingsGet", "projectId", projectId)
	return &response, nil
}

// Get a Finding by ID
// Verified against Full_api.json on April 17, 2026
func (c *ClientOps) ApiV22ClientFindingsByIdGet(id string, includeEvidence bool) (*FindingDtoAjaxResponse, error) {
	getLogger().Debug("Starting ApiV22ClientFindingsByIdGet request", "findingID", id, "includeEvidence", includeEvidence)

	q := url.Values{}
	q.Set("includeEvidence", fmt.Sprint(includeEvidence))
	path := fmt.Sprintf("/api/v2.2/client/findings/%s?%s", id, q.Encode())

	getLogger().Info("Making API request", "method", http.MethodGet, "path", path)
	var response FindingDtoAjaxResponse
	_, err := c.DoRequest(http.MethodGet, path, nil, &response)
	if err != nil {
		getLogger().Error("Failed ApiV22ClientFindingsByIdGet", "findingID", id, "includeEvidence", includeEvidence, "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully completed ApiV22ClientFindingsByIdGet", "findingID", id, "includeEvidence", includeEvidence)
	return &response, nil
}

// Update the status of a Finding
// Verified against Full_api.json on April 17, 2026
func (c *ClientOps) ApiV22ClientFindingsByIdPost(id string, triggerEvents bool, statusBody interface{}) error {
	getLogger().Debug("Starting ApiV22ClientFindingsByIdPost request", "findingID", id, "triggerEvents", triggerEvents)

	q := url.Values{}
	q.Set("triggerEvents", fmt.Sprint(triggerEvents))
	path := fmt.Sprintf("/api/v2.2/client/findings/%s?%s", id, q.Encode())

	getLogger().Info("Making API request", "method", http.MethodPost, "path", path)
	var response interface{}
	_, err := c.DoRequest(http.MethodPost, path, statusBody, &response)
	if err != nil {
		getLogger().Error("Failed ApiV22ClientFindingsByIdPost", "findingID", id, "triggerEvents", triggerEvents, "error", err)
		return err
	}

	getLogger().Debug("Successfully completed ApiV22ClientFindingsByIdPost", "findingID", id, "triggerEvents", triggerEvents)
	return nil
}

// ------------------ Assets ------------------

// Get a paginated list of Assets based on a filter
// Verified against Full_api.json on April 17, 2026
func (c *ClientOps) ApiV22ClientAssetsGet(maxResultCount, skipCount int, filter string) (*AssetDtoPagedResultDtoAjaxResponse, error) {
	getLogger().Debug("Starting ApiV22ClientAssetsGet request", "maxResultCount", maxResultCount, "skipCount", skipCount, "filter", filter)

	q := url.Values{}
	q.Set("MaxResultCount", fmt.Sprint(maxResultCount))
	q.Set("SkipCount", fmt.Sprint(skipCount))
	if filter != "" {
		q.Set("Filter", filter)
	}
	path := fmt.Sprintf("/api/v2.2/client/assets?%s", q.Encode())

	getLogger().Info("Making API request", "method", http.MethodGet, "path", path, "queryParams", q.Encode())
	var response AssetDtoPagedResultDtoAjaxResponse
	_, err := c.DoRequest(http.MethodGet, path, nil, &response)
	if err != nil {
		getLogger().Error("Failed ApiV22ClientAssetsGet", "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully completed ApiV22ClientAssetsGet")
	return &response, nil
}

// Create a new Asset
// Verified against Full_api.json on April 17, 2026
func (c *ClientOps) ApiV22ClientAssetsPost(body interface{}) (*GuidAjaxResponse, error) {
	getLogger().Debug("Starting ApiV22ClientAssetsPost request")

	path := "/api/v2.2/client/assets"

	getLogger().Info("Making API request", "method", http.MethodPost, "path", path)
	var response GuidAjaxResponse
	_, err := c.DoRequest(http.MethodPost, path, body, &response)
	if err != nil {
		getLogger().Error("Failed ApiV22ClientAssetsPost", "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully completed ApiV22ClientAssetsPost")
	return &response, nil
}

// Delete an Asset by ID
// Verified against Full_api.json on April 17, 2026
func (c *ClientOps) ApiV22ClientAssetsByIdDelete(id string) (*GuidAjaxResponse, error) {
	getLogger().Debug("Starting ApiV22ClientAssetsByIdDelete request", "assetID", id)

	path := fmt.Sprintf("/api/v2.2/client/assets/%s", id)

	getLogger().Info("Making API request", "method", http.MethodDelete, "path", path)
	var response GuidAjaxResponse
	_, err := c.DoRequest(http.MethodDelete, path, nil, &response)
	if err != nil {
		getLogger().Error("Failed ApiV22ClientAssetsByIdDelete", "assetID", id, "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully completed ApiV22ClientAssetsByIdDelete", "assetID", id)
	return &response, nil
}

// Update an existing Asset
// Verified against Full_api.json on April 17, 2026
func (c *ClientOps) ApiV22ClientAssetsByIdPut(id string, body interface{}) error {
	getLogger().Debug("Starting ApiV22ClientAssetsByIdPut request", "assetID", id)

	path := fmt.Sprintf("/api/v2.2/client/assets/%s", id)

	getLogger().Info("Making API request", "method", http.MethodPut, "path", path)
	var response interface{}
	_, err := c.DoRequest(http.MethodPut, path, body, &response)
	if err != nil {
		getLogger().Error("Failed ApiV22ClientAssetsByIdPut", "assetID", id, "error", err)
		return err
	}

	getLogger().Debug("Successfully completed ApiV22ClientAssetsByIdPut", "assetID", id)
	return nil
}

// ------------------ Users ------------------

// Get a paginated list of Users based on a filter
// Verified against Full_api.json on April 17, 2026
func (c *ClientOps) ApiV22ClientUsersGet(maxResultCount, skipCount int, filter string) (*UserDtoPagedResultDtoAjaxResponse, error) {
	getLogger().Debug("Starting ApiV22ClientUsersGet request", "maxResultCount", maxResultCount, "skipCount", skipCount, "filter", filter)

	q := url.Values{}
	q.Set("MaxResultCount", fmt.Sprint(maxResultCount))
	q.Set("SkipCount", fmt.Sprint(skipCount))
	if filter != "" {
		q.Set("Filter", filter)
	}
	path := fmt.Sprintf("/api/v2.2/client/users?%s", q.Encode())

	getLogger().Info("Making API request", "method", http.MethodGet, "path", path, "queryParams", q.Encode())
	var response UserDtoPagedResultDtoAjaxResponse
	_, err := c.DoRequest(http.MethodGet, path, nil, &response)
	if err != nil {
		getLogger().Error("Failed ApiV22ClientUsersGet", "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully completed ApiV22ClientUsersGet")
	return &response, nil
}

// Create a new User
// Verified against Full_api.json on April 17, 2026
func (c *ClientOps) ApiV22ClientUsersPost(body interface{}) (*GuidAjaxResponse, error) {
	getLogger().Debug("Starting ApiV22ClientUsersPost request")

	path := "/api/v2.2/client/users"

	getLogger().Info("Making API request", "method", http.MethodPost, "path", path)
	var response GuidAjaxResponse
	_, err := c.DoRequest(http.MethodPost, path, body, &response)
	if err != nil {
		getLogger().Error("Failed ApiV22ClientUsersPost", "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully completed ApiV22ClientUsersPost")
	return &response, nil
}
