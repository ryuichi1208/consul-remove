package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/consul/api"
)

type ConsulClient interface {
	ForceLeave(node string) error
	PruneNode(node string) error
}

type MyConsulClient struct {
	client *api.Client
}

func (m *MyConsulClient) ForceLeave(node string) error {
	agent := m.client.Agent()
	return agent.ForceLeavePrune(node)
}

func (m *MyConsulClient) PruneNode(node string) error {
	opts := &api.WriteOptions{}
	catalog := m.client.Catalog()
	_, err := catalog.Deregister(&api.CatalogDeregistration{
		Node: node,
	}, opts)
	return err
}

func main() {
	var nodeName string
	var consulAddress string

	flag.StringVar(&nodeName, "node", "", "Name of the node to force leave")
	flag.StringVar(&consulAddress, "address", "http://localhost:8500", "Consul server address")
	flag.Parse()

	if nodeName == "" {
		fmt.Println("Node name is required.")
		flag.Usage()
		os.Exit(1)
	}

	config := &api.Config{
		Address: consulAddress,
		Scheme:  "http",
	}
	config.Address = consulAddress

	// Consulクライアントを作成
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("Error creating Consul client: %v", err)
	}

	// MyConsulClientのインスタンスを作成
	myConsulClient := &MyConsulClient{client: client}

	// ノードのforce leaveを実行
	err = forceLeave(myConsulClient, nodeName)
	if err != nil {
		log.Fatalf("Error forcing node to leave: %v", err)
	}

	fmt.Printf("Node %s has been forcefully removed from the cluster\n", nodeName)

	err = pruneNode(myConsulClient, nodeName)
	if err != nil {
		log.Fatalf("Error pruning node: %v", err)
	}
	fmt.Printf("Node %s has been pruned from the cluster\n", nodeName)
}

// forceLeave 関数
func forceLeave(client ConsulClient, nodeName string) error {
	return client.ForceLeave(nodeName)
}

// pruneNode 関数
func pruneNode(client ConsulClient, nodeName string) error {
	return client.PruneNode(nodeName)
}
