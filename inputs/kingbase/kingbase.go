package kingbase

import (
	"database/sql"
	"fmt"
	"log"

	// "strings"
	"time"

	"github.com/F1997/categraf/config"
	"github.com/F1997/categraf/inputs"
	"github.com/F1997/categraf/pkg/tls"
	"github.com/F1997/categraf/types"

	// "github.com/go-sql-driver/mysql"
	_ "github.com/F1997/gokb"
)

const inputName = "kingbase"

// 定义了要查询的数据指标的配置
type QueryConfig struct {
	Mesurement    string          `toml:"mesurement"`
	LabelFields   []string        `toml:"label_fields"`
	MetricFields  []string        `toml:"metric_fields"`
	FieldToAppend string          `toml:"field_to_append"`
	Timeout       config.Duration `toml:"timeout"`
	Request       string          `toml:"request"`
}

// 插件实例的配置
type Instance struct {
	config.InstanceConfig

	Address        string `toml:"address"`
	Username       string `toml:"username"`
	Password       string `toml:"password"`
	DBname         string `toml:"dbname"`
	Sslmode        string `toml:"sslmode"`
	Parameters     string `toml:"parameters"`
	TimeoutSeconds int64  `toml:"timeout_seconds"`

	Queries       []QueryConfig `toml:"queries"`
	GlobalQueries []QueryConfig `toml:"-"`

	// ExtraStatusMetrics              bool `toml:"extra_status_metrics"`
	// ExtraInnodbMetrics              bool `toml:"extra_innodb_metrics"`
	// GatherProcessListProcessByState bool `toml:"gather_processlist_processes_by_state"`
	// GatherProcessListProcessByUser  bool `toml:"gather_processlist_processes_by_user"`
	// GatherSchemaSize                bool `toml:"gather_schema_size"`
	// GatherTableSize                 bool `toml:"gather_table_size"`
	// GatherSystemTableSize           bool `toml:"gather_system_table_size"`
	// GatherSlaveStatus               bool `toml:"gather_slave_status"`

	// DisableGlobalStatus      bool `toml:"disable_global_status"`
	// DisableGlobalVariables   bool `toml:"disable_global_variables"`
	// DisableInnodbStatus      bool `toml:"disable_innodb_status"`
	// DisableExtraInnodbStatus bool `toml:"disable_extra_innodb_status"`
	// DisablebinLogs           bool `toml:"disable_binlogs"`

	validMetrics map[string]struct{}
	dsn          string
	tls.ClientConfig
}

func (ins *Instance) Init() error {
	if ins.Address == "" {
		return types.ErrInstancesEmpty
	}

	// if ins.UseTLS {
	// 	tlsConfig, err := ins.ClientConfig.TLSConfig()
	// 	if err != nil {
	// 		return fmt.Errorf("failed to register tls config: %v", err)
	// 	}

	// 	err = mysql.RegisterTLSConfig("custom", tlsConfig)
	// 	if err != nil {
	// 		return fmt.Errorf("failed to register tls config: %v", err)
	// 	}
	// }
	// net := "tcp"
	// if strings.HasSuffix(ins.Address, ".sock") {
	// 	net = "unix"
	// }
	// ins.dsn = fmt.Sprintf("%s:%s@%s(%s)/?%s", ins.Username, ins.Password, net, ins.Address, ins.Parameters)
	// conf, err := mysql.ParseDSN(ins.dsn)
	// if err != nil {
	// 	return err
	// }
	// if conf.Timeout == 0 {
	// 	if ins.TimeoutSeconds == 0 {
	// 		ins.TimeoutSeconds = 3
	// 	}
	// 	conf.Timeout = time.Second * time.Duration(ins.TimeoutSeconds)
	// }

	// ins.dsn = conf.FormatDSN()

	ins.dsn = fmt.Sprintf("host=? user=? password=? dbname=? sslmode=?", ins.Address, ins.Username, ins.Password, ins.DBname, ins.Sslmode)

	ins.InitValidMetrics()

	return nil
}

