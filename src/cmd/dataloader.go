package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/gufranmirza/go-popular-repositories-autocomplete-api/clients"
	"github.com/gufranmirza/go-popular-repositories-autocomplete-api/dal/repositorydal"
	"github.com/spf13/cobra"
)

// serveCmd represents the loadCmd command
var loadCmd = &cobra.Command{
	Use:   "dataloader",
	Short: "Initialize database records",
	Long:  `Download popular repository from github and store into database`,
	Run: func(cmd *cobra.Command, args []string) {
		load()
	},
}

func init() {
	RootCmd.AddCommand(loadCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func load() {
	fmt.Println("Making HTTP call to Github API")
	repos, err := clients.NewClient().GetRepositories()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Collected Repositories %v Github API\n", len(repos))
	for _, repo := range repos {
		repo.CreatedTimestampUTC = time.Now().UTC()
		repo.UpdatedTimestampUTC = time.Now().UTC()
		err = repositorydal.NewRepositoryDal().Create(&repo)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
