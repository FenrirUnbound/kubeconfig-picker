package konfig

import (
	"path/filepath"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

// Konfig Wrapper object for picking specific contexts from a kube config
type Konfig struct {
	kubeconfig *api.Config
}

// ListContexts List the contexts associated with Konfig
// A sanity function for displaying what contexts are
// loaded with Konfig
func (k *Konfig) ListContexts() []string {
	result := make([]string, len(k.kubeconfig.Contexts))
	i := 0
	for key := range k.kubeconfig.Contexts {
		result[i] = key
		i++
	}

	return result
}

// SelectContexts Pick contexts by name and return an object containing the
// subset.
// If a context does not exist in the Konfig, there will be _no_ error returned
// and the context will be omitted from the returned result
func (k *Konfig) SelectContexts(contextNames []string) *api.Config {
	result := api.NewConfig()

	for _, context := range contextNames {
		selected, ok := k.kubeconfig.Contexts[context]
		if !ok {
			// skip for now. should be an error
			continue
		}
		result.Contexts[context] = selected

		clusterName := selected.Cluster
		result.Clusters[clusterName] = k.kubeconfig.Clusters[clusterName]

		user := selected.AuthInfo
		result.AuthInfos[user] = k.kubeconfig.AuthInfos[user]
	}

	return result
}

// SelectContextsAsYaml Pick contexts by name and return the kube config in its
// YAML form.
func (k *Konfig) SelectContextsAsYaml(contextNames []string) ([]byte, error) {
	config := k.SelectContexts(contextNames)

	data, err := clientcmd.Write(*config)

	return data, err
}

// NewKonfig Generate a Konfig object that wraps a kube config with the ability
// to create subset configurations.
func NewKonfig(kubeConfigFilepath string) (*Konfig, error) {
	absPath, _ := filepath.Abs(kubeConfigFilepath)

	config, err := clientcmd.LoadFromFile(absPath)
	if err != nil {
		return nil, err
	}

	result := &Konfig{
		kubeconfig: config,
	}

	return result, nil
}
