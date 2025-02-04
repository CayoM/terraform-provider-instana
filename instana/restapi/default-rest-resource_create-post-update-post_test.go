package restapi_test

import (
	"encoding/json"
	"errors"
	"testing"

	. "github.com/gessnerfl/terraform-provider-instana/instana/restapi"
	mocks "github.com/gessnerfl/terraform-provider-instana/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSuccessfulCreateOfTestObjectThroughCreatePOSTUpdatePOSTRestResource(t *testing.T) {
	executeCreateOrUpdateOperationThroughCreatePOSTUpdatePOSTRestResourceTest(t, func(t *testing.T, sut RestResource, client *mocks.MockRestClient, unmarshaller *mocks.MockJSONUnmarshaller) {
		testObject := makeTestObject()
		serializedJSON, _ := json.Marshal(testObject)

		client.EXPECT().Post(gomock.Eq(testObject), gomock.Eq(testObjectResourcePath)).Return(serializedJSON, nil)
		client.EXPECT().PostWithID(gomock.Any(), gomock.Eq(testObjectResourcePath)).Times(0)
		unmarshaller.EXPECT().Unmarshal(serializedJSON).Times(1).Return(testObject, nil)

		result, err := sut.Create(testObject)

		assert.NoError(t, err)
		assert.Equal(t, testObject, result)
	})
}

func TestShouldFailToCreateTestObjectThroughCreatePOSTUpdatePOSTRestResourceWhenErrorIsReturnedFromRestClient(t *testing.T) {
	executeCreateOrUpdateOperationThroughCreatePOSTUpdatePOSTRestResourceTest(t, func(t *testing.T, sut RestResource, client *mocks.MockRestClient, unmarshaller *mocks.MockJSONUnmarshaller) {
		testObject := makeTestObject()

		client.EXPECT().Post(gomock.Eq(testObject), gomock.Eq(testObjectResourcePath)).Return(nil, errors.New("Error during test"))
		client.EXPECT().PostWithID(gomock.Any(), gomock.Eq(testObjectResourcePath)).Times(0)
		client.EXPECT().Put(gomock.Any(), gomock.Eq(testObjectResourcePath)).Times(0)
		unmarshaller.EXPECT().Unmarshal(gomock.Any()).Times(0)

		_, err := sut.Create(testObject)

		assert.Error(t, err)
	})
}

