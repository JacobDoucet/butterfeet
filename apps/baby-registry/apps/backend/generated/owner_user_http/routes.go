package owner_user_http

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

	selectByEmailUniqueHandler, err := GetSelectByEmailUniqueHandler(handlerProps)
	if err != nil {
		return nil, err
	}
	routes.HandleFunc("/email/{email}", selectByEmailUniqueHandler)

	updateHandler, err := GetUpdateHandler(handlerProps)
	if err != nil {
		return nil, err
	}
	routes.HandleFunc("/update", updateHandler)

	return routes, nil
}
