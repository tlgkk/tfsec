package synapse

import (
	"github.com/aquasecurity/defsec/provider/azure/synapse"
	"github.com/aquasecurity/tfsec/internal/pkg/block"
)

func Adapt(modules block.Modules) synapse.Synapse {
	return synapse.Synapse{
		Workspaces: adaptWorkspaces(modules),
	}
}

func adaptWorkspaces(modules block.Modules) []synapse.Workspace {
	var workspaces []synapse.Workspace
	for _, module := range modules {
		for _, resource := range module.GetResourcesByType("azurerm_synapse_workspace") {
			workspaces = append(workspaces, adaptWorkspace(resource))
		}
	}
	return workspaces
}

func adaptWorkspace(resource *block.Block) synapse.Workspace {
	enableManagedVNAttr := resource.GetAttribute("managed_virtual_network_enabled")
	enableManagedVNVal := enableManagedVNAttr.AsBoolValueOrDefault(false, resource)

	return synapse.Workspace{
		Metadata:                    resource.Metadata(),
		EnableManagedVirtualNetwork: enableManagedVNVal,
	}
}
