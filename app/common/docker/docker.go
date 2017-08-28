// Package docker initializes a new docker client connection
package docker

import "github.com/docker/docker/client"

var (
	DockerConn *client.Client
	err        error
)

func init() {
	DockerConn, err = client.NewEnvClient()
}
