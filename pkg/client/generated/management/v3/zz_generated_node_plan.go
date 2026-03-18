package client

const (
	NodePlanType                    = "nodePlan"
	NodePlanFieldAgentCheckInterval = "agentCheckInterval"
	NodePlanFieldVersion            = "version"
)

type NodePlan struct {
	AgentCheckInterval int64 `json:"agentCheckInterval,omitempty" yaml:"agentCheckInterval,omitempty"`
	Version            int64 `json:"version,omitempty" yaml:"version,omitempty"`
}
