package server

import (
	_ "gym/docs"
	"gym/server/handler"

	"gym/server/provider"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes(server *Server) {

	//Membership routes
	server.engine.GET("/getMembership", provider.AdminAuthorization, handler.GetMembershipHandler)
	server.engine.POST("/createMembership", provider.AdminAuthorization, handler.CreateMembershipHandler)
	server.engine.PUT("/updateMembership", provider.AdminAuthorization, handler.UpdateMembershipHandler)
	server.engine.PUT("/deleteMembership", provider.AdminAuthorization, handler.DeleteMembershipHandler)

	//Equipments routes
	server.engine.POST("/createEquipment", provider.AdminAuthorization, handler.CreateEquipmentHandler)
	server.engine.GET("/getEquipments", provider.AdminAuthorization, handler.GetEquipmentHandler)

	//employee routes
	server.engine.GET("/createEmp", provider.AdminAuthorization, handler.CreateEmployeeHandler)
	server.engine.GET("/getEmp", provider.AdminAuthorization, handler.GetEmployeeHandler)
	server.engine.GET("/getEmpRole", provider.AdminAuthorization, handler.GetEmployeeRoleHandler)
	server.engine.POST("/createEmpRole", provider.AdminAuthorization, handler.EmployeeRoleHandler)
	server.engine.GET("/empWithuser", provider.AdminAuthorization, handler.GetUsersWithEmployeesHandler)
	server.engine.POST("/empAttendence", provider.AdminAuthorization, handler.EmployeeAttendenceHandler)

	//users
	server.engine.POST("/createUser", handler.CreateUserHandler)
	server.engine.POST("/getUser", provider.AdminAuthorization, handler.GetUserByIdHandler)
	server.engine.POST("/userAttendence", provider.AdminAuthorization, handler.UserAttendenceHandler)
	server.engine.GET("/getUsers", provider.AdminAuthorization, handler.GetAllUsers)

	//slots
	server.engine.POST("/slotUpdate", handler.SlotUpdateHandler)

	//subscriptions
	server.engine.POST("/createSubscription", provider.UserAuthorization, handler.CreateSubscriptionHandler)
	server.engine.DELETE("/endSubscription", provider.AdminAuthorization, handler.EndSubscriptionHandler)
	server.engine.POST("/updateSubscription", provider.AdminAuthorization, handler.UpdateMembershipHandler)

	//Payment Routes
	server.engine.POST("/createPayment", handler.MakePaymentHandler)
	server.engine.POST("/paymentResponse",handler.PaymentResponse)


	//Auth routes
	server.engine.POST("/adminRegister", handler.AdminRegisterHandler)
	server.engine.POST("/sendOtp", handler.SendOtpHandler)
	server.engine.POST("/verifyOtp", handler.VerifyOtpHandler)
	server.engine.GET("/logout", handler.LogoutHandler)

	server.engine.GET("/sad", handler.GetCookie)

	//swagger route
	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
