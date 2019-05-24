 
package cryptofs

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
    // this is encryption key found in part 2
	key := "95511870061fb3a2899aa6b2dc9838aa"

	pwd, err := os.Getwd()
	handleErr(err, t)


	f, err := os.Open("original.txt")
	handleErr(err, t)
	defer func() {
		f.Close()
	}()

	fstat, err := f.Stat()
	handleErr(err, t)

	fileInfo := &File{fstat, "txt", pwd + `/original.txt`}

	var buf []byte
	buffer := bytes.NewBuffer(buf)


	content, err := ioutil.ReadAll(buffer)
	handleErr(err, t)


	f2, err := os.Open("data")
	handleErr(err, t)
	defer func() {
		f2.Close()
	}()

	fstat, err = f2.Stat()
	handleErr(err, t)

	fileInfo = &File{fstat, "txt", pwd + `/data`}

	buffer.Reset()
	err = fileInfo.Decrypt(key, buffer)

	content, err = ioutil.ReadAll(buffer)
	handleErr(err, t)
    f3, err := os.Open("decrypted.docx")
    err = ioutil.WriteFile("decrypted.docx", content, 0600)
    handleErr(err, t)
    f3.Close()

}


func handleErr(err error, t *testing.T) {
	if err != nil {
		t.Fatalf("An error ocurred: %s", err)
	}
}
