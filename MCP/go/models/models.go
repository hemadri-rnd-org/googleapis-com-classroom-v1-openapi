package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// CourseWorkChangesInfo represents the CourseWorkChangesInfo schema from the OpenAPI specification
type CourseWorkChangesInfo struct {
	Courseid string `json:"courseId,omitempty"` // The `course_id` of the course to subscribe to work changes for.
}

// Course represents the Course schema from the OpenAPI specification
type Course struct {
	Coursestate string `json:"courseState,omitempty"` // State of the course. If unspecified, the default state is `PROVISIONED`.
	Alternatelink string `json:"alternateLink,omitempty"` // Absolute link to this course in the Classroom web UI. Read-only.
	Teacherfolder DriveFolder `json:"teacherFolder,omitempty"` // Representation of a Google Drive folder.
	Gradebooksettings GradebookSettings `json:"gradebookSettings,omitempty"` // The gradebook settings for a course. See the [help center article](https://support.google.com/edu/classroom/answer/9184995) for details.
	Creationtime string `json:"creationTime,omitempty"` // Creation time of the course. Specifying this field in a course update mask results in an error. Read-only.
	Descriptionheading string `json:"descriptionHeading,omitempty"` // Optional heading for the description. For example, "Welcome to 10th Grade Biology." If set, this field must be a valid UTF-8 string and no longer than 3600 characters.
	Calendarid string `json:"calendarId,omitempty"` // The Calendar ID for a calendar that all course members can see, to which Classroom adds events for course work and announcements in the course. The Calendar for a course is created asynchronously when the course is set to `CourseState.ACTIVE` for the first time (at creation time or when it is updated to `ACTIVE` through the UI or the API). The Calendar ID will not be populated until the creation process is completed. Read-only.
	Teachergroupemail string `json:"teacherGroupEmail,omitempty"` // The email address of a Google group containing all teachers of the course. This group does not accept email and can only be used for permissions. Read-only.
	Coursematerialsets []CourseMaterialSet `json:"courseMaterialSets,omitempty"` // Sets of materials that appear on the "about" page of this course. Read-only.
	Ownerid string `json:"ownerId,omitempty"` // The identifier of the owner of a course. When specified as a parameter of a create course request, this field is required. The identifier can be one of the following: * the numeric identifier for the user * the email address of the user * the string literal `"me"`, indicating the requesting user This must be set in a create request. Admins can also specify this field in a patch course request to transfer ownership. In other contexts, it is read-only.
	Coursegroupemail string `json:"courseGroupEmail,omitempty"` // The email address of a Google group containing all members of the course. This group does not accept email and can only be used for permissions. Read-only.
	Section string `json:"section,omitempty"` // Section of the course. For example, "Period 2". If set, this field must be a valid UTF-8 string and no longer than 2800 characters.
	Enrollmentcode string `json:"enrollmentCode,omitempty"` // Enrollment code to use when joining this course. Specifying this field in a course update mask results in an error. Read-only.
	Name string `json:"name,omitempty"` // Name of the course. For example, "10th Grade Biology". The name is required. It must be between 1 and 750 characters and a valid UTF-8 string.
	Room string `json:"room,omitempty"` // Optional room location. For example, "301". If set, this field must be a valid UTF-8 string and no longer than 650 characters.
	Id string `json:"id,omitempty"` // Identifier for this course assigned by Classroom. When creating a course, you may optionally set this identifier to an alias string in the request to create a corresponding alias. The `id` is still assigned by Classroom and cannot be updated after the course is created. Specifying this field in a course update mask results in an error.
	Description string `json:"description,omitempty"` // Optional description. For example, "We'll be learning about the structure of living creatures from a combination of textbooks, guest lectures, and lab work. Expect to be excited!" If set, this field must be a valid UTF-8 string and no longer than 30,000 characters.
	Guardiansenabled bool `json:"guardiansEnabled,omitempty"` // Whether or not guardian notifications are enabled for this course. Read-only.
	Updatetime string `json:"updateTime,omitempty"` // Time of the most recent update to this course. Specifying this field in a course update mask results in an error. Read-only.
}

// ShortAnswerSubmission represents the ShortAnswerSubmission schema from the OpenAPI specification
type ShortAnswerSubmission struct {
	Answer string `json:"answer,omitempty"` // Student response to a short-answer question.
}

// SharedDriveFile represents the SharedDriveFile schema from the OpenAPI specification
type SharedDriveFile struct {
	Drivefile DriveFile `json:"driveFile,omitempty"` // Representation of a Google Drive file.
	Sharemode string `json:"shareMode,omitempty"` // Mechanism by which students access the Drive item.
}

