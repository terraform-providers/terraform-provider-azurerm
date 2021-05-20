package validate

// NOTE: this file is generated via 'go:generate' - manual changes will be overwritten

import "testing"

func TestDomainTopicID(t *testing.T) {
	cases := []struct {
		Input string
		Valid bool
	}{

		{
			// empty
			Input: "",
			Valid: false,
		},

		{
			// missing SubscriptionId
			Input: "/",
			Valid: false,
		},

		{
			// missing value for SubscriptionId
			Input: "/subscriptions/",
			Valid: false,
		},

		{
			// missing ResourceGroup
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/",
			Valid: false,
		},

		{
			// missing value for ResourceGroup
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/",
			Valid: false,
		},

		{
			// missing DomainName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.EventGrid/",
			Valid: false,
		},

		{
			// missing value for DomainName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.EventGrid/domains/",
			Valid: false,
		},

		{
			// missing TopicName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.EventGrid/domains/domain1/",
			Valid: false,
		},

		{
			// missing value for TopicName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.EventGrid/domains/domain1/topics/",
			Valid: false,
		},

		{
			// valid
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.EventGrid/domains/domain1/topics/topic1",
			Valid: true,
		},

		{
			// upper-cased
			Input: "/SUBSCRIPTIONS/12345678-1234-9876-4563-123456789012/RESOURCEGROUPS/RESGROUP1/PROVIDERS/MICROSOFT.EVENTGRID/DOMAINS/DOMAIN1/TOPICS/TOPIC1",
			Valid: false,
		},
	}
	for _, tc := range cases {
		t.Logf("[DEBUG] Testing Value %s", tc.Input)
		_, errors := DomainTopicID(tc.Input, "test")
		valid := len(errors) == 0

		if tc.Valid != valid {
			t.Fatalf("Expected %t but got %t", tc.Valid, valid)
		}
	}
}
