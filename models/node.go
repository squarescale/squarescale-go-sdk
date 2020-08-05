package models

// ClusterType is the type used to configure
// cluster type information
type ClusterType string

var (
	// HighAvailability stand for a cluster of stateless node
	// that do has a minimal size of 3. Usefull for some productions
	// workload
	HighAvailability ClusterType = "high_availability"

	// SingleNode stand for a cluster of stateless node with only
	// one node. Usefull for development purposed or with some stateful nodes
	SingleNode ClusterType = "single_node"
)

// A Cluster is a set of node with same behavior
// and same properties.
//
// it expose information on node type and configuration
type Cluster struct {
	// Type of that cluster. The behavior and the properties
	// of this cluster heavily depends on that value
	Type ClusterType `json:"cluster_type"`

	// NodeType stand for the "class" or size of the cluster nodes.
	// Available values are provider dependant
	NodeType string `json:"node_type"`
}
