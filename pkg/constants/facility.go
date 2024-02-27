package constants

type Facility string

const (
	// Special Facilities
	AcademyFacility      Facility = "ZAE"
	HeadquartersFacility Facility = "ZHQ"
	NonMemberFacility    Facility = "ZZN"
	InactiveFacility     Facility = "ZZI"

	// ARTCC Facilities
	AlbuquerqueFacility  Facility = "ZAB"
	AnchorageFacility    Facility = "ZAN"
	AtlantaFacility      Facility = "ZTL"
	BostonFacility       Facility = "ZBW"
	ChicagoFacility      Facility = "ZAU"
	ClevelandFacility    Facility = "ZOB"
	DenverFacility       Facility = "ZDV"
	FortWorthFacility    Facility = "ZFW"
	HonoluluFacility     Facility = "HCF"
	HoustonFacility      Facility = "ZHU"
	IndianapolisFacility Facility = "ZID"
	JacksonvilleFacility Facility = "ZJX"
	KansasCityFacility   Facility = "ZKC"
	LosAngelesFacility   Facility = "ZLA"
	MemphisFacility      Facility = "ZME"
	MiamiFacility        Facility = "ZMA"
	MinneapolisFacility  Facility = "ZMP"
	NewYorkFacility      Facility = "ZNY"
	OaklandFacility      Facility = "ZOA"
	SaltLakeFacility     Facility = "ZLC"
	SeattleFacility      Facility = "ZSE"
	WashingtonFacility   Facility = "ZDC"
)

var (
	FacilityDisplayNameMap = map[Facility]string{
		AcademyFacility:      "Academy",
		HeadquartersFacility: "Headquarters",
		NonMemberFacility:    "Non-Member",
		InactiveFacility:     "Inactive",

		AlbuquerqueFacility:  "Albuquerque ARTCC",
		AnchorageFacility:    "Anchorage ARTCC",
		AtlantaFacility:      "Atlanta ARTCC",
		BostonFacility:       "Boston ARTCC",
		ChicagoFacility:      "Chicago ARTCC",
		ClevelandFacility:    "Cleveland ARTCC",
		DenverFacility:       "Denver ARTCC",
		FortWorthFacility:    "Fort Worth ARTCC",
		HonoluluFacility:     "Honolulu Control Facility",
		HoustonFacility:      "Houston ARTCC",
		IndianapolisFacility: "Indianapolis ARTCC",
		JacksonvilleFacility: "Jacksonville ARTCC",
		KansasCityFacility:   "Kansas City ARTCC",
		LosAngelesFacility:   "Los Angeles ARTCC",
		MemphisFacility:      "Memphis ARTCC",
		MiamiFacility:        "Miami ARTCC",
		MinneapolisFacility:  "Minneapolis ARTCC",
		NewYorkFacility:      "New York ARTCC",
		OaklandFacility:      "Oakland ARTCC",
		SaltLakeFacility:     "Salt Lake ARTCC",
		SeattleFacility:      "Seattle ARTCC",
		WashingtonFacility:   "Washington, D.C. ARTCC",
	}
)

func (f Facility) DisplayName() string {
	return FacilityDisplayNameMap[f]
}
