// Code generated by protoapi:go; DO NOT EDIT.

package calcsvrmain

// AddError
type AddError struct {
	Req   *AddReq `json:"req"`
	Error string  `json:"error"`
}

func (r *AddError) GetReq() *AddReq {
	if r == nil {
		var zeroVal *AddReq
		return zeroVal
	}
	return r.Req
}

func (r *AddError) GetError() string {
	if r == nil {
		var zeroVal string
		return zeroVal
	}
	return r.Error
}
