type Command interface {
	AddBitToParam(bool)
	Exec(*Runtime)
	FinishReadParam()
}
