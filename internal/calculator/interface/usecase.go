package calculatorInterface

import "gb-ui-core/internal/calculator/model"

type UseCase interface {
	GetActiveElements(doAdmin bool) (*model.GetActiveElementsResponse, error)
	GetTypes() (*model.GetTypesResponse, error)
	UpdateActiveElements(params *model.SetActiveForElementRequest) error
}
