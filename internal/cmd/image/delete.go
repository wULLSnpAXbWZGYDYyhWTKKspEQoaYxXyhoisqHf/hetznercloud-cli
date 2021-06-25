package image

import (
	"context"

	"github.com/hetznercloud/cli/internal/cmd/base"
	"github.com/hetznercloud/cli/internal/hcapi2"
	"github.com/hetznercloud/hcloud-go/hcloud"
)

var deleteCmd = base.DeleteCmd{
	ResourceNameSingular: "image",
	ShortDescription:     "Delete a image",
	NameSuggestions:      func(c hcapi2.Client) func() []string { return c.Image().Names },
	Fetch: func(ctx context.Context, client hcapi2.Client, idOrName string) (interface{}, *hcloud.Response, error) {
		return client.Image().Get(ctx, idOrName)
	},
	Delete: func(ctx context.Context, client hcapi2.Client, resource interface{}) error {
		image := resource.(*hcloud.Image)
		if _, err := client.Image().Delete(ctx, image); err != nil {
			return err
		}
		return nil
	},
}
