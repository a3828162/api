package models

import (
	"time"

	"github.com/free5gc/openapi/models"
)

type DNPerformanceAnalyticsInputData struct {
	UEIdentifier                     models.IpAddr       `json:"ue_identifier" yaml:"ue_identifier"`
	UElocation                       models.UserLocation `json:"ue_location,omitempty" yaml:"ue_location,omitempty"`
	ApplicationId                    string              `json:"application_id" yaml:"application_id"`
	IPFilter                         []models.IpAddr     `json:"ip_filter" yaml:"ip_filter"`
	LocationOfApplication            []string            `json:"location_of_application" yaml:"location_of_application"`
	ApplicationServerInstanceAddress models.IpAddr       `json:"application_server_instance_address" yaml:"application_server_instance_address"`
	PerformanceData                  []PerformanceData   `json:"performance_data" yaml:"performance_data"`
	TimeStamp                        time.Time           `json:"timestamp" yaml:"timestamp"`
}

type DNPerformance struct {
	ApplicationServerInstanceAddress models.IpAddr     `json:"application_server_instance_address" yaml:"application_server_instance_address"`
	ServingAnchorUPF                 string            `json:"serving_anchor_upf,omitempty" yaml:"serving_anchor_upf,omitempty"`
	DNAI                             string            `json:"dnai" yaml:"dnai"`
	Performance                      OutputPerformance `json:"performance_data" yaml:"performance_data"`
}

type OutputPerformance struct {
	AverageTrafficRate    float64 `json:"average_traffic_rate" yaml:"average_traffic_rate"`
	MaximumTrafficRate    float64 `json:"maximum_traffic_rate" yaml:"maximum_traffic_rate"`
	AveragePacketDelay    float64 `json:"average_packet_delay" yaml:"average_packet_delay"`
	MaximumPacketDelay    float64 `json:"maximum_packet_delay" yaml:"maximum_packet_delay"`
	AveragePakcetLossRate float64 `json:"average_packet_loss_rate" yaml:"average_packet_loss_rate"`
}