// ListStudentsResponse represents the ListStudentsResponse schema from the OpenAPI specification
type ListStudentsResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token identifying the next page of results to return. If empty, no further results are available.
	Students []Student `json:"students,omitempty"` // Students who match the list request.
}

// CloudPubsubTopic represents the CloudPubsubTopic schema from the OpenAPI specification
type CloudPubsubTopic struct {
	Topicname string `json:"topicName,omitempty"` // The `name` field of a Cloud Pub/Sub [Topic](https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.topics#Topic).
}

// Registration represents the Registration schema from the OpenAPI specification
type Registration struct {
	Cloudpubsubtopic CloudPubsubTopic `json:"cloudPubsubTopic,omitempty"` // A reference to a Cloud Pub/Sub topic. To register for notifications, the owner of the topic must grant `classroom-notifications@system.gserviceaccount.com` the `projects.topics.publish` permission.
	Expirytime string `json:"expiryTime,omitempty"` // The time until which the `Registration` is effective. This is a read-only field assigned by the server.
	Feed Feed `json:"feed,omitempty"` // A class of notifications that an application can register to receive. For example: "all roster changes for a domain".
	Registrationid string `json:"registrationId,omitempty"` // A server-generated unique identifier for this `Registration`. Read-only.
}

// Assignment represents the Assignment schema from the OpenAPI specification
type Assignment struct {
	Studentworkfolder DriveFolder `json:"studentWorkFolder,omitempty"` // Representation of a Google Drive folder.
}

// ListTeachersResponse represents the ListTeachersResponse schema from the OpenAPI specification
type ListTeachersResponse struct {
	Teachers []Teacher `json:"teachers,omitempty"` // Teachers who match the list request.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token identifying the next page of results to return. If empty, no further results are available.
}

// ReturnStudentSubmissionRequest represents the ReturnStudentSubmissionRequest schema from the OpenAPI specification
type ReturnStudentSubmissionRequest struct {
}

// Date represents the Date schema from the OpenAPI specification
type Date struct {
	Year int `json:"year,omitempty"` // Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.
	Day int `json:"day,omitempty"` // Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.
	Month int `json:"month,omitempty"` // Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.
}

// ListGuardianInvitationsResponse represents the ListGuardianInvitationsResponse schema from the OpenAPI specification
type ListGuardianInvitationsResponse struct {
	Guardianinvitations []GuardianInvitation `json:"guardianInvitations,omitempty"` // Guardian invitations that matched the list request.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token identifying the next page of results to return. If empty, no further results are available.
}

// ModifyCourseWorkAssigneesRequest represents the ModifyCourseWorkAssigneesRequest schema from the OpenAPI specification
type ModifyCourseWorkAssigneesRequest struct {
	Modifyindividualstudentsoptions ModifyIndividualStudentsOptions `json:"modifyIndividualStudentsOptions,omitempty"` // Contains fields to add or remove students from a course work or announcement where the `assigneeMode` is set to `INDIVIDUAL_STUDENTS`.
	Assigneemode string `json:"assigneeMode,omitempty"` // Mode of the coursework describing whether it will be assigned to all students or specified individual students.
}

// SubmissionHistory represents the SubmissionHistory schema from the OpenAPI specification
type SubmissionHistory struct {
	Gradehistory GradeHistory `json:"gradeHistory,omitempty"` // The history of each grade on this submission.
	Statehistory StateHistory `json:"stateHistory,omitempty"` // The history of each state this submission has been in.
}

// Attachment represents the Attachment schema from the OpenAPI specification
type Attachment struct {
	Drivefile DriveFile `json:"driveFile,omitempty"` // Representation of a Google Drive file.
	Form Form `json:"form,omitempty"` // Google Forms item.
	Link Link `json:"link,omitempty"` // URL item.
	Youtubevideo YouTubeVideo `json:"youTubeVideo,omitempty"` // YouTube video item.
}

// GradeHistory represents the GradeHistory schema from the OpenAPI specification
type GradeHistory struct {
	Gradechangetype string `json:"gradeChangeType,omitempty"` // The type of grade change at this time in the submission grade history.
	Gradetimestamp string `json:"gradeTimestamp,omitempty"` // When the grade of the submission was changed.
	Maxpoints float64 `json:"maxPoints,omitempty"` // The denominator of the grade at this time in the submission grade history.
	Pointsearned float64 `json:"pointsEarned,omitempty"` // The numerator of the grade at this time in the submission grade history.
	Actoruserid string `json:"actorUserId,omitempty"` // The teacher who made the grade change.
}

