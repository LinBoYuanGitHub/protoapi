// Code generated by protoapi; DO NOT EDIT.

package yoozoo_protoconf_ts

// GenericError
type GenericError struct {
    Message string `json:"message"`
}

func (r GenericError) Validate() *ValidateError {
    errs := []*FieldError{}
    if len(errs) > 0 {
		return &ValidateError{Errors: errs}
	}
	return nil
}
