package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/Nviswateja/CustomerTrackerService/service/protos"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	for {
		fmt.Printf("1.Enter the customer Name\n 2.Get Customer Names")
		var option int
		fmt.Scan(&option)
		if option == 1 {
			AddCustomer(conn)
		} else if option == 2 {
			GetCustomerDetails(conn)
		}
	}
}

func GetCustomerDetails(conn *grpc.ClientConn) {
	c := pb.NewCustomerServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetCustomers(ctx, &pb.GetCustomerMessageRequest{})
	if err != nil {
		log.Fatalf("could not add customer: %v", err)
	}
	for _, customer := range r.Customers {
		fmt.Println(customer)
	}
}

func AddCustomer(conn *grpc.ClientConn) {
	var name string
	fmt.Scan(&name)
	c := pb.NewCustomerServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddCustomerDetails(ctx, &pb.CustomerMessageRequest{Name: name})
	if err != nil {
		log.Fatalf("could not add customer: %v", err)
	}
	log.Printf("reply: %s", r.GetMessage())
}
