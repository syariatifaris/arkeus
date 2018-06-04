package validation

import (
	"regexp"

	"fmt"
	"reflect"
)

//Create New Validation Instance
func NewValidator() Validation {
	return &validationImpl{}
}

//Validation implementation
type Validation interface {
	//Check if an object is required
	Required(obj interface{}) *Result
	//Check if an object has met minimum value
	Min(n int, min int) *Result
	//Check if an object has met minimum value
	Min64(n int64, min int64) *Result
	//Check if an object has met maximum value
	Max(n int, max int) *Result
	//Check if an object has met range value
	Range(n, min, max int) *Result
	//Min Size of String
	MinSize(obj interface{}, min int) *Result
	//Max Size of String
	MaxSize(obj interface{}, max int) *Result
	//Length of string
	Length(obj interface{}, n int) *Result
	//Match a regular expression
	Match(str string, regex *regexp.Regexp) *Result
	//Check if the email format is correct
	Email(str string) *Result
	//Check if the phone format is correct
	Phone(str string) *Result
	//Check if the URL is valid
	URL(str string) *Result
	//Equals
	Equals(obj interface{}, compareObj interface{}) *Result
	//Validate
	Validate() (bool, []*Error)
}

//Validation implementation
type validationImpl struct {
	Errors []*Error
}

//Validation Error
type Error struct {
	Message, Key string
}

//Validation Result
type Result struct {
	Error *Error
	Ok    bool
}

//Message Set The Custom Message
func (r *Result) SetErrorMessage(message string) *Result {
	if r.Error != nil {
		r.Error.Message = message
	}

	return r
}

//Required Validation
func (v *validationImpl) Required(obj interface{}) *Result {
	return v.apply(Required{}, obj)
}

//Min Validation
func (v *validationImpl) Min(n int, min int) *Result {
	return v.apply(Min{Min: min}, n)
}

func (v *validationImpl) Min64(n int64, min int64) *Result {
	return v.apply(Min64{Min: min}, n)
}

//Max Validation
func (v *validationImpl) Max(n int, max int) *Result {
	return v.apply(Max{Max: max}, n)
}

//Range Validation
func (v *validationImpl) Range(n, min, max int) *Result {
	return v.apply(Range{Min{min}, Max{max}}, n)
}

//MinSize Validation
func (v *validationImpl) MinSize(obj interface{}, min int) *Result {
	return v.apply(MinSize{min}, obj)
}

//MaxSize Validation
func (v *validationImpl) MaxSize(obj interface{}, max int) *Result {
	return v.apply(MaxSize{max}, obj)
}

//Length Validation
func (v *validationImpl) Length(obj interface{}, n int) *Result {
	return v.apply(Length{n}, obj)
}

//Match Validation
func (v *validationImpl) Match(str string, regex *regexp.Regexp) *Result {
	return v.apply(Match{regex}, str)
}

//Email Validation
func (v *validationImpl) Email(str string) *Result {
	return v.apply(Email{Match{emailPattern}}, str)
}

//Email Validation
func (v *validationImpl) Phone(str string) *Result {
	return v.apply(Email{Match{phonePattern}}, str)
}

//URL Validation
func (v *validationImpl) URL(str string) *Result {
	return v.apply(URL{}, str)
}

//URL Validation
func (v *validationImpl) Equals(obj interface{}, compareObj interface{}) *Result {
	return v.apply(Equals{compareObj}, obj)
}

//Apply validation check validation part
func (v *validationImpl) apply(chk Validator, obj interface{}) *Result {
	if chk.IsFulfilled(obj) {
		return &Result{Ok: true}
	}

	objName := reflect.ValueOf(obj)

	err := &Error{
		Message: chk.GetDefaultMessage(),
		Key:     fmt.Sprintf("%+v", objName),
	}

	v.Errors = append(v.Errors, err)
	return &Result{
		Ok:    false,
		Error: err,
	}
}

//Validate All Validator
func (v validationImpl) Validate() (bool, []*Error) {
	return len(v.Errors) == 0, v.Errors
}
