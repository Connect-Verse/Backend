package main

import (
	"log"
	"net/http"
	"time"

	"github.com/connect-verse/internal/api"
	"github.com/connect-verse/internal/models"
	"github.com/connect-verse/internal/repository/maps"
	"github.com/connect-verse/internal/repository/user"
	"github.com/connect-verse/internal/repository/verificationToken"
	"github.com/connect-verse/internal/services"
	"github.com/connect-verse/internal/services/auth-service"
    "github.com/connect-verse/internal/services/map-service"
	"github.com/connect-verse/internal/utils"
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
	 mapRepository := maps.NewMapsRepoImpl(db)
	 if userRespository==nil && VerifyRepository==nil && mapRepository==nil{
		log.Fatal("error occured while initialising the userRepository")
        return
	 }

	 service,err := services.NewUserServiceImp(userRespository,validate)
	 utils.PanicError(err)
     authService,err := authservice.NewAuthServiceImplementation(userRespository,VerifyRepository,validate)
	 utils.PanicError(err)
	 mapService,err := mapservice.NewMapServiceImpl(mapRepository,validate)
	 utils.PanicError(err)

	 userController := handlers.NewControllerService(service,authService,mapService)
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
	authRouter.POST("/verify-User",controller.Verify)

	mapsRouter:= serve.Group("/maps")
	mapsRouter.POST("/create-map",controller.CreateMap)
	mapsRouter.DELETE("/delete-map",controller.DeleteMap)
	mapsRouter.GET("/all-maps",controller.FindAllMap)
	mapsRouter.GET("/find-map",controller.FindMap)
	mapsRouter.PATCH("/update-map",controller.UpdateMap)

	// roomsRouter:= serve.Group("/rooms")
	// roomsRouter.POST("/create-room",controller.CreateRoom)
	// roomsRouter.DELETE("/delete-room",controller.DeleteRoom)
	// roomsRouter.GET("/users-room",controller.GetUsers)
	// roomsRouter.GET("/all-rooms",controller.AllRooms)


	avatarRouter:= serve.Group("/avatar")
	avatarRouter.POST("/create-avatar",controller.CreateAvatar)
	avatarRouter.DELETE("/delete-avatar",controller.DeleteAvatar)
	avatarRouter.GET("/Update-avatar",controller.UpdateAvatar)
	avatarRouter.GET("/all-avatar",controller.AllAvatar)
	avatarRouter.GET("/find-avatar",controller.FindAvatar)




	// metaUserRouter:= serve.Group("/metaUser")
	// metaUserRouter.POST("/create-metaUser",controller.CreateMetaUser)
	// metaUserRouter.DELETE("/delete-metaUser",controller.DeleteMetaUser)
	// metaUserRouter.GET("/rooms-metaUser",controller.GetMetaUser)
	// metaUserRouter.GET("/all-metaUser",controller.AllMetaUser)

	return serve
}