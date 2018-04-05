# tools-config

Load and object from a JSON configuration file.

This process is very forgiving. If the JSON is valid the response err will be null.

If there are no matching properties then the config data will be unchanged!

If your config object implements 'Validate(filename string) error' then it will be called.

## Usage Load:
```go
type Config struct {
  Port        int
  Timeout     int64
}
....
config := Config{
  Timeout:    200,
  Port:       1080}
err := jsonconfig.LoadJson(configFileName, &config)
```
  
Note: Do **NOT** forget the '&' on the config parameter.

## Usage Store
You can also update the config file with: 

```go
err := jsonconfig.StoreJson("aFileName.json", config)
```

## Usage String
You can also get the JSON String value: 

```go
json, err := jsonconfig.StringJson(config)
```

## Validation
If the config object implements the following method:

```go
Validate(filename string) error
```

Then it WILL be called after the object is loaded. 

This is your opportunity to validate the contents of the configuration object returning an error if required. 
