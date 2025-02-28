{{.CSS}}
# Disk
---

- 操作系统支持：{{.AvailableArchs}}

disk 采集器用于主机磁盘信息采集，如磁盘存储空间、inodes 使用情况等。

![](imgs/input-disk-01.png)

## 前置条件

暂无

## 配置

进入 DataKit 安装目录下的 `conf.d/{{.Catalog}}` 目录，复制 `{{.InputName}}.conf.sample` 并命名为 `{{.InputName}}.conf`。示例如下：

```toml
{{.InputSample}}
```

配置好后，重启 DataKit 即可。

支持以环境变量的方式修改配置参数（只在 DataKit 以 K8s daemonset 方式运行时生效，主机部署的 DataKit 不支持此功能）：

| 环境变量名                            | 对应的配置参数项       | 参数示例                                                                                 |
| ---                                   | ---                    | ---                                                                                      |
| `ENV_INPUT_DISK_IGNORE_FS`            | `ignore_fs`            | `tmpfs,devtmpfs,devfs,iso9660,overlay,aufs,squashfs` 以英文逗号隔开                      |
| `ENV_INPUT_DISK_TAGS`                 | `tags`                 | `tag1=value1,tag2=value2` 如果配置文件中有同名 tag，会覆盖它                             |
| `ENV_INPUT_DISK_ONLY_PHYSICAL_DEVICE` | `only_physical_device` | 忽略非物理磁盘（如网盘、NFS 等，只采集本机硬盘/CD ROM/USB 磁盘等）任意给一个字符串值即可 |
| `ENV_INPUT_DISK_INTERVAL`             | `interval`             | `10s`                                                                                    |
| `ENV_INPUT_DISK_MOUNT_POINTS`         | `mount_points`         | `/, /path/to/point1, /path/to/point2` 以英文逗号隔开                                     |

## 指标预览

![](imgs/input-disk-02.png)

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

## 场景视图

<场景 - 新建仪表板 - 内置模板库 - Disk>

## 异常检测

<监控 - 模板新建 - 主机检测库>
