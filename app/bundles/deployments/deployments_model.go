// Package deployments handles all requests related to server deployments
package deployments

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/krishamoud/game-server/app/common/docker"
)

func createDSTService(token, maxPlayers, name string) (string, error) {
	spec := createServiceSpec(token, maxPlayers, name)
	service, err := docker.DockerConn.ServiceCreate(context.Background(), spec, types.ServiceCreateOptions{})
	if err != nil {
		return "", err
	}
	return service.ID, nil
}

func createServiceSpec(token, maxPlayers, name string) swarm.ServiceSpec {
	max := uint64(1)
	tok := "TOKEN=" + token
	players := "MAX_PLAYERS=" + maxPlayers
	n := "NAME=" + name
	envVars := []string{tok, players, n}
	spec := swarm.ServiceSpec{
		TaskTemplate: swarm.TaskSpec{
			RestartPolicy: &swarm.RestartPolicy{
				MaxAttempts: &max,
				Condition:   swarm.RestartPolicyConditionNone,
			},
			ContainerSpec: &swarm.ContainerSpec{
				Image: "dstacademy/dontstarvetogether",
				Env:   envVars,
				TTY:   true,
			},
		},
	}
	return spec
}
