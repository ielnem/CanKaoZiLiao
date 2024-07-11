package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"

	scp "github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
	"golang.org/x/crypto/ssh"
)

func ReadFile() *bytes.Reader {
	content, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return bytes.NewReader(content)
}

func doScpWork(reader *bytes.Reader, svr *Server) error {

	clientConfig, _ := auth.PasswordKey(svr.user, svr.password, ssh.InsecureIgnoreHostKey())
	addr := fmt.Sprintf("%s:%s", svr.ip, svr.port)
	scpClient := scp.NewClient(addr, &clientConfig)

	err := scpClient.Connect()
	if err != nil {
		return err
	}
	defer scpClient.Close()

	upPath := path.Join(upPath, filepath.Base(fileName))

	err = scpClient.CopyFile(context.Background(), reader, upPath, "0644")
	if err != nil {
		return err
	}
	fmt.Printf("copy file %s to %s:%s sucess \n", fileName, svr.ip, upPath)

	return nil
}
