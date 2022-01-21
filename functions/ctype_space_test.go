package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCtypeSpace(t *testing.T) {
	assert.Equal(t, true, CtypeSpace(" \n\r"))
	assert.Equal(t, false, CtypeSpace(" \n\r\\r"))

	// Irregular arguments
	assert.Equal(t, false, CtypeSpace(""))

	// Check each chars
	assert.Equal(t, false, CtypeSpace(string([]byte{0})))
	assert.Equal(t, false, CtypeSpace(string([]byte{1})))
	assert.Equal(t, false, CtypeSpace(string([]byte{2})))
	assert.Equal(t, false, CtypeSpace(string([]byte{3})))
	assert.Equal(t, false, CtypeSpace(string([]byte{4})))
	assert.Equal(t, false, CtypeSpace(string([]byte{5})))
	assert.Equal(t, false, CtypeSpace(string([]byte{6})))
	assert.Equal(t, false, CtypeSpace(string([]byte{7})))
	assert.Equal(t, false, CtypeSpace(string([]byte{8})))
	assert.Equal(t, true, CtypeSpace(string([]byte{9})))
	assert.Equal(t, true, CtypeSpace(string([]byte{10})))
	assert.Equal(t, true, CtypeSpace(string([]byte{11})))
	assert.Equal(t, true, CtypeSpace(string([]byte{12})))
	assert.Equal(t, true, CtypeSpace(string([]byte{13})))
	assert.Equal(t, false, CtypeSpace(string([]byte{14})))
	assert.Equal(t, false, CtypeSpace(string([]byte{15})))
	assert.Equal(t, false, CtypeSpace(string([]byte{16})))
	assert.Equal(t, false, CtypeSpace(string([]byte{17})))
	assert.Equal(t, false, CtypeSpace(string([]byte{18})))
	assert.Equal(t, false, CtypeSpace(string([]byte{19})))
	assert.Equal(t, false, CtypeSpace(string([]byte{20})))
	assert.Equal(t, false, CtypeSpace(string([]byte{21})))
	assert.Equal(t, false, CtypeSpace(string([]byte{22})))
	assert.Equal(t, false, CtypeSpace(string([]byte{23})))
	assert.Equal(t, false, CtypeSpace(string([]byte{24})))
	assert.Equal(t, false, CtypeSpace(string([]byte{25})))
	assert.Equal(t, false, CtypeSpace(string([]byte{26})))
	assert.Equal(t, false, CtypeSpace(string([]byte{27})))
	assert.Equal(t, false, CtypeSpace(string([]byte{28})))
	assert.Equal(t, false, CtypeSpace(string([]byte{29})))
	assert.Equal(t, false, CtypeSpace(string([]byte{30})))
	assert.Equal(t, false, CtypeSpace(string([]byte{31})))
	assert.Equal(t, true, CtypeSpace(string([]byte{32})))
	assert.Equal(t, false, CtypeSpace(string([]byte{33})))
	assert.Equal(t, false, CtypeSpace(string([]byte{34})))
	assert.Equal(t, false, CtypeSpace(string([]byte{35})))
	assert.Equal(t, false, CtypeSpace(string([]byte{36})))
	assert.Equal(t, false, CtypeSpace(string([]byte{37})))
	assert.Equal(t, false, CtypeSpace(string([]byte{38})))
	assert.Equal(t, false, CtypeSpace(string([]byte{39})))
	assert.Equal(t, false, CtypeSpace(string([]byte{40})))
	assert.Equal(t, false, CtypeSpace(string([]byte{41})))
	assert.Equal(t, false, CtypeSpace(string([]byte{42})))
	assert.Equal(t, false, CtypeSpace(string([]byte{43})))
	assert.Equal(t, false, CtypeSpace(string([]byte{44})))
	assert.Equal(t, false, CtypeSpace(string([]byte{45})))
	assert.Equal(t, false, CtypeSpace(string([]byte{46})))
	assert.Equal(t, false, CtypeSpace(string([]byte{47})))
	assert.Equal(t, false, CtypeSpace(string([]byte{48})))
	assert.Equal(t, false, CtypeSpace(string([]byte{49})))
	assert.Equal(t, false, CtypeSpace(string([]byte{50})))
	assert.Equal(t, false, CtypeSpace(string([]byte{51})))
	assert.Equal(t, false, CtypeSpace(string([]byte{52})))
	assert.Equal(t, false, CtypeSpace(string([]byte{53})))
	assert.Equal(t, false, CtypeSpace(string([]byte{54})))
	assert.Equal(t, false, CtypeSpace(string([]byte{55})))
	assert.Equal(t, false, CtypeSpace(string([]byte{56})))
	assert.Equal(t, false, CtypeSpace(string([]byte{57})))
	assert.Equal(t, false, CtypeSpace(string([]byte{58})))
	assert.Equal(t, false, CtypeSpace(string([]byte{59})))
	assert.Equal(t, false, CtypeSpace(string([]byte{60})))
	assert.Equal(t, false, CtypeSpace(string([]byte{61})))
	assert.Equal(t, false, CtypeSpace(string([]byte{62})))
	assert.Equal(t, false, CtypeSpace(string([]byte{63})))
	assert.Equal(t, false, CtypeSpace(string([]byte{64})))
	assert.Equal(t, false, CtypeSpace(string([]byte{65})))
	assert.Equal(t, false, CtypeSpace(string([]byte{66})))
	assert.Equal(t, false, CtypeSpace(string([]byte{67})))
	assert.Equal(t, false, CtypeSpace(string([]byte{68})))
	assert.Equal(t, false, CtypeSpace(string([]byte{69})))
	assert.Equal(t, false, CtypeSpace(string([]byte{70})))
	assert.Equal(t, false, CtypeSpace(string([]byte{71})))
	assert.Equal(t, false, CtypeSpace(string([]byte{72})))
	assert.Equal(t, false, CtypeSpace(string([]byte{73})))
	assert.Equal(t, false, CtypeSpace(string([]byte{74})))
	assert.Equal(t, false, CtypeSpace(string([]byte{75})))
	assert.Equal(t, false, CtypeSpace(string([]byte{76})))
	assert.Equal(t, false, CtypeSpace(string([]byte{77})))
	assert.Equal(t, false, CtypeSpace(string([]byte{78})))
	assert.Equal(t, false, CtypeSpace(string([]byte{79})))
	assert.Equal(t, false, CtypeSpace(string([]byte{80})))
	assert.Equal(t, false, CtypeSpace(string([]byte{81})))
	assert.Equal(t, false, CtypeSpace(string([]byte{82})))
	assert.Equal(t, false, CtypeSpace(string([]byte{83})))
	assert.Equal(t, false, CtypeSpace(string([]byte{84})))
	assert.Equal(t, false, CtypeSpace(string([]byte{85})))
	assert.Equal(t, false, CtypeSpace(string([]byte{86})))
	assert.Equal(t, false, CtypeSpace(string([]byte{87})))
	assert.Equal(t, false, CtypeSpace(string([]byte{88})))
	assert.Equal(t, false, CtypeSpace(string([]byte{89})))
	assert.Equal(t, false, CtypeSpace(string([]byte{90})))
	assert.Equal(t, false, CtypeSpace(string([]byte{91})))
	assert.Equal(t, false, CtypeSpace(string([]byte{92})))
	assert.Equal(t, false, CtypeSpace(string([]byte{93})))
	assert.Equal(t, false, CtypeSpace(string([]byte{94})))
	assert.Equal(t, false, CtypeSpace(string([]byte{95})))
	assert.Equal(t, false, CtypeSpace(string([]byte{96})))
	assert.Equal(t, false, CtypeSpace(string([]byte{97})))
	assert.Equal(t, false, CtypeSpace(string([]byte{98})))
	assert.Equal(t, false, CtypeSpace(string([]byte{99})))
	assert.Equal(t, false, CtypeSpace(string([]byte{100})))
	assert.Equal(t, false, CtypeSpace(string([]byte{101})))
	assert.Equal(t, false, CtypeSpace(string([]byte{102})))
	assert.Equal(t, false, CtypeSpace(string([]byte{103})))
	assert.Equal(t, false, CtypeSpace(string([]byte{104})))
	assert.Equal(t, false, CtypeSpace(string([]byte{105})))
	assert.Equal(t, false, CtypeSpace(string([]byte{106})))
	assert.Equal(t, false, CtypeSpace(string([]byte{107})))
	assert.Equal(t, false, CtypeSpace(string([]byte{108})))
	assert.Equal(t, false, CtypeSpace(string([]byte{109})))
	assert.Equal(t, false, CtypeSpace(string([]byte{110})))
	assert.Equal(t, false, CtypeSpace(string([]byte{111})))
	assert.Equal(t, false, CtypeSpace(string([]byte{112})))
	assert.Equal(t, false, CtypeSpace(string([]byte{113})))
	assert.Equal(t, false, CtypeSpace(string([]byte{114})))
	assert.Equal(t, false, CtypeSpace(string([]byte{115})))
	assert.Equal(t, false, CtypeSpace(string([]byte{116})))
	assert.Equal(t, false, CtypeSpace(string([]byte{117})))
	assert.Equal(t, false, CtypeSpace(string([]byte{118})))
	assert.Equal(t, false, CtypeSpace(string([]byte{119})))
	assert.Equal(t, false, CtypeSpace(string([]byte{120})))
	assert.Equal(t, false, CtypeSpace(string([]byte{121})))
	assert.Equal(t, false, CtypeSpace(string([]byte{122})))
	assert.Equal(t, false, CtypeSpace(string([]byte{123})))
	assert.Equal(t, false, CtypeSpace(string([]byte{124})))
	assert.Equal(t, false, CtypeSpace(string([]byte{125})))
	assert.Equal(t, false, CtypeSpace(string([]byte{126})))
	assert.Equal(t, false, CtypeSpace(string([]byte{127})))
	assert.Equal(t, false, CtypeSpace(string([]byte{128})))
	assert.Equal(t, false, CtypeSpace(string([]byte{129})))
	assert.Equal(t, false, CtypeSpace(string([]byte{130})))
	assert.Equal(t, false, CtypeSpace(string([]byte{131})))
	assert.Equal(t, false, CtypeSpace(string([]byte{132})))
	assert.Equal(t, false, CtypeSpace(string([]byte{133})))
	assert.Equal(t, false, CtypeSpace(string([]byte{134})))
	assert.Equal(t, false, CtypeSpace(string([]byte{135})))
	assert.Equal(t, false, CtypeSpace(string([]byte{136})))
	assert.Equal(t, false, CtypeSpace(string([]byte{137})))
	assert.Equal(t, false, CtypeSpace(string([]byte{138})))
	assert.Equal(t, false, CtypeSpace(string([]byte{139})))
	assert.Equal(t, false, CtypeSpace(string([]byte{140})))
	assert.Equal(t, false, CtypeSpace(string([]byte{141})))
	assert.Equal(t, false, CtypeSpace(string([]byte{142})))
	assert.Equal(t, false, CtypeSpace(string([]byte{143})))
	assert.Equal(t, false, CtypeSpace(string([]byte{144})))
	assert.Equal(t, false, CtypeSpace(string([]byte{145})))
	assert.Equal(t, false, CtypeSpace(string([]byte{146})))
	assert.Equal(t, false, CtypeSpace(string([]byte{147})))
	assert.Equal(t, false, CtypeSpace(string([]byte{148})))
	assert.Equal(t, false, CtypeSpace(string([]byte{149})))
	assert.Equal(t, false, CtypeSpace(string([]byte{150})))
	assert.Equal(t, false, CtypeSpace(string([]byte{151})))
	assert.Equal(t, false, CtypeSpace(string([]byte{152})))
	assert.Equal(t, false, CtypeSpace(string([]byte{153})))
	assert.Equal(t, false, CtypeSpace(string([]byte{154})))
	assert.Equal(t, false, CtypeSpace(string([]byte{155})))
	assert.Equal(t, false, CtypeSpace(string([]byte{156})))
	assert.Equal(t, false, CtypeSpace(string([]byte{157})))
	assert.Equal(t, false, CtypeSpace(string([]byte{158})))
	assert.Equal(t, false, CtypeSpace(string([]byte{159})))
	assert.Equal(t, false, CtypeSpace(string([]byte{160})))
	assert.Equal(t, false, CtypeSpace(string([]byte{161})))
	assert.Equal(t, false, CtypeSpace(string([]byte{162})))
	assert.Equal(t, false, CtypeSpace(string([]byte{163})))
	assert.Equal(t, false, CtypeSpace(string([]byte{164})))
	assert.Equal(t, false, CtypeSpace(string([]byte{165})))
	assert.Equal(t, false, CtypeSpace(string([]byte{166})))
	assert.Equal(t, false, CtypeSpace(string([]byte{167})))
	assert.Equal(t, false, CtypeSpace(string([]byte{168})))
	assert.Equal(t, false, CtypeSpace(string([]byte{169})))
	assert.Equal(t, false, CtypeSpace(string([]byte{170})))
	assert.Equal(t, false, CtypeSpace(string([]byte{171})))
	assert.Equal(t, false, CtypeSpace(string([]byte{172})))
	assert.Equal(t, false, CtypeSpace(string([]byte{173})))
	assert.Equal(t, false, CtypeSpace(string([]byte{174})))
	assert.Equal(t, false, CtypeSpace(string([]byte{175})))
	assert.Equal(t, false, CtypeSpace(string([]byte{176})))
	assert.Equal(t, false, CtypeSpace(string([]byte{177})))
	assert.Equal(t, false, CtypeSpace(string([]byte{178})))
	assert.Equal(t, false, CtypeSpace(string([]byte{179})))
	assert.Equal(t, false, CtypeSpace(string([]byte{180})))
	assert.Equal(t, false, CtypeSpace(string([]byte{181})))
	assert.Equal(t, false, CtypeSpace(string([]byte{182})))
	assert.Equal(t, false, CtypeSpace(string([]byte{183})))
	assert.Equal(t, false, CtypeSpace(string([]byte{184})))
	assert.Equal(t, false, CtypeSpace(string([]byte{185})))
	assert.Equal(t, false, CtypeSpace(string([]byte{186})))
	assert.Equal(t, false, CtypeSpace(string([]byte{187})))
	assert.Equal(t, false, CtypeSpace(string([]byte{188})))
	assert.Equal(t, false, CtypeSpace(string([]byte{189})))
	assert.Equal(t, false, CtypeSpace(string([]byte{190})))
	assert.Equal(t, false, CtypeSpace(string([]byte{191})))
	assert.Equal(t, false, CtypeSpace(string([]byte{192})))
	assert.Equal(t, false, CtypeSpace(string([]byte{193})))
	assert.Equal(t, false, CtypeSpace(string([]byte{194})))
	assert.Equal(t, false, CtypeSpace(string([]byte{195})))
	assert.Equal(t, false, CtypeSpace(string([]byte{196})))
	assert.Equal(t, false, CtypeSpace(string([]byte{197})))
	assert.Equal(t, false, CtypeSpace(string([]byte{198})))
	assert.Equal(t, false, CtypeSpace(string([]byte{199})))
	assert.Equal(t, false, CtypeSpace(string([]byte{200})))
	assert.Equal(t, false, CtypeSpace(string([]byte{201})))
	assert.Equal(t, false, CtypeSpace(string([]byte{202})))
	assert.Equal(t, false, CtypeSpace(string([]byte{203})))
	assert.Equal(t, false, CtypeSpace(string([]byte{204})))
	assert.Equal(t, false, CtypeSpace(string([]byte{205})))
	assert.Equal(t, false, CtypeSpace(string([]byte{206})))
	assert.Equal(t, false, CtypeSpace(string([]byte{207})))
	assert.Equal(t, false, CtypeSpace(string([]byte{208})))
	assert.Equal(t, false, CtypeSpace(string([]byte{209})))
	assert.Equal(t, false, CtypeSpace(string([]byte{210})))
	assert.Equal(t, false, CtypeSpace(string([]byte{211})))
	assert.Equal(t, false, CtypeSpace(string([]byte{212})))
	assert.Equal(t, false, CtypeSpace(string([]byte{213})))
	assert.Equal(t, false, CtypeSpace(string([]byte{214})))
	assert.Equal(t, false, CtypeSpace(string([]byte{215})))
	assert.Equal(t, false, CtypeSpace(string([]byte{216})))
	assert.Equal(t, false, CtypeSpace(string([]byte{217})))
	assert.Equal(t, false, CtypeSpace(string([]byte{218})))
	assert.Equal(t, false, CtypeSpace(string([]byte{219})))
	assert.Equal(t, false, CtypeSpace(string([]byte{220})))
	assert.Equal(t, false, CtypeSpace(string([]byte{221})))
	assert.Equal(t, false, CtypeSpace(string([]byte{222})))
	assert.Equal(t, false, CtypeSpace(string([]byte{223})))
	assert.Equal(t, false, CtypeSpace(string([]byte{224})))
	assert.Equal(t, false, CtypeSpace(string([]byte{225})))
	assert.Equal(t, false, CtypeSpace(string([]byte{226})))
	assert.Equal(t, false, CtypeSpace(string([]byte{227})))
	assert.Equal(t, false, CtypeSpace(string([]byte{228})))
	assert.Equal(t, false, CtypeSpace(string([]byte{229})))
	assert.Equal(t, false, CtypeSpace(string([]byte{230})))
	assert.Equal(t, false, CtypeSpace(string([]byte{231})))
	assert.Equal(t, false, CtypeSpace(string([]byte{232})))
	assert.Equal(t, false, CtypeSpace(string([]byte{233})))
	assert.Equal(t, false, CtypeSpace(string([]byte{234})))
	assert.Equal(t, false, CtypeSpace(string([]byte{235})))
	assert.Equal(t, false, CtypeSpace(string([]byte{236})))
	assert.Equal(t, false, CtypeSpace(string([]byte{237})))
	assert.Equal(t, false, CtypeSpace(string([]byte{238})))
	assert.Equal(t, false, CtypeSpace(string([]byte{239})))
	assert.Equal(t, false, CtypeSpace(string([]byte{240})))
	assert.Equal(t, false, CtypeSpace(string([]byte{241})))
	assert.Equal(t, false, CtypeSpace(string([]byte{242})))
	assert.Equal(t, false, CtypeSpace(string([]byte{243})))
	assert.Equal(t, false, CtypeSpace(string([]byte{244})))
	assert.Equal(t, false, CtypeSpace(string([]byte{245})))
	assert.Equal(t, false, CtypeSpace(string([]byte{246})))
	assert.Equal(t, false, CtypeSpace(string([]byte{247})))
	assert.Equal(t, false, CtypeSpace(string([]byte{248})))
	assert.Equal(t, false, CtypeSpace(string([]byte{249})))
	assert.Equal(t, false, CtypeSpace(string([]byte{250})))
	assert.Equal(t, false, CtypeSpace(string([]byte{251})))
	assert.Equal(t, false, CtypeSpace(string([]byte{252})))
	assert.Equal(t, false, CtypeSpace(string([]byte{253})))
	assert.Equal(t, false, CtypeSpace(string([]byte{254})))
}