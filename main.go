package main

import (
	"log"

	v "github.com/appscode/go/version"
	"github.com/spf13/cobra"
	"github.com/xunchangguo/chartify/pkg/cmd"
)

var (
	Version         string
	VersionStrategy string
	Os              string
	Arch            string
	CommitHash      string
	GitBranch       string
	GitTag          string
	CommitTimestamp string
	BuildTimestamp  string
	BuildHost       string
	BuildHostOs     string
	BuildHostArch   string
)

func init() {
	v.Version.Version = Version
	v.Version.VersionStrategy = VersionStrategy
	v.Version.Os = Os
	v.Version.Arch = Arch
	v.Version.CommitHash = CommitHash
	v.Version.GitBranch = GitBranch
	v.Version.GitTag = GitTag
	v.Version.CommitTimestamp = CommitTimestamp
	v.Version.BuildTimestamp = BuildTimestamp
	v.Version.BuildHost = BuildHost
	v.Version.BuildHostOs = BuildHostOs
	v.Version.BuildHostArch = BuildHostArch
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "chartify [command]",
		Short: `Generate Helm Charts from Kubernetes api objects`,
		Run: func(c *cobra.Command, args []string) {
			c.Help()
		},
	}
	rootCmd.AddCommand(cmd.NewCmdCreate())
	rootCmd.AddCommand(v.NewCmdVersion())
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
