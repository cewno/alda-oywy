package model

import (
	"fmt"
	"testing"

	_ "alda.io/client/testing"
)

func TestAttributes(t *testing.T) {
	executeScoreUpdateTestCases(
		t,
		scoreUpdateTestCase{
			label: "initial octave",
			updates: []ScoreUpdate{
				PartDeclaration{Names: []string{"piano"}},
			},
			expectations: []scoreUpdateExpectation{
				expectPart("piano", func(part *Part) error {
					if part.Octave != 4 {
						return fmt.Errorf("initial octave is %d, not 4", part.Octave)
					}

					return nil
				}),
			},
		},
		scoreUpdateTestCase{
			label: "set octave",
			updates: []ScoreUpdate{
				PartDeclaration{Names: []string{"piano"}},
				AttributeUpdate{PartUpdate: OctaveSet{OctaveNumber: 2}},
			},
			expectations: []scoreUpdateExpectation{
				expectPart("piano", func(part *Part) error {
					if part.Octave != 2 {
						return fmt.Errorf("octave is %d, not 2", part.Octave)
					}

					return nil
				}),
			},
		},
		scoreUpdateTestCase{
			label: "set octave using lisp",
			updates: []ScoreUpdate{
				PartDeclaration{Names: []string{"piano"}},
				LispList{Elements: []LispForm{
					LispSymbol{Name: "octave"},
					LispNumber{Value: 5},
				}},
			},
			expectations: []scoreUpdateExpectation{
				expectPart("piano", func(part *Part) error {
					if part.Octave != 5 {
						return fmt.Errorf("octave is %d, not 5", part.Octave)
					}

					return nil
				}),
			},
		},
		scoreUpdateTestCase{
			label: "decrement octave",
			updates: []ScoreUpdate{
				PartDeclaration{Names: []string{"piano"}},
				AttributeUpdate{PartUpdate: OctaveSet{OctaveNumber: 2}},
				AttributeUpdate{PartUpdate: OctaveDown{}},
			},
			expectations: []scoreUpdateExpectation{
				expectPart("piano", func(part *Part) error {
					if part.Octave != 1 {
						return fmt.Errorf("octave is %d, not 1", part.Octave)
					}

					return nil
				}),
			},
		},
		scoreUpdateTestCase{
			label: "decrement octave using lisp",
			updates: []ScoreUpdate{
				PartDeclaration{Names: []string{"piano"}},
				AttributeUpdate{PartUpdate: OctaveSet{OctaveNumber: 2}},
				LispList{Elements: []LispForm{
					LispSymbol{Name: "octave"},
					LispQuotedForm{Form: LispSymbol{Name: "down"}},
				}},
			},
			expectations: []scoreUpdateExpectation{
				expectPart("piano", func(part *Part) error {
					if part.Octave != 1 {
						return fmt.Errorf("octave is %d, not 1", part.Octave)
					}

					return nil
				}),
			},
		},
		scoreUpdateTestCase{
			label: "increment octave",
			updates: []ScoreUpdate{
				PartDeclaration{Names: []string{"piano"}},
				AttributeUpdate{PartUpdate: OctaveSet{OctaveNumber: 2}},
				AttributeUpdate{PartUpdate: OctaveUp{}},
			},
			expectations: []scoreUpdateExpectation{
				expectPart("piano", func(part *Part) error {
					if part.Octave != 3 {
						return fmt.Errorf("octave is %d, not 3", part.Octave)
					}

					return nil
				}),
			},
		},
		scoreUpdateTestCase{
			label: "increment octave using lisp",
			updates: []ScoreUpdate{
				PartDeclaration{Names: []string{"piano"}},
				AttributeUpdate{PartUpdate: OctaveSet{OctaveNumber: 2}},
				LispList{Elements: []LispForm{
					LispSymbol{Name: "octave"},
					LispQuotedForm{Form: LispSymbol{Name: "up"}},
				}},
			},
			expectations: []scoreUpdateExpectation{
				expectPart("piano", func(part *Part) error {
					if part.Octave != 3 {
						return fmt.Errorf("octave is %d, not 3", part.Octave)
					}

					return nil
				}),
			},
		},
		scoreUpdateTestCase{
			label: "several octave operations in a row",
			updates: []ScoreUpdate{
				PartDeclaration{Names: []string{"piano"}},
				AttributeUpdate{PartUpdate: OctaveSet{OctaveNumber: 4}},
				AttributeUpdate{PartUpdate: OctaveUp{}},
				AttributeUpdate{PartUpdate: OctaveUp{}},
				AttributeUpdate{PartUpdate: OctaveUp{}},
				AttributeUpdate{PartUpdate: OctaveDown{}},
			},
			expectations: []scoreUpdateExpectation{
				expectPart("piano", func(part *Part) error {
					if part.Octave != 6 {
						return fmt.Errorf("octave is %d, not 6", part.Octave)
					}

					return nil
				}),
			},
		},
		scoreUpdateTestCase{
			label: "initial volume",
			updates: []ScoreUpdate{
				PartDeclaration{Names: []string{"piano"}},
			},
			expectations: []scoreUpdateExpectation{
				expectPart("piano", func(part *Part) error {
					if part.Volume != 1.0 {
						return fmt.Errorf("initial volume is %f, not 1.0", part.Volume)
					}

					return nil
				}),
			},
		},
		scoreUpdateTestCase{
			label: "set volume",
			updates: []ScoreUpdate{
				PartDeclaration{Names: []string{"piano"}},
				AttributeUpdate{PartUpdate: VolumeSet{Volume: 0.85}},
			},
			expectations: []scoreUpdateExpectation{
				expectPart("piano", func(part *Part) error {
					if part.Volume != 0.85 {
						return fmt.Errorf("volume is %f, not 0.85", part.Volume)
					}

					return nil
				}),
			},
		},
		scoreUpdateTestCase{
			label: "set volume using lisp",
			updates: []ScoreUpdate{
				PartDeclaration{Names: []string{"piano"}},
				LispList{Elements: []LispForm{
					LispSymbol{Name: "volume"},
					LispNumber{Value: 82},
				}},
			},
			expectations: []scoreUpdateExpectation{
				expectPart("piano", func(part *Part) error {
					if part.Volume != 0.82 {
						return fmt.Errorf("volume is %f, not 0.82", part.Volume)
					}

					return nil
				}),
			},
		},
		scoreUpdateTestCase{
			label: "initial track volume",
			updates: []ScoreUpdate{
				PartDeclaration{Names: []string{"piano"}},
			},
			expectations: []scoreUpdateExpectation{
				expectPart("piano", func(part *Part) error {
					expected := float32(100.0 / 127)
					if part.TrackVolume != expected {
						return fmt.Errorf(
							"initial track volume is %f, not %f", part.TrackVolume, expected,
						)
					}

					return nil
				}),
			},
		},
		scoreUpdateTestCase{
			label: "set track volume",
			updates: []ScoreUpdate{
				PartDeclaration{Names: []string{"piano"}},
				AttributeUpdate{PartUpdate: TrackVolumeSet{TrackVolume: 0.85}},
			},
			expectations: []scoreUpdateExpectation{
				expectPart("piano", func(part *Part) error {
					if part.TrackVolume != 0.85 {
						return fmt.Errorf("track volume is %f, not 0.85", part.TrackVolume)
					}

					return nil
				}),
			},
		},
		scoreUpdateTestCase{
			label: "set track volume using lisp",
			updates: []ScoreUpdate{
				PartDeclaration{Names: []string{"piano"}},
				LispList{Elements: []LispForm{
					LispSymbol{Name: "track-volume"},
					LispNumber{Value: 82},
				}},
			},
			expectations: []scoreUpdateExpectation{
				expectPart("piano", func(part *Part) error {
					if part.TrackVolume != 0.82 {
						return fmt.Errorf("track volume is %f, not 0.82", part.TrackVolume)
					}

					return nil
				}),
			},
		},
		scoreUpdateTestCase{
			label: "initial panning",
			updates: []ScoreUpdate{
				PartDeclaration{Names: []string{"piano"}},
			},
			expectations: []scoreUpdateExpectation{
				expectPart("piano", func(part *Part) error {
					if part.Panning != 0.5 {
						return fmt.Errorf("initial volume is %f, not 0.5", part.Panning)
					}

					return nil
				}),
			},
		},
		scoreUpdateTestCase{
			label: "set panning",
			updates: []ScoreUpdate{
				PartDeclaration{Names: []string{"piano"}},
				AttributeUpdate{PartUpdate: PanningSet{Panning: 0.85}},
			},
			expectations: []scoreUpdateExpectation{
				expectPart("piano", func(part *Part) error {
					if part.Panning != 0.85 {
						return fmt.Errorf("panning is %f, not 0.85", part.Panning)
					}

					return nil
				}),
			},
		},
		scoreUpdateTestCase{
			label: "set panning using lisp",
			updates: []ScoreUpdate{
				PartDeclaration{Names: []string{"piano"}},
				LispList{Elements: []LispForm{
					LispSymbol{Name: "panning"},
					LispNumber{Value: 82},
				}},
			},
			expectations: []scoreUpdateExpectation{
				expectPart("piano", func(part *Part) error {
					if part.Panning != 0.82 {
						return fmt.Errorf("panning is %f, not 0.82", part.Panning)
					}

					return nil
				}),
			},
		},
		scoreUpdateTestCase{
			label: "initial quantization",
			updates: []ScoreUpdate{
				PartDeclaration{Names: []string{"piano"}},
			},
			expectations: []scoreUpdateExpectation{
				expectPart("piano", func(part *Part) error {
					if part.Quantization != 0.9 {
						return fmt.Errorf(
							"initial quantization is %f, not 0.9", part.Quantization,
						)
					}

					return nil
				}),
			},
		},
		scoreUpdateTestCase{
			label: "set quantization",
			updates: []ScoreUpdate{
				PartDeclaration{Names: []string{"piano"}},
				AttributeUpdate{PartUpdate: QuantizationSet{Quantization: 0.85}},
			},
			expectations: []scoreUpdateExpectation{
				expectPart("piano", func(part *Part) error {
					if part.Quantization != 0.85 {
						return fmt.Errorf("quantization is %f, not 0.85", part.Quantization)
					}

					return nil
				}),
			},
		},
		scoreUpdateTestCase{
			label: "set quantization using lisp",
			updates: []ScoreUpdate{
				PartDeclaration{Names: []string{"piano"}},
				LispList{Elements: []LispForm{
					LispSymbol{Name: "quant"},
					LispNumber{Value: 82},
				}},
			},
			expectations: []scoreUpdateExpectation{
				expectPart("piano", func(part *Part) error {
					if part.Quantization != 0.82 {
						return fmt.Errorf("quantization is %f, not 0.82", part.Quantization)
					}

					return nil
				}),
			},
		},
		scoreUpdateTestCase{
			label: "set quantization using lisp: value > 100",
			updates: []ScoreUpdate{
				PartDeclaration{Names: []string{"piano"}},
				LispList{Elements: []LispForm{
					LispSymbol{Name: "quant"},
					LispNumber{Value: 9001},
				}},
			},
			expectations: []scoreUpdateExpectation{
				expectPart("piano", func(part *Part) error {
					if part.Quantization != 90.01 {
						return fmt.Errorf("quantization is %f, not 90.01", part.Quantization)
					}

					return nil
				}),
			},
		},
	)
}