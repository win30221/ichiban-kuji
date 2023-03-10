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
        "/ichiban-kuji/": {
            "post": {
                "security": [
                    {
                        "Systoken": []
                    }
                ],
                "summary": "建立一番賞",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "cost",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "isLast",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "lastType",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "rewardCount",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "sellAt",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "userID",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/ichiban-kuji/reward/{ichibanKujiID}": {
            "post": {
                "security": [
                    {
                        "Systoken": []
                    }
                ],
                "summary": "建立一番賞獎品",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "name": "name",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/ichiban-kuji/{ichibanKujiID}": {
            "patch": {
                "security": [
                    {
                        "Systoken": []
                    }
                ],
                "summary": "修改一番賞",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "cost",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "name": "isLast",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "name": "lastType",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "name": "rewardCount",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "sellAt",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "name": "status",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/ichiban-kuji/{ichibanKujiRewardID}": {
            "patch": {
                "security": [
                    {
                        "Systoken": []
                    }
                ],
                "summary": "修改一番賞獎品內容",
                "parameters": [
                    {
                        "type": "string",
                        "name": "description",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "userID",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/ichiban-kuji/{userID}": {
            "get": {
                "security": [
                    {
                        "Systoken": []
                    }
                ],
                "summary": "取得一番賞獎品列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": " ",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.IchibanKuji"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.IchibanKuji": {
            "type": "object",
            "properties": {
                "cost": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isLast": {
                    "type": "integer"
                },
                "lastType": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "rewardCount": {
                    "type": "integer"
                },
                "sellAt": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Systoken": {
            "type": "apiKey",
            "name": "Systoken",
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "ichiban-kuji 模組",
	Description: "使用者管理",
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