// 初始化采集指标
func (ins *Instance) InitValidMetrics() {
	ins.validMetrics = make(map[string]struct{})

	// for key := range STATUS_VARS {
	// 	ins.validMetrics[key] = struct{}{}
	// }

	// for key := range VARIABLES_VARS {
	// 	ins.validMetrics[key] = struct{}{}
	// }

	// for key := range INNODB_VARS {
	// 	ins.validMetrics[key] = struct{}{}
	// }

	// for key := range BINLOG_VARS {
	// 	ins.validMetrics[key] = struct{}{}
	// }

	// for key := range GALERA_VARS {
	// 	ins.validMetrics[key] = struct{}{}
	// }

	// for key := range PERFORMANCE_VARS {
	// 	ins.validMetrics[key] = struct{}{}
	// }

	// for key := range SCHEMA_VARS {
	// 	ins.validMetrics[key] = struct{}{}
	// }

	// for key := range TABLE_VARS {
	// 	ins.validMetrics[key] = struct{}{}
	// }

	// for key := range REPLICA_VARS {
	// 	ins.validMetrics[key] = struct{}{}
	// }

	// for key := range GROUP_REPLICATION_VARS {
	// 	ins.validMetrics[key] = struct{}{}
	// }

	// for key := range SYNTHETIC_VARS {
	// 	ins.validMetrics[key] = struct{}{}
	// }

	// if ins.ExtraStatusMetrics {
	// 	for key := range OPTIONAL_STATUS_VARS {
	// 		ins.validMetrics[key] = struct{}{}
	// 	}
	// }

	// if ins.ExtraInnodbMetrics {
	// 	for key := range OPTIONAL_INNODB_VARS {
	// 		ins.validMetrics[key] = struct{}{}
	// 	}
	// }
}

// Catagraf 插件的配置，多个 Instance 实例，每个实例代表了一个 KingBase 数据库的连接
type KingBase struct {
	config.PluginConfig
	Instances []*Instance   `toml:"instances"`
	Queries   []QueryConfig `toml:"queries"`
}

// 初始化操作，注册 KingBase 插件
func init() {
	inputs.Add(inputName, func() inputs.Input {
		return &KingBase{}
	})
}

// 克隆 KingBase 插件的实例
func (k *KingBase) Clone() inputs.Input {
	return &KingBase{}
}

// 返回插件的名称
func (k *KingBase) Name() string {
	return inputName
}

// 返回所有 KingBase 插件实例的配置信息
func (k *KingBase) GetInstances() []inputs.Instance {
	ret := make([]inputs.Instance, len(k.Instances))
	for i := 0; i < len(k.Instances); i++ {
		k.Instances[i].GlobalQueries = k.Queries
		ret[i] = k.Instances[i]
	}
	return ret
}

// 收集 KingBase 数据库的各种指标，并将它们推送到 Catagraf 的 SampleList 中
func (ins *Instance) Gather(slist *types.SampleList) {
	tags := map[string]string{"address": ins.Address}

	begun := time.Now()

	// scrape use seconds
	defer func(begun time.Time) {
		use := time.Since(begun).Seconds()
		slist.PushSample(inputName, "scrape_use_seconds", use, tags)
	}(begun)

	db, err := sql.Open("kingbase", ins.dsn)
	if err != nil {
		slist.PushSample(inputName, "up", 0, tags)
		log.Println("E! failed to open kingbase:", err)
		return
	}

	defer db.Close()

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(time.Minute)

	if err = db.Ping(); err != nil {
		slist.PushSample(inputName, "up", 0, tags)
		log.Println("E! failed to ping kingbase:", err)
		return
	}

	slist.PushSample(inputName, "up", 1, tags)

	// cache := make(map[string]float64)

	ins.gatherInfoSchemaRuntime(slist, db, tags)

	// ins.gatherGlobalStatus(slist, db, tags, cache)
	// ins.gatherGlobalVariables(slist, db, tags, cache)
	// ins.gatherEngineInnodbStatus(slist, db, tags, cache)
	// ins.gatherEngineInnodbStatusCompute(slist, db, tags, cache)
	// ins.gatherBinlog(slist, db, tags)
	// ins.gatherProcesslistByState(slist, db, tags)
	// ins.gatherProcesslistByUser(slist, db, tags)
	// ins.gatherSchemaSize(slist, db, tags)
	// ins.gatherTableSize(slist, db, tags, false)
	// ins.gatherTableSize(slist, db, tags, true)
	// ins.gatherSlaveStatus(slist, db, tags)
	// ins.gatherCustomQueries(slist, db, tags)
}
