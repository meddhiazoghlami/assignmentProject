package cmd

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/meddhiazoghlami/assignmentProject/db"
	"github.com/meddhiazoghlami/assignmentProject/grpcserver"
	server "github.com/meddhiazoghlami/assignmentProject/httpserver"
	pb "github.com/meddhiazoghlami/assignmentProject/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var RootCmd = &cobra.Command{Use: "app"}

var runCmd = &cobra.Command{
	Use:   "with",
	Short: "Run the application",
	Run: func(cmd *cobra.Command, args []string) {
		withRest, _ := cmd.Flags().GetBool("rest")
		withGRPC, _ := cmd.Flags().GetBool("grpc")

		if withRest {
			runWithRest()
		} else if withGRPC {
			runWithGRPC()
		} else {
			fmt.Println("Please specify either --rest or --grpc")
		}
	},
}

var rest bool
var grpcs bool

func init() {
	runCmd.Flags().BoolVar(&rest, "rest", false, "Run with Rest API")
	runCmd.Flags().BoolVar(&grpcs, "grpc", false, "Run with gRPC server")
	RootCmd.AddCommand(runCmd)
}

func runWithRest() {
	fmt.Println("Running with Rest API...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	dbconfig := db.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_NAME"),
	}
	db := db.BuildDBConfig(dbconfig)
	defer db.Close()
	server := &server.Server{
		Db: db,
	}
	r := server.SetupRouter()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func runWithGRPC() {
	fmt.Println("Running with gRPC server...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	dbconfig := db.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_NAME"),
	}
	db := db.BuildDBConfig(dbconfig)
	defer db.Close()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGetBalanceServer(s, &grpcserver.GrpcServer{Db: db})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