// ListCoursesResponse represents the ListCoursesResponse schema from the OpenAPI specification
type ListCoursesResponse struct {
	Courses []Course `json:"courses,omitempty"` // Courses that match the list request.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token identifying the next page of results to return. If empty, no further results are available.
}

// CourseRosterChangesInfo represents the CourseRosterChangesInfo schema from the OpenAPI specification
type CourseRosterChangesInfo struct {
	Courseid string `json:"courseId,omitempty"` // The `course_id` of the course to subscribe to roster changes for.
}

// Name represents the Name schema from the OpenAPI specification
type Name struct {
	Givenname string `json:"givenName,omitempty"` // The user's first name. Read-only.
	Familyname string `json:"familyName,omitempty"` // The user's last name. Read-only.
	Fullname string `json:"fullName,omitempty"` // The user's full name formed by concatenating the first and last name values. Read-only.
}

// Topic represents the Topic schema from the OpenAPI specification
type Topic struct {
	Courseid string `json:"courseId,omitempty"` // Identifier of the course. Read-only.
	Name string `json:"name,omitempty"` // The name of the topic, generated by the user. Leading and trailing whitespaces, if any, are trimmed. Also, multiple consecutive whitespaces are collapsed into one inside the name. The result must be a non-empty string. Topic names are case sensitive, and must be no longer than 100 characters.
	Topicid string `json:"topicId,omitempty"` // Unique identifier for the topic. Read-only.
	Updatetime string `json:"updateTime,omitempty"` // The time the topic was last updated by the system. Read-only.
}

// AssignmentSubmission represents the AssignmentSubmission schema from the OpenAPI specification
type AssignmentSubmission struct {
	Attachments []Attachment `json:"attachments,omitempty"` // Attachments added by the student. Drive files that correspond to materials with a share mode of STUDENT_COPY may not exist yet if the student has not accessed the assignment in Classroom. Some attachment metadata is only populated if the requesting user has permission to access it. Identifier and alternate_link fields are always available, but others (for example, title) may not be.
}

// DriveFolder represents the DriveFolder schema from the OpenAPI specification
type DriveFolder struct {
	Alternatelink string `json:"alternateLink,omitempty"` // URL that can be used to access the Drive folder. Read-only.
	Id string `json:"id,omitempty"` // Drive API resource ID.
	Title string `json:"title,omitempty"` // Title of the Drive folder. Read-only.
}

// CourseMaterialSet represents the CourseMaterialSet schema from the OpenAPI specification
type CourseMaterialSet struct {
	Title string `json:"title,omitempty"` // Title for this set.
	Materials []CourseMaterial `json:"materials,omitempty"` // Materials attached to this set.
}

// ListCourseWorkMaterialResponse represents the ListCourseWorkMaterialResponse schema from the OpenAPI specification
type ListCourseWorkMaterialResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token identifying the next page of results to return. If empty, no further results are available.
	Courseworkmaterial []CourseWorkMaterial `json:"courseWorkMaterial,omitempty"` // Course work material items that match the request.
}

// CourseAlias represents the CourseAlias schema from the OpenAPI specification
type CourseAlias struct {
	Alias string `json:"alias,omitempty"` // Alias string. The format of the string indicates the desired alias scoping. * `d:` indicates a domain-scoped alias. Example: `d:math_101` * `p:` indicates a project-scoped alias. Example: `p:abc123` This field has a maximum length of 256 characters.
}

// GuardianInvitation represents the GuardianInvitation schema from the OpenAPI specification
type GuardianInvitation struct {
	State string `json:"state,omitempty"` // The state that this invitation is in.
	Studentid string `json:"studentId,omitempty"` // ID of the student (in standard format)
	Creationtime string `json:"creationTime,omitempty"` // The time that this invitation was created. Read-only.
	Invitationid string `json:"invitationId,omitempty"` // Unique identifier for this invitation. Read-only.
	Invitedemailaddress string `json:"invitedEmailAddress,omitempty"` // Email address that the invitation was sent to. This field is only visible to domain administrators.
}

// CourseMaterial represents the CourseMaterial schema from the OpenAPI specification
type CourseMaterial struct {
	Link Link `json:"link,omitempty"` // URL item.
	Youtubevideo YouTubeVideo `json:"youTubeVideo,omitempty"` // YouTube video item.
	Drivefile DriveFile `json:"driveFile,omitempty"` // Representation of a Google Drive file.
	Form Form `json:"form,omitempty"` // Google Forms item.
}

