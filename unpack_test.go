
package unpack

import . "path/filepath"
import "testing"
import "log"
import "os"

func TestUnpack(t *testing.T){
  wd, _ := os.Getwd()
  path := Join(wd, "tmp")
  tests := Join(wd, "testdata")
  os.RemoveAll(path)
  os.MkdirAll(path, 0777)

  {
    test := Join(tests, "archive.tar")
    src, _ := os.Open(test)
    dst := Join(path, "archive")
    err := UnpackTarball(src, dst, 0)
    if err != nil {
      log.Fatal(err)
    }
  }
  {
    test := Join(tests, "git.tar")
    src, _ := os.Open(test)
    dst := Join(path, "git")
    err := UnpackTarball(src, dst, 0)
    if err != nil {
      log.Fatal(err)
    }
  }

  os.RemoveAll(path)
}
