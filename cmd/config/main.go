// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"

	"github.com/spf13/cobra"
	"sigs.k8s.io/kustomize/cmd/config/cmd"
	"sigs.k8s.io/kustomize/kyaml/yaml/merge2"
	"sigs.k8s.io/kustomize/kyaml/yaml/merge3"
)

var root = &cobra.Command{
	Use:   "kyaml",
	Short: "kyaml reference comand",
	Long: `Description:
  Reference implementation for using the kyaml libraries.
`,
	Example: ``,
}

func main() {
	root.PersistentFlags().BoolVar(&cmd.StackOnError, "stack-trace", false,
		"print a stack-trace on failure")

	cmd.ExitOnError = true
	root.AddCommand(cmd.GrepCommand())
	root.AddCommand(cmd.TreeCommand())
	root.AddCommand(cmd.CatCommand())
	root.AddCommand(cmd.FmtCommand())
	root.AddCommand(cmd.MergeCommand())
	root.AddCommand(cmd.CountCommand())
	root.AddCommand(cmd.RunFnCommand())
	root.AddCommand(&cobra.Command{Use: "merge", Long: merge2.Help})
	root.AddCommand(&cobra.Command{Use: "merge3", Long: merge3.Help})

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
