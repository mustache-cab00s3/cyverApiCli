package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yourusername/cyverApiCli/internal/api/services"
)

var templatesExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Get finding templates available for export",
	Long: `Calls a non-documented web service endpoint:
/api/services/app/FindingLibrary/GetFindingTemplatesToExport

This endpoint is non-supported and intended for specific instances only.`,
	RunE: runTemplatesExport,
}

var templatesExportJSONCmd = &cobra.Command{
	Use:     "export-json",
	Aliases: []string{"exj"},
	Short:   "Download all finding templates in JSON array format",
	Long: `Downloads all finding templates in a library as a JSON array.

This command uses the non-documented web service endpoint:
/api/services/app/FindingLibrary/GetFindingTemplates`,
	RunE: runTemplatesExportJSON,
}

var templatesLibrariesCmd = &cobra.Command{
	Use:   "libraries",
	Short: "List finding template libraries",
	Long: `Calls a non-documented web service endpoint:
/api/services/app/FindingLibrary/GetFindingLibraries

This endpoint is non-supported and intended for specific instances only.`,
	RunE: runTemplatesLibraries,
}

var templatesDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download exported finding templates temp file",
	Long: `Calls a non-documented web endpoint:
/File/DownloadTempFile

This endpoint is non-supported and intended for specific instances only.`,
	RunE: runTemplatesDownload,
}

var templatesSaveTemplateCmd = &cobra.Command{
	Use:   "save-template",
	Short: "Create or update a finding template",
	Long: `Calls a non-documented web endpoint:
/api/services/app/FindingLibrary/CreateOrUpdateFindingTemplate

Provide request body via --data-json or --data-file (JSON).`,
	RunE: runTemplatesSaveTemplate,
}

var templatesSaveLibraryCmd = &cobra.Command{
	Use:   "save-library",
	Short: "Create or update a finding library",
	Long: `Calls a non-documented web endpoint:
/api/services/app/FindingLibrary/CreateOrUpdateFindingLibrary

Provide request body via --data-json or --data-file (JSON).`,
	RunE: runTemplatesSaveLibrary,
}

func init() {
	templatesExportCmd.Flags().StringP("finding-library-id", "l", "", "finding library UUID")
	templatesExportCmd.Flags().BoolP("download", "d", false, "automatically download the exported file when file metadata is returned")
	templatesExportCmd.Flags().String("download-type", "xlsx", "download format when --download is set (xlsx or json)")
	templatesExportCmd.Flags().StringP("output", "o", "", "optional output path for --download (defaults to current directory + file-name)")
	_ = templatesExportCmd.MarkFlagRequired("finding-library-id")
	templatesCmd.AddCommand(templatesExportCmd)

	templatesExportJSONCmd.Flags().StringP("finding-library-id", "l", "", "finding library UUID")
	templatesExportJSONCmd.Flags().StringP("output", "o", "", "optional output path (defaults to current directory + finding_templates_<library-id>.json)")
	_ = templatesExportJSONCmd.MarkFlagRequired("finding-library-id")
	templatesCmd.AddCommand(templatesExportJSONCmd)

	templatesLibrariesCmd.Flags().StringP("filter", "f", "", "optional search filter")
	templatesLibrariesCmd.Flags().IntP("max-result-count", "m", 10, "max number of results")
	templatesLibrariesCmd.Flags().IntP("skip-count", "s", 0, "number of records to skip")
	templatesCmd.AddCommand(templatesLibrariesCmd)

	templatesDownloadCmd.Flags().StringP("file-token", "t", "", "temporary download token from templates export")
	templatesDownloadCmd.Flags().StringP("file-name", "n", "", "download file name from templates export")
	templatesDownloadCmd.Flags().StringP("file-type", "y", "", "download MIME type from templates export")
	templatesDownloadCmd.Flags().StringP("output", "o", "", "optional output path (defaults to current directory + file-name)")
	_ = templatesDownloadCmd.MarkFlagRequired("file-token")
	_ = templatesDownloadCmd.MarkFlagRequired("file-name")
	_ = templatesDownloadCmd.MarkFlagRequired("file-type")
	templatesCmd.AddCommand(templatesDownloadCmd)

	templatesSaveTemplateCmd.Flags().StringP("data-json", "j", "", "inline JSON payload for create/update request")
	templatesSaveTemplateCmd.Flags().StringP("data-file", "f", "", "path to JSON file payload for create/update request")
	templatesCmd.AddCommand(templatesSaveTemplateCmd)

	templatesSaveLibraryCmd.Flags().StringP("data-json", "j", "", "inline JSON payload for create/update library request")
	templatesSaveLibraryCmd.Flags().StringP("data-file", "f", "", "path to JSON file payload for create/update library request")
	templatesCmd.AddCommand(templatesSaveLibraryCmd)
}

