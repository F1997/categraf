package dm

// import (
// 	"database/sql"
// 	"log"

// 	"github.com/F1997/categraf/pkg/tagx"
// 	"github.com/F1997/categraf/types"
// )

// // 采集用户状态及口令修改期限
// func (ins *Instance) gatherUserPasswdModifyPeriod(slist *types.SampleList, db *sql.DB, globalTags map[string]string) {
// 	// if !ins.GatherSchemaSize {
// 	// 	return
// 	// }

// 	rows, err := db.Query(SQL_USER_PASSWORD_MODIFY_PERIOD)
// 	if err != nil {
// 		log.Println("E! failed to get schema size of", ins.Address, err)
// 		return
// 	}

// 	defer rows.Close()

// 	labels := tagx.Copy(globalTags)

// 	for rows.Next() {
// 		var username string
// 		var life_time int64

// 		err = rows.Scan(&username, &life_time)
// 		if err != nil {
// 			log.Println("E! failed to scan rows of", ins.Address, err)
// 			return
// 		}

// 		// slist.PushFront(types.NewSample(inputName, "info_schema_basic", para_name, labels, map[string]string{"para_value": para_value}))
// 		slist.PushFront(types.NewSample(inputName, "life_time", life_time, labels, map[string]string{"username": username}))

// 	}
// }
