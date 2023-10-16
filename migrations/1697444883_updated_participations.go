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

		// remove
		collection.Schema.RemoveField("soyyqqg9")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("4p0qj6w1lln3vcw")
		if err != nil {
			return err
		}

		// add
		del_note := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "soyyqqg9",
			"name": "note",
			"type": "editor",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"convertUrls": false
			}
		}`), del_note)
		collection.Schema.AddField(del_note)

		return dao.SaveCollection(collection)
	})
}
