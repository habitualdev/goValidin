package goValidin

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

func (v *Validin) StringExtraDns(query, firstSeen, lastSeen string, limit int, wildcard, annotate bool) (DnsGeneric, error) {
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
		"wildcard":   strconv.FormatBool(wildcard),
		"annotate":   strconv.FormatBool(annotate),
	}
	resp, err := v.DoQuery(fmt.Sprintf(stringExtraDnsRecords, query), parameters)
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

func (v *Validin) StringHostResponsePivots(query, firstSeen, lastSeen string, limit int, wildcard, annotate bool) (DnsGeneric, error) {
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
		"wildcard":   strconv.FormatBool(wildcard),
	}
	resp, err := v.DoQuery(fmt.Sprintf(stringHostResponsePivots, query), parameters)
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

func (v *Validin) StringRegHistory(query, firstSeen, lastSeen string, limit int, wildcard bool) (DnsGeneric, error) {
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
		"wildcard":   strconv.FormatBool(wildcard),
	}
	resp, err := v.DoQuery(fmt.Sprintf(stringRegHistory, query), parameters)
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
