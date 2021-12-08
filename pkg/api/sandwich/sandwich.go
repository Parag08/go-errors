package sandwich

import (
	"github.com/parag08/go-errors/pkg/errors"

	"github.com/parag08/go-errors/pkg/logger"
)

// Sandwich Service contains the methods of the sandwich
type SandwichService interface {
	New() ([]string, error)
}

type sandwichService struct {
}

func NewSandWichService() SandwichService {
	return &sandwichService{}
}

func (s *sandwichService) New() ([]string, error) {
	var (
		err         error
		ingredients []string
	)
	ingredients, err = getIngredients()
	if err != nil {
		logger.Log.Debug(err)
	}
	return ingredients, err
}

func getIngredients() ([]string, error) {
	var (
		op errors.Op = "sandwich.getIngredients"
	)

	avacados, err := buyAvacados()

	if err != nil {
		return []string{}, errors.E(op, err)
	}

	boiledEggs, err := buyBoliedEggs()
	if err != nil {
		return []string{}, errors.E(op, err)
	}

	bread, err := buyBread()
	if err != nil {
		return []string{}, errors.E(op, err)
	}
	return []string{avacados, boiledEggs, bread}, nil
}

func buyAvacados() (string, error) {
	var (
		op   errors.Op   = "sandwich.buyAvacados"
		kind errors.Kind = errors.KindInternalServerError
	)
	return "Avacados", errors.E("coudn't buy avacado", op, kind)
}

func buyBoliedEggs() (string, error) {
	return "BoliedEggs", nil
}

func buyBread() (string, error) {
	return "Bread", nil
}
