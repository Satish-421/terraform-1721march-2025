package provider

import (
	"os"
	"io/ioutil"
	"log"
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLocalFile() *schema.Resource {

	return &schema.Resource {

		Description: "This resource will support create, read, update and delete file via Terraform.",

		CreateContext : resourceCreateFile,
		ReadContext  : resourceReadFile,
		UpdateContext: resourceUpdateFile,
		DeleteContext: resourceDeleteFile,

		Schema: map[string]*schema.Schema {
			"file_name": {
				Description: "Name of the file.",
		                Type:         schema.TypeString,
				Required:     true,
			},

			"file_content": {
			        Description: "Content that must be stored/retrieved to/from the file.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func resourceCreateFile( ctx context.Context, d *schema.ResourceData, meta any ) diag.Diagnostics {

	//Retrieve the inputs user provided in the terraform script .tf file
	fileName := d.Get("file_name").(string)
	content  := d.Get("file_content").(string)

	myfile, err := os.Create( fileName )

	if err != nil {
		panic(err)
	}

	myfile.WriteString( content + "\n" )
	myfile.Sync()

	d.SetId( "resource-100" )

	resourceReadFile ( ctx, d, meta )

	return nil
}

func resourceReadFile( ctx context.Context, d *schema.ResourceData, meta any ) diag.Diagnostics {
	//Retrieve the inputs user provided in the terraform script .tf file
	fileName := d.Get("file_name").(string)

	content, err := ioutil.ReadFile( fileName )

	if err != nil {
		log.Fatal(err)
	}

	d.Set( "content", content )

	return nil
}


func resourceUpdateFile( ctx context.Context, d *schema.ResourceData, meta any ) diag.Diagnostics {
	return nil
}


func resourceDeleteFile( ctx context.Context, d *schema.ResourceData, meta any ) diag.Diagnostics {
	//Retrieve the inputs user provided in the terraform script .tf file
	fileName := d.Get("file_name").(string)

	err := os.Remove(fileName)

	if err != nil {
	   log.Fatal(err)
	}

	return nil
}
