package model

type UiTypeDAO struct {
	Id              int32  `db:"id"`
	Name            string `db:"name"`
	Comment         string `db:"comment"`
	MultipleOptions bool   `db:"multiple_options"`
}

type UiInputCategoryDAO struct {
	Category string                   `json:"_id" bson:"_id"`
	Elements []*UiInputElementUnitDAO `json:"elements" bson:"elements"`
}

type UiInputElementUnitDAO struct {
	Field   string        `json:"field" bson:"field"`
	FieldId string        `json:"field_id" bson:"field_id"`
	Comment string        `json:"comment" bson:"comment"`
	Type    string        `json:"type" bson:"type"`
	Options []interface{} `json:"options" bson:"options"`
}
