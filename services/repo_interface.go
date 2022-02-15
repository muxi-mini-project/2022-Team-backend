package services

type RepoInterface interface {
	Push(filename, content string) (string, string, string)
	Del(filepath, sha string) string
}
