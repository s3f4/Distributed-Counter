package node

import (
	"coordinator/model"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

// GetFreePort asks the kernel for a free open port that is ready to use.
func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

//StartNode runs a service process service=node
func StartNode() (processID int, port int, err error) {
	freeport, err := GetFreePort()
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Print(strconv.Itoa(freeport))
	cmd := exec.Command("../service/service", "--port", strconv.Itoa(freeport))
	cmd.Stdout = os.Stdout
	err = cmd.Start()
	if err != nil {
		return 0, 0, err
	}
	fmt.Printf("Port: %v\n", strconv.Itoa(freeport))
	fmt.Printf("Service ProcessID : %v\n", cmd.Process.Pid)
	return cmd.Process.Pid, freeport, nil
}

//InitNodes initializes servers with given server count
func InitNodes(serverCount int) ([]*model.Node, error) {
	servers := make([]*model.Node, serverCount)
	for i := 0; i < serverCount; i++ {
		ProcessID, Port, err := StartNode()
		if err != nil {
			return nil, err
		}
		servers[i] = &model.Node{
			ProcessID: ProcessID,
			Port:      Port,
			DataCount: 0,
		}
	}
	return servers, nil
}

//KillNodes kills all servers or specific server
func KillNodes(processID int, servers []*model.Node) {
	if processID != 0 {
		syscall.Kill(processID, syscall.SIGINT)
	} else {
		for _, server := range servers {
			syscall.Kill(server.ProcessID, syscall.SIGINT)
		}
	}
}
