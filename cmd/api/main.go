package main

import (
	"individual-project-hacktiv8-p2/internal/delivery/http/handler"
	"individual-project-hacktiv8-p2/internal/delivery/http/middleware"
	"individual-project-hacktiv8-p2/internal/infrastructure/database"
	"individual-project-hacktiv8-p2/internal/repository/db"
	externalapi "individual-project-hacktiv8-p2/internal/repository/externalApi"
	"individual-project-hacktiv8-p2/internal/usecase/book"
	"individual-project-hacktiv8-p2/internal/usecase/paymentRecord"
	"individual-project-hacktiv8-p2/internal/usecase/rentBookUseCase"
	"individual-project-hacktiv8-p2/internal/usecase/user"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"

	_ "github.com/lib/pq"
	// "p2-gc03-Billyhaffas-3/internal/delivery/http/middleware"
)

func main() {
	// // connectDB
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, skipping", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port for local development
	}

	sqlDbConn, dbGormConn, err := database.ConnectPostgres()

	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	defer sqlDbConn.Close()

	//init repository
	userRepository := db.NewAuthDBconnection(dbGormConn)
	paymentRepository := db.NewPaymentRecordDBConnection(dbGormConn)
	bookRepository := db.NewBookDBconnection(dbGormConn)
	rentBookRepository := db.NewRentBookDBconnection(dbGormConn)

	// init repository (External API)
	httpClient := &http.Client{Timeout: 10 * time.Second}
	googleRepo := externalapi.NewGoogleBooksRepo(httpClient)

	//init usecase
	userUseCase := user.AuthUseCase(userRepository)
	paymentUseCase := paymentRecord.NewPaymentRecordUseCase(paymentRepository, userRepository)
	bookUseCase := book.NewbookUseCase(bookRepository, googleRepo)
	rentBookUseCase := rentBookUseCase.NewRentBookUseCase(*&dbGormConn, userRepository, rentBookRepository, paymentRepository, bookRepository)

	//init handler
	userHandler := handler.AuthHandler(userUseCase)
	paymentHandler := handler.PaymentHandler(paymentUseCase)
	bookHandler := handler.NewbookHandler(bookUseCase)
	rentBookHandler := handler.NewrentBookHandler(rentBookUseCase)

	echo := echo.New()
	api := echo.Group("/api")
	api.Use(middleware.AuthMiddleware())

	echo.POST("/users/register", userHandler.Register)
	echo.POST("/users/login", userHandler.Login)
	api.GET("/users", userHandler.GetMe)
	api.POST("/users/top-up", paymentHandler.PostPaymentRecord)
	api.POST("/users/rent-book", rentBookHandler.PostRentBook)
	api.GET("/users/rent-book-history", rentBookHandler.GetRentBook)
	api.GET("/users/books/:book_name", bookHandler.GetbookByName)

	log.Println("Server running on port", port)
	if err := echo.Start(":" + port); err != nil {
		log.Fatal(err)
	}

}
