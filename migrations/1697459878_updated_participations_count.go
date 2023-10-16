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

		// add
		new_participations := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "jy5e2k67",
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
			"query": "SELECT  users.id as id from users LEFT JOIN participations ON users.id == participations.owner GROUP BY owner;"
		}`), &options)
		collection.SetOptions(options)

		// remove
		collection.Schema.RemoveField("jy5e2k67")

		return dao.SaveCollection(collection)
	})
}
