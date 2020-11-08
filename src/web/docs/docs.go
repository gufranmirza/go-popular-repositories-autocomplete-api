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
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/health": {
            "get": {
                "description": "It returns the health of the service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Get health of the service",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/healthinterface.Health"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errorinterface.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errorinterface.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errorinterface.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errorinterface.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "application-specific error code",
                    "type": "integer"
                },
                "error": {
                    "description": "application-level error message, for debugging",
                    "type": "string"
                },
                "status": {
                    "description": "user-level status message",
                    "type": "string"
                }
            }
        },
        "healthinterface.Health": {
            "type": "object",
            "properties": {
                "inboundInterfaces": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/healthinterface.InboundInterface"
                    }
                },
                "outboundInterfaces": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/healthinterface.OutboundInterface"
                    }
                },
                "serviceName": {
                    "type": "string"
                },
                "serviceProvider": {
                    "type": "string"
                },
                "serviceStartTimeUTC": {
                    "type": "string"
                },
                "serviceStatus": {
                    "type": "string"
                },
                "serviceVersion": {
                    "type": "string"
                },
                "timeStampUTC": {
                    "type": "string"
                },
                "uptime": {
                    "type": "number"
                }
            }
        },
        "healthinterface.InboundInterface": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "applicationName": {
                    "type": "string"
                },
                "connectionStatus": {
                    "type": "string"
                },
                "hostname": {
                    "type": "string"
                },
                "os": {
                    "type": "string"
                },
                "timeStampUTC": {
                    "type": "string"
                }
            }
        },
        "healthinterface.OutboundInterface": {
            "type": "object",
            "properties": {
                "applicationName": {
                    "type": "string"
                },
                "connectionStatus": {
                    "type": "string"
                },
                "timeStampUTC": {
                    "type": "string"
                },
                "urls": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
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
	Version:     "2.0",
	Host:        "localhost:8001",
	BasePath:    "/go-popular/v1",
	Schemes:     []string{},
	Title:       "Go Popular API Documentation",
	Description: "Go Popular API Documentation",
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
