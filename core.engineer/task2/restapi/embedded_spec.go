// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a demo server for exercise 2.",
    "title": "demo.",
    "version": "1.0.0"
  },
  "basePath": "/",
  "paths": {
    "/base64/decode": {
      "post": {
        "consumes": [
          "text/plain"
        ],
        "produces": [
          "application/octet-stream"
        ],
        "summary": "base64 decode the supplied string",
        "parameters": [
          {
            "description": "base64 encoded string",
            "name": "base64String",
            "in": "body",
            "required": true,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "string",
              "format": "binary"
            }
          },
          "400": {
            "description": "Bad request"
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a demo server for exercise 2.",
    "title": "demo.",
    "version": "1.0.0"
  },
  "basePath": "/",
  "paths": {
    "/base64/decode": {
      "post": {
        "consumes": [
          "text/plain"
        ],
        "produces": [
          "application/octet-stream"
        ],
        "summary": "base64 decode the supplied string",
        "parameters": [
          {
            "description": "base64 encoded string",
            "name": "base64String",
            "in": "body",
            "required": true,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "string",
              "format": "binary"
            }
          },
          "400": {
            "description": "Bad request"
          }
        }
      }
    }
  }
}`))
}
