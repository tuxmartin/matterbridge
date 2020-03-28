// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// VPNConfiguration Base VPN Configuration profile.
type VPNConfiguration struct {
	// DeviceConfiguration is the base model of VPNConfiguration
	DeviceConfiguration
	// AuthenticationMethod Authentication method.
	AuthenticationMethod *VPNAuthenticationMethod `json:"authenticationMethod,omitempty"`
	// ConnectionName Connection name displayed to the user.
	ConnectionName *string `json:"connectionName,omitempty"`
	// Role Role when connection type is set to Pulse Secure.
	Role *string `json:"role,omitempty"`
	// Realm Realm when connection type is set to Pulse Secure.
	Realm *string `json:"realm,omitempty"`
	// Servers List of VPN Servers on the network. Make sure end users can access these network locations. This collection can contain a maximum of 500 elements.
	Servers []VPNServer `json:"servers,omitempty"`
}

// VPNDNSRule undocumented
type VPNDNSRule struct {
	// Object is the base model of VPNDNSRule
	Object
	// Name Name.
	Name *string `json:"name,omitempty"`
	// Servers Servers.
	Servers []string `json:"servers,omitempty"`
	// ProxyServerURI Proxy Server Uri.
	ProxyServerURI *string `json:"proxyServerUri,omitempty"`
	// AutoTrigger Automatically connect to the VPN when the device connects to this domain: Default False.
	AutoTrigger *bool `json:"autoTrigger,omitempty"`
	// Persistent Keep this rule active even when the VPN is not connected: Default False
	Persistent *bool `json:"persistent,omitempty"`
}

// VPNOnDemandRule undocumented
type VPNOnDemandRule struct {
	// Object is the base model of VPNOnDemandRule
	Object
	// Ssids Network Service Set Identifiers (SSIDs).
	Ssids []string `json:"ssids,omitempty"`
	// DNSSearchDomains DNS Search Domains.
	DNSSearchDomains []string `json:"dnsSearchDomains,omitempty"`
	// ProbeURL A URL to probe. If this URL is successfully fetched (returning a 200 HTTP status code) without redirection, this rule matches.
	ProbeURL *string `json:"probeUrl,omitempty"`
	// Action Action.
	Action *VPNOnDemandRuleConnectionAction `json:"action,omitempty"`
	// DomainAction Domain Action (Only applicable when Action is evaluate connection).
	DomainAction *VPNOnDemandRuleConnectionDomainAction `json:"domainAction,omitempty"`
	// Domains Domains (Only applicable when Action is evaluate connection).
	Domains []string `json:"domains,omitempty"`
	// ProbeRequiredURL Probe Required Url (Only applicable when Action is evaluate connection and DomainAction is connect if needed).
	ProbeRequiredURL *string `json:"probeRequiredUrl,omitempty"`
}

// VPNProxyServer undocumented
type VPNProxyServer struct {
	// Object is the base model of VPNProxyServer
	Object
	// AutomaticConfigurationScriptURL Proxy's automatic configuration script url.
	AutomaticConfigurationScriptURL *string `json:"automaticConfigurationScriptUrl,omitempty"`
	// Address Address.
	Address *string `json:"address,omitempty"`
	// Port Port. Valid values 0 to 65535
	Port *int `json:"port,omitempty"`
}

// VPNRoute undocumented
type VPNRoute struct {
	// Object is the base model of VPNRoute
	Object
	// DestinationPrefix Destination prefix (IPv4/v6 address).
	DestinationPrefix *string `json:"destinationPrefix,omitempty"`
	// PrefixSize Prefix size. (1-32). Valid values 1 to 32
	PrefixSize *int `json:"prefixSize,omitempty"`
}

// VPNServer undocumented
type VPNServer struct {
	// Object is the base model of VPNServer
	Object
	// Description Description.
	Description *string `json:"description,omitempty"`
	// Address Address (IP address, FQDN or URL)
	Address *string `json:"address,omitempty"`
	// IsDefaultServer Default server.
	IsDefaultServer *bool `json:"isDefaultServer,omitempty"`
}

// VPNTrafficRule undocumented
type VPNTrafficRule struct {
	// Object is the base model of VPNTrafficRule
	Object
	// Name Name.
	Name *string `json:"name,omitempty"`
	// Protocols Protocols (0-255). Valid values 0 to 255
	Protocols *int `json:"protocols,omitempty"`
	// LocalPortRanges Local port range can be set only when protocol is either TCP or UDP (6 or 17). This collection can contain a maximum of 500 elements.
	LocalPortRanges []NumberRange `json:"localPortRanges,omitempty"`
	// RemotePortRanges Remote port range can be set only when protocol is either TCP or UDP (6 or 17). This collection can contain a maximum of 500 elements.
	RemotePortRanges []NumberRange `json:"remotePortRanges,omitempty"`
	// LocalAddressRanges Local address range. This collection can contain a maximum of 500 elements.
	LocalAddressRanges []IPv4Range `json:"localAddressRanges,omitempty"`
	// RemoteAddressRanges Remote address range. This collection can contain a maximum of 500 elements.
	RemoteAddressRanges []IPv4Range `json:"remoteAddressRanges,omitempty"`
	// AppID App identifier, if this traffic rule is triggered by an app.
	AppID *string `json:"appId,omitempty"`
	// AppType App type, if this traffic rule is triggered by an app.
	AppType *VPNTrafficRuleAppType `json:"appType,omitempty"`
	// RoutingPolicyType When app triggered, indicates whether to enable split tunneling along this route.
	RoutingPolicyType *VPNTrafficRuleRoutingPolicyType `json:"routingPolicyType,omitempty"`
	// Claims Claims associated with this traffic rule.
	Claims *string `json:"claims,omitempty"`
}