package kernelmodule

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboy/normalizer"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Encode(kc *kernel.KernelCmd, config *razboy.Config) (*kernel.KernelCmd, error) {
	sEnc := normalizer.Encode(kc.GetStr())
	kc.SetBody(sEnc)

	return kc, nil
}

func Decode(kc *kernel.KernelCmd, config *razboy.Config) (*kernel.KernelCmd, error) {
	sDec, err := normalizer.Decode(kc.GetStr())
	kc.SetBody(sDec)

	return kc, err
}