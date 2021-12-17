// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/v1/adjacenttable": {
            "post": {
                "description": "Create adjacent table\nyou can create sector with this method just add description and coords to sector field\nalso you can just add coords fields and they will find sector\nthis endpoint also can get or create institute/profile/direction by name, because all names in this object is unique string\nif adjacent table with this sector and variant exist return bad request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/adjacenttable.CreateAdjacentTableReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/adjacenttable.CreateAdjacentTableResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/adjacenttables": {
            "post": {
                "description": "Create adjacent tables\nthis method create or institute/profile/direction but require created sector in array\nif adjacent table with this sector and variant exist return bad request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/adjacenttable.CreateAdjacentTablesReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/adjacenttable.CreateAdjacentTablesResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/direction": {
            "get": {
                "description": "return all directions",
                "produces": [
                    "application/json"
                ],
                "summary": "Get all directions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/direction.GetDirectionsResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/direction/{id}": {
            "delete": {
                "description": "Delete Direction by id",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete Direction by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of institute",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/institute": {
            "get": {
                "description": "return all institutes",
                "produces": [
                    "application/json"
                ],
                "summary": "Get all institutes",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/institute.GetInstitutesResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/institute/{id}": {
            "delete": {
                "description": "Delete Institute by id",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete Institute by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of institute",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/profile": {
            "get": {
                "description": "return all profiles",
                "produces": [
                    "application/json"
                ],
                "summary": "Get all profiles",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/profile.GetAllProfilesResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/profile/{id}": {
            "delete": {
                "description": "Delete profile by id",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete profile by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of profile",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/sector": {
            "get": {
                "description": "return all sectors",
                "produces": [
                    "application/json"
                ],
                "summary": "Get Sectors",
                "parameters": [
                    {
                        "type": "string",
                        "description": "institute name",
                        "name": "institute",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "direction name",
                        "name": "direction",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "profile name",
                        "name": "profile",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/sector.GetAllSectorsResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "create sector according to giving coords\ncoords is unique string",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create Sector",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/sector.CreateSectorReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/sector.Sector"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/sector/{id}": {
            "put": {
                "description": "update sector",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update Sector",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of sector",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/sector.UpdateSectorReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/sector.Sector"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "adjacenttable.AdjacentTable": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "sector": {
                    "$ref": "#/definitions/adjacenttable.Sector"
                },
                "variant": {
                    "$ref": "#/definitions/adjacenttable.Variant"
                }
            }
        },
        "adjacenttable.CreateAdjacentTableReq": {
            "type": "object",
            "properties": {
                "directionName": {
                    "type": "string"
                },
                "instituteName": {
                    "type": "string"
                },
                "profileName": {
                    "type": "string"
                },
                "sector": {
                    "$ref": "#/definitions/adjacenttable.CreateSectorReq"
                }
            }
        },
        "adjacenttable.CreateAdjacentTableResp": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "sector": {
                    "$ref": "#/definitions/adjacenttable.Sector"
                },
                "variant": {
                    "$ref": "#/definitions/adjacenttable.Variant"
                }
            }
        },
        "adjacenttable.CreateAdjacentTablesReq": {
            "type": "object",
            "properties": {
                "directionName": {
                    "type": "string"
                },
                "instituteName": {
                    "type": "string"
                },
                "profileName": {
                    "type": "string"
                },
                "sectors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "adjacenttable.CreateAdjacentTablesResp": {
            "type": "object",
            "properties": {
                "adjacentTables": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/adjacenttable.AdjacentTable"
                    }
                }
            }
        },
        "adjacenttable.CreateSectorReq": {
            "type": "object",
            "properties": {
                "coords": {
                    "type": "string"
                },
                "description": {
                    "type": "string",
                    "x-nullable": true
                }
            }
        },
        "adjacenttable.Sector": {
            "type": "object",
            "properties": {
                "coords": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "adjacenttable.Variant": {
            "type": "object",
            "properties": {
                "direction": {
                    "$ref": "#/definitions/variant.Direction"
                },
                "id": {
                    "type": "integer"
                },
                "institute": {
                    "$ref": "#/definitions/variant.Institute"
                },
                "profile": {
                    "$ref": "#/definitions/variant.Profile"
                }
            }
        },
        "direction.Direction": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "direction.GetDirectionsResp": {
            "type": "object",
            "properties": {
                "directions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/direction.Direction"
                    }
                }
            }
        },
        "institute.GetInstitutesResp": {
            "type": "object",
            "properties": {
                "institutes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/institute.Institute"
                    }
                }
            }
        },
        "institute.Institute": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "profile.GetAllProfilesResp": {
            "type": "object",
            "properties": {
                "profiles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/profile.Profile"
                    }
                }
            }
        },
        "profile.Profile": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "sector.CreateSectorReq": {
            "type": "object",
            "properties": {
                "coords": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                }
            }
        },
        "sector.GetAllSectorsResp": {
            "type": "object",
            "properties": {
                "sectors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/sector.Sector"
                    }
                }
            }
        },
        "sector.Sector": {
            "type": "object",
            "properties": {
                "coords": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "sector.UpdateSectorReq": {
            "type": "object",
            "properties": {
                "coords": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                }
            }
        },
        "variant.Direction": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "variant.Institute": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "variant.Profile": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "/api/magic-circle",
	Schemes:     []string{},
	Title:       "Magic-Circle API",
	Description: "This is a server to get projects from github",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
