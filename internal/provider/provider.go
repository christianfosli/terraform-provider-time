package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

func New() provider.Provider {
	return &timeProvider{}
}

var (
	_ provider.Provider             = (*timeProvider)(nil)
	_ provider.ProviderWithMetadata = (*timeProvider)(nil)
)

type timeProvider struct{}

func (p *timeProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "time"
}

func (p *timeProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

}

func (p *timeProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return nil
}

func (p *timeProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewTimeOffsetResource,
		NewTimeRotatingResource,
		NewTimeSleepResource,
		NewTimeStaticResource,
	}
}

func (p *timeProvider) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{}, nil
}
