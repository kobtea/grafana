package dtos

import (
	"time"

	"github.com/grafana/grafana/pkg/components/simplejson"
	m "github.com/grafana/grafana/pkg/models"
)

type AlertRule struct {
	Id             int64               `json:"id"`
	DashboardId    int64               `json:"dashboardId"`
	PanelId        int64               `json:"panelId"`
	Name           string              `json:"name"`
	Message        string              `json:"message"`
	State          m.AlertStateType    `json:"state"`
	Severity       m.AlertSeverityType `json:"severity"`
	NewStateDate   time.Time           `json:"newStateDate"`
	EvalDate       time.Time           `json:"evalDate"`
	ExecutionError string              `json:"executionError"`
	DashbboardUri  string              `json:"dashboardUri"`
}

type AlertNotification struct {
	Id      int64     `json:"id"`
	Name    string    `json:"name"`
	Type    string    `json:"type"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type AlertTestCommand struct {
	Dashboard *simplejson.Json `json:"dashboard" binding:"Required"`
	PanelId   int64            `json:"panelId" binding:"Required"`
}

type AlertTestResult struct {
	Firing      bool                  `json:"firing"`
	TimeMs      string                `json:"timeMs"`
	Error       string                `json:"error,omitempty"`
	EvalMatches []*EvalMatch          `json:"matches,omitempty"`
	Logs        []*AlertTestResultLog `json:"logs,omitempty"`
}

type AlertTestResultLog struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type EvalMatch struct {
	Tags   map[string]string `json:"tags,omitempty"`
	Metric string            `json:"metric"`
	Value  float64           `json:"value"`
}

type AlertHistory struct {
	AlertId   int64     `json:"alertId"`
	NewState  string    `json:"newState"`
	Timestamp time.Time `json:"timestamp"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	Metric    string    `json:"metric"`

	Data *simplejson.Json `json:"data"`
}
