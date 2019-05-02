// Code generated by protoapi:go; DO NOT EDIT.

package calcsvrmain_without

// CommonError
type CommonError struct {
	GenericError  *GenericError  `json:"genericError"`
	AuthError     *AuthError     `json:"authError"`
	ValidateError *ValidateError `json:"validateError"`
	BindError     *BindError     `json:"bindError"`
}

func (r *CommonError) GetGenericError() *GenericError {
	if r == nil {
		var zeroVal *GenericError
		return zeroVal
	}
	return r.GenericError
}

func (r *CommonError) GetAuthError() *AuthError {
	if r == nil {
		var zeroVal *AuthError
		return zeroVal
	}
	return r.AuthError
}

func (r *CommonError) GetValidateError() *ValidateError {
	if r == nil {
		var zeroVal *ValidateError
		return zeroVal
	}
	return r.ValidateError
}

func (r *CommonError) GetBindError() *BindError {
	if r == nil {
		var zeroVal *BindError
		return zeroVal
	}
	return r.BindError
}