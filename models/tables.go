package models

import (
	"time"
)

type Campuses struct {
	ID          int       `gorm:"not null;primaryKey;autoIncrement"`
	Manager_id  int       `gorm:"type:int(11)"`
	Name        string    `gorm:"not null;type:varchar(115);default:''" json:"name"`
	Type        string    `gorm:"not null;type:varchar(115);default:''" json:"type"`
	City        string    `gorm:"not null;type:varchar(115);default:''" json:"city"`
	LicenseType string    `gorm:"not null;type:varchar(115);default:''" json:"licensetype"`
	PhoneNumber string    `gorm:"not null;type:varchar(115);default:''" json:"phonenumber"`
	Address     string    `gorm:"not null;type:varchar(115);default:''" json:"address"`
	IsActive    bool      `gorm:"not null;type:tinyint(4);default:1;" json:"isactive"`
	CreatedAt   time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
	UpdatedAt   time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
}

type Managers struct {
	ID          int       `gorm:"not null;primaryKey;autoIncrement"`
	FirstName   string    `gorm:"not null;type:varchar(115);default:''" json:"firstname"`
	LastName    string    `gorm:"not null;type:varchar(115);default:''" json:"lastname"`
	PhoneNumber string    `gorm:"not null;type:varchar(115);default:''" json:"phonenumber"`
	Email       string    `gorm:"not null;type:varchar(115);default:''" json:"email"`
	Password    string    `gorm:"not null;type:varchar(115);default:''" json:"password"`
	IsActive    bool      `gorm:"not null;type:tinyint(4);default:1;" json:"isactive"`
	CreatedAt   time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
	UpdatedAt   time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
}

type Programs struct {
	ID                          int       `gorm:"not null;primaryKey;autoIncrement"`
	Name                        string    `gorm:"not null;type:varchar(115);default:''" json:"name"`
	Description                 string    `gorm:"not null;type:varchar(400);default:''" json:"description"`
	IncludedServices            string    `gorm:"not null;type:varchar(400);default:''" json:"includedservices"`
	ExcludedServices            string    `gorm:"not null;type:varchar(400);default:''" json:"excludedservices"`
	SocialAndCulturalActivities string    `gorm:"not null;type:varchar(400);default:''" json:"socialactivities"`
	Type                        string    `gorm:"not null;type:varchar(115);default:''" json:"type"`
	AgeGroup                    string    `gorm:"not null;type:varchar(115);default:''" json:"agegroup"`
	Countries                   string    `gorm:"not null;type:varchar(115);default:''" json:"countries"`
	Cities                      string    `gorm:"not null;type:varchar(200);default:''" json:"cities"`
	Department                  string    `gorm:"not null;type:varchar(115);default:''" json:"department"`
	Acquisitions                string    `gorm:"not null;type:varchar(400);default:''" json:"acquisitions"`
	TransportationType          string    `gorm:"not null;type:varchar(115);default:''" json:"transportationtype"`
	IsActive                    bool      `gorm:"not null;type:tinyint(4);default:1;" json:"isactive"`
	Agreement_file              string    `gorm:"not null;type:varchar(115);default:'/'" json:"agreement_file"`
	VisaDocument_file           string    `gorm:"not null;type:varchar(115);default:'/'" json:"visadocument_file"`
	Brochure_file               string    `gorm:"not null;type:varchar(115);default:'/'" json:"brochure_file"`
	BannerRight_file            string    `gorm:"not null;type:varchar(115);default:'/'" json:"bannerright_file"`
	BannerLeft_file             string    `gorm:"not null;type:varchar(115);default:'/'" json:"bannerleft_file"`
	BannerRightIsActive         bool      `gorm:"not null;type:tinyint(4);default:1;" json:"bannerrightisactive"`
	BannerLeftIsActive          bool      `gorm:"not null;type:tinyint(4);default:1;" json:"bannerleftisactive"`
	CreatedAt                   time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
	UpdatedAt                   time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
}

