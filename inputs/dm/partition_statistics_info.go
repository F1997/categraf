package dm

// import (
// 	"database/sql"
// 	"log"

// 	"github.com/F1997/categraf/types"
// )

// // 采集分区上的统计信息
// func (ins *Instance) gatherPartitionStatisticsInfo(slist *types.SampleList, db *sql.DB, globalTags map[string]string) {
// 	// if !ins.GatherSchemaSize {
// 	// 	return
// 	// }

// 	rows, err := db.Query(SQL_PARTITION_STATISTICS_INFO)
// 	if err != nil {
// 		log.Println("E! failed to get schema size of", ins.Address, err)
// 		return
// 	}

// 	defer rows.Close()

// 	// labels := tagx.Copy(globalTags)

// 	for rows.Next() {
// 		var status string

// 		err = rows.Scan(&status)
// 		if err != nil {
// 			log.Println("E! failed to scan rows of", ins.Address, err)
// 			return
// 		}

// 		// slist.PushFront(types.NewSample(inputName, "info_schema_basic", para_name, labels, map[string]string{"para_value": para_value}))
// 		// slist.PushFront(types.NewSample(inputName, "schema_status", status, labels))

// 	}
// }
