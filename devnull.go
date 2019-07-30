package devnull

type reader struct{
}

func (*reader) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i += 1 {
		b[i] = 0
	}
	return len(b), nil
}

var Reader *reader

func init() {
  Reader = new(reader)
}