type AreaDirector struct {
	ID          int       `gorm:"not null;primaryKey;autoIncrement"`
	FirstName   string    `gorm:"not null;type:varchar(115);default:''" json:"firstname"`
	LastName    string    `gorm:"not null;type:varchar(115);default:''" json:"lastname"`
	Email       string    `gorm:"not null;type:varchar(115);default:''" json:"email"`
	PhoneNumber string    `gorm:"not null;type:varchar(115);default:''" json:"phonenumber"`
	Password    string    `gorm:"not null;type:varchar(115);default:''" json:"password"`
	IsActive    bool      `gorm:"not null;type:tinyint(4);default:1;" json:"isactive"`
	CreatedAt   time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
	UpdatedAt   time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
}

type DepartmentHead struct {
	ID          int       `gorm:"not null;primaryKey;autoIncrement"`
	FirstName   string    `gorm:"not null;type:varchar(115);default:''" json:"firstname"`
	LastName    string    `gorm:"not null;type:varchar(115);default:''" json:"lastname"`
	Email       string    `gorm:"not null;type:varchar(115);default:''" json:"email"`
	PhoneNumber string    `gorm:"not null;type:varchar(115);default:''" json:"phonenumber"`
	Password    string    `gorm:"not null;type:varchar(115);default:''" json:"password"`
	IsActive    bool      `gorm:"not null;type:tinyint(4);default:1;" json:"isactive"`
	CreatedAt   time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
	UpdatedAt   time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
}

type WebsiteBanner struct {
	ID           int       `gorm:"not null;primaryKey;autoIncrement"`
	BannerName   string    `gorm:"not null;type:varchar(115);default:''" json:"bannername"`
	Description  string    `gorm:"not null;type:varchar(400);default:''" json:"description"`
	BannerUrl    string    `gorm:"not null;type:varchar(400);default:''" json:"banner_url"`
	Window       string    `gorm:"not null;type:varchar(115);default:''" json:"window"`
	Row          int       `gorm:"not null;type:int(11);default:0" json:"row"`
	IsActive     bool      `gorm:"not null;type:tinyint(4);default:1;" json:"isactive"`
	CategoryFile string    `gorm:"not null;type:varchar(400);default:'/'" json:"categoryfile"`
	CreatedAt    time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
	UpdatedAt    time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
}

type Tour struct {
	ID        int       `gorm:"not null;primaryKey;autoIncrement"`
	Quota     int       `gorm:"not null;type:int;default:100" json:"quota"`
	TourCost  float32   `gorm:"not null;type:decimal(7,2);default:0" json:"cost"`
	TourYear  time.Time `gorm:"not null;type:date" json:"term"`
	Type      string    `gorm:"not null;type:varchar(115);default:''" json:"type"`
	DateStart time.Time `gorm:"not null;type:date" json:"datestart"`
	DateEnd   time.Time `gorm:"not null;type:date" json:"dateend"`
	TourName  string    `gorm:"not null;type:varchar(115);default:''" json:"tourname"`
	Abroad    bool      `gorm:"not null;type:tinyint(4);default:1;" json:"abroad"`
	CreatedAt time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
	UpdatedAt time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
}

type CampusPayments struct {
	ID           int       `gorm:"not null;primaryKey;autoIncrement"`
	DateIssued   time.Time `gorm:"not null;type:date"`
	DocumentType string    `gorm:"not null;type:varchar(115);default:''" json:"documenttype"`
	ProcessType  string    `gorm:"not null;type:varchar(115);default:'Devir Bakiyesi'" json:"processtype"`
	Sum          int       `gorm:"not null;type:int(11);default:0" json:"sum"`
	Currency     string    `gorm:"not null;type:varchar(115);default:''" json:"currency"`
	CreatedAt    time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
	UpdatedAt    time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
}

type CampusesAndPayments struct {
	ID                int       `gorm:"not null;primaryKey;autoIncrement"`
	Tour_id           int       `gorm:"type:int(11)" json:"tourid"`
	Campus_id         int       `gorm:"type:int(11)" json:"campuseid"`
	CampusPayments_id int       `gorm:"type:int(11)" json:"campuspaymentsid"`
	CreatedAt         time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
	UpdatedAt         time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
}

