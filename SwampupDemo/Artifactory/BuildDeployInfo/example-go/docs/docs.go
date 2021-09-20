// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

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
        "termsOfService": "https://jfrog.com/eula/",
        "contact": {
            "name": "JFrog Support",
            "url": "https://jfrog.com/support/",
            "email": "partner-support@jfrog.com"
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
        "/bad": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    },
                    {
                        "OAuth2Application": [
                            "write",
                            "admin"
                        ]
                    }
                ],
                "description": "Example failure endpoint",
                "produces": [
                    "application/json"
                ],
                "summary": "Example failure endpoint",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.Bad"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.Error"
                        }
                    }
                }
            }
        },
        "/good": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    },
                    {
                        "OAuth2Application": [
                            "write",
                            "admin"
                        ]
                    }
                ],
                "description": "Example successful endpoint",
                "produces": [
                    "application/json"
                ],
                "summary": "Example successful endpoint",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.Good"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "server.Bad": {
            "type": "object",
            "properties": {
                "failureId": {
                    "type": "string"
                },
                "reason": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "server.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "server.Good": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "OAuth2Application": {
            "type": "oauth2",
            "flow": "application",
            "authorizationUrl": "",
            "tokenUrl": "/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "write": " Grants write access"
            }
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
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Example Go API",
	Description: "Example Go API to demo JFrog Partner Integrations",
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
