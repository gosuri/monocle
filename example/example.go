package main

import (
	"fmt"

	"github.com/gosuri/cmdns"
	"github.com/gosuri/monocle"
	"github.com/spf13/cobra"
)

var helpFunc = func(cmd *cobra.Command, args []string) { cmd.Help() }
var runFunc = func(cmd *cobra.Command, args []string) { fmt.Println("run", cmd.Name()) }

func main() {
	var verbose bool
	ovrclk := &cobra.Command{Use: "ovrclk", Short: "Utility to manage your clusters and applications"}
	ovrclk.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose mode")

	var app string
	appInfo := &cobra.Command{Use: "info", Short: "display app info", Run: runFunc}
	appInfo.Flags().StringVarP(&app, "app", "a", "", "app name")
	apps := &cobra.Command{Use: "apps", Short: "manage apps", Long: "List all apps", Run: runFunc}
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
	cmdns.Namespace(ovrclk)

	ovrclk.Execute()
}
