package dto

type Result struct {
	Success  bool        `json:"success"`
	ErrorMsg string      `json:"errormsg,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Long     int         `json:"long,omitempty"`
}

func (r *Result) Fail(errorMsg string) Result {
	r.Success = false
	r.ErrorMsg = errorMsg
	//	fmt.Println(errorMsg)
	return *r
}

func (r *Result) Ok(data interface{}) Result {
	r.Success = true
	r.Data = data
	return *r
}
