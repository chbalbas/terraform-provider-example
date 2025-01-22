//--------------------
// THIS FILE WAS MANUALLY GENERATED
//--------------------
package api_provider 

import (
	"context"
	"os"
    //TODO
    "<import-your-custom-api-client>"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ provider.Provider = New()
)

// New is a helper function to simplify provider server and testing implementation.
// offers an interface for the API
func New() provider.Provider {
	var c client.ApiClient = &client.SdkClient{}
	return &apiProvider{client: c}
}

// apiProvider is the provider implementation.
type apiProvider struct {
	client client.ApiClient
}

// apiProviderModel maps provider schema data to a Go type.
type apiProviderModel struct {
	Host     types.String `tfsdk:"host"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

// Metadata returns the provider type name.
func (p *apiProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "api"
}

// Schema defines the provider-level schema for configuration data.
func (p *apiProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Interact with API.",
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Description: "URI for API. May also be provided via API_HOST environment variable.",
				Optional:    true,
			},
			"username": schema.StringAttribute{
				Description: "Username for API. May also be provided via API_USERNAME environment variable.",
				Optional:    true,
			},
			"password": schema.StringAttribute{
				Description: "Password for API. May also be provided via API_PASSWORD environment variable.",
				Optional:    true,
				Sensitive:   true,
			},
		},
	}
}

// Configure prepares a HashiCups API client for data sources and resources.
func (p *apiProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring API client")

	// Retrieve provider data from configuration
	var config apiProviderModel
	diags := req.Config.Get(ctx, &config)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	host := os.Getenv("API_HOST")
	username := os.Getenv("API_USERNAME")
	password := os.Getenv("API_PASSWORD")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	if !config.Username.IsNull() {
		username = config.Username.ValueString()
	}

	if !config.Password.IsNull() {
		password = config.Password.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing API Host",
			"The provider cannot create the API client as there is a missing or empty value for the API host. "+
				"Set the host value in the configuration or use the API_HOST environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if username == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Missing API Username",
			"The provider cannot create the API client as there is a missing or empty value for the API username. "+
				"Set the username value in the configuration or use the API_USERNAME environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if password == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing API Password",
			"The provider cannot create the API client as there is a missing or empty value for the API password. "+
				"Set the password value in the configuration or use the API_PASSWORD environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "api_host", host)
	ctx = tflog.SetField(ctx, "api_username", username)
	ctx = tflog.SetField(ctx, "api_password", password)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "api_password")

	tflog.Debug(ctx, "Creating client")

	p.client.ConfigureClient(username, password, host)

	// type Configure methods.
	resp.DataSourceData = p.client
	resp.ResourceData = p.client

	tflog.Info(ctx, "Configured client", map[string]any{"success": true})
}

// Register DataSources defines the data sources implemented in the provider.
func (p *apiProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewOrderDataSource,
	}
}

// Register Resources defines the resources implemented in the provider.
func (p *apiProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewOrderResource,
	}
}