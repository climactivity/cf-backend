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
			"query": "SELECT COUNT(*) as participations, users.id as id from users LEFT JOIN participations ON users.id == participations.owner GROUP BY owner;"
		}`), &options)
		collection.SetOptions(options)

		// remove
		collection.Schema.RemoveField("qsiw9zfr")

		// add
		new_participations := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "qefjbpvd",
			"name": "participations",
			"type": "number",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": false
			}
		}`), new_participations)
		collection.Schema.AddField(new_participations)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("hurhq6knh0xfmlt")
		if err != nil {
			return err
		}

		options := map[string]any{}
		json.Unmarshal([]byte(`{
			"query": "SELECT COUNT(*) as participations, owner as id from participations GROUP BY owner"
		}`), &options)
		collection.SetOptions(options)

		// add
		del_participations := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "qsiw9zfr",
			"name": "participations",
			"type": "number",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": false
			}
		}`), del_participations)
		collection.Schema.AddField(del_participations)

		// remove
		collection.Schema.RemoveField("qefjbpvd")

		return dao.SaveCollection(collection)
	})
}
