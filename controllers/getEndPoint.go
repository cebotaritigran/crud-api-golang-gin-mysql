package controllers

import (
	"crud-api/models"
	"crud-api/types"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// tanımlamalar endpoint
func Users(c *gin.Context) {
	// this code is to be replaced later
	pageString := c.Param("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		fmt.Print(err)
	}
	offset := (page - 1) * 50

	/*
		This query will send only these 3 columns any other column will be ignored instead what db.select does, which sends
		all the columns and leaves every not selected column empty which wasn't ideal. The only downside is with db.raw
		is that we will have to create a lot of result structs to scan them.
	*/
	var result []types.ResultUsers
	query := `SELECT first_name, 
		last_name, 
		authlevel 
		FROM users`
	models.DB.Raw(query).Limit(50).Offset(offset).Scan(&result)
	/*
		This will return a struct with all fields like, id, first_name, last_name, password and so on. Even tho those fields are empty
		I tought it wasn't a good idea to send them to client anyway so above code will only send those 3 columns
	*/
	/* var users []models.User
	result := models.DB.Select([]string{"first_name", "last_name", "authlevel"}).Find(&users)
	if result.Error != nil {
	c.JSON(http.StatusBadRequest, gin.H{"message": "couldn't retrive data"})
	return
	}
	*/
	c.JSON(http.StatusOK, gin.H{"message": "200 ok", "data": result})
}

func Managers(c *gin.Context) {
	pageString := c.Param("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		fmt.Print(err)
	}
	offset := (page - 1) * 50
	var result []types.ResultManagers
	query := `SELECT first_name,
		last_name, 
		phone_number 
		FROM managers`
	models.DB.Raw(query).Limit(50).Offset(offset).Scan(&result)
	c.JSON(http.StatusOK, gin.H{"message": "200 ok", "data": result})
}

func Campuses(c *gin.Context) {
	pageString := c.Param("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		fmt.Print(err)
	}
	offset := (page - 1) * 150
	var result []types.ResultCampuses
	query := `SELECT name,
		city,
		type,
		phone_number
		FROM campuses`

	models.DB.Raw(query).Limit(150).Offset(offset).Scan(&result)
	c.JSON(http.StatusOK, gin.H{"message": "200 ok", "data": result})
}

func Programs(c *gin.Context) {
	pageString := c.Param("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		fmt.Print(err)
	}
	offset := (page - 1) * 150
	var result []types.ResultPrograms
	query := `SELECT name,
		description
		FROM programs`

	models.DB.Raw(query).Limit(150).Offset(offset).Scan(&result)
	c.JSON(http.StatusOK, gin.H{"message": "200 ok", "data": result})
}

func CampusManagers(c *gin.Context) {

	/*
		DEBUGING CODE
		query := `SELECT
		 campuses.name,
		 campuses.type,
		 campuses.city,
		 campuses.license_type,
		 managers.first_name
		 FROM campuses LEFT JOIN managers ON campuses.manager_id = managers.id`
		models.DB.Raw(query).Scan(&result)
	*/
	/*
		DEBUGING CODE
		err := models.DB.Table("campuses").Joins("JOIN managers ON campuses.manager_id = managers.id").Select("campuses.name,campuses.type,campuses.city,campuses.license_type,managers.first_name").Find(&result).Error
		if err != nil {
			fmt.Printf("error")
		}
		fmt.Println(result)
	*/

	pageString := c.Param("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		fmt.Print(err)
	}
	offset := (page - 1) * 150

	var result []types.ResultCampusManagers
	query := `SELECT c.name,
		c.type,
		c.city,
		c.license_type,
		CONCAT(m.first_name,' ',m.last_name) AS manager 
		FROM campuses c LEFT JOIN managers m ON c.manager_id = m.id`

	models.DB.Raw(query).Limit(150).Offset(offset).Scan(&result)

	c.JSON(http.StatusOK, gin.H{"message": "200 ok", "data": result})
}

func AreaDirectors(c *gin.Context) {
	pageString := c.Param("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		fmt.Print(err)
	}
	offset := (page - 1) * 150
	var result []types.ResultAreaDirectors
	query := `SELECT first_name,
		last_name,
		phone_number
		FROM area_directors`

	models.DB.Raw(query).Limit(150).Offset(offset).Scan(&result)

	c.JSON(http.StatusOK, gin.H{"message": "200 ok", "data": result})
}

func DepartmentHeads(c *gin.Context) {
	pageString := c.Param("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		fmt.Print(err)
	}
	offset := (page - 1) * 150
	var result []types.ResultHeadDepartments
	query := `SELECT first_name,
		last_name,
		phone_number
		FROM department_heads`

	models.DB.Raw(query).Limit(150).Offset(offset).Scan(&result)

	c.JSON(http.StatusOK, gin.H{"message": "200 ok", "data": result})
}

func WebsiteBanners(c *gin.Context) {
	pageString := c.Param("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		fmt.Print(err)
	}
	offset := (page - 1) * 150
	var result []types.ResultWebsiteBanners
	query := `SELECT banner_name,
		banner_url,
		row
		FROM website_banners`

	models.DB.Raw(query).Limit(150).Offset(offset).Scan(&result)

	c.JSON(http.StatusOK, gin.H{"message": "200 ok", "data": result})
}

// kasa endpoint
func CampusPayments(c *gin.Context) {
	pageString := c.Param("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		fmt.Print(err)
	}
	offset := (page - 1) * 150
	var result []types.ResultCampusAccountDetails
	// credit
	var resultCredit []types.ResultCampusAccountDetailsCredit
	queryCredit := `SELECT SUM(CASE
		WHEN document_type = 'Alacak' AND currency = 'Euro' THEN sum END)
		euro,
		SUM(CASE
		WHEN document_type = 'Alacak' AND currency = 'Dolar' THEN sum END)
		dolar,
		SUM(CASE
		WHEN document_type = 'Alacak' AND currency = 'Tl' THEN sum END)
		tl,
		SUM(CASE
		WHEN document_type = 'Alacak' AND currency = 'Gbp' THEN sum END)
		gbp
		FROM campus_payments`
	models.DB.Raw(queryCredit).Scan(&resultCredit)
	// debt
	var resultDebt []types.ResultCampusAccountDetailsDebt
	queryDebt := `SELECT SUM(CASE
		WHEN document_type = 'Borc' AND currency = 'Euro' THEN sum END)
		euro,
		SUM(CASE
		WHEN document_type = 'Borc' AND currency = 'Dolar' THEN sum END)
		dolar,
		SUM(CASE
		WHEN document_type = 'Borc' AND currency = 'Tl' THEN sum END)
		tl,
		SUM(CASE
		WHEN document_type = 'Borc' AND currency = 'Gbp' THEN sum END)
		gbp
		FROM campus_payments`
	models.DB.Raw(queryDebt).Scan(&resultDebt)

	query := `SELECT a.date_issued,
		c.type,
		c.name,
		p.document_type,
		p.process_type,
		p.sum,
		p.currency
		FROM campuses_and_payments a INNER JOIN campuses c ON a.campus_id = c.id
		INNER JOIN campus_payments p ON a.campus_payments_id = p.id
		INNER JOIN tours t ON a.tour_id = t.id `
	models.DB.Raw(query).Limit(150).Offset(offset).Scan(&result)
	c.JSON(http.StatusOK, gin.H{"message": "200 ok", "credit": resultCredit, "debt": resultDebt, "data": result})
}

func StudentPayments(c *gin.Context) {
	pageString := c.Param("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		fmt.Print(err)
	}
	offset := (page - 1) * 150
	var result []types.ResultStudentAccountDetails
	query := `SELECT t.datetour,
		t.tourname,
		CONCAT(s.first_name,' ',s.last_name) AS student,
		p.process_type,
		p.account,
		p.currency,
		p.payment_date,
		p.sum,
		p.sumtl,
		p.rate
		FROM students_and_payments a INNER JOIN students s
		ON a.student_id = s.id
		INNER JOIN student_payments p
		ON a.student_payments_id = p.id
		INNER JOIN tours t
		ON a.tour_id = t.id WHERE t.tourname = ?`
	tourName := "test"
	models.DB.Raw(query, tourName).Limit(150).Offset(offset).Scan(&result)
	c.JSON(http.StatusOK, gin.H{"message": "200 ok", "data": result})
}

func Tours(c *gin.Context) {
	pageString := c.Param("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		fmt.Print(err)
	}
	offset := (page - 1) * 150
	var result []types.ResultTours
	query := `SELECT tour_year,
		tour_name FROM tours`
	models.DB.Raw(query).Limit(150).Offset(offset).Scan(&result)
	c.JSON(http.StatusOK, gin.H{"message": "200 ok", "data": result})
}

func TourExpenses(c *gin.Context) {
	pageString := c.Param("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		fmt.Print(err)
	}
	offset := (page - 1) * 150
	var result []types.ResultTourCosts
	query := `SELECT t.tour_name,
		e.currency,
		e.sum,
		e.sumtl,
		e.cost_type,
		bill FROM tours_and_expenses a
		INNER JOIN tours t ON a.tour_id = t.id
		INNER JOIN tour_expenses e ON A.tour_expenses_id = e.id`
	//test
	tourId := 1
	models.DB.Raw(query, tourId).Limit(150).Offset(offset).Scan(&result)
	c.JSON(http.StatusOK, gin.H{"message": "200 ok", "data": result})
}

// 'Raporlar' enpoints
func CampusTours(c *gin.Context) {
	pageString := c.Param("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		fmt.Print(err)
	}
	offset := (page - 1) * 150
	var result []types.ResultCampusTours
	// testing
	tourId := 1
	tourDate := "2020-02-02"
	campuseId := 1
	query := `SELECT r.date_start,
		r.date_end,
		c.license_type,
		c.name,
		c.phone_number,
		r.quota,
		w.quota_request,
		CONCAT(m.first_name,' ',m.last_name) AS manager,
		m.phone_number
		FROM campus_tours t 
		INNER JOIN campuses c
		ON t.campus_id = c.id 
		INNER JOIN campus_tour_request w
		ON t.campus_tour_request_id = w.id 
		INNER JOIN managers m
		ON c.manager_id = m.id 
		INNER JOIN tours r
		ON t.tour_id = r.id WHERE t.tour_id = ? AND r.date_start = ? `
	models.DB.Raw(query, tourId, tourDate).Limit(150).Offset(offset).Scan(&result)

	var resultResponsible []types.ResultCampusToursResponsible
	queryResponsible := `SELECT r.name,
		r.phone_number 
		FROM tour_campus_manager_responsibles t 
		INNER JOIN responsibles r
		ON t.responsible_id = r.id
		INNER JOIN campuses c
		ON t.campus_id = c.id
		INNER JOIN tours o
		ON t.tour_id = o.id
		WHERE t.tour_id = ? AND t.campus_id = ?`
	models.DB.Raw(queryResponsible, tourId, campuseId).Limit(150).Offset(offset).Scan(&resultResponsible)

	var resultStudents []types.ResultCampusToursStundents
	qeuryStudents := `SELECT s.first_name,
		s.last_name,
		s.security_number,
		s.birth_date,
		s.gender,
		s.passport,
		s.phone_number 
		FROM tour_track_records r INNER JOIN students s
		ON r.studnet_id = s.id 
		INNER JOIN tours t
		ON r.tour_id = t.id 
		INNER JOIN campuses c
		ON r.campus_id = c.id
		WHERE r.tour_id = ? AND r.campus_id = ? AND canceled = 0`
	models.DB.Raw(qeuryStudents, tourId, campuseId).Limit(150).Offset(offset).Scan(&resultStudents)

	c.JSON(http.StatusOK, gin.H{"message": "200 ok", "data": result, "responsible": resultResponsible, "students": resultStudents})
}

func CampusTourRequest(c *gin.Context) {
	pageString := c.Param("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		fmt.Print(err)
	}
	offset := (page - 1) * 150
	var result []types.ResultCampusTourWanted
	//testing
	tourDate := "2020-02-02"
	campuseType := "istanbul"
	campuseName := "istanbul"
	query := `SELECT t.tourname,
		w.quota_request,
		w.attendance
		FROM campus_tour_requests w
		INNER JOIN tours t
		ON w.tour_id = t.id
		INNER JOIN campuses c
		ON w.campus_id = c.id
		WHERE t.date_start = ? AND c.type = ? AND c.name = ?`
	models.DB.Raw(query, tourDate, campuseType, campuseName).Limit(150).Offset(offset).Scan(&result)
	c.JSON(http.StatusOK, gin.H{"message": "200 ok", "data": result})
}

func TourTrackRecord(c *gin.Context) {
	pageString := c.Param("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		fmt.Print(err)
	}
	offset := (page - 1) * 150
	var result []types.ResultRecordTrack
	tourDate := "2020-02-02"
	tourName := "test"
	query := `SELECT s.first_name,
		s.last_name,
		c.type,
		c.name,
		p.name AS name_responsible,
		s.security_number,
		s.birth_date,
		s.gender,
		s.passport,
		y.currency,
		y.sum,
		y.sumtl
		FROM tour_track_records r INNER JOIN students s
		ON r.student_id = s.id
		INNER JOIN campuses c
		ON r.campus_id = c.id
		INNER JOIN students_and_payments a
		ON r.students_and_payments_id = a.id
		INNER JOIN students n
		ON a.student_id = n.id
		INNER JOIN student_payments y
		ON a.student_payments = y.id
		INNER JOIN tours t
		ON r.tour_id = t.id
		INNER JOIN tour_campus_manager_responsibles b
		ON r.tour_campus_responsible_id = b.id
		INNER JOIN responsibles p
		ON b.responsible_id = p.id
		WHERE t.date_start = ? AND 
		t.tour_name = ?`
	models.DB.Raw(query, tourDate, tourName).Limit(150).Offset(offset).Scan(&result)
	c.JSON(http.StatusOK, gin.H{"message": "200 ok", "record track": result})
}

func TourRooming(c *gin.Context) {
	pageString := c.Param("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		fmt.Print(err)
	}
	offset := (page - 1) * 150
	var result []types.ResultTourRooming
	// testing
	tourDate := "2020-02-02"
	tourName := "test"
	query := `SELECT s.first_name,
		s.last_name,
		s.security_number,
		s.birth_date,
		s.gender,
		s.passport,
		c.name,
		p.name AS responsible_name,
		r.room_number
		FROM tour_roomings r INNER JOIN students s
		ON r.student_id = s.id
		INNER JOIN campuses c
		ON r.campus_id = c.id
		INNER JOIN tours t
		ON r.tour_id = t.id
		INNER JOIN tour_campus_manager_responsibles b
		ON r.tour_campus_responsible_id = b.id
		INNER JOIN responsibles p
		ON b.responsible_id = p.id
		WHERE t.date_start = ? AND 
		t.tour_name = ?`
	models.DB.Raw(query, tourDate, tourName).Limit(150).Offset(offset).Scan(&result)
	c.JSON(http.StatusOK, gin.H{"message": "200 ok", "rooming": result})
}

func CampusTourManagerResponsible(c *gin.Context) {
	pageString := c.Param("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		fmt.Print(err)
	}
	offset := (page - 1) * 150
	var result []types.ResultCTMR
	query := `SELECT t.tour_name,
		t.date_start,
		c.type,
		c.name AS campus_name,
		c.city,
		CONCAT(m.first_name,' ',m.last_name) AS manager,
		m.email AS manager_email,
		m.phone_number AS manager_phone_number,
		p.name AS responsible_name,
		p.email AS responsible_email,
		p.phone_number AS responsible_phone,
		status
		FROM tour_campus_manager_responsibles r
		INNER JOIN tours t
		ON r.tour_id = t.id
		INNER JOIN campuses c
		ON r.campus_id = c.id
		INNER JOIN responsibles p
		ON r.responsible_id = p.id
		INNER JOIN managers m
		ON c.manager_id = m.id;`
	models.DB.Raw(query).Limit(150).Offset(offset).Scan(&result)
	c.JSON(http.StatusOK, gin.H{"message": "200 ok", "CTMR": result})
}

func AllStudentRecords(c *gin.Context) {
	pageString := c.Param("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		fmt.Print(err)
	}
	offset := (page - 1) * 150
	var result []types.ResultStudentRecords
	query := `SELECT t.abroad,
		t.tour_name,
		t.date_start,
		d.first_name,
		d.last_name,
		p.name,
		d.security_number,
		d.birth_date,
		d.gender,
		d.passport,
		i.sum
		FROM tour_track_records r
		INNER JOIN students d
		ON r.student_id = d.id
		INNER JOIN students_and_payments i
		ON d.students_and_payments = i.id
		INNER JOIN students n
		ON i.student_id = n.id
		INNER JOIN student_payments m
		ON i.student_payments = m.id
		INNER JOIN tours t
		ON r.tour_id = t.id
		INNER JOIN campuses c
		ON r.campus_id = c.id
		INNER JOIN tour_campus_manager_responsibles s
		ON r.tour_campus_responsible_id = s.id
		INNER JOIN responsibles p
		ON s.responsible_id = p.id;`
	models.DB.Raw(query).Limit(150).Offset(offset).Scan(&result)
	c.JSON(http.StatusOK, gin.H{"message": "200 ok", "CTMR": result})
}

// * * * * * * *needs to be checked more than other queries* * * * * * * * *
func AllToursSummary(c *gin.Context) {
	pageString := c.Param("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		fmt.Print(err)
	}
	offset := (page - 1) * 150
	var result []types.ResultAllToursSummary
	query := `SELECT t.date_start,
		t.tour_name,
		SUM(cw.quota_request) AS quota_request,
		t.quota,
		t.tour_cost,
		quota_request * t.tour_cost AS total
		FROM campus_tours ct
		INNER JOIN campuses c
		ON ct.campus_id = c.id
		INNER JOIN tours t
		ON ct.tour_id = t.id
		INNER JOIN campus_tour_requests cw
		ON ct.campus_tour_request_id = cw.id
		GROUP BY ct.id`
	models.DB.Raw(query).Limit(150).Offset(offset).Scan(&result)
	c.JSON(http.StatusOK, gin.H{"message": "200 ok", "Summary": result})
}

func LogStudents(c *gin.Context) {
	pageString := c.Param("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		fmt.Print(err)
	}
	offset := (page - 1) * 150
	var result []types.LogStudents
	query := `SELECT c.name,
		s.first_name,
		s.last_name,
		s.security_number,
		s.birth_date,
		s.address,
		s.phone_number,
		s.email,
		p.name AS parent_name,
		p.phone_number AS parent_phone,
		p.email AS parent_email,
		s.status
		FROM students s
		INNER JOIN campuses c
		ON s.campus_id = c.id
		INNER JOIN parents p
		ON s.parent_id = p.id`
	models.DB.Raw(query).Limit(150).Offset(offset).Scan(&result)
	c.JSON(http.StatusOK, gin.H{"message": "200 ok", "log students": result})
}

func RecordTrackReport(c *gin.Context) {
	pageString := c.Param("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		fmt.Print(err)
	}
	offset := (page - 1) * 150
	var result []types.RecordTrackReport
	query := `SELECT s.first_name,
		s.last_name,
		c.name AS campus_name,
		r.advisor,
		r.country,
		s.departmant,
		r.language,
		p.name AS parent_name,
		p.phone_number AS parent_phone
		FROM record_track_reports r
		INNER JOIN students s
		ON r.student_id = s.id
		INNER JOIN parents p
		ON s.parent_id = p.id
		INNER JOIN campuses c
		ON s.campus_id = c.id`
	models.DB.Raw(query).Limit(150).Offset(offset).Scan(&result)
	c.JSON(http.StatusOK, gin.H{"message": "200 ok", "report": result})
}
