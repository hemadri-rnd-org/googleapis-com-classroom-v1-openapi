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

func Classroom_courses_coursework_patchHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		var requestBody models.CourseWork
		
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
		url := fmt.Sprintf("%s/v1/courses/%s/courseWork/%s%s", cfg.BaseURL, courseId, id, queryString)
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
		var result models.CourseWork
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

func CreateClassroom_courses_coursework_patchTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_v1_courses_courseId_courseWork_id",
		mcp.WithDescription("Updates one or more fields of a course work. See google.classroom.v1.CourseWork for details of which fields may be updated and who may change them. This request must be made by the Developer Console project of the [OAuth client ID](https://support.google.com/cloud/answer/6158849) used to create the corresponding course work item. This method returns the following error codes: * `PERMISSION_DENIED` if the requesting developer project did not create the corresponding course work, if the user is not permitted to make the requested modification to the student submission, or for access errors. * `INVALID_ARGUMENT` if the request is malformed. * `FAILED_PRECONDITION` if the requested course work has already been deleted. * `NOT_FOUND` if the requested course, course work, or student submission does not exist."),
		mcp.WithString("courseId", mcp.Required(), mcp.Description("Identifier of the course. This identifier can be either the Classroom-assigned identifier or an alias.")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Identifier of the course work.")),
		mcp.WithString("updateMask", mcp.Description("Mask that identifies which fields on the course work to update. This field is required to do an update. The update fails if invalid fields are specified. If a field supports empty values, it can be cleared by specifying it in the update mask and not in the `CourseWork` object. If a field that does not support empty values is included in the update mask and not set in the `CourseWork` object, an `INVALID_ARGUMENT` error is returned. The following fields may be specified by teachers: * `title` * `description` * `state` * `due_date` * `due_time` * `max_points` * `scheduled_time` * `submission_modification_mode` * `topic_id`")),
		mcp.WithString("assigneeMode", mcp.Description("Input parameter: Assignee mode of the coursework. If unspecified, the default value is `ALL_STUDENTS`.")),
		mcp.WithObject("gradeCategory", mcp.Description("Input parameter: Details for a grade category in a course. Coursework may have zero or one grade category, and the category may be used in computing the overall grade. See the [help center article](https://support.google.com/edu/classroom/answer/9184995) for details.")),
		mcp.WithBoolean("associatedWithDeveloper", mcp.Description("Input parameter: Whether this course work item is associated with the Developer Console project making the request. See CreateCourseWork for more details. Read-only.")),
		mcp.WithObject("assignment", mcp.Description("Input parameter: Additional details for assignments.")),
		mcp.WithString("scheduledTime", mcp.Description("Input parameter: Optional timestamp when this course work is scheduled to be published.")),
		mcp.WithObject("individualStudentsOptions", mcp.Description("Input parameter: Assignee details about a coursework/announcement. This field is set if and only if `assigneeMode` is `INDIVIDUAL_STUDENTS`.")),
		mcp.WithArray("materials", mcp.Description("Input parameter: Additional materials. CourseWork must have no more than 20 material items.")),
		mcp.WithObject("dueDate", mcp.Description("Input parameter: Represents a whole or partial calendar date, such as a birthday. The time of day and time zone are either specified elsewhere or are insignificant. The date is relative to the Gregorian Calendar. This can represent one of the following: * A full date, with non-zero year, month, and day values. * A month and day, with a zero year (for example, an anniversary). * A year on its own, with a zero month and a zero day. * A year and month, with a zero day (for example, a credit card expiration date). Related types: * google.type.TimeOfDay * google.type.DateTime * google.protobuf.Timestamp")),
		mcp.WithString("id", mcp.Description("Input parameter: Classroom-assigned identifier of this course work, unique per course. Read-only.")),
		mcp.WithString("state", mcp.Description("Input parameter: Status of this course work. If unspecified, the default state is `DRAFT`.")),
		mcp.WithString("workType", mcp.Description("Input parameter: Type of this course work. The type is set when the course work is created and cannot be changed.")),
		mcp.WithObject("dueTime", mcp.Description("Input parameter: Represents a time of day. The date and time zone are either not significant or are specified elsewhere. An API may choose to allow leap seconds. Related types are google.type.Date and `google.protobuf.Timestamp`.")),
		mcp.WithString("creatorUserId", mcp.Description("Input parameter: Identifier for the user that created the coursework. Read-only.")),
		mcp.WithObject("multipleChoiceQuestion", mcp.Description("Input parameter: Additional details for multiple-choice questions.")),
		mcp.WithString("creationTime", mcp.Description("Input parameter: Timestamp when this course work was created. Read-only.")),
		mcp.WithString("description", mcp.Description("Input parameter: Optional description of this course work. If set, the description must be a valid UTF-8 string containing no more than 30,000 characters.")),
		mcp.WithString("maxPoints", mcp.Description("Input parameter: Maximum grade for this course work. If zero or unspecified, this assignment is considered ungraded. This must be a non-negative integer value.")),
		mcp.WithString("courseId", mcp.Description("Input parameter: Identifier of the course. Read-only.")),
		mcp.WithString("title", mcp.Description("Input parameter: Title of this course work. The title must be a valid UTF-8 string containing between 1 and 3000 characters.")),
		mcp.WithString("updateTime", mcp.Description("Input parameter: Timestamp of the most recent change to this course work. Read-only.")),
		mcp.WithString("submissionModificationMode", mcp.Description("Input parameter: Setting to determine when students are allowed to modify submissions. If unspecified, the default value is `MODIFIABLE_UNTIL_TURNED_IN`.")),
		mcp.WithString("topicId", mcp.Description("Input parameter: Identifier for the topic that this coursework is associated with. Must match an existing topic in the course.")),
		mcp.WithString("alternateLink", mcp.Description("Input parameter: Absolute link to this course work in the Classroom web UI. This is only populated if `state` is `PUBLISHED`. Read-only.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Classroom_courses_coursework_patchHandler(cfg),
	}
}
