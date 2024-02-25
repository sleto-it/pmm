// Copyright (C) 2023 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/AlekSi/pointer"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/percona/pmm/admin/agentlocal"
	"github.com/percona/pmm/api/inventory/v1/json/client"
	agents "github.com/percona/pmm/api/inventory/v1/json/client/agents_service"
	services "github.com/percona/pmm/api/inventory/v1/json/client/services_service"
	"github.com/percona/pmm/api/inventory/v1/types"
)

var listResultT = ParseTemplate(`
Service type{{"\t"}}Service name{{"\t"}}Address and port{{"\t"}}Service ID
{{ range .Services }}
{{- .HumanReadableServiceType }}{{"\t"}}{{ .ServiceName }}{{"\t"}}{{ .AddressPort }}{{"\t"}}{{ .ServiceID }}
{{ end }}
Agent type{{"\t"}}Status{{"\t"}}Metrics Mode{{"\t"}}Agent ID{{"\t"}}Service ID{{"\t"}}Port
{{ range .Agents }}
{{- .HumanReadableAgentType }}{{"\t"}}{{ .NiceAgentStatus }}{{"\t"}}{{ .MetricsMode }}{{"\t"}}{{ .AgentID }}{{"\t"}}{{ .ServiceID }}{{"\t"}}{{ .Port }} 
{{ end }}
`)

type listResultAgent struct {
	AgentType   string `json:"agent_type"`
	AgentID     string `json:"agent_id"`
	ServiceID   string `json:"service_id"`
	Status      string `json:"status"`
	Disabled    bool   `json:"disabled"`
	MetricsMode string `json:"push_metrics_enabled"`
	Port        int64  `json:"port,omitempty"`
}

func (a listResultAgent) HumanReadableAgentType() string {
	return types.AgentTypeName(a.AgentType)
}

func (a listResultAgent) NiceAgentStatus() string {
	res := a.Status
	if res == "" {
		res = "unknown" //nolint:goconst
	}
	res = cases.Title(language.English).String(strings.ToLower(res))
	if a.Disabled {
		res += " (disabled)"
	}
	return res
}

type listResultService struct {
	ServiceType string `json:"service_type"`
	ServiceID   string `json:"service_id"`
	ServiceName string `json:"service_name"`
	AddressPort string `json:"address_port"`
	Group       string `json:"external_group"`
}

func (s listResultService) HumanReadableServiceType() string {
	serviceTypeName := types.ServiceTypeName(s.ServiceType)

	if s.ServiceType == types.ServiceTypeExternalService {
		return fmt.Sprintf("%s:%s", serviceTypeName, s.Group)
	}

	return serviceTypeName
}

type listResult struct {
	Services []listResultService `json:"service"`
	Agents   []listResultAgent   `json:"agent"`
}

func (res *listResult) Result() {}

func (res *listResult) String() string {
	template := RenderTemplate(listResultT, res)
	formattedTemplate, err := convertTabs(template)
	if err != nil {
		logrus.Panicf("Failed to render response: %s", err)
	}
	return formattedTemplate
}

func convertTabs(template string) (string, error) {
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 4, 4, 8, ' ', tabwriter.TabIndent)
	if _, err := io.WriteString(w, template); err != nil {
		return "", err
	}
	if err := w.Flush(); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// ListCommand is used by Kong for CLI flags and commands.
type ListCommand struct {
	NodeID string `help:"Node ID (default is autodetected)"`
}

// RunCmd executes the List command and returns the result.
func (cmd *ListCommand) RunCmd() (Result, error) {
	if cmd.NodeID == "" {
		status, err := agentlocal.GetStatus(agentlocal.DoNotRequestNetworkInfo)
		if err != nil {
			return nil, err
		}
		cmd.NodeID = status.NodeID
	}

	servicesRes, err := client.Default.ServicesService.ListServices(&services.ListServicesParams{
		Body: services.ListServicesBody{
			NodeID: cmd.NodeID,
		},
		Context: Ctx,
	})
	if err != nil {
		return nil, err
	}

	agentsRes, err := client.Default.AgentsService.ListAgents(&agents.ListAgentsParams{
		Context: Ctx,
	})
	if err != nil {
		return nil, err
	}

	return &listResult{
		Services: servicesList(servicesRes),
		Agents:   agentsList(agentsRes, cmd.NodeID),
	}, nil
}

