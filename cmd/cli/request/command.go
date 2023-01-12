package request

import (
	_ "embed"

	"github.com/everitosan/sniim-scrapper/internal/app/consult"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const allFlag = "all"
const indexFlag = "index"
const saveFlag = "save"

func Command(sniimAddr string, queryRepo consult.QueryRepository, responseRepo consult.ConsultResponseRepository) *cobra.Command {
	requestCommand := &cobra.Command{
		Use:   "request",
		Short: "Consulta de información",
		Long:  "Consulta la información de un query",
		Run: func(cmd *cobra.Command, args []string) {

			index, _ := cmd.Flags().GetInt32(indexFlag)
			shouldSave, _ := cmd.Flags().GetBool(saveFlag)
			queries, err := queryRepo.GetAll()

			if err != nil {
				logrus.Fatal(err)
			}

			switch {
			case index != -1:
				/*
				* Case for makinng a single request
				 */
				if int(index) >= len(queries) {
					logrus.Warnf("No existe consulta número %d", index)
					return
				}

				selectedConsult := queries[index]
				results, err := consult.Scrap(sniimAddr, selectedConsult)

				if err != nil {
					logrus.Fatal(err)
				}

				if len(results) == 0 {
					logrus.Warn("No hay resultados de la búsqueda.")
					return
				}

				if shouldSave {
					err = responseRepo.Save(results)
				} else {

					PrintResultTable(results)
				}

				if err != nil {
					logrus.Fatal(err)
				}
				return

			}

		},
	}

	requestCommand.Flags().Int32P(indexFlag, "i", -1, "Realiza una consulta basado en el índice de query")
	requestCommand.Flags().BoolP(allFlag, "a", false, "Realiza todas las consultas de los queries guardados")
	requestCommand.Flags().BoolP(saveFlag, "s", false, "Guarda el resultado de los queries ejecutados")

	return requestCommand
}
