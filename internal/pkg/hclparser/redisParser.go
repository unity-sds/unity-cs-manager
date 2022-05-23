package hclparser

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func parseRedis(f *hclwrite.File, bl string) error {

	// ACL NAME
	f, err := addAttribute(f, fmt.Sprintf("%v.acl_name", bl), "open-access", false)
	if err != nil {
		fmt.Printf("%v", err)
	}
	// NAME
	f, err = addAttribute(f, fmt.Sprintf("%v.name", bl), "unity-cluster", false)
	if err != nil {
		fmt.Printf("%v", err)
	}
	// Node Type
	f, err = addAttribute(f, fmt.Sprintf("%v.node_type", bl), "db.t4g.small", false)
	if err != nil {
		fmt.Printf("%v", err)
	}
	// Num Shards
	f, err = addAttribute(f, fmt.Sprintf("%v.num_shards", bl), "2", false)
	if err != nil {
		fmt.Printf("%v", err)
	}
	// Snapshot Retention Limit
	f, err = addAttribute(f, fmt.Sprintf("%v.snapshot_retention_limit", bl), "7", false)
	if err != nil {
		fmt.Printf("%v", err)
	}
	//TODO
	// Security Group IDs
	f, err = addAttribute(f, fmt.Sprintf("%v.security_group_ids", bl), "open-access", false)
	if err != nil {
		fmt.Printf("%v", err)
	}
	//TODO
	// Subnet Group Name
	f, err = addAttribute(f, fmt.Sprintf("%v.subnet_group_name", bl), "open-access", false)
	if err != nil {
		fmt.Printf("%v", err)
	}
	// Tags

	return nil
}