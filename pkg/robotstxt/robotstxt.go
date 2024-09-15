package robotstxt

import "github.com/jimsmart/grobotstxt"

// NewParser Create new Parser
func NewParser(robotsTxtBody string, agentNames ...string) (*Parser, error) {
	var agentName string
	if agentNames != nil && agentNames[0] != "" {
		agentName = agentNames[0]
	}
	return &Parser{
		AgentName:  agentName,
		RobotsBody: robotsTxtBody,
		Sitemaps:   grobotstxt.Sitemaps(robotsTxtBody),
	}, nil
}

// Parser Robots.txt Parser
type Parser struct {
	// RobotsBody a body of robots.txt
	RobotsBody string
	// AgentName Name of the Agent
	AgentName string
	// Sitemaps a list of sitemaps
	Sitemaps []string
}

// Allowed Check given path is allowed
func (p *Parser) Allowed(path string) bool {
	return grobotstxt.AgentAllowed(p.RobotsBody, p.AgentName, path)
}

// GetSitemaps Get Sitemaps
func (p *Parser) GetSitemaps() []string {
	return p.Sitemaps
}
