package media

//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=MediaService -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Media/mediaservices/account1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Transform -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Media/mediaservices/account1/transforms/transform1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Asset -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Media/mediaservices/account1/assets/asset1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=StreamingEndpoint -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Media/mediaservices/account1/streamingendpoints/endpoint1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=ContentKeyPolicy -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Media/mediaservices/account1/contentkeypolicies/policy1
