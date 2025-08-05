package surgeon

import (
	"testing"

	"github.com/mdw-go/surgeon/internal/should"
)

func Test(t *testing.T) {
	operation := &Stringectomy{Text: []byte("The quick brown fox jumps over the lazy dog.")}

	should.So(t, operation.Contains("jumps over"), should.BeTrue)
	should.So(t, operation.Contains("JUMPS OVER"), should.BeFalse)

	should.So(t, operation.CutPrefix("The"), should.BeTrue)
	should.So(t, operation.CutPrefix("The"), should.BeFalse)
	should.So(t, string(operation.Text), should.Equal, " quick brown fox jumps over the lazy dog.")

	should.So(t, operation.DiscardNLeft(1), should.BeTrue)
	should.So(t, string(operation.Text), should.Equal, "quick brown fox jumps over the lazy dog.")

	should.So(t, operation.TrimRight("."), should.BeTrue)
	should.So(t, operation.TrimRight("."), should.BeFalse)
	should.So(t, string(operation.Text), should.Equal, "quick brown fox jumps over the lazy dog")

	should.So(t, operation.ExciseAll("o"), should.BeTrue)
	should.So(t, operation.ExciseAll("o"), should.BeFalse)
	should.So(t, string(operation.Text), should.Equal, "quick brwn fx jumps ver the lazy dg")

	should.So(t, operation.DiscardNRight(1), should.BeTrue)
	should.So(t, string(operation.Text), should.Equal, "quick brwn fx jumps ver the lazy d")

	should.So(t, operation.Excise("the ", 1), should.BeTrue)
	should.So(t, operation.Excise("the ", 1), should.BeFalse)
	should.So(t, string(operation.Text), should.Equal, "quick brwn fx jumps ver lazy d")

	should.So(t, operation.TrimLeft("kciuq"), should.BeTrue)
	should.So(t, operation.TrimLeft("kciuq"), should.BeFalse)
	should.So(t, string(operation.Text), should.Equal, " brwn fx jumps ver lazy d")

	should.So(t, operation.CutAfter("lazy"), should.BeTrue)
	should.So(t, operation.CutAfter("lazy"), should.BeFalse)
	should.So(t, string(operation.Text), should.Equal, " brwn fx jumps ver ")

	should.So(t, operation.CutBefore("fx"), should.BeTrue)
	should.So(t, operation.CutBefore("fx"), should.BeFalse)
	should.So(t, string(operation.Text), should.Equal, " jumps ver ")

	should.So(t, operation.CutSuffix("ver "), should.BeTrue)
	should.So(t, operation.CutSuffix("ver "), should.BeFalse)
	should.So(t, string(operation.Text), should.Equal, " jumps ")

	should.So(t, operation.DiscardNLeft(len(" jumps ")), should.BeTrue)
	should.So(t, operation.DiscardNLeft(1), should.BeFalse)
	should.So(t, operation.DiscardNRight(1), should.BeFalse)
	should.So(t, string(operation.Text), should.Equal, "")
}
