package restapi

import "errors"

//NewCreatePUTUpdatePUTRestResource creates a new REST resource using the provided unmarshaller function to convert the response from the REST API to the corresponding InstanaDataObject. The REST resource is using PUT as operation for create and update
func NewCreatePUTUpdatePUTRestResource(resourcePath string, unmarshaller JSONUnmarshaller, client RestClient) RestResource {
	return &defaultRestResource{
		mode:         DefaultRestResourceModeCreateAndUpdatePUT,
		resourcePath: resourcePath,
		unmarshaller: unmarshaller,
		client:       client,
	}
}

//NewCreatePOSTUpdatePUTRestResource creates a new REST resource using the provided unmarshaller function to convert the response from the REST API to the corresponding InstanaDataObject. The REST resource is using POST as operation for create and PUT for update
func NewCreatePOSTUpdatePUTRestResource(resourcePath string, unmarshaller JSONUnmarshaller, client RestClient) RestResource {
	return &defaultRestResource{
		mode:         DefaultRestResourceModeCreatePOSTUpdatePUT,
		resourcePath: resourcePath,
		unmarshaller: unmarshaller,
		client:       client,
	}
}

//NewCreatePOSTUpdatePOSTRestResource creates a new REST resource using the provided unmarshaller function to convert the response from the REST API to the corresponding InstanaDataObject. The REST resource is using POST as operation for create and update
func NewCreatePOSTUpdatePOSTRestResource(resourcePath string, unmarshaller JSONUnmarshaller, client RestClient) RestResource {
	return &defaultRestResource{
		mode:         DefaultRestResourceModeCreateAndUpdatePOST,
		resourcePath: resourcePath,
		unmarshaller: unmarshaller,
		client:       client,
	}
}

//DefaultRestResourceMode custom type for create/update behavior of the defaultRestResource
type DefaultRestResourceMode string

type restClientOperation func(InstanaDataObject, string) ([]byte, error)

const (
	//DefaultRestResourceModeCreateAndUpdatePUT constant value for the DefaultRestResourceMode CREATE_PUT_UPDATE_PUT where create and update is implemented as an upsert using HTTP PUT method only
	DefaultRestResourceModeCreateAndUpdatePUT = DefaultRestResourceMode("CREATE_PUT_UPDATE_PUT")
	//DefaultRestResourceModeCreatePOSTUpdatePUT constant value for the DefaultRestResourceMode CREATE_POST_UPDATE_PUT where create is implemented as an HTTP POST method and update is implemented as HTTP PUT method
	DefaultRestResourceModeCreatePOSTUpdatePUT = DefaultRestResourceMode("CREATE_POST_UPDATE_PUT")
	//DefaultRestResourceModeCreateAndUpdatePOST constant value for the DefaultRestResourceMode CREATE_POST_UPDATE_POST where create is implemented as an HTTP POST method and update is implemented as HTTP PUT method
	DefaultRestResourceModeCreateAndUpdatePOST = DefaultRestResourceMode("CREATE_POST_UPDATE_POST")
)

type defaultRestResource struct {
	mode         DefaultRestResourceMode
	resourcePath string
	unmarshaller JSONUnmarshaller
	client       RestClient
}

func (r *defaultRestResource) GetOne(id string) (InstanaDataObject, error) {
	data, err := r.client.GetOne(id, r.resourcePath)
	if err != nil {
		return nil, err
	}
	return r.validateResponseAndConvertToStruct(data)
}

func (r *defaultRestResource) Create(data InstanaDataObject) (InstanaDataObject, error) {
	if r.mode == DefaultRestResourceModeCreateAndUpdatePUT {
		return r.upsert(data, r.client.Put)
	}
	return r.upsert(data, r.client.Post)
}

func (r *defaultRestResource) Update(data InstanaDataObject) (InstanaDataObject, error) {
	if r.mode == DefaultRestResourceModeCreateAndUpdatePOST {
		return r.upsert(data, r.client.PostWithID)
	}
	return r.upsert(data, r.client.Put)
}

func (r *defaultRestResource) upsert(data InstanaDataObject, operation restClientOperation) (InstanaDataObject, error) {
	if err := data.Validate(); err != nil {
		return data, err
	}
	response, err := operation(data, r.resourcePath)
	if err != nil {
		return data, err
	}
	return r.validateResponseAndConvertToStruct(response)
}

func (r *defaultRestResource) validateResponseAndConvertToStruct(data []byte) (InstanaDataObject, error) {
	object, err := r.unmarshaller.Unmarshal(data)
	if err != nil {
		return nil, err
	}
	dataObject, ok := object.(InstanaDataObject)
	if !ok {
		return dataObject, errors.New("unmarshalled object does not implement InstanaDataObject")
	}

	if err := dataObject.Validate(); err != nil {
		return dataObject, err
	}
	return dataObject, nil
}

func (r *defaultRestResource) Delete(data InstanaDataObject) error {
	return r.DeleteByID(data.GetIDForResourcePath())
}

func (r *defaultRestResource) DeleteByID(id string) error {
	return r.client.Delete(id, r.resourcePath)
}
