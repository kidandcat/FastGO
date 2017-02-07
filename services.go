package main

import "github.com/gocraft/web"

func Services(router *web.Router) {
	//Services
	//                            Redis Address
	//             endpoint name       |     Redis Password
	//                   |             |          |
	//                   v             v          v
	Controller(router, "user", "localhost:6666", "")
	Controller(router, "call", "localhost:6667", "")
}

//      Redis Example Configuration
// pidfile /var/run/redis/redis-server1.pid
// port 6666
// logfile /var/log/redis/redis-server1.log
// dbfilename dump1.rdb
// appendfilename "appendonly1.aof"
//
//
// daemonize yes
// tcp-backlog 511
// bind 0.0.0.0
// timeout 0
// tcp-keepalive 60
// loglevel notice
// databases 2
// save 900 1
// save 300 10
// save 60 10000
// stop-writes-on-bgsave-error yes
// rdbcompression yes
// rdbchecksum yes
// dir /var/lib/redis
// appendonly yes
// appendfsync everysec
// no-appendfsync-on-rewrite no
// auto-aof-rewrite-percentage 100
// auto-aof-rewrite-min-size 64mb
// aof-load-truncated yes
// lua-time-limit 5000
// slowlog-log-slower-than 10000
// slowlog-max-len 128
// latency-monitor-threshold 0
// notify-keyspace-events ""
// aof-rewrite-incremental-fsync yes
