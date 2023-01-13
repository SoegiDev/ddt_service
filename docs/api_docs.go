// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplateapi = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Fajar soegi",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user/:ID": {
            "get": {
                "description": "Profile Meta User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "API SERVICE",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Search User By ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.MetaUser"
                        }
                    }
                }
            }
        },
        "/user/:ID/assign_roles": {
            "get": {
                "description": "User Assign Role User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "User Assign Role",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Assign User By ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Assign User Schema ",
                        "name": "Register",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.AssignUserRole"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.MsgResponse"
                        }
                    }
                }
            }
        },
        "/user/employee": {
            "post": {
                "description": "Employee Create New",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Employee Add New",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Employee Add New Schema ",
                        "name": "Register",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.JsonEmployeeCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.MsgResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "schema.AssignUserRole": {
            "type": "object",
            "properties": {
                "role": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "schema.JsonEmployeeCreate": {
            "type": "object",
            "required": [
                "address",
                "application",
                "code",
                "company_id",
                "department",
                "division",
                "email",
                "fullname",
                "nickname",
                "nik",
                "office_number",
                "password",
                "phone_number",
                "picture",
                "role",
                "role_application",
                "username"
            ],
            "properties": {
                "address": {
                    "type": "string"
                },
                "application": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "code": {
                    "type": "string"
                },
                "company_id": {
                    "type": "string"
                },
                "department": {
                    "type": "string"
                },
                "division": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "nik": {
                    "type": "string"
                },
                "office_number": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                },
                "role": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "role_application": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "schema.MetaAccount": {
            "type": "object",
            "properties": {
                "application": {
                    "$ref": "#/definitions/schema.MetaApp"
                },
                "id": {
                    "type": "string"
                },
                "roleApplication": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.MetaRoleApp"
                    }
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "schema.MetaActivities": {
            "type": "object",
            "properties": {
                "application": {
                    "type": "string"
                },
                "client": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "duration": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "method": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "pathOp": {
                    "type": "string"
                },
                "referrer": {
                    "type": "string"
                },
                "reqBody": {
                    "type": "string"
                },
                "requestId": {
                    "type": "string"
                },
                "source": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "schema.MetaApp": {
            "type": "object",
            "properties": {
                "appPackage": {
                    "type": "string"
                },
                "appPackageClass": {
                    "type": "string"
                },
                "assetApk": {
                    "type": "string"
                },
                "assetIcon": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "isActive": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "updatedNote": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "schema.MetaCompany": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "code": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "domain": {
                    "type": "string"
                },
                "fax": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "sector": {
                    "type": "string"
                }
            }
        },
        "schema.MetaDivision": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "gang": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.MetaGang"
                    }
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "schema.MetaEmployee": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "code": {
                    "type": "string"
                },
                "company": {
                    "$ref": "#/definitions/schema.MetaCompany"
                },
                "companyId": {
                    "type": "string"
                },
                "department": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "expired": {
                    "type": "integer"
                },
                "fullName": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "nickName": {
                    "type": "string"
                },
                "nik": {
                    "type": "string"
                },
                "officeNumber": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "schema.MetaEstate": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "company": {
                    "$ref": "#/definitions/schema.MetaCompany"
                },
                "description": {
                    "type": "string"
                },
                "division": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.MetaDivision"
                    }
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "schema.MetaGang": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "schema.MetaRoleApp": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "schema.MetaRoles": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "schema.MetaUser": {
            "type": "object",
            "properties": {
                "accounts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.MetaAccount"
                    }
                },
                "activityLogs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.MetaActivities"
                    }
                },
                "code": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "employees": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.MetaEmployee"
                    }
                },
                "estates": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.MetaEstate"
                    }
                },
                "id": {
                    "type": "string"
                },
                "isActive": {
                    "type": "boolean"
                },
                "password": {
                    "type": "string"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.MetaRoles"
                    }
                },
                "updatedAt": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "schema.MsgResponse": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfoapi holds exported Swagger Info so clients can modify it
var SwaggerInfoapi = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Swagger User Service",
	Description:      "This is User Service Test.",
	InfoInstanceName: "api",
	SwaggerTemplate:  docTemplateapi,
}

func init() {
	swag.Register(SwaggerInfoapi.InstanceName(), SwaggerInfoapi)
}
