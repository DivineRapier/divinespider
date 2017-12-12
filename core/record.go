/*
 * @Author: divinerapier
 * @Date: 2017-11-18 22:38:18
 * @Last Modified by: divinerapier
 * @Last Modified time: 2017-11-18 23:32:57
 */

package core

type (
	data struct {
		URL     string `json:"url"`
		Data    string `json:"data"`
		URLrule string `json:"url_rule"`
		RuleCnt int32  `json:"rule_cnt"`
	}

	// Record 代表爬虫获取到的一条记录
	record struct {
		Type   string `json:"type"`
		Method string `json:"method"`
		Data   data   `json:"data"`
	}

	// Records 结果集合
	records []record
)

// NewRecord 创建一条记录
func NewRecord(typ, mtd, u, d string) record {
	return record{
		Type:   typ,
		Method: mtd,
		Data: data{
			URL:     u,
			Data:    d,
			URLrule: "",
			RuleCnt: 0,
		},
	}
}

// Len 长度
func (r records) Len() int {
	return len(r)
}

// Less 比较
func (r records) Less(i, j int) bool {
	return (r[i].Data.URL < r[j].Data.URL) ||
		(r[i].Data.URL == r[j].Data.URL && r[i].Method < r[j].Method) ||
		(r[i].Data.URL == r[j].Data.URL && r[i].Method == r[j].Method && r[i].Data.Data < r[j].Data.Data)
}

// Swap 交换
func (r records) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
