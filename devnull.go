package devnull

type reader struct{}

func (reader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 0
	}
	return len(b), nil
}

var Reader *reader

func Read(b []byte) (n int, err error) {
	return Reader.Read(b)
}

var Writer *writer

type writer struct{}

func (writer) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func Write(p []byte) (n int, err error) {
	return Writer.Write(p)
}

func init() {
	Reader = new(reader)
	Writer = new(writer)
}
