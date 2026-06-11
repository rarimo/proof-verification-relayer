package figure

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3"
	"reflect"

	"math/big"

	"net/url"

	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

var (
	// BaseHooks set of default hooks for common types
	BaseHooks = Hooks{
		"string": func(value interface{}) (reflect.Value, error) {
			result, err := cast.ToStringE(value)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to parse string")
			}
			return reflect.ValueOf(result), nil
		},
		"*string": func(value interface{}) (reflect.Value, error) {
			result, err := cast.ToStringE(value)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to parse string")
			}
			return reflect.ValueOf(&result), nil
		},
		"[]int64": func(value interface{}) (reflect.Value, error) {
			var a []int64

			switch v := value.(type) {
			case []int64:
				return reflect.ValueOf(value), nil
			case []int:
				for _, intValue := range v {
					a = append(a, int64(intValue))
				}
				return reflect.ValueOf(a), nil
			case []interface{}:
				for i, u := range v {
					int64Value, err := cast.ToInt64E(u)
					if err != nil {
						return reflect.Value{}, errors.Errorf("failed to cast slice element number %d: %#v of type %T into int64", i, value, value)
					}
					a = append(a, int64Value)
				}
				return reflect.ValueOf(a), nil
			case interface{}:
				int64Value, err := cast.ToInt64E(value)
				if err != nil {
					return reflect.Value{}, errors.Errorf("failed to cast %#v of type %T to int64", value, value)
				}
				return reflect.ValueOf([]int64{int64Value}), nil
			default:
				return reflect.Value{}, errors.Errorf("failed to cast %#v of type %T to []int64", value, value)
			}
		},
		"[]string": func(value interface{}) (reflect.Value, error) {
			result, err := cast.ToStringSliceE(value)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to parse []string")
			}
			return reflect.ValueOf(result), nil
		},
		"int": func(value interface{}) (reflect.Value, error) {
			result, err := cast.ToIntE(value)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to parse int")
			}
			return reflect.ValueOf(result), nil
		},
		"int32": func(value interface{}) (reflect.Value, error) {
			result, err := cast.ToInt32E(value)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to parse int32")
			}
			return reflect.ValueOf(result), nil
		},
		"int64": func(value interface{}) (reflect.Value, error) {
			result, err := cast.ToInt64E(value)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to parse int64")
			}
			return reflect.ValueOf(result), nil
		},
		"uint": func(value interface{}) (reflect.Value, error) {
			result, err := cast.ToUintE(value)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to parse uint")
			}
			return reflect.ValueOf(result), nil
		},
		"uint32": func(value interface{}) (reflect.Value, error) {
			result, err := cast.ToUint32E(value)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to parse uint32")
			}
			return reflect.ValueOf(result), nil
		},
		"uint64": func(value interface{}) (reflect.Value, error) {
			result, err := cast.ToUint64E(value)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to parse uint64")
			}
			return reflect.ValueOf(result), nil
		},
		"float64": func(value interface{}) (reflect.Value, error) {
			result, err := cast.ToFloat64E(value)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to parse float64")
			}
			return reflect.ValueOf(result), nil
		},
		"bool": func(value interface{}) (reflect.Value, error) {
			result, err := cast.ToBoolE(value)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to parse bool")
			}
			return reflect.ValueOf(result), nil
		},
		"*bool": func(value interface{}) (reflect.Value, error) {
			result, err := cast.ToBoolE(value)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to parse bool")
			}
			return reflect.ValueOf(&result), nil
		},
		"time.Time": func(value interface{}) (reflect.Value, error) {
			result, err := cast.ToTimeE(value)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to parse time")
			}
			return reflect.ValueOf(result), nil
		},
		"*time.Time": func(value interface{}) (reflect.Value, error) {
			result, err := cast.ToTimeE(value)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to parse time pointer")
			}
			return reflect.ValueOf(&result), nil
		},
		"time.Duration": func(value interface{}) (reflect.Value, error) {
			result, err := cast.ToDurationE(value)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to parse duration")
			}
			return reflect.ValueOf(result), nil
		},
		"*time.Duration": func(value interface{}) (reflect.Value, error) {
			if value == nil {
				return reflect.ValueOf(nil), nil
			}
			result, err := cast.ToDurationE(value)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to parse duration")
			}
			return reflect.ValueOf(&result), nil
		},
		"*big.Int": func(value interface{}) (reflect.Value, error) {
			switch v := value.(type) {
			case string:
				i, ok := new(big.Int).SetString(v, 10)
				if !ok {
					return reflect.Value{}, errors.New("failed to parse")
				}
				return reflect.ValueOf(i), nil
			case int:
				return reflect.ValueOf(big.NewInt(int64(v))), nil
			default:
				return reflect.Value{}, errors.Errorf("unsupported conversion from %T", value)
			}
		},
		"logan.Level": func(value interface{}) (reflect.Value, error) {
			switch v := value.(type) {
			case string:
				lvl, err := logan.ParseLevel(v)
				if err != nil {
					return reflect.Value{}, errors.Wrap(err, "failed to parse log level")
				}
				return reflect.ValueOf(lvl), nil
			default:
				return reflect.Value{}, errors.Errorf("unsupported conversion from %T", value)
			}
		},
		"*uint64": func(value interface{}) (reflect.Value, error) {
			switch v := value.(type) {
			case string:
				puint, err := cast.ToUint64E(v)
				if err != nil {
					return reflect.Value{}, errors.New("failed to parse")
				}
				return reflect.ValueOf(&puint), nil
			default:
				return reflect.Value{}, errors.Errorf("unsupported conversion from %T", value)
			}
		},
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
				return reflect.Value{}, errors.Errorf("unsupported conversion from %T", value)
			}
		},
		"json.RawMessage": func(value interface{}) (reflect.Value, error) {
			if value == nil {
				return reflect.Value{}, nil
			}

			var params map[string]interface{}
			switch s := value.(type) {
			case map[interface{}]interface{}:
				params = make(map[string]interface{})
				for key, value := range s {
					params[key.(string)] = value
				}
			case map[string]interface{}:
				params = make(map[string]interface{})
				for key, value := range s {
					params[key] = value
				}
			default:
				return reflect.Value{}, errors.New("unexpected type while figure []json.RawMessage")
			}

			result, err := json.Marshal(params)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to parse json.RawMessage")
			}
			return reflect.ValueOf(json.RawMessage(result)), nil
		},
		"map[string]interface {}": func(value interface{}) (reflect.Value, error) {
			if value == nil {
				return reflect.Value{}, nil
			}

			var params map[string]interface{}
			switch s := value.(type) {
			case map[interface{}]interface{}:
				params = make(map[string]interface{})
				for key, value := range s {
					params[key.(string)] = value
				}
			case map[string]interface{}:
				params = s
			default:
				return reflect.Value{}, errors.New("unexpected type while figure map[string]interface{}")
			}

			return reflect.ValueOf(params), nil
		},
	}
)

// Merge does not modify any Hooks, only produces new Hooks.
// If duplicated keys - the value from the last Hooks with such key will be taken.
func Merge(manyHooks ...Hooks) Hooks {
	if len(manyHooks) == 1 {
		return manyHooks[0]
	}

	merged := Hooks{}

	for _, hooks := range manyHooks {
		for key, hook := range hooks {
			merged[key] = hook
		}
	}

	return merged
}
