# kubeconfig-picker

Pick out a subset of Kube config contexts

## Usage

```
import (
    "fmt"

    konfig "github.com/fenrirunbound/kubeconfig-picker"
)

func main() {
    k, _ := konfig.NewKonfig("/home/me/.kube/config")
    
    availableContexts := k.ListContexts()
    // A list of contexts contained in the kube config
    fmt.Printf("%v\n", availableContexts)

    config, _ := k.SelectContextsAsYaml([]string{"context1", "anotherContext"})
    // A kube config with only "context1" and "anotherContext" derived from the main config
    fmt.Printf("%v\n", string(config))
}
```