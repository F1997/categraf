package dm

import (
	"database/sql"
	"log"

	"github.com/F1997/categraf/pkg/tagx"
	"github.com/F1997/categraf/types"
)

func (ins *Instance) gatherInfoSchemaBasic(slist *types.SampleList, db *sql.DB, globalTags map[string]string) {
	// if !ins.GatherSchemaSize {
	// 	return
	// }

	rows, err := db.Query(SQL_INFO_SCHEMA_BASIC)
	if err != nil {
		log.Println("E! failed to get schema size of", ins.Address, err)
		return
	}

	defer rows.Close()

	labels := tagx.Copy(globalTags)

	for rows.Next() {
		var para_name string
		var para_value string

		err = rows.Scan(&para_name, &para_value)
		if err != nil {
			log.Println("E! failed to scan rows of", ins.Address, err)
			return
		}

		// slist.PushFront(types.NewSample(inputName, "info_schema_basic", para_name, labels, map[string]string{"para_value": para_value}))
		slist.PushFront(types.NewSample(inputName, para_name, para_value, labels))

	}
}
