// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

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
                                            "$ref": "#/definitions/github_com_adeyemialameen04_unwind-be_internal_domain.AuthResponse"
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
                "security": [
                    {
                        "RefreshTokenBearer": []
                    }
                ],
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
                            "$ref": "#/definitions/github_com_adeyemialameen04_unwind-be_internal_domain.RegisterRequest"
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
                                            "$ref": "#/definitions/github_com_adeyemialameen04_unwind-be_internal_domain.AuthResponse"
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
        "/user/profile": {
            "patch": {
                "security": [
                    {
                        "AccessTokenBearer": []
                    }
                ],
                "description": "Updates a user profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update Profile",
                "parameters": [
                    {
                        "description": "Profile Data",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_adeyemialameen04_unwind-be_internal_db_repository.UpdateProfileParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_adeyemialameen04_unwind-be_core_server.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/github_com_adeyemialameen04_unwind-be_internal_db_repository.Profile"
                                        }
                                    }
                                }
                            ]
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
                },
                "profileId": {
                    "type": "string"
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
        "github_com_adeyemialameen04_unwind-be_internal_db_repository.Profile": {
            "type": "object",
            "required": [
                "userId",
                "username"
            ],
            "properties": {
                "coverPic": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "profilePic": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                },
                "username": {
                    "type": "string",
                    "minLength": 8
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
        },
        "github_com_adeyemialameen04_unwind-be_internal_db_repository.UpdateProfileParams": {
            "type": "object",
            "required": [
                "username"
            ],
            "properties": {
                "coverPic": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "profilePic": {
                    "type": "string"
                },
                "username": {
                    "type": "string",
                    "minLength": 8
                }
            }
        },
        "github_com_adeyemialameen04_unwind-be_internal_domain.AuthResponse": {
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
        "github_com_adeyemialameen04_unwind-be_internal_domain.RegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
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
	BasePath:         "/api/v1",
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
