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

		// add
		new_completion := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "vejqttum",
			"name": "completion",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "4p0qj6w1lln3vcw",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), new_completion)
		collection.Schema.AddField(new_completion)

		// update
		edit_field := &schema.SchemaField{}
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
		}`), edit_field)
		collection.Schema.AddField(edit_field)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("lu6j1fm6h29sgl9")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("vejqttum")

		// update
		edit_field := &schema.SchemaField{}
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
					"missed"
				]
			}
		}`), edit_field)
		collection.Schema.AddField(edit_field)

		return dao.SaveCollection(collection)
	})
}
