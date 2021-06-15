package main

import (
	"github.com/bigg01/terraform-provider-cmdb/cmdb"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cmdb.Provider})
}
