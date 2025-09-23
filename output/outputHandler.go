package output

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	log "github.com/yourusername/cyverApiCli/logger"
)

func PrintJSONResponse(data interface{}, verboseLevel int) error {
	// Check if the data has a RawJSON field (for full API responses)
	if response, ok := data.(interface{ GetRawJSON() interface{} }); ok {
		if rawJSON := response.GetRawJSON(); rawJSON != nil {
			// Use the raw JSON response directly
			if rawBytes, ok := rawJSON.([]byte); ok {
				fmt.Println(string(rawBytes))
				return nil
			}
			// If it's not []byte, try to marshal it as JSON
			if rawJSONStr, ok := rawJSON.(string); ok {
				fmt.Println(rawJSONStr)
				return nil
			}
		}
	}

	// Fall back to marshaling the data structure
	prettyJSON, err := json.MarshalIndent(data, "", "  ") // Use two spaces for indentation
	if err != nil {
		log.GetLogger(verboseLevel).Error("failed to format JSON output: %v\n", err)
		return fmt.Errorf("failed to format JSON output: %w", err)
	}
	log.GetLogger(verboseLevel).Info("JSON response formatted successfully\n")
	fmt.Println(string(prettyJSON)) // Always print output
	return nil
}

// SimpleProject represents a project with only ID and Name fields
type SimpleProject struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// printSimpleProjectsList extracts only id and name from paged project results and prints them
func PrintSimpleProjectsList(data interface{}, verboseLevel int) error {
	var items []interface{}

	// Try to handle both struct format and raw JSON format
	if dataMap, ok := data.(map[string]interface{}); ok {
		// Handle raw JSON format (when unmarshaling fails)
		if result, ok := dataMap["result"].(map[string]interface{}); ok {
			if itemsList, ok := result["items"].([]interface{}); ok {
				items = itemsList
			}
		}
	} else {
		// Handle struct format (when unmarshaling succeeds)
		val := reflect.ValueOf(data)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}

		// Try to access the Result field which should contain the paged data
		resultField := val.FieldByName("Result")
		if !resultField.IsValid() {
			return fmt.Errorf("no Result field found in response data")
		}

		if resultField.Kind() == reflect.Ptr {
			resultField = resultField.Elem()
		}

		// Try to access the Items field which should contain the project list
		itemsField := resultField.FieldByName("Items")
		if !itemsField.IsValid() {
			return fmt.Errorf("no Items field found in Result")
		}

		// Convert items to a slice of interfaces
		if itemsField.Kind() == reflect.Slice {
			for i := 0; i < itemsField.Len(); i++ {
				items = append(items, itemsField.Index(i).Interface())
			}
		}
	}

	// Extract only id and name from each item
	var simpleProjects []SimpleProject
	for _, item := range items {
		// Handle both map[string]interface{} and struct cases
		var id, name string

		switch itemVal := item.(type) {
		case map[string]interface{}:
			// Handle map case (most likely scenario)
			if idVal, ok := itemVal["id"]; ok {
				id = fmt.Sprintf("%v", idVal)
			} else if idVal, ok := itemVal["ID"]; ok {
				id = fmt.Sprintf("%v", idVal)
			} else if idVal, ok := itemVal["Id"]; ok {
				id = fmt.Sprintf("%v", idVal)
			}

			// Handle pointer to string fields properly
			if nameVal, ok := itemVal["name"]; ok && nameVal != nil {
				name = fmt.Sprintf("%v", nameVal)
			} else if nameVal, ok := itemVal["Name"]; ok && nameVal != nil {
				name = fmt.Sprintf("%v", nameVal)
			}
		default:
			// Handle struct case using reflection
			itemReflect := reflect.ValueOf(item)
			if itemReflect.Kind() == reflect.Ptr {
				itemReflect = itemReflect.Elem()
			}

			if itemReflect.Kind() == reflect.Struct {
				// Try to get ID field (case insensitive)
				idField := itemReflect.FieldByName("ID")
				if !idField.IsValid() {
					idField = itemReflect.FieldByName("Id")
				}
				if idField.IsValid() {
					id = fmt.Sprintf("%v", idField.Interface())
				}

				// Try to get Name field - handle pointer to string case
				nameField := itemReflect.FieldByName("Name")
				if nameField.IsValid() {
					// Handle pointer to string case
					if nameField.Kind() == reflect.Ptr && !nameField.IsNil() {
						name = fmt.Sprintf("%v", nameField.Elem().Interface())
					} else if nameField.Kind() != reflect.Ptr {
						name = fmt.Sprintf("%v", nameField.Interface())
					}
				}
			}
		}

		// Only add if we have both id and name
		if id != "" && name != "" {
			simpleProject := SimpleProject{
				ID:   id,
				Name: name,
			}
			simpleProjects = append(simpleProjects, simpleProject)
		}
	}

	// Print the simplified list
	prettyJSON, err := json.MarshalIndent(simpleProjects, "", "  ")
	if err != nil {
		log.GetLogger(verboseLevel).Error("failed to format simple projects JSON output: %v\n", err)
		return fmt.Errorf("failed to format simple projects JSON output: %w", err)
	}
	log.GetLogger(verboseLevel).Info("Simple projects list formatted successfully\n")
	fmt.Println(string(prettyJSON))
	return nil
}

