// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/url": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "operationId": "create-entities.Tashilcar",
                "parameters": [
                    {
                        "description": "entities.Tashilcar data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.Tashilcar"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Tashilcar"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            }
        },
        "/url/:id": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "enable cheack health a entities.Tashilcar item by ID",
                "operationId": "enable-check-health-entities.Tashilcar-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "entities.Tashilcar.ID ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "delete a entities.Tashilcar item by ID",
                "operationId": "delete-entities.Tashilcar-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "entities.Tashilcar.ID ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    }
                }
            }
        },
        "/urls": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "get all items in the entities.Tashilcar list",
                "operationId": "get-all-entities.Tashilcar",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.Tashilcar"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entities.Tashilcar": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "integer"
                },
                "health_check": {
                    "type": "boolean"
                },
                "health_check_time_interval": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "request_body": {
                    "type": "string"
                },
                "request_headers": {
                    "type": "string"
                },
                "request_http_method": {
                    "type": "string"
                },
                "request_url": {
                    "type": "string"
                },
                "response_status": {
                    "type": "string"
                },
                "updated_at": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
