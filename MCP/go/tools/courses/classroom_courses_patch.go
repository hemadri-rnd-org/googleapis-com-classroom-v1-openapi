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

func Classroom_courses_patchHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
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
		var requestBody models.Course
		
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
		url := fmt.Sprintf("%s/v1/courses/%s%s", cfg.BaseURL, id, queryString)
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
		var result models.Course
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

func CreateClassroom_courses_patchTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_v1_courses_id",
		mcp.WithDescription("Updates one or more fields in a course. This method returns the following error codes: * `PERMISSION_DENIED` if the requesting user is not permitted to modify the requested course or for access errors. * `NOT_FOUND` if no course exists with the requested ID. * `INVALID_ARGUMENT` if invalid fields are specified in the update mask or if no update mask is supplied. * `FAILED_PRECONDITION` for the following request errors: * CourseNotModifiable * InactiveCourseOwner * IneligibleOwner"),
		mcp.WithString("id", mcp.Required(), mcp.Description("Identifier of the course to update. This identifier can be either the Classroom-assigned identifier or an alias.")),
		mcp.WithString("updateMask", mcp.Description("Mask that identifies which fields on the course to update. This field is required to do an update. The update will fail if invalid fields are specified. The following fields are valid: * `name` * `section` * `descriptionHeading` * `description` * `room` * `courseState` * `ownerId` Note: patches to ownerId are treated as being effective immediately, but in practice it may take some time for the ownership transfer of all affected resources to complete. When set in a query parameter, this field should be specified as `updateMask=,,...`")),
		mcp.WithString("descriptionHeading", mcp.Description("Input parameter: Optional heading for the description. For example, \"Welcome to 10th Grade Biology.\" If set, this field must be a valid UTF-8 string and no longer than 3600 characters.")),
		mcp.WithString("calendarId", mcp.Description("Input parameter: The Calendar ID for a calendar that all course members can see, to which Classroom adds events for course work and announcements in the course. The Calendar for a course is created asynchronously when the course is set to `CourseState.ACTIVE` for the first time (at creation time or when it is updated to `ACTIVE` through the UI or the API). The Calendar ID will not be populated until the creation process is completed. Read-only.")),
		mcp.WithString("teacherGroupEmail", mcp.Description("Input parameter: The email address of a Google group containing all teachers of the course. This group does not accept email and can only be used for permissions. Read-only.")),
		mcp.WithArray("courseMaterialSets", mcp.Description("Input parameter: Sets of materials that appear on the \"about\" page of this course. Read-only.")),
		mcp.WithString("ownerId", mcp.Description("Input parameter: The identifier of the owner of a course. When specified as a parameter of a create course request, this field is required. The identifier can be one of the following: * the numeric identifier for the user * the email address of the user * the string literal `\"me\"`, indicating the requesting user This must be set in a create request. Admins can also specify this field in a patch course request to transfer ownership. In other contexts, it is read-only.")),
		mcp.WithString("courseGroupEmail", mcp.Description("Input parameter: The email address of a Google group containing all members of the course. This group does not accept email and can only be used for permissions. Read-only.")),
		mcp.WithString("section", mcp.Description("Input parameter: Section of the course. For example, \"Period 2\". If set, this field must be a valid UTF-8 string and no longer than 2800 characters.")),
		mcp.WithString("enrollmentCode", mcp.Description("Input parameter: Enrollment code to use when joining this course. Specifying this field in a course update mask results in an error. Read-only.")),
		mcp.WithString("name", mcp.Description("Input parameter: Name of the course. For example, \"10th Grade Biology\". The name is required. It must be between 1 and 750 characters and a valid UTF-8 string.")),
		mcp.WithString("room", mcp.Description("Input parameter: Optional room location. For example, \"301\". If set, this field must be a valid UTF-8 string and no longer than 650 characters.")),
		mcp.WithString("id", mcp.Description("Input parameter: Identifier for this course assigned by Classroom. When creating a course, you may optionally set this identifier to an alias string in the request to create a corresponding alias. The `id` is still assigned by Classroom and cannot be updated after the course is created. Specifying this field in a course update mask results in an error.")),
		mcp.WithString("description", mcp.Description("Input parameter: Optional description. For example, \"We'll be learning about the structure of living creatures from a combination of textbooks, guest lectures, and lab work. Expect to be excited!\" If set, this field must be a valid UTF-8 string and no longer than 30,000 characters.")),
		mcp.WithBoolean("guardiansEnabled", mcp.Description("Input parameter: Whether or not guardian notifications are enabled for this course. Read-only.")),
		mcp.WithString("updateTime", mcp.Description("Input parameter: Time of the most recent update to this course. Specifying this field in a course update mask results in an error. Read-only.")),
		mcp.WithString("courseState", mcp.Description("Input parameter: State of the course. If unspecified, the default state is `PROVISIONED`.")),
		mcp.WithString("alternateLink", mcp.Description("Input parameter: Absolute link to this course in the Classroom web UI. Read-only.")),
		mcp.WithObject("teacherFolder", mcp.Description("Input parameter: Representation of a Google Drive folder.")),
		mcp.WithObject("gradebookSettings", mcp.Description("Input parameter: The gradebook settings for a course. See the [help center article](https://support.google.com/edu/classroom/answer/9184995) for details.")),
		mcp.WithString("creationTime", mcp.Description("Input parameter: Creation time of the course. Specifying this field in a course update mask results in an error. Read-only.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Classroom_courses_patchHandler(cfg),
	}
}
