package commands

import (
	"fmt"

	cons "github.com/dhyanio/kubefront/internal/config/constants"
	models "github.com/dhyanio/kubefront/internal/config/models"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Kubefront version details.",
	Long:  `Fetching KubeFront's current server and client versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		versionServer, versionClient := fetchKubefrontVersion()
		fmt.Printf(cons.Discription)
		fmt.Printf("KubeFront-Server: %s \n", versionServer)
		fmt.Printf("KubeFront-Client: %s \n", versionClient)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func fetchKubefrontVersion() (string, string) {
	var details models.TomlDetails
	if _, err := toml.DecodeFile("details.toml", &details); err != nil {
		fmt.Println(err)
	}
	return details.VersionServer, details.VersionClient
}
