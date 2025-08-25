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

func Classroom_courses_coursework_studentsubmissions_patchHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		var requestBody models.StudentSubmission
		
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
		url := fmt.Sprintf("%s/v1/courses/%s/courseWork/%s/studentSubmissions/%s%s", cfg.BaseURL, courseId, courseWorkId, id, queryString)
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

func CreateClassroom_courses_coursework_studentsubmissions_patchTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_v1_courses_courseId_courseWork_courseWorkId_studentSubmissions_id",
		mcp.WithDescription("Updates one or more fields of a student submission. See google.classroom.v1.StudentSubmission for details of which fields may be updated and who may change them. This request must be made by the Developer Console project of the [OAuth client ID](https://support.google.com/cloud/answer/6158849) used to create the corresponding course work item. This method returns the following error codes: * `PERMISSION_DENIED` if the requesting developer project did not create the corresponding course work, if the user is not permitted to make the requested modification to the student submission, or for access errors. * `INVALID_ARGUMENT` if the request is malformed. * `NOT_FOUND` if the requested course, course work, or student submission does not exist."),
		mcp.WithString("courseId", mcp.Required(), mcp.Description("Identifier of the course. This identifier can be either the Classroom-assigned identifier or an alias.")),
		mcp.WithString("courseWorkId", mcp.Required(), mcp.Description("Identifier of the course work.")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Identifier of the student submission.")),
		mcp.WithString("updateMask", mcp.Description("Mask that identifies which fields on the student submission to update. This field is required to do an update. The update fails if invalid fields are specified. The following fields may be specified by teachers: * `draft_grade` * `assigned_grade`")),
		mcp.WithString("id", mcp.Description("Input parameter: Classroom-assigned Identifier for the student submission. This is unique among submissions for the relevant course work. Read-only.")),
		mcp.WithString("updateTime", mcp.Description("Input parameter: Last update time of this submission. This may be unset if the student has not accessed this item. Read-only.")),
		mcp.WithString("userId", mcp.Description("Input parameter: Identifier for the student that owns this submission. Read-only.")),
		mcp.WithString("courseWorkId", mcp.Description("Input parameter: Identifier for the course work this corresponds to. Read-only.")),
		mcp.WithBoolean("associatedWithDeveloper", mcp.Description("Input parameter: Whether this student submission is associated with the Developer Console project making the request. See CreateCourseWork for more details. Read-only.")),
		mcp.WithString("draftGrade", mcp.Description("Input parameter: Optional pending grade. If unset, no grade was set. This value must be non-negative. Decimal (that is, non-integer) values are allowed, but are rounded to two decimal places. This is only visible to and modifiable by course teachers.")),
		mcp.WithBoolean("late", mcp.Description("Input parameter: Whether this submission is late. Read-only.")),
		mcp.WithString("courseId", mcp.Description("Input parameter: Identifier of the course. Read-only.")),
		mcp.WithString("courseWorkType", mcp.Description("Input parameter: Type of course work this submission is for. Read-only.")),
		mcp.WithString("alternateLink", mcp.Description("Input parameter: Absolute link to the submission in the Classroom web UI. Read-only.")),
		mcp.WithString("state", mcp.Description("Input parameter: State of this submission. Read-only.")),
		mcp.WithArray("submissionHistory", mcp.Description("Input parameter: The history of the submission (includes state and grade histories). Read-only.")),
		mcp.WithObject("multipleChoiceSubmission", mcp.Description("Input parameter: Student work for a multiple-choice question.")),
		mcp.WithObject("shortAnswerSubmission", mcp.Description("Input parameter: Student work for a short answer question.")),
		mcp.WithString("creationTime", mcp.Description("Input parameter: Creation time of this submission. This may be unset if the student has not accessed this item. Read-only.")),
		mcp.WithString("assignedGrade", mcp.Description("Input parameter: Optional grade. If unset, no grade was set. This value must be non-negative. Decimal (that is, non-integer) values are allowed, but are rounded to two decimal places. This may be modified only by course teachers.")),
		mcp.WithObject("assignmentSubmission", mcp.Description("Input parameter: Student work for an assignment.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Classroom_courses_coursework_studentsubmissions_patchHandler(cfg),
	}
}
