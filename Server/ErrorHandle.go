package Server

type ErrorHandle struct{}

func (e ErrorHandle) CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
