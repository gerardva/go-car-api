package server

import (
	"log"
	"net"

	"github.com/gerardva/go-api/controllers"
	pb "github.com/gerardva/go-api/grpc/protos/carserver"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func Init() {
	initGin()
	initGrpc()
}

func initGrpc() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCarManagerServer(s, &grpcserver.carServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initGin() {
	r := gin.Default()
	health := new(controllers.HealthController)
	r.GET("/health", health.Check)

	carGroup := r.Group("car")
	{
		car := new(controllers.CarController)
		carGroup.GET("/", car.GetAllCars)
		carGroup.POST("/", car.CreateCar)
		carGroup.GET("/:id", car.GetCarById)
		carGroup.DELETE("/:id", car.DeleteCar)
		carGroup.PATCH("/:id", car.UpdateCar)
	}

	r.Run()
}
