package filestorage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/everitosan/sniim-scrapper/internal/app/consult"
)

type queryFileRepository struct {
	dst      string
	fileName string
}

func NewQueryFileRepository(dst, fileName string) (*queryFileRepository, error) {
	var repo queryFileRepository
	err := initDir(dst)
	if err == nil {
		repo.dst = dst
		repo.fileName = fileName
	}
	return &repo, err
}

func (pR *queryFileRepository) SaveOne(content consult.Consult) error {
	var all []consult.Consult
	fileName := filepath.Join(pR.dst, pR.fileName+".json")

	all, err := pR.GetAll()
	if err != nil {
		all = make([]consult.Consult, 0, 1)
	}

	all = append(all, content)

	str, err := json.Marshal(all)
	if err != nil {
		return err
	}

	return saveJsonStrToFile(string(str), fileName)

}

func (pR *queryFileRepository) DeleteOne(index int) error {
	var all []consult.Consult
	fileName := filepath.Join(pR.dst, pR.fileName+".json")

	all, err := pR.GetAll()
	if err != nil {
		all = make([]consult.Consult, 0, 1)
	}

	allLength := len(all)

	if index > allLength-1 {
		return fmt.Errorf("invalid index")
	}

	copy(all[index:], all[index+1:])
	all[allLength-1] = consult.Consult{}
	all = all[:allLength-1]

	str, err := json.Marshal(all)
	if err != nil {
		return err
	}

	return saveJsonStrToFile(string(str), fileName)

}

func (pR *queryFileRepository) GetAll() ([]consult.Consult, error) {
	var formParams []consult.Consult
	fileName := filepath.Join(pR.dst, pR.fileName+".json")

	content, err := os.ReadFile(fileName)

	if err != nil {
		return formParams, err
	}

	err = json.Unmarshal(content, &formParams)
	return formParams, err
}
