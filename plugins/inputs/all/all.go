package inputs

import (
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/aliyunactiontrail"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/aliyuncdn"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/aliyuncms"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/aliyuncost"

	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/aliyunddos"
	// _ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/aliyunfc"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/aliyunlog"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/aliyunobject"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/aliyunprice"

	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/aliyunrdsslowlog"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/aliyunsecurity"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/awsbill"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/awscloudtrail"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/azurecms"

	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/baiduIndex"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/cloudflare"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/confluence"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/containerd"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/coredns"

	// _ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/dataclean"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/docker_containers"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/dockerlog"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/druid"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/envoy"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/etcd"

	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/expressjs"
	// _ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/external"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/flink"
	// _ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/fluentdlog"
	// _ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/gitlab"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/goruntime"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/harborMonitor"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/hostobject"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/httpstat"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/huaweiyunces"
	// _ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/jira"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/kong"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/lighttpd"
	// _ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/mock"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/mongodboplog"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/mysqlmonitor"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/neo4j"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/nfsstat"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/pgreplication"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/prom"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/puppetagent"

	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/scanport"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/self"
	// _ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/squid"
	// _ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/ssh"
	// _ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/statsd"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/systemd"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/tailf"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/tencentcms"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/tencentcost"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/tencentobject"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/tidb"

	// _ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/timezone"
	// _ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/traceSkywalking"
	// _ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/traceZipkin"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/tracerouter"
	// _ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/traefik"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/ucmon"
	// _ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/wechatminiprogram"
	// _ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/yarn"
	// _ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/zabbix"
	//
	// // removed
	// //_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/tcpdump"
	//
	// // only windows
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/wmi"
	//
	// // 32bit disabled, only 64 bit available
	// _ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/binlog"
	//
	// // external inputs wrap
	// _ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/ansible"
	// _ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/csvmetric"
	// _ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/csvobject"
	// _ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/oraclemonitor"
	// //
	// // Buggy inputs
	// // with dll/so dependencies, and also 32bit disabled
	// // BUG: within vendor/github.com/ericchiang/k8s/watch/versioned/generated.pb.go, we should replace
	// // github.com/ericchiang.k8s.watch.versioned.Event -> k8s.io.kubernetes.pkg.watch.versioned.Event
	// //_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/prometheus"
)
