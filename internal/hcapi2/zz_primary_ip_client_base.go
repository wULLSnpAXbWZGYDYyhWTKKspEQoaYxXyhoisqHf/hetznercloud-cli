// Code generated by interfacer; DO NOT EDIT

package hcapi2

import (
	"context"

	"github.com/hetznercloud/hcloud-go/hcloud"
)

// PrimaryIPClientBase is an interface generated for "github.com/hetznercloud/hcloud-go/hcloud.PrimaryIPClient".
type PrimaryIPClientBase interface {
	All(context.Context) ([]*hcloud.PrimaryIP, error)
	Assign(context.Context, hcloud.PrimaryIPAssignOpts) (*hcloud.Action, *hcloud.Response, error)
	ChangeDNSPtr(context.Context, hcloud.PrimaryIPChangeDNSPtrOpts) (*hcloud.Action, *hcloud.Response, error)
	ChangeProtection(context.Context, hcloud.PrimaryIPChangeProtectionOpts) (*hcloud.Action, *hcloud.Response, error)
	Create(context.Context, hcloud.PrimaryIPCreateOpts) (*hcloud.PrimaryIPCreateResult, *hcloud.Response, error)
	Delete(context.Context, *hcloud.PrimaryIP) (*hcloud.Response, error)
	Get(context.Context, string) (*hcloud.PrimaryIP, *hcloud.Response, error)
	GetByID(context.Context, int) (*hcloud.PrimaryIP, *hcloud.Response, error)
	GetByIP(context.Context, string) (*hcloud.PrimaryIP, *hcloud.Response, error)
	GetByName(context.Context, string) (*hcloud.PrimaryIP, *hcloud.Response, error)
	List(context.Context, hcloud.PrimaryIPListOpts) ([]*hcloud.PrimaryIP, *hcloud.Response, error)
	Unassign(context.Context, int) (*hcloud.Action, *hcloud.Response, error)
	Update(context.Context, *hcloud.PrimaryIP, hcloud.PrimaryIPUpdateOpts) (*hcloud.PrimaryIP, *hcloud.Response, error)
}
