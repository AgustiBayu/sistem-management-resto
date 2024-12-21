package helper

import (
	"SistemManagementResto/model/web"
	"fmt"
)

func CreateLinksForItem(id int, resource string) []web.Link {
	return []web.Link{
		{Rel: "self", Href: fmt.Sprintf("/%s/%d", resource, id)},
		{Rel: "update", Href: fmt.Sprintf("/%s/%d", resource, id)},
		{Rel: "delete", Href: fmt.Sprintf("/%s/%d", resource, id)},
	}
}

func CreateLinksForItems(ids []int, entity string) []web.Link {
	var allLinks []web.Link
	for _, id := range ids {
		allLinks = append(allLinks, CreateLinksForItem(id, entity)...)
	}
	return allLinks
}
