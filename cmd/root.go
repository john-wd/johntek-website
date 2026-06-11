/*
Copyright © 2026 JohnW (john@johntekconsulting.com)
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/gohugoio/hugo/commands"
	"github.com/gohugoio/hugo/common/herrors"
	"github.com/gohugoio/hugo/common/loggers"
	"github.com/john-wd/johntek-website/cmd/components/util"
	"github.com/john-wd/johntek-website/cmd/internal/resume"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "webctl",
	Short: "Manages Hugo website resources.",
	Long: `Wrapper on the Hugo CLI extending resource management.

It exposes commands to manage publishing content in social media,
tracking published status and compiling general templates into static
PDF files for download.`,
	PersistentPreRunE: checkRootHugo,
}

var hugoCmd = &cobra.Command{
	Use:                "hugo",
	Short:              "Runs the Hugo CLI.",
	DisableFlagParsing: true, // disable cobra's parsing to properly pass them to hugo
	Run: func(cmd *cobra.Command, args []string) {
		log.SetFlags(0)
		err := commands.Execute(os.Args[2:])
		if err != nil {
			for _, e := range herrors.Errors(err) {
				loggers.Log().Errorf("%s", e)
			}
			os.Exit(1)
		}
	},
}

func checkRootHugo(cmd *cobra.Command, args []string) error {
	if !util.FileExists("hugo.yaml") && !util.FileExists("hugo.toml") {
		return fmt.Errorf("Not in a Hugo project directory. Please run this command from the root of your Hugo project.")
	}
	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.AddCommand(hugoCmd)
	rootCmd.AddCommand(resume.Command)
}
