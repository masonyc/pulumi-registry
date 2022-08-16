package main

import (
	"github.com/masonyc/pulumi-registry-bridge/sdk/go/registry"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		resource, err := registry.RegistryResource(ctx, "__example.resource.id_1", nil)

		return nil
	})
}