// printSimpleProjectsTable extracts only id and name from paged project results and prints them as a table
func PrintSimpleProjectsTable(data interface{}, verboseLevel int) error {
	var items []interface{}

	// Try to handle both struct format and raw JSON format
	if dataMap, ok := data.(map[string]interface{}); ok {
		// Handle raw JSON format (when unmarshaling fails)
		if result, ok := dataMap["result"].(map[string]interface{}); ok {
			if itemsList, ok := result["items"].([]interface{}); ok {
				items = itemsList
			}
		}
	} else {
		// Handle struct format (when unmarshaling succeeds)
		val := reflect.ValueOf(data)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}

		// Try to access the Result field which should contain the paged data
		resultField := val.FieldByName("Result")
		if !resultField.IsValid() {
			return fmt.Errorf("no Result field found in response data")
		}

		if resultField.Kind() == reflect.Ptr {
			resultField = resultField.Elem()
		}

		// Try to access the Items field which should contain the project list
		itemsField := resultField.FieldByName("Items")
		if !itemsField.IsValid() {
			return fmt.Errorf("no Items field found in Result")
		}

		// Convert items to a slice of interfaces
		if itemsField.Kind() == reflect.Slice {
			for i := 0; i < itemsField.Len(); i++ {
				items = append(items, itemsField.Index(i).Interface())
			}
		}
	}

	// Extract only id and name from each item
	var simpleProjects []SimpleProject
	for _, item := range items {
		// Handle both map[string]interface{} and struct cases
		var id, name string

		switch itemVal := item.(type) {
		case map[string]interface{}:
			// Handle map case (most likely scenario)
			if idVal, ok := itemVal["id"]; ok {
				id = fmt.Sprintf("%v", idVal)
			} else if idVal, ok := itemVal["ID"]; ok {
				id = fmt.Sprintf("%v", idVal)
			} else if idVal, ok := itemVal["Id"]; ok {
				id = fmt.Sprintf("%v", idVal)
			}

			// Handle pointer to string fields properly
			if nameVal, ok := itemVal["name"]; ok && nameVal != nil {
				name = fmt.Sprintf("%v", nameVal)
			} else if nameVal, ok := itemVal["Name"]; ok && nameVal != nil {
				name = fmt.Sprintf("%v", nameVal)
			}
		default:
			// Handle struct case using reflection
			itemReflect := reflect.ValueOf(item)
			if itemReflect.Kind() == reflect.Ptr {
				itemReflect = itemReflect.Elem()
			}

			if itemReflect.Kind() == reflect.Struct {
				// Try to get ID field (case insensitive)
				idField := itemReflect.FieldByName("ID")
				if !idField.IsValid() {
					idField = itemReflect.FieldByName("Id")
				}
				if idField.IsValid() {
					id = fmt.Sprintf("%v", idField.Interface())
				}

				// Try to get Name field - handle pointer to string case
				nameField := itemReflect.FieldByName("Name")
				if nameField.IsValid() {
					// Handle pointer to string case
					if nameField.Kind() == reflect.Ptr && !nameField.IsNil() {
						name = fmt.Sprintf("%v", nameField.Elem().Interface())
					} else if nameField.Kind() != reflect.Ptr {
						name = fmt.Sprintf("%v", nameField.Interface())
					}
				}
			}
		}

		// Only add if we have both id and name
		if id != "" && name != "" {
			simpleProject := SimpleProject{
				ID:   id,
				Name: name,
			}
			simpleProjects = append(simpleProjects, simpleProject)
		}
	}

	// Print the table
	if len(simpleProjects) == 0 {
		fmt.Println("No projects found.")
		return nil
	}

	// Calculate column widths
	maxIDWidth := len("ID")
	maxNameWidth := len("NAME")

	for _, project := range simpleProjects {
		if len(project.ID) > maxIDWidth {
			maxIDWidth = len(project.ID)
		}
		if len(project.Name) > maxNameWidth {
			maxNameWidth = len(project.Name)
		}
	}

	// Ensure minimum widths
	if maxIDWidth < 10 {
		maxIDWidth = 10
	}
	if maxNameWidth < 20 {
		maxNameWidth = 20
	}

	// Print header
	header := fmt.Sprintf("| %-*s | %-*s |", maxIDWidth, "ID", maxNameWidth, "NAME")
	separator := fmt.Sprintf("|%s|%s|",
		strings.Repeat("-", maxIDWidth+2),
		strings.Repeat("-", maxNameWidth+2))

	fmt.Println(separator)
	fmt.Println(header)
	fmt.Println(separator)

	// Print rows
	for _, project := range simpleProjects {
		// Truncate name if too long
		displayName := project.Name
		if len(displayName) > maxNameWidth {
			displayName = displayName[:maxNameWidth-3] + "..."
		}

		row := fmt.Sprintf("| %-*s | %-*s |", maxIDWidth, project.ID, maxNameWidth, displayName)
		fmt.Println(row)
	}

	fmt.Println(separator)
	log.GetLogger(verboseLevel).Info("Simple projects table formatted successfully\n")
	return nil
}

// SimpleFinding represents a finding with only ID and Title fields
type SimpleFinding struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// printSimpleFindingsList extracts only id and title from paged finding results and prints them
func PrintSimpleFindingsList(data interface{}, verboseLevel int) error {
	// Use reflection to safely access the data structure
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Try to access the Result field which should contain the paged data
	resultField := val.FieldByName("Result")
	if !resultField.IsValid() {
		return fmt.Errorf("no Result field found in response data")
	}

	if resultField.Kind() == reflect.Ptr {
		resultField = resultField.Elem()
	}

	// Try to access the Items field which should contain the finding list
	itemsField := resultField.FieldByName("Items")
	if !itemsField.IsValid() {
		return fmt.Errorf("no Items field found in Result")
	}

	// Convert to slice
	items := itemsField.Interface()
	itemsSlice := reflect.ValueOf(items)
	if itemsSlice.Kind() != reflect.Slice {
		return fmt.Errorf("items field is not a slice")
	}

	// Convert items to a slice of interfaces
	var itemsList []interface{}
	if itemsSlice.Kind() == reflect.Slice {
		for i := 0; i < itemsSlice.Len(); i++ {
			itemsList = append(itemsList, itemsSlice.Index(i).Interface())
		}
	}

	// Extract only id and title from each item
	var simpleFindings []SimpleFinding
	for _, item := range itemsList {
		// Handle both map[string]interface{} and struct cases
		var id, title string

		switch itemVal := item.(type) {
		case map[string]interface{}:
			// Handle map case (most likely scenario)
			if idVal, ok := itemVal["id"]; ok {
				id = fmt.Sprintf("%v", idVal)
			} else if idVal, ok := itemVal["ID"]; ok {
				id = fmt.Sprintf("%v", idVal)
			} else if idVal, ok := itemVal["Id"]; ok {
				id = fmt.Sprintf("%v", idVal)
			}

			// For findings, the field is "name" not "title"
			if nameVal, ok := itemVal["name"]; ok && nameVal != nil {
				title = fmt.Sprintf("%v", nameVal)
			} else if titleVal, ok := itemVal["title"]; ok && titleVal != nil {
				title = fmt.Sprintf("%v", titleVal)
			} else if titleVal, ok := itemVal["Title"]; ok && titleVal != nil {
				title = fmt.Sprintf("%v", titleVal)
			} else if nameVal, ok := itemVal["Name"]; ok && nameVal != nil {
				title = fmt.Sprintf("%v", nameVal)
			}
		default:
			// Handle struct case using reflection
			itemReflect := reflect.ValueOf(item)
			if itemReflect.Kind() == reflect.Ptr {
				itemReflect = itemReflect.Elem()
			}

			if itemReflect.Kind() == reflect.Struct {
				// Try to get ID field (case insensitive)
				idField := itemReflect.FieldByName("ID")
				if !idField.IsValid() {
					idField = itemReflect.FieldByName("Id")
				}
				if idField.IsValid() {
					id = fmt.Sprintf("%v", idField.Interface())
				}

				// Try to get Name/Title field - for findings, it's "Name"
				titleField := itemReflect.FieldByName("Name")
				if !titleField.IsValid() {
					titleField = itemReflect.FieldByName("Title")
				}
				if titleField.IsValid() {
					// Handle pointer to string case
					if titleField.Kind() == reflect.Ptr && !titleField.IsNil() {
						title = fmt.Sprintf("%v", titleField.Elem().Interface())
					} else if titleField.Kind() != reflect.Ptr {
						title = fmt.Sprintf("%v", titleField.Interface())
					}
				}
			}
		}

		// Only add if we have both id and title
		if id != "" && title != "" {
			simpleFinding := SimpleFinding{
				ID:    id,
				Title: title,
			}
			simpleFindings = append(simpleFindings, simpleFinding)
		}
	}

	// Print the JSON
	prettyJSON, err := json.MarshalIndent(simpleFindings, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to format JSON output: %w", err)
	}

	fmt.Println(string(prettyJSON))
	log.GetLogger(verboseLevel).Info("Simple findings list formatted successfully\n")
	return nil
}

