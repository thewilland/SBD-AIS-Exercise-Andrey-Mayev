package client

import (
	"context"
	"exc8/pb"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcClient struct {
	client pb.OrderServiceClient
}

func NewGrpcClient() (*GrpcClient, error) {
	conn, err := grpc.NewClient(":4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewOrderServiceClient(conn)
	return &GrpcClient{client: client}, nil
}

func (c *GrpcClient) Run() error {
	// todo
	ctx := context.Background()

	// 1. List drinks
	fmt.Println("Available drinks: ")
	drinksResp, err := c.client.GetDrinks(ctx, &emptypb.Empty{})
	if err != nil {
		return fmt.Errorf("GetDrinks failed: %w", err)
	}
	for _, d := range drinksResp.Drinks {
		fmt.Printf("[%d] %s - %f euros (%s)\n", d.Id, d.Name, d.Price, d.Description)
	}

	// 2. Order a few drinks
	fmt.Println("\nOrdering drinks: ")

	order1 := &pb.Order{DrinkId: 1, Quantity: 2}
	_, err = c.client.OrderDrink(ctx, order1)
	if err != nil {
		return fmt.Errorf("OrderDrink failed: %w", err)
	}
	fmt.Println("Ordered 2 Waters")

	order2 := &pb.Order{DrinkId: 2, Quantity: 1}
	_, err = c.client.OrderDrink(ctx, order2)
	if err != nil {
		return fmt.Errorf("OrderDrink failed: %w", err)
	}
	fmt.Println("Ordered 1 Evil Water")

	// 3. Order more drinks
	order3 := &pb.Order{DrinkId: 3, Quantity: 4}
	_, err = c.client.OrderDrink(ctx, order3)
	if err != nil {
		return fmt.Errorf("OrderDrink failed: %w", err)
	}
	fmt.Println("Ordered 4 Teas")

	order4 := &pb.Order{DrinkId: 1, Quantity: 2}
	_, err = c.client.OrderDrink(ctx, order4)
	if err != nil {
		return fmt.Errorf("OrderDrink failed: %w", err)
	}
	fmt.Println("Ordered 2 more Waters")

	// 4. Get order total
	fmt.Println("\nOrders: ")
	ordersResp, err := c.client.GetOrders(ctx, &emptypb.Empty{})
	if err != nil {
		return fmt.Errorf("GetOrders failed: %w", err)
	}

	// Calculating total using drink prices
	total := 0.0
	for _, o := range ordersResp.Items {
		// find drink
		var drink *pb.Drink
		for _, d := range drinksResp.Drinks {
			if d.Id == o.DrinkId {
				drink = d
			}
		}

		if drink != nil {
			lineTotal := float64(drink.Price) * float64(o.Quantity)
			total += lineTotal
			fmt.Printf("%d * %s = %f euros\n", o.Quantity, drink.Name, lineTotal)
		} else {
			fmt.Printf("%d * (unknown drink id %d)\n", o.Quantity, o.DrinkId)
		}
	}

	fmt.Printf("\nTotal order price: %f euros\n", total)

	//
	// print responses after each call
	return nil
}
