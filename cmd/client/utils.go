package client

import (
	"fmt"

	"github.com/yourusername/cyverApiCli/cmd/shared"
	"github.com/yourusername/cyverApiCli/internal/api/versions/v2_2"
)

// =============================================================================
// STRING UTILITIES
// =============================================================================

// getStringValue safely gets string value from pointer
func getStringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// =============================================================================
// OUTPUT FORMATTING UTILITIES
// =============================================================================

// printSimpleUsersList prints a simple users list
func printSimpleUsersList(users *v2_2.UserDtoPagedResultDtoAjaxResponse) error {
	if users == nil || users.Result == nil || len(users.Result.Items) == 0 {
		fmt.Println("[]")
		return nil
	}

	simpleList := make([]map[string]string, len(users.Result.Items))
	for i, user := range users.Result.Items {
		simpleList[i] = map[string]string{
			"id":    user.ID,
			"name":  getStringValue(user.Name),
			"email": getStringValue(user.Email),
		}
	}
	return shared.PrintJSONResponse(simpleList)
}

// printSimpleUsersTable prints a simple users table
func printSimpleUsersTable(users *v2_2.UserDtoPagedResultDtoAjaxResponse) error {
	if users == nil || users.Result == nil || len(users.Result.Items) == 0 {
		fmt.Println("No users found")
		return nil
	}

	fmt.Println("┌─────────────────────────────────────────────────────────────┐")
	fmt.Println("│                           USERS                            │")
	fmt.Println("├─────────────────────────────────────────────────────────────┤")
	for _, user := range users.Result.Items {
		fmt.Printf("│ %-20s │ %-35s │\n", user.ID, getStringValue(user.Name))
	}
	fmt.Println("└─────────────────────────────────────────────────────────────┘")
	return nil
}

// printSimpleAssetsList prints a simple assets list
func printSimpleAssetsList(assets *v2_2.AssetDtoPagedResultDtoAjaxResponse) error {
	if assets == nil || assets.Result == nil || len(assets.Result.Items) == 0 {
		fmt.Println("[]")
		return nil
	}

	simpleList := make([]map[string]string, len(assets.Result.Items))
	for i, asset := range assets.Result.Items {
		simpleList[i] = map[string]string{
			"id":   asset.ID,
			"name": getStringValue(asset.Title),
		}
	}
	return shared.PrintJSONResponse(simpleList)
}

// printSimpleAssetsTable prints a simple assets table
func printSimpleAssetsTable(assets *v2_2.AssetDtoPagedResultDtoAjaxResponse) error {
	if assets == nil || assets.Result == nil || len(assets.Result.Items) == 0 {
		fmt.Println("No assets found")
		return nil
	}

	fmt.Println("┌─────────────────────────────────────────────────────────────┐")
	fmt.Println("│                           ASSETS                           │")
	fmt.Println("├─────────────────────────────────────────────────────────────┤")
	for _, asset := range assets.Result.Items {
		fmt.Printf("│ %-20s │ %-35s │\n", asset.ID, getStringValue(asset.Title))
	}
	fmt.Println("└─────────────────────────────────────────────────────────────┘")
	return nil
}

// printSimpleRequestFormsList prints a simple request forms list
func printSimpleRequestFormsList(requestForms *v2_2.RequestProjectFormDtoPagedResultDtoAjaxResponse) error {
	if requestForms == nil || len(requestForms.Data) == 0 {
		fmt.Println("[]")
		return nil
	}

	simpleList := make([]map[string]string, len(requestForms.Data))
	for i, form := range requestForms.Data {
		simpleList[i] = map[string]string{
			"id":   form.ID,
			"name": form.Name,
		}
	}
	return shared.PrintJSONResponse(simpleList)
}

// printSimpleRequestFormsTable prints a simple request forms table
func printSimpleRequestFormsTable(requestForms *v2_2.RequestProjectFormDtoPagedResultDtoAjaxResponse) error {
	if requestForms == nil || len(requestForms.Data) == 0 {
		fmt.Println("No request forms found")
		return nil
	}

	fmt.Println("┌─────────────────────────────────────────────────────────────┐")
	fmt.Println("│                      REQUEST FORMS                         │")
	fmt.Println("├─────────────────────────────────────────────────────────────┤")
	for _, form := range requestForms.Data {
		fmt.Printf("│ %-20s │ %-35s │\n", form.ID, form.Name)
	}
	fmt.Println("└─────────────────────────────────────────────────────────────┘")
	return nil
}

