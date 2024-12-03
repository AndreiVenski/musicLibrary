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
            "name": "Andrei Venski",
            "url": "https://github.com/andrew967",
            "email": "venskiandrei32@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/info": {
            "post": {
                "description": "Retrieve information about the music library with optional pagination.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Library"
                ],
                "summary": "Get library information",
                "parameters": [
                    {
                        "description": "Filter parameters",
                        "name": "filter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SongFullDataWithLimitAndOffsetRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of songs",
                        "schema": {
                            "$ref": "#/definitions/models.SongsResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/music": {
            "post": {
                "description": "Add a new song to the music library.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Library"
                ],
                "summary": "Add a new song",
                "parameters": [
                    {
                        "description": "Song data to add",
                        "name": "songData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SongRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Added song information",
                        "schema": {
                            "$ref": "#/definitions/models.SongResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a song from the library using song details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Library"
                ],
                "summary": "Delete a song",
                "parameters": [
                    {
                        "description": "Song to delete",
                        "name": "songData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SongRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Song deleted successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Invalid request data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/music/id": {
            "delete": {
                "description": "Delete a song from the library using its unique ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Library"
                ],
                "summary": "Delete a song by ID",
                "parameters": [
                    {
                        "description": "Song ID to delete",
                        "name": "songID",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SongID"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Song deleted successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Invalid request data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/music/text": {
            "post": {
                "description": "Retrieve verse information for a specific song.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Library"
                ],
                "summary": "Get music text information",
                "parameters": [
                    {
                        "description": "Verse request",
                        "name": "verseReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.VerseRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Verse information",
                        "schema": {
                            "$ref": "#/definitions/models.VerseResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/music/update": {
            "put": {
                "description": "Update details of a song in the library.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Library"
                ],
                "summary": "Update music information",
                "parameters": [
                    {
                        "description": "Song data to update",
                        "name": "updateData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SongFullDataRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated song information",
                        "schema": {
                            "$ref": "#/definitions/models.SongResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/music/update/id": {
            "put": {
                "description": "Update details of a song in the library using its unique ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Library"
                ],
                "summary": "Update music information by ID",
                "parameters": [
                    {
                        "description": "Song data to update",
                        "name": "updateData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SongFullDataRequestWithID"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated song information",
                        "schema": {
                            "$ref": "#/definitions/models.SongResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.SongFullDataRequest": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "releaseDate": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "models.SongFullDataRequestWithID": {
            "type": "object",
            "required": [
                "songId"
            ],
            "properties": {
                "group": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "releaseDate": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                },
                "songId": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "models.SongFullDataWithLimitAndOffsetRequest": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "limit": {
                    "type": "integer"
                },
                "link": {
                    "type": "string"
                },
                "offset": {
                    "type": "integer"
                },
                "releaseDate": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "models.SongID": {
            "type": "object",
            "properties": {
                "songId": {
                    "type": "integer"
                }
            }
        },
        "models.SongRequest": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                }
            }
        },
        "models.SongResponse": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "releaseDate": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                },
                "songId": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "models.SongsResponse": {
            "type": "object",
            "properties": {
                "songs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SongsResponse"
                    }
                }
            }
        },
        "models.VerseRequest": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                },
                "verseID": {
                    "type": "integer"
                }
            }
        },
        "models.VerseResponse": {
            "type": "object",
            "properties": {
                "verse": {
                    "type": "string"
                },
                "verseID": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/lib",
	Schemes:          []string{},
	Title:            "Song library API",
	Description:      "This is API for song library",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}