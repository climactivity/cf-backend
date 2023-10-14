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
			"id": "8l6d3yrwfsfj68t",
			"created": "2023-10-13 11:07:31.636Z",
			"updated": "2023-10-13 11:07:31.636Z",
			"name": "usercount",
			"type": "view",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "yaeessj1",
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
				"query": "SELECT COUNT(*) as count, (ROW_NUMBER() OVER()) as id FROM users;\n"
			}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("8l6d3yrwfsfj68t")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
