package model

type GetActiveElementsResponse struct {
	Elements []*UiCategoryLogic `json:"elements"`
}

type GetTypesResponse struct {
	Elements []*UiTypeLogic `json:"elements"`
}