// printSimpleFindingsTable extracts only id and title from paged finding results and prints them in table format
func PrintSimpleFindingsTable(data interface{}, verboseLevel int) error {
	// Use reflection to safely access the data structure
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Try to access the Result field which should contain the paged data
	resultField := val.FieldByName("Result")
	if !resultField.IsValid() {
		return fmt.Errorf("no Result field found in response data")
	}

	if resultField.Kind() == reflect.Ptr {
		resultField = resultField.Elem()
	}

	// Try to access the Items field which should contain the finding list
	itemsField := resultField.FieldByName("Items")
	if !itemsField.IsValid() {
		return fmt.Errorf("no Items field found in Result")
	}

	// Convert to slice
	items := itemsField.Interface()
	itemsSlice := reflect.ValueOf(items)
	if itemsSlice.Kind() != reflect.Slice {
		return fmt.Errorf("items field is not a slice")
	}

	// Convert items to a slice of interfaces
	var itemsList []interface{}
	if itemsSlice.Kind() == reflect.Slice {
		for i := 0; i < itemsSlice.Len(); i++ {
			itemsList = append(itemsList, itemsSlice.Index(i).Interface())
		}
	}

	// Extract only id and title from each item
	var simpleFindings []SimpleFinding
	for _, item := range itemsList {
		// Handle both map[string]interface{} and struct cases
		var id, title string

		switch itemVal := item.(type) {
		case map[string]interface{}:
			// Handle map case (most likely scenario)
			if idVal, ok := itemVal["id"]; ok {
				id = fmt.Sprintf("%v", idVal)
			} else if idVal, ok := itemVal["ID"]; ok {
				id = fmt.Sprintf("%v", idVal)
			} else if idVal, ok := itemVal["Id"]; ok {
				id = fmt.Sprintf("%v", idVal)
			}

			// For findings, the field is "name" not "title"
			if nameVal, ok := itemVal["name"]; ok && nameVal != nil {
				title = fmt.Sprintf("%v", nameVal)
			} else if titleVal, ok := itemVal["title"]; ok && titleVal != nil {
				title = fmt.Sprintf("%v", titleVal)
			} else if titleVal, ok := itemVal["Title"]; ok && titleVal != nil {
				title = fmt.Sprintf("%v", titleVal)
			} else if nameVal, ok := itemVal["Name"]; ok && nameVal != nil {
				title = fmt.Sprintf("%v", nameVal)
			}
		default:
			// Handle struct case using reflection
			itemReflect := reflect.ValueOf(item)
			if itemReflect.Kind() == reflect.Ptr {
				itemReflect = itemReflect.Elem()
			}

			if itemReflect.Kind() == reflect.Struct {
				// Try to get ID field (case insensitive)
				idField := itemReflect.FieldByName("ID")
				if !idField.IsValid() {
					idField = itemReflect.FieldByName("Id")
				}
				if idField.IsValid() {
					id = fmt.Sprintf("%v", idField.Interface())
				}

				// Try to get Name/Title field - for findings, it's "Name"
				titleField := itemReflect.FieldByName("Name")
				if !titleField.IsValid() {
					titleField = itemReflect.FieldByName("Title")
				}
				if titleField.IsValid() {
					// Handle pointer to string case
					if titleField.Kind() == reflect.Ptr && !titleField.IsNil() {
						title = fmt.Sprintf("%v", titleField.Elem().Interface())
					} else if titleField.Kind() != reflect.Ptr {
						title = fmt.Sprintf("%v", titleField.Interface())
					}
				}
			}
		}

		// Only add if we have both id and title
		if id != "" && title != "" {
			simpleFinding := SimpleFinding{
				ID:    id,
				Title: title,
			}
			simpleFindings = append(simpleFindings, simpleFinding)
		}
	}

	// Print the table
	if len(simpleFindings) == 0 {
		fmt.Println("No findings found.")
		return nil
	}

	// Calculate column widths
	maxIDWidth := len("ID")
	maxTitleWidth := len("TITLE")

	for _, finding := range simpleFindings {
		if len(finding.ID) > maxIDWidth {
			maxIDWidth = len(finding.ID)
		}
		if len(finding.Title) > maxTitleWidth {
			maxTitleWidth = len(finding.Title)
		}
	}

	// Ensure minimum widths
	if maxIDWidth < 10 {
		maxIDWidth = 10
	}
	if maxTitleWidth < 20 {
		maxTitleWidth = 20
	}

	// Print header
	header := fmt.Sprintf("| %-*s | %-*s |", maxIDWidth, "ID", maxTitleWidth, "TITLE")
	separator := fmt.Sprintf("|%s|%s|",
		strings.Repeat("-", maxIDWidth+2),
		strings.Repeat("-", maxTitleWidth+2))

	fmt.Println(separator)
	fmt.Println(header)
	fmt.Println(separator)

	// Print rows
	for _, finding := range simpleFindings {
		// Truncate title if too long
		displayTitle := finding.Title
		if len(displayTitle) > maxTitleWidth {
			displayTitle = displayTitle[:maxTitleWidth-3] + "..."
		}

		row := fmt.Sprintf("| %-*s | %-*s |", maxIDWidth, finding.ID, maxTitleWidth, displayTitle)
		fmt.Println(row)
	}

	fmt.Println(separator)
	log.GetLogger(verboseLevel).Info("Simple findings table formatted successfully\n")
	return nil
}

