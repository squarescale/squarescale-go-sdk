package models

// Project is the main concept of the Squarescale
// plateform. It carries the
type Project struct {
	Name                       string                      `json:"name"`
	UUID                       string                      `json:"uuid"`
	Provider                   string                      `json:"provider"`
	Region                     string                      `json:"region"`
	CredentialName             string                      `json:"credential_name"`
	Clusters                   []Cluster                   `json:"clusters"`
	NotificationConfigurations []NotificationConfiguration `json:"notification_configuration"`
	Organization               *Organization               `json:"organization"`
	MonitoringConfiguration    *MonitoringConfiguration    `json:"monitoring_configuration"`
	ManagedServices            ManagedServices             `json:"managed_services"`
}

type MonitoringConfiguration struct {
	Engine string `json:"engine"`
}

type ManagedServices struct {
	Databases []Database `json:"databases"`
}

type NotificationConfiguration struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}
