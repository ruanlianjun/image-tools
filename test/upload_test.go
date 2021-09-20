package test

import (
	"context"
	"github.com/smallnest/rpcx/log"
	"github.com/smallnest/rpcx/share"
	"testing"

	"github.com/smallnest/rpcx/client"
)

func TestUpload(t *testing.T) {
	discovery, err := client.NewPeer2PeerDiscovery("tcp@localhost:8080", "")
	if err != nil {
		log.Fatalf("discovery rpc server err:%#v", err)
	}
	xClient := client.NewXClient(share.SendFileServiceName, client.Failfast, client.RandomSelect, discovery, client.DefaultOption)
	err = xClient.SendFile(context.Background(), "./file/a.txt", 10000, map[string]string{
		"upload": "uploadName",
	})
	if err != nil {
		log.Fatalf("upload sendFile err:%#v", err)
	}
	log.Info("send file success")
}
