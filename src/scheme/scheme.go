package scheme

type Scheme interface {
	GetSchemeID() string
}

type dataScheme struct {
}

func (sc *dataScheme) GetSchemeID() string {
	return "data:"
}

func Data() *dataScheme {
	return &dataScheme{}
}

type mediastreamScheme struct {
}

func (sc *mediastreamScheme) GetSchemeID() string {
	return "mediastream:"
}

func Mediastream() *mediastreamScheme {
	return &mediastreamScheme{}
}

type blobScheme struct {
}

func (sc *blobScheme) GetSchemeID() string {
	return "blob:"
}

func Blob() *blobScheme {
	return &blobScheme{}
}

type fileSystemScheme struct {
}

func (sc *fileSystemScheme) GetSchemeID() string {
	return "filesystem:"
}

func FileSystem() *fileSystemScheme {
	return &fileSystemScheme{}
}
