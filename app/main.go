package main

import (
	"log"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	_routes "widyaWicaraBackend/app/routes"

	_middleware "widyaWicaraBackend/app/middleware"
	_productUseCase "widyaWicaraBackend/businesses/products"
	_userUseCase "widyaWicaraBackend/businesses/users"
	_productController "widyaWicaraBackend/controllers/products"
	_userController "widyaWicaraBackend/controllers/users"
	_userRepo "widyaWicaraBackend/drivers/databases/users"
	_productRepo "widyaWicaraBackend/drivers/databases/products"
	postgre "widyaWicaraBackend/drivers/postgree"
)

func init() {
	viper.SetConfigFile(`config/config.json`)
	err:= viper.ReadInConfig()
	if err != nil{
		panic(err)
	}
	if viper.GetBool(`debug`){
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB){
	db.AutoMigrate(
		&_userRepo.User{},
		&_productRepo.Product{},

	)
}
func main(){
configDB := postgre.ConfigDB{
	DB_Host:viper.GetString(`database.host`),
	DB_User:viper.GetString(`database.user`),
	DB_Password    :viper.GetString(`database.password`),
	DB_Name   :viper.GetString(`database.name`),
	DB_Port:viper.GetString(`database.port`),
}
db:= configDB.InitialDB()
	dbMigrate(db)

	jwt := _middleware.ConfigJWT{
		SecretJWT : viper.GetString(`jwt.secret`),
		ExpiresDuration : viper.GetInt(`jwt.expired`),
}
timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

 e := echo.New()
e.Use(middleware.CORS())

 userRepo := _userRepo.NewUserRepository(db)
userUseCase := _userUseCase.NewUserUsecase(userRepo,timeoutContext,&jwt)
userController := _userController.NewUserController(userUseCase)

productRepo := _productRepo.NewProductRepository(db)
productUsecase := _productUseCase.NewProductUsecase(productRepo,timeoutContext)
productController := _productController.NewProductController(productUsecase)

routesInit := _routes.ControllerList{
	UserController : *userController,
	ProductController : *productController,
	JWTConfig:		jwt.Init(),
}
routesInit.RouteRegister(e)
log.Fatal(e.Start(viper.GetString("server.address")))
}