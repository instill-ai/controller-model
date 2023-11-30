package util

import (
	"fmt"
)

func ConvertUIDToResourcePermalink(uid string, resourceType string) string {
	resourceName := fmt.Sprintf("resources/%s/types/%s", uid, resourceType)

	return resourceName
}

func ConvertServiceToResourceName(serviceName string) string {
	resourceName := fmt.Sprintf("resources/%s/types/%s", serviceName, RESOURCE_TYPE_SERVICE)

	return resourceName
}

func ConvertResourcePermalinkToResourceRetryName(resourcePermalink string) string {
	resourceWorkflowId := fmt.Sprintf("%s/retry", resourcePermalink)

	return resourceWorkflowId
}

func ConvertResourcePermalinkToWorkflowName(resourcePermalink string) string {
	resourceWorkflowId := fmt.Sprintf("%s/workflow", resourcePermalink)

	return resourceWorkflowId
}
