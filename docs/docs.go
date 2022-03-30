// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Hanmur",
            "url": "https://hanmur.cn/",
            "email": "wenyt8@mail2.edu.cn"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/article": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取多个文章",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章管理"
                ],
                "summary": "获取多个文章",
                "parameters": [
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "文章名称",
                        "name": "title",
                        "in": "query"
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "default": 1,
                        "description": "状态",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.Article"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "创建文章",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章管理"
                ],
                "summary": "创建文章",
                "parameters": [
                    {
                        "maxLength": 100,
                        "minLength": 1,
                        "type": "string",
                        "description": "文章标题",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "minLength": 1,
                        "type": "string",
                        "description": "文章描述",
                        "name": "desc",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "minLength": 1,
                        "type": "string",
                        "description": "文章内容",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "封面路径",
                        "name": "cover_image_url",
                        "in": "formData"
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "default": 1,
                        "description": "状态",
                        "name": "state",
                        "in": "formData"
                    },
                    {
                        "maxLength": 100,
                        "minLength": 1,
                        "type": "string",
                        "description": "创建者",
                        "name": "created_by",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/article/{id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "更新文章",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章管理"
                ],
                "summary": "更新文章",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "文章 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "文章标题",
                        "name": "title",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "文章描述",
                        "name": "desc",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "文章内容",
                        "name": "content",
                        "in": "formData"
                    },
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "封面路径",
                        "name": "cover_image_url",
                        "in": "formData"
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "default": 1,
                        "description": "状态",
                        "name": "state",
                        "in": "formData"
                    },
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "type": "string",
                        "description": "修改者",
                        "name": "modified_by",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Article"
                            }
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "删除文章",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章管理"
                ],
                "summary": "删除文章",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "文章 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/article/{title}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取单个文章",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章管理"
                ],
                "summary": "获取单个文章",
                "parameters": [
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "文章名称",
                        "name": "title",
                        "in": "query"
                    },
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "文章作者",
                        "name": "created_by",
                        "in": "query"
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "default": 1,
                        "description": "状态",
                        "name": "state",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.Article"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/tags": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取多个标签",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "标签管理"
                ],
                "summary": "获取多个标签",
                "parameters": [
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "标签名称",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "default": 1,
                        "description": "状态",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.Tag"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "创建一个新标签",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "标签管理"
                ],
                "summary": "新增标签",
                "parameters": [
                    {
                        "maxLength": 100,
                        "minLength": 1,
                        "type": "string",
                        "description": "标签名称",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "default": 1,
                        "description": "状态",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "maxLength": 100,
                        "minLength": 1,
                        "type": "string",
                        "description": "创建者",
                        "name": "created_by",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.Tag"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/tags/{id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "更新标签",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "标签管理"
                ],
                "summary": "更新标签",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "标签 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "minLength": 1,
                        "type": "string",
                        "description": "标签名称",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "default": 1,
                        "description": "状态",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "maxLength": 100,
                        "minLength": 1,
                        "type": "string",
                        "description": "修改者",
                        "name": "modified_by",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Tag"
                            }
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "删除标签",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "标签管理"
                ],
                "summary": "删除标签",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "标签 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/tags/{name}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取单个标签，未实现",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "标签管理"
                ],
                "summary": "获取单个标签",
                "parameters": [
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "标签名称",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "default": 1,
                        "description": "状态",
                        "name": "state",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.Tag"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "登录，获取Token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "账户管理"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Hanmur",
                        "description": "认证账号",
                        "name": "auth_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "Hanmur_Goose",
                        "description": "认证密码",
                        "name": "auth_code",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    }
                }
            }
        },
        "/auth/modifyCode": {
            "put": {
                "description": "修改密码",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "账户管理"
                ],
                "summary": "修改密码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "账号",
                        "name": "auth_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "原密码",
                        "name": "auth_code",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "新密码",
                        "name": "new_code",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "检验验证码和账号密码格式，进行登录",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "账户管理"
                ],
                "summary": "注册账号",
                "parameters": [
                    {
                        "type": "string",
                        "description": "账号",
                        "name": "auth_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "auth_code",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "验证码",
                        "name": "check_code",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "A response",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    },
                    "default": {
                        "description": "A response",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/resetCode": {
            "put": {
                "description": "检测验证码后重置密码",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "账户管理"
                ],
                "summary": "重置密码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "验证码",
                        "name": "check_code",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    }
                }
            }
        },
        "/auth/sendCheck": {
            "post": {
                "description": "在Redis中生成验证码并发送该验证码至Redis",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "账户管理"
                ],
                "summary": "发送验证码",
                "parameters": [
                    {
                        "type": "string",
                        "default": "1466046208@qq.com",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功"
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    }
                }
            }
        },
        "/upload/file": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "文件上传，目前仅支持图片",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "上传"
                ],
                "summary": "文件上传",
                "parameters": [
                    {
                        "type": "file",
                        "description": "文件路径",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "enum": [
                            1
                        ],
                        "type": "integer",
                        "description": "文件类型",
                        "name": "type",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errorCode.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errorCode.Error": {
            "type": "object"
        },
        "model.Article": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "cover_image_url": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "created_on": {
                    "type": "integer"
                },
                "deleted_on": {
                    "type": "integer"
                },
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_del": {
                    "type": "integer"
                },
                "modified_by": {
                    "type": "string"
                },
                "modified_on": {
                    "type": "integer"
                },
                "state": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.Tag": {
            "type": "object",
            "properties": {
                "created_by": {
                    "type": "string"
                },
                "created_on": {
                    "type": "integer"
                },
                "deleted_on": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "is_del": {
                    "type": "integer"
                },
                "modified_by": {
                    "type": "string"
                },
                "modified_on": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "state": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{"http"},
	Title:            "Goose谷声",
	Description:      "简单的API描述文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
