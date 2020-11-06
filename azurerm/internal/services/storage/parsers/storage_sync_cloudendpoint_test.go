package parsers

import "testing"

func TestSyncCloudEndpointID(t *testing.T) {
	testData := []struct {
		Name     string
		Input    string
		Expected *SyncCloudEndpointId
	}{
		{
			Name:     "Empty",
			Input:    "",
			Expected: nil,
		},
		{
			Name:     "No Resource Groups Segment",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000",
			Expected: nil,
		},
		{
			Name:     "No Resource Groups Value",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/",
			Expected: nil,
		},
		{
			Name:     "Resource Group ID",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/foo/",
			Expected: nil,
		},
		{
			Name:     "Missing Sync Group",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.StorageSync/storageSyncServices/sync1",
			Expected: nil,
		},
		{
			Name:     "Missing Sync Group Value",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.StorageSync/storageSyncServices/sync1/syncGroups/",
			Expected: nil,
		},
		{
			Name:     "Missing Cloud Endpoint",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.StorageSync/storageSyncServices/sync1/syncGroups/group1",
			Expected: nil,
		},
		{
			Name:     "Missing Cloud Endpoint Value",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.StorageSync/storageSyncServices/sync1/syncGroups/group1/cloudEndpoints",
			Expected: nil,
		},
		{
			Name:  "Sync Group Id",
			Input: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.StorageSync/storageSyncServices/sync1/syncGroups/group1/cloudEndpoints/cloudEndpoint1",
			Expected: &SyncCloudEndpointId{
				Name:             "cloudEndpoint1",
				StorageSyncGroup: "group1",
				StorageSyncName:  "sync1",
				ResourceGroup:    "resGroup1",
			},
		},
		{
			Name:     "Wrong Casing",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.StorageSync/storageSyncServices/sync1/syncGroups/group1/CloudEndpoints/cloudEndpoints1",
			Expected: nil,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Name)

		actual, err := SyncCloudEndpointID(v.Input)
		if err != nil {
			if v.Expected == nil {
				continue
			}

			t.Fatalf("Expected a value but got an error: %s", err)
		}

		if actual.Name != v.Expected.Name {
			t.Fatalf("Expected %q but got %q for Name", v.Expected.Name, actual.Name)
		}

		if actual.StorageSyncGroup != v.Expected.StorageSyncGroup {
			t.Fatalf("Expected %q but got %q for Storage Sync Group", v.Expected.Name, actual.Name)
		}

		if actual.StorageSyncName != v.Expected.StorageSyncName {
			t.Fatalf("Expected %q but got %q for Storage Sync Name", v.Expected.Name, actual.Name)
		}

		if actual.ResourceGroup != v.Expected.ResourceGroup {
			t.Fatalf("Expected %q but got %q for Resource Group", v.Expected.ResourceGroup, actual.ResourceGroup)
		}
	}
}
