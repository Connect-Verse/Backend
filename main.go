package main

import (
	"log"
	"net/http"
	"time"

	"github.com/connect-verse/internal/api"
	"github.com/connect-verse/internal/models"
	"github.com/connect-verse/internal/repository/avatars"
	"github.com/connect-verse/internal/repository/maps"
	"github.com/connect-verse/internal/repository/metaUser"
	"github.com/connect-verse/internal/repository/rooms"
	"github.com/connect-verse/internal/repository/user"
	"github.com/connect-verse/internal/repository/verificationToken"
	"github.com/connect-verse/internal/services"
	"github.com/connect-verse/internal/services/auth-service"
	"github.com/connect-verse/internal/services/avatar-service"
	"github.com/connect-verse/internal/services/map-service"
	"github.com/connect-verse/internal/services/meta-users"
	"github.com/connect-verse/internal/services/room-service"
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

	 db.AutoMigrate(&models.User{}, &models.VerificationToken{}, &models.Avatars{}, &models.Maps{},&models.Rooms{},&models.MetaUsers{},&models.PlayerPosition{})


	 userRespository := user.NewUserImplementation(db)
     VerifyRepository := verificationtoken.NewVerifyImplementation(db)
	 mapRepository := maps.NewMapsRepoImpl(db)
	 avatarRepository := avatars.NewAvatarRepoImpl(db)
	 roomRepository := rooms.NewRoomImplementation(db)
	 metaRepository := metauser.NewMetaRepoImpl(db)
	 if metaRepository==nil && roomRepository==nil && avatarRepository==nil && userRespository==nil && VerifyRepository==nil && mapRepository==nil{
		log.Fatal("error occured while initialising the userRepository")
        return
	 }
     //services
	 service,err := services.NewUserServiceImp(userRespository,validate)
	 avatarservice,err := avatarservice.NewAvatarServImpl(avatarRepository,validate)
     authService,err := authservice.NewAuthServiceImplementation(userRespository,VerifyRepository,validate)
	 mapService,err := mapservice.NewMapServiceImpl(mapRepository,validate)
	 roomService,err := roomservice.NewRoomServiceImpl(roomRepository,validate)
	 metaService, err := metaservice.NewMetaServiceImpl(metaRepository,validate)
	 utils.PanicError(err)
	 
	 //controllers
	 userController := handlers.NewControllerService(metaService,roomService,avatarservice,service,authService,mapService)
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

	roomsRouter:= serve.Group("/rooms")
	roomsRouter.POST("/create-room",controller.CreateRoom)
	roomsRouter.DELETE("/delete-room",controller.DeleteRoom)
	roomsRouter.GET("/users-room",controller.MyRoom)
	roomsRouter.GET("/all-rooms",controller.FindAllRooms)


	avatarRouter:= serve.Group("/avatar")
	avatarRouter.POST("/create-avatar",controller.CreateAvatar)
	avatarRouter.DELETE("/delete-avatar",controller.DeleteAvatar)
	avatarRouter.GET("/Update-avatar",controller.UpdateAvatar)
	avatarRouter.GET("/all-avatar",controller.FindAllAvatar)
	avatarRouter.GET("/find-avatar",controller.FindAvatar)


	metaUserRouter:= serve.Group("/metaUser")
	metaUserRouter.POST("/create-metaUser",controller.CreateMeta)
	metaUserRouter.DELETE("/delete-metaUser",controller.DeleteMetaUser)
	metaUserRouter.GET("/find-metaUser",controller.FindById)

	// playerPosRouter:= serve.Group("/player-position")
	// playerPosRouter.POST("/create-metaUser",controller.CreateMetaUser)
	// playerPosRouter.DELETE("/delete-metaUser",controller.DeleteMetaUser)
	// playerPosRouter.GET("/rooms-metaUser",controller.GetMetaUser)
	// playerPosRouter.GET("/all-metaUser",controller.AllMetaUser)

	return serve
}