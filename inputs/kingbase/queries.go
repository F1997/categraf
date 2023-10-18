package kingbase

// (C) Datadog, Inc. 2020-present
// All rights reserved
// Licensed under Simplified BSD License (see LICENSE)

const (
	// 返回包含 "userstat" 和 "userstat_running" 这两个全局变量的当前值的结果集
	SQL_GLOBAL_VARIABLES = `SHOW GLOBAL VARIABLES WHERE Variable_Name='userstat' OR Variable_Name='userstat_running'`

	// 查询全局状态
	SQL_GLOBAL_STATUS = "SHOW GLOBAL STATUS;"

	// 查看实例启动时间
	SQL_START_TIME = "select sys_postmaster_start_time();"

	// 查看版本
	// SQL_VERSION = "select SVR_VERSION from v$instance"

	// 查看KES无故障运行时长
	SQL_UP_TIME  = "select date_trunc('second',current_timestamp - sys_postmaster_start_time()) as uptime;"
	SQL_RUN_TIME = "select extract(epoch from (current_timestamp - sys_postmaster_start_time())) as runningtime"
	// 统计所有数据库占用的磁盘空间总量
	SQL_SYS_DATABASE_SIZE = "select (sum(sys_database_size(datname))/1024/1024) || 'MB'  MB from sys_database;"
	// 查看所有会话执行的SQL信息
	SQL_ALL_QUERY_INFO = "select datname,usename,client_addr,client_port from sys_stat_activity;"
)
