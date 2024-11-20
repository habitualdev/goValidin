package goValidin

import "errors"

type DnsGeneric struct {
	QueryOpts struct {
		Limit             int      `json:"limit"`
		Wildcard          bool     `json:"wildcard"`
		Type              string   `json:"type"`
		Annotate          bool     `json:"annotate"`
		CategoriesInclude []string `json:"categories_include"`
		CategoriesExclude []string `json:"categories_exclude"`
	} `json:"query_opts"`
	QueryKey      string `json:"query_key"`
	Informational struct {
		Structure struct {
			Fqdn   string `json:"fqdn"`
			Domain string `json:"domain"`
			Etld   string `json:"etld"`
		} `json:"structure"`
	} `json:"informational"`
	Status  string `json:"status"`
	Records map[string][]struct {
		Description  string   `json:"description"`
		Key          string   `json:"key"`
		Value        string   `json:"value"`
		ValueType    string   `json:"value_type"`
		Catagory     string   `json:"catagory"`
		RiskCategory string   `json:"risk_cat"`
		Score        float64  `json:"score"`
		Weight       float64  `json:"weight"`
		Title        string   `json:"title"`
		HelpText     string   `json:"help_text"`
		FirstSeen    int      `json:"first_seen"`
		LastSeen     int      `json:"last_seen"`
		KeyTags      []string `json:"key_tags"`
		ValTags      []string `json:"val_tags"`
	} `json:"records"`
	RecordsReturned int  `json:"records_returned"`
	Limited         bool `json:"limited"`
}

func (d *DnsGeneric) ListRecordKeys() []string {
	keys := make([]string, 0)

	for key := range d.Records {
		keys = append(keys, key)
	}
	return keys
}
func (d *DnsGeneric) GetRecordsByKey(key string) ([]struct {
	Description  string   `json:"description"`
	Key          string   `json:"key"`
	Value        string   `json:"value"`
	ValueType    string   `json:"value_type"`
	Catagory     string   `json:"catagory"`
	RiskCategory string   `json:"risk_cat"`
	Score        float64  `json:"score"`
	Weight       float64  `json:"weight"`
	Title        string   `json:"title"`
	HelpText     string   `json:"help_text"`
	FirstSeen    int      `json:"first_seen"`
	LastSeen     int      `json:"last_seen"`
	KeyTags      []string `json:"key_tags"`
	ValTags      []string `json:"val_tags"`
}, error) {
	if records, ok := d.Records[key]; ok {
		return records, nil
	}
	return nil, errors.New("key not found")
}
