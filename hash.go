package goValidin

import (
	"encoding/json"
	"fmt"
)

func (v *Validin) HashHostResponsePivots(hash string) (DnsGeneric, error) {
	parameters := map[string]string{}
	resp, err := v.DoQuery(fmt.Sprintf(hashHostResponsePivots, hash), parameters)
	if err != nil {
		return DnsGeneric{}, err
	}

	var dnsGeneric DnsGeneric
	err = json.Unmarshal(resp, &dnsGeneric)
	if err != nil {
		return DnsGeneric{}, err
	}
	return dnsGeneric, nil
}
