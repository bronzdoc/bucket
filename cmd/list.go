package cmd

import (
	"github.com/bronzdoc/bucket/lib/s3"
	"github.com/spf13/cobra"
)

var sortBy string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List AWS buckets",
	Run: func(cmd *cobra.Command, args []string) {

		s3 := s3.New()
		s3.ListBuckets()
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	listCmd.Flags().StringVarP(
		&sortBy,
		"sort",
		"s",
		"",
		`
	- dateCreated
	- dateModified
	- alphaAsc
	- alphaDesc

	Example: bucket list --sort dateCreated
		`,
	)
}
