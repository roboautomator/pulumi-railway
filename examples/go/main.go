package main

import (
	"example.com/pulumi-railway/sdk/go/railway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myRandomResource, err := railway.NewRandom(ctx, "myRandomResource", &railway.RandomArgs{
			Length: pulumi.Int(24),
		})
		if err != nil {
			return err
		}
		_, err = railway.NewRandomComponent(ctx, "myRandomComponent", &railway.RandomComponentArgs{
			Length: pulumi.Int(24),
		})
		if err != nil {
			return err
		}
		ctx.Export("output", pulumi.StringMap{
			"value": myRandomResource.Result,
		})
		return nil
	})
}
