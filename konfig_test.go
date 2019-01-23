package konfig_test

import (
	"io/ioutil"
	"testing"

	konfig "github.com/fenrirunbound/kubeconfig-picker"
	"github.com/stretchr/testify/assert"
)

func TestNewKonfig(t *testing.T) {
	kubeconfig, err := ioutil.ReadFile("./test_files/valid_file.yaml")
	if err != nil {
		assert.Fail(t, err.Error())
	}

	cfg, err := konfig.NewKonfig(kubeconfig)
	assert.Nil(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, cfg.ListContexts(), []string{"myContextName"})
}

func TestNewKonfigFromFile(t *testing.T) {
	testFilepath := "./test_files/valid_file.yaml"

	cfg, err := konfig.NewKonfigFromFile(testFilepath)

	assert.Nil(t, err)
	assert.NotNil(t, cfg)
}

func TestNewConfigFromFileNonexistingConfig(t *testing.T) {
	fakeFilepath := "./no_file_here.yaml"

	cfg, err := konfig.NewKonfigFromFile(fakeFilepath)

	assert.NotNil(t, err)
	assert.Nil(t, cfg)
}

func TestListContexts(t *testing.T) {
	testFilepath := "./test_files/valid_file.yaml"
	cfg, _ := konfig.NewKonfigFromFile(testFilepath)

	contexts := cfg.ListContexts()

	assert.Equal(t, contexts, []string{"myContextName"})
}

func TestPickContexts(t *testing.T) {
	testFilepath := "./test_files/bigger_file.yaml"
	cfg, _ := konfig.NewKonfigFromFile(testFilepath)

	slimmerConfig := cfg.SelectContexts([]string{"lastContextName"})
	assert.NotNil(t, slimmerConfig)

	assert.Equal(t, "randomCluster", slimmerConfig.Contexts["lastContextName"].Cluster)
	assert.Equal(t, "https://somewhere.com", slimmerConfig.Clusters["randomCluster"].Server)
	assert.Equal(t, "someUser", slimmerConfig.Contexts["lastContextName"].AuthInfo)
	assert.Equal(t, []byte("dummyData"), slimmerConfig.AuthInfos["someUser"].ClientKeyData)
}

func TestPickNonexistingContext(t *testing.T) {
	testFilepath := "./test_files/bigger_file.yaml"
	cfg, _ := konfig.NewKonfigFromFile(testFilepath)

	slimmerConfig := cfg.SelectContexts([]string{"thisContextCannotExist"})
	assert.NotNil(t, slimmerConfig)
	assert.Len(t, slimmerConfig.Contexts, 0)
}

func TestPickContextsAsYaml(t *testing.T) {
	testFilepath := "./test_files/bigger_file.yaml"
	cfg, _ := konfig.NewKonfigFromFile(testFilepath)

	slimmerConfig, err := cfg.SelectContextsAsYaml([]string{"lastContextName"})
	assert.Nil(t, err)

	targetTestContent, _ := ioutil.ReadFile("./test_files/verify_bigger_file.yaml")
	output := string(slimmerConfig)
	assert.Equal(t, string(targetTestContent), output)
}
