package objects

type Blob struct {
	Content []byte
}

func NewBlob(content []byte) *Blob{
	return &Blob{Content: content}
}

func (b *Blob) Type() ObjectType {
	return BlobType
}

func (b *Blob) Serialize() []byte{
	return b.Content
}

func (b *Blob) Hash() string{
	return  HashContent(BlobType, b.Content)
}