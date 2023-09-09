package scheme

type Scheme string

const (
	Data        Scheme = "data:"
	Mediastream Scheme = "mediastream:"
	Blob        Scheme = "blob:"
	Filesystem  Scheme = "filesystem:"
)
