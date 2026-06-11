package figure

import (
	"fmt"
	"reflect"

	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	keyTag   = "fig"
	ignore   = "-"        // Field is omitted
	required = "required" // Value must be presented in config
	nonZero  = "non_zero" // Field must have non zero value. Note: zero value for pointer is nil
)

var (
	ErrNonZeroValue  = errors.New("you must set non zero value to this field")
	ErrRequiredValue = errors.New("you must set the value in field")
	ErrNoHook        = errors.New("no such hook")
	ErrNotValid      = errors.New("not valid value")
)

type Validatable interface {
	// Validate validates the data and returns an error if validation fails.
	Validate() error
}

// Hook signature for custom hooks.
// Takes raw value expected to return target value
type Hook func(value interface{}) (reflect.Value, error)

// Hooks is mapping raw type -> `Hook` instance
type Hooks map[string]Hook

// With accepts hooks to be used for figuring out target from raw values.
// `BaseHooks` will be used implicitly if no hooks are provided
func (f *Figurator) With(hooks ...Hooks) *Figurator {
	merged := f.hooks
	for _, partial := range hooks {
		for key, hook := range partial {
			merged[key] = hook
		}
	}
	f.hooks = merged
	return f
}

// Figurator holds state for chained call
type Figurator struct {
	from   interface{}
	hooks  Hooks
	target interface{}
}

// Out is main entry point for package, used to start figure out chain
func Out(target interface{}) *Figurator {
	hooks := make(Hooks)
	for key, hook := range BaseHooks {
		hooks[key] = hook
	}
	return &Figurator{
		target: target,
		hooks:  hooks,
	}
}

// From takes raw config values to be used in figure out process
func (f *Figurator) From(values map[string]interface{}) *Figurator {
	f.from = values
	return f
}

// FromInterface takes raw config values to be used in figure out process
func (f *Figurator) FromInterface(values interface{}) *Figurator {
	f.from = values
	return f
}

// Please exit point for figure out chain.
// Will modify target partially in case of error
func (f *Figurator) Please() error {
	if f.target == nil {
		return errors.New("can't decode to nil target")
	}

	vv := reflect.ValueOf(f.target)
	if vv.Kind() != reflect.Pointer {
		return fmt.Errorf("can't decode to non-pointer '%v' - use & operator", vv.Type().String())
	}
	if vv.IsNil() && !vv.CanSet() {
		return fmt.Errorf("can't decode to unsettable '%v' - use & operator", vv.Type().String())
	}

	return f.setValue(vv.Elem(), f.from)
}

func (f *Figurator) setValue(vv reflect.Value, from interface{}) error {
	if from == nil {
		return nil
	}

	hook, hasHook := f.hooks[vv.Type().String()]
	if hasHook {
		value, err := hook(from)
		if err != nil {
			return err
		}
		vv.Set(value)
		return nil
	}

	var err error
	switch vv.Kind() {
	case reflect.Bool:
		err = setBool(vv, from)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		err = setInt(vv, from)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		err = setUint(vv, from)
	case reflect.Float32, reflect.Float64:
		err = setFloat(vv, from)
	case reflect.String:
		err = setString(vv, from)
	case reflect.Pointer:
		err = f.setPointer(vv, from)
	case reflect.Array:
		err = f.setArray(vv, from)
	case reflect.Slice:
		err = f.setSlice(vv, from)
	case reflect.Map:
		err = f.setMap(vv, from)
	case reflect.Struct:
		err = f.setStruct(vv, from)
	default:
		return errors.New(fmt.Sprintf("%s types are not supported", vv.Type().String()))
	}

	if val, ok := f.target.(Validatable); ok {
		return val.Validate()
	}

	return err
}

func setBool(vv reflect.Value, from interface{}) error {
	v, err := cast.ToBoolE(from)
	vv.SetBool(v)
	return err
}

func setInt(vv reflect.Value, from interface{}) error {
	v, err := cast.ToInt64E(from)
	vv.SetInt(v)
	return err
}

func setUint(vv reflect.Value, from interface{}) error {
	v, err := cast.ToUint64E(from)
	vv.SetUint(v)
	return err
}

func setFloat(vv reflect.Value, from interface{}) error {
	v, err := cast.ToFloat64E(from)
	vv.SetFloat(v)
	return err
}

func setString(vv reflect.Value, from interface{}) error {
	v, err := cast.ToStringE(from)
	vv.SetString(v)
	return err
}

func (f *Figurator) setPointer(vv reflect.Value, from interface{}) error {
	if err := allocPtrIfNil(&vv); err != nil {
		return err
	}
	return f.setValue(vv.Elem(), from)
}

func (f *Figurator) setArray(vv reflect.Value, from interface{}) error {
	fromValue := reflect.ValueOf(from)
	if fromValue.Kind() != reflect.Array && fromValue.Kind() != reflect.Slice {
		return errors.New(fmt.Sprintf("can't set array from non-array value: expected type %s, actual type: %s",
			vv.Type().String(), fromValue.Type().String()))
	}
	if vv.Len() != fromValue.Len() {
		return errors.New("array length mismatch")
	}

	for i := 0; i < fromValue.Len(); i++ {
		if err := f.setValue(vv.Index(i), fromValue.Index(i).Interface()); err != nil {
			return errors.Wrap(err, fmt.Sprintf("can't set array element %d", i))
		}
	}

	return nil
}

