# go-appsettings
App settings for Golang applications

## Install

If you have Go installed and configured (i.e. with `$GOPATH/bin` in your `$PATH`):

```
go get -u github.com/bkot88/go-appsettings
```

## Usage
```
builder := appsettings.NewConfigurationBuilder()
builder.AddEnvironmentVariables()
builder.AddJSONFile("appsettings.local.json")
builder.AddInMemoryCollection(map[string]interface{}{"hello:world": "world"})
configuration := builder.Build()

log.Println(configuration.GetValue("HOME"))
log.Println(configuration.GetValue("hello:world"))
log.Println(configuration.GetValue("hello__world"))
log.Println(configuration.GetValue("connectionString"))
```
