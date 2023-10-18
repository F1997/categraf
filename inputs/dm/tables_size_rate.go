package dm

import (
	"database/sql"
	"log"

	"github.com/F1997/categraf/pkg/tagx"
	"github.com/F1997/categraf/types"
)

// 采集表空间率
func (ins *Instance) gatherTableSizeRate(slist *types.SampleList, db *sql.DB, globalTags map[string]string) {
	// if !ins.GatherSchemaSize {
	// 	return
	// }

	rows, err := db.Query(SQL_TABLES_SIZE_RATE)
	if err != nil {
		log.Println("E! failed to get schema size of", ins.Address, err)
		return
	}

	defer rows.Close()

	labels := tagx.Copy(globalTags)

	for rows.Next() {
		var tablespace_name string
		var used_gb float64
		var free_gb float64
		var total_gb float64
		var per_free string

		err = rows.Scan(&tablespace_name, &used_gb, &free_gb, &total_gb, &per_free)
		if err != nil {
			log.Println("E! failed to scan rows of", ins.Address, err)
			return
		}

		// slist.PushFront(types.NewSample(inputName, "info_schema_basic", para_name, labels, map[string]string{"para_value": para_value}))
		slist.PushFront(types.NewSample("dm_table_size", "used_gb", used_gb, labels, map[string]string{"tablespace_name": tablespace_name}))
		slist.PushFront(types.NewSample("dm_table_size", "free_gb", free_gb, labels, map[string]string{"tablespace_name": tablespace_name}))
		slist.PushFront(types.NewSample("dm_table_size", "total_gb", total_gb, labels, map[string]string{"tablespace_name": tablespace_name}))
		slist.PushFront(types.NewSample("dm_table_size", "pre_free", per_free, labels, map[string]string{"tablespace_name": tablespace_name}))

	}
}

// 采集表空间使用率
func (ins *Instance) gatherTableSizeUseRate(slist *types.SampleList, db *sql.DB, globalTags map[string]string) {
	// if !ins.GatherSchemaSize {
	// 	return
	// }

	rows, err := db.Query(SQL_TABLES_SIZE_USE_RATE)
	if err != nil {
		log.Println("E! failed to get schema size of", ins.Address, err)
		return
	}

	defer rows.Close()

	labels := tagx.Copy(globalTags)

	for rows.Next() {
		var tablespace_name string
		var total_mb float64
		var free_mb float64
		var usage string

		err = rows.Scan(&tablespace_name, &total_mb, &free_mb, &usage)
		if err != nil {
			log.Println("E! failed to scan rows of", ins.Address, err)
			return
		}

		// slist.PushFront(types.NewSample(inputName, "info_schema_basic", para_name, labels, map[string]string{"para_value": para_value}))
		slist.PushFront(types.NewSample("dm_table_size_use", "total_mb", total_mb, labels, map[string]string{"tablespace_name": tablespace_name}))
		slist.PushFront(types.NewSample("dm_table_size_use", "free_mb", free_mb, labels, map[string]string{"tablespace_name": tablespace_name}))
		slist.PushFront(types.NewSample("dm_table_size", "usage", usage, labels, map[string]string{"tablespace_name": tablespace_name}))

	}
}
