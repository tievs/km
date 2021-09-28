package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tievs/km/controllers"
	"github.com/tievs/km/db"
	"net"
	"strings"
)

func init() {
	db.Init()
}

func main()  {
	//ip:=GetOutboundIP()
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORS())

	e.Static("/", "public")

	e.GET("/doctor", controllers.GetDoctorFiles)
	e.POST("/doctor", controllers.PostDoctorFiles)
	e.PUT("/doctor/:id", controllers.UpdateDoctorFiles)
	e.PUT("/doctor/:id/content", controllers.UpdateDoctorContent)
	e.DELETE("/doctor/:id", controllers.DeleteDoctorFiles)
	e.GET("/doctor/:id", controllers.GetDoctorFile)

	e.GET("/patient", controllers.GetPatientFiles)
	e.POST("/patient", controllers.PostPatientFiles)
	e.PUT("/patient/:id", controllers.UpdatePatientFiles)
	e.PUT("/patient/:id/content", controllers.UpdatePatientContent)
	e.DELETE("/patient/:id", controllers.DeletePatientFiles)
	e.GET("/patient/:id", controllers.GetPatientFile)

	e.GET("/drug", controllers.GetDrugFiles)
	e.POST("/drug", controllers.PostDrugFiles)
	e.PUT("/drug/:id", controllers.UpdateDrugFiles)
	e.PUT("/drug/:id/content", controllers.UpdateDrugContent)
	e.DELETE("/drug/:id", controllers.DeleteDrugFiles)
	e.GET("/drug/:id", controllers.GetDrugFile)

	e.GET("/nurse", controllers.GetNurseFiles)
	e.POST("/nurse", controllers.PostNurseFiles)
	e.PUT("/nurse/:id", controllers.UpdateNurseFiles)
	e.PUT("/nurse/:id/content", controllers.UpdateNurseContent)
	e.DELETE("/nurse/:id", controllers.DeleteNurseFiles)
	e.GET("/nurse/:id", controllers.GetNurseFile)

	//e.Logger.Fatal(e.Start(ip+":3001"))
	e.Logger.Fatal(e.Start("localhost:3001"))
}

// GetOutboundIP Get preferred outbound ip of this machine
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "localhost"
	}
	defer conn.Close()
	localAddr := strings.Split(conn.LocalAddr().String(),":")
	return localAddr[0]
}