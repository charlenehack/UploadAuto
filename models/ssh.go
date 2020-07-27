package models

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
)

type ClientConfig struct {
	Host string
	Port int64
	Username string
	Password string
	Client *ssh.Client
	Session *ssh.Session
	err error
	Result string
}

func (cc *ClientConfig) SSHConnect(host, username, password string, port int64) {
	cc.Host, cc.Port, cc.Username, cc.Password = host, port, username, password
	config := ssh.ClientConfig{
		User: cc.Username,
		Auth: []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	addr := fmt.Sprintf("%s:%d", cc.Host, cc.Port)
	if cc.Client, cc.err = ssh.Dial("tcp", addr, &config); cc.err != nil {
		fmt.Println("err:", cc.err)
	}
}

func (cc *ClientConfig) RunCMD(cmd string) error  {
	if cc.Session, cc.err = cc.Client.NewSession(); cc.err != nil {
		return  cc.err
	}
	defer cc.Session.Close()

	if output, err := cc.Session.CombinedOutput(cmd); err != nil {
		return  err
	} else {
		cc.Result = string(output)
	}

	return nil
}