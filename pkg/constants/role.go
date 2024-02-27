package constants

import "slices"

type Role string

const (
	AirTrafficManagerRole       Role = "ATM"
	DeputyAirTrafficManagerRole Role = "DATM"
	TrainingAdministratorRole   Role = "TA"
	EventCoordinatorRole        Role = "EC"
	FacilityEngineerRole        Role = "FE"
	WebMasterRole               Role = "WM"
	InstructorRole              Role = "INS"
	MentorRole                  Role = "MTR"
	DeveloperTeamRole           Role = "DEVELOPER"
	DivisionStaffRole           Role = "DIVISION_STAFF"
	DivisionManagementRole      Role = "DIVISION_MANAGEMENT"
	AceTeamRole                 Role = "ACE"
	NTMSRole                    Role = "NTMS"
	NTMTRole                    Role = "NTMT"
	SocialMediaTeam             Role = "SMT"
	EmailUser                   Role = "EMAIL"
	TrainingContentTeam         Role = "TCT"
	AcademyMaterialEditor       Role = "CBT"
	FacilityMaterialEditor      Role = "FACCBT"
)

var (
	RoleDisplayNameMap = map[Role]string{
		AirTrafficManagerRole:       "Air Traffic Manager",
		DeputyAirTrafficManagerRole: "Deputy Air Traffic Manager",
		TrainingAdministratorRole:   "Training Administrator",
		EventCoordinatorRole:        "Event Coordinator",
		FacilityEngineerRole:        "Facility Engineer",
		WebMasterRole:               "Webmaster",
		InstructorRole:              "Instructor",
		MentorRole:                  "Mentor",
		DeveloperTeamRole:           "Developer Team",
		DivisionStaffRole:           "Division Staff",
		DivisionManagementRole:      "Division Management",
		AceTeamRole:                 "ACE Team",
		NTMSRole:                    "National Traffic Management Supervisor",
		NTMTRole:                    "National Traffic Management Supervisor",
		SocialMediaTeam:             "Social Media Team",
		EmailUser:                   "Email User",
		TrainingContentTeam:         "Training Content Team",
		AcademyMaterialEditor:       "Academy Material Editor (Academy)",
		FacilityMaterialEditor:      "Academy Material Editor (Facility)",
	}
	FacilityBasedRoles = []Role{
		AirTrafficManagerRole,
		DeputyAirTrafficManagerRole,
		TrainingAdministratorRole,
		EventCoordinatorRole,
		FacilityEngineerRole,
		WebMasterRole,
		InstructorRole,
		MentorRole,
		FacilityMaterialEditor,
		EmailUser,
	}
	FacilityStaffRoles = []Role{
		AirTrafficManagerRole,
		DeputyAirTrafficManagerRole,
		TrainingAdministratorRole,
		EventCoordinatorRole,
		FacilityEngineerRole,
		WebMasterRole,
	}
	FacilityManagementRoles = []Role{
		AirTrafficManagerRole,
		DeputyAirTrafficManagerRole,
		TrainingAdministratorRole,
	}
	// FacilityManagedRoles can be assigned by AirTrafficManagerRole / DeputyAirTrafficManagerRole
	FacilityManagedRoles = []Role{
		EventCoordinatorRole,
		FacilityEngineerRole,
		WebMasterRole,
		MentorRole,
		FacilityMaterialEditor,
		EmailUser,
	}

	// AssignmentRestrictedRoles can only be assigned by DivisionManagementRole
	AssignmentRestrictedRoles = []Role{
		DivisionStaffRole,
		DivisionManagementRole,
	}
)

func (r Role) DisplayName() string {
	return RoleDisplayNameMap[r]
}

func (r Role) isFacilityBased() bool {
	return slices.Contains(FacilityBasedRoles, r)
}

func (r Role) isFacilityStaff() bool {
	return slices.Contains(FacilityStaffRoles, r)
}

func (r Role) isFacilityManagement() bool {
	return slices.Contains(FacilityManagementRoles, r)
}

func (r Role) isFacilityManaged() bool {
	return slices.Contains(FacilityManagedRoles, r)
}

func (r Role) isAssignmentRestricted() bool {
	return slices.Contains(AssignmentRestrictedRoles, r)
}
