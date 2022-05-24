package main

import (
	"context"
	"log"
	"time"

	pb "productinfo/client/ecommerce"

	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProductInfoClient(conn)
	oc := pb.NewOrderManagementClient(conn)

	dialProductInfo(c)
	dialOrderManagement(oc)
}

func dialOrderManagement(c pb.OrderManagementClient) {
	// Get Order
	retrievedOrder, err := c.GetOrder(context.Background(), &wrappers.StringValue{Value: "106"})
	if err != nil {
		log.Fatalf("Could not get Order: %v", err)
	}
	log.Print("GetOrder Response -> : ", retrievedOrder)
}

func dialProductInfo(c pb.ProductInfoClient) {
	name := "Apple iPhone 11"
	description := `Meet Apple iPhone 11. All-new dual-camera system with
					Ultra Wide and Night mode.`
	price := float32(1000.0)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddProduct(ctx,
		&pb.Product{Name: name, Description: description, Price: price})
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	log.Printf("Product ID: %s added successfully", r.Value)

	product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	if err != nil {
		log.Fatalf("Could not get product: %v", err)
	}
	log.Println("Product:", product)
}
