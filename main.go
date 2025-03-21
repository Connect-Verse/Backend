package main

import (
	"log"
	"net/http"
	"time"
	"github.com/connect-verse/internal/api"
	"github.com/connect-verse/internal/repository/user"
	"github.com/connect-verse/internal/repository/verificationToken"
	"github.com/connect-verse/internal/services"
	"github.com/connect-verse/internal/services/auth-service"
	"github.com/connect-verse/internal/utils"
	"github.com/connect-verse/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main(){
	 db:= utils.DatabaseConnection()
	 validate := validator.New()

	 if db==nil{
		log.Fatal("error occured while initialising the database")
	 }

	 db.AutoMigrate(&models.User{}, &models.VerificationToken{})


	 userRespository := user.NewUserImplementation(db)
     VerifyRepository := verificationtoken.NewVerifyImplementation(db)
	 if userRespository==nil && VerifyRepository==nil{
		log.Fatal("error occured while initialising the userRepository")
        return
	 }

	 service,err := services.NewUserServiceImp(userRespository,validate)
	 utils.PanicError(err)
     authService,err := authservice.NewAuthServiceImplementation(userRespository,VerifyRepository,validate)
	 utils.PanicError(err)

	 userController := handlers.NewControllerService(service,authService)
	 routes:=userRouter(userController)


	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	   }
	  
	   err = server.ListenAndServe()
	   utils.PanicError(err)
 
}

func userRouter(controller *handlers.Controller) *gin.Engine {
	serve:=gin.Default()
	router:= serve.Group("/user")
	router.POST("/check",controller.Check)
	router.POST("/delete")

	authRouter:= serve.Group("/auth")
	authRouter.POST("/login",controller.Login)
	authRouter.POST("/signUp",controller.Signup)

	return serve
}