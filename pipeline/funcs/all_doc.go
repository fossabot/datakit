// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package funcs

import "strings"

type PLDoc struct {
	Doc             string `json:"doc"`
	Prototype       string `json:"prototype"`
	Description     string `json:"description"`
	Deprecated      bool   `json:"deprecated"`
	RequiredVersion string `json:"required_version"`
}

var PipelineFunctionDocs = map[string]*PLDoc{
	"add_key()":            &addKeyMarkdown,
	"add_pattern()":        &addPatternMarkdown,
	"adjust_timezone()":    &adjustTimezoneMarkdown,
	"cast()":               &castMarkdown,
	"cover()":              &coverMarkdown,
	"datetime()":           &datetimeMarkdown,
	"default_time()":       &defaultTimeMarkdown,
	"drop()":               &dropMarkdown,
	"drop_key()":           &dropKeyMarkdown,
	"drop_origin_data()":   &dropOriginDataMarkdown,
	"duration_precision()": &durationPrecisionMarkdown,
	"exit()":               &exitMarkdown,
	"geoip()":              &geoIPMarkdown,
	"grok()":               &grokMarkdown,
	"group_between()":      &groupBetweenMarkdown,
	"group_in()":           &groupInMarkdown,
	"json()":               &jsonMarkdown,
	"lowercase()":          &lowercaseMarkdown,
	"nullif()":             &nullIfMarkdown,
	"parse_date()":         &parseDateMarkdown,
	"parse_duration()":     &parseDurationMarkdown,
	"rename()":             &renameMarkdown,
	"replace()":            &replaceMarkdown,
	"set_tag()":            &setTagMarkdown,
	"strfmt()":             &strfmtMarkdown,
	"uppercase()":          &uppercaseMarkdown,
	"url_decode()":         &URLDecodeMarkdown,
	"user_agent()":         &userAgentMarkdown,
	"match":                &matchMarkdown,
	"decode":               &decodeMarkdown,
}

