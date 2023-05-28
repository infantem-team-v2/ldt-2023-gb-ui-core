package uiInterface

import (
	calcModel "gb-ui-core/internal/calculator/model"
	"gb-ui-core/internal/ui/model"
)

type RelationalRepository interface {
	GetUiTypes() ([]*model.UiTypeDAO, error)
}

type NonRelationalRepository interface {
	GetActiveCalculatorElements() ([]*model.UiInputElementUnitDAO, error)
	GetActiveElementsByCategory(doAdmin bool) ([]*model.UiInputCategoryDAO, error)
	UpdateActiveElements(params *calcModel.SetActiveForElementRequest) error
}
