package helper

import (
	"SistemManagementResto/model/web"
	"fmt"
)

func CreateLinksForMenuItem(menuItem web.MenuItemResponse) []web.Link {
	return []web.Link{
		{Rel: "self", Href: fmt.Sprintf("/menuItems/%d", menuItem.Id)},
		{Rel: "update", Href: fmt.Sprintf("/menuItems/%d", menuItem.Id)},
		{Rel: "delete", Href: fmt.Sprintf("/menuItems/%d", menuItem.Id)},
	}
}

func CreateLinksForMenuItems(menuItems []web.MenuItemResponse) []web.Link {
	var links []web.Link
	for _, menuItem := range menuItems {
		links = append(links, CreateLinksForMenuItem(menuItem)...)
	}
	return links
}
