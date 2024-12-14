// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

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
        "/api/v1/cronjob/start": {
            "get": {
                "description": "Scheduler Start",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sms"
                ],
                "summary": "Scheduler Start",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/cronjob/stop": {
            "get": {
                "description": "Scheduler Stop",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sms"
                ],
                "summary": "Scheduler Stop",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/notifications/sms": {
            "get": {
                "description": "List Sms",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sms"
                ],
                "summary": "List Sms",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Sms"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "send sms",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SMS"
                ],
                "summary": "send sms",
                "parameters": [
                    {
                        "description": "SMS you want to create",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.SmsRequestV1"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "The newly created SMS",
                        "schema": {
                            "$ref": "#/definitions/entity.Sms"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Sms": {
            "type": "object",
            "properties": {
                "activeStatus": {
                    "type": "boolean"
                },
                "content": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string",
                    "example": "00000000-0000-0000-0000-000000000000"
                },
                "messageId": {
                    "type": "string"
                },
                "to": {
                    "type": "string"
                }
            }
        },
        "request.SmsRequestV1": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "to": {
                    "type": "string"
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
