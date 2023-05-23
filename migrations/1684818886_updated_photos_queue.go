package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("fyrlbiwqhdbdhtf")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("@request.auth.role = 'viewer' || @request.auth.role = 'uploader'")

		collection.ViewRule = types.Pointer("@request.auth.role = 'viewer' || @request.auth.role = 'uploader'")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("fyrlbiwqhdbdhtf")
		if err != nil {
			return err
		}

		collection.ListRule = nil

		collection.ViewRule = nil

		return dao.SaveCollection(collection)
	})
}
