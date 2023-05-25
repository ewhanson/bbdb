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
		jsonData := `{
			"id": "de4vrswhzc1s8u9",
			"created": "2023-05-23 21:29:25.314Z",
			"updated": "2023-05-23 21:29:25.314Z",
			"name": "tags",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "kosrr7a7",
					"name": "name",
					"type": "text",
					"required": true,
					"unique": false,
					"options": {
						"min": 3,
						"max": null,
						"pattern": ""
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
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("de4vrswhzc1s8u9")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
