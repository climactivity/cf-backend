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

		collection, err := dao.FindCollectionByNameOrId("lu6j1fm6h29sgl9")
		if err != nil {
			return err
		}

		// update
		edit_state := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "yzwkb36w",
			"name": "state",
			"type": "select",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"dont_remind",
					"remind",
					"pending",
					"missed",
					"completed"
				]
			}
		}`), edit_state)
		collection.Schema.AddField(edit_state)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("lu6j1fm6h29sgl9")
		if err != nil {
			return err
		}

		// update
		edit_state := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "yzwkb36w",
			"name": "field",
			"type": "select",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"dont_remind",
					"remind",
					"pending",
					"missed",
					"completed"
				]
			}
		}`), edit_state)
		collection.Schema.AddField(edit_state)

		return dao.SaveCollection(collection)
	})
}