type Students struct {
	ID             int       `gorm:"not null;primaryKey;autoIncrement"`
	Campus_id      int       `gorm:"type:int(11)" json:"campuseid"`
	Parent_id      int       `gorm:"type:int(11)" json:"parentid"`
	FirstName      string    `gorm:"not null;type:varchar(115);default:''" json:"firstname"`
	LastName       string    `gorm:"not null;type:varchar(115);default:''" json:"lastname"`
	SecurityNumber string    `gorm:"not null;type:varchar(115);default:''" json:"securitynumber"`
	BirthDate      time.Time `gorm:"not null;type:date"`
	Address        string    `gorm:"not null;type:varchar(115);default:''" json:"address"`
	PhoneNumber    string    `gorm:"not null;type:varchar(115);default:''" json:"phonenumber"`
	Email          string    `gorm:"not null;type:varchar(115);default:''" json:"email"`
	Passport       string    `gorm:"not null;type:varchar(115);default:'Secilmemis'" json:"passport"`
	Gender         string    `gorm:"not null; type:varchar(40);default:''" json:"gender"`
	Departmant     string    `gorm:"not null;type:varchar(115);default:''" json:"departmant"`
	Status         bool      `gorm:"not null;type:tinyint(4);default:0;" json:"status"`
	CreatedAt      time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
	UpdatedAt      time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
}

type StudentPayments struct {
	ID          int       `gorm:"not null;primaryKey;autoIncrement"`
	ProcessType string    `gorm:"not null;type:varchar(115);default:''" json:"processtype"`
	Currency    string    `gorm:"not null;type:varchar(115);default:''" json:"currency"`
	Account     string    `gorm:"not null;type:varchar(115);default:'Tur Odeme'" json:"account"`
	Sum         float32   `gorm:"not null;type:decimal(7,2);default:0" json:"sum"`
	SumTl       float32   `gorm:"not null;type:decimal(7,2);default:0" json:"sumtl"`
	Rate        float32   `gorm:"not null;type:decimal(5,2);default:1" json:"rate"`
	PaymentDate time.Time `gorm:"not null;type:date" json:"paymentdate"`
	CreatedAt   time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
	UpdatedAt   time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
}

type StudentsAndPayments struct {
	ID                 int       `gorm:"not null;primaryKey;autoIncrement"`
	Tour_id            int       `gorm:"type:int(11)" json:"tourid"`
	Student_id         int       `gorm:"type:int(11)" json:"studentid"`
	StudentPayments_id int       `gorm:"type:int(11)" json:"studentpaymentsid"`
	CreatedAt          time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
	UpdatedAt          time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
}

type TourExpenses struct {
	ID          int       `gorm:"not null;primaryKey;autoIncrement"`
	Date        time.Time `gorm:"not null;type:date" json:"date"`
	CostType    string    `gorm:"not null;type:varchar(115);default:''" json:"costtype"`
	Currency    string    `gorm:"not null;type:varchar(115);default:''" json:"currency"`
	Sum         float32   `gorm:"not null;type:decimal(7,2);default:0" json:"sum"`
	Sumtl       float32   `gorm:"not null;type:decimal(7,2);default:0" json:"sumtl"`
	Description string    `gorm:"not null;type:varchar(400);default:''" json:"description"`
	Bill        bool      `gorm:"not null;type:tinyint(4);default:1;" json:"bill"`
	CreatedAt   time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
	UpdatedAt   time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
}

type ToursAndExpenses struct {
	ID              int       `gorm:"not null;primaryKey;autoIncrement"`
	Tour_id         int       `gorm:"type:int(11)" json:"tourid"`
	TourExpenses_id int       `gorm:"type:int(11)" json:"tourexpensesid"`
	CreatedAt       time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
	UpdatedAt       time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
}

type CampusTour struct {
	ID                   int `gorm:"not null;primaryKey;autoIncrement"`
	Campus_id            int `gorm:"type:int(11)" json:"campuseid"`
	CampusTourRequest_id int `gorm:"type:int(11)" json:"campusetourrequestid"`
	Tour_id              int `gorm:"type:int(11)" json:"tourid"`
	// foreign key to Tour Record track to check if student is going
	TourTrackRecord_id int       `gorm:"type:int(11)"`
	Wanted             int       `gorm:"not null;type:int;default:1" json:"wanted"`
	CreatedAt          time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
	UpdatedAt          time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
}

