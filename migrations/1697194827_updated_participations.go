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

		collection, err := dao.FindCollectionByNameOrId("4p0qj6w1lln3vcw")
		if err != nil {
			return err
		}

		// update
		edit_note := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "y8ueplkj",
			"name": "note",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_note)
		collection.Schema.AddField(edit_note)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("4p0qj6w1lln3vcw")
		if err != nil {
			return err
		}

		// update
		edit_note := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "y8ueplkj",
			"name": "note",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_note)
		collection.Schema.AddField(edit_note)

		return dao.SaveCollection(collection)
	})
}
