package constants

import (
	"github.com/VATUSA/primary-api/pkg/database/models"
	"slices"
)

type RoleID string

type Role struct {
	Name        string
	RolesCanAdd []RoleID
}

const (
	// ARTCC Roles
	AirTrafficManagerRole       RoleID = "ATM"
	DeputyAirTrafficManagerRole RoleID = "DATM"
	TrainingAdministratorRole   RoleID = "TA"
	EventCoordinatorRole        RoleID = "EC"
	AssistantEventCoordinator   RoleID = "AEC"
	FacilityEngineerRole        RoleID = "FE"
	AssistantFacilityEngineer   RoleID = "AFE"
	WebMasterRole               RoleID = "WM"
	AssistantWebMasterRole      RoleID = "AWM"
	InstructorRole              RoleID = "INS"
	MentorRole                  RoleID = "MTR"

	// Division Roles
	DivisionDirectorRole        RoleID = "USA1"
	AirTrafficServicesRole      RoleID = "USA2"
	TrainingServicesRole        RoleID = "USA3"
	SupportServicesRole         RoleID = "USA4"
	EventsManagerRole           RoleID = "USA5"
	TechnicalManagerRole        RoleID = "USA6"
	StaffDevelopmentManagerRole RoleID = "USA7"
	TrainingServicesManagerRole RoleID = "USA8"
	TrainingContentManagerRole  RoleID = "USA9"

	// Other Roles
	DeveloperTeamRole      RoleID = "DEV"
	AceTeamRole            RoleID = "ACE"
	NTMSRole               RoleID = "NTMS"
	NTMTRole               RoleID = "NTMT"
	SocialMediaTeam        RoleID = "SMT"
	TrainingContentTeam    RoleID = "TCT"
	AcademyMaterialEditor  RoleID = "CBT"
	FacilityMaterialEditor RoleID = "FACCBT"

	// Misc Roles
	EmailUser RoleID = "EMAIL"
)

var Roles = map[RoleID]Role{
	// ARTCC Roles
	AirTrafficManagerRole: {
		Name: "Air Traffic Manager",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			SupportServicesRole,
			TechnicalManagerRole,
			DeveloperTeamRole,
		},
	},
	DeputyAirTrafficManagerRole: {
		Name: "Deputy Air Traffic Manager",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			SupportServicesRole,
			TechnicalManagerRole,
			DeveloperTeamRole,
			AirTrafficManagerRole,
		},
	},
	TrainingAdministratorRole: {
		Name: "Training Administrator",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			TrainingServicesRole,
			SupportServicesRole,
			TechnicalManagerRole,
			TrainingServicesManagerRole,
			DeveloperTeamRole,
		},
	},
	EventCoordinatorRole: {
		Name: "Event Coordinator",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			SupportServicesRole,
			EventsManagerRole,
			TechnicalManagerRole,
			DeveloperTeamRole,
			AirTrafficManagerRole,
			DeputyAirTrafficManagerRole,
		},
	},
	AssistantEventCoordinator: {
		Name: "Assistant Event Coordinator",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			SupportServicesRole,
			EventsManagerRole,
			TechnicalManagerRole,
			DeveloperTeamRole,
			AirTrafficManagerRole,
			DeputyAirTrafficManagerRole,
			EventCoordinatorRole,
		},
	},
	FacilityEngineerRole: {
		Name: "Facility Engineer",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			SupportServicesRole,
			TechnicalManagerRole,
			DeveloperTeamRole,
			AirTrafficManagerRole,
			DeputyAirTrafficManagerRole,
		},
	},
	AssistantFacilityEngineer: {
		Name: "Assistant Facility Engineer",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			SupportServicesRole,
			TechnicalManagerRole,
			DeveloperTeamRole,
			AirTrafficManagerRole,
			DeputyAirTrafficManagerRole,
			FacilityEngineerRole,
		},
	},
	WebMasterRole: {
		Name: "Webmaster",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			SupportServicesRole,
			TechnicalManagerRole,
			DeveloperTeamRole,
			AirTrafficManagerRole,
			DeputyAirTrafficManagerRole,
		},
	},
	AssistantWebMasterRole: {
		Name: "Assistant Webmaster",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			SupportServicesRole,
			TechnicalManagerRole,
			DeveloperTeamRole,
			AirTrafficManagerRole,
			DeputyAirTrafficManagerRole,
			WebMasterRole,
		},
	},
	InstructorRole: {
		Name: "Instructor",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			TrainingServicesRole,
			TrainingServicesManagerRole,
			TechnicalManagerRole,
			DeveloperTeamRole,
		},
	},
	MentorRole: {
		Name: "Mentor",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			TrainingServicesRole,
			TrainingServicesManagerRole,
			TechnicalManagerRole,
			DeveloperTeamRole,
			AirTrafficManagerRole,
			DeputyAirTrafficManagerRole,
			TrainingAdministratorRole,
		},
	},

	// Division Roles
	DivisionDirectorRole: {
		Name: "Division Director",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			TrainingServicesRole,
			TechnicalManagerRole,
		},
	},
	AirTrafficServicesRole: {
		Name: "Deputy Director Air Traffic Services",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			TechnicalManagerRole,
		},
	},
	TrainingServicesRole: {
		Name: "Deputy Director Training Services",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			TrainingServicesRole,
			TechnicalManagerRole,
		},
	},
	SupportServicesRole: {
		Name: "Deputy Director Support Services",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			SupportServicesRole,
			TechnicalManagerRole,
		},
	},
	EventsManagerRole: {
		Name: "Events Manager",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			SupportServicesRole,
			EventsManagerRole,
			TechnicalManagerRole,
		},
	},
	TechnicalManagerRole: {
		Name: "Technical Manager",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			SupportServicesRole,
			TechnicalManagerRole,
		},
	},
	StaffDevelopmentManagerRole: {
		Name: "Staff Development Manager",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			SupportServicesRole,
			TechnicalManagerRole,
		},
	},
	TrainingServicesManagerRole: {
		Name: "Training Services Manager",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			TrainingServicesRole,
			TechnicalManagerRole,
		},
	},
	TrainingContentManagerRole: {
		Name: "Training Content Manager",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			TrainingServicesRole,
			TechnicalManagerRole,
		},
	},

	// Other Roles
	DeveloperTeamRole: {
		Name: "Developer Team",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			SupportServicesRole,
			TechnicalManagerRole,
		},
	},
	AceTeamRole: {
		Name: "ACE Team",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			SupportServicesRole,
			EventsManagerRole,
			TechnicalManagerRole,
		},
	},
	NTMSRole: {
		Name: "National Traffic Management Supervisor",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			SupportServicesRole,
			EventsManagerRole,
			TechnicalManagerRole,
		},
	},
	NTMTRole: {
		Name: "National Training Management Supervisor",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			TrainingServicesRole,
			EventsManagerRole,
			TechnicalManagerRole,
		},
	},
	SocialMediaTeam: {
		Name: "Social Media Team",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			SupportServicesRole,
			TechnicalManagerRole,
		},
	},
	TrainingContentTeam: {
		Name: "Training Content Team",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			TrainingServicesRole,
			TrainingServicesManagerRole,
			TrainingContentManagerRole,
			TechnicalManagerRole,
		},
	},
	AcademyMaterialEditor: {
		Name: "Academy Material Editor (Academy)",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			TrainingServicesRole,
			TrainingServicesManagerRole,
			TrainingContentManagerRole,
			TechnicalManagerRole,
		},
	},
	FacilityMaterialEditor: {
		Name: "Academy Material Editor (Facility)",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			TrainingServicesRole,
			TrainingServicesManagerRole,
			TrainingContentManagerRole,
			TechnicalManagerRole,
			AirTrafficManagerRole,
			DeputyAirTrafficManagerRole,
			TrainingAdministratorRole,
		},
	},

	// Misc Roles
	EmailUser: {
		Name: "Email User",
		RolesCanAdd: []RoleID{
			DivisionDirectorRole,
			AirTrafficServicesRole,
			TrainingServicesRole,
			SupportServicesRole,
			EventsManagerRole,
			TechnicalManagerRole,
			StaffDevelopmentManagerRole,
			TrainingServicesManagerRole,
			TrainingContentManagerRole,

			AirTrafficManagerRole,
			DeputyAirTrafficManagerRole,
			TrainingAdministratorRole,
			EventCoordinatorRole,
			FacilityEngineerRole,
			WebMasterRole,
		},
	},
}

