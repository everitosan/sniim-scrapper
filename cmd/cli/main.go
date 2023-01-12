package main

import (
	"fmt"

	"github.com/everitosan/sniim-scrapper/cmd/cli/initial"
	"github.com/everitosan/sniim-scrapper/cmd/cli/query"
	"github.com/everitosan/sniim-scrapper/cmd/cli/request"
	"github.com/everitosan/sniim-scrapper/internal/config"
	"github.com/everitosan/sniim-scrapper/internal/transport/repository"
	"github.com/everitosan/sniim-scrapper/internal/transport/repository/filestorage"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	config := config.LoadConfig()

	if config.DEBUG {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug("Debug Log level")
	}

	// Repositories
	marketRepo, _ := filestorage.NewMarketFileRepository(config.CATALOGUE_SRC)
	paramsRepo, _ := filestorage.NewParamsFileRepository(config.CATALOGUE_SRC, "params")
	queryRepo, _ := filestorage.NewQueryFileRepository(config.CATALOGUE_SRC, "queries")
	consultResponseRepo, _ := filestorage.NewConsultResponseFileRepository(config.CATALOGUE_SRC, "responses")

	productRepo, _ := filestorage.NewProductFileRepository(config.CATALOGUE_SRC, "product")
	productSourceRepo, _ := filestorage.NewOptionSelectFileRepository(config.CATALOGUE_SRC, "productSource")
	productDestinyRepo, _ := filestorage.NewOptionSelectFileRepository(config.CATALOGUE_SRC, "productDestiny")
	pricePresentationRepo, _ := filestorage.NewOptionSelectFileRepository(config.CATALOGUE_SRC, "pricePresentation")
	weekRepo, _ := filestorage.NewOptionSelectFileRepository(config.CATALOGUE_SRC, "week")
	monthRepo, _ := filestorage.NewOptionSelectFileRepository(config.CATALOGUE_SRC, "month")
	yearRepo, _ := filestorage.NewOptionSelectFileRepository(config.CATALOGUE_SRC, "year")

	rContainer := repository.Repository{
		Market:            marketRepo,
		Params:            paramsRepo,
		Query:             queryRepo,
		ConsultResponse:   consultResponseRepo,
		Product:           productRepo,
		ProductSource:     productSourceRepo,
		ProductDestiny:    productDestinyRepo,
		PricePresentation: pricePresentationRepo,
		Week:              weekRepo,
		Month:             monthRepo,
		Year:              yearRepo,
	}

	rootCmd := &cobra.Command{
		Use: "sniim-cli",
		Run: func(cmd *cobra.Command, args []string) {
			version, _ := cmd.Flags().GetBool("version")
			if version {
				fmt.Println(`
██╗░░░██╗░█████╗░░░░░█████╗░░░░░░███╗░░
██║░░░██║██╔══██╗░░░██╔══██╗░░░░████║░░
╚██╗░██╔╝██║░░██║░░░██║░░██║░░░██╔██║░░
░╚████╔╝░██║░░██║░░░██║░░██║░░░╚═╝██║░░
░░╚██╔╝░░╚█████╔╝██╗╚█████╔╝██╗███████╗
░░░╚═╝░░░░╚════╝░╚═╝░╚════╝░╚═╝╚══════╝`)
			}
		},
	}
	rootCmd.Flags().BoolP("version", "v", false, "Show version of the cli")

	rootCmd.AddCommand(initial.Command(config.SNIIM_ADDR, rContainer))
	rootCmd.AddCommand(query.Command(config.SNIIM_ADDR, rContainer))
	rootCmd.AddCommand(request.Command(config.SNIIM_ADDR, rContainer.Query, rContainer.ConsultResponse))
	rootCmd.Execute()
}
