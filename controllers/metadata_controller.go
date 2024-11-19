package controllers

import (
	"fmt"

	"github.com/ssebs/go-mmp/models"
	"github.com/ssebs/go-mmp/views"
)

type MetadataController struct {
	*models.Metadata
	*views.MetadataEditorView
}

func NewMetadataController(m *models.Metadata, v *views.MetadataEditorView) *MetadataController {
	mc := &MetadataController{
		Metadata:           m,
		MetadataEditorView: v,
	}

	mc.MetadataEditorView.SetOnSubmit(func(m models.Metadata) {
		mc.UpdateAllFields(m)

		fmt.Println("Updated")
		fmt.Println(mc.Metadata)
	})

	mc.UpdateMetadataView()

	return mc
}

func (mc *MetadataController) UpdateMetadataView() {
	mc.MetadataEditorView.SetMetadata(mc.Metadata)
}