// printProjectTable prints project details in a formatted table
func PrintProjectTable(project interface{}, verboseLevel int) error {
	// Extract project fields
	var id, name, description, status, clientID string
	var labelIDs []string

	// Handle both map[string]interface{} and struct cases
	switch projectVal := project.(type) {
	case map[string]interface{}:
		// Handle raw JSON format (when unmarshaling fails)
		if result, ok := projectVal["result"].(map[string]interface{}); ok {
			// Extract from result field
			if idVal, ok := result["id"]; ok {
				id = fmt.Sprintf("%v", idVal)
			}
			if nameVal, ok := result["name"]; ok && nameVal != nil {
				name = fmt.Sprintf("%v", nameVal)
			}
			if descVal, ok := result["description"]; ok && descVal != nil {
				description = fmt.Sprintf("%v", descVal)
			}
			if statusVal, ok := result["status"]; ok {
				status = fmt.Sprintf("%v", statusVal)
			}
			if clientVal, ok := result["clientId"]; ok && clientVal != nil {
				clientID = fmt.Sprintf("%v", clientVal)
			}
			if labelsVal, ok := result["labelIds"]; ok {
				if labelsSlice, ok := labelsVal.([]interface{}); ok {
					for _, label := range labelsSlice {
						labelIDs = append(labelIDs, fmt.Sprintf("%v", label))
					}
				}
			}
		} else {
			// Handle direct map case
			if idVal, ok := projectVal["id"]; ok {
				id = fmt.Sprintf("%v", idVal)
			} else if idVal, ok := projectVal["ID"]; ok {
				id = fmt.Sprintf("%v", idVal)
			}

			if nameVal, ok := projectVal["name"]; ok && nameVal != nil {
				name = fmt.Sprintf("%v", nameVal)
			} else if nameVal, ok := projectVal["Name"]; ok && nameVal != nil {
				name = fmt.Sprintf("%v", nameVal)
			}

			if descVal, ok := projectVal["description"]; ok && descVal != nil {
				description = fmt.Sprintf("%v", descVal)
			} else if descVal, ok := projectVal["Description"]; ok && descVal != nil {
				description = fmt.Sprintf("%v", descVal)
			}

			if statusVal, ok := projectVal["status"]; ok {
				status = fmt.Sprintf("%v", statusVal)
			} else if statusVal, ok := projectVal["Status"]; ok {
				status = fmt.Sprintf("%v", statusVal)
			}

			if clientVal, ok := projectVal["clientId"]; ok && clientVal != nil {
				clientID = fmt.Sprintf("%v", clientVal)
			} else if clientVal, ok := projectVal["ClientID"]; ok && clientVal != nil {
				clientID = fmt.Sprintf("%v", clientVal)
			}

			if labelsVal, ok := projectVal["labelIds"]; ok {
				if labelsSlice, ok := labelsVal.([]interface{}); ok {
					for _, label := range labelsSlice {
						labelIDs = append(labelIDs, fmt.Sprintf("%v", label))
					}
				}
			} else if labelsVal, ok := projectVal["LabelIDs"]; ok {
				if labelsSlice, ok := labelsVal.([]interface{}); ok {
					for _, label := range labelsSlice {
						labelIDs = append(labelIDs, fmt.Sprintf("%v", label))
					}
				}
			}
		}
	default:
		// Handle struct case using reflection
		val := reflect.ValueOf(project)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}

		if val.Kind() == reflect.Struct {
			// Try to get ID field (case insensitive)
			idField := val.FieldByName("ID")
			if idField.IsValid() {
				id = fmt.Sprintf("%v", idField.Interface())
			}

			// Try to get Name field - handle pointer to string case
			nameField := val.FieldByName("Name")
			if nameField.IsValid() {
				if nameField.Kind() == reflect.Ptr && !nameField.IsNil() {
					name = fmt.Sprintf("%v", nameField.Elem().Interface())
				} else if nameField.Kind() != reflect.Ptr {
					name = fmt.Sprintf("%v", nameField.Interface())
				}
			}

			// Try to get Description field - handle pointer to string case
			descField := val.FieldByName("Description")
			if descField.IsValid() {
				if descField.Kind() == reflect.Ptr && !descField.IsNil() {
					description = fmt.Sprintf("%v", descField.Elem().Interface())
				} else if descField.Kind() != reflect.Ptr {
					description = fmt.Sprintf("%v", descField.Interface())
				}
			}

			// Try to get Status field
			statusField := val.FieldByName("Status")
			if statusField.IsValid() {
				status = fmt.Sprintf("%v", statusField.Interface())
			}

			// Try to get ClientID field - handle pointer to string case
			clientField := val.FieldByName("ClientID")
			if clientField.IsValid() {
				if clientField.Kind() == reflect.Ptr && !clientField.IsNil() {
					clientID = fmt.Sprintf("%v", clientField.Elem().Interface())
				} else if clientField.Kind() != reflect.Ptr {
					clientID = fmt.Sprintf("%v", clientField.Interface())
				}
			}

			// Try to get LabelIDs field
			labelsField := val.FieldByName("LabelIDs")
			if labelsField.IsValid() && labelsField.Kind() == reflect.Slice {
				for i := 0; i < labelsField.Len(); i++ {
					labelIDs = append(labelIDs, fmt.Sprintf("%v", labelsField.Index(i).Interface()))
				}
			}
		}
	}

	// Print the table
	fmt.Println("┌─────────────────────────────────────────────────────────────┐")
	fmt.Println("│                        PROJECT DETAILS                      │")
	fmt.Println("├─────────────────────────────────────────────────────────────┤")

	// Project ID
	fmt.Printf("│ %-15s │ %-40s │\n", "ID", truncateString(id, 40))
	fmt.Println("├─────────────────────────────────────────────────────────────┤")

	// Project Name
	fmt.Printf("│ %-15s │ %-40s │\n", "Name", truncateString(name, 40))
	fmt.Println("├─────────────────────────────────────────────────────────────┤")

	// Project Description
	if description != "" {
		// Split description into multiple lines if too long
		descLines := splitString(description, 40)
		for i, line := range descLines {
			if i == 0 {
				fmt.Printf("│ %-15s │ %-40s │\n", "Description", line)
			} else {
				fmt.Printf("│ %-15s │ %-40s │\n", "", line)
			}
		}
		fmt.Println("├─────────────────────────────────────────────────────────────┤")
	}

	// Project Status
	fmt.Printf("│ %-15s │ %-40s │\n", "Status", truncateString(status, 40))
	fmt.Println("├─────────────────────────────────────────────────────────────┤")

	// Client ID
	fmt.Printf("│ %-15s │ %-40s │\n", "Client ID", truncateString(clientID, 40))
	fmt.Println("├─────────────────────────────────────────────────────────────┤")

	// Label IDs
	if len(labelIDs) > 0 {
		labelsStr := strings.Join(labelIDs, ", ")
		labelLines := splitString(labelsStr, 40)
		for i, line := range labelLines {
			if i == 0 {
				fmt.Printf("│ %-15s │ %-40s │\n", "Label IDs", line)
			} else {
				fmt.Printf("│ %-15s │ %-40s │\n", "", line)
			}
		}
	} else {
		fmt.Printf("│ %-15s │ %-40s │\n", "Label IDs", "None")
	}

	fmt.Println("└─────────────────────────────────────────────────────────────┘")
	log.GetLogger(verboseLevel).Info("Project table formatted successfully\n")
	return nil
}

