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

		collection, err := dao.FindCollectionByNameOrId("omsg2dsu8extwlx")
		if err != nil {
			return err
		}

		options := map[string]any{}
		json.Unmarshal([]byte(`{
			"query": "select count, id from\n  (\n  SELECT COUNT(*) as count, region as id FROM users GROUP BY region\n  UNION ALL\n  SELECT COUNT(*) as count, \"usercount\" as id FROM users\n  )\n;"
		}`), &options)
		collection.SetOptions(options)

		// remove
		collection.Schema.RemoveField("8floosgj")

		// add
		new_count := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "gwqkioh0",
			"name": "count",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), new_count)
		collection.Schema.AddField(new_count)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("omsg2dsu8extwlx")
		if err != nil {
			return err
		}

		options := map[string]any{}
		json.Unmarshal([]byte(`{
			"query": "SELECT COUNT(*)as count, \"count\" AS id FROM users;"
		}`), &options)
		collection.SetOptions(options)

		// add
		del_count := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "8floosgj",
			"name": "count",
			"type": "number",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": false
			}
		}`), del_count)
		collection.Schema.AddField(del_count)

		// remove
		collection.Schema.RemoveField("gwqkioh0")

		return dao.SaveCollection(collection)
	})
}
