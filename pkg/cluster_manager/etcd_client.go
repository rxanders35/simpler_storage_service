package cluster_manager

import (
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdClient struct{}

func NewEtcdClient() *EtcdClient {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
	}
	defer cli.Close()
}

// goal:
// build out a fully working etcd client that implements the MetaDataServer interface.
// define structs for cluster/object metadata
// build replication algorithm
