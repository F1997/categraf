package dm

import (
	"database/sql"
	"log"

	"github.com/F1997/categraf/pkg/tagx"
	"github.com/F1997/categraf/types"
)

// 采集数据库连接数
func (ins *Instance) gatherSchemaConnectionNums(slist *types.SampleList, db *sql.DB, globalTags map[string]string) {
	// if !ins.GatherSchemaSize {
	// 	return
	// }

	rows, err := db.Query(SQL_CONNECTION_NUMS)
	if err != nil {
		log.Println("E! failed to get schema size of", ins.Address, err)
		return
	}

	defer rows.Close()

	labels := tagx.Copy(globalTags)

	for rows.Next() {
		var connection_nums string

		err = rows.Scan(&connection_nums)
		if err != nil {
			log.Println("E! failed to scan rows of", ins.Address, err)
			return
		}

		// slist.PushFront(types.NewSample(inputName, "info_schema_basic", para_name, labels, map[string]string{"para_value": para_value}))
		slist.PushFront(types.NewSample(inputName, "connection_nums", connection_nums, labels))

	}
}