// CourseWorkMaterial represents the CourseWorkMaterial schema from the OpenAPI specification
type CourseWorkMaterial struct {
	Creatoruserid string `json:"creatorUserId,omitempty"` // Identifier for the user that created the course work material. Read-only.
	Id string `json:"id,omitempty"` // Classroom-assigned identifier of this course work material, unique per course. Read-only.
	State string `json:"state,omitempty"` // Status of this course work material. If unspecified, the default state is `DRAFT`.
	Updatetime string `json:"updateTime,omitempty"` // Timestamp of the most recent change to this course work material. Read-only.
	Courseid string `json:"courseId,omitempty"` // Identifier of the course. Read-only.
	Description string `json:"description,omitempty"` // Optional description of this course work material. The text must be a valid UTF-8 string containing no more than 30,000 characters.
	Materials []Material `json:"materials,omitempty"` // Additional materials. A course work material must have no more than 20 material items.
	Title string `json:"title,omitempty"` // Title of this course work material. The title must be a valid UTF-8 string containing between 1 and 3000 characters.
	Scheduledtime string `json:"scheduledTime,omitempty"` // Optional timestamp when this course work material is scheduled to be published.
	Topicid string `json:"topicId,omitempty"` // Identifier for the topic that this course work material is associated with. Must match an existing topic in the course.
	Assigneemode string `json:"assigneeMode,omitempty"` // Assignee mode of the course work material. If unspecified, the default value is `ALL_STUDENTS`.
	Creationtime string `json:"creationTime,omitempty"` // Timestamp when this course work material was created. Read-only.
	Individualstudentsoptions IndividualStudentsOptions `json:"individualStudentsOptions,omitempty"` // Assignee details about a coursework/announcement. This field is set if and only if `assigneeMode` is `INDIVIDUAL_STUDENTS`.
	Alternatelink string `json:"alternateLink,omitempty"` // Absolute link to this course work material in the Classroom web UI. This is only populated if `state` is `PUBLISHED`. Read-only.
}

// ListGuardiansResponse represents the ListGuardiansResponse schema from the OpenAPI specification
type ListGuardiansResponse struct {
	Guardians []Guardian `json:"guardians,omitempty"` // Guardians on this page of results that met the criteria specified in the request.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token identifying the next page of results to return. If empty, no further results are available.
}

// TimeOfDay represents the TimeOfDay schema from the OpenAPI specification
type TimeOfDay struct {
	Nanos int `json:"nanos,omitempty"` // Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	Seconds int `json:"seconds,omitempty"` // Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds.
	Hours int `json:"hours,omitempty"` // Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value "24:00:00" for scenarios like business closing time.
	Minutes int `json:"minutes,omitempty"` // Minutes of hour of day. Must be from 0 to 59.
}

// TurnInStudentSubmissionRequest represents the TurnInStudentSubmissionRequest schema from the OpenAPI specification
type TurnInStudentSubmissionRequest struct {
}

// YouTubeVideo represents the YouTubeVideo schema from the OpenAPI specification
type YouTubeVideo struct {
	Title string `json:"title,omitempty"` // Title of the YouTube video. Read-only.
	Alternatelink string `json:"alternateLink,omitempty"` // URL that can be used to view the YouTube video. Read-only.
	Id string `json:"id,omitempty"` // YouTube API resource ID.
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // URL of a thumbnail image of the YouTube video. Read-only.
}

// ListAnnouncementsResponse represents the ListAnnouncementsResponse schema from the OpenAPI specification
type ListAnnouncementsResponse struct {
	Announcements []Announcement `json:"announcements,omitempty"` // Announcement items that match the request.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token identifying the next page of results to return. If empty, no further results are available.
}

// Material represents the Material schema from the OpenAPI specification
type Material struct {
	Drivefile SharedDriveFile `json:"driveFile,omitempty"` // Drive file that is used as material for course work.
	Form Form `json:"form,omitempty"` // Google Forms item.
	Link Link `json:"link,omitempty"` // URL item.
	Youtubevideo YouTubeVideo `json:"youtubeVideo,omitempty"` // YouTube video item.
}

// Teacher represents the Teacher schema from the OpenAPI specification
type Teacher struct {
	Profile UserProfile `json:"profile,omitempty"` // Global information for a user.
	Userid string `json:"userId,omitempty"` // Identifier of the user. When specified as a parameter of a request, this identifier can be one of the following: * the numeric identifier for the user * the email address of the user * the string literal `"me"`, indicating the requesting user
	Courseid string `json:"courseId,omitempty"` // Identifier of the course. Read-only.
}

