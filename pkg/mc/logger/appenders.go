
type Appender interface {
	write(msg string) error
}

type FileAppender struct {
}

func (f *FileAppender) write(msg string) error {

}

type ConsoleAppender struct {
}

func (f *ConsoleAppender) write(msg string) error {

}


