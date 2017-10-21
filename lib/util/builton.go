package util

import (
  "path"
  "strings"
  "reflect"
  "strconv"
  "fmt"
  "errors"
)

// ErrNotAStructPointer self explanatory
var ErrNotAStructPointer = errors.New("given interface is not a struct pointer")

// Shift returns the head and tail of a string slice, given an empty list it will return ("", nil).
func Shift(slice []string) (head string, tail []string) {
  switch len(slice) {
  case 0:
    return "", nil
  case 1:
    return slice[0], nil
  default:
    return slice[0], slice[1:]
  }
}

// ShiftPath returns the head and tail of a string of form "a/b/c/d", in this case "a" and "/b/c/d"
func ShiftPath(p string) (head, tail string) {
    p = path.Clean("/" + p)
    i := strings.Index(p[1:], "/") + 1
    if i <= 0 {
        return p[1:], "/"
    }
    return p[1:i], p[i:]
}

// Override takes a pointer to a struct, a field and a value, and overrides the field of the struct with the value.
func Override(in interface{}, fieldName string, inval interface{}) (error) {
  val := reflect.ValueOf(in)
  if val.Kind() != reflect.Ptr {
    return fmt.Errorf("Override only accepts pointers to structs; got %T", val)
  }
  val = val.Elem()
  if val.Kind() != reflect.Struct {
    return fmt.Errorf("Override only accepts pointers to structs; got %T", val)
  }
  field := val.FieldByName(fieldName)
  if !field.IsValid() {
    return fmt.Errorf("No field by name: %s, in struct", field)
  }
  value := reflect.ValueOf(inval)
  if !value.Type().AssignableTo(field.Type()) {
    return fmt.Errorf("The type of the value (%s) is not assignable to the field(%s)", field.Kind().String(), value.Kind().String())
  }
  field.Set(value)
  return nil
}

// TypeString returns the name of the type.
func TypeString(in interface{}) (string) {
  val := reflect.ValueOf(in)
  if val.Kind() != reflect.Ptr {
    return val.Type().String()
  }
  val = val.Elem()
  return val.Type().String()
}

// FieldType takes a pointer to a struct and a field name, returns the type of the field in the struct.
func FieldType(in interface{}, fieldName string) (reflect.Type, error) {
  val := reflect.ValueOf(in)
  if val.Kind() != reflect.Ptr {
    return nil, fmt.Errorf("FieldType only accepts pointers to structs; got %T", val)
  }
  val = val.Elem()
  if val.Kind() != reflect.Struct {
    return nil, fmt.Errorf("FieldType only accepts pointers to structs; got %T", val)
  }
  field := val.FieldByName(fieldName)
  if !field.IsValid() {
    return nil, fmt.Errorf("No field by name: %s, in struct", field)
  }
  return field.Type(), nil
}

// HasField returns true if given a struct that has a field by the name
// fieldName, returns false otherwise.
func HasField(in interface{}, fieldName string) bool {
  val := reflect.ValueOf(in)
  if val.Kind() != reflect.Ptr {
    return false
  }
  val = val.Elem()
  if val.Kind() != reflect.Struct {
    return false
  }
  field := val.FieldByName(fieldName)
  return field.IsValid()
}

// GetStringField given a struct and a fieldName corresponding to a string value
// returns the string of the field. If field is not a string, will return string
// value of field.
func GetStringField(in interface{}, fieldName string) (string, error) {
  val := reflect.ValueOf(in)
  if val.Kind() != reflect.Ptr {
    return "", ErrNotAStructPointer
  }
  val = val.Elem()
  if val.Kind() != reflect.Struct {
    return "", ErrNotAStructPointer
  }
  field := val.FieldByName(fieldName)
  return field.String(), nil
}

// StringToType will given a string and a type try to return a value of the given type derived from the string.
// not robust, only works for string, bool and int.
func StringToType(value string, tpe reflect.Type) (interface{}, error){
  switch tpe.Kind() {
  case reflect.Bool:
    return strconv.ParseBool(value)
  case reflect.Int:
    return strconv.ParseInt(value, 10, 64)
  case reflect.String:
    return value, nil
  }
  return nil, fmt.Errorf("Type is neither bool, int or string; got %T", tpe)
}
