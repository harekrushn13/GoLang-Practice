package linux_plugin

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"strings"
	"time"
)

type DiscoveryStatus struct {
	Status        bool   `json:"status"`
	StatusMessage string `json:"status_message"`
}

func Discovery(ip, port, username, password string) DiscoveryStatus {

	client, errMsg := checkSSH(ip, port, username, password)

	if errMsg != "" {
		return DiscoveryStatus{
			Status:        false,
			StatusMessage: errMsg,
		}
	}
	defer client.Close()

	session, err := client.NewSession()

	if err != nil {
		return DiscoveryStatus{
			Status:        false,
			StatusMessage: fmt.Sprintf("Error creating session: %s", err),
		}
	}
	defer session.Close()

	var out bytes.Buffer

	session.Stdout = &out

	if err := session.Run("whoami"); err != nil {
		return DiscoveryStatus{
			Status:        false,
			StatusMessage: fmt.Sprintf("Error running whoami: %s", err),
		}
	}

	session, err = client.NewSession()
	if err != nil {
		return DiscoveryStatus{
			Status:        false,
			StatusMessage: fmt.Sprintf("Error creating second session: %s", err),
		}
	}
	defer session.Close()

	out.Reset()
	session.Stdout = &out

	if err := session.Run("uname -a"); err != nil {
		return DiscoveryStatus{
			Status:        false,
			StatusMessage: fmt.Sprintf("Error running uname -a: %s", err),
		}
	}

	unameOutput := strings.ToLower(out.String())

	if !strings.Contains(unameOutput, "linux") {
		return DiscoveryStatus{
			Status:        false,
			StatusMessage: fmt.Sprintf("%s is not linux", unameOutput),
		}
	}

	return DiscoveryStatus{
		Status:        true,
		StatusMessage: fmt.Sprintf("All check passed"),
	}
}

func checkSSH(ip, port, username, password string) (*ssh.Client, string) {

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", ip, port), config)

	if err != nil {
		return nil, fmt.Sprintf("Failed to connect to %s: %s", ip, err)
	}

	return client, ""
}
