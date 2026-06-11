# Figure

Library to parse interface{}, map[string]interface{}, etc. to Go structs. Mostly used to parse configs.

# Usage

```go
package main

import "gitlab.com/distributed_lab/figure/v3"

type Config struct {
    Name string `fig:"name,required"`
    Age  int    `fig:"age"`
}

func main() {
	// parse somehow your config to variable values
	
    config := Config{}
    err := figure.
        Out(&config).// variable to parse to
        From(values).// map[string]interface{} to parse from
        Please() // parse
    if err != nil {
        panic(err)
    }
}
```

## Parsing slices and primitive types

If your raw config []interface{} or interface{} you can use `FromInterface` function

## Setting custom hooks

If you want to enable parsing custom types with custom rules, you can use `WithHooks` function

```go
func main() {
    config := Config{}
    err := figure.
        Out(&config).
        From(values).
        WithHooks(figure.Hooks{
            "*url.URL": func(value interface{}) (reflect.Value, error) {
                switch v := value.(type) {
                case string:
                    u, err := url.Parse(v)
                    if err != nil {
                        return reflect.Value{}, errors.Wrap(err, "failed to parse url")
                    }
                    return reflect.ValueOf(u), nil
                case nil:
                    return reflect.ValueOf(nil), nil
                default:
                    return reflect.Value{}, fmt.Errorf("unsupported conversion from %T", value)
                },
            },
        }).
        Please()
    if err != nil {
        panic(err)
    }
}
```