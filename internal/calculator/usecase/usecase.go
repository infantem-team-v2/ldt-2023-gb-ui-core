package usecase

import (
	"gb-ui-core/internal/calculator/model"
	uiInterface "gb-ui-core/internal/ui/interface"
	"github.com/sarulabs/di"
)

type CalculatorUseCase struct {
	UiPostgresRepo uiInterface.RelationalRepository    `di:"uiPostgresRepo"`
	UiMongoRepo    uiInterface.NonRelationalRepository `di:"uiMongoRepo"`
}

func BuildCalculatorUseCase(ctn di.Container) (interface{}, error) {
	return &CalculatorUseCase{
		UiPostgresRepo: ctn.Get("uiPostgresRepo").(uiInterface.RelationalRepository),
		UiMongoRepo:    ctn.Get("uiMongoRepo").(uiInterface.NonRelationalRepository),
	}, nil
}

func (cus *CalculatorUseCase) GetActiveElements() (*model.GetActiveElementsResponse, error) {
	elementsDAO, err := cus.UiMongoRepo.GetActiveElementsByCategory()
	if err != nil {
		return nil, err
	}
	elementsDTO := make([]*model.UiCategoryLogic, 0, len(elementsDAO))

	for _, e := range elementsDAO {
		innerElements := make([]*model.UiElementLogic, 0, len(e.Elements))
		for _, innerE := range e.Elements {
			innerElements = append(innerElements, &model.UiElementLogic{
				Field:   innerE.Field,
				FieldId: innerE.FieldId,
				Comment: innerE.Comment,
				Type:    innerE.Type,
				Options: innerE.Options,
			})
		}
		elementsDTO = append(elementsDTO, &model.UiCategoryLogic{
			Category: e.Category,
			Elements: innerElements,
		})
	}
	return &model.GetActiveElementsResponse{
		Elements: elementsDTO,
	}, nil
}

func (cus *CalculatorUseCase) GetTypes() (*model.GetTypesResponse, error) {
	typesDAO, err := cus.UiPostgresRepo.GetUiTypes()
	if err != nil {
		return nil, err
	}

	typesDTO := make([]*model.UiTypeLogic, 0, len(typesDAO))
	for _, t := range typesDAO {
		typesDTO = append(typesDTO, &model.UiTypeLogic{
			Name:            t.Name,
			Comment:         t.Comment,
			MultipleOptions: t.MultipleOptions,
		})
	}
	return &model.GetTypesResponse{
		Elements: typesDTO,
	}, nil
}
