package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/orbitdb/go-orbit-db/http"
)

var incCmd = &cobra.Command{
	Use:   "inc <database> [<increment>]",
	Aliases: []string{"increment"},
	Short: "Increments a counter.",
	Long: `
The inc command increments a counter data store. If no <increment> is provided
counter is incremented by 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: orbitdb inc <database> []<increment>]")
			os.Exit(1)
		}

		var r http.Request

		r.SetURL(httpHost, "/db/", args[0], "/inc")

		if len(args) == 2 {
			r.SetURL(r.Url, "/", args[1])
		}

		fmt.Println(r.Post())
	},
}

func init() {
	rootCmd.AddCommand(incCmd)
}
