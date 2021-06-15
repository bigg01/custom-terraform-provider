# Blog Posting

## About Boxboat

This source code is for the blog on [boxboat.com](https://boxboat.com). At Boxboat, we bring vast industry experience to the table, and our engineers
have worked across different clouds, architectures, and environments. Reach out if you're interested in talking Terraform, DevOps, or anything
cloud!

## Blog Post

The blog post talks about how to write your own custom Terraform Provider, and why you might do that.

Check out the full blog post [here](https://boxboat.com/2020/02/04/writing-a-custom-terraform-provider).


```
terraform providers mirror ~/.terraform.d/plugins

TF_LOG=debug terraform init

```

https://www.terraform.io/docs/cli/config/config-file.html#development-overrides-for-provider-developers

```
export TF_CLI_CONFIG_FILE=provider_overwrite.tfrc 

```

https://learn.hashicorp.com/tutorials/terraform/provider-setup


1082  make build
1083  mkdir -p ~/.terraform.d/plugins/registry.terraform.io/bigg01/cmdb/0.1/linux_amd64/
1084  mv terraform-provider-cmdb ~/.terraform.d/plugins/registry.terraform.io/bigg01/cmdb/0.1/linux_amd64/
1085  ls  ~/.terraform.d/plugins/registry.terraform.io/bigg01/cmdb/0.1/linux_amd64/
