package bitmap

type Header struct {
	FileType          [2]byte
	FileSize          uint32
	Reserved1         uint16
	Reserved2         uint16
	StartingAddress   uint32
	HeaderSize        uint32
	Width             uint32
	Height            uint32
	ColorPlanes       uint16
	BitsPerPixel      uint16
	CompressionMethod uint32
	ImageSize         uint32
	XPixelsPerMeter   int32
	YPixelsPerMeter   int32
	TotalColors       uint32
	ImportantColors   uint32
}

type Pixel struct {
	b byte
	g byte
	r byte
}

type BMPFile struct {
	head  Header
	image [][]Pixel
}
