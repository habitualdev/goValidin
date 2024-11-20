package goValidin

import "encoding/json"

func (v *Validin) ApiUsage() (ApiUsageResp, error) {
	resp, err := v.DoQuery(utilApiUsage, nil)
	if err != nil {
		return ApiUsageResp{}, err
	}
	var apiUsage ApiUsageResp

	err = json.Unmarshal(resp, &apiUsage)
	if err != nil {
		return ApiUsageResp{}, err
	}
	return apiUsage, nil
}

type ApiUsageResp struct {
	Usage struct {
		Total struct {
			Daily   int `json:"daily"`
			Monthly int `json:"monthly"`
		} `json:"total"`
	} `json:"usage"`
	Remaining struct {
		Daily   int `json:"daily"`
		Monthly int `json:"monthly"`
	} `json:"remaining"`
}

func (v *Validin) AllPaths() (AllPathsResp, error) {
	resp, err := v.DoQuery(utilAllPaths, nil)
	if err != nil {
		return AllPathsResp{}, err
	}
	var allPaths AllPathsResp

	err = json.Unmarshal(resp, &allPaths)
	if err != nil {
		return AllPathsResp{}, err
	}
	return allPaths, nil
}

type AllPathsResp struct {
	ApiAxonLiveScanStart struct {
		Path        string      `json:"path"`
		Verb        string      `json:"verb"`
		Title       interface{} `json:"title"`
		Description string      `json:"description"`
		QueryParams struct {
			Wildcard struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     bool   `json:"default"`
			} `json:"wildcard"`
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			FirstSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"first_seen"`
			LastSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"last_seen"`
			Annotate struct {
				Type           string   `json:"type"`
				AcceptedValues []string `json:"accepted_values"`
				Description    string   `json:"description"`
			} `json:"annotate"`
		} `json:"query_params"`
		PathParams interface{} `json:"path_params"`
		Cache      string      `json:"cache"`
		Roles      []string    `json:"roles"`
		Category   string      `json:"category"`
		Options    struct {
		} `json:"options"`
		Id         string      `json:"id"`
		Type       string      `json:"type"`
		Visible    bool        `json:"visible"`
		BodyParams interface{} `json:"body_params"`
	} `json:"/api/axon/live/scan/start"`
	ApiProfileUsage struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
		} `json:"query_params"`
		PathParams interface{} `json:"path_params"`
		Cache      string      `json:"cache"`
		Roles      interface{} `json:"roles"`
		Category   string      `json:"category"`
		Options    struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/profile/usage"`
	ApiProfileToken struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
		} `json:"query_params"`
		PathParams interface{} `json:"path_params"`
		Cache      string      `json:"cache"`
		Roles      interface{} `json:"roles"`
		Category   string      `json:"category"`
		Options    struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/profile/token"`
	ApiPaths struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
		} `json:"query_params"`
		PathParams interface{} `json:"path_params"`
		Cache      string      `json:"cache"`
		Roles      interface{} `json:"roles"`
		Category   string      `json:"category"`
		Options    struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/paths"`
	ApiPing struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
		} `json:"query_params"`
		PathParams interface{} `json:"path_params"`
		Cache      string      `json:"cache"`
		Roles      interface{} `json:"roles"`
		Category   string      `json:"category"`
		Options    struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/ping"`
	ApiProjects struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
		} `json:"query_params"`
		PathParams interface{} `json:"path_params"`
		Cache      string      `json:"cache"`
		Roles      []string    `json:"roles"`
		Category   string      `json:"category"`
		Options    struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/projects"`
	ApiProjectsProjectId struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
		} `json:"query_params"`
		PathParams struct {
			ProjectId struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
				Required    bool   `json:"required"`
			} `json:"project_id"`
		} `json:"path_params"`
		Cache    string   `json:"cache"`
		Roles    []string `json:"roles"`
		Category string   `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/projects/:project_id"`
	ApiProjectsProjectIdIndicators struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
		} `json:"query_params"`
		PathParams struct {
			ProjectId struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
				Required    bool   `json:"required"`
			} `json:"project_id"`
		} `json:"path_params"`
		Cache    string   `json:"cache"`
		Roles    []string `json:"roles"`
		Category string   `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/projects/:project_id/indicators"`
	ApiProjectsProjectIdIndicatorsAdd struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
		} `json:"query_params"`
		PathParams interface{} `json:"path_params"`
		Cache      string      `json:"cache"`
		Roles      []string    `json:"roles"`
		Category   string      `json:"category"`
		Options    struct {
		} `json:"options"`
		Id         string      `json:"id"`
		Type       string      `json:"type"`
		Visible    bool        `json:"visible"`
		BodyParams interface{} `json:"body_params"`
	} `json:"/api/projects/:project_id/indicators/add"`
	ApiProjectsProjectIdIndicatorsDelete struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
		} `json:"query_params"`
		PathParams interface{} `json:"path_params"`
		Cache      string      `json:"cache"`
		Roles      []string    `json:"roles"`
		Category   string      `json:"category"`
		Options    struct {
		} `json:"options"`
		Id         string      `json:"id"`
		Type       string      `json:"type"`
		Visible    bool        `json:"visible"`
		BodyParams interface{} `json:"body_params"`
	} `json:"/api/projects/:project_id/indicators/delete"`
	ApiProjectsProjectIdAlertsDates struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
		} `json:"query_params"`
		PathParams struct {
			ProjectId struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
				Required    bool   `json:"required"`
			} `json:"project_id"`
		} `json:"path_params"`
		Cache    string   `json:"cache"`
		Roles    []string `json:"roles"`
		Category string   `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/projects/:project_id/alerts/dates"`
	ApiProjectsProjectIdAlertsLatest struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
		} `json:"query_params"`
		PathParams struct {
			ProjectId struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
				Required    bool   `json:"required"`
			} `json:"project_id"`
		} `json:"path_params"`
		Cache    string   `json:"cache"`
		Roles    []string `json:"roles"`
		Category string   `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/projects/:project_id/alerts/latest"`
	ApiProjectsProjectIdAlertsDate struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
		} `json:"query_params"`
		PathParams struct {
			ProjectId struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
				Required    bool   `json:"required"`
			} `json:"project_id"`
			Date struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
				Required    bool   `json:"required"`
			} `json:"date"`
		} `json:"path_params"`
		Cache    string   `json:"cache"`
		Roles    []string `json:"roles"`
		Category string   `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/projects/:project_id/alerts/:date"`
	ApiAxonDomainDnsHistoryDomain struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Wildcard struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     bool   `json:"default"`
			} `json:"wildcard"`
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			FirstSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"first_seen"`
			LastSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"last_seen"`
			Annotate struct {
				Type           string   `json:"type"`
				AcceptedValues []string `json:"accepted_values"`
				Description    string   `json:"description"`
			} `json:"annotate"`
		} `json:"query_params"`
		PathParams struct {
			Domain struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"domain"`
		} `json:"path_params"`
		Cache    string      `json:"cache"`
		Roles    interface{} `json:"roles"`
		Category string      `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/domain/dns/history/:domain"`
	ApiAxonDomainDnsHistoryDomainA struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Wildcard struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     bool   `json:"default"`
			} `json:"wildcard"`
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			FirstSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"first_seen"`
			LastSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"last_seen"`
			Annotate struct {
				Type           string   `json:"type"`
				AcceptedValues []string `json:"accepted_values"`
				Description    string   `json:"description"`
			} `json:"annotate"`
		} `json:"query_params"`
		PathParams struct {
			Domain struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"domain"`
		} `json:"path_params"`
		Cache    string      `json:"cache"`
		Roles    interface{} `json:"roles"`
		Category string      `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/domain/dns/history/:domain/A"`
	ApiAxonDomainDnsHistoryDomainAAAA struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Wildcard struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     bool   `json:"default"`
			} `json:"wildcard"`
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			FirstSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"first_seen"`
			LastSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"last_seen"`
			Annotate struct {
				Type           string   `json:"type"`
				AcceptedValues []string `json:"accepted_values"`
				Description    string   `json:"description"`
			} `json:"annotate"`
		} `json:"query_params"`
		PathParams struct {
			Domain struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"domain"`
		} `json:"path_params"`
		Cache    string      `json:"cache"`
		Roles    interface{} `json:"roles"`
		Category string      `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/domain/dns/history/:domain/AAAA"`
	ApiAxonDomainDnsHistoryDomainNS struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Wildcard struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     bool   `json:"default"`
			} `json:"wildcard"`
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			FirstSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"first_seen"`
			LastSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"last_seen"`
			Annotate struct {
				Type           string   `json:"type"`
				AcceptedValues []string `json:"accepted_values"`
				Description    string   `json:"description"`
			} `json:"annotate"`
		} `json:"query_params"`
		PathParams struct {
			Domain struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"domain"`
		} `json:"path_params"`
		Cache    string      `json:"cache"`
		Roles    interface{} `json:"roles"`
		Category string      `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/domain/dns/history/:domain/NS"`
	ApiAxonDomainDnsHistoryDomainNSFOR struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Wildcard struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     bool   `json:"default"`
			} `json:"wildcard"`
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			FirstSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"first_seen"`
			LastSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"last_seen"`
			Annotate struct {
				Type           string   `json:"type"`
				AcceptedValues []string `json:"accepted_values"`
				Description    string   `json:"description"`
			} `json:"annotate"`
		} `json:"query_params"`
		PathParams struct {
			Domain struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"domain"`
		} `json:"path_params"`
		Cache    string      `json:"cache"`
		Roles    interface{} `json:"roles"`
		Category string      `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/domain/dns/history/:domain/NS_FOR"`
	ApiAxonIpDnsHistoryIp struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			FirstSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"first_seen"`
			LastSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"last_seen"`
			Annotate struct {
				Type           string   `json:"type"`
				AcceptedValues []string `json:"accepted_values"`
				Description    string   `json:"description"`
			} `json:"annotate"`
		} `json:"query_params"`
		PathParams struct {
			Ip struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"ip"`
		} `json:"path_params"`
		Cache    string      `json:"cache"`
		Roles    interface{} `json:"roles"`
		Category string      `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/ip/dns/history/:ip"`
	ApiAxonIpDnsHistoryIpCidr struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			FirstSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"first_seen"`
			LastSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"last_seen"`
			Annotate struct {
				Type           string   `json:"type"`
				AcceptedValues []string `json:"accepted_values"`
				Description    string   `json:"description"`
			} `json:"annotate"`
		} `json:"query_params"`
		PathParams struct {
			Ip struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"ip"`
			Cidr struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"cidr"`
		} `json:"path_params"`
		Cache    string      `json:"cache"`
		Roles    interface{} `json:"roles"`
		Category string      `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/ip/dns/history/:ip/:cidr"`
	ApiAxonDomainDnsHostnameDomain struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Wildcard struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     bool   `json:"default"`
			} `json:"wildcard"`
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			FirstSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"first_seen"`
			LastSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"last_seen"`
			Annotate struct {
				Type           string   `json:"type"`
				AcceptedValues []string `json:"accepted_values"`
				Description    string   `json:"description"`
			} `json:"annotate"`
		} `json:"query_params"`
		PathParams struct {
			Domain struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"domain"`
		} `json:"path_params"`
		Cache    string      `json:"cache"`
		Roles    interface{} `json:"roles"`
		Category string      `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/domain/dns/hostname/:domain"`
	ApiAxonIpDnsHostnameIp struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			FirstSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"first_seen"`
			LastSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"last_seen"`
			Annotate struct {
				Type           string   `json:"type"`
				AcceptedValues []string `json:"accepted_values"`
				Description    string   `json:"description"`
			} `json:"annotate"`
		} `json:"query_params"`
		PathParams struct {
			Ip struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"ip"`
		} `json:"path_params"`
		Cache    string      `json:"cache"`
		Roles    interface{} `json:"roles"`
		Category string      `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/ip/dns/hostname/:ip"`
	ApiAxonIpDnsHostnameIpCidr struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			FirstSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"first_seen"`
			LastSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"last_seen"`
			Annotate struct {
				Type           string   `json:"type"`
				AcceptedValues []string `json:"accepted_values"`
				Description    string   `json:"description"`
			} `json:"annotate"`
		} `json:"query_params"`
		PathParams struct {
			Ip struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"ip"`
			Cidr struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"cidr"`
		} `json:"path_params"`
		Cache    string      `json:"cache"`
		Roles    interface{} `json:"roles"`
		Category string      `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/ip/dns/hostname/:ip/:cidr"`
	ApiAxonDomainDnsExtraDomain struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Wildcard struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     bool   `json:"default"`
			} `json:"wildcard"`
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			FirstSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"first_seen"`
			LastSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"last_seen"`
			Annotate struct {
				Type           string   `json:"type"`
				AcceptedValues []string `json:"accepted_values"`
				Description    string   `json:"description"`
			} `json:"annotate"`
		} `json:"query_params"`
		PathParams struct {
			Domain struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"domain"`
		} `json:"path_params"`
		Cache    string      `json:"cache"`
		Roles    interface{} `json:"roles"`
		Category string      `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/domain/dns/extra/:domain"`
	ApiAxonIpDnsExtraIp struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			FirstSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"first_seen"`
			LastSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"last_seen"`
			Annotate struct {
				Type           string   `json:"type"`
				AcceptedValues []string `json:"accepted_values"`
				Description    string   `json:"description"`
			} `json:"annotate"`
		} `json:"query_params"`
		PathParams struct {
			Ip struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"ip"`
		} `json:"path_params"`
		Cache    string      `json:"cache"`
		Roles    interface{} `json:"roles"`
		Category string      `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/ip/dns/extra/:ip"`
	ApiAxonIpDnsExtraIpCidr struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			FirstSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"first_seen"`
			LastSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"last_seen"`
			Annotate struct {
				Type           string   `json:"type"`
				AcceptedValues []string `json:"accepted_values"`
				Description    string   `json:"description"`
			} `json:"annotate"`
		} `json:"query_params"`
		PathParams struct {
			Ip struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"ip"`
			Cidr struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"cidr"`
		} `json:"path_params"`
		Cache    string      `json:"cache"`
		Roles    interface{} `json:"roles"`
		Category string      `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/ip/dns/extra/:ip/:cidr"`
	ApiAxonStringDnsExtra struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Wildcard struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     bool   `json:"default"`
			} `json:"wildcard"`
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			FirstSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"first_seen"`
			LastSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"last_seen"`
			Annotate struct {
				Type           string   `json:"type"`
				AcceptedValues []string `json:"accepted_values"`
				Description    string   `json:"description"`
			} `json:"annotate"`
		} `json:"query_params"`
		PathParams struct {
			String struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"string"`
		} `json:"path_params"`
		Cache    string      `json:"cache"`
		Roles    interface{} `json:"roles"`
		Category string      `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/string/dns/extra/*"`
	ApiAxonDomainSubdomainsDomain struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			Annotate struct {
				Type           string   `json:"type"`
				AcceptedValues []string `json:"accepted_values"`
				Description    string   `json:"description"`
			} `json:"annotate"`
		} `json:"query_params"`
		PathParams struct {
			Domain struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"domain"`
		} `json:"path_params"`
		Cache    string      `json:"cache"`
		Roles    interface{} `json:"roles"`
		Category string      `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/domain/subdomains/:domain"`
	ApiAxonDomainOsintHistoryDomain struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Wildcard struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     bool   `json:"default"`
			} `json:"wildcard"`
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			FirstSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"first_seen"`
			LastSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"last_seen"`
		} `json:"query_params"`
		PathParams struct {
			Domain struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"domain"`
		} `json:"path_params"`
		Cache    string      `json:"cache"`
		Roles    interface{} `json:"roles"`
		Category string      `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/domain/osint/history/:domain"`
	ApiAxonIpOsintHistoryIp struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			FirstSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"first_seen"`
			LastSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"last_seen"`
		} `json:"query_params"`
		PathParams struct {
			Ip struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"ip"`
		} `json:"path_params"`
		Cache    string      `json:"cache"`
		Roles    interface{} `json:"roles"`
		Category string      `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/ip/osint/history/:ip"`
	ApiAxonIpOsintHistoryIpCidr struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			FirstSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"first_seen"`
			LastSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"last_seen"`
		} `json:"query_params"`
		PathParams struct {
			Ip struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"ip"`
			Cidr struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"cidr"`
		} `json:"path_params"`
		Cache    string      `json:"cache"`
		Roles    interface{} `json:"roles"`
		Category string      `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/ip/osint/history/:ip/:cidr"`
	ApiAxonDomainOsintContextDomain struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Wildcard struct {
				Type        string `json:"type"`
				Description string `json:"description"`
			} `json:"wildcard"`
		} `json:"query_params"`
		PathParams struct {
			Domain struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"domain"`
		} `json:"path_params"`
		Cache    string      `json:"cache"`
		Roles    interface{} `json:"roles"`
		Category string      `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/domain/osint/context/:domain"`
	ApiAxonBulkOsintContext struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
		} `json:"query_params"`
		PathParams interface{} `json:"path_params"`
		Cache      string      `json:"cache"`
		Roles      []string    `json:"roles"`
		Category   string      `json:"category"`
		Options    struct {
		} `json:"options"`
		Id         string `json:"id"`
		Type       string `json:"type"`
		Visible    bool   `json:"visible"`
		BodyParams struct {
			Domains struct {
				Type        string `json:"type"`
				Subtype     string `json:"subtype"`
				Description string `json:"description"`
			} `json:"domains"`
			Ipv4Addresses struct {
				Type        string `json:"type"`
				Subtype     string `json:"subtype"`
				Description string `json:"description"`
			} `json:"ipv4_addresses"`
			Ipv6Addresses struct {
				Type        string `json:"type"`
				Subtype     string `json:"subtype"`
				Description string `json:"description"`
			} `json:"ipv6_addresses"`
		} `json:"body_params"`
	} `json:"/api/axon/bulk/osint/context"`
	ApiAxonDomainReputationQuickDomain struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
		} `json:"query_params"`
		PathParams struct {
			Domain struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"domain"`
		} `json:"path_params"`
		Cache    string   `json:"cache"`
		Roles    []string `json:"roles"`
		Category string   `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/domain/reputation/quick/:domain"`
	ApiAxonIpReputationQuickIp struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
		} `json:"query_params"`
		PathParams struct {
			Ip struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"ip"`
		} `json:"path_params"`
		Cache    string   `json:"cache"`
		Roles    []string `json:"roles"`
		Category string   `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/ip/reputation/quick/:ip"`
	ApiAxonIpOsintContextIp struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
		} `json:"query_params"`
		PathParams struct {
			Ip struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"ip"`
		} `json:"path_params"`
		Cache    string      `json:"cache"`
		Roles    interface{} `json:"roles"`
		Category string      `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/ip/osint/context/:ip"`
	ApiAxonDomainPivotsDomain struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Wildcard struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     bool   `json:"default"`
			} `json:"wildcard"`
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			FirstSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"first_seen"`
			LastSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"last_seen"`
			Annotate struct {
				Type           string   `json:"type"`
				AcceptedValues []string `json:"accepted_values"`
				Description    string   `json:"description"`
			} `json:"annotate"`
		} `json:"query_params"`
		PathParams struct {
			Domain struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"domain"`
		} `json:"path_params"`
		Cache    string      `json:"cache"`
		Roles    interface{} `json:"roles"`
		Category string      `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/domain/pivots/:domain"`
	ApiAxonIpPivotsIp struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			FirstSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"first_seen"`
			LastSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"last_seen"`
			Annotate struct {
				Type           string   `json:"type"`
				AcceptedValues []string `json:"accepted_values"`
				Description    string   `json:"description"`
			} `json:"annotate"`
		} `json:"query_params"`
		PathParams struct {
			Ip struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"ip"`
		} `json:"path_params"`
		Cache    string      `json:"cache"`
		Roles    interface{} `json:"roles"`
		Category string      `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/ip/pivots/:ip"`
	ApiAxonIpPivotsIpCidr struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			FirstSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"first_seen"`
			LastSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"last_seen"`
			Annotate struct {
				Type           string   `json:"type"`
				AcceptedValues []string `json:"accepted_values"`
				Description    string   `json:"description"`
			} `json:"annotate"`
		} `json:"query_params"`
		PathParams struct {
			Ip struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"ip"`
			Cidr struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"cidr"`
		} `json:"path_params"`
		Cache    string      `json:"cache"`
		Roles    interface{} `json:"roles"`
		Category string      `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/ip/pivots/:ip/:cidr"`
	ApiAxonHashPivotsHash struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Wildcard struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     bool   `json:"default"`
			} `json:"wildcard"`
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			FirstSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"first_seen"`
			LastSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"last_seen"`
			Annotate struct {
				Type           string   `json:"type"`
				AcceptedValues []string `json:"accepted_values"`
				Description    string   `json:"description"`
			} `json:"annotate"`
		} `json:"query_params"`
		PathParams struct {
			Hash struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
				Required    bool   `json:"required"`
			} `json:"hash"`
		} `json:"path_params"`
		Cache    string      `json:"cache"`
		Roles    interface{} `json:"roles"`
		Category string      `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/hash/pivots/:hash"`
	ApiAxonStringPivots struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Wildcard struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     bool   `json:"default"`
			} `json:"wildcard"`
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			FirstSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"first_seen"`
			LastSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"last_seen"`
			Annotate struct {
				Type           string   `json:"type"`
				AcceptedValues []string `json:"accepted_values"`
				Description    string   `json:"description"`
			} `json:"annotate"`
		} `json:"query_params"`
		PathParams struct {
			String struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"string"`
		} `json:"path_params"`
		Cache    string      `json:"cache"`
		Roles    interface{} `json:"roles"`
		Category string      `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/string/pivots/*"`
	ApiAxonDomainCertificatesDomain struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Wildcard struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     bool   `json:"default"`
			} `json:"wildcard"`
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			FirstSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"first_seen"`
			LastSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"last_seen"`
			Annotate struct {
				Type           string   `json:"type"`
				AcceptedValues []string `json:"accepted_values"`
				Description    string   `json:"description"`
			} `json:"annotate"`
		} `json:"query_params"`
		PathParams struct {
			Domain struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"domain"`
		} `json:"path_params"`
		Cache    string      `json:"cache"`
		Roles    interface{} `json:"roles"`
		Category string      `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/domain/certificates/:domain"`
	ApiAxonDomainRegistrationHistoryDomain struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			FirstSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"first_seen"`
			LastSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"last_seen"`
		} `json:"query_params"`
		PathParams struct {
			Domain struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"domain"`
		} `json:"path_params"`
		Cache    string   `json:"cache"`
		Roles    []string `json:"roles"`
		Category string   `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/domain/registration/history/:domain"`
	ApiAxonDomainRegistrationLiveDomain struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Normalize struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     bool   `json:"default"`
			} `json:"normalize"`
		} `json:"query_params"`
		PathParams struct {
			Domain struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"domain"`
		} `json:"path_params"`
		Cache    string   `json:"cache"`
		Roles    []string `json:"roles"`
		Category string   `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/domain/registration/live/:domain"`
	ApiAxonStringRegistrationHistory struct {
		Path        string `json:"path"`
		Verb        string `json:"verb"`
		Title       string `json:"title"`
		Description string `json:"description"`
		QueryParams struct {
			Wildcard struct {
				Type        string `json:"type"`
				Description string `json:"description"`
			} `json:"wildcard"`
			Limit struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     int    `json:"default"`
				Min         int    `json:"min"`
				Max         int    `json:"max"`
			} `json:"limit"`
			FirstSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"first_seen"`
			LastSeen struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Format      string `json:"format"`
			} `json:"last_seen"`
		} `json:"query_params"`
		PathParams struct {
			String struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
			} `json:"string"`
		} `json:"path_params"`
		Cache    string   `json:"cache"`
		Roles    []string `json:"roles"`
		Category string   `json:"category"`
		Options  struct {
		} `json:"options"`
		Id      string `json:"id"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"/api/axon/string/registration/history/*"`
}

func (v *Validin) Ping() (PingResp, error) {
	resp, err := v.DoQuery(utilPing, nil)
	if err != nil {
		return PingResp{}, err
	}
	pingResp := PingResp{}
	err = json.Unmarshal(resp, &pingResp)
	if err != nil {
		return PingResp{}, err
	}
	return pingResp, nil

}

type PingResp struct {
	Status string `json:"status"`
}
