package appsettings

// Configuration interface
type Configuration interface {
	GetValue(string) interface{}
}

// ConfigurationImpl encapsulates the app settings after they are loaded
type ConfigurationImpl struct {
	values *appsettingsCollection
}

// GetValue returns configuraiton values
func (c *ConfigurationImpl) GetValue(key string) interface{} {
	return c.values.get(key)
}