// StateHistory represents the StateHistory schema from the OpenAPI specification
type StateHistory struct {
	Actoruserid string `json:"actorUserId,omitempty"` // The teacher or student who made the change.
	State string `json:"state,omitempty"` // The workflow pipeline stage.
	Statetimestamp string `json:"stateTimestamp,omitempty"` // When the submission entered this state.
}

// ListTopicResponse represents the ListTopicResponse schema from the OpenAPI specification
type ListTopicResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token identifying the next page of results to return. If empty, no further results are available.
	Topic []Topic `json:"topic,omitempty"` // Topic items that match the request.
}

// ModifyIndividualStudentsOptions represents the ModifyIndividualStudentsOptions schema from the OpenAPI specification
type ModifyIndividualStudentsOptions struct {
	Addstudentids []string `json:"addStudentIds,omitempty"` // IDs of students to be added as having access to this coursework/announcement.
	Removestudentids []string `json:"removeStudentIds,omitempty"` // IDs of students to be removed from having access to this coursework/announcement.
}

// IndividualStudentsOptions represents the IndividualStudentsOptions schema from the OpenAPI specification
type IndividualStudentsOptions struct {
	Studentids []string `json:"studentIds,omitempty"` // Identifiers for the students that have access to the coursework/announcement.
}

// Feed represents the Feed schema from the OpenAPI specification
type Feed struct {
	Courserosterchangesinfo CourseRosterChangesInfo `json:"courseRosterChangesInfo,omitempty"` // Information about a `Feed` with a `feed_type` of `COURSE_ROSTER_CHANGES`.
	Courseworkchangesinfo CourseWorkChangesInfo `json:"courseWorkChangesInfo,omitempty"` // Information about a `Feed` with a `feed_type` of `COURSE_WORK_CHANGES`.
	Feedtype string `json:"feedType,omitempty"` // The type of feed.
}

// ModifyAttachmentsRequest represents the ModifyAttachmentsRequest schema from the OpenAPI specification
type ModifyAttachmentsRequest struct {
	Addattachments []Attachment `json:"addAttachments,omitempty"` // Attachments to add. A student submission may not have more than 20 attachments. Form attachments are not supported.
}

// GradebookSettings represents the GradebookSettings schema from the OpenAPI specification
type GradebookSettings struct {
	Calculationtype string `json:"calculationType,omitempty"` // Indicates how the overall grade is calculated.
	Displaysetting string `json:"displaySetting,omitempty"` // Indicates who can see the overall grade..
	Gradecategories []GradeCategory `json:"gradeCategories,omitempty"` // Grade categories that are available for coursework in the course.
}

// Link represents the Link schema from the OpenAPI specification
type Link struct {
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // URL of a thumbnail image of the target URL. Read-only.
	Title string `json:"title,omitempty"` // Title of the target of the URL. Read-only.
	Url string `json:"url,omitempty"` // URL to link to. This must be a valid UTF-8 string containing between 1 and 2024 characters.
}

// ModifyAnnouncementAssigneesRequest represents the ModifyAnnouncementAssigneesRequest schema from the OpenAPI specification
type ModifyAnnouncementAssigneesRequest struct {
	Modifyindividualstudentsoptions ModifyIndividualStudentsOptions `json:"modifyIndividualStudentsOptions,omitempty"` // Contains fields to add or remove students from a course work or announcement where the `assigneeMode` is set to `INDIVIDUAL_STUDENTS`.
	Assigneemode string `json:"assigneeMode,omitempty"` // Mode of the announcement describing whether it is accessible by all students or specified individual students.
}

// ReclaimStudentSubmissionRequest represents the ReclaimStudentSubmissionRequest schema from the OpenAPI specification
type ReclaimStudentSubmissionRequest struct {
}

