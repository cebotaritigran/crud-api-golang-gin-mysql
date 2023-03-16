package types

/*
	stucts variable names should be identical to the table's column names
*/
type ResultUsers struct {
	First_name string
	Last_name  string
	Authlevel  int
}

type ResultManagers struct {
	First_name  string
	Last_name   string
	PhoneNumber string
}

type ResultCampuses struct {
	Name        string
	City        string
	Type        string
	PhoneNumber string
}

type ResultPrograms struct {
	Name        string
	Description string
}

type ResultCampusManagers struct {
	Name         string
	Type         string
	City         string
	License_type string
	Manager      string
}

type ResultAreaDirectors struct {
	First_name  string
	Last_name   string
	PhoneNumber string
}

type ResultHeadDepartments struct {
	First_name  string
	Last_name   string
	PhoneNumber string
}

type ResultWebsiteBanners struct {
	Banner_name string
	Banner_url  string
	Row         int
}

type ResultCampusAccountDetails struct {
	DateIssued   string
	Type         string
	Name         string
	DocumentType string
	ProcessType  string
	Sum          int
	Currency     string
}

type ResultCampusAccountDetailsCredit struct {
	Euro  int
	Tl    int
	Dolar int
	Gbp   int
}

type ResultCampusAccountDetailsDebt struct {
	Euro  int
	Tl    int
	Dolar int
	Gbp   int
}

type ResultStudentAccountDetails struct {
	Datetour    string
	Tourname    string
	Student     string
	ProcessType string
	Account     string
	Currency    string
	PaymentDate string
	Sum         float32
	Sumtl       float32
	Rate        float32
}

type ResultTours struct {
	Term     string
	Tourname string
}

type ResultTourCosts struct {
	Date     string
	Currency string
	Sum      float32
	Sumtl    float32
	CostType string
	Bill     bool
}

type ResultCampusTours struct {
	DateStart    string
	DateEnd      string
	LicenseType  string
	Name         string
	PhoneNumber  string
	Quota        string
	QuotaRequest string
	Manager      string
	//manager phone number
	Phone_Number string
}

// testing
type ResultCampusToursResponsible struct {
	Name        string
	PhoneNumber string
}

type ResultCampusToursStundents struct {
	First_name     string
	Last_name      string
	SecurityNumber string
	BirthDate      string
	Gender         string
	Passport       string
	PhoneNumber    string
}

type ResultCampusTourWanted struct {
	Tourname     string
	QuotaRequest int
	Attendance   int
}

type ResultRecordTrack struct {
	First_name      string
	Last_name       string
	Type            string
	Name            string
	NameResponsible string
	SecurityNumber  string
	BirthDate       string
	Gender          string
	Passport        string
	Currency        string
	Sum             float32
	Sumtl           float32
}

type ResultTourRooming struct {
	First_name      string
	Last_name       string
	SecurityNumber  string
	BirthDate       string
	Gender          string
	Passport        string
	Name            string
	ResponsibleName string
	RoomNumber      int
}

type ResultCTMR struct {
	Tourname           string
	Datetour           string
	Type               string
	CampusName         string
	City               string
	Manager            string
	ManagerEmail       string
	ManagerPhoneNumber string
	ResponsibleName    string
	ResponsibleEmail   string
	ResponsiblePhone   string
	Status             int
}

type ResultStudentRecords struct {
	Abroad         bool
	Tourname       string
	Datetour       string
	First_name     string
	Last_name      string
	Name           string
	SecurityNumber string
	BirthDate      string
	Gender         string
	Passport       string
	Sum            float32
}

type ResultAllToursSummary struct {
	Datetour     string
	Tourname     string
	QuotaRequest int
	Quota        int
	CostOfTour   float32
	Total        float32
}

type LogStudents struct {
	Name           string
	First_name     string
	Last_name      string
	SecurityNumber string
	BirthDate      string
	Address        string
	Phone_Number   string
	Email          string
	ParentName     string
	ParentPhone    string
	ParentEmail    string
	Status         int
}

type RecordTrackReport struct {
	First_name  string
	Last_name   string
	CampusName  string
	Advisor     string
	Country     string
	Departmant  string
	Language    string
	ParentName  string
	ParentPhone string
}
