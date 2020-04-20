package appsettings

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

// Provider interface
type Provider interface {
	load() *appsettingsCollection
}

type appsettingsCollection struct {
	values map[string]interface{}
}

func newAppsettingsCollection() *appsettingsCollection {
	return &appsettingsCollection{values: make(map[string]interface{})}
}

func (a *appsettingsCollection) set(key string, value interface{}) {
	/*
		The : separator doesn't work with environment variable hierarchical
		keys on all platforms. __, the double underscore, is.

		Automatically replaced by a :
	*/
	newKey := strings.ReplaceAll(key, "__", ":")
	a.values[newKey] = value
}

func (a *appsettingsCollection) get(key string) interface{} {
	/*
		The : separator doesn't work with environment variable hierarchical
		keys on all platforms. __, the double underscore, is.

		Automatically replaced by a :
	*/
	newKey := strings.ReplaceAll(key, "__", ":")
	value := a.values[newKey]
	return value
}

func (a *appsettingsCollection) mergeMap(m map[string]interface{}) {
	for k, v := range m {
		a.values[k] = v
	}
}

func (a *appsettingsCollection) merge(other *appsettingsCollection) {
	a.mergeMap(other.values)
}

type environmentVariablesProvider struct{}

func (p *environmentVariablesProvider) load() *appsettingsCollection {
	result := newAppsettingsCollection()

	// loop through all environment variablee
	for _, e := range os.Environ() {

		// split them by '='
		pair := strings.SplitN(e, "=", 2)
		key, value := pair[0], pair[1]

		// copy values using set for safe keeping
		result.set(key, value)
	}

	return result
}

type inMemoryProvider struct {
	data map[string]interface{}
}

func (p *inMemoryProvider) load() *appsettingsCollection {
	result := newAppsettingsCollection()

	// copy values using set for safe keeping
	for k, v := range p.data {
		result.set(k, v)
	}

	return result
}

type jsonProvider struct {
	path string
}

func (p *jsonProvider) load() *appsettingsCollection {
	result := newAppsettingsCollection()

	// read json file
	jsonFile, err := os.Open(p.path)

	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var jsonData map[string]interface{}
	json.Unmarshal([]byte(byteValue), &jsonData)

	// copy values using set for safe keeping
	for k, v := range jsonData {
		result.set(k, v)
	}

	return result
}
