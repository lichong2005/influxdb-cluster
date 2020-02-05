package main

import (
	"github.com/influxdb-cluster/cmd/influxd-ctl/command"
)

func main() {
	command := command.NewCommand()
	command.Execute()
}
