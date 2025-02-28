{{.CSS}}
Tomcat
---

- 操作系统支持：{{.AvailableArchs}}

采集 tomcat 指标

## 前置条件

下载 [Jolokia](https://search.maven.org/remotecontent?filepath=org/jolokia/jolokia-war/1.6.2/jolokia-war-1.6.2.war){:target="_blank"}, 重命名为 jolokia.war, 并放置于 tomcat 的 webapps 目录下。也可从Datakit 的安装目录下的 data 目录下获取 jolokia war 包。
编辑 tomcat 的 conf 目录下的 tomcat-users.xml，增加 role 为 jolokia 的用户。

以 apache-tomcat-9.0.45 为例（示例中的 jolokia user 的 username 和 password 请务必修改！！！）:

```ssh
$ cd apache-tomcat-9.0.45/

$ export tomcat_dir=`pwd`

$ wget https://search.maven.org/remotecontent?filepath=org/jolokia/jolokia-war/1.6.2/jolokia-war-1.6.2.war \
-O $tomcat_dir/webapps/jolokia.war

$ vim $tomcat_dir/conf/tomcat-users.xml

 37 <!--
 38   <role rolename="tomcat"/>
 39   <role rolename="role1"/>
 40   <user username="tomcat" password="<must-be-changed>" roles="tomcat"/>
 41   <user username="both" password="<must-be-changed>" roles="tomcat,role1"/>
 42   <user username="role1" password="<must-be-changed>" roles="role1"/>
 43 -->
 44   <role rolename="jolokia"/>
 45   <user username="jolokia_user" password="secPassWd@123" roles="jolokia"/>
 46
 47 </tomcat-users>


$ $tomcat_dir/bin/startup.sh

 ...
 Tomcat started.
```

前往 http://localhost:8080/jolokia 查看是否配置成功

## 配置

进入 DataKit 安装目录下的 `conf.d/{{.Catalog}}` 目录，复制 `{{.InputName}}.conf.sample` 并命名为 `{{.InputName}}.conf`。示例如下：

```toml
{{.InputSample}}
```

## 指标集

以下所有数据采集，默认会追加名为 `host` 的全局 tag（tag 值为 DataKit 所在主机名），也可以在配置中通过 `[inputs.{{.InputName}}.tags]` 指定其它标签：

``` toml
 [inputs.{{.InputName}}.tags]
  # some_tag = "some_value"
  # more_tag = "some_other_value"
  # ...
```

{{ range $i, $m := .Measurements }}

### `{{$m.Name}}`

-  标签

{{$m.TagsMarkdownTable}}

- 指标列表

{{$m.FieldsMarkdownTable}}

{{ end }}

## 日志采集

如需采集 Tomcat 的日志，可在 {{.InputName}}.conf 中 将 `files` 打开，并写入 Tomcat 日志文件的绝对路径。比如：

``` toml
  [inputs.tomcat.log]
    files = ["/path_to_tomcat/logs/*"]
```

开启日志采集以后，默认会产生日志来源（`source`）为 `tomcat` 的日志。

**字段说明**

* Access Log

日志示例：

```
0:0:0:0:0:0:0:1 - admin [24/Feb/2015:15:57:10 +0530] "GET /manager/images/tomcat.gif HTTP/1.1" 200 2066
```



切割后的字段列表如下：

|字段名|字段值|说明|
|---|---|---|
|time|1424773630000000000|日志产生的时间|
|status|OK|日志等级|
|client_ip|0:0:0:0:0:0:0:1|客户端 ip|
|http_auth|admin|通过 HTTP Basic 认证的授权用户|
|http_method|GET|HTTP 方法|
|http_url|/manager/images/tomcat.gif|客户端请求地址|
|http_version|1.1|HTTP 协议版本|
|status_code|200|HTTP 状态码|
|bytes|2066|HTTP 响应 body 的字节数|

* Cataline / Host-manager / Localhost / Manager Log

日志示例：

```
06-Sep-2021 22:33:30.513 INFO [main] org.apache.catalina.startup.VersionLoggerListener.log Command line argument: -Xmx256m
```

切割后的字段列表如下：

| 字段名        | 字段值                                                | 说明                 |
| ---           | ---                                                   | ---                  |
| time          | 1630938810513000000                                   | 日志产生的时间       |
| status        | INFO                                                  | 日志等级             |
| thread_name   | main                                                  | 线程名               |
| report_source | org.apache.catalina.startup.VersionLoggerListener.log | ClassName.MethodName |
| msg           | Command line argument: -Xmx256m                       | 消息                 |

**注意**

- 日志采集仅支持采集已安装 DataKit 主机上的日志
