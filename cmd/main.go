package main

import (
	"log"
	"os"

	"github.com/vlayco/hex-go/internal/adapters/app/api"
	"github.com/vlayco/hex-go/internal/adapters/core/arithmetic"
	"github.com/vlayco/hex-go/internal/adapters/framework/right/db"
	"github.com/vlayco/hex-go/internal/ports"

	gRPC "github.com/vlayco/hex-go/internal/adapters/framework/left/grpc"
)

func main() {
	var err error

	// ports
	var dbaseAdapter ports.DBPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate DB connection")
	}
	defer dbaseAdapter.CloseDBConnection()

	core = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(dbaseAdapter, core)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)
	gRPCAdapter.Run()

}
