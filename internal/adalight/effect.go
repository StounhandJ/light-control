package adalight

type Effect interface {
	Init(strip *Strip, fps int) (name string)
	Frame(n int) (done bool)
}
