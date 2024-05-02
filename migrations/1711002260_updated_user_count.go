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

		collection, err := dao.FindCollectionByNameOrId("omsg2dsu8extwlx")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("czph7pzs")

		// add
		new_count := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "rluhjmam",
			"name": "count",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSize": 1
			}
		}`), new_count)
		collection.Schema.AddField(new_count)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("omsg2dsu8extwlx")
		if err != nil {
			return err
		}

		// add
		del_count := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "czph7pzs",
			"name": "count",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSize": 2000000
			}
		}`), del_count)
		collection.Schema.AddField(del_count)

		// remove
		collection.Schema.RemoveField("rluhjmam")

		return dao.SaveCollection(collection)
	})
}
