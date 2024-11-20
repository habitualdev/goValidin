package goValidin

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

func (v *Validin) DnsInfrastructure(domain, firstSeen, lastSeen string, limit int, wildcard, annotate bool) (DnsGeneric, error) {
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

	if annotate {
		parameters["annotate"] = "tags"
	}

	resp, err := v.DoQuery(fmt.Sprintf(dnsInfra, domain), parameters)
	if err != nil {
		return DnsGeneric{}, err
	}
	var result DnsGeneric
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return DnsGeneric{}, err
	}
	return result, nil
}

func (v *Validin) DnsARecords(domain, firstSeen, lastSeen string, limit int, wildcard, annotate bool) (DnsGeneric, error) {
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
	resp, err := v.DoQuery(fmt.Sprintf(dnsInfraA, domain), parameters)
	if err != nil {
		return DnsGeneric{}, err
	}
	fmt.Println(string(resp))
	var result DnsGeneric
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return DnsGeneric{}, err
	}
	return result, nil
}

func (v *Validin) DnsAAAARecords(domain, firstSeen, lastSeen string, limit int, wildcard, annotate bool) (DnsGeneric, error) {
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
	resp, err := v.DoQuery(fmt.Sprintf(dnsInfraAAAA, domain), parameters)
	if err != nil {
		return DnsGeneric{}, err
	}
	var result DnsGeneric
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return DnsGeneric{}, err
	}
	return result, nil
}

func (v *Validin) DnsNSRecords(domain, firstSeen, lastSeen string, limit int, wildcard, annotate bool) (DnsGeneric, error) {
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
	resp, err := v.DoQuery(fmt.Sprintf(dnsInfraNS, domain), parameters)
	if err != nil {
		return DnsGeneric{}, err
	}
	var result DnsGeneric
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return DnsGeneric{}, err
	}
	return result, nil
}

func (v *Validin) DnsNSFORRecords(domain, firstSeen, lastSeen string, limit int, wildcard, annotate bool) (DnsGeneric, error) {
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
	resp, err := v.DoQuery(fmt.Sprintf(dnsInfraNSFOR, domain), parameters)
	if err != nil {
		return DnsGeneric{}, err
	}
	var result DnsGeneric
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return DnsGeneric{}, err
	}
	return result, nil
}

func (v *Validin) DnsPTRRecords(domain, firstSeen, lastSeen string, limit int, wildcard, annotate bool) (DnsGeneric, error) {
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
	resp, err := v.DoQuery(fmt.Sprintf(dnsPtr, domain), parameters)
	if err != nil {
		return DnsGeneric{}, err
	}
	var result DnsGeneric
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return DnsGeneric{}, err
	}
	return result, nil
}

func (v *Validin) DnsExtra(domain, firstSeen, lastSeen string, limit int, wildcard, annotate bool) (DnsGeneric, error) {
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
	resp, err := v.DoQuery(fmt.Sprintf(dnsExtra, domain), parameters)
	if err != nil {
		return DnsGeneric{}, err
	}
	var result DnsGeneric
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return DnsGeneric{}, err
	}
	return result, nil
}

func (v *Validin) DnsSubdomains(domain string, limit int, annotate bool) (DnsGeneric, error) {
	if limit < 1 || limit > 20000 {
		return DnsGeneric{}, errors.New("invalid limit")
	}
	parameters := map[string]string{
		"limit":    strconv.Itoa(limit),
		"annotate": strconv.FormatBool(annotate),
	}
	resp, err := v.DoQuery(fmt.Sprintf(dnsSubDomain, domain), parameters)
	if err != nil {
		return DnsGeneric{}, err
	}
	var result DnsGeneric
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return DnsGeneric{}, err
	}
	return result, nil
}