func servicesList(servicesRes *services.ListServicesOK) []listResultService {
	getSocketOrHost := func(socket, address string, port int64) string {
		if socket != "" {
			return socket
		}
		return net.JoinHostPort(address, strconv.FormatInt(port, 10))
	}

	l := len(servicesRes.Payload.Mysql) + len(servicesRes.Payload.Mongodb) + len(servicesRes.Payload.Postgresql)
	l += len(servicesRes.Payload.Proxysql) + len(servicesRes.Payload.Haproxy) + len(servicesRes.Payload.External)
	servicesList := make([]listResultService, 0, l)

	for _, s := range servicesRes.Payload.Mysql {
		servicesList = append(servicesList, listResultService{
			ServiceType: types.ServiceTypeMySQLService,
			ServiceID:   s.ServiceID,
			ServiceName: s.ServiceName,
			AddressPort: getSocketOrHost(s.Socket, s.Address, s.Port),
		})
	}
	for _, s := range servicesRes.Payload.Mongodb {
		servicesList = append(servicesList, listResultService{
			ServiceType: types.ServiceTypeMongoDBService,
			ServiceID:   s.ServiceID,
			ServiceName: s.ServiceName,
			AddressPort: getSocketOrHost(s.Socket, s.Address, s.Port),
		})
	}
	for _, s := range servicesRes.Payload.Postgresql {
		servicesList = append(servicesList, listResultService{
			ServiceType: types.ServiceTypePostgreSQLService,
			ServiceID:   s.ServiceID,
			ServiceName: s.ServiceName,
			AddressPort: getSocketOrHost(s.Socket, s.Address, s.Port),
		})
	}
	for _, s := range servicesRes.Payload.Proxysql {
		servicesList = append(servicesList, listResultService{
			ServiceType: types.ServiceTypeProxySQLService,
			ServiceID:   s.ServiceID,
			ServiceName: s.ServiceName,
			AddressPort: getSocketOrHost(s.Socket, s.Address, s.Port),
		})
	}
	for _, s := range servicesRes.Payload.Haproxy {
		servicesList = append(servicesList, listResultService{
			ServiceType: types.ServiceTypeHAProxyService,
			ServiceID:   s.ServiceID,
			ServiceName: s.ServiceName,
		})
	}
	for _, s := range servicesRes.Payload.External {
		servicesList = append(servicesList, listResultService{
			ServiceType: types.ServiceTypeExternalService,
			ServiceID:   s.ServiceID,
			ServiceName: s.ServiceName,
			Group:       s.Group,
		})
	}

	return servicesList
}

