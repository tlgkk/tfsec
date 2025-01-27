package elb

import "github.com/aquasecurity/defsec/types"

type ELB struct {
	types.Metadata
	LoadBalancers []LoadBalancer
}

const (
	TypeApplication = "application"
	TypeGateway     = "gateway"
	TypeNetwork     = "network"
)

type LoadBalancer struct {
	types.Metadata
	Type                    types.StringValue
	DropInvalidHeaderFields types.BoolValue
	Internal                types.BoolValue
	Listeners               []Listener
}

type Listener struct {
	types.Metadata
	Protocol      types.StringValue
	TLSPolicy     types.StringValue
	DefaultAction Action
}

type Action struct {
	types.Metadata
	Type types.StringValue
}

func (c *LoadBalancer) GetMetadata() *types.Metadata {
	return &c.Metadata
}

func (c *LoadBalancer) GetRawValue() interface{} {
	return nil
}

func (c *Listener) GetMetadata() *types.Metadata {
	return &c.Metadata
}

func (c *Listener) GetRawValue() interface{} {
	return nil
}

func (e *ELB) GetMetadata() *types.Metadata {
	return &e.Metadata
}

func (e *ELB) GetRawValue() interface{} {
	return nil
}

func (a *Action) GetMetadata() *types.Metadata {
	return &a.Metadata
}

func (a *Action) GetRawValue() interface{} {
	return nil
}
