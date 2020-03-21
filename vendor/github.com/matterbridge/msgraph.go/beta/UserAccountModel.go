// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// UserAccount undocumented
type UserAccount struct {
	// Object is the base model of UserAccount
	Object
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LastSeenDateTime undocumented
	LastSeenDateTime *time.Time `json:"lastSeenDateTime,omitempty"`
	// RiskScore undocumented
	RiskScore *string `json:"riskScore,omitempty"`
	// Service undocumented
	Service *string `json:"service,omitempty"`
	// SigninName undocumented
	SigninName *string `json:"signinName,omitempty"`
	// Status undocumented
	Status *AccountStatus `json:"status,omitempty"`
}