type Responsible struct {
	ID          int       `gorm:"not null;primaryKey;autoIncrement"`
	Name        string    `gorm:"not null;type:varchar(115);default:''" json:"name"`
	Email       string    `gorm:"not null;type:varchar(115);default:''" json:"email"`
	PhoneNumber string    `gorm:"not null;type:varchar(115);default:''" json:"phonenumber"`
	CreatedAt   time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
	UpdatedAt   time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
}

type Parents struct {
	ID          int       `gorm:"not null;primaryKey;autoIncrement"`
	Name        string    `gorm:"not null;type:varchar(115);default:''" json:"name"`
	PhoneNumber string    `gorm:"not null;type:varchar(115);default:''" json:"phonenumber"`
	Email       string    `gorm:"not null;type:varchar(115);default:''" json:"email"`
	CreatedAt   time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
	UpdatedAt   time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
}

type CampusTourRequest struct {
	ID           int       `gorm:"not null;primaryKey;autoIncrement"`
	Campus_id    int       `gorm:"type:int(11)" json:"campuseid"`
	Tour_id      int       `gorm:"type:int(11)" json:"tourid"`
	QuotaRequest int       `gorm:"not null;type:int(11);default:0" json:"quotarequest"`
	Attendance   int       `gorm:"not null;type:int(11);default:0" json:"attendance"`
	CreatedAt    time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
	UpdatedAt    time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
}

type TourTrackRecord struct {
	ID                          int       `gorm:"not null;primaryKey;autoIncrement"`
	Student_id                  int       `gorm:"type:int(11)" json:"studentid"`
	Campus_id                   int       `gorm:"type:int(11)" json:"campuseid"`
	Tour_id                     int       `gorm:"type:int(11)" json:"tourid"`
	Tour_campuse_responsible_id int       `gorm:"type:int(11)" json:"tourcampuserid"`
	StudentsAndPayments_id      int       `gorm:"type:int(11)" json:"studentandpaymentsid"`
	Canceled                    bool      `gorm:"not null;type:tinyint(4);default:0;" json:"canceled"`
	CreatedAt                   time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
	UpdatedAt                   time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
}

type TourRooming struct {
	ID         int `gorm:"not null;primaryKey;autoIncrement"`
	Student_id int `gorm:"type:int(11)" json:"studentid"`
	Campus_id  int `gorm:"type:int(11)" json:"campuseid"`
	Tour_id    int `gorm:"type:int(11)" json:"tourid"`
	// foreign key to relational database between responsibles tours and campuses
	Tour_campuse_responsible_id int       `gorm:"type:int(11)" json:"tourcampuserid"`
	RoomNumber                  string    `gorm:"not null;type:varchar(115);default:''" json:"roomnumber"`
	CreatedAt                   time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
	UpdatedAt                   time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
}

type TourCampusManagerResponsible struct {
	ID int `gorm:"not null;primaryKey;autoIncrement"`
	// foreign key to tours
	Tour_id int `gorm:"type:int(11)" json:"tourid"`
	//foreign key to campuses
	Campus_id int `gorm:"type:int(11)" json:"campuseid"`
	//foreign key to responsibles
	Responsible_id int       `gorm:"type:int(11)" json:"responsibleid"`
	Status         bool      `gorm:"not null;type:tinyint(4);default:0;" json:"status"`
	CreatedAt      time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
	UpdatedAt      time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
}

type AdvisorTrackReport struct {
	ID         int       `gorm:"not null;primaryKey;autoIncrement"`
	Student_id int       `gorm:"type:int(11)" json:"studentid"`
	Advisor    string    `gorm:"not null;type:varchar(115);default:''" json:"advisor"`
	Countries  string    `gorm:"not null;type:varchar(115);default:''" json:"countries"`
	Language   string    `gorm:"not null;type:varchar(115);default:''" json:"language"`
	Departmant string    `gorm:"not null;type:varchar(115);default:''" json:"departmant"`
	CreatedAt  time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
	UpdatedAt  time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
}
