package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"fmt"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

const (
	resourceGroupLocation = "westeurope"
)

// Information loaded from the authorization file to identify the client
type clientInfo struct {
	SubscriptionID string
	VMPassword     string
}

var (
	ctx        = context.Background()
	clientData clientInfo
	authorizer autorest.Authorizer
)

// Authenticate with the Azure services using file-based authentication
func init() {
	var err error

	if err != nil {
		log.Fatalf("Failed to get OAuth config: %v", err)
	}

	authorizer, err = auth.NewAuthorizerFromFile(azure.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		log.Fatalf("Failed to get OAuth config: %v", err)
	}

	authInfo, err := readJSON(os.Getenv("AZURE_AUTH_LOCATION"))
	if err != nil {
		log.Fatalf("Failed to read JSON: %+v", err)
	}

	clientData.SubscriptionID = (*authInfo)["subscriptionId"].(string)
	clientData.VMPassword = (*authInfo)["clientSecret"].(string)
}

func main() {

	args := os.Args[1:]
	
	if len(args) <= 0 {
		log.Fatalf("Usage: openshift-azure-cleanup <cluster_id>")
	}
	clusterID := args[0]

	resourceGroupName := clusterID + "-rg"
	fmt.Println(resourceGroupName)

	lbPublicMaster := clusterID + "-public-lb"
	fmt.Println(lbPublicMaster)

	lb, err := GetLoadBalancer(resourceGroupName, lbPublicMaster)
	if err != nil {
		log.Fatalf("failed to get lb: %v", err)
	}
	log.Printf("Public loadbalancer: %v", *lb.Location)

	deletedLbStatus, err := DeleteLoadBalancer(resourceGroupName, lbPublicMaster)
	if err != nil {
		log.Fatalf("failed to delete lb: %v", err)
	}

	log.Printf("Public loadbalancer deleted: %v", deletedLbStatus.Future)
}


func readJSON(path string) (*map[string]interface{}, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	contents := make(map[string]interface{})
	_ = json.Unmarshal(data, &contents)
	return &contents, nil
}