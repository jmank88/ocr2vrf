package point_translation

import (
	"github.com/smartcontractkit/ocr2vrf/altbn_128"
	"go.dedis.ch/kyber/v3"
)

type PubKeyTranslation interface {
	TranslateKey(share kyber.Scalar) (kyber.Point, error)

	VerifyTranslation(pk1, pk2 kyber.Point) error

	Name() string

	TargetGroup(sourceGroup kyber.Group) (targetGroup kyber.Group, err error)
}

var TranslatorRegistry = map[string]PubKeyTranslation{

	"translator from AltBN-128 G₁ to AltBN-128 G₂": &PairingTranslation{
		&altbn_128.PairingSuite{},
	},

	"trivial": &TrivialTranslation{},

	"bad translator": &BadTranslator{PairingTranslation{&altbn_128.PairingSuite{}}},
}
