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
		collection.Schema.RemoveField("iwinmvvm")

		// remove
		collection.Schema.RemoveField("zpjy4de3")

		// remove
		collection.Schema.RemoveField("0ev1arqw")

		// add
		new_count := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "prvxcbkr",
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

		collection, err := dao.FindCollectionByNameOrId("hurhq6knh0xfmlt")
		if err != nil {
			return err
		}

		options := map[string]any{}
		json.Unmarshal([]byte(`{
			"query": "SELECT\n    email, owner, \n    ifnull(count,0) as count,\n    users.id as id \nfrom users\nLEFT JOIN \n    (SELECT count(*) as count, owner from participations group by owner) \n    on owner = users.id\n;"
		}`), &options)
		collection.SetOptions(options)

		// add
		del_email := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "iwinmvvm",
			"name": "email",
			"type": "email",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"exceptDomains": null,
				"onlyDomains": null
			}
		}`), del_email)
		collection.Schema.AddField(del_email)

		// add
		del_owner := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "zpjy4de3",
			"name": "owner",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), del_owner)
		collection.Schema.AddField(del_owner)

		// add
		del_count := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "0ev1arqw",
			"name": "count",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), del_count)
		collection.Schema.AddField(del_count)

		// remove
		collection.Schema.RemoveField("prvxcbkr")

		return dao.SaveCollection(collection)
	})
}
