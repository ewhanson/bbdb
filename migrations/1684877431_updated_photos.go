package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("KbXls5Rim5vsjZr")
		if err != nil {
			return err
		}

		// add
		new_tags := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "qc2y0sxe",
			"name": "tags",
			"type": "relation",
			"required": false,
			"unique": false,
			"options": {
				"collectionId": "de4vrswhzc1s8u9",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": null,
				"displayFields": [
					"name"
				]
			}
		}`), new_tags)
		collection.Schema.AddField(new_tags)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("KbXls5Rim5vsjZr")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("qc2y0sxe")

		return dao.SaveCollection(collection)
	})
}
