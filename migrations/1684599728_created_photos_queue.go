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
			"id": "fyrlbiwqhdbdhtf",
			"created": "2023-05-20 16:22:08.343Z",
			"updated": "2023-05-20 16:22:08.343Z",
			"name": "photos_queue",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "6myqwebd",
					"name": "photo",
					"type": "relation",
					"required": false,
					"unique": false,
					"options": {
						"collectionId": "KbXls5Rim5vsjZr",
						"cascadeDelete": true,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": []
					}
				},
				{
					"system": false,
					"id": "srm8xhrp",
					"name": "is_pending",
					"type": "bool",
					"required": false,
					"unique": false,
					"options": {}
				}
			],
			"indexes": [],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
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
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("fyrlbiwqhdbdhtf")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
