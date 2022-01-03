package utilities

import (
	"fmt"

	"github.com/onsi/gomega"
)

func AssertContain(Str string, Substr string) (err error) {

	gomega.RegisterFailHandler(func(message string, _ ...int) {
		panic(message)
	})
	gomega.Expect(Str).Should(gomega.ContainSubstring(Substr))

	defer FailHandler(&err)
	return err
}

func AssertEqual(Expected string, Actual string) (err error) {
	gomega.RegisterFailHandler(func(message string, _ ...int) {
		panic(message)
	})
	gomega.Expect(Expected).Should(gomega.Equal(Actual))
	defer FailHandler(&err)
	return err
}

func AssertEqualInt(Expected int, Actual int) (err error) {
	gomega.RegisterFailHandler(func(message string, _ ...int) {
		panic(message)
	})
	gomega.Expect(Expected).Should(gomega.Equal(Actual))
	defer FailHandler(&err)
	return err
}

func FailHandler(err *error) {
	if r := recover(); r != nil {
		*err = fmt.Errorf("%s", r)
	}
}
