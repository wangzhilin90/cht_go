package zkclient

import (
	"testing"
)

func TestgetLocalIP(t *testing.T) {
	ips, err := GetLocalIP()
	if err != nil {
		t.Fatalf("TestgetLocalIP failed", err)
	}
	t.Logf("getLocalIP res", ips)
}

func TestRegisterNode(t *testing.T) {
	zkServers := []string{"192.168.8.212:2181", "192.168.8.213:2181", "192.168.8.214:2181"}
	conn, err := ConnectZk(zkServers)
	if err != nil {
		t.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	servicename := "/testnode/providers"
	err = RegisterNode(conn, servicename, "127.0.0.1:30002")
	// err = createPermanentNode(conn, servicename)
	if err != nil {
		t.Fatalf("RegisterNode failed", err)
	}
}

// func TestWatchNode(t *testing.T) {
// 	zkServers := []string{"192.168.8.212:2181", "192.168.8.213:2181", "192.168.8.214:2181"}
// 	conn, err := ConnectZk(zkServers)
// 	if err != nil {
// 		t.Fatalf("connect zk failed %v ", err)
// 	}
// 	defer conn.Close()
// }

// func TestInitSevice(t *testing.T) {
// 	zkServers := []string{"192.168.8.212:2181", "192.168.8.213:2181", "192.168.8.214:2181"}
// 	conn, err := ConnectZk(zkServers)
// 	if err != nil {
// 		t.Fatalf("connect zk failed %v ", err)
// 	}
// 	defer conn.Close()

// 	servicename = "/service/providers"
// 	err = InitSevice(conn, servicename, servicemap)
// 	if err != nil {
// 		t.Fatalf("InitSevice failed", err)
// 	}
// }

// func TestCallService(t *testing.T) {
// 	zkServers := []string{"192.168.8.212:2181", "192.168.8.213:2181", "192.168.8.214:2181"}
// 	conn, err := ConnectZk(zkServers)
// 	if err != nil {
// 		t.Fatalf("connect zk failed %v ", err)
// 	}
// 	defer conn.Close()

// 	servicePath = "/service/providers"
// 	err = CallService(conn, servicePath, servicemap)
// 	if err != nil {
// 		t.Fatalf("CallService failed", err)
// 	}
// }

// func TestDeleteNode(t *testing.T) {
// 	zkServers := []string{"192.168.8.212:2181", "192.168.8.213:2181", "192.168.8.214:2181"}
// 	conn, err := ConnectZk(zkServers)
// 	if err != nil {
// 		t.Fatalf("connect zk failed %v ", err)
// 	}
// 	defer conn.Close()

// 	servicePath := "/service/providers"
// 	err = DeleteNode(conn, servicePath, servicemap)
// 	if err != nil {
// 		t.Fatalf("DeleteNode failed", err)
// 	}
// }
