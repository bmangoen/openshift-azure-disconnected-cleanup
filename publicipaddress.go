package main

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-04-01/network"
)

func getIPClient() network.PublicIPAddressesClient {
	ipClient := network.NewPublicIPAddressesClient(clientData.SubscriptionID)
	ipClient.Authorizer = authorizer
	//ipClient.AddToUserAgent(config.UserAgent())
	return ipClient
}

// GetPublicIP returns an existing public IP
func GetPublicIP(resourceGroupName, ipName string) (network.PublicIPAddress, error) {
	ipClient := getIPClient()
	return ipClient.Get(ctx, resourceGroupName, ipName, "")
}

// DeletePublicIP gets info on a loadbalancer
func DeletePublicIP(resourceGroupName, ipName string) (network.PublicIPAddressesDeleteFuture, error) {
	ipClient := getIPClient()
	return ipClient.Delete(
		ctx,
		resourceGroupName,
		ipName)
}