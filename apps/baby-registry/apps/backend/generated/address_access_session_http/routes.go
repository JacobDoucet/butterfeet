package address_access_session_http

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

	selectByTokenUniqueHandler, err := GetSelectByTokenUniqueHandler(handlerProps)
	if err != nil {
		return nil, err
	}
	routes.HandleFunc("/tokenHash/{tokenHash}", selectByTokenUniqueHandler)

	aggregateHandler, err := GetAggregateHandler(handlerProps)
	if err != nil {
		return nil, err
	}
	routes.HandleFunc("/aggregate", aggregateHandler)

	return routes, nil
}
