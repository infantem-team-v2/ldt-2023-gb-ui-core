// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Docs developer",
            "url": "https://t.me/KlenoviySirop",
            "email": "KlenoviySir@yandex.ru"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/calc/element/active": {
            "get": {
                "description": "Get active UI elements for calculator",
                "tags": [
                    "Calculator"
                ],
                "summary": "Get active UI elements",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GetActiveElementsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            },
            "patch": {
                "description": "Set state of activity for element",
                "tags": [
                    "Calculator",
                    "Admin"
                ],
                "summary": "Set active/inactive state for element",
                "parameters": [
                    {
                        "description": "Fields and their states",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SetActiveForElementRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/calc/types": {
            "get": {
                "description": "Get possible UI elements for calculator",
                "tags": [
                    "Calculator"
                ],
                "summary": "Get UI types for calculator (soon deprecated)",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GetTypesResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.Response": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "internal_code": {
                    "type": "integer"
                }
            }
        },
        "model.GetActiveElementsResponse": {
            "type": "object",
            "properties": {
                "categories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.UiCategoryLogic"
                    }
                }
            }
        },
        "model.GetTypesResponse": {
            "type": "object",
            "properties": {
                "elements": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.UiTypeLogic"
                    }
                }
            }
        },
        "model.SetActiveForElementRequest": {
            "type": "object",
            "properties": {
                "elements": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.UiChangeElementLogic"
                    }
                }
            }
        },
        "model.UiCategoryLogic": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "category_id": {
                    "type": "string"
                },
                "elements": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.UiElementLogic"
                    }
                }
            }
        },
        "model.UiChangeElementLogic": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "field_id": {
                    "type": "string"
                }
            }
        },
        "model.UiElementLogic": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "comment": {
                    "type": "string"
                },
                "field": {
                    "type": "string"
                },
                "field_id": {
                    "type": "string"
                },
                "options": {
                    "type": "array",
                    "items": {}
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "model.UiTypeLogic": {
            "type": "object",
            "properties": {
                "hint": {
                    "type": "string"
                },
                "multiple_options": {
                    "type": "boolean"
                },
                "type": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "ui.gb.ldt2023.infantem.tech",
	BasePath:         "",
	Schemes:          []string{"https"},
	Title:            "Backend-Driven-UI",
	Description:      "Service to provide UI specification frontend from backend",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
