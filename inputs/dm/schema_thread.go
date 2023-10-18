package dm

import (
	"database/sql"
	"log"

	"github.com/F1997/categraf/pkg/tagx"
	"github.com/F1997/categraf/types"
)

func (ins *Instance) gatherInfoSchemaThread(slist *types.SampleList, db *sql.DB, globalTags map[string]string) {
	// if !ins.GatherSchemaSize {
	// 	return
	// }

	rows, err := db.Query(SQL_INFO_SCHEMA_THREAD)
	if err != nil {
		log.Println("E! failed to get schema size of", ins.Address, err)
		return
	}

	defer rows.Close()

	labels := tagx.Copy(globalTags)

	for rows.Next() {
		var name string
		var num string

		err = rows.Scan(&name, &num)
		if err != nil {
			log.Println("E! failed to scan rows of", ins.Address, err)
			return
		}

		// slist.PushFront(types.NewSample(inputName, "info_schema_thread", name, labels, map[string]string{"num": num}))
		slist.PushFront(types.NewSample("", name, num, labels))

	}
}
