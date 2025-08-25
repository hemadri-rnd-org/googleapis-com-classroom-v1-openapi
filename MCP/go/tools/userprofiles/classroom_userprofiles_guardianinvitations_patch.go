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

func Classroom_userprofiles_guardianinvitations_patchHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		invitationIdVal, ok := args["invitationId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: invitationId"), nil
		}
		invitationId, ok := invitationIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: invitationId"), nil
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
		var requestBody models.GuardianInvitation
		
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
		url := fmt.Sprintf("%s/v1/userProfiles/%s/guardianInvitations/%s%s", cfg.BaseURL, studentId, invitationId, queryString)
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
		var result models.GuardianInvitation
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

func CreateClassroom_userprofiles_guardianinvitations_patchTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_v1_userProfiles_studentId_guardianInvitations_invitationId",
		mcp.WithDescription("Modifies a guardian invitation. Currently, the only valid modification is to change the `state` from `PENDING` to `COMPLETE`. This has the effect of withdrawing the invitation. This method returns the following error codes: * `PERMISSION_DENIED` if the current user does not have permission to manage guardians, if guardians are not enabled for the domain in question or for other access errors. * `FAILED_PRECONDITION` if the guardian link is not in the `PENDING` state. * `INVALID_ARGUMENT` if the format of the student ID provided cannot be recognized (it is not an email address, nor a `user_id` from this API), or if the passed `GuardianInvitation` has a `state` other than `COMPLETE`, or if it modifies fields other than `state`. * `NOT_FOUND` if the student ID provided is a valid student ID, but Classroom has no record of that student, or if the `id` field does not refer to a guardian invitation known to Classroom."),
		mcp.WithString("studentId", mcp.Required(), mcp.Description("The ID of the student whose guardian invitation is to be modified.")),
		mcp.WithString("invitationId", mcp.Required(), mcp.Description("The `id` field of the `GuardianInvitation` to be modified.")),
		mcp.WithString("updateMask", mcp.Description("Mask that identifies which fields on the course to update. This field is required to do an update. The update fails if invalid fields are specified. The following fields are valid: * `state` When set in a query parameter, this field should be specified as `updateMask=,,...`")),
		mcp.WithString("state", mcp.Description("Input parameter: The state that this invitation is in.")),
		mcp.WithString("studentId", mcp.Description("Input parameter: ID of the student (in standard format)")),
		mcp.WithString("creationTime", mcp.Description("Input parameter: The time that this invitation was created. Read-only.")),
		mcp.WithString("invitationId", mcp.Description("Input parameter: Unique identifier for this invitation. Read-only.")),
		mcp.WithString("invitedEmailAddress", mcp.Description("Input parameter: Email address that the invitation was sent to. This field is only visible to domain administrators.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Classroom_userprofiles_guardianinvitations_patchHandler(cfg),
	}
}
