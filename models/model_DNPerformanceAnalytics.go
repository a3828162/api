package models

import (
	"github.com/free5gc/openapi/models"
)

type DNPerformanceAnalyticsFilter struct {
	ApplicationName          string                                     `json:"application_name" yaml:"application_name"`
	Snssai                   models.Snssai                              `json:"snssai" yaml:"snssai"`
	NsiIds                   []string                                   `json:"nsi_id" yaml:"nsi_id"`
	AreaOfInterest           map[string]models.AreaOfInterestEventState `json:"area_of_interest,omitempty" yaml:"area_of_interest,omitempty"`
	UPFAnchorIdentity        string
	DNN                      string        `json:"dnn" yaml:"dnn"`
	DNAI                     string        `json:"dnai" yaml:"dnai"`
	ApplicationServerAddress models.IpAddr `json:"application_server_address" yaml:"application_server_address"`
	// ListOfAnalyticsSubsets   []string      `json:"list_of_analytics_subsets" yaml:"list_of_analytics_subsets"`
}

type DNPerformanceAnalyticsStatistic struct {
	ApplicationId             string                  `json:"application_id" yaml:"application_id"`
	Snssai                    models.Snssai           `json:"snssai" yaml:"snssai"`
	DNN                       string                  `json:"dnn" yaml:"dnn"`
	DNPerformances            []DNPerformance         `json:"dn_performances" yaml:"dn_performances"`
	SpatialValidityCondition  models.SpatialValidity  `json:"spatial_validity_condition,omitempty" yaml:"spatial_validity_condition,omitempty"`
	TemporalValidityCondition models.TemporalValidity `json:"temporal_validity_condition,omitempty" yaml:"temporal_validity_condition,omitempty"`
}

type DNPerformanceAnalyticsPrediction struct {
	ApplicationId             string                  `json:"application_id" yaml:"application_id"`
	Snssai                    models.Snssai           `json:"snssai" yaml:"snssai"`
	DNN                       string                  `json:"dnn" yaml:"dnn"`
	DNPerformances            []DNPerformance         `json:"dn_performances" yaml:"dn_performances"`
	SpatialValidityCondition  models.SpatialValidity  `json:"spatial_validity_condition,omitempty" yaml:"spatial_validity_condition,omitempty"`
	TemporalValidityCondition models.TemporalValidity `json:"temporal_validity_condition,omitempty" yaml:"temporal_validity_condition,omitempty"`
	Confidence                float64                 `json:"confidence" yaml:"confidence"`
}
