package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("omsg2dsu8extwlx")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("")

		collection.ViewRule = types.Pointer("")

		// remove
		collection.Schema.RemoveField("gwqkioh0")

		// add
		new_count := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "vkberteo",
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

		collection.ListRule = nil

		collection.ViewRule = nil

		// add
		del_count := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "gwqkioh0",
			"name": "count",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), del_count)
		collection.Schema.AddField(del_count)

		// remove
		collection.Schema.RemoveField("vkberteo")

		return dao.SaveCollection(collection)
	})
}
