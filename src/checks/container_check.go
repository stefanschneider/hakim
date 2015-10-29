package checks

import (
	"errors"

	"github.com/cloudfoundry-incubator/garden"
	gclient "github.com/cloudfoundry-incubator/garden/client"
	gconnection "github.com/cloudfoundry-incubator/garden/client/connection"
)

func ContainerCheck(gardenAddr string) error {
	client := gclient.New(gconnection.New("tcp", gardenAddr))
	container, err := client.Create(garden.ContainerSpec{})
	if container != nil {
		defer client.Destroy(container.Handle())
	}

	if err != nil {
		return errors.New("Failed to create container\n" + err.Error())
	}
	return nil
}
