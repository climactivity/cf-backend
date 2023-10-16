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

		// add
		new_date := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "n0llf7l5",
			"name": "date",
			"type": "date",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), new_date)
		collection.Schema.AddField(new_date)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("4p0qj6w1lln3vcw")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("n0llf7l5")

		return dao.SaveCollection(collection)
	})
}
