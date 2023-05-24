package model

type GetActiveElementsResponse struct {
	Elements []*UiCategoryLogic `json:"categories"`
}

type GetTypesResponse struct {
	Elements []*UiTypeLogic `json:"elements"`
}
