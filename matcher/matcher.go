package matcher

type Matcher struct {
  params *MatcherParams

  FileList *FileList
}

func New(params *MatcherParams) *Matcher {
  return &Matcher {
    FileList: NewFileList(params),
    params: params }
}
