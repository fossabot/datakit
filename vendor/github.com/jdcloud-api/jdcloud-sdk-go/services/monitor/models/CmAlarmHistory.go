// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package models


type CmAlarmHistory struct {

    /* 统计方法：平均值=avg、最大值=max、最小值=min (Optional) */
    Calculation string `json:"calculation"`

    /*  (Optional) */
    ContactGroups []string `json:"contactGroups"`

    /*  (Optional) */
    ContactPersons []string `json:"contactPersons"`

    /* 该规则是否已经被删除，1表示已经被删除，0表示未删除，被删除的规则，在使用查询规则的接口时，将不会被检索到 (Optional) */
    Deleted int64 `json:"deleted"`

    /* 启用禁用 1启用，0禁用 (Optional) */
    Enabled int64 `json:"enabled"`

    /* 规则id (Optional) */
    Id string `json:"id"`

    /* 监控项 (Optional) */
    Metric string `json:"metric"`

    /* 规则id监控项名称 (Optional) */
    MetricName string `json:"metricName"`

    /* 命名空间 (Optional) */
    Namespace string `json:"namespace"`

    /* 命名空间id (Optional) */
    NamespaceUID string `json:"namespaceUID"`

    /* 通知周期 单位：小时 (Optional) */
    NoticePeriod int64 `json:"noticePeriod"`

    /*  (Optional) */
    NoticeTime string `json:"noticeTime"`

    /* 对象 (Optional) */
    Obj string `json:"obj"`

    /* 对象id (Optional) */
    ObjUID string `json:"objUID"`

    /* >=、>、<、<=、=、！= (Optional) */
    Operation string `json:"operation"`

    /* 统计周期（单位：分钟） (Optional) */
    Period int64 `json:"period"`

    /* 地域信息 (Optional) */
    Region string `json:"region"`

    /* 此规则所应用的资源id (Optional) */
    ResourceId string `json:"resourceId"`

    /* root rule id (Optional) */
    RootRuleId int64 `json:"rootRuleId"`

    /* rule id (Optional) */
    RuleId int64 `json:"ruleId"`

    /* 规则名称 (Optional) */
    RuleName string `json:"ruleName"`

    /* 报警规则对应的产品 (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* 监控项附属信息 (Optional) */
    Tag string `json:"tag"`

    /* 阈值 (Optional) */
    Threshold float64 `json:"threshold"`

    /* 连续多少次后报警 (Optional) */
    Times int64 `json:"times"`

    /* 报警值 (Optional) */
    Value float64 `json:"value"`
}
