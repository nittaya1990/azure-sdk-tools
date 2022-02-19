//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package test_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2020-12-01/examples/GetDiskRestorePointResources.json
func ExampleDiskRestorePointClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := test.NewDiskRestorePointClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<restore-point-collection-name>",
		"<vm-restore-point-name>",
		"<disk-restore-point-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DiskRestorePointClientGetResult)
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2020-12-01/examples/ListDiskRestorePointsInVmRestorePoint.json
func ExampleDiskRestorePointClient_ListByRestorePoint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := test.NewDiskRestorePointClient("<subscription-id>", cred, nil)
	pager := client.ListByRestorePoint("<resource-group-name>",
		"<restore-point-collection-name>",
		"<vm-restore-point-name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}