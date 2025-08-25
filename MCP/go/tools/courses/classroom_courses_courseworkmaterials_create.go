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

func Classroom_courses_courseworkmaterials_createHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		var requestBody models.CourseWorkMaterial
		
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
		url := fmt.Sprintf("%s/v1/courses/%s/courseWorkMaterials%s", cfg.BaseURL, courseId, queryString)
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
		var result models.CourseWorkMaterial
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

func CreateClassroom_courses_courseworkmaterials_createTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_v1_courses_courseId_courseWorkMaterials",
		mcp.WithDescription("Creates a course work material. This method returns the following error codes: * `PERMISSION_DENIED` if the requesting user is not permitted to access the requested course, create course work material in the requested course, share a Drive attachment, or for access errors. * `INVALID_ARGUMENT` if the request is malformed or if more than 20 * materials are provided. * `NOT_FOUND` if the requested course does not exist. * `FAILED_PRECONDITION` for the following request error: * AttachmentNotVisible"),
		mcp.WithString("courseId", mcp.Required(), mcp.Description("Identifier of the course. This identifier can be either the Classroom-assigned identifier or an alias.")),
		mcp.WithArray("materials", mcp.Description("Input parameter: Additional materials. A course work material must have no more than 20 material items.")),
		mcp.WithString("title", mcp.Description("Input parameter: Title of this course work material. The title must be a valid UTF-8 string containing between 1 and 3000 characters.")),
		mcp.WithString("scheduledTime", mcp.Description("Input parameter: Optional timestamp when this course work material is scheduled to be published.")),
		mcp.WithString("topicId", mcp.Description("Input parameter: Identifier for the topic that this course work material is associated with. Must match an existing topic in the course.")),
		mcp.WithString("assigneeMode", mcp.Description("Input parameter: Assignee mode of the course work material. If unspecified, the default value is `ALL_STUDENTS`.")),
		mcp.WithString("creationTime", mcp.Description("Input parameter: Timestamp when this course work material was created. Read-only.")),
		mcp.WithObject("individualStudentsOptions", mcp.Description("Input parameter: Assignee details about a coursework/announcement. This field is set if and only if `assigneeMode` is `INDIVIDUAL_STUDENTS`.")),
		mcp.WithString("alternateLink", mcp.Description("Input parameter: Absolute link to this course work material in the Classroom web UI. This is only populated if `state` is `PUBLISHED`. Read-only.")),
		mcp.WithString("creatorUserId", mcp.Description("Input parameter: Identifier for the user that created the course work material. Read-only.")),
		mcp.WithString("id", mcp.Description("Input parameter: Classroom-assigned identifier of this course work material, unique per course. Read-only.")),
		mcp.WithString("state", mcp.Description("Input parameter: Status of this course work material. If unspecified, the default state is `DRAFT`.")),
		mcp.WithString("updateTime", mcp.Description("Input parameter: Timestamp of the most recent change to this course work material. Read-only.")),
		mcp.WithString("courseId", mcp.Description("Input parameter: Identifier of the course. Read-only.")),
		mcp.WithString("description", mcp.Description("Input parameter: Optional description of this course work material. The text must be a valid UTF-8 string containing no more than 30,000 characters.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Classroom_courses_courseworkmaterials_createHandler(cfg),
	}
}
