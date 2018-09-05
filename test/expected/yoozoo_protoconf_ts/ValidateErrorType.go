// Code generated by protoapi; DO NOT EDIT.

package yoozoo_protoconf_ts

type ValidateErrorType int

const (
    INVALID_EMAIL ValidateErrorType = 0
    FIELD_REQUIRED ValidateErrorType = 1
)

func (code ValidateErrorType) String() string {
	names := map[ValidateErrorType]string{
        INVALID_EMAIL: "INVALID_EMAIL",
        FIELD_REQUIRED: "FIELD_REQUIRED",
	}

	return names[code]
}

func (code ValidateErrorType) Code() int {
	return (int)(code)
}
