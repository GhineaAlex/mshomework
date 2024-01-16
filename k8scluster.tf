provider "azurerm" {
  features {}
  subscription_id = var.subscription_id
  app_id          = var.app_id
  password        = var.password
  tenant_id       = var.tenant_id
}

resource "azurerm_resource_group" "aks_rg" {
  name     = "mshome"
  location = "North Europe"
}

resource "azurerm_kubernetes_cluster" "aks_cluster" {
  name                = "msAKScluster"
  location            = azurerm_resource_group.aks_rg.location
  resource_group_name = azurerm_resource_group.aks_rg.name
  dns_prefix          = "msakscluster"

  default_node_pool {
    name       = "default"
    node_count = 1
    vm_size    = "Standard_B2s"
  }

  identity {
    type = "SystemAssigned"
  }

  role_based_access_control {
    enabled = true
  }

  service_principal {
    client_id     = var.client_id
    client_secret = var.client_secret
  }

  tags = {
    Environment = "Production"
  }
}

output "client_certificate" {
  value = azurerm_kubernetes_cluster.aks_cluster.kube_config.0.client_certificate
}

output "kube_config" {
  value = azurerm_kubernetes_cluster.aks_cluster.kube_config_raw
}

