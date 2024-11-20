package goValidin

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

func (v *Validin) IpDnsHistory(ip, firstSeen, lastSeen string, limit int, annotate bool) (DnsGeneric, error) {
	_, err := time.Parse("2006-01-02", firstSeen)
	if err != nil && firstSeen != "" {
		return DnsGeneric{}, errors.New("invalid firstSeen date")
	}
	_, err = time.Parse("2006-01-02", lastSeen)
	if err != nil && lastSeen != "" {
		return DnsGeneric{}, errors.New("invalid lastSeen date")
	}
	if limit < 1 || limit > 20000 {
		return DnsGeneric{}, errors.New("invalid limit")
	}
	parameters := map[string]string{
		"first_seen": firstSeen,
		"last_seen":  lastSeen,
		"limit":      strconv.Itoa(limit),
		"annotate":   strconv.FormatBool(annotate),
	}
	resp, err := v.DoQuery(fmt.Sprintf(ipDnsHistory, ip), parameters)

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

func (v *Validin) IpDnsHistoryCidr(ip, firstSeen, lastSeen string, cidr, limit int, annotate bool) (DnsGeneric, error) {
	_, err := time.Parse("2006-01-02", firstSeen)
	if err != nil && firstSeen != "" {
		return DnsGeneric{}, errors.New("invalid firstSeen date")
	}
	_, err = time.Parse("2006-01-02", lastSeen)
	if err != nil && lastSeen != "" {
		return DnsGeneric{}, errors.New("invalid lastSeen date")
	}
	if limit < 1 || limit > 20000 {
		return DnsGeneric{}, errors.New("invalid limit")
	}
	parameters := map[string]string{
		"first_seen": firstSeen,
		"last_seen":  lastSeen,
		"limit":      strconv.Itoa(limit),
		"annotate":   strconv.FormatBool(annotate),
	}
	resp, err := v.DoQuery(fmt.Sprintf(ipDnsHistoryCidr, ip, strconv.Itoa(cidr)), parameters)

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

func (v *Validin) IpPtrRecords(ip, firstSeen, lastSeen string, limit int, annotate bool) (DnsGeneric, error) {
	_, err := time.Parse("2006-01-02", firstSeen)
	if err != nil && firstSeen != "" {
		return DnsGeneric{}, errors.New("invalid firstSeen date")
	}
	_, err = time.Parse("2006-01-02", lastSeen)
	if err != nil && lastSeen != "" {
		return DnsGeneric{}, errors.New("invalid lastSeen date")
	}
	if limit < 1 || limit > 20000 {
		return DnsGeneric{}, errors.New("invalid limit")
	}
	parameters := map[string]string{
		"first_seen": firstSeen,
		"last_seen":  lastSeen,
		"limit":      strconv.Itoa(limit),
		"annotate":   strconv.FormatBool(annotate),
	}
	resp, err := v.DoQuery(fmt.Sprintf(ipPtrRecords, ip), parameters)

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

func (v *Validin) IpPtrRecordsCidr(ip, firstSeen, lastSeen string, cidr, limit int, annotate bool) (DnsGeneric, error) {
	_, err := time.Parse("2006-01-02", firstSeen)
	if err != nil && firstSeen != "" {
		return DnsGeneric{}, errors.New("invalid firstSeen date")
	}
	_, err = time.Parse("2006-01-02", lastSeen)
	if err != nil && lastSeen != "" {
		return DnsGeneric{}, errors.New("invalid lastSeen date")
	}
	if limit < 1 || limit > 20000 {
		return DnsGeneric{}, errors.New("invalid limit")
	}
	parameters := map[string]string{
		"first_seen": firstSeen,
		"last_seen":  lastSeen,
		"limit":      strconv.Itoa(limit),
		"annotate":   strconv.FormatBool(annotate),
	}
	resp, err := v.DoQuery(fmt.Sprintf(ipPtrRecordsCidr, ip, strconv.Itoa(cidr)), parameters)
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

func (v *Validin) IpExtraRecords(ip, firstSeen, lastSeen string, limit int, annotate bool) (DnsGeneric, error) {
	_, err := time.Parse("2006-01-02", firstSeen)
	if err != nil && firstSeen != "" {
		return DnsGeneric{}, errors.New("invalid firstSeen date")
	}
	_, err = time.Parse("2006-01-02", lastSeen)
	if err != nil && lastSeen != "" {
		return DnsGeneric{}, errors.New("invalid lastSeen date")
	}
	if limit < 1 || limit > 20000 {
		return DnsGeneric{}, errors.New("invalid limit")
	}
	parameters := map[string]string{
		"first_seen": firstSeen,
		"last_seen":  lastSeen,
		"limit":      strconv.Itoa(limit),
		"annotate":   strconv.FormatBool(annotate),
	}
	resp, err := v.DoQuery(fmt.Sprintf(ipExtra, ip), parameters)
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

func (v *Validin) IpExtraRecordsCidr(ip, firstSeen, lastSeen string, cidr, limit int, annotate bool) (DnsGeneric, error) {
	_, err := time.Parse("2006-01-02", firstSeen)
	if err != nil && firstSeen != "" {
		return DnsGeneric{}, errors.New("invalid firstSeen date")
	}
	_, err = time.Parse("2006-01-02", lastSeen)
	if err != nil && lastSeen != "" {
		return DnsGeneric{}, errors.New("invalid lastSeen date")
	}
	if limit < 1 || limit > 20000 {
		return DnsGeneric{}, errors.New("invalid limit")
	}
	parameters := map[string]string{
		"first_seen": firstSeen,
		"last_seen":  lastSeen,
		"limit":      strconv.Itoa(limit),
		"annotate":   strconv.FormatBool(annotate),
	}
	resp, err := v.DoQuery(fmt.Sprintf(ipExtraCidr, ip, strconv.Itoa(cidr)), parameters)

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

func (v *Validin) IpOsintHistory(ip, firstSeen, lastSeen string, limit int) (DnsGeneric, error) {
	_, err := time.Parse("2006-01-02", firstSeen)
	if err != nil && firstSeen != "" {
		return DnsGeneric{}, errors.New("invalid firstSeen date")
	}
	_, err = time.Parse("2006-01-02", lastSeen)
	if err != nil && lastSeen != "" {
		return DnsGeneric{}, errors.New("invalid lastSeen date")
	}
	if limit < 1 || limit > 20000 {
		return DnsGeneric{}, errors.New("invalid limit")
	}
	parameters := map[string]string{
		"first_seen": firstSeen,
		"last_seen":  lastSeen,
		"limit":      strconv.Itoa(limit),
	}
	resp, err := v.DoQuery(fmt.Sprintf(ipOsintHistory, ip), parameters)
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

func (v *Validin) IpOsintHistoryCidr(ip, firstSeen, lastSeen string, cidr, limit int) (DnsGeneric, error) {
	_, err := time.Parse("2006-01-02", firstSeen)
	if err != nil && firstSeen != "" {
		return DnsGeneric{}, errors.New("invalid firstSeen date")
	}
	_, err = time.Parse("2006-01-02", lastSeen)
	if err != nil && lastSeen != "" {
		return DnsGeneric{}, errors.New("invalid lastSeen date")
	}
	if limit < 1 || limit > 20000 {
		return DnsGeneric{}, errors.New("invalid limit")
	}
	parameters := map[string]string{
		"first_seen": firstSeen,
		"last_seen":  lastSeen,
		"limit":      strconv.Itoa(limit),
	}
	resp, err := v.DoQuery(fmt.Sprintf(ipOsintHistoryCidr, ip, strconv.Itoa(cidr)), parameters)
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

func (v *Validin) IpFastReport(ip string) (DnsGeneric, error) {
	resp, err := v.DoQuery(fmt.Sprintf(ipFastRep, ip), map[string]string{})
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

func (v *Validin) IpOsintContext(ip string) (DnsGeneric, error) {
	resp, err := v.DoQuery(fmt.Sprintf(ipOsintContext, ip), map[string]string{})
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

func (v *Validin) IpHostResponse(ip, firstSeen, lastSeen string, limit int, annotate bool) (DnsGeneric, error) {
	_, err := time.Parse("2006-01-02", firstSeen)
	if err != nil && firstSeen != "" {
		return DnsGeneric{}, errors.New("invalid firstSeen date")
	}
	_, err = time.Parse("2006-01-02", lastSeen)
	if err != nil && lastSeen != "" {
		return DnsGeneric{}, errors.New("invalid lastSeen date")
	}
	if limit < 1 || limit > 20000 {
		return DnsGeneric{}, errors.New("invalid limit")
	}
	parameters := map[string]string{
		"first_seen": firstSeen,
		"last_seen":  lastSeen,
		"limit":      strconv.Itoa(limit),
		"annotate":   strconv.FormatBool(annotate),
	}
	resp, err := v.DoQuery(fmt.Sprintf(ipHostResponsePivots, ip), parameters)

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

func (v *Validin) IpHostResponseCidr(ip, firstSeen, lastSeen string, cidr, limit int, annotate bool) (DnsGeneric, error) {
	_, err := time.Parse("2006-01-02", firstSeen)
	if err != nil && firstSeen != "" {
		return DnsGeneric{}, errors.New("invalid firstSeen date")
	}
	_, err = time.Parse("2006-01-02", lastSeen)
	if err != nil && lastSeen != "" {
		return DnsGeneric{}, errors.New("invalid lastSeen date")
	}
	if limit < 1 || limit > 20000 {
		return DnsGeneric{}, errors.New("invalid limit")
	}
	parameters := map[string]string{
		"first_seen": firstSeen,
		"last_seen":  lastSeen,
		"limit":      strconv.Itoa(limit),
		"annotate":   strconv.FormatBool(annotate),
	}
	resp, err := v.DoQuery(fmt.Sprintf(ipHostResponsePivotsCidr, ip, strconv.Itoa(cidr)), parameters)

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
