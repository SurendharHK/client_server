package main

import (
	"context"
	"fmt"
	"net"

	"github.com/Grpc_bank/bank_customer_config/config"
	"github.com/Grpc_bank/bank_customer_config/constants"
	pro "github.com/Grpc_bank/bank_customer_proto/netxd_customer"
	"github.com/Grpc_bank/bank_customer_server/controllers"
	"github.com/Grpc_bank/bank_customer_service/services"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initDatabase(client *mongo.Client) {
	
	CustomerCollection := config.GetCollection(client, "bankdb", "customers")
	transactionCollection := config.GetCollection(client, "bankdb", "transactions")
	controllers.CustomerService = services.InitCustomerService(CustomerCollection, transactionCollection, context.Background(),client)
}

func main() {
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)
	lis, err := net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	pro.RegisterCustomerServiceServer(s, &controllers.RPCServer{})

	fmt.Println("Server listening on", constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
