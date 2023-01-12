// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplateauth = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login Schema ",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.LoginResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "schema.Login": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "schema.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfoauth holds exported Swagger Info so clients can modify it
var SwaggerInfoauth = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:8080",
	BasePath:         "/auth",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server.",
	InfoInstanceName: "auth",
	SwaggerTemplate:  docTemplateauth,
}

func init() {
	swag.Register(SwaggerInfoauth.InstanceName(), SwaggerInfoauth)
}
