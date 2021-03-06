// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2021-09-27 05:35:38.185131974 -0300 -03 m=+0.032329308

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Carlos Nóbrega",
            "email": "nobreqacarlosjr@gmail.com"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/files": {
            "post": {
                "security": [
                    {
                        "userIdAuthentication": []
                    }
                ],
                "description": "Upload a new File",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "summary": "Upload a file",
                "parameters": [
                    {
                        "type": "file",
                        "description": "file to be uploaded",
                        "name": "files",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "Indicates if the file is public or private",
                        "name": "isPublic",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Gives a custom name to the file",
                        "name": "filename",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Indicates the file directory",
                        "name": "directory",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.FileResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    }
                }
            }
        },
        "/files/grep": {
            "get": {
                "security": [
                    {
                        "userIdAuthentication": []
                    }
                ],
                "description": "List files by passing a directory",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "summary": "List files by passing a directory",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Attribute to get files under a directory",
                        "name": "dir",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.FileResponse"
                        }
                    }
                }
            }
        },
        "/files/list/private": {
            "get": {
                "security": [
                    {
                        "userIdAuthentication": []
                    }
                ],
                "description": "List public files",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "summary": "List public files",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Limit of page records ",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Page number",
                        "name": "page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.FileResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    }
                }
            }
        },
        "/files/list/public": {
            "get": {
                "description": "List public files",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "summary": "List public files",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The identifier for the file's user",
                        "name": "user_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Limit of page records ",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Page number",
                        "name": "page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.FileResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    }
                }
            }
        },
        "/files/{id}": {
            "get": {
                "security": [
                    {
                        "userIdAuthentication": []
                    }
                ],
                "description": "Find a  File",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "summary": "Find a file",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The identifier for the file",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.FilePublic"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "userIdAuthentication": []
                    }
                ],
                "description": "Deletes a  File",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "summary": "Deletes a file",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The identifier for the file",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.FileList"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "userIdAuthentication": []
                    }
                ],
                "description": "Update a file directory",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "summary": "Update a file directory",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The identifier for the file",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The identifier for the file",
                        "name": "content",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usecases.UpdateFileDirectoryByIdParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.FilePublic"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    }
                }
            }
        },
        "/files/{id}/replace": {
            "patch": {
                "security": [
                    {
                        "userIdAuthentication": []
                    }
                ],
                "description": "Replace a blob from a existent file",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "summary": "Replace a blob from a existent file",
                "parameters": [
                    {
                        "type": "file",
                        "description": "file to be uploaded",
                        "name": "files",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "file to be replaced",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.FileResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    }
                }
            }
        },
        "/static/{fullfilepath}": {
            "get": {
                "security": [
                    {
                        "userIdAuthentication": []
                    }
                ],
                "description": "Download static file (doesnt working on swagger)",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "summary": "Download static file (doesnt working on swagger)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer \u003ctoken\u003e is required to download/view",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "router example: http://localhost:8090/static/example/updated/fix/splunk.pdf",
                        "name": "fullfilepath",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    }
                }
            }
        },
        "/users/": {
            "put": {
                "security": [
                    {
                        "userIdAuthentication": []
                    }
                ],
                "description": "Updates a user",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Updates a user",
                "parameters": [
                    {
                        "description": "Object for update the user",
                        "name": "content",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.User"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new user",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Creates a new user",
                "parameters": [
                    {
                        "description": "Object for creating the user",
                        "name": "content",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.User"
                        }
                    }
                }
            }
        },
        "/users/findone": {
            "get": {
                "description": "Find specific user",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Find specific user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Attribute to get specific user",
                        "name": "user_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.User"
                        }
                    }
                }
            }
        },
        "/users/list": {
            "get": {
                "description": "list users",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "list users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Users"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "Sign in a user",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Sign in a user",
                "parameters": [
                    {
                        "description": "Object for sign in the user",
                        "name": "content",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helpers.TokenResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.Directories": {
            "type": "object",
            "properties": {
                "directories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Directory"
                    }
                }
            }
        },
        "dto.Directory": {
            "type": "object",
            "properties": {
                "path": {
                    "type": "string"
                }
            }
        },
        "dto.FileList": {
            "type": "object",
            "properties": {
                "files": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "dto.FilePublic": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "directory": {
                    "type": "string"
                },
                "extension": {
                    "type": "string"
                },
                "fileId": {
                    "type": "string"
                },
                "fileSize": {
                    "type": "string"
                },
                "isPublic": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "dto.FileResponse": {
            "type": "object",
            "properties": {
                "files": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.FilePublic"
                    }
                }
            }
        },
        "dto.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "dto.UserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.Users": {
            "type": "object",
            "properties": {
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.User"
                    }
                }
            }
        },
        "helpers.TokenResponse": {
            "type": "object",
            "properties": {
                "acessToken": {
                    "type": "string"
                },
                "expiryDate": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "routes.file": {
            "$ref": "#/definitions/io.Writer"
        },
        "usecases.UpdateFileDirectoryByIdParam": {
            "type": "object",
            "properties": {
                "fileId": {
                    "type": "string"
                },
                "newDirectory": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "userIdAuthentication": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0.0",
	Host:        "localhost:8090",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "File API",
	Description: "API documentation for File API",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
