# monocle [![GoDoc](https://godoc.org/github.com/gosuri/monocle?status.svg)](https://godoc.org/github.com/gosuri/monocle) [![Build Status](https://travis-ci.org/gosuri/monocle.svg?branch=master)](https://travis-ci.org/gosuri/monocle)

Monocle is a go library for advanced command line help. It extends [Cobra](https://github.com/spf13/cobra) library and provides a set of powerful features to customize generate help for a root command.

Monocle, along with [cmdns](https://github.com/gosuri/cmdns) powers most of [OvrClk](http://ovrclk.com) toolchain.

## Example Usage

The source code for the below example is available at [example/example.go](example/example.go)

```go
ovrclk := &cobra.Command{Use: "ovrclk", Short: "Utility to manage your clusters and applications"}

var app string
appInfo := &cobra.Command{Use: "info", Short: "display app info", Run: runFunc}
appInfo.Flags().StringVarP(&app, "app", "a", "", "app name")
apps := &cobra.Command{Use: "apps", Short: "manage apps", Long: "list all apps", Run: runFunc}
apps.AddCommand(appInfo)
ovrclk.AddCommand(apps)

var cluster string
clusterInfo := &cobra.Command{Use: "info", Short: "display cluster info", Run: runFunc}
clusterInfo.Flags().StringVarP(&cluster, "clusters", "c", "", "cluster name")
clusters := &cobra.Command{Use: "clusters", Short: "manage clusters", Long: "list all clusters", Run: runFunc}
clusters.AddCommand(clusterInfo)
ovrclk.AddCommand(clusters)

version := &cobra.Command{Use: "version", Short: "display version", Run: runFunc}
ovrclk.AddCommand(version)

// Enable monocle
monocle.Enable(ovrclk)

// Set primary topics
monocle.Primary(apps, clusters)

ovrclk.Execute()
```

Will produce the help like:

```sh
Utility to manage your clusters and applications

Usage: ovrclk COMMAND 

Primary help topics, type "ovrclk help TOPIC" for more details:

  apps          manage apps
  clusters      manage clusters 

Additional topics:

  version       display version
```
