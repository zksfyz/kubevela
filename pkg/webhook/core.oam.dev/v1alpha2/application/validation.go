package application

import (
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/oam-dev/kubevela/apis/core.oam.dev/v1alpha2"
	"github.com/oam-dev/kubevela/pkg/appfile"
)

// ValidateCreate validates the Application on creation
func (h *ValidatingHandler) ValidateCreate(app *v1alpha2.Application) field.ErrorList {
	var componentErrs field.ErrorList
	// try to generate an app file
	appParser := appfile.NewApplicationParser(h.Client, h.dm)
	if _, err := appParser.GenerateAppFile(app.Name, app); err != nil {
		componentErrs = append(componentErrs, field.Invalid(field.NewPath("spec"), app, err.Error()))
	}
	return componentErrs
}

// ValidateUpdate validates the Application on update
func (h *ValidatingHandler) ValidateUpdate(newApp, oldApp *v1alpha2.Application) field.ErrorList {
	// check if the newApp is valid
	componentErrs := h.ValidateCreate(newApp)
	// TODO: add more validating
	return componentErrs
}
