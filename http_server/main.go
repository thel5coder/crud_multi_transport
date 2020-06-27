package main

import (
	"crud_multi_transport/db"
	"crud_multi_transport/helpers/jwe"
	"crud_multi_transport/helpers/jwt"
	"crud_multi_transport/helpers/str"
	"crud_multi_transport/http_server/bootstrap"
	"crud_multi_transport/usecase"
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
	"github.com/go-redis/redis/v7"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"os"
)

var (
	validatorDriver *validator.Validate
	uni             *ut.UniversalTranslator
	translator      ut.Translator
)

func main(){
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading ..env file")
	}

	//jwe

	jweCredential := jwe.Credential{
		KeyLocation: os.Getenv("PRIVATE_KEY"),
		Passphrase:  os.Getenv("PASSPHRASE"),
	}

	//setup redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	pong, err := redisClient.Ping().Result()
	fmt.Println("Redis ping status: "+pong, err)

	//setup db connection
	dbInfo := db.Connection{
		Host:     os.Getenv("DB_HOST"),
		DbName:   os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		SslMode:  os.Getenv("DB_SSL_MODE"),
	}

	database, err := dbInfo.DbConnect()
	if err != nil {
		panic(err)
	}

	//jwtconfig
	jwtConfig := middleware.JWTConfig{
		Claims:     &jwt.CustomClaims{},
		SigningKey: []byte(os.Getenv("SECRET")),
	}

	//jwt credential
	jwtCred := jwt.JwtCredential{
		TokenSecret:         os.Getenv("SECRET"),
		ExpiredToken:        str.StringToInt(os.Getenv("TOKEN_EXP_TIME")),
		RefreshTokenSecret:  os.Getenv("SECRET_REFRESH_TOKEN"),
		ExpiredRefreshToken: str.StringToInt(os.Getenv("REFRESH_TOKEN_EXP_TIME")),
	}

	//init validator
	validatorInit()

	//echo instance
	e := echo.New()
	useCaseContract := usecase.UcContract{
		DB:        database,
		Redis:     redisClient,
		Jwe:       jweCredential,
		JwtConfig: jwtConfig,
		JwtCred:   jwtCred,
	}

	bootApp := bootstrap.Bootstrap{
		E:               e,
		UseCaseContract: useCaseContract,
		Validator:       validatorDriver,
		Translator:      translator,
	}

	bootApp.E.Use(middleware.Logger())
	bootApp.E.Use(middleware.Recover())
	bootApp.E.Use(middleware.CORS())
	bootApp.RegisterRouters()

	bootApp.E.Logger.Fatal(bootApp.E.Start(os.Getenv("APP_HOST_SERVER")))
}

func validatorInit() {
	en := en.New()
	id := id.New()
	uni = ut.New(en, id)

	transEN, _ := uni.GetTranslator("en")
	transID, _ := uni.GetTranslator("id")

	validatorDriver = validator.New()

	enTranslations.RegisterDefaultTranslations(validatorDriver, transEN)
	idTranslations.RegisterDefaultTranslations(validatorDriver, transID)

	switch os.Getenv("APP_LOCALE") {
	case "en":
		translator = transEN
	case "id":
		translator = transID
	}
}