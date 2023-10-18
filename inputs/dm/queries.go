package dm

// (C) Datadog, Inc. 2020-present
// All rights reserved
// Licensed under Simplified BSD License (see LICENSE)

const (
	//
	SQL_INFO_SCHEMA_BASIC = `
		select PARA_NAME, PARA_VALUE from SYS."V$DM_INI" where PARA_NAME = 'MAX_SESSIONS' or PARA_NAME = 'PORT_NUM';
	`
	//
	SQL_INFO_SCHEMA_THREAD = `
		SELECT DISTINCT NAME as NAME, COUNT(*)  AS NUM FROM V$THREADS GROUP BY NAME ORDER BY NUM DESC;
	`
	// // 查看达梦数据库版本
	// SQL = `
	// 	select SVR_VERSION from v$instance;
	// `

	// 查看达梦数据库运行状态
	SQL_DM_STATUS = `
		select status$ as status from v$instance;
	`
	// 查询服务器信息
	SQL_SYSTEM_INFO = `
		select * from  V$SYSTEMINFO;
	`
	// 查看数据库是否有阻塞
	SQL_SCHEMA_LOCK = `
		select o.name,l.* from v$lock l,sysobjects o where l.table_id=o.id and blocked=1;
	`
	// 等待事件
	SQL_WAIT = `
		with locks as(
			select o.name,l.*,s.sess_id,s.sql_text,s.clnt_ip,s.last_send_time  from v$lock l,sysobjects o,v$sessions s
			where l.table_id=o.id and l.trx_id=s.trx_id ),
		lock_tr as (select trx_id wt_trxid,row_idx blk_trxid from locks where blocked=1),
		res as(select sysdate stattime,t1.name,t1.sess_id wt_sessid,s.wt_trxid,
			t2.sess_id blk_sessid,s.blk_trxid,t2.clnt_ip,SF_GET_SESSION_SQL(t1.sess_id) fulsql,
			datediff(ss,t1.last_send_time,sysdate) ss,t1.sql_text wt_sql  from lock_tr s,locks t1,locks t2
			where t1.ltype='OBJECT'  and t1.table_id<>0   and t2.ltype='OBJECT'  and t2.table_id<>0
			and s.wt_trxid=t1.trx_id  and s.blk_trxid=t2.trx_id)
		select distinct wt_sql,clnt_ip,ss,wt_trxid,blk_trxid  from res;
	`

	// 表空间率
	SQL_TABLES_SIZE_RATE = `
		SELECT F.TABLESPACE_NAME,
			ROUND((T.TOTAL_SPACE - F.FREE_SPACE) / 1024,2) as  "USED(GB)",
			ROUND(F.FREE_SPACE / 1024,2) "FREE (GB)",
			ROUND(T.TOTAL_SPACE / 1024,2)  "TOTAL(GB)",
			ROUND((F.FREE_SPACE / T.TOTAL_SPACE) * 100) ||  '% '  as  "PER_FREE"
		FROM (SELECT TABLESPACE_NAME,
					ROUND(SUM(BLOCKS *
								(SELECT PARA_VALUE / 1024
									FROM V$DM_INI
								WHERE PARA_NAME = 'GLOBAL_PAGE_SIZE') / 1024)) FREE_SPACE
				FROM DBA_FREE_SPACE
				GROUP BY TABLESPACE_NAME) F,
			(SELECT TABLESPACE_NAME, ROUND(SUM(BYTES / 1048576)) TOTAL_SPACE
				FROM DBA_DATA_FILES
				GROUP BY TABLESPACE_NAME) T
		WHERE F.TABLESPACE_NAME = T.TABLESPACE_NAME;
	`
	// 表空间使用率
	SQL_TABLES_SIZE_USE_RATE = `
		SELECT A.TABLESPACE_NAME,
			A.TOTAL_MB,
			ROUND(B.FREE_MB, 2) FREE_MB,
			TO_CHAR(ROUND((A.TOTAL_MB - B.FREE_MB) / A.TOTAL_MB * 100, 2), '990.99') || '%' "USAGE %" 
		FROM (SELECT TABLESPACE_NAME, SUM(BYTES) / 1024 / 1024 TOTAL_MB
			FROM DBA_DATA_FILES GROUP BY TABLESPACE_NAME) A,
			(SELECT TABLESPACE_NAME, SUM(BYTES) / 1024 / 1024 FREE_MB
			FROM DBA_FREE_SPACE GROUP BY TABLESPACE_NAME) B,
			DBA_TABLESPACES D
		WHERE A.TABLESPACE_NAME = B.TABLESPACE_NAME(+)
		AND A.TABLESPACE_NAME = D.TABLESPACE_NAME(+)
		ORDER BY 4 DESC;
	`
	// // 查看表统计信息最后一次收集时间
	// SQL_TABLE_STATISTICS_LAST_COLLECTED = `
	// 	select LAST_GATHERED from sysstats where id=(select  id from sysobjects where name='APP_SYS_INTEGRAL_LOG');
	// `
	// // 慢SQL
	// SQL_LOW_SQL = `
	// 	select SQL_TXT,EXEC_TIME,MAX_MEM_USED/1024 from V$SQL_STAT_HISTORY where EXEC_TIME>300  order by 3 desc;
	// `
	// // 慢SQL top10
	// SQL_LOW_SQL_TOP_10 = `
	// 	select * from (SELECT sess_id,sql_text,datediff(ss,last_send_time,sysdate) ss,SF_GET_SESSION_SQL(SESS_ID) fullsql FROM V$SESSIONS WHERE STATE='ACTIVE' and sess_id <> sessid())where ROWNUM >=10
	// `

	// 内存池使用情况
	SQL_MEMORY_USAGE = `
		select 'TOTAL' as NAME,(sum(total_size)/1024/1024) as TOTAL_MB ,NULL as mem_pool 
		from v$mem_pool 
		union all
		select * from (
		select name,sum(total_size)/1024/1024 as TOTAL,count(1) as count 
		from v$mem_pool group by name order by 2 desc
		) where TOTAL>0;
	`

	// // 用户状态
	// SQL_USER_STATUS = `
	// 	select username,account_status from dba_users where account_status <> 'OPEN';
	// `
	// // 用户状态及口令修改期限
	// SQL_USER_PASSWORD_MODIFY_PERIOD = `
	// 	select b.username,a.life_time from SYSUSERS a,SYS.DBA_USERS b where a.id=b.user_id;
	// `

	// // 安全补丁
	// SQL_SECURITY_PATCH = `
	// 	select id_code();
	// `

	// // 分区上的统计信息
	// SQL_PARTITION_STATISTICS_INFO = `
	// 	select
	// 		a.table_name,
	// 		a.num_rows  ,
	// 		a.blocks    ,
	// 		a.last_analyzed,
	// 		b.partitioning_type,
	// 		b.subpartitioning_type
	// 	from
	// 		dba_tables a ,
	// 		dba_part_tables b
	// 	where
	// 		a.table_name = b.table_name;
	// `

	// // 大对象排查
	// SQL_LARGE_OBJECT = `
	// 	select type,"V$SEGMENT_INFOS".OBJ_ID from V$SEGMENT_INFOS where type in ('BLOB_SEG','CLOB_SEG');
	// `

	// // 分区表分区检查
	// SQL_PARTITION_CHECK = `
	// 	select TABLE_OWNER,TABLE_NAME,COMPOSITE,PARTITION_NAME,TABLESPACE_NAME,LAST_ANALYZED,GLOBAL_STATS,USER_STATS from dba_tab_partitions order by PARTITION_NAME;
	// `

	// // 安全相关的检查，包含弱密码、普通用户是否具有DBA权限、口令到期时间等
	// SQL_SAFETY_CHECK = `
	// 	select GRANTEE 用户,LISTAGG(GRANTED_ROLE,' ') WITHIN group (order by GRANTED_ROLE) 权限
	// 		from dba_role_privs where grantee in(
	// 	select username  from dba_users where username not in ('SYSSSO','SYSDBA','SYS','SYSAUDITOR'))
	// 	group BY DBA_ROLE_PRIVS.GRANTEE;
	// `

	// // 主备同步差异时间
	// SQL_MASTAR_BACKUP_SYNC_DIFF_TIME = `
	// 	select APPLY_CMT_TIME-LAST_CMT_TIME from v$rapply_stat;
	// `
	// // 等待事件 top5
	// SQL_WAIT_EVENT_TOP_5 = `
	// 	select top 5 * from SYS."V$SESSION_EVENT";
	// `

	// // 索引状态
	// SQL_INDEX_STATUS = `
	// 	select owner,index_name,index_type from SYS.DBA_INDEXES where status <> 'VALID' GROUP by owner,index_name,index_type;
	// `

	// 连接数
	SQL_CONNECTION_NUMS = `
		select count(*) as connection_nums from v$sessions;
	`

	// // 三个月未更新
	// SQL_NO_UPDATE_FOR_THREE_MONTH = `
	// 	select  a.owner,b.id,a.object_name,b.LAST_GATHERED from dba_objects a,sysstats b where a.OBJECT_ID=b.id and owner not in('DMHS','DMHS2','SYS','SYSDBA','SYSJOB','SYSSSO','SYSAUDITOR') and datediff(MONTH,b.LAST_GATHERED,sysdate)>3;
	// `
)
