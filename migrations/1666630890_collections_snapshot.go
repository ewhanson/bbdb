package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

// Auto generated migration with the most recent collections configuration.
func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `[
			{
				"id": "De8dLzRmJYuZ4RE",
				"created": "2022-07-29 02:39:09.804",
				"updated": "2022-07-29 02:43:01.370",
				"name": "profiles",
				"system": true,
				"schema": [
					{
						"system": true,
						"id": "mbsrmqsc",
						"name": "userId",
						"type": "user",
						"required": true,
						"unique": true,
						"options": {
							"maxSelect": 1,
							"cascadeDelete": true
						}
					},
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
							"thumbs": null
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
				"listRule": "userId = @request.user.id",
				"viewRule": "userId = @request.user.id",
				"createRule": "userId = @request.user.id",
				"updateRule": "userId = @request.user.id",
				"deleteRule": null
			},
			{
				"id": "KbXls5Rim5vsjZr",
				"created": "2022-07-29 02:42:24.606",
				"updated": "2022-09-12 18:18:23.443",
				"name": "photos",
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
							]
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
				"listRule": "@request.user.profile.role = 'viewer'",
				"viewRule": "@request.user.profile.role = 'viewer'",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null
			},
			{
				"id": "eo0m00o3tcctqwo",
				"created": "2022-09-08 19:49:36.011",
				"updated": "2022-09-08 19:49:36.011",
				"name": "subscribers",
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
					}
				],
				"listRule": null,
				"viewRule": null,
				"createRule": "@request.user.id!=\"\"",
				"updateRule": null,
				"deleteRule": null
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		// no revert since the configuration on the environment, on which
		// the migration was executed, could have changed via the UI/API
		return nil
	})
}
