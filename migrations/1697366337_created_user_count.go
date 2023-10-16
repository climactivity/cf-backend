package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "omsg2dsu8extwlx",
			"created": "2023-10-15 10:38:57.148Z",
			"updated": "2023-10-15 10:38:57.148Z",
			"name": "user_count",
			"type": "view",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "8floosgj",
					"name": "count",
					"type": "number",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"noDecimal": false
					}
				}
			],
			"indexes": [],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {
				"query": "SELECT COUNT(*)as count, \"count\" AS id FROM users;"
			}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("omsg2dsu8extwlx")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
