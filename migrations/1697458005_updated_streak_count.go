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

		collection, err := dao.FindCollectionByNameOrId("gltspnud37g5bx8")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("id = @request.auth.id")

		collection.ViewRule = types.Pointer("id = @request.auth.id")

		// remove
		collection.Schema.RemoveField("kg3teagj")

		// add
		new_length := &schema.SchemaField{}
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
		}`), new_length)
		collection.Schema.AddField(new_length)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("gltspnud37g5bx8")
		if err != nil {
			return err
		}

		collection.ListRule = nil

		collection.ViewRule = nil

		// add
		del_length := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "kg3teagj",
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
		collection.Schema.RemoveField("0devqua6")

		return dao.SaveCollection(collection)
	})
}
