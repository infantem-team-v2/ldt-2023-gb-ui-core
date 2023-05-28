package http

import (
	calculatorInterface "gb-ui-core/internal/calculator/interface"
	"gb-ui-core/internal/calculator/model"
	"gb-ui-core/internal/pkg/common"
	"gb-ui-core/pkg/terrors"
	"gb-ui-core/pkg/thttp/server"
	"github.com/gofiber/fiber/v2"
)

type CalculatorHandler struct {
	CalculatorUC calculatorInterface.UseCase `di:"calcUC"`
	prefix       string
	router       fiber.Router
}

func NewCalculatorHandler(app *fiber.App) server.IHandler {
	return &CalculatorHandler{
		prefix: "calc",
		router: app.Group("calc"),
	}
}

func (ch *CalculatorHandler) GetRouter() fiber.Router {
	return ch.router
}

func (ch *CalculatorHandler) GetPrefix() string {
	return ch.prefix
}

// GetActiveElements godoc
// @Summary Get active UI elements
// @Description Get active UI elements for calculator
// @Tags Calculator
// @Success 200 {object} model.GetActiveElementsResponse
// @Failure 400 {object} common.Response
// @Failure 401 {object} common.Response
// @Failure 403 {object} common.Response
// @Router /calc/element/active [get]
func (ch *CalculatorHandler) GetActiveElements() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		source := ctx.Query("source", "")
		doAdmin := false
		if source == "admin" {
			doAdmin = true
		}
		response, err := ch.CalculatorUC.GetActiveElements(doAdmin)
		if err != nil {
			return err
		}

		return ctx.JSON(response)
	}
}

// UpdateActiveElements godoc
// @Summary Set active/inactive state for element
// @Description Set state of activity for element
// @Tags Calculator, Admin
// @Param data body model.SetActiveForElementRequest true "Fields and their states"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Failure 401 {object} common.Response
// @Failure 403 {object} common.Response
// @Failure 422 {object} common.Response
// @Failure 409 {object} common.Response
// @Router /calc/element/active [patch]
func (ch *CalculatorHandler) UpdateActiveElements() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params model.SetActiveForElementRequest
		if err := server.ReadRequest(ctx, &params); err != nil {
			return terrors.Raise(err, 100001)
		}

		err := ch.CalculatorUC.UpdateActiveElements(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(common.Response{
			InternalCode: 200,
			Message:      "Successfully updated elements",
		})
	}
}

// GetTypes godoc
// @Summary Get UI types for calculator (soon deprecated)
// @Description Get possible UI elements for calculator
// @Tags Calculator
// @Success 200 {object} model.GetTypesResponse
// @Failure 400 {object} common.Response
// @Failure 401 {object} common.Response
// @Failure 403 {object} common.Response
// @Router /calc/types [get]
func (ch *CalculatorHandler) GetTypes() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		response, err := ch.CalculatorUC.GetTypes()
		if err != nil {
			return err
		}

		return ctx.JSON(response)
	}
}