func (v *Validin) DnsOsintHistory(domain, firstSeen, lastSeen string, limit int, wildcard bool) (DnsGeneric, error) {
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
	resp, err := v.DoQuery(fmt.Sprintf(dnsOsintHistory, domain), parameters)
	if err != nil {
		return DnsGeneric{}, err
	}
	var result DnsGeneric
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return DnsGeneric{}, err
	}
	return result, nil
}

func (v *Validin) DnsOsintContext(domain string, wildcard bool) ([]byte, error) {
	parameters := map[string]string{
		"annotate": strconv.FormatBool(wildcard),
	}
	return v.DoQuery(fmt.Sprintf(dnsOsintContext, domain), parameters)
}

func (v *Validin) DnsFastRep(domain string) ([]byte, error) {
	parameters := map[string]string{}
	return v.DoQuery(fmt.Sprintf(dnsFastRep, domain), parameters)
}

func (v *Validin) DnsHostResponsePivots(domain, firstSeen, lastSeen string, limit int, wildcard, annotate bool) ([]byte, error) {
	_, err := time.Parse("2006-01-02", firstSeen)
	if err != nil {
		return nil, errors.New("invalid firstSeen date")
	}
	_, err = time.Parse("2006-01-02", lastSeen)
	if err != nil {
		return nil, errors.New("invalid lastSeen date")
	}
	if limit < 1 || limit > 250 {
		return nil, errors.New("invalid limit")
	}
	parameters := map[string]string{
		"first_seen": firstSeen,
		"last_seen":  lastSeen,
		"limit":      strconv.Itoa(limit),
		"wildcard":   strconv.FormatBool(wildcard),
		"annotate":   strconv.FormatBool(annotate),
	}
	return v.DoQuery(fmt.Sprintf(HostResponsePivots, domain), parameters)
}

func (v *Validin) DnsCtStreamLogs(domain, firstSeen, lastSeen string, limit int, wildcard, annotate bool) ([]byte, error) {
	_, err := time.Parse("2006-01-02", firstSeen)
	if err != nil {
		return nil, errors.New("invalid firstSeen date")
	}
	_, err = time.Parse("2006-01-02", lastSeen)
	if err != nil {
		return nil, errors.New("invalid lastSeen date")
	}
	if limit < 1 || limit > 250 {
		return nil, errors.New("invalid limit")
	}
	parameters := map[string]string{
		"first_seen": firstSeen,
		"last_seen":  lastSeen,
		"limit":      strconv.Itoa(limit),
		"wildcard":   strconv.FormatBool(wildcard),
		"annotate":   strconv.FormatBool(annotate),
	}
	return v.DoQuery(fmt.Sprintf(dnsCtStreamLogs, domain), parameters)
}

func (v *Validin) DnsRegHistory(domain, firstSeen, lastSeen string, limit int) (RegistrationHistory, error) {
	_, err := time.Parse("2006-01-02", firstSeen)
	if err != nil {
		return RegistrationHistory{}, errors.New("invalid firstSeen date")
	}
	_, err = time.Parse("2006-01-02", lastSeen)
	if err != nil {
		return RegistrationHistory{}, errors.New("invalid lastSeen date")
	}
	if limit < 1 || limit > 250 {
		return RegistrationHistory{}, errors.New("invalid limit")
	}
	parameters := map[string]string{
		"first_seen": firstSeen,
		"last_seen":  lastSeen,
		"limit":      strconv.Itoa(limit),
	}
	resp, err := v.DoQuery(fmt.Sprintf(dnsRegStreamLogs, domain), parameters)
	if err != nil {
		return RegistrationHistory{}, err
	}
	var result RegistrationHistory
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return RegistrationHistory{}, err
	}
	return result, nil
}

