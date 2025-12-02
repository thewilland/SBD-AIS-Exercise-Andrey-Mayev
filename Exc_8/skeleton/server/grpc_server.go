package server

import (
	"context"
	"errors"
	"exc8/pb"
	"fmt"
	//"log/slog"
	"net"

	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

type GRPCService struct {
	pb.UnimplementedOrderServiceServer
	drinks *pb.Drinks
	orders *pb.Orders
}

func StartGrpcServer() error {
	// Create a new gRPC server.
	srv := grpc.NewServer()
	// Create grpc service
	grpcService := &GRPCService{}
	// Register our service implementation with the gRPC server.
	pb.RegisterOrderServiceServer(srv, grpcService)
	// Serve gRPC server on port 4000.
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		return err
	}
	err = srv.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}

// todo implement functions

func (c *GRPCService) AddDrinks() {
	if c.drinks != nil {
		return
	}

	c.drinks = &pb.Drinks{
		Drinks: []*pb.Drink{
			{Id: 1, Name: "Water", Price: 1.80, Description: "Yay water"},
			{Id: 2, Name: "Evil Water", Price: 6.66, Description: "AAAA OH NO"},
			{Id: 3, Name: "Leaf Water Hot", Price: 2.80, Description: "Oh that's tea mmmm AAAAA HOT OH NO"},
		},
	}
}

func (c *GRPCService) GetDrinks(ctx context.Context, in *emptypb.Empty) (*pb.Drinks, error) {
	c.AddDrinks()
	return c.drinks, nil
}

func (c *GRPCService) OrderDrink(ctx context.Context, in *pb.Order) (*wrapperspb.BoolValue, error) {
	if in == nil {
		return &wrapperspb.BoolValue{Value: false}, errors.New("order is nil")
	}

	c.AddDrinks()

	// Check if drink with that ID exists
	var exists bool
	for _, d := range c.drinks.Drinks {
		if d.Id == in.DrinkId {
			exists = true
			break
		}
	}
	if !exists {
		return &wrapperspb.BoolValue{Value: false}, fmt.Errorf("drink with id %d not found", in.DrinkId)
	}

	// Initialize order list if needed
	if c.orders == nil {
		c.orders = &pb.Orders{}
	}

	// Append order
	c.orders.Items = append(c.orders.Items, in)

	return &wrapperspb.BoolValue{Value: true}, nil
}

func (c *GRPCService) GetOrders(ctx context.Context, in *emptypb.Empty) (*pb.Orders, error) {
	if c.orders == nil || len(c.orders.Items) == 0 {
		return nil, errors.New("no orders")
	}
	return c.orders, nil
}