// Announcement represents the Announcement schema from the OpenAPI specification
type Announcement struct {
	Courseid string `json:"courseId,omitempty"` // Identifier of the course. Read-only.
	Creationtime string `json:"creationTime,omitempty"` // Timestamp when this announcement was created. Read-only.
	Materials []Material `json:"materials,omitempty"` // Additional materials. Announcements must have no more than 20 material items.
	Scheduledtime string `json:"scheduledTime,omitempty"` // Optional timestamp when this announcement is scheduled to be published.
	Id string `json:"id,omitempty"` // Classroom-assigned identifier of this announcement, unique per course. Read-only.
	State string `json:"state,omitempty"` // Status of this announcement. If unspecified, the default state is `DRAFT`.
	Alternatelink string `json:"alternateLink,omitempty"` // Absolute link to this announcement in the Classroom web UI. This is only populated if `state` is `PUBLISHED`. Read-only.
	Assigneemode string `json:"assigneeMode,omitempty"` // Assignee mode of the announcement. If unspecified, the default value is `ALL_STUDENTS`.
	Creatoruserid string `json:"creatorUserId,omitempty"` // Identifier for the user that created the announcement. Read-only.
	Text string `json:"text,omitempty"` // Description of this announcement. The text must be a valid UTF-8 string containing no more than 30,000 characters.
	Individualstudentsoptions IndividualStudentsOptions `json:"individualStudentsOptions,omitempty"` // Assignee details about a coursework/announcement. This field is set if and only if `assigneeMode` is `INDIVIDUAL_STUDENTS`.
	Updatetime string `json:"updateTime,omitempty"` // Timestamp of the most recent change to this announcement. Read-only.
}

// DriveFile represents the DriveFile schema from the OpenAPI specification
type DriveFile struct {
	Title string `json:"title,omitempty"` // Title of the Drive item. Read-only.
	Alternatelink string `json:"alternateLink,omitempty"` // URL that can be used to access the Drive item. Read-only.
	Id string `json:"id,omitempty"` // Drive API resource ID.
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // URL of a thumbnail image of the Drive item. Read-only.
}

// Invitation represents the Invitation schema from the OpenAPI specification
type Invitation struct {
	Id string `json:"id,omitempty"` // Identifier assigned by Classroom. Read-only.
	Role string `json:"role,omitempty"` // Role to invite the user to have. Must not be `COURSE_ROLE_UNSPECIFIED`.
	Userid string `json:"userId,omitempty"` // Identifier of the invited user. When specified as a parameter of a request, this identifier can be set to one of the following: * the numeric identifier for the user * the email address of the user * the string literal `"me"`, indicating the requesting user
	Courseid string `json:"courseId,omitempty"` // Identifier of the course to invite the user to.
}

// UserProfile represents the UserProfile schema from the OpenAPI specification
type UserProfile struct {
	Permissions []GlobalPermission `json:"permissions,omitempty"` // Global permissions of the user. Read-only.
	Photourl string `json:"photoUrl,omitempty"` // URL of user's profile photo. Must request `https://www.googleapis.com/auth/classroom.profile.photos` scope for this field to be populated in a response body. Read-only.
	Verifiedteacher bool `json:"verifiedTeacher,omitempty"` // Represents whether a Google Workspace for Education user's domain administrator has explicitly verified them as being a teacher. This field is always false if the user is not a member of a Google Workspace for Education domain. Read-only
	Emailaddress string `json:"emailAddress,omitempty"` // Email address of the user. Must request `https://www.googleapis.com/auth/classroom.profile.emails` scope for this field to be populated in a response body. Read-only.
	Id string `json:"id,omitempty"` // Identifier of the user. Read-only.
	Name Name `json:"name,omitempty"` // Details of the user's name.
}

// ListCourseWorkResponse represents the ListCourseWorkResponse schema from the OpenAPI specification
type ListCourseWorkResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token identifying the next page of results to return. If empty, no further results are available.
	Coursework []CourseWork `json:"courseWork,omitempty"` // Course work items that match the request.
}

// ListInvitationsResponse represents the ListInvitationsResponse schema from the OpenAPI specification
type ListInvitationsResponse struct {
	Invitations []Invitation `json:"invitations,omitempty"` // Invitations that match the list request.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token identifying the next page of results to return. If empty, no further results are available.
}

// Form represents the Form schema from the OpenAPI specification
type Form struct {
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // URL of a thumbnail image of the Form. Read-only.
	Title string `json:"title,omitempty"` // Title of the Form. Read-only.
	Formurl string `json:"formUrl,omitempty"` // URL of the form.
	Responseurl string `json:"responseUrl,omitempty"` // URL of the form responses document. Only set if responses have been recorded and only when the requesting user is an editor of the form. Read-only.
}

// Empty represents the Empty schema from the OpenAPI specification
type Empty struct {
}