// printFindingTable prints a formatted table view of a finding
func PrintFindingTable(finding interface{}, verboseLevel int) error {
	// Use reflection to safely access the finding data
	val := reflect.ValueOf(finding)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Extract finding fields
	var id, name, description, projectID string
	var severity, status int

	// Handle both map[string]interface{} and struct cases
	switch findingVal := finding.(type) {
	case map[string]interface{}:
		if v, ok := findingVal["id"]; ok {
			id = fmt.Sprintf("%v", v)
		}
		if v, ok := findingVal["name"]; ok {
			name = fmt.Sprintf("%v", v)
		}
		if v, ok := findingVal["description"]; ok {
			description = fmt.Sprintf("%v", v)
		}
		if v, ok := findingVal["severity"]; ok {
			if s, ok := v.(int); ok {
				severity = s
			} else if s, ok := v.(float64); ok {
				severity = int(s)
			}
		}
		if v, ok := findingVal["status"]; ok {
			if s, ok := v.(int); ok {
				status = s
			} else if s, ok := v.(float64); ok {
				status = int(s)
			}
		}
		if v, ok := findingVal["projectId"]; ok {
			projectID = fmt.Sprintf("%v", v)
		}
	default:
		// Try to access struct fields
		if idField := val.FieldByName("ID"); idField.IsValid() {
			id = idField.String()
		}
		if nameField := val.FieldByName("Name"); nameField.IsValid() {
			name = nameField.String()
		}
		if descField := val.FieldByName("Description"); descField.IsValid() {
			description = descField.String()
		}
		if severityField := val.FieldByName("Severity"); severityField.IsValid() {
			severity = int(severityField.Int())
		}
		if statusField := val.FieldByName("Status"); statusField.IsValid() {
			status = int(statusField.Int())
		}
		if projectIDField := val.FieldByName("ProjectID"); projectIDField.IsValid() {
			projectID = projectIDField.String()
		}
	}

	if id == "" {
		fmt.Println("No finding information available")
		return nil
	}

	fmt.Println("┌─────────────────────────────────────────────────────────────┐")
	fmt.Println("│                        FINDING INFO                        │")
	fmt.Println("├─────────────────────────────────────────────────────────────┤")
	fmt.Printf("│ ID:          %-45s │\n", id)
	fmt.Printf("│ Name:        %-45s │\n", name)
	fmt.Printf("│ Description: %-45s │\n", description)
	fmt.Printf("│ Severity:    %-45d │\n", severity)
	fmt.Printf("│ Status:      %-45d │\n", status)
	fmt.Printf("│ Project ID:  %-45s │\n", projectID)
	fmt.Println("└─────────────────────────────────────────────────────────────┘")
	return nil
}

// Helper functions for string formatting
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

func splitString(s string, maxLen int) []string {
	if len(s) <= maxLen {
		return []string{s}
	}

	var lines []string
	words := strings.Fields(s)
	currentLine := ""

	for _, word := range words {
		// If a single word is longer than maxLen, we need to break it
		if len(word) > maxLen {
			// First, add the current line if it has content
			if currentLine != "" {
				lines = append(lines, currentLine)
				currentLine = ""
			}

			// Break the long word into chunks
			for len(word) > maxLen {
				lines = append(lines, word[:maxLen])
				word = word[maxLen:]
			}
			if word != "" {
				currentLine = word
			}
		} else if len(currentLine)+len(word)+1 <= maxLen {
			// Normal case: add word to current line
			if currentLine == "" {
				currentLine = word
			} else {
				currentLine += " " + word
			}
		} else {
			// Current line is full, start a new one
			if currentLine != "" {
				lines = append(lines, currentLine)
			}
			currentLine = word
		}
	}

	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	return lines
}

// printPentesterInfoTable prints pentester information in table format
func PrintPentesterInfoTable(pentester interface{}, verboseLevel int) error {
	// Use reflection to safely access the data structure
	val := reflect.ValueOf(pentester)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Print header
	fmt.Println("┌─────────────────────────────────────────────────────────────┐")
	fmt.Println("│                    Pentester Information                    │")
	fmt.Println("├─────────────────────────────────────────────────────────────┤")

	// Print fields
	if idField := val.FieldByName("ID"); idField.IsValid() {
		fmt.Printf("│ ID:          %-45s │\n", truncateString(fmt.Sprintf("%v", idField.Interface()), 45))
	}
	if nameField := val.FieldByName("Name"); nameField.IsValid() {
		fmt.Printf("│ Name:        %-45s │\n", truncateString(fmt.Sprintf("%v", nameField.Interface()), 45))
	}
	if emailField := val.FieldByName("Email"); emailField.IsValid() {
		fmt.Printf("│ Email:       %-45s │\n", truncateString(fmt.Sprintf("%v", emailField.Interface()), 45))
	}
	if descField := val.FieldByName("Description"); descField.IsValid() {
		fmt.Printf("│ Description: %-45s │\n", truncateString(fmt.Sprintf("%v", descField.Interface()), 45))
	}

	fmt.Println("└─────────────────────────────────────────────────────────────┘")
	return nil
}

// printClientsTable prints a list of clients in table format
func PrintClientsTable(clients interface{}, verboseLevel int) error {
	// Use reflection to safely access the data structure
	val := reflect.ValueOf(clients)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Handle ClientDtoPagedResultDtoAjaxResponse structure
	var items reflect.Value
	if val.Kind() == reflect.Struct {
		// Try to access the Result field
		resultField := val.FieldByName("Result")
		if !resultField.IsValid() {
			return fmt.Errorf("no Result field found in response data")
		}

		if resultField.Kind() == reflect.Ptr {
			resultField = resultField.Elem()
		}

		// Try to access the Items field from Result
		itemsField := resultField.FieldByName("Items")
		if !itemsField.IsValid() {
			return fmt.Errorf("no Items field found in Result")
		}

		items = itemsField
	} else if val.Kind() == reflect.Slice {
		items = val
	} else {
		return fmt.Errorf("expected slice or ClientDtoPagedResultDtoAjaxResponse, got %T", clients)
	}

	if items.Kind() != reflect.Slice {
		return fmt.Errorf("items field is not a slice")
	}

	if items.Len() == 0 {
		fmt.Println("No clients found.")
		return nil
	}

	// Print header
	fmt.Println("┌─────────────────────────────────────────────────────────────────────────────────────┐")
	fmt.Println("│                                        Clients                                        │")
	fmt.Println("├─────────────────────────────────────────────────────────────────────────────────────┤")
	fmt.Printf("│ %-36s │ %-20s │ %-30s │ %-10s │\n", "ID", "Name", "Email", "Status")
	fmt.Println("├─────────────────────────────────────────────────────────────────────────────────────┤")

	// Print each client
	for i := 0; i < items.Len(); i++ {
		client := items.Index(i)
		if client.Kind() == reflect.Ptr {
			client = client.Elem()
		}

		var id, name, email, status string

		if idField := client.FieldByName("ID"); idField.IsValid() {
			id = fmt.Sprintf("%v", idField.Interface())
		}
		if nameField := client.FieldByName("Name"); nameField.IsValid() {
			if nameField.Kind() == reflect.Ptr && !nameField.IsNil() {
				name = fmt.Sprintf("%v", nameField.Elem().Interface())
			} else if nameField.Kind() != reflect.Ptr {
				name = fmt.Sprintf("%v", nameField.Interface())
			}
		}
		if emailField := client.FieldByName("Email"); emailField.IsValid() {
			if emailField.Kind() == reflect.Ptr && !emailField.IsNil() {
				email = fmt.Sprintf("%v", emailField.Elem().Interface())
			} else if emailField.Kind() != reflect.Ptr {
				email = fmt.Sprintf("%v", emailField.Interface())
			}
		}
		if statusField := client.FieldByName("Status"); statusField.IsValid() {
			if statusField.Kind() == reflect.Ptr && !statusField.IsNil() {
				status = fmt.Sprintf("%v", statusField.Elem().Interface())
			} else if statusField.Kind() != reflect.Ptr {
				status = fmt.Sprintf("%v", statusField.Interface())
			}
		}

		fmt.Printf("│ %-36s │ %-20s │ %-30s │ %-10s │\n",
			truncateString(id, 36),
			truncateString(name, 20),
			truncateString(email, 30),
			truncateString(status, 10))
	}

	fmt.Println("└─────────────────────────────────────────────────────────────────────────────────────┘")
	return nil
}

