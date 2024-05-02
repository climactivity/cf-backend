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
			"id": "lu6j1fm6h29sgl9",
			"created": "2024-03-21 06:36:28.207Z",
			"updated": "2024-03-21 06:36:28.207Z",
			"name": "weeks",
			"type": "base",
			"system": false,
			"schema": [
				{
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
				},
				{
					"system": false,
					"id": "zn3mhnew",
					"name": "weeknumber",
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
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("lu6j1fm6h29sgl9")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
