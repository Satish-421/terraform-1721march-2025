package provider

import (
	"context"
	"log"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDockerContainer() *schema.Resource {

	return &schema.Resource {

		Description: "This resource will support create, read, update and delete docker containers via Terraform.",

		CreateContext : resourceCreateDockerContainer,
		ReadContext   : resourceReadDockerContainer,
		UpdateContext : resourceUpdateDockerContainer,
		DeleteContext : resourceDeleteDockerContainer,

		Schema: map[string]*schema.Schema {
			"container_name": {
				Description: "Name of the container.",
		                Type:         schema.TypeString,
				Required:     true,
			},

			"image_name": {
			        Description: "Name of the docker image.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"host_name": {
			        Description: "Hostname of the docker container.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func resourceCreateDockerContainer( ctx context.Context, d *schema.ResourceData, meta any ) diag.Diagnostics {

	//Retreive docker container parameters from terraform .tf scripts

	imageName     		:= d.Get("image_name").(string)
	containerName 		:= d.Get("container_name").(string)
	containerHostname	:= d.Get("host_name").(string)

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()
/*
	reader, err := cli.ImagePull(ctx, imageName, image.PullOptions{})
	if err != nil {
		panic(err)
	}
*/


	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		Hostname: containerHostname,
	}, nil, nil, nil, containerName)
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		panic(err)
	}

	d.Set("container_name", containerName )
	d.Set("host_name", containerName )
	d.SetId(resp.ID)

	return resourceReadDockerContainer(ctx,d,meta)
}

func resourceReadDockerContainer( ctx context.Context, d *schema.ResourceData, meta any ) diag.Diagnostics {

	return nil
}


func resourceUpdateDockerContainer( ctx context.Context, d *schema.ResourceData, meta any ) diag.Diagnostics {
	return nil
}


func resourceDeleteDockerContainer( ctx context.Context, d *schema.ResourceData, meta any ) diag.Diagnostics {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	containerName 		:= d.Get("container_name").(string)
	removeOptions		:= container.RemoveOptions{RemoveVolumes: true, Force: true}

	err = cli.ContainerStop(ctx, containerName, container.StopOptions{})
	if err != nil {
		log.Printf("Unable to stop container: %s", err)
		return nil
	}

	err = cli.ContainerRemove(ctx, containerName, removeOptions)
	if err != nil {
		log.Printf("Unable to delete container: %s", err)
		return nil
	}

	return nil
}
