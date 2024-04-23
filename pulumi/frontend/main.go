package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		image, err := docker.NewImage(ctx, "frontend", &docker.ImageArgs{
			Build: docker.DockerBuildArgs{
				Context:    pulumi.String("/Users/taradruffel/.workspace/repos/frontend/."),
				Dockerfile: pulumi.String("/Users/taradruffel/.workspace/repos/frontend/DockerfileDev"),
			},
			ImageName: pulumi.String("frontend"),
			SkipPush:  pulumi.Bool(true),
		})
		if err != nil {
			return err
		}


		_, err = docker.NewContainer(ctx, "frontend", &docker.ContainerArgs{
			Image: image.BaseImageName,
			Ports: docker.ContainerPortArray{
				&docker.ContainerPortArgs{
					Internal: pulumi.Int(4200),
					External: pulumi.Int(4200),
				},
			},
			Volumes: docker.ContainerVolumeArray{
				&docker.ContainerVolumeArgs{
					ContainerPath: pulumi.String("/app"),
					HostPath:      pulumi.String("/Users/taradruffel/.workspace/repos/frontend/."),
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{image}))

		if err != nil {
			return err
		}

		return nil
	})
}
