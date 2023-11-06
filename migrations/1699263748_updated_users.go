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

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		// add
		new_notification_email := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "78cdriem",
			"name": "notification_email",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), new_notification_email)
		collection.Schema.AddField(new_notification_email)

		// add
		new_notification_push := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "tbqhhzt2",
			"name": "notification_push",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), new_notification_push)
		collection.Schema.AddField(new_notification_push)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("78cdriem")

		// remove
		collection.Schema.RemoveField("tbqhhzt2")

		return dao.SaveCollection(collection)
	})
}