type RegistrationHistory struct {
	QueryOpts struct {
		Limit             int      `json:"limit"`
		Wildcard          bool     `json:"wildcard"`
		Type              string   `json:"type"`
		CategoriesInclude []string `json:"categories_include"`
		Annotate          bool     `json:"annotate"`
	} `json:"query_opts"`
	QueryKey string `json:"query_key"`
	Status   string `json:"status"`
	Records  struct {
		Registration []struct {
			Found  bool     `json:"found"`
			Date   string   `json:"date"`
			Status []string `json:"status"`
			SDNS   bool     `json:"sDNS"`
			Roles  struct {
				ABUSE struct {
					Tel   string `json:"tel"`
					Email string `json:"email"`
				} `json:"ABUSE,omitempty"`
				Abuse struct {
					Tel   []string `json:"tel"`
					Email []string `json:"email"`
				} `json:"abuse,omitempty"`
			} `json:"roles"`
			Key         string      `json:"key"`
			KeyType     string      `json:"key_type"`
			Changed     []time.Time `json:"changed"`
			Updated     []time.Time `json:"updated"`
			Registrar   []string    `json:"registrar"`
			Registered  []time.Time `json:"registered"`
			Expires     []time.Time `json:"expires"`
			Nameservers []string    `json:"nameservers"`
			Source      string      `json:"source"`
			Related     []string    `json:"related"`
			Domain      string      `json:"domain"`
		} `json:"registration"`
	} `json:"records"`
	RecordsReturned int  `json:"records_returned"`
	Limited         bool `json:"limited"`
}

func (v *Validin) DnsLiveReg(domain string, normalize bool) ([]byte, error) {
	parameters := map[string]string{
		"normalize": strconv.FormatBool(normalize),
	}
	return v.DoQuery(fmt.Sprintf(dnsLiveReg, domain), parameters)
}

type LiveRegResp struct {
	QueryOpts struct {
		Limit    int  `json:"limit"`
		Wildcard bool `json:"wildcard"`
		Parent   bool `json:"parent"`
	} `json:"query_opts"`
	QueryKey string `json:"query_key"`
	Records  []struct {
		Status     string `json:"status"`
		CreateTime int    `json:"create_time"`
		Result     struct {
			Domain   string `json:"domain"`
			Host     string `json:"host"`
			Ip       string `json:"ip"`
			Response string `json:"response"`
			Found    bool   `json:"found"`
			Date     string `json:"date"`
			Rdap     struct {
				ObjectClassName string `json:"objectClassName"`
				Handle          string `json:"handle"`
				LdhName         string `json:"ldhName"`
				Links           []struct {
					Value string `json:"value"`
					Rel   string `json:"rel"`
					Href  string `json:"href"`
					Type  string `json:"type"`
				} `json:"links"`
				Status   []string `json:"status"`
				Entities []struct {
					ObjectClassName string   `json:"objectClassName"`
					Handle          string   `json:"handle"`
					Roles           []string `json:"roles"`
					PublicIds       []struct {
						Type       string `json:"type"`
						Identifier string `json:"identifier"`
					} `json:"publicIds"`
					VcardArray []interface{} `json:"vcardArray"`
					Entities   []struct {
						ObjectClassName string        `json:"objectClassName"`
						Roles           []string      `json:"roles"`
						VcardArray      []interface{} `json:"vcardArray"`
					} `json:"entities"`
				} `json:"entities"`
				Events []struct {
					EventAction string    `json:"eventAction"`
					EventDate   time.Time `json:"eventDate"`
				} `json:"events"`
				SecureDNS struct {
					DelegationSigned bool `json:"delegationSigned"`
				} `json:"secureDNS"`
				Nameservers []struct {
					ObjectClassName string `json:"objectClassName"`
					LdhName         string `json:"ldhName"`
				} `json:"nameservers"`
				RdapConformance []string `json:"rdapConformance"`
				Notices         []struct {
					Title       string   `json:"title"`
					Description []string `json:"description"`
					Links       []struct {
						Href string `json:"href"`
						Type string `json:"type"`
					} `json:"links"`
				} `json:"notices"`
			} `json:"rdap"`
		} `json:"result"`
	} `json:"records"`
}
