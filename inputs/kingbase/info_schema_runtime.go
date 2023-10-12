package kingbase

import (
	"database/sql"
	"log"

	"github.com/F1997/categraf/pkg/tagx"
	"github.com/F1997/categraf/types"
)

func (ins *Instance) gatherInfoSchemaRuntime(slist *types.SampleList, db *sql.DB, globalTags map[string]string) {
	if !ins.GatherSchemaSize {
		return
	}

	rows, err := db.Query(SQL_RUN_TIME)
	if err != nil {
		log.Println("E! failed to get schema size of", ins.Address, err)
		return
	}

	defer rows.Close()

	labels := tagx.Copy(globalTags)

	for rows.Next() {
		var runtime string

		err = rows.Scan(&runtime)
		if err != nil {
			log.Println("E! failed to scan rows of", ins.Address, err)
			return
		}

		slist.PushFront(types.NewSample(inputName, "info_schema_runtime", runtime))
	}
}
