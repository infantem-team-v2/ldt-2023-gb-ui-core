package uiInterface

import "gb-ui-core/internal/ui/model"

type RelationalRepository interface {
	GetUiTypes() ([]*model.UiTypeDAO, error)
}

type NonRelationalRepository interface {
	GetActiveCalculatorElements() ([]*model.UiInputElementUnitDAO, error)
}
