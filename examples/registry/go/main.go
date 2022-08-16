package main

import (
	"fmt"
	"github.com/masonyc/pulumi-registry-bridge/sdk/go/registry"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		resource, err := registry.NewRegistryResource(ctx, "__example.resource.id_1",
			&registry.RegistryResourceArgs{
				Name: pulumi.String("__example.resource.id_1"),
			},
		)
		if err != nil {
			fmt.Println(resource)
		}

		return nil
	})
}
