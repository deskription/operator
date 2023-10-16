package v1alpha1

type ResourceGroupSelector struct {
	ApiGroup   string `json:"apiGroup,omitempty"`
	ApiVersion string `json:"apiVersion,omitempty"`
}

type ResourceKindSelector struct {
	ApiGroup   string `json:"apiGroup,omitempty"`
	ApiVersion string `json:"apiVersion,omitempty"`
	Kind       string `json:"kind,omitempty"`
}
