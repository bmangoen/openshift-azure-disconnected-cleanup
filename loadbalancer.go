package main

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-04-01/network"
	"github.com/Azure/go-autorest/autorest/to"
)


func getLBClient() network.LoadBalancersClient {
	lbClient := network.NewLoadBalancersClient(clientData.SubscriptionID)
	lbClient.Authorizer = authorizer
	//lbClient.AddToUserAgent(config.UserAgent())
	return lbClient
}

// GetLoadBalancer gets info on a loadbalancer
func GetLoadBalancer(resourceGroupName, lbName string) (network.LoadBalancer, error) {
	lbClient := getLBClient()
	return lbClient.Get(ctx,resourceGroupName, lbName, "")
}


// UpdateLoadBalancer updates []LoadBalacingRules to nil (does not work...)
func UpdateLoadBalancer(resourceGroupName, lbName string) (network.LoadBalancersCreateOrUpdateFuture, error) {
	lbClient := getLBClient()
	return lbClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		lbName,
		network.LoadBalancer{
			Location: to.StringPtr("westeurope"),
			Sku: &network.LoadBalancerSku{
				Name: "Standard",
			},
			LoadBalancerPropertiesFormat: &network.LoadBalancerPropertiesFormat{
				LoadBalancingRules: nil,
			},
		})
}

// DeleteLoadBalancer gets info on a loadbalancer
func DeleteLoadBalancer(resourceGroupName, lbName string) (network.LoadBalancersDeleteFuture, error) {
	lbClient := getLBClient()
	return lbClient.Delete(
		ctx,
		resourceGroupName,
		lbName)
}
