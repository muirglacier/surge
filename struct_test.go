package surge_test

import (
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/muirglacier/surge/surgeutil"
)

type MyStruct struct {
	MyString string
	MyInt    uint64
	MyFloat  float64
	MyBool   bool
	MySlice  []byte
}

var _ = Describe("Struct", func() {

	numTrials := 100

	ts := []reflect.Type{
		reflect.TypeOf(MyStruct{}),
	}

	Context("when marshaling and then unmarshaling", func() {
		It("should return itself", func() {
			for trial := 0; trial < numTrials; trial++ {
				for _, t := range ts {
					Expect(surgeutil.MarshalUnmarshalCheck(t)).To(Succeed())
				}
			}
		})
	})

	Context("when fuzzing", func() {
		It("should not panic", func() {
			for trial := 0; trial < numTrials; trial++ {
				for _, t := range ts {
					Expect(func() { surgeutil.Fuzz(t) }).ToNot(Panic())
				}
			}
		})
	})

	Context("when marshaling", func() {
		Context("when the buffer is too small", func() {
			It("should return itself", func() {
				for trial := 0; trial < numTrials; trial++ {
					for _, t := range ts {
						Expect(surgeutil.MarshalBufTooSmall(t)).To(Succeed())
					}
				}
			})
		})

		Context("when the remaining memory quota is too small", func() {
			It("should return itself", func() {
				for trial := 0; trial < numTrials; trial++ {
					for _, t := range ts {
						Expect(surgeutil.MarshalRemTooSmall(t)).To(Succeed())
					}
				}
			})
		})
	})

	Context("when unmarshaling", func() {
		Context("when the buffer is too small", func() {
			It("should return itself", func() {
				for trial := 0; trial < numTrials; trial++ {
					for _, t := range ts {
						Expect(surgeutil.UnmarshalBufTooSmall(t)).To(Succeed())
					}
				}
			})
		})

		Context("when the remaining memory quota is too small", func() {
			It("should return itself", func() {
				for trial := 0; trial < numTrials; trial++ {
					for _, t := range ts {
						Expect(surgeutil.UnmarshalRemTooSmall(t)).To(Succeed())
					}
				}
			})
		})
	})
})
