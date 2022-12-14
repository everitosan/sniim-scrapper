package initial

import (
	"github.com/everitosan/sniim-scrapper/internal/app/scraper"
	"github.com/everitosan/sniim-scrapper/internal/transport/repository"
	"github.com/spf13/cobra"
)

func Command(sniimAddr string, rContainer repository.Repository) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Crea los catálogos",
		Long:  "Obtiene la información de la fuente y crea los catálogos",
		Run: func(cmd *cobra.Command, args []string) {
			scraper.InitCatalogues(sniimAddr, rContainer)
		},
	}
}