func agentsList(agentsRes *agents.ListAgentsOK, nodeID string) []listResultAgent { //nolint:cyclop
	getStatus := func(s *string) string {
		res := pointer.GetString(s)
		if res == "" {
			res = "unknown"
		}
		return strings.ToUpper(res)
	}
	getMetricsMode := func(s bool) string {
		if s {
			return "push"
		}

		return "pull"
	}
	pmmAgentIDs := make(map[string]struct{})
	agentsList := []listResultAgent{}
	for _, a := range agentsRes.Payload.PMMAgent {
		if a.RunsOnNodeID == nodeID {
			pmmAgentIDs[a.AgentID] = struct{}{}

			status := "disconnected"
			if a.Connected {
				status = "connected"
			}
			agentsList = append(agentsList, listResultAgent{
				AgentType: types.AgentTypePMMAgent,
				AgentID:   a.AgentID,
				Status:    strings.ToUpper(status),
			})
		}
	}
	for _, a := range agentsRes.Payload.NodeExporter {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agentsList = append(agentsList, listResultAgent{
				AgentType:   types.AgentTypeNodeExporter,
				AgentID:     a.AgentID,
				Status:      getStatus(a.Status),
				Disabled:    a.Disabled,
				MetricsMode: getMetricsMode(a.PushMetricsEnabled),
				Port:        a.ListenPort,
			})
		}
	}
	for _, a := range agentsRes.Payload.MysqldExporter {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agentsList = append(agentsList, listResultAgent{
				AgentType:   types.AgentTypeMySQLdExporter,
				AgentID:     a.AgentID,
				ServiceID:   a.ServiceID,
				Status:      getStatus(a.Status),
				Disabled:    a.Disabled,
				MetricsMode: getMetricsMode(a.PushMetricsEnabled),
				Port:        a.ListenPort,
			})
		}
	}
	for _, a := range agentsRes.Payload.MongodbExporter {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agentsList = append(agentsList, listResultAgent{
				AgentType:   types.AgentTypeMongoDBExporter,
				AgentID:     a.AgentID,
				ServiceID:   a.ServiceID,
				Status:      getStatus(a.Status),
				Disabled:    a.Disabled,
				MetricsMode: getMetricsMode(a.PushMetricsEnabled),
				Port:        a.ListenPort,
			})
		}
	}
	for _, a := range agentsRes.Payload.PostgresExporter {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agentsList = append(agentsList, listResultAgent{
				AgentType:   types.AgentTypePostgresExporter,
				AgentID:     a.AgentID,
				ServiceID:   a.ServiceID,
				Status:      getStatus(a.Status),
				Disabled:    a.Disabled,
				MetricsMode: getMetricsMode(a.PushMetricsEnabled),
				Port:        a.ListenPort,
			})
		}
	}
	for _, a := range agentsRes.Payload.ProxysqlExporter {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agentsList = append(agentsList, listResultAgent{
				AgentType:   types.AgentTypeProxySQLExporter,
				AgentID:     a.AgentID,
				ServiceID:   a.ServiceID,
				Status:      getStatus(a.Status),
				Disabled:    a.Disabled,
				MetricsMode: getMetricsMode(a.PushMetricsEnabled),
				Port:        a.ListenPort,
			})
		}
	}
	for _, a := range agentsRes.Payload.RDSExporter {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agentsList = append(agentsList, listResultAgent{
				AgentType:   types.AgentTypeRDSExporter,
				AgentID:     a.AgentID,
				Status:      getStatus(a.Status),
				Disabled:    a.Disabled,
				MetricsMode: getMetricsMode(a.PushMetricsEnabled),
				Port:        a.ListenPort,
			})
		}
	}
	for _, a := range agentsRes.Payload.QANMysqlPerfschemaAgent {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agentsList = append(agentsList, listResultAgent{
				AgentType: types.AgentTypeQANMySQLPerfSchemaAgent,
				AgentID:   a.AgentID,
				ServiceID: a.ServiceID,
				Status:    getStatus(a.Status),
				Disabled:  a.Disabled,
			})
		}
	}
	for _, a := range agentsRes.Payload.QANMysqlSlowlogAgent {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agentsList = append(agentsList, listResultAgent{
				AgentType: types.AgentTypeQANMySQLSlowlogAgent,
				AgentID:   a.AgentID,
				ServiceID: a.ServiceID,
				Status:    getStatus(a.Status),
				Disabled:  a.Disabled,
			})
		}
	}
	for _, a := range agentsRes.Payload.QANMongodbProfilerAgent {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agentsList = append(agentsList, listResultAgent{
				AgentType: types.AgentTypeQANMongoDBProfilerAgent,
				AgentID:   a.AgentID,
				ServiceID: a.ServiceID,
				Status:    getStatus(a.Status),
				Disabled:  a.Disabled,
			})
		}
	}
	for _, a := range agentsRes.Payload.QANPostgresqlPgstatementsAgent {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agentsList = append(agentsList, listResultAgent{
				AgentType: types.AgentTypeQANPostgreSQLPgStatementsAgent,
				AgentID:   a.AgentID,
				ServiceID: a.ServiceID,
				Status:    getStatus(a.Status),
				Disabled:  a.Disabled,
			})
		}
	}
	for _, a := range agentsRes.Payload.QANPostgresqlPgstatmonitorAgent {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agentsList = append(agentsList, listResultAgent{
				AgentType: types.AgentTypeQANPostgreSQLPgStatMonitorAgent,
				AgentID:   a.AgentID,
				ServiceID: a.ServiceID,
				Status:    getStatus(a.Status),
				Disabled:  a.Disabled,
			})
		}
	}
	for _, a := range agentsRes.Payload.ExternalExporter {
		if a.RunsOnNodeID == nodeID {
			agentsList = append(agentsList, listResultAgent{
				AgentType:   types.AgentTypeExternalExporter,
				AgentID:     a.AgentID,
				ServiceID:   a.ServiceID,
				Status:      getStatus(nil),
				Disabled:    a.Disabled,
				MetricsMode: getMetricsMode(a.PushMetricsEnabled),
				Port:        a.ListenPort,
			})
		}
	}
	for _, a := range agentsRes.Payload.VMAgent {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agentsList = append(agentsList, listResultAgent{
				AgentType:   types.AgentTypeVMAgent,
				AgentID:     a.AgentID,
				Status:      getStatus(a.Status),
				MetricsMode: getMetricsMode(true),
				Port:        a.ListenPort,
			})
		}
	}

	return agentsList
}
