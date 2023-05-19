package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `[
			{
				"id": "KbXls5Rim5vsjZr",
				"created": "2022-07-29 02:42:24.606Z",
				"updated": "2023-03-16 04:36:13.349Z",
				"name": "photos",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "sirr1kxd",
						"name": "file",
						"type": "file",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [
								"image/jpg",
								"image/jpeg",
								"image/png",
								"image/svg+xml",
								"image/gif"
							],
							"thumbs": [
								"378x504",
								"504x378"
							],
							"protected": false
						}
					},
					{
						"system": false,
						"id": "xpylwb2l",
						"name": "description",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "mdp3httm",
						"name": "dateTaken",
						"type": "date",
						"required": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.role = 'viewer' || @request.auth.role = 'uploader'",
				"viewRule": "@request.auth.role = 'viewer' || @request.auth.role = 'uploader'",
				"createRule": "@request.auth.role = 'uploader'",
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "eo0m00o3tcctqwo",
				"created": "2022-09-08 19:49:36.011Z",
				"updated": "2023-02-23 02:41:32.791Z",
				"name": "subscribers",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "ur7amurz",
						"name": "email",
						"type": "email",
						"required": true,
						"unique": true,
						"options": {
							"exceptDomains": null,
							"onlyDomains": null
						}
					},
					{
						"system": false,
						"id": "91l5rdgm",
						"name": "name",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": 90,
							"pattern": ""
						}
					}
				],
				"indexes": [],
				"listRule": null,
				"viewRule": null,
				"createRule": "@request.auth.id!=\"\"",
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "De8dLzRmJYuZ4RE",
				"created": "2023-01-29 01:19:56.846Z",
				"updated": "2023-01-29 01:20:14.092Z",
				"name": "users",
				"type": "auth",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "yljyw2gh",
						"name": "name",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "gemminvv",
						"name": "avatar",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [
								"image/jpg",
								"image/jpeg",
								"image/png",
								"image/svg+xml",
								"image/gif"
							],
							"thumbs": null,
							"protected": false
						}
					},
					{
						"system": false,
						"id": "mtl7ahpx",
						"name": "role",
						"type": "select",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"viewer",
								"uploader"
							]
						}
					}
				],
				"indexes": [],
				"listRule": "id = @request.auth.id",
				"viewRule": "id = @request.auth.id",
				"createRule": "",
				"updateRule": "id = @request.auth.id",
				"deleteRule": null,
				"options": {
					"allowEmailAuth": true,
					"allowOAuth2Auth": true,
					"allowUsernameAuth": false,
					"exceptEmailDomains": null,
					"manageRule": null,
					"minPasswordLength": 8,
					"onlyEmailDomains": null,
					"requireEmail": true
				}
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
