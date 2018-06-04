package validation

import (
	"fmt"
	"reflect"
	"regexp"
	"time"
	"unicode/utf8"
)

var (
	//The Email Pattern Regex
	emailPattern = regexp.MustCompile("^[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?$")
	//The Domain Patter Regex
	domainPattern = regexp.MustCompile(`^(([a-zA-Z0-9-\p{L}]{1,63}\.)?(xn--)?[a-zA-Z0-9\p{L}]+(-[a-zA-Z0-9\p{L}]+)*\.)+[a-zA-Z\p{L}]{2,63}$`)
	//The URL Pattern Regex
	urlPattern = regexp.MustCompile(`^((((https?|ftps?|gopher|telnet|nntp)://)|(mailto:|news:))(%[0-9A-Fa-f]{2}|[-()_.!~*';/?:@#&=+$,A-Za-z0-9\p{L}])+)([).!';/?:,][[:blank:]])?$`)
	//The Phone Pattern Reges
	phonePattern = regexp.MustCompile("^\\+{0,1}0{0,1}62[0-9]+$")
)

//Validator Contract
type Validator interface {
	//Check if an object is fulfilled
	IsFulfilled(interface{}) bool
	//Get the validator default message
	GetDefaultMessage() string
}

//Max Validator
type Max struct {
	Max int
}

//IsFulfilled Check if a max integer length is fulfilled
func (m Max) IsFulfilled(obj interface{}) bool {
	num, ok := obj.(int)
	if ok {
		return num <= m.Max
	}
	return false
}

//GetDefaultMessage Get Max Validation Default Message
func (m Max) GetDefaultMessage() string {
	return fmt.Sprintln("Maximum value is", m.Max)
}

//Validator Min
type Min struct {
	Min int
}

//IsFulfilled Check if the value does not exceed the min value
func (m Min) IsFulfilled(obj interface{}) bool {
	num, ok := obj.(int)
	if ok {
		return num >= m.Min
	}
	return false
}

//GetDefaultMessage Get the default message
func (m Min) GetDefaultMessage() string {
	return fmt.Sprintln("Minimum value is", m.Min)
}

//Validator Min
type Min64 struct {
	Min int64
}

//IsFulfilled Check if the value does not exceed the min value
func (m Min64) IsFulfilled(obj interface{}) bool {
	num, ok := obj.(int64)
	if ok {
		return num >= m.Min
	}
	return false
}

//GetDefaultMessage Get the default message
func (m Min64) GetDefaultMessage() string {
	return fmt.Sprintln("Minimum value is", m.Min)
}

// Range requires an integer to be within Min, Max inclusive.
type Range struct {
	Min
	Max
}

//IsFulfilled Check if an integer between max and min indicator
func (r Range) IsFulfilled(obj interface{}) bool {
	return r.Min.IsFulfilled(obj) && r.Max.IsFulfilled(obj)
}

//GetDefaultMessage Get the default message
func (r Range) GetDefaultMessage() string {
	return fmt.Sprintln("Range is", r.Min.Min, "to", r.Max.Max)
}

//Required Validation
type Required struct {
}

//IsFulfilled Check if required validation is fulfilled
func (Required) IsFulfilled(obj interface{}) bool {
	if obj == nil {
		return false
	}

	v := reflect.ValueOf(obj)
	switch v.Kind() {
	case reflect.String:
		if str, ok := obj.(string); ok {
			return utf8.RuneCountInString(str) > 0
		}
	case reflect.Bool:
		if b, ok := obj.(bool); ok {
			return b
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i := v.Int()
		return i != 0
	case reflect.Slice:
		return v.Len() > 0
	case reflect.Struct:
		if v.Type() == reflect.TypeOf(time.Time{}) {
			if t, ok := obj.(time.Time); ok {
				return !t.IsZero()
			}
		}
	}
	return true
}

//GetDefaultMessage Get Required Default Message
func (Required) GetDefaultMessage() string {
	return "Required"
}

//MinSize Validator
type MinSize struct {
	Min int
}

//IsFulfilled Check if minimum size validation is fulfilled
func (m MinSize) IsFulfilled(obj interface{}) bool {
	if str, ok := obj.(string); ok {
		return utf8.RuneCountInString(str) >= m.Min
	}
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Slice {
		return v.Len() >= m.Min
	}
	return false
}

//Get the default message
func (m MinSize) GetDefaultMessage() string {
	return fmt.Sprintln("Minimum size is", m.Min)
}

// MaxSize requires an array or string to be at most a given length.
type MaxSize struct {
	Max int
}

//IsFulfilled Check if maximum size validation is fulfilled
func (m MaxSize) IsFulfilled(obj interface{}) bool {
	if str, ok := obj.(string); ok {
		return utf8.RuneCountInString(str) <= m.Max
	}
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Slice {
		return v.Len() <= m.Max
	}
	return false
}

//Get the default message
func (m MaxSize) GetDefaultMessage() string {
	return fmt.Sprintln("Maximum size is", m.Max)
}

// Length requires an array or string to be exactly a given length.
type Length struct {
	N int
}

//IsFulfilled Check if length of a string is fulfilled
func (s Length) IsFulfilled(obj interface{}) bool {
	if str, ok := obj.(string); ok {
		return utf8.RuneCountInString(str) == s.N
	}
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Slice {
		return v.Len() == s.N
	}
	return false
}

//GetDefaultMessage Get the Default Message
func (s Length) GetDefaultMessage() string {
	return fmt.Sprintln("Required length is", s.N)
}

// Match requires a string to match a given regex.
type Match struct {
	Regexp *regexp.Regexp
}

//IsFulfilled Check if a string match with the regular expression
func (m Match) IsFulfilled(obj interface{}) bool {
	str := obj.(string)
	return m.Regexp.MatchString(str)
}

//Get the default message
func (m Match) GetDefaultMessage() string {
	return fmt.Sprintln("Must match", m.Regexp)
}

//Email Validator
type Email struct {
	Match
}

//GetDeafaultMessage Get the default message
func (e Email) GetDefaultMessage() string {
	return fmt.Sprintln("Must be a valid email address")
}

// Requires a Domain string to be exactly
type Domain struct {
	Regexp *regexp.Regexp
}

//IsFulfilled Check if a domain validation is fulfilled
func (d Domain) IsFulfilled(obj interface{}) bool {
	if str, ok := obj.(string); ok {
		l := len(str)
		//can't exceed 253 chars.
		if l > 253 {
			return false
		}

		//first and last char must be alphanumeric
		if str[l-1] == 46 || str[0] == 46 {
			return false
		}

		return domainPattern.MatchString(str)
	}

	return false
}

//GetDefaultMessage Get the default message
func (d Domain) GetDefaultMessage() string {
	return fmt.Sprintln("Must be a vaild domain address")
}

//URL Validation
type URL struct {
	Domain
}

//IsFulfilled Check if an URL validation is fulfilled
func (u URL) IsFulfilled(obj interface{}) bool {
	if str, ok := obj.(string); ok {
		// TODO : Required lot of testing
		return urlPattern.MatchString(str)
	}

	return false
}

//GetDefaultMessage Get the default message
func (u URL) GetDefaultMessage() string {
	return fmt.Sprintln("Must be a vaild URL address")
}

//Equals Validation
type Equals struct {
	CompareObj interface{}
}

//IsFulfilled Checks if the object is fulfilled
func (e Equals) IsFulfilled(obj interface{}) bool {
	return reflect.DeepEqual(obj, e.CompareObj)
}

//GetDefaultMessage returns the default message
func (e Equals) GetDefaultMessage() string {
	return fmt.Sprintln("Must be equal")
}
