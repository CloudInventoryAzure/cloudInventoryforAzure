package azurelib

import (
        "context"
        "fmt"
        "testing"
        "time"
)

const subscriptionID = "282160c0-3c83-43f1-bff1-9356b1678ffb"
const autorizationError = "Failed to authorize: %v"
const getallVmserror = "Failed to  get all VMs: %v"
const networkInterfaceerror = "Failed to  get the network interface of Vm  %v : %v"
const resourceGrouperror = "Failed to  get resource group of VM %v : %v"

func GetAuthorizedclients(subscriptionID string) (client Clients, err error) {
        clients := GetNewClients(subscriptionID)
        client, err = AuthorizeClients(clients)
        return
}

//TestGetallVMS tests function TestGetallVMS
func TestGetallVMS(t *testing.T) {

        client, err := GetAuthorizedclients(subscriptionID)
        if err != nil {
                t.Errorf(autorizationError, err)
        }
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
        defer cancel()
        Vmlist, err := GetallVMS(client, ctx)
        if err != nil {
                t.Errorf(getallVmserror, err)
        } else {
                t.Logf("GetallVMS successful")
                fmt.Println(Vmlist)
        }
}

//TestGetVmnetworkinterface tests function GetVmnetworkinterface
func TestGetVmnetworkinterface(t *testing.T) {
        client, err := GetAuthorizedclients(subscriptionID)
        if err != nil {
                t.Errorf(autorizationError, err)
        }
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
        defer cancel()
        Vmlist, err := GetallVMS(client, ctx)
        if err != nil {
                t.Errorf("Failed to  get all VMs: %v", err)
        }
        for i := 0; i < len(Vmlist); i++ {
                networkInterface, err := GetVmnetworkinterface(Vmlist[i])
                if err != nil {
                        t.Errorf(networkInterfaceerror, *Vmlist[i].Name, err)
                } else {
                        t.Logf("GetVmnetworkinterface successful  Vm name and its network interface  %v : %v", *Vmlist[i].Name, networkInterface)
                }
        }
}

//TestGetPrivateIP tests function GetPrivateIP
func TestGetPrivateIP(t *testing.T) {
        client, err := GetAuthorizedclients(subscriptionID)
        if err != nil {
                t.Errorf(autorizationError, err)
        }
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
        defer cancel()
        Vmlist, err := GetallVMS(client, ctx)
        if err != nil {
                t.Errorf(getallVmserror, err)
        }
        for i := 0; i < len(Vmlist); i++ {
                networkInterface, err := GetVmnetworkinterface(Vmlist[i])
                if err != nil {
                        t.Errorf(networkInterfaceerror, *Vmlist[i].Name, err)
                }
                resourceGroup, err := GetVMResourcegroup(Vmlist[i])
                if err != nil {
                        t.Errorf(resourceGrouperror, *Vmlist[i].Name, err)
                }
                privateIPaddress, IPconfig, err := GetPrivateIP(client, ctx, resourceGroup, networkInterface, "")
                if err != nil {
                        t.Errorf("Failed to  get the privateIPaddress of Vm  %v : %v", *Vmlist[i].Name, err)
                } else {
                        t.Logf("GetPrivateIP successful  Vm name and its IP configuration , privateIPaddress  %v :%v, %v", *Vmlist[i].Name, IPconfig, privateIPaddress)
                }
        }
}

