package main

import (
	"context"
	pb "elasticsearch/example/grpc"
	"elasticsearch/example/mongo"
	"fmt"
	"log"
	"net"

	// "log"
	// "net"
	"google.golang.org/grpc"
)

type server struct {
	pb.ProductServiceServer
}

func (s *server) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	fmt.Println("Call api create product...", in)
	mongo.CreateProduct(in.Product)
	return &pb.CreateProductResponse{
		Product: in.Product,
	}, nil
}
func (s *server) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {

	fmt.Println("Call api get product...")
	return &pb.GetProductResponse{
		Product: mongo.GetProduct(req.ID),
	}, nil
}
func (s *server) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {

	fmt.Println("Call api update product...")
	return &pb.UpdateProductResponse{
		Product: &pb.Product{
			ID:          "occaecat nisi elit amet",
			Price:       -51385620.69486148,
			Productname: "commodo et Excepteur",
			Quantity:    -990649183,
		},
	}, nil
}
func main() {
	log.Println("Start server...")
	port := ":8080"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, new(server))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
