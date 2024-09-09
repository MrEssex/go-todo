package database

import (
	"github.com/kubex/keystone-go/keystone"
	"github.com/kubex/keystone-go/proto"
	"github.com/mressex/go-todo/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	vendorID    = "mressex"
	appID       = "todo"
	accessToken = "e185b3pB0tQ3i7uetPPUA5WSXN-2YXVYuQG8qmELPI75wKCPU" //Obv use env but it's local :D
)

var keystoneConnection *keystone.Connection
var ksGrpcConn *grpc.ClientConn
var pClient proto.KeystoneClient

func InitKeyStone() {
	var err error

	ksGrpcConn, err = grpc.NewClient("localhost:51051", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithIdleTimeout(time.Minute*5), grpc.WithConnectParams(grpc.ConnectParams{MinConnectTimeout: time.Second * 5}))

	if err != nil {
		log.Fatalf("Failed to create gRPC connection: %v", err)
	}

	pClient = proto.NewKeystoneClient(ksGrpcConn)
	keystoneConnection = keystone.NewConnection(pClient, vendorID, appID, accessToken)

	keystoneConnection.RegisterTypes(models.Todo{})
	keystoneConnection.SyncSchema().Wait()
}

func CloseKeyStone() error {
	if ksGrpcConn != nil {
		return ksGrpcConn.Close()
	}

	return nil
}

func Conn() *keystone.Connection {
	return keystoneConnection
}

func Actor() *keystone.Actor {
	a := keystoneConnection.Actor("todo", "21.21.21.21", "todo", "UserAgent")
	return &a
}