func TestShouldFailToCreateTestObjectThroughCreatePOSTUpdatePOSTRestResourceWhenResponseCannotBeUnmarshalled(t *testing.T) {
	executeCreateOrUpdateOperationThroughCreatePOSTUpdatePOSTRestResourceTest(t, func(t *testing.T, sut RestResource, client *mocks.MockRestClient, unmarshaller *mocks.MockJSONUnmarshaller) {
		testObject := makeTestObject()
		expectedError := errors.New("test")

		client.EXPECT().Post(gomock.Eq(testObject), gomock.Eq(testObjectResourcePath)).Return(invalidResponse, nil)
		client.EXPECT().PostWithID(gomock.Any(), gomock.Eq(testObjectResourcePath)).Times(0)
		client.EXPECT().Put(gomock.Any(), gomock.Eq(testObjectResourcePath)).Times(0)
		unmarshaller.EXPECT().Unmarshal(invalidResponse).Times(1).Return(nil, expectedError)

		_, err := sut.Create(testObject)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestShouldFailToCreateTestObjectThroughCreatePOSTUpdatePOSTRestResourceWhenTheUnmarshalledResponseIsNotImplementingInstanaDataObject(t *testing.T) {
	executeCreateOrUpdateOperationThroughCreatePOSTUpdatePOSTRestResourceTest(t, func(t *testing.T, sut RestResource, client *mocks.MockRestClient, unmarshaller *mocks.MockJSONUnmarshaller) {
		testObject := makeTestObject()
		response := []byte("{ \"invalid\" : \"testObject\" }")

		client.EXPECT().Post(gomock.Eq(testObject), gomock.Eq(testObjectResourcePath)).Return(response, nil)
		client.EXPECT().PostWithID(gomock.Any(), gomock.Eq(testObjectResourcePath)).Times(0)
		client.EXPECT().Put(gomock.Any(), gomock.Eq(testObjectResourcePath)).Times(0)
		unmarshaller.EXPECT().Unmarshal(response).Times(1).Return(&InvalidInstanaDataObject{}, nil)

		_, err := sut.Create(testObject)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unmarshalled object does not implement InstanaDataObject")
	})
}

func TestShouldFailedToCreateTestObjectThroughCreatePOSTUpdatePOSTRestResourceWhenAnInvalidTestObjectIsProvided(t *testing.T) {
	executeCreateOrUpdateOperationThroughCreatePOSTUpdatePOSTRestResourceTest(t, func(t *testing.T, sut RestResource, client *mocks.MockRestClient, unmarshaller *mocks.MockJSONUnmarshaller) {
		testObject := &testObject{
			ID:   "some id",
			Name: "invalid name",
		}

		client.EXPECT().Post(gomock.Eq(testObject), gomock.Eq(testObjectResourcePath)).Times(0)
		client.EXPECT().PostWithID(gomock.Any(), gomock.Eq(testObjectResourcePath)).Times(0)
		client.EXPECT().Put(gomock.Any(), gomock.Eq(testObjectResourcePath)).Times(0)
		unmarshaller.EXPECT().Unmarshal(gomock.Any()).Times(0)

		_, err := sut.Create(testObject)

		assert.Error(t, err)
	})
}

func TestShouldFailedToCreateTestObjectThroughCreatePOSTUpdatePOSTRestResourceWhenAnInvalidTestObjectIsReceived(t *testing.T) {
	executeCreateOrUpdateOperationThroughCreatePOSTUpdatePOSTRestResourceTest(t, func(t *testing.T, sut RestResource, client *mocks.MockRestClient, unmarshaller *mocks.MockJSONUnmarshaller) {
		object := makeTestObject()

		client.EXPECT().Post(gomock.Eq(object), gomock.Eq(testObjectResourcePath)).Times(1).Return(invalidResponse, nil)
		client.EXPECT().PostWithID(gomock.Any(), gomock.Eq(testObjectResourcePath)).Times(0)
		client.EXPECT().Put(gomock.Any(), gomock.Eq(testObjectResourcePath)).Times(0)
		unmarshaller.EXPECT().Unmarshal(invalidResponse).Times(1).Return(&testObject{ID: object.ID, Name: "invalid"}, nil)

		_, err := sut.Create(object)

		assert.Error(t, err)
	})
}

func TestSuccessfulUpdateOfTestObjectThroughCreatePOSTUpdatePOSTRestResource(t *testing.T) {
	executeCreateOrUpdateOperationThroughCreatePOSTUpdatePOSTRestResourceTest(t, func(t *testing.T, sut RestResource, client *mocks.MockRestClient, unmarshaller *mocks.MockJSONUnmarshaller) {
		testObject := makeTestObject()
		serializedJSON, _ := json.Marshal(testObject)

		client.EXPECT().PostWithID(gomock.Eq(testObject), gomock.Eq(testObjectResourcePath)).Return(serializedJSON, nil)
		client.EXPECT().Post(gomock.Any(), gomock.Eq(testObjectResourcePath)).Times(0)
		client.EXPECT().Put(gomock.Any(), gomock.Eq(testObjectResourcePath)).Times(0)
		unmarshaller.EXPECT().Unmarshal(serializedJSON).Times(1).Return(testObject, nil)

		result, err := sut.Update(testObject)

		assert.NoError(t, err)
		assert.Equal(t, testObject, result)
	})
}

func TestShouldFailToUpdateTestObjectThroughCreatePOSTUpdatePOSTRestResourceWhenErrorIsReturnedFromRestClient(t *testing.T) {
	executeCreateOrUpdateOperationThroughCreatePOSTUpdatePOSTRestResourceTest(t, func(t *testing.T, sut RestResource, client *mocks.MockRestClient, unmarshaller *mocks.MockJSONUnmarshaller) {
		testObject := makeTestObject()

		client.EXPECT().PostWithID(gomock.Eq(testObject), gomock.Eq(testObjectResourcePath)).Return(nil, errors.New("Error during test"))
		client.EXPECT().Post(gomock.Any(), gomock.Eq(testObjectResourcePath)).Times(0)
		client.EXPECT().Put(gomock.Any(), gomock.Eq(testObjectResourcePath)).Times(0)
		unmarshaller.EXPECT().Unmarshal(gomock.Any()).Times(0)

		_, err := sut.Update(testObject)

		assert.Error(t, err)
	})
}

func TestShouldFailToUpdateTestObjectThroughCreatePOSTUpdatePOSTRestResourceWhenResponseCannotBeUnmarshalled(t *testing.T) {
	executeCreateOrUpdateOperationThroughCreatePOSTUpdatePOSTRestResourceTest(t, func(t *testing.T, sut RestResource, client *mocks.MockRestClient, unmarshaller *mocks.MockJSONUnmarshaller) {
		testObject := makeTestObject()
		expectedError := errors.New("test")

		client.EXPECT().PostWithID(gomock.Eq(testObject), gomock.Eq(testObjectResourcePath)).Return(invalidResponse, nil)
		client.EXPECT().Post(gomock.Any(), gomock.Eq(testObjectResourcePath)).Times(0)
		client.EXPECT().Put(gomock.Any(), gomock.Eq(testObjectResourcePath)).Times(0)
		unmarshaller.EXPECT().Unmarshal(invalidResponse).Times(1).Return(nil, expectedError)

		_, err := sut.Update(testObject)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestShouldFailToUpdateTestObjectThroughCreatePOSTUpdatePOSTRestResourceWhenTheUnmarshalledResponseIsNotImplementingInstanaDataObject(t *testing.T) {
	executeCreateOrUpdateOperationThroughCreatePOSTUpdatePOSTRestResourceTest(t, func(t *testing.T, sut RestResource, client *mocks.MockRestClient, unmarshaller *mocks.MockJSONUnmarshaller) {
		testObject := makeTestObject()
		response := []byte("{ \"invalid\" : \"testObject\" }")

		client.EXPECT().PostWithID(gomock.Eq(testObject), gomock.Eq(testObjectResourcePath)).Return(response, nil)
		client.EXPECT().Post(gomock.Any(), gomock.Eq(testObjectResourcePath)).Times(0)
		client.EXPECT().Put(gomock.Any(), gomock.Eq(testObjectResourcePath)).Times(0)
		unmarshaller.EXPECT().Unmarshal(response).Times(1).Return(&InvalidInstanaDataObject{}, nil)

		_, err := sut.Update(testObject)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unmarshalled object does not implement InstanaDataObject")
	})
}

func TestShouldFailedToUpdateTestObjectThroughCreatePOSTUpdatePOSTRestResourceWhenAnInvalidTestObjectIsProvided(t *testing.T) {
	executeCreateOrUpdateOperationThroughCreatePOSTUpdatePOSTRestResourceTest(t, func(t *testing.T, sut RestResource, client *mocks.MockRestClient, unmarshaller *mocks.MockJSONUnmarshaller) {
		testObject := &testObject{
			ID:   "some id",
			Name: "invalid name",
		}

		client.EXPECT().PostWithID(gomock.Any(), gomock.Eq(testObjectResourcePath)).Times(0)
		client.EXPECT().Post(gomock.Any(), gomock.Eq(testObjectResourcePath)).Times(0)
		client.EXPECT().Put(gomock.Eq(testObject), gomock.Eq(testObjectResourcePath)).Times(0)
		unmarshaller.EXPECT().Unmarshal(gomock.Any()).Times(0)

		_, err := sut.Update(testObject)

		assert.Error(t, err)
	})
}

func TestShouldFailedToUpdateTestObjectThroughCreatePOSTUpdatePOSTRestResourceWhenAnInvalidTestObjectIsReceived(t *testing.T) {
	executeCreateOrUpdateOperationThroughCreatePOSTUpdatePOSTRestResourceTest(t, func(t *testing.T, sut RestResource, client *mocks.MockRestClient, unmarshaller *mocks.MockJSONUnmarshaller) {
		object := makeTestObject()

		client.EXPECT().PostWithID(gomock.Eq(object), gomock.Eq(testObjectResourcePath)).Times(1).Return(invalidResponse, nil)
		client.EXPECT().Post(gomock.Any(), gomock.Eq(testObjectResourcePath)).Times(0)
		client.EXPECT().Put(gomock.Any(), gomock.Eq(testObjectResourcePath)).Times(0)
		unmarshaller.EXPECT().Unmarshal(invalidResponse).Times(1).Return(&testObject{ID: object.ID, Name: "invalid"}, nil)

		_, err := sut.Update(object)

		assert.Error(t, err)
	})
}

func executeCreateOrUpdateOperationThroughCreatePOSTUpdatePOSTRestResourceTest(t *testing.T, testFunction func(t *testing.T, sut RestResource, client *mocks.MockRestClient, unmarshaller *mocks.MockJSONUnmarshaller)) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)
	unmarshaller := mocks.NewMockJSONUnmarshaller(ctrl)

	sut := NewCreatePOSTUpdatePOSTRestResource(testObjectResourcePath, unmarshaller, client)

	testFunction(t, sut, client, unmarshaller)
}
