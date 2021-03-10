package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	server "github.com/will-rowe/archer/pkg/protocol/grpc"
	"github.com/will-rowe/archer/pkg/service/v1"
)

// launchCmd represents the launch command
var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Launch the Archer service",
	Long:  `Launch the Archer service.`,
	Run: func(cmd *cobra.Command, args []string) {
		launchArcher()
	},
}

func init() {
	grpcAddr = launchCmd.Flags().String("grpcAddress", DefaultServerAddress, "address to announce on")
	grpcPort = launchCmd.Flags().String("grpcPort", DefaultgRPCport, "TCP port to listen to by the gRPC server")
	rootCmd.AddCommand(launchCmd)
}

// launchArcher sets up and runs the gRPC Archer service
func launchArcher() {

	// get top level context
	ctx := context.Background()

	// get the service API
	serverAPI := service.NewArcher()

	// run the server until shutdown signal received
	addr := fmt.Sprintf("%s:%s", *grpcAddr, *grpcPort)
	if err := server.Launch(ctx, serverAPI, addr, *logFile); err != nil {
		log.Fatal(err)
	}

	// clean up the service API
	if err := serverAPI.Shutdown(); err != nil {
		log.Fatalf("could not shutdown the Archer service: %v", err)
	}
}
