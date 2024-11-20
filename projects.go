package goValidin

import (
	"encoding/json"
	"fmt"
)

type AllProjects struct {
	Status   string   `json:"status"`
	Users    []string `json:"users"`
	Projects map[string]struct {
		Details struct {
			Name        string        `json:"name"`
			Description string        `json:"description"`
			Tags        []interface{} `json:"tags"`
			Status      string        `json:"status"`
			CreateTime  int           `json:"create_time"`
			SharedUsers map[string]struct {
				Edit  string `json:"edit"`
				Alert bool   `json:"alert"`
			} `json:"shared_users"`
			AutoAlerts  bool `json:"auto_alerts"`
			AuthorAlert bool `json:"author_alert"`
		} `json:"details"`
		Author      string `json:"author"`
		Permissions struct {
			Edit   string `json:"edit"`
			Alert  bool   `json:"alert"`
			Shared bool   `json:"shared"`
		} `json:"permissions"`
		IndicatorsCount int `json:"indicators_count"`
	} `json:"projects"`
}

func (v *Validin) AllProjects() (AllProjects, error) {
	resp, err := v.DoQuery(projectAll, nil)
	if err != nil {
		return AllProjects{}, err
	}

	var allProjects AllProjects
	err = json.Unmarshal(resp, &allProjects)
	if err != nil {
		return AllProjects{}, err
	}
	return allProjects, nil
}

type ProjectInfo struct {
	Status    string   `json:"status"`
	Users     []string `json:"users"`
	ProjectId string   `json:"project_id"`
	Project   struct {
		Details struct {
			Name        string        `json:"name"`
			Description string        `json:"description"`
			Tags        []interface{} `json:"tags"`
			Status      string        `json:"status"`
			CreateTime  int           `json:"create_time"`
			SharedUsers map[string]struct {
				Edit  string `json:"edit"`
				Alert bool   `json:"alert"`
			} `json:"shared_users"`
			AutoAlerts  bool `json:"auto_alerts"`
			AuthorAlert bool `json:"author_alert"`
		} `json:"details"`
		Author      string `json:"author"`
		Permissions struct {
			Edit   string `json:"edit"`
			Alert  bool   `json:"alert"`
			Shared bool   `json:"shared"`
		} `json:"permissions"`
		IndicatorsCount int `json:"indicators_count"`
	} `json:"project"`
}

func (v *Validin) ProjectInfo(projectId string) (ProjectInfo, error) {
	resp, err := v.DoQuery(fmt.Sprintf(projectInfo, projectId), nil)
	if err != nil {
		return ProjectInfo{}, err
	}

	var projectInfo ProjectInfo
	err = json.Unmarshal(resp, &projectInfo)
	if err != nil {
		return ProjectInfo{}, err
	}
	return projectInfo, nil

}

type AllIndicators struct {
	Status     string   `json:"status"`
	Users      []string `json:"users"`
	ProjectId  string   `json:"project_id"`
	Indicators map[string]struct {
		RefId   interface{} `json:"ref_id"`
		AddedBy string      `json:"added_by"`
		Time    int         `json:"time"`
		Type    string      `json:"type"`
		Alert   bool        `json:"alert"`
	} `json:"indicators"`
}

func (v *Validin) AllIndicators(projectId string) (AllIndicators, error) {
	resp, err := v.DoQuery(fmt.Sprintf(projectAllIndicators, projectId), nil)
	if err != nil {
		return AllIndicators{}, err
	}

	var allIndicators AllIndicators
	err = json.Unmarshal(resp, &allIndicators)
	if err != nil {
		return AllIndicators{}, err
	}
	return allIndicators, nil
}