// StudentSubmission represents the StudentSubmission schema from the OpenAPI specification
type StudentSubmission struct {
	Assignedgrade float64 `json:"assignedGrade,omitempty"` // Optional grade. If unset, no grade was set. This value must be non-negative. Decimal (that is, non-integer) values are allowed, but are rounded to two decimal places. This may be modified only by course teachers.
	Assignmentsubmission AssignmentSubmission `json:"assignmentSubmission,omitempty"` // Student work for an assignment.
	Id string `json:"id,omitempty"` // Classroom-assigned Identifier for the student submission. This is unique among submissions for the relevant course work. Read-only.
	Updatetime string `json:"updateTime,omitempty"` // Last update time of this submission. This may be unset if the student has not accessed this item. Read-only.
	Userid string `json:"userId,omitempty"` // Identifier for the student that owns this submission. Read-only.
	Courseworkid string `json:"courseWorkId,omitempty"` // Identifier for the course work this corresponds to. Read-only.
	Associatedwithdeveloper bool `json:"associatedWithDeveloper,omitempty"` // Whether this student submission is associated with the Developer Console project making the request. See CreateCourseWork for more details. Read-only.
	Draftgrade float64 `json:"draftGrade,omitempty"` // Optional pending grade. If unset, no grade was set. This value must be non-negative. Decimal (that is, non-integer) values are allowed, but are rounded to two decimal places. This is only visible to and modifiable by course teachers.
	Late bool `json:"late,omitempty"` // Whether this submission is late. Read-only.
	Courseid string `json:"courseId,omitempty"` // Identifier of the course. Read-only.
	Courseworktype string `json:"courseWorkType,omitempty"` // Type of course work this submission is for. Read-only.
	Alternatelink string `json:"alternateLink,omitempty"` // Absolute link to the submission in the Classroom web UI. Read-only.
	State string `json:"state,omitempty"` // State of this submission. Read-only.
	Submissionhistory []SubmissionHistory `json:"submissionHistory,omitempty"` // The history of the submission (includes state and grade histories). Read-only.
	Multiplechoicesubmission MultipleChoiceSubmission `json:"multipleChoiceSubmission,omitempty"` // Student work for a multiple-choice question.
	Shortanswersubmission ShortAnswerSubmission `json:"shortAnswerSubmission,omitempty"` // Student work for a short answer question.
	Creationtime string `json:"creationTime,omitempty"` // Creation time of this submission. This may be unset if the student has not accessed this item. Read-only.
}

// GradeCategory represents the GradeCategory schema from the OpenAPI specification
type GradeCategory struct {
	Defaultgradedenominator int `json:"defaultGradeDenominator,omitempty"` // Default value of denominator. Only applicable when grade calculation type is TOTAL_POINTS.
	Id string `json:"id,omitempty"` // ID of the grade category.
	Name string `json:"name,omitempty"` // Name of the grade category.
	Weight int `json:"weight,omitempty"` // The weight of the category average as part of overall average. A weight of 12.34% is represented as 123400 (100% is 1,000,000). The last two digits should always be zero since we use two decimal precision. Only applicable when grade calculation type is WEIGHTED_CATEGORIES.
}

