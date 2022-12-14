package query

import (
	"fmt"
	"log"

	"github.com/everitosan/sniim-scrapper/cmd/cli/request"
	"github.com/everitosan/sniim-scrapper/internal/app/consult"
	"github.com/everitosan/sniim-scrapper/internal/transport/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const createCosultFlag = "create"
const listCosultFlag = "list"
const deleteCosultFlag = "delete"
const saveConsultFlag = "save"

func Command(sniiimAddr string, rContainer repository.Repository) *cobra.Command {
	consultCommand := &cobra.Command{
		Use:   "query",
		Short: "Administra los queries",
		Long:  "Puede crear, guardar y listar los queries",
		Run: func(cmd *cobra.Command, args []string) {

			list, _ := cmd.Flags().GetBool(listCosultFlag)
			create, _ := cmd.Flags().GetBool(createCosultFlag)
			delete, _ := cmd.Flags().GetInt16(deleteCosultFlag)
			save, _ := cmd.Flags().GetBool(saveConsultFlag)

			switch {
			case list:
				/*
				* Case for listing consults
				 */
				consults, err := rContainer.Consult.GetAll()
				if err != nil {
					logrus.Fatal(err)
				}

				for index, consult := range consults {
					fmt.Printf("(%d) - %s\n", index, consult.String())
				}
				return
			case create:
				/*
				* Case create a consult
				 */
				// Ask category and subcategory
				newConsult := askBreadCrumb(rContainer.Params)

				// Ask for required inputs
				dateDetected := askInputs(rContainer, newConsult)

				// Ask dates
				if !dateDetected {
					askDates(newConsult)
				}

				results, err := consult.Scrap(sniiimAddr, *newConsult)

				if err != nil {
					logrus.Fatal(err)
				}

				if len(results) == 0 {
					logrus.Warn("No se encontraron resultados")
				} else {
					request.PrintResultTable(results)
				}

				if save {
					rContainer.Consult.SaveOne(*newConsult)
				} else {
					res, err := confirmPrompt("¿Desea guardar la consulta?")
					if err != nil {
						logrus.Warn("No se guardará el query")
					}
					if res == "y" {
						rContainer.Consult.SaveOne(*newConsult)
					}
				}

				return
			case delete != -1:
				/*
				* Delete case
				 */
				err := rContainer.Consult.DeleteOne(int(delete))
				if err != nil {
					log.Fatal(err)
				}
				return
			}

		},
	}

	consultCommand.Flags().BoolP(createCosultFlag, "c", false, "Crea un query")
	consultCommand.Flags().BoolP(saveConsultFlag, "s", false, "Bandera que indica si debe guardarse el query")
	consultCommand.Flags().BoolP(listCosultFlag, "l", false, "Muestra todos los queries guardados")
	consultCommand.Flags().Int16P(deleteCosultFlag, "d", -1, "Elimina el query indicado")

	return consultCommand
}
