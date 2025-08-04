package routes
import (
	"github.com/gin-gonic/gin"
    "github.com/Anujc07/real-estate/internal/handler"
    "github.com/Anujc07/real-estate/internal/services"
    "github.com/Anujc07/real-estate/internal/repository"
	"github.com/Anujc07/real-estate/config"
)

func SetupRoutes(router *gin.Engine) {

    // User properties routea
    propertyRepo := repository.NewPropertyRepository(config.DB)
    propertyService := service.NewPropertyService(propertyRepo)
    propertyHandler := handler.NewPropertyHandler(propertyService)

    property := router.Group("/properties")
    {
        property.POST("/", propertyHandler.Add)
        property.GET("/", propertyHandler.GetAll)
        property.GET("/:id", propertyHandler.GetByID)
    }

    // User bookings routes
    bookingRepo := repository.NewBookingRepository(config.DB)
    bookingService := service.NewBookingService(bookingRepo)
    bookingHandler := handler.NewBookingHandler(bookingService)

    booking := router.Group("/bookings")
    {
        booking.POST("/", bookingHandler.BookProperty)
        booking.GET("/", bookingHandler.GetAll)
        booking.GET("/user/:user_id", bookingHandler.GetByUser)
    }

    // User enquiries routes
    enquiryRepo := repository.NewEnquiryRepository(config.DB)
    enquiryService := service.NewEnquiryService(enquiryRepo)
    enquiryHandler := handler.NewEnquiryHandler(enquiryService)
    enquiry := router.Group("/enquiries")
    {
        enquiry.POST("/", enquiryHandler.Create)
        enquiry.GET("/property/:id", enquiryHandler.GetByPropertyID)
    }

    // User reviews routes
    reviewRepo := repository.NewReviewRepository(config.DB)
    reviewService := service.NewReviewService(reviewRepo)
    reviewHandler := handler.NewReviewHandler(reviewService)
    review := router.Group("/reviews")
    {
        review.POST("/", reviewHandler.Create)
        review.GET("/property/:id", reviewHandler.GetByProperty)
    }
}
