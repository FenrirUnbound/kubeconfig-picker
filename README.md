# kubeconfig-picker

[![wercker status][wercker-image]][wercker-url] [![Go Report Card][goreport-image]][goreport-url]

Pick out a subset of Kube config contexts

## Build and Install

```console
go install
```

## Usage

### Run

```console
kubeconfig-picker
```

### Import

```
import (
    "fmt"

    konfig "github.com/fenrirunbound/kubeconfig-picker/konfig"
)

func main() {
    k, _ := konfig.NewKonfigFromFile("/home/me/.kube/config")
    
    availableContexts := k.ListContexts()
    // A list of contexts contained in the kube config
    fmt.Printf("%v\n", availableContexts)

    config, _ := k.SelectContextsAsYaml([]string{"context1", "anotherContext"})
    // A kube config with only "context1" and "anotherContext" derived from the main config
    fmt.Printf("%v\n", string(config))
}
```

[goreport-image]: https://goreportcard.com/badge/github.com/fenrirunbound/kubeconfig-picker
[goreport-url]: https://goreportcard.com/report/github.com/fenrirunbound/kubeconfig-picker
[wercker-image]: https://app.wercker.com/status/b97a9ce6c5baee376f4d54ceed0b7c98/s/master
[wercker-url]: https://app.wercker.com/project/byKey/b97a9ce6c5baee376f4d54ceed0b7c98
