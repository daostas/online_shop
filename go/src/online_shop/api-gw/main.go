package main

import (
	"fmt"
	"log"
	"online_shop/api-gw/config"
	"online_shop/api-gw/user"
	"os"

	"github.com/iris-contrib/middleware/cors"
	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// @title Quick Shop API
// @version 1.0
// @description Quick shop and related service API documentation
// @termsOfService https://postiv.kz

// @contact.name Michael Studzitsky
// @contact.url https://positiv.kz
// @contact.email info@positiv.kz

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9012
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {

	app := iris.New()
	logf, err := os.OpenFile("./api-gw/log/qs.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		app.Logger().Fatalf("Cant open log file: %v", err)
	}
	app.Logger().SetPrefix("ONLINE_SHOP: ")
	app.Logger().SetLevel("debug")
	app.Logger().SetOutput(logf)
	cfg, err := config.LoadConfig("./api-gw/config")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"HEAD", "GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	})

	//SWAGGER
	swaggerUI := swagger.Handler(swaggerFiles.Handler,
		swagger.URL("localhost:9012/swagger/doc.json"),
		swagger.DeepLinking(true),
		swagger.Prefix("/swagger"),
	)
	app.Get("/swagger", swaggerUI)
	app.Get("/swagger/{any:path}", swaggerUI)

	part := mvc.New(app.Party("/user", c).AllowMethods(iris.MethodOptions))
	user.SetupUser(part, &cfg)

	app.Logger().Println("API_GW on", cfg.Port)
	fmt.Printf("Api-gw started on port %s", cfg.Port)
	err = app.Listen(cfg.Port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nApi-gw stopped")
	app.Logger().Println("API_GW stopped", cfg.Port)
}
