package mail

type Packet struct {
	FFrom    string
	SSubject string
	TText    string
	TTo      string
}

func (p Packet) From(from string) Packet {
	p.FFrom = from

	return p
}

func (p Packet) Subject(subject string) Packet {
	p.SSubject = subject

	return p
}

func (p Packet) Text(text string) Packet {
	p.TText = text

	return p
}

func (p Packet) To(to string) Packet {
	p.TTo = to

	return p
}