func newTemplatesServiceClient() (*services.NonSupportedServiceClient, error) {
	_, baseURL, _, err := LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	timeoutSeconds := viper.GetInt("client.timeout")
	if timeoutSeconds <= 0 {
		timeoutSeconds = 30
	}

	client := services.NewNonSupportedServiceClient(baseURL, time.Duration(timeoutSeconds)*time.Second)
	client.Token = viper.GetString("token.access_token")
	return client, nil
}

func runTemplatesExport(cmd *cobra.Command, args []string) error {
	findingLibraryID, _ := cmd.Flags().GetString("finding-library-id")
	autoDownload, _ := cmd.Flags().GetBool("download")
	downloadType, _ := cmd.Flags().GetString("download-type")
	exportOutputPath, _ := cmd.Flags().GetString("output")
	downloadType = strings.ToLower(strings.TrimSpace(downloadType))
	if downloadType == "" {
		downloadType = "xlsx"
	}
	if downloadType != "xlsx" && downloadType != "json" {
		return fmt.Errorf("invalid --download-type %q (supported: xlsx, json)", downloadType)
	}

	client, err := newTemplatesServiceClient()
	if err != nil {
		return err
	}

	responseModel, body, _, err := client.GetFindingTemplatesToExportModel(context.Background(), findingLibraryID)
	if err != nil {
		if len(body) > 0 {
			return fmt.Errorf("%w: %s", err, string(body))
		}
		return err
	}

	pretty, err := json.MarshalIndent(responseModel, "", "  ")
	if err != nil {
		fmt.Println(string(body))
		return nil
	}
	fmt.Println(string(pretty))

	if responseModel.Result != nil {
		fmt.Println()
		fmt.Println("Download metadata:")
		fmt.Printf("  fileName : %s\n", responseModel.Result.FileName)
		fmt.Printf("  fileType : %s\n", responseModel.Result.FileType)
		fmt.Printf("  fileToken: %s\n", responseModel.Result.FileToken)

		if !autoDownload {
			fmt.Println()
			fmt.Println("Next step: use this fileToken with templates download, or rerun with --download.")
			return nil
		}

		if downloadType == "json" {
			allTemplates, err := fetchAllFindingTemplatesForLibrary(context.Background(), client, findingLibraryID)
			if err != nil {
				return fmt.Errorf("json export failed: %w", err)
			}

			jsonBytes, err := json.MarshalIndent(allTemplates, "", "  ")
			if err != nil {
				return fmt.Errorf("encode json export: %w", err)
			}

			if strings.TrimSpace(exportOutputPath) == "" {
				exportOutputPath = filepath.Join(".", fmt.Sprintf("finding_templates_%s.json", findingLibraryID))
			}
			if err := os.WriteFile(exportOutputPath, jsonBytes, 0600); err != nil {
				return fmt.Errorf("write json export file: %w", err)
			}

			fmt.Println()
			fmt.Printf("Auto-downloaded %s (%d templates)\n", exportOutputPath, len(allTemplates))
			return nil
		}

		downloadReq := services.TempFileDownloadRequest{
			FileToken: responseModel.Result.FileToken,
			FileName:  responseModel.Result.FileName,
			FileType:  responseModel.Result.FileType,
		}
		result, _, err := client.DownloadTempFile(context.Background(), downloadReq)
		if err != nil {
			return fmt.Errorf("auto-download failed: %w", err)
		}

		if strings.TrimSpace(exportOutputPath) == "" {
			exportOutputPath = filepath.Join(".", result.FileName)
		}
		if err := os.WriteFile(exportOutputPath, result.Body, 0600); err != nil {
			return fmt.Errorf("write auto-downloaded file: %w", err)
		}

		fmt.Println()
		fmt.Printf("Auto-downloaded %s (%d bytes)\n", exportOutputPath, result.ContentLength)
		fmt.Printf("Content-Type: %s\n", result.ContentType)
		return nil
	}

	return nil
}

