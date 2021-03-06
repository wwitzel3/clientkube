package clientkube

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"go.uber.org/multierr"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery/cached/disk"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/bryanl/clientkube/pkg/cluster"
)

// OutOfClusterClient is a client that be used out of cluster.
type OutOfClusterClient struct {
	client          dynamic.Interface
	dir             string
	discoveryClient *disk.CachedDiscoveryClient
}

var _ cluster.Client = &OutOfClusterClient{}

// NewOutOfClusterClient creates an instance of OutOfClusterClient.
func NewOutOfClusterClient(kubeconfig string) (*OutOfClusterClient, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("build config: %w", err)
	}

	client, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("create cluster client: %w", err)
	}

	dir, err := ioutil.TempDir("", "clientkube")
	if err != nil {
		return nil, fmt.Errorf("create temporary directory")
	}

	discoveryClient, err := disk.NewCachedDiscoveryClientForConfig(
		config,
		dir,
		dir,
		180*time.Second,
	)
	if err != nil {
		return nil, fmt.Errorf("create discovery client")
	}

	c := OutOfClusterClient{
		dir:             dir,
		client:          client,
		discoveryClient: discoveryClient,
	}

	return &c, nil
}

// Close closes the client and cleans up its resources.
func (c *OutOfClusterClient) Close() error {
	var err error

	if c.dir != "" {
		if rErr := os.RemoveAll(c.dir); rErr != nil {
			err = multierr.Append(err, fmt.Errorf("remove temporary directory: %w", err))
		}
	}

	return err
}

// Resources lists the resources available in the cluster.
func (c *OutOfClusterClient) Resources() (cluster.Resources, error) {
	resourceLists, err := c.discoveryClient.ServerPreferredResources()
	if err != nil {
		return nil, fmt.Errorf("get server preferred resources: %w", err)
	}

	var list cluster.Resources

	for _, resourceList := range resourceLists {
		groupVersion, err := schema.ParseGroupVersion(resourceList.GroupVersion)
		if err != nil {
			return nil, fmt.Errorf("parse group version: %w", err)
		}

		for _, apiResource := range resourceList.APIResources {
			list = append(list, newResource(groupVersion, apiResource))
		}
	}

	return list, nil
}

// List lists objects in the cluster.
func (c *OutOfClusterClient) List(
	ctx context.Context,
	res schema.GroupVersionResource,
	options cluster.ListOptions) (*unstructured.UnstructuredList, error) {
	if options.Namespace == "" {
		return c.client.Resource(res).List(ctx, options.ListOptions)
	}

	return c.client.Resource(res).Namespace(options.Namespace).List(ctx, options.ListOptions)
}

// Watch watches a resource.
func (c *OutOfClusterClient) Watch(
	ctx context.Context,
	res schema.GroupVersionResource,
	options cluster.ListOptions) (cluster.Watch, error) {
	if options.Namespace == "" {
		return c.client.Resource(res).Watch(ctx, options.ListOptions)
	}

	return c.client.Resource(res).Namespace(options.Namespace).Watch(ctx, options.ListOptions)
}