var (
	FacilityGroup = []RoleID{
		AirTrafficManagerRole,
		DeputyAirTrafficManagerRole,
		TrainingAdministratorRole,
		EventCoordinatorRole,
		AssistantEventCoordinator,
		FacilityEngineerRole,
		AssistantFacilityEngineer,
		WebMasterRole,
		AssistantWebMasterRole,
		InstructorRole,
		MentorRole,
		FacilityMaterialEditor,
		EmailUser,
	}

	FacilityStaffGroup = []RoleID{
		AirTrafficManagerRole,
		DeputyAirTrafficManagerRole,
		TrainingAdministratorRole,
		EventCoordinatorRole,
		FacilityEngineerRole,
		WebMasterRole,
	}
	FacilitySeniorStaffGroup = []RoleID{
		AirTrafficManagerRole,
		DeputyAirTrafficManagerRole,
		TrainingAdministratorRole,
	}

	FacilityJuniorStaffGroup = []RoleID{
		EventCoordinatorRole,
		FacilityEngineerRole,
		WebMasterRole,
	}

	FacilityAssistantGroup = []RoleID{
		AssistantEventCoordinator,
		AssistantFacilityEngineer,
		AssistantWebMasterRole,
	}

	DivisionSeniorStaffGroup = []RoleID{
		DivisionDirectorRole,
		AirTrafficServicesRole,
		TrainingServicesRole,
		SupportServicesRole,
	}

	DivisionJuniorStaffGroup = []RoleID{
		EventsManagerRole,
		TechnicalManagerRole,
		StaffDevelopmentManagerRole,
		TrainingServicesManagerRole,
		TrainingContentManagerRole,
	}
)

func (r RoleID) DisplayName() string {
	return Roles[r].Name
}

func CanModifyRole(user *models.User, role RoleID) bool {
	if _, ok := Roles[role]; !ok {
		return false
	}

	return HasRoleList(user, Roles[role].RolesCanAdd)
}

func HasRoleList(user *models.User, roles []RoleID) bool {
	for _, role := range roles {
		if HasRole(user, role) {
			return true
		}
	}
	return false
}

func HasRole(user *models.User, role RoleID) bool {
	for _, r := range user.Roles {
		if r.RoleID == role {
			return true
		}
	}
	return false
}

func (r RoleID) isFacilityStaff() bool {
	return slices.Contains(FacilityGroup, r)
}

func (r RoleID) isFacilitySeniorStaff() bool {
	return slices.Contains(FacilitySeniorStaffGroup, r)
}

func (r RoleID) isFacilityJuniorStaff() bool {
	return slices.Contains(FacilityJuniorStaffGroup, r)
}

func IsValidRole(role RoleID) bool {
	_, ok := Roles[role]
	return ok
}
