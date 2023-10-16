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

		collection, err := dao.FindCollectionByNameOrId("gltspnud37g5bx8")
		if err != nil {
			return err
		}

		options := map[string]any{}
		json.Unmarshal([]byte(`{
			"query": "SELECT length as length, owner as id from streaks ORDER BY lastweek LIMIT 1;"
		}`), &options)
		collection.SetOptions(options)

		// remove
		collection.Schema.RemoveField("r9jzc7ip")

		// add
		new_length := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "msxelkqg",
			"name": "length",
			"type": "number",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": false
			}
		}`), new_length)
		collection.Schema.AddField(new_length)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("gltspnud37g5bx8")
		if err != nil {
			return err
		}

		options := map[string]any{}
		json.Unmarshal([]byte(`{
			"query": "SELECT length as length, owner as id from streaks ORDER BY lastweek;"
		}`), &options)
		collection.SetOptions(options)

		// add
		del_length := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "r9jzc7ip",
			"name": "length",
			"type": "number",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": false
			}
		}`), del_length)
		collection.Schema.AddField(del_length)

		// remove
		collection.Schema.RemoveField("msxelkqg")

		return dao.SaveCollection(collection)
	})
}
