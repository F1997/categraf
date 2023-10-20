package dm

import (
	"database/sql"
	"log"

	"github.com/F1997/categraf/pkg/tagx"
	"github.com/F1997/categraf/types"
)

// 采集内存池使用情况
func (ins *Instance) gatherMemoryUsage(slist *types.SampleList, db *sql.DB, globalTags map[string]string) {
	// if !ins.GatherSchemaSize {
	// 	return
	// }

	rows, err := db.Query(SQL_MEMORY_USAGE)
	if err != nil {
		log.Println("E! failed to get schema size of", ins.Address, err)
		return
	}

	defer rows.Close()

	labels := tagx.Copy(globalTags)

	for rows.Next() {
		var name string
		var total_mb sql.NullInt64
		var mem_pool sql.NullInt64

		err = rows.Scan(&name, &total_mb, &mem_pool)
		if err != nil {
			log.Println("E! failed to scan rows of", ins.Address, err)
			return
		}

		// slist.PushFront(types.NewSample(inputName, "info_schema_basic", para_name, labels, map[string]string{"para_value": para_value}))
		slist.PushFront(types.NewSample(inputName, "mem_total_mb", total_mb, labels, map[string]string{"name": name}))
		slist.PushFront(types.NewSample(inputName, "mem_pool", mem_pool, labels, map[string]string{"name": name}))

	}
}
