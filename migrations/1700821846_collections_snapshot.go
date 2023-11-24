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
		jsonData := `[
			{
				"id": "_pb_users_auth_",
				"created": "2023-10-12 10:00:56.334Z",
				"updated": "2023-11-06 09:42:28.704Z",
				"name": "users",
				"type": "auth",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "users_name",
						"name": "name",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "users_avatar",
						"name": "avatar",
						"type": "file",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [
								"image/jpeg",
								"image/png",
								"image/svg+xml",
								"image/gif",
								"image/webp"
							],
							"thumbs": null,
							"protected": false
						}
					},
					{
						"system": false,
						"id": "wmsntlgv",
						"name": "region",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "78cdriem",
						"name": "notification_email",
						"type": "bool",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "tbqhhzt2",
						"name": "notification_push",
						"type": "bool",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {}
					}
				],
				"indexes": [],
				"listRule": "id = @request.auth.id",
				"viewRule": "id = @request.auth.id",
				"createRule": "",
				"updateRule": "id = @request.auth.id",
				"deleteRule": "id = @request.auth.id",
				"options": {
					"allowEmailAuth": true,
					"allowOAuth2Auth": true,
					"allowUsernameAuth": true,
					"exceptEmailDomains": null,
					"manageRule": null,
					"minPasswordLength": 8,
					"onlyEmailDomains": null,
					"requireEmail": false
				}
			},
			{
				"id": "4p0qj6w1lln3vcw",
				"created": "2023-10-13 10:58:05.499Z",
				"updated": "2023-10-18 15:40:37.649Z",
				"name": "participations",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "apn1do78",
						"name": "owner",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "4kfsyjy2",
						"name": "option",
						"type": "select",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"button",
								"sticker",
								"banner",
								"flag",
								"window_foil"
							]
						}
					},
					{
						"system": false,
						"id": "n0llf7l5",
						"name": "date",
						"type": "date",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "xr0ooyvm",
						"name": "note",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "8xwrmm2k",
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
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_lvMokuh` + "`" + ` ON ` + "`" + `participations` + "`" + ` (` + "`" + `owner` + "`" + `)"
				],
				"listRule": "owner.id = @request.auth.id",
				"viewRule": "owner.id = @request.auth.id",
				"createRule": "owner.id = @request.auth.id",
				"updateRule": "owner.id = @request.auth.id",
				"deleteRule": "owner.id = @request.auth.id",
				"options": {}
			},
			{
				"id": "omsg2dsu8extwlx",
				"created": "2023-10-15 10:38:57.148Z",
				"updated": "2023-10-15 10:43:43.956Z",
				"name": "user_count",
				"type": "view",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "vkberteo",
						"name": "count",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {}
					}
				],
				"indexes": [],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {
					"query": "select count, id from\n  (\n  SELECT COUNT(*) as count, region as id FROM users GROUP BY region\n  UNION ALL\n  SELECT COUNT(*) as count, \"usercount\" as id FROM users\n  )\n;"
				}
			},
			{
				"id": "hurhq6knh0xfmlt",
				"created": "2023-10-16 11:24:35.173Z",
				"updated": "2023-10-16 13:04:27.636Z",
				"name": "participations_count",
				"type": "view",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "prvxcbkr",
						"name": "count",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {}
					}
				],
				"indexes": [],
				"listRule": null,
				"viewRule": "id = @request.auth.id",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {
					"query": "SELECT \n    ifnull(count,0) as count,\n    users.id as id \nfrom users\nLEFT JOIN \n    (SELECT count(*) as count, owner from participations group by owner) \n    on owner = users.id\n;"
				}
			},
			{
				"id": "abmhdkjt7rp4pof",
				"created": "2023-10-16 11:53:53.215Z",
				"updated": "2023-10-18 15:40:53.667Z",
				"name": "streaks",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "1hzcy25a",
						"name": "owner",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "m3rqxkt4",
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
					},
					{
						"system": false,
						"id": "kxhuzye4",
						"name": "lastweek",
						"type": "date",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "lwldhrpe",
						"name": "firstweek",
						"type": "date",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					}
				],
				"indexes": [],
				"listRule": "owner.id = @request.auth.id",
				"viewRule": "owner.id = @request.auth.id",
				"createRule": "owner.id = @request.auth.id",
				"updateRule": "owner.id = @request.auth.id",
				"deleteRule": "owner.id = @request.auth.id",
				"options": {}
			},
			{
				"id": "gltspnud37g5bx8",
				"created": "2023-10-16 12:00:52.669Z",
				"updated": "2023-10-16 13:13:37.944Z",
				"name": "streak_count",
				"type": "view",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "7iotwjgx",
						"name": "length",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {}
					}
				],
				"indexes": [],
				"listRule": "id = @request.auth.id",
				"viewRule": "id = @request.auth.id",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {
					"query": "SELECT\n     users.id as id,  ifnull(len, 0) as length\nfrom users\nLEFT JOIN \n    (SELECT length as len, owner from streaks ORDER BY lastweek DESC LIMIT 1) \n    on owner = users.id\n;"
				}
			},
			{
				"id": "rsnh1vpv4rzoca8",
				"created": "2023-11-17 10:36:39.859Z",
				"updated": "2023-11-17 10:36:39.859Z",
				"name": "notification_settings",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "spuxnps3",
						"name": "user",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "j0orsxuq",
						"name": "push_settings",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {}
					}
				],
				"indexes": [],
				"listRule": null,
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
