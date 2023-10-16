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
			"id": "gltspnud37g5bx8",
			"created": "2023-10-16 12:00:52.669Z",
			"updated": "2023-10-16 12:00:52.669Z",
			"name": "streak_count",
			"type": "view",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "r9jzc7ip",
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
				}
			],
			"indexes": [],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {
				"query": "SELECT length as length, owner as id from streaks ORDER BY lastweek;"
			}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("gltspnud37g5bx8")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
