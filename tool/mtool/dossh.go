package main

import (
	"fmt"
	"time"

	"golang.org/x/crypto/ssh"
)

func createSSHClient(svr *Server) *ssh.Client {

	config := &ssh.ClientConfig{
		Timeout:         5 * time.Second,
		User:            svr.user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth:            []ssh.AuthMethod{ssh.Password(svr.password)},
	}

	addr := fmt.Sprintf("%s:%s", svr.ip, svr.port)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return sshClient
}

func runCMD(cmd string, sshClient *ssh.Client) error {
	session, err := sshClient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	ret, err := session.CombinedOutput(cmd)
	if err != nil {
		return err
	}
	fmt.Printf("run cmd [%s] on [%s] sucess:[%s] \n", cmd, sshClient.RemoteAddr(), string(ret))

	return nil

}
