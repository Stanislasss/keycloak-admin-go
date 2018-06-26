//go:generate gomodifytags -file $GOFILE -struct UserRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct UserConsentRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct CredentialRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct FederatedIdentityRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate easyjson -all $GOFILE

package keycloak

import (
	"net/url"
	"time"
)

// UserConsentRepresentation represents client consents
type UserConsentRepresentation struct {
	ClientID            string     `json:"clientId,omitempty"`
	CreatedDate         *time.Time `json:"createdDate,omitempty"`
	GrantedClientScopes []string   `json:"grantedClientScopes,omitempty"`
	LastUpdatedDate     *time.Time `json:"lastUpdatedDate,omitempty"`
}

// CredentialRepresentation represents credentials for a user or client
type CredentialRepresentation struct {
	Algorithm         string     `json:"algorithm,omitempty"`
	Counter           int32      `json:"counter,omitempty"`
	CreatedDate       *time.Time `json:"createdDate,omitempty"`
	Device            string     `json:"device,omitempty"`
	Digits            int32      `json:"digits,omitempty"`
	HashIterations    int32      `json:"hashIterations,omitempty"`
	HashedSaltedValue string     `json:"hashedSaltedValue,omitempty"`
	Period            int32      `json:"period,omitempty"`
	Salt              string     `json:"salt,omitempty"`
	Temporary         bool       `json:"temporary,omitempty"`
	Type              string     `json:"type,omitempty"`
	Value             string     `json:"value,omitempty"`
}

// FederatedIdentityRepresentation represents a federated identity
type FederatedIdentityRepresentation struct {
	IdentityProvider string `json:"identityProvider,omitempty"`
	UserID           string `json:"userId,omitempty"`
	UserName         string `json:"userName,omitempty"`
}

// UserRepresentation represents a realm user in Keycloak
type UserRepresentation struct {
	Access                 AttributeMap                      `json:"access,omitempty"`
	Attributes             AttributeMap                      `json:"attributes,omitempty"`
	ClientRoles            AttributeMap                      `json:"clientRoles,omitempty"`
	ClientConsents         []UserConsentRepresentation       `json:"clientConsents,omitempty"`
	CreatedTimestamp       *time.Time                        `json:"createdTimestamp,omitempty"`
	Credentials            []CredentialRepresentation        `json:"credentials,omitempty"`
	DisableCredentialTypes []string                          `json:"disableCredentialTypes,omitempty"`
	Email                  string                            `json:"email,omitempty"`
	EmailVerified          bool                              `json:"emailVerified,omitempty"`
	Enabled                bool                              `json:"enabled,omitempty"`
	FederatedIDentities    []FederatedIdentityRepresentation `json:"federatedIdentities,omitempty"`
	FederationLink         *url.URL                          `json:"federationLink,omitempty"`
	FirstName              string                            `json:"firstName,omitempty"`
	Groups                 []string                          `json:"groups,omitempty"`
	ID                     string                            `json:"id,omitempty"`
	LastName               string                            `json:"lastName,omitempty"`
	NotBefore              *time.Time                        `json:"notBefore,omitempty"`
	Origin                 string                            `json:"origin,omitempty"`
	RealmRoles             []string                          `json:"realmRoles,omitempty"`
	RequiredActions        []string                          `json:"requiredActions,omitempty"`
	Self                   string                            `json:"self,omitempty"`
	ServiceAccountClientID string                            `json:"serviceAccountClientId,omitempty"`
	Username               string                            `json:"username,omitempty"`
}