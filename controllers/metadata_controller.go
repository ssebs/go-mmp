package controllers

import (
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

	mc.UpdateMetadataView()
	return mc
}
func (mc *MetadataController) SetOnSubmit(f func(m models.Metadata)) {
	mc.MetadataEditorView.SetOnSubmit(f)
}
func (mc *MetadataController) UpdateMetadataView() {
	mc.MetadataEditorView.SetMetadata(mc.Metadata)
}
