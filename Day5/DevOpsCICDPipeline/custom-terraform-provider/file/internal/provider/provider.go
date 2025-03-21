package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown
}

type FileConfig struct {
  name string
  content string
}

//New is the name of function that take version as a string and returns a function 
//func() * schema.Provider is the function return by the New function
// *schema.Provider - the returned function is returning a schema.Provider address
func New(version string) func() *schema.Provider {

	return func() *schema.Provider {
		p := &schema.Provider {
			DataSourcesMap: map[string]*schema.Resource {

			},
			ResourcesMap: map[string]*schema.Resource {
				"localfile": resourceLocalFile(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure( version string, p *schema.Provider) func( context.Context, *schema.ResourceData) ( any, diag.Diagnostics ) {
	return func(context.Context, *schema.ResourceData) ( any, diag.Diagnostics) {
		fileConfig := FileConfig{}

		return &fileConfig, nil
	}
}
