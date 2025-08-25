package main

import (
	"github.com/google-classroom-api/mcp-server/config"
	"github.com/google-classroom-api/mcp-server/models"
	tools_courses "github.com/google-classroom-api/mcp-server/tools/courses"
	tools_invitations "github.com/google-classroom-api/mcp-server/tools/invitations"
	tools_userprofiles "github.com/google-classroom-api/mcp-server/tools/userprofiles"
	tools_registrations "github.com/google-classroom-api/mcp-server/tools/registrations"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_courses.CreateClassroom_courses_coursework_studentsubmissions_turninTool(cfg),
		tools_courses.CreateClassroom_courses_announcements_modifyassigneesTool(cfg),
		tools_courses.CreateClassroom_courses_courseworkmaterials_deleteTool(cfg),
		tools_courses.CreateClassroom_courses_courseworkmaterials_getTool(cfg),
		tools_courses.CreateClassroom_courses_courseworkmaterials_patchTool(cfg),
		tools_courses.CreateClassroom_courses_announcements_listTool(cfg),
		tools_courses.CreateClassroom_courses_announcements_createTool(cfg),
		tools_invitations.CreateClassroom_invitations_listTool(cfg),
		tools_invitations.CreateClassroom_invitations_createTool(cfg),
		tools_courses.CreateClassroom_courses_topics_listTool(cfg),
		tools_courses.CreateClassroom_courses_topics_createTool(cfg),
		tools_courses.CreateClassroom_courses_coursework_studentsubmissions_getTool(cfg),
		tools_courses.CreateClassroom_courses_coursework_studentsubmissions_patchTool(cfg),
		tools_userprofiles.CreateClassroom_userprofiles_guardians_listTool(cfg),
		tools_courses.CreateClassroom_courses_coursework_studentsubmissions_returnTool(cfg),
		tools_courses.CreateClassroom_courses_aliases_listTool(cfg),
		tools_courses.CreateClassroom_courses_aliases_createTool(cfg),
		tools_courses.CreateClassroom_courses_aliases_deleteTool(cfg),
		tools_courses.CreateClassroom_courses_coursework_modifyassigneesTool(cfg),
		tools_courses.CreateClassroom_courses_coursework_createTool(cfg),
		tools_courses.CreateClassroom_courses_coursework_listTool(cfg),
		tools_courses.CreateClassroom_courses_announcements_deleteTool(cfg),
		tools_courses.CreateClassroom_courses_announcements_getTool(cfg),
		tools_courses.CreateClassroom_courses_announcements_patchTool(cfg),
		tools_courses.CreateClassroom_courses_coursework_studentsubmissions_modifyattachmentsTool(cfg),
		tools_courses.CreateClassroom_courses_coursework_studentsubmissions_listTool(cfg),
		tools_courses.CreateClassroom_courses_students_deleteTool(cfg),
		tools_courses.CreateClassroom_courses_students_getTool(cfg),
		tools_courses.CreateClassroom_courses_teachers_listTool(cfg),
		tools_courses.CreateClassroom_courses_teachers_createTool(cfg),
		tools_courses.CreateClassroom_courses_coursework_studentsubmissions_reclaimTool(cfg),
		tools_registrations.CreateClassroom_registrations_createTool(cfg),
		tools_courses.CreateClassroom_courses_deleteTool(cfg),
		tools_courses.CreateClassroom_courses_getTool(cfg),
		tools_courses.CreateClassroom_courses_patchTool(cfg),
		tools_courses.CreateClassroom_courses_updateTool(cfg),
		tools_courses.CreateClassroom_courses_coursework_deleteTool(cfg),
		tools_courses.CreateClassroom_courses_coursework_getTool(cfg),
		tools_courses.CreateClassroom_courses_coursework_patchTool(cfg),
		tools_courses.CreateClassroom_courses_topics_patchTool(cfg),
		tools_courses.CreateClassroom_courses_topics_deleteTool(cfg),
		tools_courses.CreateClassroom_courses_topics_getTool(cfg),
		tools_courses.CreateClassroom_courses_students_listTool(cfg),
		tools_courses.CreateClassroom_courses_students_createTool(cfg),
		tools_courses.CreateClassroom_courses_teachers_deleteTool(cfg),
		tools_courses.CreateClassroom_courses_teachers_getTool(cfg),
		tools_registrations.CreateClassroom_registrations_deleteTool(cfg),
		tools_userprofiles.CreateClassroom_userprofiles_guardianinvitations_listTool(cfg),
		tools_userprofiles.CreateClassroom_userprofiles_guardianinvitations_createTool(cfg),
		tools_userprofiles.CreateClassroom_userprofiles_guardians_getTool(cfg),
		tools_userprofiles.CreateClassroom_userprofiles_guardians_deleteTool(cfg),
		tools_userprofiles.CreateClassroom_userprofiles_guardianinvitations_getTool(cfg),
		tools_userprofiles.CreateClassroom_userprofiles_guardianinvitations_patchTool(cfg),
		tools_invitations.CreateClassroom_invitations_acceptTool(cfg),
		tools_userprofiles.CreateClassroom_userprofiles_getTool(cfg),
		tools_invitations.CreateClassroom_invitations_deleteTool(cfg),
		tools_invitations.CreateClassroom_invitations_getTool(cfg),
		tools_courses.CreateClassroom_courses_listTool(cfg),
		tools_courses.CreateClassroom_courses_createTool(cfg),
		tools_courses.CreateClassroom_courses_courseworkmaterials_createTool(cfg),
		tools_courses.CreateClassroom_courses_courseworkmaterials_listTool(cfg),
	}
}
