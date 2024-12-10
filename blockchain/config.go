package blockchain

import (
	"fmt"
	"net"
	"strings"
)

type NetworkConfig struct {
	// Listen address (e.g., "0.0.0.0" for all interfaces)
	ListenHost string
	// Listen port
	ListenPort int
	// Bootstrap nodes (e.g., ["ip:port", "ip:port"])
	BootstrapNodes []string
	// External IP (optional, for NAT traversal)
	ExternalIP string
	// Enable DHT server mode
	DHTServerMode bool
	// NAT configuration
	NATEnabled bool
	// UPnP port mapping
	UPnPEnabled bool
	// STUN server addresses
	STUNServers []string
	// TURN server configuration
	TURNServers []TURNConfig
	// Blockchain components
	Blockchain *Blockchain
	// Stake pool	
	StakePool  *StakePool
	// Mempool
	Mempool    *Mempool
	// UTXO pool
	UTXOPool   *UTXOPool
}

// TURNConfig holds TURN server configuration
type TURNConfig struct {
	Address  string
	Username string
	Password string
}

type NetworkPorts struct {
	BootstrapNodeBasePort int
	P2PBasePort           int
	RPCBasePort           int
}

var DefaultNetworkPorts = NetworkPorts{
	BootstrapNodeBasePort: 50500, // Base port for bootstrap nodes
	P2PBasePort:           51500, // Base port for peer-to-peer communication
	RPCBasePort:           52500, // Base port for RPC services
}

var DefaultBootnodeConfig = NetworkConfig{
	ListenHost:     "0.0.0.0",
	ListenPort:     50505, // Set to the fixed port for bootnode
	BootstrapNodes: []string{},
	ExternalIP:     "49.204.110.41", // Set to the provided public IP
	DHTServerMode:  true,
	NATEnabled:     true,
	UPnPEnabled:    true,
	STUNServers:    []string{},
	TURNServers:    []TURNConfig{},
}

var DefaultNetworkConfig = NetworkConfig{
	ListenHost:     "0.0.0.0",
	ListenPort:     9000,
	BootstrapNodes: []string{},
	DHTServerMode:  true,
	NATEnabled:     true,
	UPnPEnabled:    true,
	STUNServers: []string{
		"stun.l.google.com:19302",
		"stun1.l.google.com:19302",
	},
	TURNServers: []TURNConfig{},
}

func (np NetworkPorts) GetBootstrapNodePort(instanceID int) int {
	return np.BootstrapNodeBasePort + instanceID
}

func (np NetworkPorts) GetP2PPort(instanceID int) int {
	return np.P2PBasePort + instanceID
}

func (np NetworkPorts) GetRPCPort(instanceID int) int {
	return np.RPCBasePort + instanceID
}

func NewDefaultConfig() *NetworkConfig {
	return &NetworkConfig{
		ListenHost:     "0.0.0.0",
		ListenPort:     9000,
		BootstrapNodes: []string{},
		DHTServerMode:  true,
		NATEnabled:     true,
		UPnPEnabled:    true,
		STUNServers: []string{
			"stun.l.google.com:19302",
			"stun1.l.google.com:19302",
		},
		TURNServers: []TURNConfig{},
	}
}

func (c *NetworkConfig) GetMultiaddr() string {
	return fmt.Sprintf("/ip4/%s/tcp/%d", c.ListenHost, c.ListenPort)
}

func (c *NetworkConfig) ValidateConfig() error {
	// Validate listen host
	if c.ListenHost != "0.0.0.0" {
		if net.ParseIP(c.ListenHost) == nil {
			return fmt.Errorf("invalid listen host: %s", c.ListenHost)
		}
	}
	// Validate port
	if c.ListenPort < 0 || (c.ListenPort > 65535) || (c.ListenPort > 0 && c.ListenPort < 1024) {
		return fmt.Errorf("invalid port number: %d", c.ListenPort)
	}
	return nil
}

func isValidHostname(hostname string) bool {
	if len(hostname) > 255 {
		return false
	}
	for _, part := range strings.Split(hostname, ".") {
		if len(part) > 63 {
			return false
		}
	}
	return true
}