//TestGetPublicIPAddress tests function GetPublicIPAddress
func TestGetPublicIPAddress(t *testing.T) {
        client, err := GetAuthorizedclients(subscriptionID)
        if err != nil {
                t.Errorf(autorizationError, err)
        }
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
        defer cancel()
        Vmlist, err := GetallVMS(client, ctx)
        if err != nil {
                t.Errorf(getallVmserror, err)
        }
        for i := 0; i < len(Vmlist); i++ {
                networkInterface, err := GetVmnetworkinterface(Vmlist[i])
                if err != nil {
                        t.Errorf(networkInterfaceerror, *Vmlist[i].Name, err)
                }

                resourceGroup, err := GetVMResourcegroup(Vmlist[i])
                if err != nil {
                        t.Errorf(resourceGrouperror, *Vmlist[i].Name, err)
                }
                publicIPname, err := GetPublicIPAddressID(client, ctx, resourceGroup, networkInterface, "")
                if err != nil {
                        t.Errorf("Failed to  get the publicIPname of Vm  %v : %v", *Vmlist[i].Name, err)
                }
                publicIPaddress, err := GetPublicIPAddress(client, ctx, resourceGroup, publicIPname, "")
                if err != nil {
                        t.Errorf("Failed to  get PublicIPaddress of VM %v : %v", *Vmlist[i].Name, err)
                } else {
                        t.Logf("GetPublicIPAddress successful  Vm name and its publicIPaddress  %v : %v", *Vmlist[i].Name, publicIPaddress)
                }
        }
}

//TestGetSubnetandvirtualnetwork tests the function GetSubnetandvirtualnetwork
func TestGetSubnetandvirtualnetwork(t *testing.T) {
        client, err := GetAuthorizedclients(subscriptionID)
        if err != nil {
                t.Errorf(autorizationError, err)
        }
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
        defer cancel()
        Vmlist, err := GetallVMS(client, ctx)
        if err != nil {
                t.Errorf(getallVmserror, err)
        }
        for i := 0; i < len(Vmlist); i++ {
                networkInterface, err := GetVmnetworkinterface(Vmlist[i])
                if err != nil {
                        t.Errorf(networkInterfaceerror, *Vmlist[i].Name, err)
                }
                resourceGroup, err := GetVMResourcegroup(Vmlist[i])
                if err != nil {
                        t.Errorf(resourceGrouperror, *Vmlist[i].Name, err)
                }
                subnetAndvirtualnetwork, err := GetSubnetandvirtualnetwork(client, ctx, resourceGroup, networkInterface, "")
                if err != nil {
                        t.Errorf("Failed to  get the Subnetandvirtualnetwork of Vm  %v : %v", *Vmlist[i].Name, err)
                } else {
                        t.Logf("GetSubnetandvirtualnetwork successful  Vm name and its subnet and virtual network  %v : %v", *Vmlist[i].Name, subnetAndvirtualnetwork)
                }
        }
}

//TestGetDNS tests the function GetDNS
func TestGetDNS(t *testing.T) {
        clients := GetNewClients(subscriptionID)
        client, err := AuthorizeClients(clients)
        if err != nil {
                t.Errorf(autorizationError, err)
        }
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
        defer cancel()
        Vmlist, err := GetallVMS(client, ctx)
        if err != nil {
                t.Errorf(getallVmserror, err)
        }
        for i := 0; i < len(Vmlist); i++ {
                networkInterface, err := GetVmnetworkinterface(Vmlist[i])
                if err != nil {
                        t.Errorf(networkInterfaceerror, *Vmlist[i].Name, err)
                }

                resourceGroup, err := GetVMResourcegroup(Vmlist[i])
                if err != nil {
                        t.Errorf(resourceGrouperror, *Vmlist[i].Name, err)
                }
                publicIPname, err := GetPublicIPAddressID(client, ctx, resourceGroup, networkInterface, "")
                if err != nil {
                        t.Errorf("Failed to  get the publicIPname of Vm  %v : %v", *Vmlist[i].Name, err)
                }
                VMDNS, err := GetDNS(client, ctx, resourceGroup, publicIPname, "")
                if err != nil {
                        t.Errorf("Failed to  get DNS of VM %v : %v", *Vmlist[i].Name, err)
                } else {
                        t.Logf("GetDNS successful  Vm name and its DNS  %v : %v", *Vmlist[i].Name, VMDNS)
                }
        }
}