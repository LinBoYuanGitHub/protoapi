// Code generated by protoapi; DO NOT EDIT.

package yoozoo_protoconf_ts

// ServiceListRequest
type ServiceListRequest struct {
    Tag_ids []int `json:"tag_ids"`
    Env_id int `json:"env_id"`
    Offset int `json:"offset"`
    Limit int `json:"limit"`
}

func (r ServiceListRequest) Validate() *ValidateError {
    errs := []*FieldError{}
    if len(errs) > 0 {
		return &ValidateError{Errors: errs}
	}
	return nil
}