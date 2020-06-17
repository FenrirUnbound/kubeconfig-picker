package main

import (
	"fmt"
	"log"
	"os/user"
	"path"

	"github.com/FenrirUnbound/kubeconfig-picker/konfig"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	k, err := konfig.NewKonfigFromFile(path.Join(usr.HomeDir, ".kube/config"))
	if err != nil {
		log.Fatal(err)
	}

	availableContexts := k.ListContexts()
	// A list of contexts contained in the kube config
	fmt.Printf("%v\n", availableContexts)

	config, _ := k.SelectContextsAsYaml([]string{"context1", "anotherContext"})
	// A kube config with only "context1" and "anotherContext" derived from the main config
	fmt.Printf("%v\n", string(config))
}
