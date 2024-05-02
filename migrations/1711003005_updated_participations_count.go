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

		collection, err := dao.FindCollectionByNameOrId("hurhq6knh0xfmlt")
		if err != nil {
			return err
		}

		options := map[string]any{}
		json.Unmarshal([]byte(`{
			"query": "SELECT \n    ifnull(count,0) as count,\n    users.id as id \nfrom users\nLEFT JOIN \n    (SELECT count(*) as count, owner from participations group by owner) \n    on owner = users.id\n;"
		}`), &options)
		collection.SetOptions(options)

		// remove
		collection.Schema.RemoveField("fv4papjx")

		// add
		new_count := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "tk4xfntm",
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

		collection, err := dao.FindCollectionByNameOrId("hurhq6knh0xfmlt")
		if err != nil {
			return err
		}

		options := map[string]any{}
		json.Unmarshal([]byte(`{
			"query": "SELECT \n    ifnull(count,0) as count,\n    users.id as id \nfrom users\nLEFT JOIN \n    (SELECT count(*) as count, owner from participation_completions group by owner) \n    on owner = users.id\n;"
		}`), &options)
		collection.SetOptions(options)

		// add
		del_count := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "fv4papjx",
			"name": "count",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSize": 1
			}
		}`), del_count)
		collection.Schema.AddField(del_count)

		// remove
		collection.Schema.RemoveField("tk4xfntm")

		return dao.SaveCollection(collection)
	})
}
