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
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("lu6j1fm6h29sgl9")
		if err != nil {
			return err
		}

		// add
		new_year := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "yppvhw2z",
			"name": "year",
			"type": "number",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": false
			}
		}`), new_year)
		collection.Schema.AddField(new_year)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("lu6j1fm6h29sgl9")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("yppvhw2z")

		return dao.SaveCollection(collection)
	})
}
