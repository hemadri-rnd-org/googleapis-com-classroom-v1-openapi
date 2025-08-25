package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/google-classroom-api/mcp-server/config"
	"github.com/google-classroom-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Classroom_userprofiles_guardians_listHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		studentIdVal, ok := args["studentId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: studentId"), nil
		}
		studentId, ok := studentIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: studentId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["invitedEmailAddress"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("invitedEmailAddress=%v", val))
		}
		if val, ok := args["pageSize"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("pageSize=%v", val))
		}
		if val, ok := args["pageToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("pageToken=%v", val))
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
		url := fmt.Sprintf("%s/v1/userProfiles/%s/guardians%s", cfg.BaseURL, studentId, queryString)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.ListGuardiansResponse
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

func CreateClassroom_userprofiles_guardians_listTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_userProfiles_studentId_guardians",
		mcp.WithDescription("Returns a list of guardians that the requesting user is permitted to view, restricted to those that match the request. To list guardians for any student that the requesting user may view guardians for, use the literal character `-` for the student ID. This method returns the following error codes: * `PERMISSION_DENIED` if a `student_id` is specified, and the requesting user is not permitted to view guardian information for that student, if `"-"` is specified as the `student_id` and the user is not a domain administrator, if guardians are not enabled for the domain in question, if the `invited_email_address` filter is set by a user who is not a domain administrator, or for other access errors. * `INVALID_ARGUMENT` if a `student_id` is specified, but its format cannot be recognized (it is not an email address, nor a `student_id` from the API, nor the literal string `me`). May also be returned if an invalid `page_token` is provided. * `NOT_FOUND` if a `student_id` is specified, and its format can be recognized, but Classroom has no record of that student."),
		mcp.WithString("studentId", mcp.Required(), mcp.Description("Filter results by the student who the guardian is linked to. The identifier can be one of the following: * the numeric identifier for the user * the email address of the user * the string literal `\"me\"`, indicating the requesting user * the string literal `\"-\"`, indicating that results should be returned for all students that the requesting user has access to view.")),
		mcp.WithString("invitedEmailAddress", mcp.Description("Filter results by the email address that the original invitation was sent to, resulting in this guardian link. This filter can only be used by domain administrators.")),
		mcp.WithNumber("pageSize", mcp.Description("Maximum number of items to return. Zero or unspecified indicates that the server may assign a maximum. The server may return fewer than the specified number of results.")),
		mcp.WithString("pageToken", mcp.Description("nextPageToken value returned from a previous list call, indicating that the subsequent page of results should be returned. The list request must be otherwise identical to the one that resulted in this token.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Classroom_userprofiles_guardians_listHandler(cfg),
	}
}
