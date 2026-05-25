package registry_approved_guest_http

import (
	"net/http"
	"strings"
)

func formatEndpoint(rootPath, endpoint string) string {
	return strings.ReplaceAll(rootPath+"/"+endpoint, "//", "/")
}

func RegisterRoutes(handlerProps HandlerProps) (*http.ServeMux, error) {
	if err := handlerProps.Validate(); err != nil {
		return nil, err
	}

	routes := http.NewServeMux()

	searchHandler, err := GetSearchHandler(handlerProps)
	if err != nil {
		return nil, err
	}
	routes.HandleFunc("/search", searchHandler)

	selectByIdHandler, err := GetSelectByIdHandler(handlerProps)
	if err != nil {
		return nil, err
	}
	routes.HandleFunc("/id/{id}", selectByIdHandler)

	selectByRegistryEmailUniqueHandler, err := GetSelectByRegistryEmailUniqueHandler(handlerProps)
	if err != nil {
		return nil, err
	}
	routes.HandleFunc("/registryId/{registryId}/emailHash/{emailHash}", selectByRegistryEmailUniqueHandler)

	createHandler, err := GetCreateHandler(handlerProps)
	if err != nil {
		return nil, err
	}
	routes.HandleFunc("/create", createHandler)

	updateHandler, err := GetUpdateHandler(handlerProps)
	if err != nil {
		return nil, err
	}
	routes.HandleFunc("/update", updateHandler)

	deleteHandler, err := GetDeleteHandler(handlerProps)
	if err != nil {
		return nil, err
	}
	routes.HandleFunc("/delete/{id}", deleteHandler)

	return routes, nil
}
