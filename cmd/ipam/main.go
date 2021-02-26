package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/nordix/nvip/pkg/ipam"
	"github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)

	port, err := strconv.Atoi(os.Getenv("IPAM_PORT"))
	if err != nil || port <= 0 {
		port = 100
	}

	i, _ := ipam.NewIpam(port)

	i.Start()
}