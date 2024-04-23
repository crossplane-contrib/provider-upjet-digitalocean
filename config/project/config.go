package project

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("digitalocean_project", func(r *config.Resource) {
		r.ShortGroup = "project"
		//r.ExternalName.GetIDFn = func(_ context.Context, externalName string, parameters map[string]any, terraformProviderConfig map[string]any) (string, error) {
		//	//return externalName, nil
		//
		//	id, err := uuid.NewUUID()
		//	if err != nil {
		//		return "", err
		//	}
		//	return id.String(), nil
		//}
		//r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		//	if id, ok := tfstate["id"].(string); ok && id != "" {
		//		return id, nil
		//	}
		//	return "", errors.New("cannot find id in tfstate")
		//}
		//r.ExternalName.OmittedFields = []string{}
		//r.ExternalName = config.ExternalName{
		//	SetIdentifierArgumentFn: func(base map[string]any, name string) {
		//		base["name"] = name
		//	},
		//	GetExternalNameFn: func(tfstate map[string]any) (string, error) {
		//		if id, ok := tfstate["id"].(string); ok && id != "" {
		//			return id, nil
		//		}
		//		return "", errors.New("cannot find id in tfstate")
		//	},
		//	GetIDFn: func(_ context.Context, externalName string, parameters map[string]any, terraformProviderConfig map[string]any) (string, error) {
		//		return externalName, nil
		//	},
		//	OmittedFields:          []string{},
		//	DisableNameInitializer: true,
		//	IdentifierFields:       nil,
		//}
	})
}
