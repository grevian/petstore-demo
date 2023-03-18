/*
 * Swagger Petstore - OpenAPI 3.0
 *
 * This is a sample Pet Store Server based on the OpenAPI 3.0 specification.  You can find out more about Swagger at [http://swagger.io](http://swagger.io). In the third iteration of the pet store, we've switched to the design first approach! You can now help us improve the API whether it's by making changes to the definition itself or to the code. That way, with time, we can improve the API in general, and expose some of the new features in OAS3.  Some useful links: - [The Pet Store repository](https://github.com/swagger-api/swagger-petstore) - [The source API definition for the Pet Store](https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml)
 *
 * API version: 1.0.17
 * Contact: apiteam@swagger.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Pet struct {
	Id int64 `json:"id,omitempty"`

	Name string `json:"name"`

	Category Category `json:"category,omitempty"`

	PhotoUrls []string `json:"photoUrls"`

	Tags []Tag `json:"tags,omitempty"`

	// pet status in the store
	Status string `json:"status,omitempty"`
}

// AssertPetRequired checks if the required fields are not zero-ed
func AssertPetRequired(obj Pet) error {
	elements := map[string]interface{}{
		"name":      obj.Name,
		"photoUrls": obj.PhotoUrls,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertCategoryRequired(obj.Category); err != nil {
		return err
	}
	for _, el := range obj.Tags {
		if err := AssertTagRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecursePetRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Pet (e.g. [][]Pet), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePetRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPet, ok := obj.(Pet)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPetRequired(aPet)
	})
}
