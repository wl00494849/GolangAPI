package Model

type Sys_Code struct {
	SC_type   string
	SC_Code   string
	SC_Parent string
	SC_Level  int
	SC_Desc   string
	SC_Order  int
}

type Sys_Setting struct {
	SS_Key   string
	SS_Value string
	SS_Type  string
	SS_Desc  string
	SS_Order int
}