//nolint:lll
var (
	addPatternMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `add_pattern()`",
			"",
			"函数原型：`add_pattern(name=required, pattern=required)`",
			"",
			"函数说明：创建自定义 grok 模式。grok 模式有作用域限制, 如在 if else 语句内将产生新的作用域, 该 pattern 仅在此作用域内有效。该函数不可覆盖同一作用域或者上一作用域已经存在的 grok 模式",
			"",
			"参数:",
			"",
			"- `name`：模式命名",
			"- `pattern`: 自定义模式内容",
			"",
			"示例:",
			"",
			"```python",
			`# 待处理数据: "11,abc,end1", "22,abc,end1", "33,abc,end3"`,
			"",
			"# pipline脚本",
			`add_pattern("aa", "\\d{2}")`,
			`grok(_, "%{aa:aa}")`,
			`if false {`,
			``,
			`} else {`,
			`    add_pattern("bb", "[a-z]{3}")`,
			`    if aa == "11" {`,
			`        add_pattern("cc", "end1")`,
			`        grok(_, "%{aa:aa},%{bb:bb},%{cc:cc}")`,
			`    } elif aa == "22" {`,
			`        # 此处使用 pattern cc 将导致编译失败: no pattern found for %{cc}`,
			`        grok(_, "%{aa:aa},%{bb:bb},%{INT:cc}")`,
			`    } elif aa == "33" {`,
			`        add_pattern("bb", "[\\d]{5}")	# 此处覆盖 bb 失败`,
			`        add_pattern("cc", "end3")`,
			`        grok(_, "%{aa:aa},%{bb:bb},%{cc:cc}")`,
			`    }`,
			`}`,
			"",
			"# 处理结果",
			`{`,
			`    "aa":      "11"`,
			`    "bb":      "abc"`,
			`    "cc":      "end1"`,
			`    "message": "11,abc,end1"`,
			`}`,
			`{`,
			`    "aa":      "22"`,
			`	 "message": "22,abc,end1"`,
			`}`,
			`{`,
			`    "aa":      "33"`,
			`    "bb":      "abc"`,
			`    "cc":      "end3"`,
			`    "message": "33,abc,end3"`,
			`}`,
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	grokMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `grok()`",
			"",
			"函数原型：`grok(input=required, pattern=required)`",
			"",
			"函数说明：通过 `pattern` 提取文本串 `input` 中的内容。",
			"",
			"参数:",
			"",
			"- `input`：待提取文本，可以是原始文本（`_`）或经过初次提取之后的某个 `key`",
			"- `pattern`: grok 表达式",
			"",
			"```python",
			"grok(_, pattern)    # 直接使用输入的文本作为原始数据",
			"grok(key, pattern)  # 对之前已经提取出来的某个 key，做再次 grok",
			"```",
			"",
			"示例:",
			"",
			"```python",
			`# 待处理数据: "21:13:14"`,
			"",
			"# pipline脚本",
			`add_pattern("_second", "(?:(?:[0-5]?[0-9]|60)(?:[:.,][0-9]+)?)")`,
			`add_pattern("_minute", "(?:[0-5][0-9])")`,
			`add_pattern("_hour", "(?:2[0123]|[01]?[0-9])")`,
			`add_pattern("time", "([^0-9]?)%{HOUR:hour}:%{MINUTE:minute}(?::%{SECOND:second})([^0-9]?)")`,
			`grok(_, "%{time}")`,
			"",
			"# 处理结果",
			"{",
			`    "hour":"12",`,
			`    "minute":"13",`,
			`    "second":"14",`,
			`    "message":"21:13:14"`,
			"}",
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	jsonMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `json()`",
			"",
			"函数原型：`json(input=required, jsonPath=required, newkey=optional)`",
			"",
			"函数说明：提取 json 中的指定字段，并可将其命名成新的字段。",
			"",
			"参数:",
			"",
			"- `input`: 待提取 json，可以是原始文本（`_`）或经过初次提取之后的某个 `key`",
			"- `jsonPath`: json 路径信息",
			"- `newkey`：提取后数据写入新 key",
			"",
			"```python",
			"# 直接提取原始输入 json 中的x.y字段，并可将其命名成新字段abc",
			"json(_, x.y, abc)",
			"",
			"# 已提取出的某个 `key`，对其再提取一次 `x.y`，提取后字段名为 `x.y`",
			"json(key, x.y) ",
			"```",
			"",
			"示例一:",
			"",
			"```python",
			`# 待处理数据: {"info": {"age": 17, "name": "zhangsan", "height": 180}}`,
			"",
			"# 处理脚本",
			`json(_, info, "zhangsan")`,
			"json(zhangsan, name)",
			`json(zhangsan, age, "年龄")`,
			"",
			"# 处理结果",
			"{",
			`    "message": "{\"info\": {\"age\": 17, \"name\": \"zhangsan\", \"height\": 180}}",`,
			`    "zhangsan": {`,
			`        "age": 17,`,
			`        "height": 180,`,
			`        "name": "zhangsan"`,
			"    }",
			"}",
			"```",
			"",
			"示例二:",
			"",
			"```python",
			"# 待处理数据",
			"#    data = {",
			`#        "name": {"first": "Tom", "last": "Anderson"},`,
			`#        "age":37,`,
			`#        "children": ["Sara","Alex","Jack"],`,
			`#        "fav.movie": "Deer Hunter",`,
			`#        "friends": [`,
			`#            {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},`,
			`#            {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},`,
			`#            {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}`,
			"#        ]",
			"#    }",
			"",
			"# 处理脚本",
			"json(_, name) json(name, first)",
			"```",
			"",
			"示例三:",
			"",
			"```python",
			"# 待处理数据",
			"#    [",
			`#            {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},`,
			`#            {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},`,
			`#            {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}`,
			"#    ]",
			"    ",
			"# 处理脚本, json数组处理",
			"json(_, [0].nets[-1])",
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	renameMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `rename()`",
			"",
			"函数原型：`rename(new-key=required, old-key=required)`",
			"",
			"函数说明：将已提取的字段重新命名",
			"",
			"参数:",
			"",
			"- `new-key`: 新字段名",
			"- `old-key`: 已提取的字段名",
			"",
			"```python",
			"# 把已提取的 abc 字段重新命名为 abc1",
			"rename('abc1', abc)",
			"",
			"# or ",
			"",
			"rename(abc1, abc)",
			"```",
			"",
			"示例：",
			"",
			"```python",
			`# 待处理数据: {"info": {"age": 17, "name": "zhangsan", "height": 180}}`,
			"",
			"# 处理脚本",
			`json(_, info.name, "姓名")`,
			"",
			"# 处理结果",
			"{",
			`  "message": "{\"info\": {\"age\": 17, \"name\": \"zhangsan\", \"height\": 180}}",`,
			`  "zhangsan": {`,
			`    "age": 17,`,
			`    "height": 180,`,
			`    "姓名": "zhangsan"`,
			"  }",
			"}",
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	URLDecodeMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `url_decode()`",
			"",
			"函数原型：`url_decode(key=required)`",
			"",
			"函数说明：将已提取 `key` 中的 URL 解析成明文",
			"",
			"参数:",
			"",
			"- `key`: 已经提取的某个 `key`",
			"",
			"示例：",
			"",
			"```python",
			`# 待处理数据: {"url":"http%3a%2f%2fwww.baidu.com%2fs%3fwd%3d%e6%b5%8b%e8%af%95"}`,
			"",
			"# 处理脚本",
			"json(_, url) url_decode(url)",
			"",
			"# 处理结果",
			"{",
			`  "message": "{"url":"http%3a%2f%2fwww.baidu.com%2fs%3fwd%3d%e6%b5%8b%e8%af%95"}",`,
			`  "url": "http://www.baidu.com/s?wd=测试"`,
			"}",
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	geoIPMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `geoip()`",
			"",
			"函数原型：`geoip(key=required)`",
			"",
			"函数说明：在 IP 上追加更对 geo 信息。 `geoip()` 会额外产生多个字段，如：",
			"",
			"- `isp`: 运营商",
			"- `city`: 城市",
			"- `province`: 省份",
			"- `country`: 国家",
			"",
			"参数:",
			"",
			"- `key`: 已经提取出来的 IP 字段，支持 IPv4/6",
			"",
			"示例：",
			"",
			"```python",
			`# 待处理数据: {"ip":"116.228.89.206"}`,
			"",
			"# 处理脚本",
			"json(_, ip) geoip(ip)",
			"",
			"# 处理结果",
			"{",
			`  "message": "{"ip":"116.228.89.206"}",`,
			`  "isp": "xxxx",`,
			`  "city": "xxxx",`,
			`  "province": "xxxx",`,
			`  "country": "xxxx"`,
			"}",
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	datetimeMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `datetime()`",
			"",
			"函数原型：`datetime(key=required, precision=required, fmt=required)`",
			"",
			"函数说明：将时间戳转成指定日期格式",
			"",
			"函数参数",
			"",
			"- `key`: 已经提取的时间戳 (必选参数)",
			"- `precision`：输入的时间戳精度(s, ms)",
			"- `fmt`：日期格式，时间格式, 支持以下模版",
			"",
			"```python",
			`ANSIC       = "Mon Jan _2 15:04:05 2006"`,
			`UnixDate    = "Mon Jan _2 15:04:05 MST 2006"`,
			`RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"`,
			`RFC822      = "02 Jan 06 15:04 MST"`,
			`RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone`,
			`RFC850      = "Monday, 02-Jan-06 15:04:05 MST"`,
			`RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"`,
			`RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone`,
			`RFC3339     = "2006-01-02T15:04:05Z07:00"`,
			`RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"`,
			`Kitchen     = "3:04PM"`,
			"```",
			"",
			"示例:",
			"",
			"```python",
			"# 待处理数据: ",
			"#    {",
			`#        "a":{`,
			`#            "timestamp": "1610960605000",`,
			`#            "second":2`,
			"#        },",
			`#        "age":47`,
			"#    }",
			"",
			"# 处理脚本",
			"json(_, a.timestamp) datetime(a.timestamp, 'ms', 'RFC3339')",
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	castMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `cast()`",
			"",
			"函数原型：`cast(key=required, type=required)`",
			"",
			"函数说明：将 key 值转换拆成指定类型",
			"",
			"函数参数",
			"",
			"- `key`: 已提取的某字段",
			"- `type`：转换的目标类型，支持 `\"str\", \"float\", \"int\", \"bool\"` 这几种，目标类型需要用英文状态双引号括起来",
			"",
			"示例:",
			"",
			"```python",
			`# 待处理数据: {"first": 1,"second":2,"third":"aBC","forth":true}`,
			"",
			"# 处理脚本",
			`json(_, first) cast(first, "str")`,
			"",
			"# 处理结果",
			"{",
			`  "first":"1"`,
			"}",
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	groupBetweenMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `group_between()`",
			"",
			"函数原型：`group_between(key=required, between=required, new-value=required, new-key=optional)`",
			"",
			"函数说明：如果 `key` 值在指定范围 `between` 内（注意：只能是单个区间，如 `[0,100]`），则可创建一个新字段，并赋予新值。若不提供新字段，则覆盖原字段值",
			"",
			"示例一:",
			"",
			"```python",
			`# 待处理数据: {"http_status": 200, "code": "success"}`,
			"",
			"json(_, http_status)",
			"",
			`# 如果字段 http_status 值在指定范围内，则将其值改为 "OK"`,
			`group_between(http_status, [200, 300], "OK")`,
			"`",
			"",
			"# 处理结果",
			"{",
			`    "http_status": "OK"`,
			"}",
			"```",
			"",
			"示例二:",
			"",
			"```python",
			`# 待处理数据: {"http_status": 200, "code": "success"}`,
			"",
			"json(_, http_status)",
			"",
			`# 如果字段 http_status 值在指定范围内，则新建 status 字段，其值为 "OK"`,
			`group_between(http_status, [200, 300], "OK", status)`,
			"",
			"# 处理结果",
			"{",
			`    "http_status": 200,`,
			`    "status": "OK"`,
			"}",
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	groupInMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `group_in()`",
			"",
			"函数原型：`group_in(key=required, in=required, new-value=required, new-key=optional)`",
			"",
			"函数说明：如果 `key` 值在列表 `in` 中，则可创建一个新字段，并赋予新值。若不提供新字段，则覆盖原字段值",
			"",
			"示例:",
			"",
			"```python",
			`# 如果字段 log_level 值在列表中，则将其值改为 "OK"`,
			`group_in(log_level, ["info", "debug"], "OK")`,
			"",
			`# 如果字段 http_status 值在指定列表中，则新建 status 字段，其值为 "not-ok"`,
			`group_in(log_level, ["error", "panic"], "not-ok", status)`,
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	uppercaseMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `uppercase()`",
			"",
			"函数原型：`uppercase(key=required)`",
			"",
			"函数说明：将已提取 key 中内容转换成大写",
			"",
			"函数参数",
			"",
			"- `key`: 指定已提取的待转换字段名，将 `key` 内容转成大写",
			"",
			"示例:",
			"",
			"```python",
			`# 待处理数据: {"first": "hello","second":2,"third":"aBC","forth":true}`,
			"",
			"# 处理脚本",
			"json(_, first) uppercase(first)",
			"",
			"# 处理结果",
			"{",
			`   "first": "HELLO"`,
			"}",
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	lowercaseMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `lowercase()`",
			"",
			"函数原型：`lowercase(key=required)`",
			"",
			"函数说明：将已提取 key 中内容转换成小写",
			"",
			"函数参数",
			"",
			"- `key`: 指定已提取的待转换字段名",
			"",
			"示例:",
			"",
			"```python",
			`# 待处理数据: {"first": "HeLLo","second":2,"third":"aBC","forth":true}`,
			"",
			"# 处理脚本",
			"json(_, first) lowercase(first)",
			"",
			"# 处理结果",
			"{",
			`    "first": "hello"`,
			"}",
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	nullIfMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `nullif()`",
			"",
			"函数原型：`nullif(key=required, value=required)`",
			"",
			"函数说明：若已提取 `key` 指定的字段内容等于 `value` 值，则删除此字段",
			"",
			"函数参数",
			"",
			"- `key`: 指定字段",
			"- `value`: 目标值",
			"",
			"示例:",
			"",
			"```python",
			`# 待处理数据: {"first": 1,"second":2,"third":"aBC","forth":true}`,
			"",
			"# 处理脚本",
			`json(_, first) json(_, second) nullif(first, "1")`,
			"",
			"# 处理结果",
			"{",
			`    "second":2`,
			"}",
			"```",
			"",
			"> 注：该功能可通过 `if/else` 语义来实现：",
			"",
			"```python",
			`if first == "1" {`,
			"	drop_key(first)",
			"}",
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	strfmtMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `strfmt()`",
			"",
			"函数原型：`strfmt(key=required, fmt=required, key1=optional, key2, ...)`",
			"",
			"函数说明：对已提取 `key1,key2...` 指定的字段内容根据 `fmt` 进行格式化，并把格式化后的内容写入 `key` 字段中",
			"",
			"函数参数",
			"",
			"- `key`: 指定格式化后数据写入字段名",
			"- `fmt`: 格式化字符串模板",
			"- `key1，key2`:已提取待格式化字段名",
			"",
			"示例:",
			"",
			"```python",
			`# 待处理数据: {"a":{"first":2.3,"second":2,"third":"abc","forth":true},"age":47}`,
			"",
			"# 处理脚本",
			"json(_, a.second)",
			"json(_, a.thrid)",
			`cast(a.second, "int")`,
			"json(_, a.forth)",
			`strfmt(bb, "%v %s %v", a.second, a.thrid, a.forth)`,
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	dropOriginDataMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `drop_origin_data()`",
			"",
			"函数原型：`drop_origin_data()`",
			"",
			"函数说明：丢弃初始化文本，否则初始文本放在 message 字段中",
			"",
			"示例:",
			"",
			"```python",
			`# 待处理数据: {"age": 17, "name": "zhangsan", "height": 180}`,
			"",
			"# 结果集中删除 message 内容",
			"drop_origin_data()",
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	addKeyMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `add_key()`",
			"",
			"函数原型：`add_key(key-name=required, key-value=required)`",
			"",
			"函数说明：增加一个字段",
			"",
			"函数参数",
			"",
			"- `key-name`: 新增的 key 名称",
			"- `key-value`：key 值（只能是 string/number/bool/nil 这几种类型）",
			"",
			"示例:",
			"",
			"```python",
			`# 待处理数据: {"age": 17, "name": "zhangsan", "height": 180}`,
			"",
			"# 处理脚本",
			`add_key(city, "shanghai")`,
			"",
			"# 处理结果",
			"{",
			`    "age": 17,`,
			`    "height": 180,`,
			`    "name": "zhangsan",`,
			`    "city": "shanghai"`,
			"}",
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	defaultTimeMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `default_time()`",
			"",
			"函数原型：`default_time(key=required, timezone=optional)`",
			"",
			"函数说明：以提取的某个字段作为最终数据的时间戳",
			"",
			"函数参数",
			"",
			"- `key`: 指定的 key",
			"- `timezone`: 指定的时区，默认为本机当前时区",
			"",
			"待处理数据支持以下格式化时间",
			"",
			"| 日期格式                                           | 日期格式                                                | 日期格式                                       | 日期格式                          |",
			"| -----                                              | ----                                                    | ----                                           | ----                              |",
			"| `2014-04-26 17:24:37.3186369`                      | `May 8, 2009 5:57:51 PM`                                | `2012-08-03 18:31:59.257000000`                | `oct 7, 1970`                     |",
			"| `2014-04-26 17:24:37.123`                          | `oct 7, '70`                                            | `2013-04-01 22:43`                             | `oct. 7, 1970`                    |",
			"| `2013-04-01 22:43:22`                              | `oct. 7, 70`                                            | `2014-12-16 06:20:00 UTC`                      | `Mon Jan  2 15:04:05 2006`        |",
			"| `2014-12-16 06:20:00 GMT`                          | `Mon Jan  2 15:04:05 MST 2006`                          | `2014-04-26 05:24:37 PM`                       | `Mon Jan 02 15:04:05 -0700 2006`  |",
			"| `2014-04-26 13:13:43 +0800`                        | `Monday, 02-Jan-06 15:04:05 MST`                        | `2014-04-26 13:13:43 +0800 +08`                | `Mon, 02 Jan 2006 15:04:05 MST`   |",
			"| `2014-04-26 13:13:44 +09:00`                       | `Tue, 11 Jul 2017 16:28:13 +0200 (CEST)`                | `2012-08-03 18:31:59.257000000 +0000 UTC`      | `Mon, 02 Jan 2006 15:04:05 -0700` |",
			"| `2015-09-30 18:48:56.35272715 +0000 UTC`           | `Thu, 4 Jan 2018 17:53:36 +0000`                        | `2015-02-18 00:12:00 +0000 GMT`                | `Mon 30 Sep 2018 09:09:09 PM UTC` |",
			"| `2015-02-18 00:12:00 +0000 UTC`                    | `Mon Aug 10 15:44:11 UTC+0100 2015`                     | `2015-02-08 03:02:00 +0300 MSK m=+0.000000001` | `Thu, 4 Jan 2018 17:53:36 +0000`  |",
			"| `2015-02-08 03:02:00.001 +0300 MSK m=+0.000000001` | `Fri Jul 03 2015 18:04:07 GMT+0100 (GMT Daylight Time)` | `2017-07-19 03:21:51+00:00`                    | `September 17, 2012 10:09am`      |",
			"| `2014-04-26`                                       | `September 17, 2012 at 10:09am PST-08`                  | `2014-04`                                      | `September 17, 2012, 10:10:09`    |",
			"| `2014`                                             | `2014:3:31`                                             | `2014-05-11 08:20:13,787`                      | `2014:03:31`                      |",
			"| `3.31.2014`                                        | `2014:4:8 22:05`                                        | `03.31.2014`                                   | `2014:04:08 22:05`                |",
			"| `08.21.71`                                         | `2014:04:2 03:00:51`                                    | `2014.03`                                      | `2014:4:02 03:00:51`              |",
			"| `2014.03.30`                                       | `2012:03:19 10:11:59`                                   | `20140601`                                     | `2012:03:19 10:11:59.3186369`     |",
			"| `20140722105203`                                   | `2014年04月08日`                                        | `1332151919`                                   | `2006-01-02T15:04:05+0000`        |",
			"| `1384216367189`                                    | `2009-08-12T22:15:09-07:00`                             | `1384216367111222`                             | `2009-08-12T22:15:09`             |",
			"| `1384216367111222333`                              | `2009-08-12T22:15:09Z`                                  |",
			"",
			"JSON 提取示例:",
			"",
			"```python",
			"# 原始 json",
			"{",
			`    "time":"06/Jan/2017:16:16:37 +0000",`,
			`    "second":2,`,
			`    "third":"abc",`,
			`    "forth":true`,
			"}",
			"",
			"# pipeline 脚本",
			"json(_, time)      # 提取 time 字段",
			"default_time(time) # 将提取到的 time 字段转换成时间戳",
			"",
			"# 处理结果",
			"{",
			`  "time": 1483719397000000000,`,
			"}",
			"```",
			"",
			"文本提取示例:",
			"",
			"```python",
			"# 原始日志文本",
			"2021-01-11T17:43:51.887+0800  DEBUG io  io/io.go:458  post cost 6.87021ms",
			"",
			"# pipeline 脚本",
			"grok(_, '%{TIMESTAMP_ISO8601:log_time}')   # 提取日志时间，并将字段命名为 log_time",
			"default_time(log_time)                     # 将提取到的 log_time 字段转换成时间戳",
			"",
			"# 处理结果",
			"{",
			`  "log_time": 1610358231887000000,`,
			"}",
			"",
			"# 对于 logging 采集的数据，最好将时间字段命名为 time，否则 logging 采集器会以当前时间填充",
			`rename("time", log_time)`,
			"",
			"# 处理结果",
			"{",
			`  "time": 1610358231887000000,`,
			"}",
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	adjustTimezoneMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `adjust_timezone()`",
			"",
			"函数原型：`adjust_timezone(key=required)`",
			"",
			"函数说明：自动选择时区，校准时间戳。用于校准日志中的时间格式不带时区信息，且与 pipeline 时间处理函数默认的本地时区不一致时使得时间戳出现数小时的偏差，适用于时间偏差小于24小时",
			"",
			"函数参数",
			"",
			"- `key`: 纳秒时间戳，如 default_time(time) 函数处理后得到的时间戳",
			"",
			"示例:",
			"",
			"```python",
			"# 原始 json",
			"{",
			`    "time":"10 Dec 2021 03:49:20.937", `,
			`    "second":2,`,
			`    "third":"abc",`,
			`    "forth":true`,
			"}",
			"",
			"# pipeline 脚本",
			"json(_, time)      # 提取 time 字段 (若容器中时区 UTC+0000)",
			"default_time(time) # 将提取到的 time 字段转换成时间戳 ",
			"                   # (对无时区数据使用本地时区 UTC+0800/UTC+0900...解析)",
			"adjust_timezone(time)",
			"                   # 自动(重新)选择时区，校准时间偏差",
			"# 处理结果",
			"{",
			`  "time": 1639108160937000000,`,
			"}",
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	dropKeyMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `drop_key()`",
			"",
			"函数原型：`drop_key(key=required)`",
			"",
			"函数说明：删除已提取字段",
			"",
			"函数参数",
			"",
			"- `key`: 待删除字段名",
			"",
			"示例:",
			"",
			"```python",
			"data = `{\"age\": 17, \"name\": \"zhangsan\", \"height\": 180}`",
			"",
			"# 处理脚本",
			"json(_, age,)",
			"json(_, name)",
			"json(_, height)",
			"drop_key(height)",
			"",
			"# 处理结果",
			"{",
			`    "age": 17,`,
			`    "name": "zhangsan"`,
			"}",
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	userAgentMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `user_agent()`",
			"",
			"函数原型：`user_agent(key=required)`",
			"",
			"函数说明：对指定字段上获取客户端信息",
			"",
			"函数参数",
			"",
			"- `key`: 待提取字段",
			"",
			"`user_agent()` 会生产多个字段，如：",
			"",
			"- `os`: 操作系统",
			"- `browser`: 浏览器",
			"",
			"示例:",
			"",
			"```python",
			"# 待处理数据",
			"#    {",
			`#        "userAgent" : "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.125 Safari/537.36",`,
			`#        "second"    : 2,`,
			`#        "third"     : "abc",`,
			`#        "forth"     : true`,
			"#    }",
			"",
			"json(_, userAgent) user_agent(userAgent)",
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	parseDurationMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `parse_duration()`",
			"",
			"函数原型：`parse_duration(key=required)`",
			"",
			"函数说明：如果 `key` 的值是一个 golang 的 duration 字符串（如 `123ms`），则自动将 `key` 解析成纳秒为单位的整数",
			"",
			"目前 golang 中的 duration 单位如下：",
			"",
			"- `ns` 纳秒",
			"- `us/µs` 微秒",
			"- `ms` 毫秒",
			"- `s` 秒",
			"- `m` 分钟",
			"- `h` 小时",
			"",
			"函数参数",
			"",
			"- `key`: 待解析的字段",
			"",
			"示例:",
			"",
			"```python",
			`# 假定 abc = "3.5s"`,
			"parse_duration(abc) # 结果 abc = 3500000000",
			"",
			`# 支持负数: abc = "-3.5s"`,
			"parse_duration(abc) # 结果 abc = -3500000000",
			"",
			`# 支持浮点: abc = "-2.3s"`,
			"parse_duration(abc) # 结果 abc = -2300000000",
			"",
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	parseDateMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `parse_date()`",
			"",
			"函数原型：`parse_date(new-key=required, yy=require, MM=require, dd=require, hh=require, mm=require, ss=require, ms=require, zone=require)`",
			"",
			"函数说明：将传入的日期字段各部分的值转化为时间戳",
			"",
			"函数参数",
			"",
			"- `key`: 新插入的字段",
			"- `yy` : 年份数字字符串，支持四位或两位数字字符串，为空字符串，则处理时取当前年份",
			"- `MM`:  月份字符串, 支持数字，英文，英文缩写",
			"- `dd`: 日字符串",
			"- `hh`: 小时字符串",
			"- `mm`: 分钟字符串",
			"- `ss`: 秒字符串",
			"- `ms`: 毫秒字符串",
			"- `zone`: 时区字符串，“+8”或\"Asia/Shanghai\"形式",
			"",
			"示例:",
			"",
			"```python",
			`parse_date(aa, "2021", "May", "12", "10", "10", "34", "", "Asia/Shanghai") # 结果 aa=1620785434000000000`,
			"",
			`parse_date(aa, "2021", "12", "12", "10", "10", "34", "", "Asia/Shanghai") # 结果 aa=1639275034000000000`,
			"",
			`parse_date(aa, "2021", "12", "12", "10", "10", "34", "100", "Asia/Shanghai") # 结果 aa=1639275034000000100`,
			"",
			`parse_date(aa, "20", "February", "12", "10", "10", "34", "", "+8") 结果 aa=1581473434000000000`,
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	coverMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `cover()`",
			"",
			"函数原型：`cover(key=required, range=require)`",
			"",
			"函数说明：对指定字段上获取的字符串数据，按范围进行数据脱敏处理",
			"",
			"函数参数",
			"",
			"- `key`: 待提取字段",
			"- `range`: 脱敏字符串的索引范围（`[start,end]`），如 `[3,5]` 这种形式，这里是一个闭合区间, **下标从 1 开始**。索引不区分中文半角和全角。另外，`start` 必须大于 `end`，如果长度不固定，可以用一个极大的整数来表示，如 `[1, 999999999]` 表示整个字符串都覆盖掉。另外，**脱敏只对字符串值有效，对其它类型的值脱敏，其行为暂未定义**",
			"",
			"示例:",
			"",
			"```python",
			`# 待处理数据 {"str": "13789123014"}`,
			"json(_, str) cover(str, [8, 13])",
			"",
			`# 待处理数据 {"str": "13789123014"}`,
			"json(_, str) cover(str, [2, 4])",
			"",
			`# 待处理数据 {"str": "13789123014"}`,
			"json(_, str) cover(str, [1, 1])",
			"",
			`# 待处理数据 {"str": "小阿卡"}`,
			"json(_, str) cover(str, [2, 2])",
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	replaceMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `replace()`",
			"",
			"函数原型：`replace(key=required, regex=required, replaceStr=required)`",
			"",
			"函数说明：对指定字段上获取的字符串数据按正则进行替换",
			"",
			"函数参数",
			"",
			"- `key`: 待提取字段",
			"- `regex`: 正则表达式",
			"- `replaceStr`: 替换的字符串",
			"",
			"示例:",
			"",
			"```python",
			`# 电话号码：{"str": "13789123014"}`,
			"json(_, str)",
			`replace(str, "(1[0-9]{2})[0-9]{4}([0-9]{4})", "$1****$2")`,
			"",
			`# 英文名 {"str": "zhang san"}`,
			"json(_, str)",
			`replace(str, "([a-z]*) \\w*", "$1 ***")`,
			"",
			`# 身份证号 {"str": "362201200005302565"}`,
			"json(_, str)",
			`replace(str, "([1-9]{4})[0-9]{10}([0-9]{4})", "$1**********$2")`,
			"",
			`# 中文名 {"str": "小阿卡"}`,
			"json(_, str)",
			`replace(str, '([\u4e00-\u9fa5])[\u4e00-\u9fa5]([\u4e00-\u9fa5])', "$1＊$2")`,
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	setTagMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `set_tag()`",
			"",
			"函数原型：`set_tag(key=required, value=optional)`",
			"",
			"函数说明：对指定字段标记为 tag 输出，设置为 tag 后，其他函数仍可对该变量操作。如果被置为 tag 的 key 是已经切割出来的 field，那么它将不会在 field 中出现，这样可以避免切割出来的 field key 跟已有数据上的 tag key 重名",
			"",
			"函数参数",
			"",
			"- `key`: 待标记为 tag 的字段",
			"- `value`: 可以为字符串字面量或者变量",
			"",
			"```python",
			`# in << {"str": "13789123014"}`,
			"set_tag(str)",
			`json(_, str)          # str == "13789123014"`,
			`replace(str, "(1[0-9]{2})[0-9]{4}([0-9]{4})", "$1****$2")`,
			"# Extracted data(drop: false, cost: 49.248µs):",
			"# {",
			`#   "message": "{\"str\": \"13789123014\", \"str_b\": \"3\"}",`,
			`#   "str#": "137****3014"`,
			"# }",
			"# * 字符 `#` 仅为 datakit --pl <path> --txt <str> 输出展示时字段为 tag 的标记",
			"",
			`# in << {"str_a": "2", "str_b": "3"}`,
			"json(_, str_a)",
			`set_tag(str_a, "3")   # str_a == 3`,
			"# Extracted data(drop: false, cost: 30.069µs):",
			"# {",
			`#   "message": "{\"str_a\": \"2\", \"str_b\": \"3\"}",`,
			`#   "str_a#": "3"`,
			"# }",
			"",
			"",
			`# in << {"str_a": "2", "str_b": "3"}`,
			"json(_, str_a)",
			"json(_, str_b)",
			`set_tag(str_a, str_b) # str_a == str_b == "3"`,
			"# Extracted data(drop: false, cost: 32.903µs):",
			"# {",
			`#   "message": "{\"str_a\": \"2\", \"str_b\": \"3\"}",`,
			`#   "str_a#": "3",`,
			`#   "str_b": "3"`,
			"# }",
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	dropMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `drop()`",
			"",
			"函数原型：`drop()`",
			"",
			"函数说明：丢弃整条日志，不进行上传",
			"",
			"```python",
			`# in << {"str_a": "2", "str_b": "3"}`,
			"json(_, str_a)",
			`if str_a == "2"{`,
			"  drop()",
			"  exit()",
			"}",
			"json(_, str_b)",
			"",
			"# Extracted data(drop: true, cost: 30.02µs):",
			"# {",
			`#   "message": "{\"str_a\": \"2\", \"str_b\": \"3\"}",`,
			`#   "str_a": "2"`,
			"# }",
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	exitMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `exit()`",
			"",
			"函数原型：`exit()`",
			"",
			"函数说明：结束当前一条日志的解析，若未调用函数 drop() 仍会输出已经解析的部分",
			"",
			"```python",
			`# in << {"str_a": "2", "str_b": "3"}`,
			"json(_, str_a)",
			`if str_a == "2"{`,
			"  exit()",
			"}",
			"json(_, str_b)",
			"",
			"# Extracted data(drop: false, cost: 48.233µs):",
			"# {",
			`#   "message": "{\"str_a\": \"2\", \"str_b\": \"3\"}",`,
			`#   "str_a": "2"`,
			"# }",
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	durationPrecisionMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `duration_precision()`",
			"",
			"函数原型：`duration_precision(key=required, old_precision=require, new_precision=require)`",
			"",
			"函数说明：进行 duration 精度的转换，通过参数指定当前精度和目标精度。支持在 s, ms, us, ns 间转换。",
			"",
			"```python",
			`# in << {"ts":12345}`,
			"json(_, ts)",
			`cast(ts, "int")`,
			`duration_precision(ts, "ms", "ns")`,
			"",
			"# Extracted data(drop: false, cost: 33.279µs):",
			"# {",
			`#   "message": "{\"ts\":12345}",`,
			`#   "ts": 12345000000`,
			"# }",
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	matchMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `match()`",
			"",
			"函数原型：`match(words,pattern)`",
			"",
			"函数说明：看words是否符合pattern正则出来的",
			"",
			"```python",
			`# in << {"str_a": "2", "str_b": "3"}`,
			`if match("peech", "p([a-z]+)ch") {`,
			`json(_, str_a)`,
			`if str_a == "2"{`,
			`drop()`,
			`exit()`,
			`}`,
			`json(_, str_b)`,
			`}`,
			"",
			"# Extracted data(drop: false, cost: 33.279µs):",
			"# {",
			`#   "message": "{\"str_a\": \"2\", \"str_b\": \"3\"}",`,
			`#   "str_a": "2"`,
			"# }",
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
	decodeMarkdown = PLDoc{
		Doc: strings.Join([]string{
			"### `decode()`",
			"",
			"函数原型：`decode(text,textCode)`",
			"",
			"函数说明：把text变成utf-8编码",
			"",
			"```python",
			`decode("wwwwww","gbk")`,
			"",
			"# Extracted data(drop: false, cost: 33.279µs):",
			"# {",
			`#   "message": "wwwwww",`,
			`#   "changed": "wwwwww"`,
			"# }",
			"```",
			"",
		}, "\n"),
		Deprecated: false,
	}
)