// printPentestersTable prints a list of pentesters in table format
func PrintPentestersTable(pentesters interface{}, verboseLevel int) error {
	// Use reflection to safely access the data structure
	val := reflect.ValueOf(pentesters)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Handle ClientDtoPagedResultDtoAjaxResponse structure
	var items reflect.Value
	if val.Kind() == reflect.Struct {
		// Try to access the Result field
		resultField := val.FieldByName("Result")
		if !resultField.IsValid() {
			return fmt.Errorf("no Result field found in response data")
		}

		if resultField.Kind() == reflect.Ptr {
			resultField = resultField.Elem()
		}

		// Try to access the Items field from Result
		itemsField := resultField.FieldByName("Items")
		if !itemsField.IsValid() {
			return fmt.Errorf("no Items field found in Result")
		}

		items = itemsField
	} else if val.Kind() == reflect.Slice {
		items = val
	} else {
		return fmt.Errorf("expected slice or ClientDtoPagedResultDtoAjaxResponse, got %T", pentesters)
	}

	if items.Kind() != reflect.Slice {
		return fmt.Errorf("items field is not a slice")
	}

	if items.Len() == 0 {
		fmt.Println("No pentesters found.")
		return nil
	}

	// Print header
	fmt.Println("┌─────────────────────────────────────────────────────────────────────────────────────┐")
	fmt.Println("│                                    Pentesters                                        │")
	fmt.Println("├─────────────────────────────────────────────────────────────────────────────────────┤")
	fmt.Printf("│ %-36s │ %-20s │ %-30s │\n", "ID", "Name", "Email")
	fmt.Println("├─────────────────────────────────────────────────────────────────────────────────────┤")

	// Print each pentester
	for i := 0; i < items.Len(); i++ {
		pentester := items.Index(i)
		if pentester.Kind() == reflect.Ptr {
			pentester = pentester.Elem()
		}

		var id, name, email string

		if idField := pentester.FieldByName("ID"); idField.IsValid() {
			id = fmt.Sprintf("%v", idField.Interface())
		}
		if nameField := pentester.FieldByName("Name"); nameField.IsValid() {
			if nameField.Kind() == reflect.Ptr && !nameField.IsNil() {
				name = fmt.Sprintf("%v", nameField.Elem().Interface())
			} else if nameField.Kind() != reflect.Ptr {
				name = fmt.Sprintf("%v", nameField.Interface())
			}
		}
		if emailField := pentester.FieldByName("Email"); emailField.IsValid() {
			if emailField.Kind() == reflect.Ptr && !emailField.IsNil() {
				email = fmt.Sprintf("%v", emailField.Elem().Interface())
			} else if emailField.Kind() != reflect.Ptr {
				email = fmt.Sprintf("%v", emailField.Interface())
			}
		}

		fmt.Printf("│ %-36s │ %-20s │ %-30s │\n",
			truncateString(id, 36),
			truncateString(name, 20),
			truncateString(email, 30))
	}

	fmt.Println("└─────────────────────────────────────────────────────────────────────────────────────┘")
	return nil
}

// printChecklistsTable prints a list of checklists in table format
func PrintChecklistsTable(checklists interface{}, verboseLevel int) error {
	// Use reflection to safely access the data structure
	val := reflect.ValueOf(checklists)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Slice {
		return fmt.Errorf("expected slice, got %T", checklists)
	}

	if val.Len() == 0 {
		fmt.Println("No checklists found.")
		return nil
	}

	// Print header
	fmt.Println("┌─────────────────────────────────────────────────────────────────────────────────────┐")
	fmt.Println("│                                    Checklists                                        │")
	fmt.Println("├─────────────────────────────────────────────────────────────────────────────────────┤")
	fmt.Printf("│ %-36s │ %-20s │ %-30s │\n", "ID", "Name", "Status")
	fmt.Println("├─────────────────────────────────────────────────────────────────────────────────────┤")

	// Print each checklist
	for i := 0; i < val.Len(); i++ {
		checklist := val.Index(i)
		if checklist.Kind() == reflect.Ptr {
			checklist = checklist.Elem()
		}

		var id, name, status string

		if idField := checklist.FieldByName("ID"); idField.IsValid() {
			id = fmt.Sprintf("%v", idField.Interface())
		}
		if nameField := checklist.FieldByName("Name"); nameField.IsValid() {
			name = fmt.Sprintf("%v", nameField.Interface())
		}
		if statusField := checklist.FieldByName("Status"); statusField.IsValid() {
			status = fmt.Sprintf("%v", statusField.Interface())
		}

		fmt.Printf("│ %-36s │ %-20s │ %-30s │\n",
			truncateString(id, 36),
			truncateString(name, 20),
			truncateString(status, 30))
	}

	fmt.Println("└─────────────────────────────────────────────────────────────────────────────────────┘")
	return nil
}

// printComplianceNormsTable prints a list of compliance norms in table format
func PrintComplianceNormsTable(complianceNorms interface{}, verboseLevel int) error {
	// Use reflection to safely access the data structure
	val := reflect.ValueOf(complianceNorms)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Slice {
		return fmt.Errorf("expected slice, got %T", complianceNorms)
	}

	if val.Len() == 0 {
		fmt.Println("No compliance norms found.")
		return nil
	}

	// Print header
	fmt.Println("┌─────────────────────────────────────────────────────────────────────────────────────┐")
	fmt.Println("│                                 Compliance Norms                                     │")
	fmt.Println("├─────────────────────────────────────────────────────────────────────────────────────┤")
	fmt.Printf("│ %-36s │ %-20s │ %-30s │\n", "ID", "Name", "Status")
	fmt.Println("├─────────────────────────────────────────────────────────────────────────────────────┤")

	// Print each compliance norm
	for i := 0; i < val.Len(); i++ {
		norm := val.Index(i)
		if norm.Kind() == reflect.Ptr {
			norm = norm.Elem()
		}

		var id, name, status string

		if idField := norm.FieldByName("ID"); idField.IsValid() {
			id = fmt.Sprintf("%v", idField.Interface())
		}
		if nameField := norm.FieldByName("Name"); nameField.IsValid() {
			name = fmt.Sprintf("%v", nameField.Interface())
		}
		if statusField := norm.FieldByName("Status"); statusField.IsValid() {
			status = fmt.Sprintf("%v", statusField.Interface())
		}

		fmt.Printf("│ %-36s │ %-20s │ %-30s │\n",
			truncateString(id, 36),
			truncateString(name, 20),
			truncateString(status, 30))
	}

	fmt.Println("└─────────────────────────────────────────────────────────────────────────────────────┘")
	return nil
}