// CourseWork represents the CourseWork schema from the OpenAPI specification
type CourseWork struct {
	Creationtime string `json:"creationTime,omitempty"` // Timestamp when this course work was created. Read-only.
	Description string `json:"description,omitempty"` // Optional description of this course work. If set, the description must be a valid UTF-8 string containing no more than 30,000 characters.
	Maxpoints float64 `json:"maxPoints,omitempty"` // Maximum grade for this course work. If zero or unspecified, this assignment is considered ungraded. This must be a non-negative integer value.
	Courseid string `json:"courseId,omitempty"` // Identifier of the course. Read-only.
	Title string `json:"title,omitempty"` // Title of this course work. The title must be a valid UTF-8 string containing between 1 and 3000 characters.
	Updatetime string `json:"updateTime,omitempty"` // Timestamp of the most recent change to this course work. Read-only.
	Submissionmodificationmode string `json:"submissionModificationMode,omitempty"` // Setting to determine when students are allowed to modify submissions. If unspecified, the default value is `MODIFIABLE_UNTIL_TURNED_IN`.
	Topicid string `json:"topicId,omitempty"` // Identifier for the topic that this coursework is associated with. Must match an existing topic in the course.
	Alternatelink string `json:"alternateLink,omitempty"` // Absolute link to this course work in the Classroom web UI. This is only populated if `state` is `PUBLISHED`. Read-only.
	Assigneemode string `json:"assigneeMode,omitempty"` // Assignee mode of the coursework. If unspecified, the default value is `ALL_STUDENTS`.
	Gradecategory GradeCategory `json:"gradeCategory,omitempty"` // Details for a grade category in a course. Coursework may have zero or one grade category, and the category may be used in computing the overall grade. See the [help center article](https://support.google.com/edu/classroom/answer/9184995) for details.
	Associatedwithdeveloper bool `json:"associatedWithDeveloper,omitempty"` // Whether this course work item is associated with the Developer Console project making the request. See CreateCourseWork for more details. Read-only.
	Assignment Assignment `json:"assignment,omitempty"` // Additional details for assignments.
	Scheduledtime string `json:"scheduledTime,omitempty"` // Optional timestamp when this course work is scheduled to be published.
	Individualstudentsoptions IndividualStudentsOptions `json:"individualStudentsOptions,omitempty"` // Assignee details about a coursework/announcement. This field is set if and only if `assigneeMode` is `INDIVIDUAL_STUDENTS`.
	Materials []Material `json:"materials,omitempty"` // Additional materials. CourseWork must have no more than 20 material items.
	Duedate Date `json:"dueDate,omitempty"` // Represents a whole or partial calendar date, such as a birthday. The time of day and time zone are either specified elsewhere or are insignificant. The date is relative to the Gregorian Calendar. This can represent one of the following: * A full date, with non-zero year, month, and day values. * A month and day, with a zero year (for example, an anniversary). * A year on its own, with a zero month and a zero day. * A year and month, with a zero day (for example, a credit card expiration date). Related types: * google.type.TimeOfDay * google.type.DateTime * google.protobuf.Timestamp
	Id string `json:"id,omitempty"` // Classroom-assigned identifier of this course work, unique per course. Read-only.
	State string `json:"state,omitempty"` // Status of this course work. If unspecified, the default state is `DRAFT`.
	Worktype string `json:"workType,omitempty"` // Type of this course work. The type is set when the course work is created and cannot be changed.
	Duetime TimeOfDay `json:"dueTime,omitempty"` // Represents a time of day. The date and time zone are either not significant or are specified elsewhere. An API may choose to allow leap seconds. Related types are google.type.Date and `google.protobuf.Timestamp`.
	Creatoruserid string `json:"creatorUserId,omitempty"` // Identifier for the user that created the coursework. Read-only.
	Multiplechoicequestion MultipleChoiceQuestion `json:"multipleChoiceQuestion,omitempty"` // Additional details for multiple-choice questions.
}

// ListCourseAliasesResponse represents the ListCourseAliasesResponse schema from the OpenAPI specification
type ListCourseAliasesResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token identifying the next page of results to return. If empty, no further results are available.
	Aliases []CourseAlias `json:"aliases,omitempty"` // The course aliases.
}

// MultipleChoiceQuestion represents the MultipleChoiceQuestion schema from the OpenAPI specification
type MultipleChoiceQuestion struct {
	Choices []string `json:"choices,omitempty"` // Possible choices.
}

// Student represents the Student schema from the OpenAPI specification
type Student struct {
	Courseid string `json:"courseId,omitempty"` // Identifier of the course. Read-only.
	Profile UserProfile `json:"profile,omitempty"` // Global information for a user.
	Studentworkfolder DriveFolder `json:"studentWorkFolder,omitempty"` // Representation of a Google Drive folder.
	Userid string `json:"userId,omitempty"` // Identifier of the user. When specified as a parameter of a request, this identifier can be one of the following: * the numeric identifier for the user * the email address of the user * the string literal `"me"`, indicating the requesting user
}

// GlobalPermission represents the GlobalPermission schema from the OpenAPI specification
type GlobalPermission struct {
	Permission string `json:"permission,omitempty"` // Permission value.
}

// ListStudentSubmissionsResponse represents the ListStudentSubmissionsResponse schema from the OpenAPI specification
type ListStudentSubmissionsResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token identifying the next page of results to return. If empty, no further results are available.
	Studentsubmissions []StudentSubmission `json:"studentSubmissions,omitempty"` // Student work that matches the request.
}

// MultipleChoiceSubmission represents the MultipleChoiceSubmission schema from the OpenAPI specification
type MultipleChoiceSubmission struct {
	Answer string `json:"answer,omitempty"` // Student's select choice.
}

// Guardian represents the Guardian schema from the OpenAPI specification
type Guardian struct {
	Studentid string `json:"studentId,omitempty"` // Identifier for the student to whom the guardian relationship applies.
	Guardianid string `json:"guardianId,omitempty"` // Identifier for the guardian.
	Guardianprofile UserProfile `json:"guardianProfile,omitempty"` // Global information for a user.
	Invitedemailaddress string `json:"invitedEmailAddress,omitempty"` // The email address to which the initial guardian invitation was sent. This field is only visible to domain administrators.
}
