// Package docs GENERATED BY SWAG; DO NOT EDIT
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
            "name": "Emanuel Vilela",
            "email": "evs10@aluno.ifal.edu.br"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/anonymous/authenticate": {
            "get": {
                "description": "Rota que permite que um usuário se autentique no Echofy com seus dados de sua conta do Spotify.\t\t\t\t\t\t\t\t\t\t  |",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rotas anônimas"
                ],
                "summary": "Fazer a autenticação no sistema",
                "operationId": "Login",
                "responses": {
                    "400": {
                        "description": "Requisição mal formulada.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "401": {
                        "description": "Usuário não autorizado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "403": {
                        "description": "Acesso negado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "422": {
                        "description": "Algum dado informado não pôde ser processado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Ocorreu um erro inesperado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "503": {
                        "description": "A base de dados não está disponível.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/user/playlist/{playlistID}": {
            "get": {
                "description": "Rota que permite que se busque todas as informações de uma playlist",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rotas do usuário"
                ],
                "summary": "Buscar os dados de uma playlist",
                "operationId": "GetPlaylistID",
                "parameters": [
                    {
                        "type": "string",
                        "default": "7pCvSVfjcnOw6AFJNZZ4bN",
                        "description": "ID da playlist.",
                        "name": "playlistID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Requisição realizada com sucesso.",
                        "schema": {
                            "$ref": "#/definitions/response.PlaylistDTO"
                        }
                    },
                    "401": {
                        "description": "Usuário não autorizado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "403": {
                        "description": "Acesso negado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "422": {
                        "description": "Algum dado informado não pôde ser processado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Ocorreu um erro inesperado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "503": {
                        "description": "A base de dados não está disponível.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/user/playlist/{playlistID}/songs": {
            "get": {
                "description": "Rota que permite que se busque todas as músicas de uma determinada playlist",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rotas do usuário"
                ],
                "summary": "Buscar todas as músicas de uma playlist",
                "operationId": "GetSongsByPlaylistID",
                "parameters": [
                    {
                        "type": "string",
                        "default": "7pCvSVfjcnOw6AFJNZZ4bN",
                        "description": "ID da playlist.",
                        "name": "playlistID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Requisição realizada com sucesso.",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.SongDTO"
                            }
                        }
                    },
                    "401": {
                        "description": "Usuário não autorizado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "403": {
                        "description": "Acesso negado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "422": {
                        "description": "Algum dado informado não pôde ser processado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Ocorreu um erro inesperado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "503": {
                        "description": "A base de dados não está disponível.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "response.ArtistDTO": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "founded_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "members": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.ArtistDTO"
                    }
                },
                "name": {
                    "type": "string"
                },
                "spotify_url": {
                    "type": "string"
                },
                "super_artist_id": {
                    "type": "string"
                },
                "terminated_at": {
                    "type": "string"
                }
            }
        },
        "response.ErrorMessage": {
            "type": "object",
            "properties": {
                "duplicated_fields": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "invalid_fields": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.InvalidField"
                    }
                },
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "response.InvalidField": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "field_name": {
                    "type": "string"
                }
            }
        },
        "response.PlaylistDTO": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "followers_amount": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "song_amount": {
                    "type": "integer"
                }
            }
        },
        "response.SongDTO": {
            "type": "object",
            "properties": {
                "album_id": {
                    "type": "string"
                },
                "artists": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.ArtistDTO"
                    }
                },
                "duration": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "release_date": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "bearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "ECHOFY API",
	Description:      "Aplicação de artistas do spotify",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
