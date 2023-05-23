package model

type UiTypeDAO struct {
	Id              int32  `db:"id"`
	Name            string `db:"name"`
	Comment         string `db:"comment"`
	MultipleOptions bool   `db:"multiple_options"`
}

type UiInputElementUnitDAO struct {
	Field   string        `json:"field"`
	Comment string        `json:"comment"`
	Type    string        `json:"type"`
	Options []interface{} `json:"options"`
}
