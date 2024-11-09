// Package __parseDependencyLevel Code generated by swaggo/swag. DO NOT EDIT
package __parseDependencyLevel

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Al-Ameen Adeyemi",
            "url": "https://github.com/adeyemialameen04",
            "email": "adeyemialameen04@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "Logs a user into his/her account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login to your account",
                "parameters": [
                    {
                        "description": "Login data",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_adeyemialameen04_unwind-be_internal_db_repository.RegisterUserParams"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Login success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_adeyemialameen04_unwind-be_core_server.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/core_handlers_auth.RegisterResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid request data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Failed to start transaction or insert book",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/auth/refresh": {
            "get": {
                "description": "Refreshes token to get new token pair",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Refreh Token",
                "responses": {}
            }
        },
        "/auth/signup": {
            "post": {
                "description": "Create an account on unwind",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Create an account",
                "parameters": [
                    {
                        "description": "Signup data",
                        "name": "EmailAndPassword",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_adeyemialameen04_unwind-be_internal_db_repository.RegisterUserParams"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User created successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_adeyemialameen04_unwind-be_core_server.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/core_handlers_auth.RegisterResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid request data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Failed to start transaction or insert book",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/books": {
            "get": {
                "description": "Get a list of all books",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Get all books",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Books retrieved successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_adeyemialameen04_unwind-be_core_server.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/github_com_adeyemialameen04_unwind-be_internal_db_repository.Book"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "AccessTokenBearer": []
                    }
                ],
                "description": "Insert a new book into the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Create a new book",
                "parameters": [
                    {
                        "description": "Book data",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_adeyemialameen04_unwind-be_internal_db_repository.InsertBookParams"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Book created successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_adeyemialameen04_unwind-be_core_server.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/github_com_adeyemialameen04_unwind-be_internal_db_repository.Book"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid request data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Failed to start transaction or insert book",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "EmailID": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "core_handlers_auth.RegisterResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/EmailID"
                }
            }
        },
        "github_com_adeyemialameen04_unwind-be_core_server.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "errors": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "github_com_adeyemialameen04_unwind-be_internal_db_repository.Book": {
            "type": "object",
            "required": [
                "author",
                "title"
            ],
            "properties": {
                "author": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "github_com_adeyemialameen04_unwind-be_internal_db_repository.InsertBookParams": {
            "type": "object",
            "required": [
                "author",
                "title"
            ],
            "properties": {
                "author": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "github_com_adeyemialameen04_unwind-be_internal_db_repository.RegisterUserParams": {
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
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "AccessTokenBearer": {
            "description": "AccessTokenBearer Authentication",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "RefreshTokenBearer": {
            "description": "RefreshTokenBearer Authentication",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
    "tags": [
        {
            "description": "Authentication endpoints",
            "name": "Auth"
        }
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Unwind API",
	Description:      "API for Unwind",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