// printReportVersionsTable prints a list of report versions in table format
func PrintReportVersionsTable(reportVersions interface{}, verboseLevel int) error {
	// Use reflection to safely access the data structure
	val := reflect.ValueOf(reportVersions)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Slice {
		return fmt.Errorf("expected slice, got %T", reportVersions)
	}

	if val.Len() == 0 {
		fmt.Println("No report versions found.")
		return nil
	}

	// Print header
	fmt.Println("┌─────────────────────────────────────────────────────────────────────────────────────┐")
	fmt.Println("│                                 Report Versions                                      │")
	fmt.Println("├─────────────────────────────────────────────────────────────────────────────────────┤")
	fmt.Printf("│ %-36s │ %-15s │ %-20s │ %-15s │\n", "ID", "Version", "Created At", "Status")
	fmt.Println("├─────────────────────────────────────────────────────────────────────────────────────┤")

	// Print each report version
	for i := 0; i < val.Len(); i++ {
		version := val.Index(i)
		if version.Kind() == reflect.Ptr {
			version = version.Elem()
		}

		var id, versionStr, createdAt, status string

		if idField := version.FieldByName("ID"); idField.IsValid() {
			id = fmt.Sprintf("%v", idField.Interface())
		}
		if versionField := version.FieldByName("Version"); versionField.IsValid() {
			versionStr = fmt.Sprintf("%v", versionField.Interface())
		}
		if createdAtField := version.FieldByName("CreatedAt"); createdAtField.IsValid() {
			createdAt = fmt.Sprintf("%v", createdAtField.Interface())
		}
		if statusField := version.FieldByName("Status"); statusField.IsValid() {
			status = fmt.Sprintf("%v", statusField.Interface())
		}

		fmt.Printf("│ %-36s │ %-15s │ %-20s │ %-15s │\n",
			truncateString(id, 36),
			truncateString(versionStr, 15),
			truncateString(createdAt, 20),
			truncateString(status, 15))
	}

	fmt.Println("└─────────────────────────────────────────────────────────────────────────────────────┘")
	return nil
}

// printReportTable prints report information in table format
func PrintReportTable(report interface{}, verboseLevel int) error {
	// Use reflection to safely access the data structure
	val := reflect.ValueOf(report)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Print header
	fmt.Println("┌─────────────────────────────────────────────────────────────┐")
	fmt.Println("│                      Report Information                     │")
	fmt.Println("├─────────────────────────────────────────────────────────────┤")

	// Print fields
	if idField := val.FieldByName("ID"); idField.IsValid() {
		fmt.Printf("│ ID:          %-45s │\n", truncateString(fmt.Sprintf("%v", idField.Interface()), 45))
	}
	if versionField := val.FieldByName("Version"); versionField.IsValid() {
		fmt.Printf("│ Version:     %-45s │\n", truncateString(fmt.Sprintf("%v", versionField.Interface()), 45))
	}
	if createdAtField := val.FieldByName("CreatedAt"); createdAtField.IsValid() {
		fmt.Printf("│ Created At:  %-45s │\n", truncateString(fmt.Sprintf("%v", createdAtField.Interface()), 45))
	}
	if publishedAtField := val.FieldByName("PublishedAt"); publishedAtField.IsValid() {
		fmt.Printf("│ Published:   %-45s │\n", truncateString(fmt.Sprintf("%v", publishedAtField.Interface()), 45))
	}
	if contentField := val.FieldByName("Content"); contentField.IsValid() {
		content := fmt.Sprintf("%v", contentField.Interface())
		if len(content) > 45 {
			content = content[:42] + "..."
		}
		fmt.Printf("│ Content:     %-45s │\n", content)
	}

	fmt.Println("└─────────────────────────────────────────────────────────────┘")
	return nil
}

// CustomTableConfig holds configuration for custom table output
type CustomTableConfig struct {
	MaxColumns     int
	SelectedFields []string
}

// printCustomTable creates an interactive custom table with user-selected fields
func PrintCustomTable(data interface{}, maxColumns int, verboseLevel int) error {
	// Extract items from the response data
	items, err := extractItemsFromResponse(data)
	if err != nil {
		return fmt.Errorf("failed to extract items from response: %w", err)
	}

	if len(items) == 0 {
		fmt.Println("No data found to display.")
		return nil
	}

	// Get available fields from the first item
	availableFields := getAvailableFields(items[0])
	if len(availableFields) == 0 {
		return fmt.Errorf("no fields found in data")
	}

	// Let user select fields
	selectedFields, err := selectFields(availableFields, maxColumns)
	if err != nil {
		return fmt.Errorf("failed to select fields: %w", err)
	}

	if len(selectedFields) == 0 {
		fmt.Println("No fields selected.")
		return nil
	}

	// Generate and print the custom table
	return generateCustomTable(items, selectedFields, verboseLevel)
}

// extractItemsFromResponse extracts items from various response formats
func extractItemsFromResponse(data interface{}) ([]interface{}, error) {
	var items []interface{}

	// Try to handle both struct format and raw JSON format
	if dataMap, ok := data.(map[string]interface{}); ok {
		// Handle raw JSON format (when unmarshaling fails)
		if result, ok := dataMap["result"].(map[string]interface{}); ok {
			if itemsList, ok := result["items"].([]interface{}); ok {
				items = itemsList
			}
		} else if itemsList, ok := dataMap["items"].([]interface{}); ok {
			items = itemsList
		}
	} else {
		// Handle struct format (when unmarshaling succeeds)
		val := reflect.ValueOf(data)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}

		// Try to access the Result field which should contain the paged data
		resultField := val.FieldByName("Result")
		if resultField.IsValid() {
			if resultField.Kind() == reflect.Ptr {
				resultField = resultField.Elem()
			}

			// Try to access the Items field which should contain the list
			itemsField := resultField.FieldByName("Items")
			if itemsField.IsValid() {
				// Convert items to a slice of interfaces
				if itemsField.Kind() == reflect.Slice {
					for i := 0; i < itemsField.Len(); i++ {
						items = append(items, itemsField.Index(i).Interface())
					}
				}
			}
		} else {
			// Check if the data itself is a slice
			if val.Kind() == reflect.Slice {
				for i := 0; i < val.Len(); i++ {
					items = append(items, val.Index(i).Interface())
				}
			}
		}
	}

	return items, nil
}

// getAvailableFields extracts all available field names from an item
func getAvailableFields(item interface{}) []string {
	var fields []string

	switch itemVal := item.(type) {
	case map[string]interface{}:
		// Handle map case
		for fieldName := range itemVal {
			fields = append(fields, fieldName)
		}
	default:
		// Handle struct case using reflection
		itemReflect := reflect.ValueOf(item)
		if itemReflect.Kind() == reflect.Ptr {
			itemReflect = itemReflect.Elem()
		}

		if itemReflect.Kind() == reflect.Struct {
			for i := 0; i < itemReflect.NumField(); i++ {
				field := itemReflect.Type().Field(i)
				fields = append(fields, field.Name)
			}
		}
	}

	return fields
}

