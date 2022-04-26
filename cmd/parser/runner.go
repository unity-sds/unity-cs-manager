package main

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"io/ioutil"
)

func main() {

	fbyte, err := ioutil.ReadFile("test/ec2.tf")
	if err != nil {
		fmt.Println(err)
	}
	fwr, diag := hclwrite.ParseConfig(fbyte, "example.tf", hcl.Pos{Line: 1, Column: 1})
	if diag.HasErrors() {
		fmt.Println("Couldn't parse file")
	}
	//fmt.Printf("%v", string(fwr.Bytes()))
	//aaf := NewAttributeAppendFilter("resource.aws_instance.unity-ec2-instance.tags.unityname", "mytag", false)
	//fwr, err = aaf.Filter(fwr)
	err = addTagsToBlocks(fwr)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%v", string(fwr.Bytes()))

}

func getFiles() {

}

func getBlocks(f *hclwrite.File) ([]string, error) {

	bl := NewBlockListSink()
	return bl.Sink(f)

}

func addTagsToBlocks(f *hclwrite.File) error {
	blocks, err := getBlocks(f)
	if err != nil {
		return err
	}
	for _, bl := range blocks {
		aaf := NewAttributeAppendFilter(fmt.Sprintf("%v.tags.unityname", bl), "myunitydeployment", false)
		f, err = aaf.Filter(f)
		if err != nil {
			return err
		}
	}

	baf := NewBlockAppendFilter("resource.aws_eip.ip-test-env", "tags", true)
	f, err = baf.Filter(f)
	if err != nil {
		return err
	}
	aaf := NewAttributeAppendFilter(fmt.Sprintf("%v.tags.unityname", "resource.aws_eip.ip-test-env"), "myunitydeployment", false)
	f, err = aaf.Filter(f)
	if err != nil {
		return err
	}
	/*as := attributeSet{
		address: "resource.aws_eip.ip-test-env.tags.unityname",
		value:   "myunitydeployment",
	}
	f, err = as.Filter(f)*/
	return err
}