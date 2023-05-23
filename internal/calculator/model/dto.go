package model

type GetActiveElementsResponse struct {
	Elements []*UiElementLogic `json:"elements"`
}

type GetTypesResponse struct {
	Elements []*UiTypeLogic `json:"elements"`
}
