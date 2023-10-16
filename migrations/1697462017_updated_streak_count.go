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
			"query": "SELECT\n     users.id as id,  ifnull(len, 0) as length\nfrom users\nLEFT JOIN \n    (SELECT length as len, owner from streaks ORDER BY lastweek DESC LIMIT 1) \n    on owner = users.id\n;"
		}`), &options)
		collection.SetOptions(options)

		// remove
		collection.Schema.RemoveField("0devqua6")

		// add
		new_length := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "7iotwjgx",
			"name": "length",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
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
			"query": "SELECT length as length, owner as id from streaks ORDER BY lastweek DESC LIMIT 1;"
		}`), &options)
		collection.SetOptions(options)

		// add
		del_length := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "0devqua6",
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
		collection.Schema.RemoveField("7iotwjgx")

		return dao.SaveCollection(collection)
	})
}
