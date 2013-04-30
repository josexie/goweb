package handlers

import (
	"github.com/stretchrcom/goweb/context"
	"github.com/stretchrcom/goweb/paths"
)

func (h *HttpHandler) Map(pathPattern string, executor func(context.Context) error) error {

	path, pathErr := paths.NewPathPattern(pathPattern)

	if pathErr != nil {
		return pathErr
	}

	handler := &PathMatchHandler{path, executor}
	h.AppendHandler(handler)

	// ok
	return nil

}