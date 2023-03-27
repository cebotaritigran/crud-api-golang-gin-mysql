package routes

import (
	"crud-api/controllers"
	"crud-api/middlewares"

	"github.com/gin-gonic/gin"
)

// middlewared routes by jwt
func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middlewares.JwtAuthMiddleware())
	i := incomingRoutes
	// 'tanimlamalar' api endpoints maybe could be grouped to 'tanimlamalar'
	i.GET("/user", controllers.CurrentUser)
	i.GET("/users", controllers.Users)
	// pagination ot this api endpoint
	i.GET("/users/:page", controllers.Users)
	i.GET("/managers", controllers.Managers)
	i.GET("/managers:page", controllers.Managers)
	i.GET("/campuses", controllers.Campuses)
	i.GET("/campuses:page", controllers.Campuses)
	i.GET("/programs", controllers.Programs)
	i.GET("/programs:page", controllers.Programs)
	i.GET("/campuse-managers", controllers.CampusManagers)
	i.GET("/campuse-managers:page", controllers.CampusManagers)
	i.GET("/department-heads", controllers.DepartmentHeads)
	i.GET("/department-heads:page", controllers.DepartmentHeads)
	i.GET("/area-directors", controllers.AreaDirectors)
	i.GET("/area-directors:page", controllers.AreaDirectors)
	i.GET("/website-banners", controllers.WebsiteBanners)
	i.GET("/website-banners:page", controllers.WebsiteBanners)
	//'Kasa' api endpoints
	i.GET("/campuse-account", controllers.CampusPayments)
	i.GET("/campuse-account:page", controllers.CampusPayments)
	i.GET("/student-account", controllers.StudentPayments)
	i.GET("/student-account:page", controllers.StudentPayments)
	i.GET("/tours", controllers.Tours)
	i.GET("/tours:page", controllers.Tours)
	i.GET("/tour-costs", controllers.TourExpenses)
	i.GET("/tour-costs:page", controllers.TourExpenses)
	//'Raporlar' api endpoints
	i.GET("/campuse-tours", controllers.CampusTours)
	i.GET("/campuse-tours:page", controllers.CampusTours)
	i.GET("/campuse-tours-wanted", controllers.CampusTourRequest)
	i.GET("/campuse-tours-wanted:page", controllers.CampusTourRequest)
	i.GET("/tour-record-track", controllers.TourTrackRecord)
	i.GET("/tour-record-track:page", controllers.TourTrackRecord)
	i.GET("/tour-rooming", controllers.TourRooming)
	i.GET("/tour-rooming:page", controllers.TourRooming)
	i.GET("/campuse-tour-manager-responsible", controllers.CampusTourManagerResponsible)
	i.GET("/campuse-tour-manager-responsible:page", controllers.CampusTourManagerResponsible)
	i.GET("/student-records", controllers.AllStudentRecords)
	i.GET("/student-records:page", controllers.AllStudentRecords)
	i.GET("/tour-summary", controllers.AllToursSummary)
	i.GET("/tour-summary:page", controllers.AllToursSummary)
	i.GET("/log-students", controllers.LogStudents)
	i.GET("/log-students:page", controllers.LogStudents)
	i.GET("/record-track-report", controllers.RecordTrackReport)
	i.GET("/record-track-report:page", controllers.RecordTrackReport)
	// post test
	i.POST("/manager-insert", controllers.PostManagers)
}
