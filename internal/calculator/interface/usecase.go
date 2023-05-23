package calculatorInterface

import "gb-ui-core/internal/calculator/model"

type UseCase interface {
	GetActiveElements() (*model.GetActiveElementsResponse, error)
	GetTypes() (*model.GetTypesResponse, error)
}
