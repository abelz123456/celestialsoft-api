// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/auth/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "User Authentication",
                "parameters": [
                    {
                        "description": "Login info",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.PayloadLogin"
                        }
                    }
                ],
                "responses": {
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseProperties"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "User Registration",
                "parameters": [
                    {
                        "description": "Login info",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.PayloadRegister"
                        }
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseProperties"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseProperties"
                        }
                    }
                }
            }
        },
        "/bank": {
            "get": {
                "description": "Get Bank Collection",
                "tags": [
                    "Bank"
                ],
                "summary": "Get Bank Collection",
                "parameters": [
                    {
                        "type": "string",
                        "description": "With value 'Bearer {authToken}'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ok",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseProperties"
                        }
                    },
                    "400": {
                        "description": "General Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseProperties"
                        }
                    },
                    "401": {
                        "description": "Authentication Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseProperties"
                        }
                    }
                }
            },
            "post": {
                "description": "Create New Bank",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bank"
                ],
                "summary": "Create New Bank",
                "parameters": [
                    {
                        "type": "string",
                        "description": "With value 'Bearer {authToken}'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Bank Information",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.CreateBankDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseProperties"
                        }
                    },
                    "400": {
                        "description": "General Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseProperties"
                        }
                    },
                    "401": {
                        "description": "Authentication Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseProperties"
                        }
                    }
                }
            }
        },
        "/bank/{oid}": {
            "get": {
                "description": "Detail Bank",
                "tags": [
                    "Bank"
                ],
                "summary": "Detail Bank",
                "parameters": [
                    {
                        "type": "string",
                        "description": "oid of Bank",
                        "name": "oid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "With value 'Bearer {authToken}'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseProperties"
                        }
                    },
                    "400": {
                        "description": "General Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseProperties"
                        }
                    },
                    "401": {
                        "description": "Authentication Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseProperties"
                        }
                    }
                }
            },
            "delete": {
                "description": "Detail Bank",
                "tags": [
                    "Bank"
                ],
                "summary": "Detail Bank",
                "parameters": [
                    {
                        "type": "string",
                        "description": "oid of Bank",
                        "name": "oid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "With value 'Bearer {authToken}'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseProperties"
                        }
                    },
                    "400": {
                        "description": "General Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseProperties"
                        }
                    },
                    "401": {
                        "description": "Authentication Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseProperties"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update Bank",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bank"
                ],
                "summary": "Update Bank",
                "parameters": [
                    {
                        "type": "string",
                        "description": "oid of Bank",
                        "name": "oid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "With value 'Bearer {authToken}'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "New Bank Information",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UpdateBankDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseProperties"
                        }
                    },
                    "400": {
                        "description": "General Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseProperties"
                        }
                    },
                    "401": {
                        "description": "Authentication Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseProperties"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.CreateBankDto": {
            "type": "object",
            "required": [
                "bankCode",
                "bankName"
            ],
            "properties": {
                "bankCode": {
                    "type": "string"
                },
                "bankName": {
                    "type": "string"
                }
            }
        },
        "domain.PayloadLogin": {
            "type": "object",
            "required": [
                "emailName",
                "password"
            ],
            "properties": {
                "emailName": {
                    "type": "string",
                    "example": "me@mail.com"
                },
                "password": {
                    "type": "string",
                    "example": "securePassword"
                }
            }
        },
        "domain.PayloadRegister": {
            "type": "object",
            "required": [
                "emailName",
                "password"
            ],
            "properties": {
                "emailName": {
                    "type": "string",
                    "example": "me@mail.com"
                },
                "password": {
                    "type": "string",
                    "example": "securePassword"
                }
            }
        },
        "domain.UpdateBankDto": {
            "type": "object",
            "required": [
                "bankName"
            ],
            "properties": {
                "bankName": {
                    "type": "string"
                }
            }
        },
        "response.ResponseProperties": {
            "type": "object",
            "properties": {
                "data": {},
                "httpStatus": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "resultCode": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Celestialsoftware API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
