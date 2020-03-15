Terraform Provider for Databricks
==================

[![Build Status](https://travis-ci.com/innovationnorway/terraform-provider-databricks.svg?branch=master)](https://travis-ci.com/innovationnorway/terraform-provider-databricks)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.12.x
-	[Go](https://golang.org/doc/install) 1.13

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/innovationnorway/terraform-provider-databricks`

```sh
$ mkdir -p $GOPATH/src/github.com/innovationnorway; cd $GOPATH/src/github.com/innovationnorway
$ git clone https://github.com/innovationnorway/terraform-provider-databricks.git
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/innovationnorway/terraform-provider-databricks
$ make build
```

Using the provider
----------------------

```hcl
provider "databricks" {
  host  = var.databricks_host
  token = var.databricks_token
}

resource "databricks_cluster" "example" {
  cluster_name  = "example"
  spark_version = "6.3.x-scala2.11"
  node_type_id  = "Standard_DS3_v2"

  autoscale {
    min_workers = 2
    max_workers = 8
  }

  autotermination_minutes = 120
}
```

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-databricks
...
```

You can also cross-compile if necessary:

```sh
GOOS=windows GOARCH=amd64 make build
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

In order to run a subset of Acceptance tests, you can run:

```sh
make testacc TESTARGS='-run=TestAccDatabricksCluster'
```

The following environment variables must be set prior to running acceptance tests:

- `DATABRICKS_HOST`
- `DATABRICKS_TOKEN`
