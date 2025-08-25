package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/google-classroom-api/mcp-server/config"
	"github.com/google-classroom-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Classroom_courses_coursework_studentsubmissions_modifyattachmentsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		courseIdVal, ok := args["courseId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: courseId"), nil
		}
		courseId, ok := courseIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: courseId"), nil
		}
		courseWorkIdVal, ok := args["courseWorkId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: courseWorkId"), nil
		}
		courseWorkId, ok := courseWorkIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: courseWorkId"), nil
		}
		idVal, ok := args["id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: id"), nil
		}
		id, ok := idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: id"), nil
		}
		queryParams := make([]string, 0)
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("access_token=%s", cfg.BearerToken))
		}
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("key=%s", cfg.APIKey))
		}
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("oauth_token=%s", cfg.BearerToken))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.ModifyAttachmentsRequest
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/v1/courses/%s/courseWork/%s/studentSubmissions/%s:modifyAttachments%s", cfg.BaseURL, courseId, courseWorkId, id, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		// API key already added to query string
		// API key already added to query string
		// API key already added to query string
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.StudentSubmission
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateClassroom_courses_coursework_studentsubmissions_modifyattachmentsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_v1_courses_courseId_courseWork_courseWorkId_studentSubmissions_id_modifyAttachments",
		mcp.WithDescription("Modifies attachments of student submission. Attachments may only be added to student submissions belonging to course work objects with a `workType` of `ASSIGNMENT`. This request must be made by the Developer Console project of the [OAuth client ID](https://support.google.com/cloud/answer/6158849) used to create the corresponding course work item. This method returns the following error codes: * `PERMISSION_DENIED` if the requesting user is not permitted to access the requested course or course work, if the user is not permitted to modify attachments on the requested student submission, or for access errors. * `INVALID_ARGUMENT` if the request is malformed. * `NOT_FOUND` if the requested course, course work, or student submission does not exist."),
		mcp.WithString("courseId", mcp.Required(), mcp.Description("Identifier of the course. This identifier can be either the Classroom-assigned identifier or an alias.")),
		mcp.WithString("courseWorkId", mcp.Required(), mcp.Description("Identifier of the course work.")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Identifier of the student submission.")),
		mcp.WithArray("addAttachments", mcp.Description("Input parameter: Attachments to add. A student submission may not have more than 20 attachments. Form attachments are not supported.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Classroom_courses_coursework_studentsubmissions_modifyattachmentsHandler(cfg),
	}
}
