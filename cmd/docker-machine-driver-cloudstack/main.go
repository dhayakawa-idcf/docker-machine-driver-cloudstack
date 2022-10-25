package main

import (
	cloudstack "github.com/dhayakawa-idcf/docker-machine-driver-cloudstack"
	"github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(cloudstack.NewDriver("", ""))
}
