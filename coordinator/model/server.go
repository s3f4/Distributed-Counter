package model

import (
	"coordinator/util"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

//Server keeps servers ports and process IDs
type Server struct {
	Port      int
	ProcessID int
	DataCount int // this property is keeping server's total data count
}

//StartServer runs a service process service=server
func StartServer() (processID int, port int, err error) {
	freeport, err := util.GetFreePort()
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

//InitServers initializes servers with given server count
func InitServers(serverCount int) ([]*Server, error) {
	servers := make([]*Server, serverCount)
	for i := 0; i < serverCount; i++ {
		ProcessID, Port, err := StartServer()
		if err != nil {
			return nil, err
		}
		servers[i] = &Server{
			ProcessID: ProcessID,
			Port:      Port,
			DataCount: 0,
		}
	}
	return servers, nil
}

//KillServers kills all servers or specific server
func KillServers(processID int, servers []*Server) {
	if processID != 0 {
		syscall.Kill(processID, syscall.SIGINT)
	} else {
		for _, server := range servers {
			syscall.Kill(server.ProcessID, syscall.SIGINT)
		}
	}
}
