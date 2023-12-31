// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CISKubeBenchReportDataApplyConfiguration represents an declarative configuration of the CISKubeBenchReportData type for use
// with apply.
type CISKubeBenchReportDataApplyConfiguration struct {
	UpdateTimestamp *v1.Time                                `json:"updateTimestamp,omitempty"`
	Scanner         *ScannerApplyConfiguration              `json:"scanner,omitempty"`
	Summary         *CISKubeBenchSummaryApplyConfiguration  `json:"summary,omitempty"`
	Sections        []CISKubeBenchSectionApplyConfiguration `json:"sections,omitempty"`
}

// CISKubeBenchReportDataApplyConfiguration constructs an declarative configuration of the CISKubeBenchReportData type for use with
// apply.
func CISKubeBenchReportData() *CISKubeBenchReportDataApplyConfiguration {
	return &CISKubeBenchReportDataApplyConfiguration{}
}

// WithUpdateTimestamp sets the UpdateTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UpdateTimestamp field is set to the value of the last call.
func (b *CISKubeBenchReportDataApplyConfiguration) WithUpdateTimestamp(value v1.Time) *CISKubeBenchReportDataApplyConfiguration {
	b.UpdateTimestamp = &value
	return b
}

// WithScanner sets the Scanner field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Scanner field is set to the value of the last call.
func (b *CISKubeBenchReportDataApplyConfiguration) WithScanner(value *ScannerApplyConfiguration) *CISKubeBenchReportDataApplyConfiguration {
	b.Scanner = value
	return b
}

// WithSummary sets the Summary field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Summary field is set to the value of the last call.
func (b *CISKubeBenchReportDataApplyConfiguration) WithSummary(value *CISKubeBenchSummaryApplyConfiguration) *CISKubeBenchReportDataApplyConfiguration {
	b.Summary = value
	return b
}

// WithSections adds the given value to the Sections field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Sections field.
func (b *CISKubeBenchReportDataApplyConfiguration) WithSections(values ...*CISKubeBenchSectionApplyConfiguration) *CISKubeBenchReportDataApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithSections")
		}
		b.Sections = append(b.Sections, *values[i])
	}
	return b
}
