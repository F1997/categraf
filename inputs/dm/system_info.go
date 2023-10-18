package dm

import (
	"database/sql"
	"log"

	"github.com/F1997/categraf/pkg/tagx"
	"github.com/F1997/categraf/types"
)

// 采集服务器信息
func (ins *Instance) gatherSystemInfo(slist *types.SampleList, db *sql.DB, globalTags map[string]string) {
	// if !ins.GatherSchemaSize {
	// 	return
	// }

	rows, err := db.Query(SQL_SYSTEM_INFO)
	if err != nil {
		log.Println("E! failed to get schema size of", ins.Address, err)
		return
	}

	defer rows.Close()

	labels := tagx.Copy(globalTags)

	for rows.Next() {
		var n_cpu int16

		var total_phy_size int64
		var free_phy_size int64
		var total_vir_size int64
		var free_vir_size int64
		var total_disk_size int64
		var free_disk_size int64

		var driver_name sql.NullString
		var driver_total_size sql.NullInt64
		var driver_free_size sql.NullInt64
		var load_one_average int64
		var load_five_average int64
		var load_fifteen_average int64

		var cpu_user_rate int64
		var cpu_system_rate int64
		var cpu_idle_rate int64
		var send_bytes_total int64
		var receive_bytes_total int64
		var send_bytes_per_second int64
		var receive_bytes_per_second int64
		var send_packages_per_second int64
		var receive_packages_per_second int64
		var used_phy_size int64

		err = rows.Scan(
			&n_cpu,
			&total_phy_size,
			&free_phy_size,
			&total_vir_size,
			&free_vir_size,
			&total_disk_size,
			&free_disk_size,
			&driver_name,
			&driver_total_size,
			&driver_free_size,
			&load_one_average,
			&load_five_average,
			&load_fifteen_average,
			&cpu_user_rate,
			&cpu_system_rate,
			&cpu_idle_rate,
			&send_bytes_total,
			&receive_bytes_total,
			&send_bytes_per_second,
			&receive_bytes_per_second,
			&send_packages_per_second,
			&receive_packages_per_second,
			&used_phy_size,
		)
		if err != nil {
			log.Println("E! failed to scan rows of", ins.Address, err)
			return
		}

		slist.PushFront(types.NewSample(inputName, "n_cpu", n_cpu, labels))
		slist.PushFront(types.NewSample(inputName, "total_phy_size", total_phy_size, labels))
		slist.PushFront(types.NewSample(inputName, "free_phy_size", free_phy_size, labels))
		slist.PushFront(types.NewSample(inputName, "total_vir_size", total_vir_size, labels))
		slist.PushFront(types.NewSample(inputName, "free_vir_size", free_vir_size, labels))
		slist.PushFront(types.NewSample(inputName, "total_disk_size", total_disk_size, labels))
		slist.PushFront(types.NewSample(inputName, "free_disk_size", free_disk_size, labels))
		slist.PushFront(types.NewSample(inputName, "driver_name", driver_name, labels))
		slist.PushFront(types.NewSample(inputName, "driver_total_size", driver_total_size, labels))
		slist.PushFront(types.NewSample(inputName, "driver_free_size", driver_free_size, labels))
		slist.PushFront(types.NewSample(inputName, "load_one_average", load_one_average, labels))
		slist.PushFront(types.NewSample(inputName, "load_five_average", load_five_average, labels))
		slist.PushFront(types.NewSample(inputName, "load_fifteen_average", load_fifteen_average, labels))
		slist.PushFront(types.NewSample(inputName, "cpu_user_rate", cpu_user_rate, labels))
		slist.PushFront(types.NewSample(inputName, "cpu_system_rate", cpu_system_rate, labels))
		slist.PushFront(types.NewSample(inputName, "cpu_idle_rate", cpu_idle_rate, labels))
		slist.PushFront(types.NewSample(inputName, "send_bytes_total", send_bytes_total, labels))
		slist.PushFront(types.NewSample(inputName, "receive_bytes_total", receive_bytes_total, labels))
		slist.PushFront(types.NewSample(inputName, "send_bytes_per_second", send_bytes_per_second, labels))
		slist.PushFront(types.NewSample(inputName, "receive_bytes_per_second", receive_bytes_per_second, labels))
		slist.PushFront(types.NewSample(inputName, "send_packages_per_second", send_packages_per_second, labels))
		slist.PushFront(types.NewSample(inputName, "receive_packages_per_second", receive_packages_per_second, labels))
		slist.PushFront(types.NewSample(inputName, "used_phy_size", used_phy_size, labels))
	}
}
