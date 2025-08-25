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

func Classroom_courses_topics_patchHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		idVal, ok := args["id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: id"), nil
		}
		id, ok := idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: id"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["updateMask"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("updateMask=%v", val))
		}
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
		var requestBody models.Topic
		
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
		url := fmt.Sprintf("%s/v1/courses/%s/topics/%s%s", cfg.BaseURL, courseId, id, queryString)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
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
		var result models.Topic
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

func CreateClassroom_courses_topics_patchTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_v1_courses_courseId_topics_id",
		mcp.WithDescription("Updates one or more fields of a topic. This method returns the following error codes: * `PERMISSION_DENIED` if the requesting developer project did not create the corresponding topic or for access errors. * `INVALID_ARGUMENT` if the request is malformed. * `NOT_FOUND` if the requested course or topic does not exist"),
		mcp.WithString("courseId", mcp.Required(), mcp.Description("Identifier of the course. This identifier can be either the Classroom-assigned identifier or an alias.")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Identifier of the topic.")),
		mcp.WithString("updateMask", mcp.Description("Mask that identifies which fields on the topic to update. This field is required to do an update. The update fails if invalid fields are specified. If a field supports empty values, it can be cleared by specifying it in the update mask and not in the Topic object. If a field that does not support empty values is included in the update mask and not set in the Topic object, an `INVALID_ARGUMENT` error is returned. The following fields may be specified: * `name`")),
		mcp.WithString("updateTime", mcp.Description("Input parameter: The time the topic was last updated by the system. Read-only.")),
		mcp.WithString("courseId", mcp.Description("Input parameter: Identifier of the course. Read-only.")),
		mcp.WithString("name", mcp.Description("Input parameter: The name of the topic, generated by the user. Leading and trailing whitespaces, if any, are trimmed. Also, multiple consecutive whitespaces are collapsed into one inside the name. The result must be a non-empty string. Topic names are case sensitive, and must be no longer than 100 characters.")),
		mcp.WithString("topicId", mcp.Description("Input parameter: Unique identifier for the topic. Read-only.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Classroom_courses_topics_patchHandler(cfg),
	}
}
