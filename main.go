package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb1 "github.com/Nviswateja/CustomerTrackerService/service/protos/Version1Protos"
	pb2 "github.com/Nviswateja/CustomerTrackerService/service/protos/Version2Protos"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	for {
		fmt.Printf("1.Enter the customer Name\n 2.Get Customer Names\n3.Get customer with number\n")
		var option int
		fmt.Scan(&option)
		if option == 1 {
			AddCustomer(conn)
		} else if option == 2 {
			GetCustomerDetails(conn)
		} else if option ==3{
			GetCustomerDetailsWithNumber(conn)
		}
	}
}

func GetCustomerDetailsWithNumber(conn *grpc.ClientConn) {
	c := pb2.NewCustomerServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetCustomerWithName(ctx, &pb2.GetCustomerMessageRequest{})
	if err != nil {
		log.Fatalf("could not add customer: %v", err)
	}
	for _, customer := range r.Customers {
		fmt.Println(customer.Name+"--"+customer.PhoneNumber)
	}
}

func GetCustomerDetails(conn *grpc.ClientConn) {
	c := pb1.NewCustomerServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetCustomers(ctx, &pb1.GetCustomerMessageRequest{})
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
	c := pb1.NewCustomerServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddCustomerDetails(ctx, &pb1.CustomerMessageRequest{Name: name})
	if err != nil {
		log.Fatalf("could not add customer: %v", err)
	}
	log.Printf("reply: %s", r.GetMessage())
}
