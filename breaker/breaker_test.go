package breaker_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/Oncilla/break-bazel/breaker"
	"github.com/Oncilla/break-bazel/breaker/mock_breaker"
)

func TestSmasherSmash(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSmasher := mock_breaker.NewMockSmasher(ctrl)
	mockSmasher.EXPECT().Smash().DoAndReturn(func() (breaker.Broken, error) {
		return false, errors.New("i don't want to break anything")
	})
	if _, err := (breaker.Breaker{Smasher: mockSmasher}).Break(); err == nil {
		t.Error("Smash should fail")
	}
}
