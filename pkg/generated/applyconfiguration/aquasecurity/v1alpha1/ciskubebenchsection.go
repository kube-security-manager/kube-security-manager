// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// CISKubeBenchSectionApplyConfiguration represents an declarative configuration of the CISKubeBenchSection type for use
// with apply.
type CISKubeBenchSectionApplyConfiguration struct {
	ID        *string                               `json:"id,omitempty"`
	Version   *string                               `json:"version,omitempty"`
	Text      *string                               `json:"text,omitempty"`
	NodeType  *string                               `json:"node_type,omitempty"`
	TotalPass *int                                  `json:"total_pass,omitempty"`
	TotalFail *int                                  `json:"total_fail,omitempty"`
	TotalWarn *int                                  `json:"total_warn,omitempty"`
	TotalInfo *int                                  `json:"total_info,omitempty"`
	Tests     []CISKubeBenchTestsApplyConfiguration `json:"tests,omitempty"`
}

// CISKubeBenchSectionApplyConfiguration constructs an declarative configuration of the CISKubeBenchSection type for use with
// apply.
func CISKubeBenchSection() *CISKubeBenchSectionApplyConfiguration {
	return &CISKubeBenchSectionApplyConfiguration{}
}

// WithID sets the ID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ID field is set to the value of the last call.
func (b *CISKubeBenchSectionApplyConfiguration) WithID(value string) *CISKubeBenchSectionApplyConfiguration {
	b.ID = &value
	return b
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *CISKubeBenchSectionApplyConfiguration) WithVersion(value string) *CISKubeBenchSectionApplyConfiguration {
	b.Version = &value
	return b
}

// WithText sets the Text field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Text field is set to the value of the last call.
func (b *CISKubeBenchSectionApplyConfiguration) WithText(value string) *CISKubeBenchSectionApplyConfiguration {
	b.Text = &value
	return b
}

// WithNodeType sets the NodeType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeType field is set to the value of the last call.
func (b *CISKubeBenchSectionApplyConfiguration) WithNodeType(value string) *CISKubeBenchSectionApplyConfiguration {
	b.NodeType = &value
	return b
}

// WithTotalPass sets the TotalPass field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TotalPass field is set to the value of the last call.
func (b *CISKubeBenchSectionApplyConfiguration) WithTotalPass(value int) *CISKubeBenchSectionApplyConfiguration {
	b.TotalPass = &value
	return b
}

// WithTotalFail sets the TotalFail field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TotalFail field is set to the value of the last call.
func (b *CISKubeBenchSectionApplyConfiguration) WithTotalFail(value int) *CISKubeBenchSectionApplyConfiguration {
	b.TotalFail = &value
	return b
}

// WithTotalWarn sets the TotalWarn field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TotalWarn field is set to the value of the last call.
func (b *CISKubeBenchSectionApplyConfiguration) WithTotalWarn(value int) *CISKubeBenchSectionApplyConfiguration {
	b.TotalWarn = &value
	return b
}

// WithTotalInfo sets the TotalInfo field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TotalInfo field is set to the value of the last call.
func (b *CISKubeBenchSectionApplyConfiguration) WithTotalInfo(value int) *CISKubeBenchSectionApplyConfiguration {
	b.TotalInfo = &value
	return b
}

// WithTests adds the given value to the Tests field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tests field.
func (b *CISKubeBenchSectionApplyConfiguration) WithTests(values ...*CISKubeBenchTestsApplyConfiguration) *CISKubeBenchSectionApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTests")
		}
		b.Tests = append(b.Tests, *values[i])
	}
	return b
}
