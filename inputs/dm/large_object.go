package dm

// import (
// 	"database/sql"
// 	"log"

// 	"github.com/F1997/categraf/pkg/tagx"
// 	"github.com/F1997/categraf/types"
// )

// // 采集大对象数目以及大对象ID
// func (ins *Instance) gatherLargeObject(slist *types.SampleList, db *sql.DB, globalTags map[string]string) {
// 	// if !ins.GatherSchemaSize {
// 	// 	return
// 	// }

// 	rows, err := db.Query(SQL_LARGE_OBJECT)
// 	if err != nil {
// 		log.Println("E! failed to get schema size of", ins.Address, err)
// 		return
// 	}

// 	defer rows.Close()

// 	labels := tagx.Copy(globalTags)

// 	large_object_num := 0
// 	large_blob_seg_num := 0
// 	large_clob_seg_num := 0

// 	for rows.Next() {
// 		var obj_type string
// 		var obj_id int64

// 		err = rows.Scan(&obj_type, &obj_id)
// 		if err != nil {
// 			log.Println("E! failed to scan rows of", ins.Address, err)
// 			return
// 		}
// 		large_object_num += 1
// 		switch obj_type {
// 		case "BLOB_SEG":
// 			slist.PushFront(types.NewSample(inputName, "large_obj_id", obj_id, labels, map[string]string{"type": obj_type}))
// 			large_blob_seg_num += 1
// 		case "CLOB_SEG":
// 			slist.PushFront(types.NewSample(inputName, "large_obj_id", obj_id, labels, map[string]string{"type": obj_type}))
// 			large_clob_seg_num += 1
// 		}
// 	}

// 	slist.PushFront(types.NewSample(inputName, "large_object_num", large_object_num, labels))
// 	slist.PushFront(types.NewSample(inputName, "large_blob_seg_num", large_object_num, labels))
// 	slist.PushFront(types.NewSample(inputName, "large_clob_seg_num", large_object_num, labels))

// }
