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
			"query": "SELECT\n    email, owner, \n    ifnull(count,0) as count,\n    users.id as id \nfrom users\nLEFT JOIN \n    (SELECT count(*) as count, owner from participations group by owner) \n    on owner = users.id\n;"
		}`), &options)
		collection.SetOptions(options)

		// remove
		collection.Schema.RemoveField("wmniga9i")

		// add
		new_email := &schema.SchemaField{}
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
		}`), new_email)
		collection.Schema.AddField(new_email)

		// add
		new_owner := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "zpjy4de3",
			"name": "owner",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), new_owner)
		collection.Schema.AddField(new_owner)

		// add
		new_count := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "0ev1arqw",
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
			"query": "SELECT (COUNT(*)) as participations, users.id as id from users LEFT JOIN participations ON users.id == participations.owner GROUP BY owner;"
		}`), &options)
		collection.SetOptions(options)

		// add
		del_participations := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "wmniga9i",
			"name": "participations",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), del_participations)
		collection.Schema.AddField(del_participations)

		// remove
		collection.Schema.RemoveField("iwinmvvm")

		// remove
		collection.Schema.RemoveField("zpjy4de3")

		// remove
		collection.Schema.RemoveField("0ev1arqw")

		return dao.SaveCollection(collection)
	})
}
