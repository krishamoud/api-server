// Package docker initializes a new docker client connection
package docker

import "github.com/docker/docker/client"

var (
	// DockerConn is the persistent connection to the docker client
	DockerConn *client.Client
	err        error
)

func init() {
	DockerConn, err = client.NewEnvClient()
}
