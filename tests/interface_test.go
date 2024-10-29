package tests

import (
	funcs "github.com/BelehovEgor/go-fuzz-targets/examples"
	"github.com/BelehovEgor/go-fuzz-targets/mocks"
	"go.uber.org/mock/gomock"
	"testing"
)

func FuzzBuyCar(f *testing.F) {
	f.Fuzz(func(t *testing.T, money int, price int, bytes []byte) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		// generate mock by go.uber.org/mock/mockgen
		carMock := mocks.NewMockCar(mockCtrl)

		// setup
		carMock.
			EXPECT().
			Price().       // setup as any? if args needed
			Return(price). // setup as args of regular functions?
			AnyTimes()

		err := funcs.BuyCar(carMock, -123)
		if err == nil {
			// some result processing
			return
		}
	})
}
