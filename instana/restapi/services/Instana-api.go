package services

import "github.com/gessnerfl/terraform-provider-instana/instana/restapi"
import "github.com/gessnerfl/terraform-provider-instana/instana/restapi/resources"

//NewInstanaAPI creates a new instance of the instana API
func NewInstanaAPI(apiToken string, endpoint string, validateServerCertificate bool) restapi.InstanaAPI {
	client := NewClient(apiToken, endpoint, validateServerCertificate)
	return &baseInstanaAPI{client: client}
}

type baseInstanaAPI struct {
	client restapi.RestClient
}

//Rules implementation of InstanaAPI interface
func (api baseInstanaAPI) Rules() restapi.RuleResource {
	return resources.NewRuleResource(api.client)
}

//Rules implementation of InstanaAPI interface
func (api baseInstanaAPI) RuleBindings() restapi.RuleBindingResource {
	return resources.NewRuleBindingResource(api.client)
}