package instana_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/require"

	. "github.com/gessnerfl/terraform-provider-instana/instana"
	"github.com/gessnerfl/terraform-provider-instana/instana/restapi"
	"github.com/gessnerfl/terraform-provider-instana/testutils"
)

const resourceAlertingChannelSplunkDefinitionTemplate = `
resource "instana_alerting_channel_splunk" "example" {
	name  = "name %d"
	url   = "url"
	token = "token"
}
`

const alertingChannelSplunkServerResponseTemplate = `
{
	"id"    : "%s",
	"name"  : "prefix name %d suffix",
	"kind"  : "SPLUNK",
	"url"   : "url",
	"token" : "token"
}
`

const testAlertingChannelSplunkDefinition = "instana_alerting_channel_splunk.example"

func TestCRUDOfAlertingChannelSplunkResourceWithMockServer(t *testing.T) {
	httpServer := createMockHttpServerForResource(restapi.AlertingChannelsResourcePath, alertingChannelSplunkServerResponseTemplate)
	httpServer.Start()
	defer httpServer.Close()

	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			createAlertingChannelSplunkResourceTestStep(httpServer.GetPort(), 0),
			testStepImport(testAlertingChannelSplunkDefinition),
			createAlertingChannelSplunkResourceTestStep(httpServer.GetPort(), 1),
			testStepImport(testAlertingChannelSplunkDefinition),
		},
	})
}

func createAlertingChannelSplunkResourceTestStep(httpPort int, iteration int) resource.TestStep {
	config := appendProviderConfig(fmt.Sprintf(resourceAlertingChannelSplunkDefinitionTemplate, iteration), httpPort)
	return resource.TestStep{
		Config: config,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(testAlertingChannelSplunkDefinition, "id"),
			resource.TestCheckResourceAttr(testAlertingChannelSplunkDefinition, AlertingChannelFieldName, formatResourceName(iteration)),
			resource.TestCheckResourceAttr(testAlertingChannelSplunkDefinition, AlertingChannelFieldFullName, formatResourceFullName(iteration)),
			resource.TestCheckResourceAttr(testAlertingChannelSplunkDefinition, AlertingChannelSplunkFieldURL, "url"),
			resource.TestCheckResourceAttr(testAlertingChannelSplunkDefinition, AlertingChannelSplunkFieldToken, "token"),
		),
	}
}

func TestResourceAlertingChannelSplunkDefinition(t *testing.T) {
	resource := NewAlertingChannelSplunkResourceHandle()

	schemaMap := resource.MetaData().Schema

	schemaAssert := testutils.NewTerraformSchemaAssert(schemaMap, t)
	schemaAssert.AssertSchemaIsRequiredAndOfTypeString(AlertingChannelFieldName)
	schemaAssert.AssertSchemaIsComputedAndOfTypeString(AlertingChannelFieldFullName)
	schemaAssert.AssertSchemaIsRequiredAndOfTypeString(AlertingChannelSplunkFieldURL)
	schemaAssert.AssertSchemaIsRequiredAndOfTypeString(AlertingChannelSplunkFieldToken)
}

func TestShouldUpdateResourceStateForAlertingChanneSplunk(t *testing.T) {
	testHelper := NewTestHelper(t)
	resourceHandle := NewAlertingChannelSplunkResourceHandle()
	resourceData := testHelper.CreateEmptyResourceDataForResourceHandle(resourceHandle)
	url := "url"
	token := "token"
	data := restapi.AlertingChannel{
		ID:    "id",
		Name:  resourceFullName,
		URL:   &url,
		Token: &token,
	}

	err := resourceHandle.UpdateState(resourceData, &data, testHelper.ResourceFormatter())

	require.Nil(t, err)
	require.Equal(t, "id", resourceData.Id())
	require.Equal(t, "name", resourceData.Get(AlertingChannelFieldName))
	require.Equal(t, resourceFullName, resourceData.Get(AlertingChannelFieldFullName))
	require.Equal(t, url, resourceData.Get(AlertingChannelSplunkFieldURL))
	require.Equal(t, token, resourceData.Get(AlertingChannelSplunkFieldToken))
}

func TestShouldConvertStateOfAlertingChannelSplunkToDataModel(t *testing.T) {
	testHelper := NewTestHelper(t)
	resourceHandle := NewAlertingChannelSplunkResourceHandle()
	url := "url"
	token := "token"
	resourceData := testHelper.CreateEmptyResourceDataForResourceHandle(resourceHandle)
	resourceData.SetId("id")
	resourceData.Set(AlertingChannelFieldName, "name")
	resourceData.Set(AlertingChannelFieldFullName, resourceFullName)
	resourceData.Set(AlertingChannelSplunkFieldURL, url)
	resourceData.Set(AlertingChannelSplunkFieldToken, token)

	model, err := resourceHandle.MapStateToDataObject(resourceData, testHelper.ResourceFormatter())

	require.Nil(t, err)
	require.IsType(t, &restapi.AlertingChannel{}, model, "Model should be an alerting channel")
	require.Equal(t, "id", model.GetIDForResourcePath())
	require.Equal(t, resourceFullName, model.(*restapi.AlertingChannel).Name, "name should be equal to full name")
	require.Equal(t, url, *model.(*restapi.AlertingChannel).URL, "url should be equal")
	require.Equal(t, token, *model.(*restapi.AlertingChannel).Token, "token should be equal")
}

func TestAlertingChannelSplunkkShouldHaveSchemaVersionZero(t *testing.T) {
	require.Equal(t, 0, NewAlertingChannelSplunkResourceHandle().MetaData().SchemaVersion)
}

func TestAlertingChannelSplunkShouldHaveNoStateUpgrader(t *testing.T) {
	require.Equal(t, 0, len(NewAlertingChannelSplunkResourceHandle().StateUpgraders()))
}

func TestShouldReturnCorrectResourceNameForAlertingChannelSplunk(t *testing.T) {
	name := NewAlertingChannelSplunkResourceHandle().MetaData().ResourceName

	require.Equal(t, name, "instana_alerting_channel_splunk")
}
