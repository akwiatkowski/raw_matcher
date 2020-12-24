package matcher

type Matcher struct {
  FileList *FileList
}

func New() *Matcher {
  return &Matcher{FileList: NewFileList()}
}
