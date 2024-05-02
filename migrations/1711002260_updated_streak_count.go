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

		collection, err := dao.FindCollectionByNameOrId("gltspnud37g5bx8")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("n1woeoyr")

		// add
		new_length := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "ufrdv0mr",
			"name": "length",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSize": 1
			}
		}`), new_length)
		collection.Schema.AddField(new_length)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("gltspnud37g5bx8")
		if err != nil {
			return err
		}

		// add
		del_length := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "n1woeoyr",
			"name": "length",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSize": 2000000
			}
		}`), del_length)
		collection.Schema.AddField(del_length)

		// remove
		collection.Schema.RemoveField("ufrdv0mr")

		return dao.SaveCollection(collection)
	})
}
