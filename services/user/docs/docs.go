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
        "/v1/login/email": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AuthByEmailRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "success response",
                        "schema": {
                            "$ref": "#/definitions/response.AuthResponse"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/v1/login/phone": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AuthByPhoneRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "success response",
                        "schema": {
                            "$ref": "#/definitions/response.AuthResponse"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/v1/register/email": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AuthByEmailRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "success response",
                        "schema": {
                            "$ref": "#/definitions/response.AuthResponse"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/v1/register/phone": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AuthByPhoneRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "success response",
                        "schema": {
                            "$ref": "#/definitions/response.AuthResponse"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/v1/user": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserProfile"
                ],
                "responses": {
                    "200": {
                        "description": "success response",
                        "schema": {
                            "$ref": "#/definitions/response.UserResponse"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserProfile"
                ],
                "parameters": [
                    {
                        "description": "Update Profile Payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateUserProfileRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success response",
                        "schema": {
                            "$ref": "#/definitions/response.UserResponse"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/v1/user/link/email": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserProfile"
                ],
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LinkEmailRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success response",
                        "schema": {
                            "$ref": "#/definitions/response.UserResponse"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/v1/user/link/phone": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserProfile"
                ],
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LinkPhoneRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success response",
                        "schema": {
                            "$ref": "#/definitions/response.UserResponse"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.AuthByEmailRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 8
                }
            }
        },
        "request.AuthByPhoneRequest": {
            "type": "object",
            "required": [
                "password",
                "phone"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 8
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "request.LinkEmailRequest": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "request.LinkPhoneRequest": {
            "type": "object",
            "required": [
                "phone"
            ],
            "properties": {
                "phone": {
                    "type": "string"
                }
            }
        },
        "request.UpdateUserProfileRequest": {
            "type": "object",
            "required": [
                "bankAccountHolder",
                "bankAccountName",
                "bankAccountNumber"
            ],
            "properties": {
                "bankAccountHolder": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 4
                },
                "bankAccountName": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 4
                },
                "bankAccountNumber": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 4
                },
                "fileId": {
                    "type": "string"
                }
            }
        },
        "response.AuthResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "response.UserResponse": {
            "type": "object",
            "properties": {
                "bankAccountHolder": {
                    "type": "string"
                },
                "bankAccountName": {
                    "type": "string"
                },
                "bankAccountNumber": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "fileId": {
                    "type": "string"
                },
                "fileThumbnailUri": {
                    "type": "string"
                },
                "fileUri": {
                    "type": "string"
                },
                "phone": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
