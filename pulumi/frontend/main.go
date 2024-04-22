package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Build and publish the Docker image
		image, err := docker.NewImage(ctx, "my-image", &docker.ImageArgs{
			Build: docker.DockerBuildArgs{
				Context: pulumi.String("./app"),
			},
			ImageName: pulumi.String("myapp:latest"),
		})
		if err != nil {
			return err
		}

		// Start a container using the built Docker image
		_, err = docker.NewContainer(ctx, "my-container", &docker.ContainerArgs{
			Image: image.BaseImageName,
			Ports: docker.ContainerPortArray{
				&docker.ContainerPortArgs{
					Internal: pulumi.Int(80),
					External: pulumi.Int(80),
				},
			},
		})
		if err != nil {
			return err
		}

		return nil
	})
}