// =============================================================================
// CONTINUOUS PROJECTS UTILITIES
// =============================================================================

// printSimpleContinuousProjectsList prints a simple continuous projects list
func printSimpleContinuousProjectsList(continuousProjects *v2_2.ContinuousProjectDtoPagedResultDtoAjaxResponse) error {
	if continuousProjects == nil || len(continuousProjects.Data) == 0 {
		fmt.Println("[]")
		return nil
	}

	simpleList := make([]map[string]string, len(continuousProjects.Data))
	for i, project := range continuousProjects.Data {
		simpleList[i] = map[string]string{
			"id":   project.ID,
			"name": project.Name,
		}
	}
	return shared.PrintJSONResponse(simpleList)
}

// printSimpleContinuousProjectsTable prints a simple continuous projects table
func printSimpleContinuousProjectsTable(continuousProjects *v2_2.ContinuousProjectDtoPagedResultDtoAjaxResponse) error {
	if continuousProjects == nil || len(continuousProjects.Data) == 0 {
		fmt.Println("No continuous projects found")
		return nil
	}

	fmt.Println("┌─────────────────────────────────────────────────────────────┐")
	fmt.Println("│                    CONTINUOUS PROJECTS                     │")
	fmt.Println("├─────────────────────────────────────────────────────────────┤")
	for _, project := range continuousProjects.Data {
		fmt.Printf("│ %-20s │ %-35s │\n", project.ID, project.Name)
	}
	fmt.Println("└─────────────────────────────────────────────────────────────┘")
	return nil
}

// printContinuousProjectTable prints continuous project information in table format
func printContinuousProjectTable(continuousProject *v2_2.ContinuousProjectDtoAjaxResponse) error {
	if continuousProject == nil || continuousProject.Data.ID == "" {
		fmt.Println("No continuous project information available")
		return nil
	}

	fmt.Println("┌─────────────────────────────────────────────────────────────┐")
	fmt.Println("│                  CONTINUOUS PROJECT INFO                   │")
	fmt.Println("├─────────────────────────────────────────────────────────────┤")
	fmt.Printf("│ ID:          %-45s │\n", continuousProject.Data.ID)
	fmt.Printf("│ Name:        %-45s │\n", continuousProject.Data.Name)
	fmt.Printf("│ Description: %-45s │\n", continuousProject.Data.Description)
	fmt.Printf("│ Status:      %-45s │\n", continuousProject.Data.Status)
	fmt.Println("└─────────────────────────────────────────────────────────────┘")
	return nil
}

// printSimpleContinuousRequestFormsList prints a simple continuous request forms list
func printSimpleContinuousRequestFormsList(requestForms *v2_2.RequestProjectFormDtoPagedResultDtoAjaxResponse) error {
	if requestForms == nil || len(requestForms.Data) == 0 {
		fmt.Println("[]")
		return nil
	}

	simpleList := make([]map[string]string, len(requestForms.Data))
	for i, form := range requestForms.Data {
		simpleList[i] = map[string]string{
			"id":   form.ID,
			"name": form.Name,
		}
	}
	return shared.PrintJSONResponse(simpleList)
}

// printSimpleContinuousRequestFormsTable prints a simple continuous request forms table
func printSimpleContinuousRequestFormsTable(requestForms *v2_2.RequestProjectFormDtoPagedResultDtoAjaxResponse) error {
	if requestForms == nil || len(requestForms.Data) == 0 {
		fmt.Println("No continuous project request forms found")
		return nil
	}

	fmt.Println("┌─────────────────────────────────────────────────────────────┐")
	fmt.Println("│                CONTINUOUS REQUEST FORMS                    │")
	fmt.Println("├─────────────────────────────────────────────────────────────┤")
	for _, form := range requestForms.Data {
		fmt.Printf("│ %-20s │ %-35s │\n", form.ID, form.Name)
	}
	fmt.Println("└─────────────────────────────────────────────────────────────┘")
	return nil
}
