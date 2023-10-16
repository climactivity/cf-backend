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

		collection, err := dao.FindCollectionByNameOrId("hurhq6knh0xfmlt")
		if err != nil {
			return err
		}

		collection.ViewRule = types.Pointer("id = @request.auth.id")

		// remove
		collection.Schema.RemoveField("slis3azf")

		// add
		new_participations := &schema.SchemaField{}
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
		}`), new_participations)
		collection.Schema.AddField(new_participations)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("hurhq6knh0xfmlt")
		if err != nil {
			return err
		}

		collection.ViewRule = nil

		// add
		del_participations := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "slis3azf",
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
		collection.Schema.RemoveField("qsiw9zfr")

		return dao.SaveCollection(collection)
	})
}
