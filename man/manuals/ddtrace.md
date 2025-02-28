{{.CSS}}
# DDTrace
---

- 操作系统支持：{{.AvailableArchs}}

Datakit 内嵌的 DDTrace Agent 用于接收，运算，分析 DataDog Tracing 协议数据。

## DDTrace SDK {#sdk}

### 文档 {#docs}

- [Java](https://docs.datadoghq.com/tracing/setup_overview/setup/java?tab=containers){:target="_blank"} 
- [Python](https://docs.datadoghq.com/tracing/setup_overview/setup/python?tab=containers){:target="_blank"}
- [Ruby](https://docs.datadoghq.com/tracing/setup_overview/setup/ruby){:target="_blank"}
- [Golang](https://docs.datadoghq.com/tracing/setup_overview/setup/go?tab=containers){:target="_blank"}
- [NodeJS](https://docs.datadoghq.com/tracing/setup_overview/setup/nodejs?tab=containers){:target="_blank"}
- [PHP](https://docs.datadoghq.com/tracing/setup_overview/setup/php?tab=containers){:target="_blank"}
- [C++](https://docs.datadoghq.com/tracing/setup_overview/setup/cpp?tab=containers){:target="_blank"}
- [.Net Core](https://docs.datadoghq.com/tracing/setup_overview/setup/dotnet-core?tab=windows){:target="_blank"}
- [.Net Framework](https://docs.datadoghq.com/tracing/setup_overview/setup/dotnet-framework?tab=windows){:target="_blank"}

### 源码 {#source-code}

- [Java](https://github.com/DataDog/dd-trace-java){:target="_blank"}
- [Python](https://github.com/DataDog/dd-trace-py){:target="_blank"}
- [Ruby](https://github.com/DataDog/dd-trace-rb){:target="_blank"}
- [Golang](https://github.com/DataDog/dd-trace-go){:target="_blank"}
- [NodeJS](https://github.com/DataDog/dd-trace-js){:target="_blank"}
- [PHP](https://github.com/DataDog/dd-trace-php){:target="_blank"}
- [C++](https://github.com/opentracing/opentracing-cpp){:target="_blank"}
- [.Net](https://github.com/DataDog/dd-trace-dotnet){:target="_blank"}

> Java： DataKit 安装目录 `data` 目录下，有预先准备好的 `dd-java-agent.jar`（推荐使用）。也可以直接去 [Maven 下载](https://mvnrepository.com/artifact/com.datadoghq/dd-java-agent){:target="_blank"}

## 采集器配置 {#config}

进入 DataKit 安装目录下的 `conf.d/{{.Catalog}}` 目录，复制 `{{.InputName}}.conf.sample` 并命名为 `{{.InputName}}.conf`。示例如下：

```toml
{{.InputSample}}
```

???+ attention

    不要修改这里的 `endpoints` 列表。

    ```toml
    endpoints = ["/v0.3/traces", "/v0.4/traces", "/v0.5/traces"]
    ```

### HTTP 设置 {#http}

如果 Trace 数据是跨机器发送过来的，那么需要设置 [DataKit 的 HTTP 设置](../datakit/datakit-conf.md#config-http-server)。

如果有 ddtrace 数据发送给 DataKit，那么在 [DataKit 的 monitor](../datakit/datakit-tools-how-to.md#monitor) 上能看到：

<figure markdown>
  ![](imgs/input-ddtrace-monitor.png){ width="800" }
  <figcaption> DDtrace 将数据发送给了 /v0.4/traces 接口</figcaption>
</figure>

### 开启磁盘缓存 {#disk-cache}

如果 Trace 数据量很大，为避免给主机造成大量的资源开销，可以将 Trace 数据临时缓存到磁盘中，延迟处理：

``` toml
[inputs.ddtrace.storage]
  path = "/path/to/ddtrace-disk-storage"
  capacity = 5120
```

## DDtrace SDK 配置

配置完采集器之后，还可以对 DDtrace SDK 端做一些配置。

### 环境变量设置 {#dd-envs}

- `DD_TRACE_ENABLED`: Enable global tracer (部分语言平台支持)
- `DD_AGENT_HOST`: DDtrace agent host address
- `DD_TRACE_AGENT_PORT`: DDtrace agent host port
- `DD_SERVICE`: Service name
- `DD_TRACE_SAMPLE_RATE`: Set sampling rate
- `DD_VERSION`: Application version (optional)
- `DD_TRACE_STARTUP_LOGS`: DDtrace logger
- `DD_TRACE_DEBUG`: DDtrace debug mode
- `DD_ENV`: Application env values
- `DD_TAGS`: Application

除了在应用初始化时设置项目名，环境名以及版本号外，还可通过如下两种方式设置：

- 通过命令行注入环境变量

```shell
DD_TAGS="project:your_project_name,env=test,version=v1" ddtrace-run python app.py
```

- 在 ddtrace.conf 中直接配置自定义标签。这种方式会影响**所有**发送给 DataKit tracing 服务的数据，需慎重考虑：

```toml
## tags is ddtrace configed key value pairs
[inputs.ddtrace.tags]
  some_tag = "some_value"
  more_tag = "some_other_value"
```

### 在代码中添加业务 tag {#add-tags}

在应用代码中，可通过诸如 `span.SetTag(some-tag-key, some-tag-value)`（不同语言方式不同） 这样的方式来设置业务自定义 tag。对于这些业务自定义 tag，可通过在 ddtrace.conf 中配置 `customer_tags` 来识别并提取：

```toml
customer_tags = [
  "order_id",
  "task_id",
  "some.key",  # 被重命名为 some_key
]
```

注意，这些 tag-key 中不能包含英文字符 '.'，带 `.` 的 tag-key 会替换为 `_`。

???+ attention "应用代码中添加业务 tag 注意事项"

    - 在应用代码中添加了对应的 tag 后，必须在 ddtrace.conf 的 `customer_tags` 中也同步添加对应的 tag-key 列表，否则 DataKit 不会对这些业务 tag 进行提取
    - 在开启了采样的情况下，部分添加了 tag 的 span 有可能被舍弃

## 指标集 {#measurements}

{{ range $i, $m := .Measurements }}

{{if eq $m.Type "tracing"}}

### `{{$m.Name}}`

{{$m.Desc}}

- 标签

{{$m.TagsMarkdownTable}}

- 指标列表

{{$m.FieldsMarkdownTable}}
{{end}}

{{ end }}

## 延伸阅读 {#more-reading}

- [DataKit Tracing 字段定义](datakit-tracing-struct.md)
- [DataKit 通用 Tracing 数据采集说明](datakit-tracing.md)
- [正确使用正则表达式来配置](../datakit/datakit-input-conf.md#debug-regex) 
