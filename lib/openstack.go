package main

import (
	"fmt"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
	"github.com/rackspace/gophercloud/openstack/compute/v2/extensions/secgroups"
)

//SG is struct for slice SecurityGroup struct from gophercloud openstack package
type SG secgroups.SecurityGroup

//SR is struct for security rule for a security group
type SR secgroups.Rule

// CreateRO provides struct for create rule options
type CreateRO secgroups.CreateRuleOpts

// CreateGO provides struct for create group options
type CreateGO secgroups.GroupOpts

// CreateClient retuns a openstack client
func CreateClient(*gophercloud.ServiceClient) {
	opts, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		fmt.Printf("%v", err)
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		panic(err)
	}

	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{Region: "jp-east2-a-prod"})
	if err != nil {
		panic(err)
	}

	return client
}

// GetSecGroups fetches all security groups
func GetSecGroups(client gophercloud.ServiceClient) ([]SG, error) {
	pages, err := secgroups.List(client).AllPages()
	if err != nil {
		panic(err)
	}
	groups, err := secgroups.ExtractSecurityGroups(pages)
	if err != nil {
		panic(err)
	}
	return groups, err
}

//CreateSecGroup creates security group
func CreateSecGroup(client gophercloud.ServiceClient, opts CreateGO) (SG, error) {
	group, err := secgroups.Create(client, opts).Extract()
	if err != nil {
		panic(err)
	}
	return group, nil
}
