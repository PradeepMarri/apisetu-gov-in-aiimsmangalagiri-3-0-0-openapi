package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// AcademicCertificateSchema represents the AcademicCertificateSchema schema from the OpenAPI specification
type AcademicCertificateSchema struct {
	Validfromdate string `json:"validFromDate"`
	Issuedat string `json:"issuedAt"`
	Status string `json:"status"`
	TypeField string `json:"type"`
	Issuedto map[string]interface{} `json:"IssuedTo"`
	Language string `json:"language"`
	Name string `json:"name"`
	Issuedby map[string]interface{} `json:"IssuedBy"`
	Issuedate string `json:"issueDate"`
	Number int `json:"number"`
	Certificatedata map[string]interface{} `json:"CertificateData"`
}

// ConsentArtifactSchema represents the ConsentArtifactSchema schema from the OpenAPI specification
type ConsentArtifactSchema struct {
	Signature map[string]interface{} `json:"signature"`
	Consent map[string]interface{} `json:"consent"`
}