func (f *Figurator) setSlice(vv reflect.Value, from interface{}) error {
	fromValue := reflect.ValueOf(from)
	if fromValue.Kind() != reflect.Array && fromValue.Kind() != reflect.Slice {
		return errors.Wrap(ErrNotValid, fmt.Sprintf("can't set slice from non-array value: expected type %s, actual type: %s",
			vv.Type().String(), fromValue.Type().String()))
	}

	slice := reflect.MakeSlice(vv.Type(), fromValue.Len(), fromValue.Len())
	for i := 0; i < fromValue.Len(); i++ {
		if err := f.setValue(slice.Index(i), fromValue.Index(i).Interface()); err != nil {
			return errors.Wrap(err, fmt.Sprintf("can't set slice element %d", i))
		}
	}

	vv.Set(slice)
	return nil
}

func (f *Figurator) setMap(vv reflect.Value, from interface{}) error {
	if vv.Type().Key().Kind() != reflect.String {
		return errors.Wrap(ErrNotValid, fmt.Sprintf("map key type must be string or its alias, actual is %s", vv.Type().Key().String()))
	}

	fromValue := reflect.ValueOf(from)
	if fromValue.Kind() != reflect.Map {
		return errors.Wrap(ErrNotValid, fmt.Sprintf("can't set map from non-map value: expected type %s, actual type: %s",
			vv.Type().String(), fromValue.Type().String()))
	}

	m := reflect.MakeMap(vv.Type())
	for _, key := range fromValue.MapKeys() {
		val := reflect.New(vv.Type().Elem()).Elem()
		if err := f.setValue(val, fromValue.MapIndex(key).Interface()); err != nil {
			return errors.Wrap(err, fmt.Sprintf("can't set map element %s", key.String()))
		}
		keyStr := reflect.New(vv.Type().Key()).Elem()
		err := setString(keyStr, key.Interface())
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("can't set map key %s", key.String()))
		}
		m.SetMapIndex(keyStr, val)
	}

	vv.Set(m)
	return nil
}

func (f *Figurator) setStruct(vv reflect.Value, from interface{}) error {
	v, ok := from.(map[string]interface{})
	if !ok {
		rawV, ok := from.(map[interface{}]interface{})
		if !ok {
			return errors.Wrap(ErrNotValid, fmt.Sprintf("can't cast to map %s", reflect.TypeOf(from).String()))
		}
		v = make(map[string]interface{})
		for k, val := range rawV {
			v[fmt.Sprintf("%v", k)] = val
		}
	}

	for i := 0; i < vv.NumField(); i++ {
		vf := vv.Field(i)

		tag, err := parseFieldTag(vv.Type().Field(i), keyTag)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("can't parse field tag for field %s", vv.Type().Field(i).Name))
		}
		if tag == nil {
			continue
		}

		value, hasValue := v[tag.Key]

		if hasValue {
			if err := f.setValue(vf, value); err != nil {
				return errors.Wrap(err, fmt.Sprintf("can't set field %s", tag.Key))
			}
		}

		if tag.Required && !hasValue {
			return errors.Wrap(ErrRequiredValue, fmt.Sprintf("required field '%s' is missing", tag.Key))
		}

		if tag.NonZero && isZero(vf) {
			return errors.Wrap(ErrNonZeroValue, fmt.Sprintf("field %s has a zero value", tag.Key))
		}
	}

	return nil
}

func allocPtrIfNil(v *reflect.Value) error {
	if v.Kind() != reflect.Pointer {
		return fmt.Errorf("value is not a pointer: '%v'", v.Type().String())
	}
	isNil := v.IsNil()
	if isNil && !v.CanSet() {
		return fmt.Errorf("unable to allocate pointer for '%v'", v.Type().String())
	}
	if isNil {
		v.Set(reflect.New(v.Type().Elem()))
	}
	return nil
}

// IsZeroer is used to check whether an object is zero to
// determine whether it should be omitted when marshaling
// with the omitempty flag. One notable implementation
// is time.Time.
type IsZeroer interface {
	IsZero() bool
}

func isZero(v reflect.Value) bool {
	kind := v.Kind()
	if z, ok := v.Interface().(IsZeroer); ok {
		if (kind == reflect.Pointer || kind == reflect.Interface) && v.IsNil() {
			return true
		}
		return z.IsZero()
	}
	switch kind {
	case reflect.String:
		return len(v.String()) == 0
	case reflect.Interface, reflect.Pointer:
		return v.IsNil()
	case reflect.Slice:
		return v.Len() == 0
	case reflect.Map:
		return v.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Struct:
		vt := v.Type()
		for i := v.NumField() - 1; i >= 0; i-- {
			if vt.Field(i).PkgPath != "" {
				continue // Private field
			}
			if !isZero(v.Field(i)) {
				return false
			}
		}
		return true
	}
	return false
}