func fetchAllFindingTemplatesForLibrary(ctx context.Context, client *services.NonSupportedServiceClient, findingLibraryID string) ([]services.FindingTemplateDto, error) {
	const pageSize = 100
	var allTemplates []services.FindingTemplateDto
	skipCount := 0

	for {
		responseModel, body, _, err := client.GetFindingTemplates(ctx, findingLibraryID, "", "", pageSize, skipCount)
		if err != nil {
			if len(body) > 0 {
				return nil, fmt.Errorf("%w: %s", err, string(body))
			}
			return nil, err
		}
		if responseModel.Result == nil {
			break
		}

		items := responseModel.Result.Items
		allTemplates = append(allTemplates, items...)
		if len(items) == 0 || len(allTemplates) >= responseModel.Result.TotalCount {
			break
		}

		skipCount += len(items)
	}

	return allTemplates, nil
}

func runTemplatesExportJSON(cmd *cobra.Command, args []string) error {
	findingLibraryID, _ := cmd.Flags().GetString("finding-library-id")
	outputPath, _ := cmd.Flags().GetString("output")

	client, err := newTemplatesServiceClient()
	if err != nil {
		return err
	}

	allTemplates, err := fetchAllFindingTemplatesForLibrary(context.Background(), client, findingLibraryID)
	if err != nil {
		return fmt.Errorf("json export failed: %w", err)
	}

	jsonBytes, err := json.MarshalIndent(allTemplates, "", "  ")
	if err != nil {
		return fmt.Errorf("encode json export: %w", err)
	}

	if strings.TrimSpace(outputPath) == "" {
		outputPath = filepath.Join(".", fmt.Sprintf("finding_templates_%s.json", findingLibraryID))
	}
	if err := os.WriteFile(outputPath, jsonBytes, 0600); err != nil {
		return fmt.Errorf("write json export file: %w", err)
	}

	fmt.Printf("Downloaded %s (%d templates)\n", outputPath, len(allTemplates))
	return nil
}

func runTemplatesLibraries(cmd *cobra.Command, args []string) error {
	filter, _ := cmd.Flags().GetString("filter")
	maxResultCount, _ := cmd.Flags().GetInt("max-result-count")
	skipCount, _ := cmd.Flags().GetInt("skip-count")

	client, err := newTemplatesServiceClient()
	if err != nil {
		return err
	}

	responseModel, body, _, err := client.GetFindingLibraries(context.Background(), filter, maxResultCount, skipCount)
	if err != nil {
		if len(body) > 0 {
			return fmt.Errorf("%w: %s", err, string(body))
		}
		return err
	}

	pretty, err := json.MarshalIndent(responseModel, "", "  ")
	if err != nil {
		fmt.Println(string(body))
		return nil
	}
	fmt.Println(string(pretty))

	if responseModel.Result != nil {
		fmt.Printf("\nReturned %d libraries (totalCount=%d)\n", len(responseModel.Result.Items), responseModel.Result.TotalCount)
	}

	return nil
}

