package appsettings

// Builder interface
type Builder interface {
	AddEnvironmentVariables() *Builder
	AddJSONFile() *Builder
	AddInMemoryCollection(map[string]interface{}) *Builder
	Build() *Provider
}

// BuilderImpl implements the Builder interface
type BuilderImpl struct {
	providers []Provider
}

// AddEnvironmentVariables reads system environment variables
func (b *BuilderImpl) AddEnvironmentVariables() *BuilderImpl {
	b.providers = append(b.providers, &environmentVariablesProvider{})
	return b
}

// AddInMemoryCollection loads from a map
func (b *BuilderImpl) AddInMemoryCollection(m map[string]interface{}) *BuilderImpl {
	b.providers = append(b.providers, &inMemoryProvider{data: m})
	return b
}

// AddJSONFile configuration loader
func (b *BuilderImpl) AddJSONFile(path string) *BuilderImpl {
	b.providers = append(b.providers, &jsonProvider{path: path})
	return b
}

// Build the configuration
func (b *BuilderImpl) Build() *ConfigurationImpl {
	mergedData := newAppsettingsCollection()
	for _, p := range b.providers {
		data := p.load()
		mergedData.merge(data)
	}
	return &ConfigurationImpl{values: mergedData}
}

// NewConfigurationBuilder constructor for BuilderImpl
func NewConfigurationBuilder() *BuilderImpl {
	return &BuilderImpl{}
}
