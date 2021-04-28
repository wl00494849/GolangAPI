package Server

type Err struct{}

func (e *Err) checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