func runTemplatesDownload(cmd *cobra.Command, args []string) error {
	fileToken, _ := cmd.Flags().GetString("file-token")
	fileName, _ := cmd.Flags().GetString("file-name")
	fileType, _ := cmd.Flags().GetString("file-type")
	outputPath, _ := cmd.Flags().GetString("output")

	client, err := newTemplatesServiceClient()
	if err != nil {
		return err
	}

	request := services.TempFileDownloadRequest{
		FileToken: fileToken,
		FileName:  fileName,
		FileType:  fileType,
	}

	result, _, err := client.DownloadTempFile(context.Background(), request)
	if err != nil {
		return err
	}

	if strings.TrimSpace(outputPath) == "" {
		outputPath = filepath.Join(".", result.FileName)
	}

	if err := os.WriteFile(outputPath, result.Body, 0600); err != nil {
		return fmt.Errorf("write downloaded file: %w", err)
	}

	fmt.Printf("Downloaded %s (%d bytes)\n", outputPath, result.ContentLength)
	fmt.Printf("Content-Type: %s\n", result.ContentType)
	return nil
}

func runTemplatesSaveTemplate(cmd *cobra.Command, args []string) error {
	dataJSON, _ := cmd.Flags().GetString("data-json")
	dataFile, _ := cmd.Flags().GetString("data-file")

	if strings.TrimSpace(dataJSON) == "" && strings.TrimSpace(dataFile) == "" {
		return fmt.Errorf("provide one payload source: --data-json or --data-file")
	}
	if strings.TrimSpace(dataJSON) != "" && strings.TrimSpace(dataFile) != "" {
		return fmt.Errorf("use only one payload source: --data-json or --data-file")
	}

	var raw []byte
	var err error
	if strings.TrimSpace(dataFile) != "" {
		raw, err = os.ReadFile(dataFile)
		if err != nil {
			return fmt.Errorf("read --data-file: %w", err)
		}
	} else {
		raw = []byte(dataJSON)
	}

	var payload interface{}
	if err := json.Unmarshal(raw, &payload); err != nil {
		return fmt.Errorf("parse JSON payload: %w", err)
	}

	client, err := newTemplatesServiceClient()
	if err != nil {
		return err
	}

	responseModel, body, _, err := client.CreateOrUpdateFindingTemplate(context.Background(), payload)
	if err != nil {
		if len(body) > 0 {
			return fmt.Errorf("%w: %s", err, string(body))
		}
		return err
	}

	pretty, err := json.MarshalIndent(responseModel, "", "  ")
	if err != nil {
		fmt.Println(string(body))
		return nil
	}
	fmt.Println(string(pretty))
	return nil
}

func runTemplatesSaveLibrary(cmd *cobra.Command, args []string) error {
	dataJSON, _ := cmd.Flags().GetString("data-json")
	dataFile, _ := cmd.Flags().GetString("data-file")

	if strings.TrimSpace(dataJSON) == "" && strings.TrimSpace(dataFile) == "" {
		return fmt.Errorf("provide one payload source: --data-json or --data-file")
	}
	if strings.TrimSpace(dataJSON) != "" && strings.TrimSpace(dataFile) != "" {
		return fmt.Errorf("use only one payload source: --data-json or --data-file")
	}

	var raw []byte
	var err error
	if strings.TrimSpace(dataFile) != "" {
		raw, err = os.ReadFile(dataFile)
		if err != nil {
			return fmt.Errorf("read --data-file: %w", err)
		}
	} else {
		raw = []byte(dataJSON)
	}

	var payload interface{}
	if err := json.Unmarshal(raw, &payload); err != nil {
		return fmt.Errorf("parse JSON payload: %w", err)
	}

	client, err := newTemplatesServiceClient()
	if err != nil {
		return err
	}

	responseModel, body, _, err := client.CreateOrUpdateFindingLibrary(context.Background(), payload)
	if err != nil {
		if len(body) > 0 {
			return fmt.Errorf("%w: %s", err, string(body))
		}
		return err
	}

	pretty, err := json.MarshalIndent(responseModel, "", "  ")
	if err != nil {
		fmt.Println(string(body))
		return nil
	}
	fmt.Println(string(pretty))
	return nil
}
