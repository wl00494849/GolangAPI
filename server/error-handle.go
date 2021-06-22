package server

type ErrorHandle struct{}

var er = new(ErrorHandle)

func (e ErrorHandle) CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
