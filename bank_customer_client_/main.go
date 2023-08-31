package main

import (
	"context"
	"fmt"
	"log"
	pb "github.com/Grpc_bank/bank_customer_proto/netxd_customer"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)

	// response, err := client.CreateCustomer(context.Background(), &pb.Customer{CustomerId: 12,Balance:3000 })
	// if err != nil {
	// 	log.Fatalf("Failed to call SayHello: %v", err)
	// }

	// fmt.Printf("Response: %s\n", response)


	response1, err := client.Transaction(context.Background(), &pb.Transfer{
		From:   11,
		To:     12,
		Amount: 500,
	})
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}

	fmt.Printf("Response: %s\n", response1)
}
