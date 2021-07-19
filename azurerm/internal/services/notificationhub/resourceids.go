package notificationhub

//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Namespace -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.NotificationHubs/namespaces/namespace1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=NotificationHub -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.NotificationHubs/namespaces/namespace1/NotificationHubs/hub1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=NotificationHubAuthorizationRule -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.NotificationHubs/namespaces/namespace1/NotificationHubs/hub1/AuthorizationRules/authorizationRule1
