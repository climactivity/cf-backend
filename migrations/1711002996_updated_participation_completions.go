package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("4p0qj6w1lln3vcw")
		if err != nil {
			return err
		}

		collection.Name = "participations"

		json.Unmarshal([]byte(`[
			"CREATE INDEX `+"`"+`idx_lvMokuh`+"`"+` ON `+"`"+`participations`+"`"+` (`+"`"+`owner`+"`"+`)"
		]`), &collection.Indexes)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("4p0qj6w1lln3vcw")
		if err != nil {
			return err
		}

		collection.Name = "participation_completions"

		json.Unmarshal([]byte(`[
			"CREATE INDEX `+"`"+`idx_lvMokuh`+"`"+` ON `+"`"+`participation_completions`+"`"+` (`+"`"+`owner`+"`"+`)"
		]`), &collection.Indexes)

		return dao.SaveCollection(collection)
	})
}
