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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/image/v2/image": {
            "post": {
                "description": "Загрузить изображение и сохранить на директории сервера",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Image"
                ],
                "summary": "Загрузить изображение",
                "parameters": [
                    {
                        "description": "-",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/imagerouter.LoadFromNetDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "image url",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/image/v2/image/optimize": {
            "post": {
                "description": "Загрузить изображение и сохранить на директории сервера, сделать ресайз, и перевести в нужный формат",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Image"
                ],
                "summary": "Оптимизировать изображение",
                "parameters": [
                    {
                        "description": "-",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/imagerouter.OptimizeDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "image urls",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/image/v2/image/optimize/load": {
            "post": {
                "description": "Загрузить изображение, сделать ресайз, и перевести в нужный формат",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Image"
                ],
                "summary": "Загрузить и оптимизировать изображение",
                "parameters": [
                    {
                        "description": "Загрузить FormData c файлом",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/imagerouter.RequestLoadOptimize"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "файл",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/image/v2/image/resize": {
            "post": {
                "description": "Загрузить изображение и сохранить на директории сервера, сделать ресайз",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Image"
                ],
                "summary": "Изменить размер изображения",
                "parameters": [
                    {
                        "description": "-",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/imagerouter.LoadFromNetDtoAndResize"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "image url",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/image/v2/image/resize/convert": {
            "post": {
                "description": "Загрузить изображение и сохранить на директории сервера, сделать ресайз, и перевести в нужный формат",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Image"
                ],
                "summary": "Конвертировать изображение",
                "parameters": [
                    {
                        "description": "-",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/imagerouter.LoadFromNetDtoAndResizeAndConvert"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "image url",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "image.Point": {
            "type": "object",
            "required": [
                "format"
            ],
            "properties": {
                "format": {
                    "type": "string"
                },
                "height": {
                    "type": "integer"
                },
                "width": {
                    "type": "integer"
                }
            }
        },
        "imagerouter.LoadFromNetDto": {
            "type": "object",
            "required": [
                "url"
            ],
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        },
        "imagerouter.LoadFromNetDtoAndResize": {
            "type": "object",
            "required": [
                "url"
            ],
            "properties": {
                "height": {
                    "type": "integer"
                },
                "url": {
                    "type": "string"
                },
                "width": {
                    "type": "integer"
                }
            }
        },
        "imagerouter.LoadFromNetDtoAndResizeAndConvert": {
            "type": "object",
            "required": [
                "format",
                "url"
            ],
            "properties": {
                "format": {
                    "type": "string"
                },
                "height": {
                    "type": "integer"
                },
                "url": {
                    "type": "string"
                },
                "width": {
                    "type": "integer"
                }
            }
        },
        "imagerouter.OptimizeDto": {
            "type": "object",
            "required": [
                "points",
                "url"
            ],
            "properties": {
                "points": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/image.Point"
                    }
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "imagerouter.RequestLoadOptimize": {
            "type": "object",
            "properties": {
                "file": {
                    "type": "string"
                },
                "format": {
                    "type": "string"
                },
                "height": {
                    "type": "integer"
                },
                "width": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