// selectFields allows user to interactively select fields for the table
func selectFields(availableFields []string, maxColumns int) ([]string, error) {
	if len(availableFields) == 0 {
		return nil, fmt.Errorf("no fields available")
	}

	fmt.Printf("\nAvailable fields (%d total):\n", len(availableFields))
	for i, field := range availableFields {
		fmt.Printf("  %d. %s\n", i+1, field)
	}

	fmt.Printf("\nSelect up to %d fields for your custom table:\n", maxColumns)
	fmt.Println("Enter field numbers separated by commas (e.g., 1,3,5) or 'all' for all fields:")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("failed to read input: %w", err)
	}

	input = strings.TrimSpace(input)

	if input == "all" {
		if len(availableFields) > maxColumns {
			fmt.Printf("Warning: You selected all fields (%d), but maximum is %d. Showing first %d fields.\n",
				len(availableFields), maxColumns, maxColumns)
			return availableFields[:maxColumns], nil
		}
		return availableFields, nil
	}

	// Parse comma-separated numbers
	parts := strings.Split(input, ",")
	var selectedFields []string
	selectedIndices := make(map[int]bool)

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		index, err := strconv.Atoi(part)
		if err != nil {
			fmt.Printf("Invalid input '%s'. Please enter numbers only.\n", part)
			continue
		}

		if index < 1 || index > len(availableFields) {
			fmt.Printf("Invalid field number %d. Please select between 1 and %d.\n", index, len(availableFields))
			continue
		}

		if !selectedIndices[index-1] {
			selectedIndices[index-1] = true
			selectedFields = append(selectedFields, availableFields[index-1])
		}
	}

	if len(selectedFields) > maxColumns {
		fmt.Printf("Warning: You selected %d fields, but maximum is %d. Showing first %d fields.\n",
			len(selectedFields), maxColumns, maxColumns)
		selectedFields = selectedFields[:maxColumns]
	}

	if len(selectedFields) == 0 {
		return nil, fmt.Errorf("no valid fields selected")
	}

	return selectedFields, nil
}

// generateCustomTable creates and prints the custom table with selected fields
func generateCustomTable(items []interface{}, selectedFields []string, verboseLevel int) error {
	if len(items) == 0 {
		fmt.Println("No data to display.")
		return nil
	}

	// Calculate column widths
	columnWidths := make([]int, len(selectedFields))

	// Initialize with header widths
	for i, field := range selectedFields {
		columnWidths[i] = len(field)
	}

	// Calculate maximum width for each column
	for _, item := range items {
		for i, field := range selectedFields {
			value := getFieldValue(item, field)
			valueStr := fmt.Sprintf("%v", value)
			if len(valueStr) > columnWidths[i] {
				columnWidths[i] = len(valueStr)
			}
		}
	}

	// Ensure minimum widths
	for i := range columnWidths {
		if columnWidths[i] < 10 {
			columnWidths[i] = 10
		}
	}

	// Print table header
	fmt.Println()
	printTableSeparator(columnWidths)
	printTableHeader(selectedFields, columnWidths)
	printTableSeparator(columnWidths)

	// Print table rows
	for _, item := range items {
		printTableRow(item, selectedFields, columnWidths)
	}

	printTableSeparator(columnWidths)
	log.GetLogger(verboseLevel).Info("Custom table formatted successfully\n")
	return nil
}

// getFieldValue extracts the value of a field from an item
func getFieldValue(item interface{}, fieldName string) interface{} {
	switch itemVal := item.(type) {
	case map[string]interface{}:
		// Handle map case
		if value, ok := itemVal[fieldName]; ok {
			return dereferenceValue(value)
		}
		// Try case-insensitive match
		for key, value := range itemVal {
			if strings.EqualFold(key, fieldName) {
				return dereferenceValue(value)
			}
		}
	default:
		// Handle struct case using reflection
		itemReflect := reflect.ValueOf(item)
		if itemReflect.Kind() == reflect.Ptr {
			itemReflect = itemReflect.Elem()
		}

		if itemReflect.Kind() == reflect.Struct {
			// Try exact match first
			field := itemReflect.FieldByName(fieldName)
			if field.IsValid() {
				return dereferenceValue(field.Interface())
			}

			// Try case-insensitive match
			for i := 0; i < itemReflect.NumField(); i++ {
				structField := itemReflect.Type().Field(i)
				if strings.EqualFold(structField.Name, fieldName) {
					return dereferenceValue(itemReflect.Field(i).Interface())
				}
			}
		}
	}

	return ""
}

// dereferenceValue safely dereferences pointers and returns the actual value
func dereferenceValue(value interface{}) interface{} {
	if value == nil {
		return ""
	}

	val := reflect.ValueOf(value)

	// If it's a pointer, dereference it
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return ""
		}
		val = val.Elem()
	}

	// Handle different types appropriately
	switch val.Kind() {
	case reflect.String:
		return val.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return val.Uint()
	case reflect.Float32, reflect.Float64:
		return val.Float()
	case reflect.Bool:
		return val.Bool()
	case reflect.Slice, reflect.Array:
		// For slices/arrays, convert to a readable format
		if val.Len() == 0 {
			return "[]"
		}
		var elements []string
		for i := 0; i < val.Len(); i++ {
			elements = append(elements, fmt.Sprintf("%v", dereferenceValue(val.Index(i).Interface())))
		}
		return fmt.Sprintf("[%s]", strings.Join(elements, ", "))
	case reflect.Map:
		// For maps, convert to a readable format
		if val.Len() == 0 {
			return "{}"
		}
		var pairs []string
		for _, key := range val.MapKeys() {
			pairs = append(pairs, fmt.Sprintf("%v: %v", dereferenceValue(key.Interface()), dereferenceValue(val.MapIndex(key).Interface())))
		}
		return fmt.Sprintf("{%s}", strings.Join(pairs, ", "))
	case reflect.Struct:
		// For structs, try to get a meaningful string representation
		// Check if it has a String() method
		if val.CanInterface() {
			if stringer, ok := val.Interface().(fmt.Stringer); ok {
				return stringer.String()
			}
		}
		// If no String() method, return the type name
		return val.Type().Name()
	default:
		// For any other type, try to convert to string
		return fmt.Sprintf("%v", val.Interface())
	}
}

// printTableSeparator prints a separator line for the table
func printTableSeparator(columnWidths []int) {
	fmt.Print("|")
	for _, width := range columnWidths {
		fmt.Printf("%s|", strings.Repeat("-", width+2))
	}
	fmt.Println()
}

// printTableHeader prints the table header
func printTableHeader(fields []string, columnWidths []int) {
	fmt.Print("|")
	for i, field := range fields {
		fmt.Printf(" %-*s |", columnWidths[i], strings.ToUpper(field))
	}
	fmt.Println()
}

// printTableRow prints a single table row
func printTableRow(item interface{}, fields []string, columnWidths []int) {
	fmt.Print("|")
	for i, field := range fields {
		value := getFieldValue(item, field)
		valueStr := fmt.Sprintf("%v", value)

		// Truncate if too long
		if len(valueStr) > columnWidths[i] {
			valueStr = valueStr[:columnWidths[i]-3] + "..."
		}

		fmt.Printf(" %-*s |", columnWidths[i], valueStr)
	}
	fmt.Println()
}
