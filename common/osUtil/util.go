package osUtil

import "os"

func GetStaticPath() string {
	root, _ := os.Getwd()
	static := root + "/static"
	return static
}

func MakeDirAll(dir string) error {
	return os.MkdirAll(dir, 0755)
}
