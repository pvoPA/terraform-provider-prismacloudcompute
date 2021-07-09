package prismacloudcompute

import (
	"bytes"
	"encoding/base64"
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-compute-go"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/collection"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceCollections() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCollectionsRead,

		Schema: map[string]*schema.Schema{

			// Output.
			"accountids": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of account IDs.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"appids": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of application IDs.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"clusters": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of Kubernetes cluster names.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"coderepos": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of code repositories.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"color": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A hex color code for a collection.",
			},
			"containers": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of containers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A free-form text description of the collection.",
			},
			"functions": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of functions.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"hosts": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of hosts.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"images": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of images.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"labels": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of labels.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"modified": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Date/time when the collection was last modified.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique collection name.",
			},
			"namespaces": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of Kubernetes namespaces.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"owner": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User who created or last modified the collection.",
			},
			"prisma": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If set to 'true', this collection originates from Prisma Cloud.",
			},
			"system": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If set to 'true', this collection was created by the system (i.e., a non-user). Otherwise (false) a real user.",
			},
		},
	}
}

func dataSourceCollectionsRead(d *schema.ResourceData, meta interface{}) error {
	var buf bytes.Buffer
	client := meta.(*pc.Client)

	items, err := collection.List(client)
	if err != nil {
		return err
	}

	if buf.Len() == 0 {
		d.SetId("all")
	} else {
		d.SetId(base64.StdEncoding.EncodeToString(buf.Bytes()))
	}
	d.Set("total", len(items))

	list := make([]interface{}, 0, len(items))
	for _, i := range items {
		list = append(list, map[string]interface{}{
			"accountIDs":  StringSliceToSet(i.AccountIDs),
			"appIDs":      StringSliceToSet(i.AppIDs),
			"clusters":    StringSliceToSet(i.Clusters),
			"codeRepos":   StringSliceToSet(i.CodeRepos),
			"color":       i.Color,
			"containers":  StringSliceToSet(i.Containers),
			"description": i.Description,
			"functions":   StringSliceToSet(i.Functions),
			"hosts":       StringSliceToSet(i.Hosts),
			"images":      StringSliceToSet(i.Images),
			"labels":      StringSliceToSet(i.Labels),
			"modified":    i.Modified,
			"name":        i.Name,
			"namespaces":  StringSliceToSet(i.Namespaces),
			"owner":       i.Owner,
			"prisma":      i.Prisma,
			"system":      i.System,
		})
	}

	if err := d.Set("listing", list); err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}

	return nil
}