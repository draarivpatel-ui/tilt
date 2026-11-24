package yaml

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	noSep = `this: is
some: yaml`
	endSep = `iLoveYaml: meh
betterThanXml: true
---`
	startSep = `---
someMap:
    stuff: yes
    things: also yes
`
)

func TestConcatYamlNoSeparators(t *testing.T) {
	expected := `---
someMap:
    stuff: yes
    things: also yes
---
this: is
some: yaml`
	assert.Equal(t, expected, ConcatYAML(startSep, noSep))
}

func TestConcatYamlBothSeparators(t *testing.T) {
	expected := `iLoveYaml: meh
betterThanXml: true
---
someMap:
    stuff: yes
    things: also yes`
	assert.Equal(t, expected, ConcatYAML(endSep, startSep))
}

func TestConcatYamlEndSep(t *testing.T) {
	expected := `iLoveYaml: meh
betterThanXml: true
---
this: is
some: yaml`
	assert.Equal(t, expected, ConcatYAML(endSep, noSep))
}

func TestConcatYamlStartSep(t *testing.T) {
	expected := `---
someMap:
    stuff: yes
    things: also yes
---
someMap:
    stuff: yes
    things: also yes`
	assert.Equal(t, expected, ConcatYAML(startSep, startSep))
}

func TestConcatYamlDoesNotTreatSeparatorPrefixAsMarker(t *testing.T) {
	scalar := `---not-a-marker`
	expected := `this: is
some: yaml
---
---not-a-marker`
	assert.Equal(t, expected, ConcatYAML(noSep, scalar))
}

func TestConcatYamlDoesNotTreatSeparatorSuffixAsMarker(t *testing.T) {
	scalar := `value: foo---`
	expected := `value: foo---
---
this: is
some: yaml`
	assert.Equal(t, expected, ConcatYAML(scalar, noSep))
}

func TestConcatYamlSeparatorComment(t *testing.T) {
	endComment := `iLoveYaml: meh
--- # end of document`
	startComment := `--- # start of document
someMap:
    stuff: yes`
	expected := `iLoveYaml: meh
--- # start of document
someMap:
    stuff: yes`
	assert.Equal(t, expected, ConcatYAML(endComment, startComment))
}

func TestConcatManyYamls(t *testing.T) {
	expected := `---
someMap:
    stuff: yes
    things: also yes
---
this: is
some: yaml
---
iLoveYaml: meh
betterThanXml: true
---`

	assert.Equal(t, expected, ConcatYAML(startSep, noSep, endSep))
}

func TestNoopConcatYaml(t *testing.T) {
	assert.Equal(t, "", ConcatYAML())

	assert.Equal(t, noSep, ConcatYAML(noSep))
}
