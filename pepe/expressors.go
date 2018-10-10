package pepe

import "cryptopepe.io/cryptopepe-svg/builder/look"

// Expressor type
type GeneExpressor func(geneDNA uint32, look *look.PepeLook)

/**
   Gene expressors
   ----------------------------------------
 */

var skinGradient = GradientTable{
  {MustHexColor("#68984E"), 0.0},
  {MustHexColor("#42ad40"), 0.3},
  {MustHexColor("#2c994f"), 0.35},
  {MustHexColor("#50992c"), 0.4},
  {MustHexColor("#20a01c"), 0.5},
  {MustHexColor("#89a830"), 0.55},
  {MustHexColor("#78ab5e"), 0.6},
  {MustHexColor("#508a31"), 0.7},
  {MustHexColor("#68984E"), 0.9},
  {MustHexColor("#6174bf"), 0.91},
  {MustHexColor("#A8549D"), 0.92},
  {MustHexColor("#7063c8"), 0.93},
  {MustHexColor("#449ba3"), 0.94},
  {MustHexColor("#90a343"), 0.95},
  {MustHexColor("#33e0ec"), 0.96},
  {MustHexColor("#7D9C81"), 0.97},
  {MustHexColor("#68984E"), 1.0},
}

func SkinColorExpressor(geneDNA uint32, look *look.PepeLook) {
  bit10Expr.ExpressColor(geneDNA, &look.Skin.Color, &skinGradient)
}

var eyeGradient = GradientTable{
  {MustHexColor("#0B4AA8"), 0.0},
  {MustHexColor("#1C4280"), 0.4},
  {MustHexColor("#4C6652"), 0.7},
  {MustHexColor("#25998B"), 0.8},
  {MustHexColor("#54164B"), 0.82},
  {MustHexColor("#fa0a0a"), 0.83},
  {MustHexColor("#40313F"), 0.84},
  {MustHexColor("#0a22a8"), 0.85},
  {MustHexColor("#5009a8"), 0.86},
  {MustHexColor("#0F3B54"), 0.9},
  {MustHexColor("#0B4AA8"), 1.0},
}

func EyesColorExpressor(geneDNA uint32, look *look.PepeLook) {
  bit10Expr.ExpressColor(geneDNA, &look.Head.Eyes.EyeColor, &eyeGradient)
}

var eyeData = ExpressorData{
  Ranges: []ExpressorRange{
    bit12Expr.Expression("eyes>colored_eyes", 0.0, 0.48, false),
    bit12Expr.Expression("eyes>closed_eyes", 0.48, 0.49, false),
    bit12Expr.Expression("eyes>future_robot_eyes", 0.49, 0.505, false),
    bit12Expr.Expression("eyes>scary_eyes", 0.505, 0.53, false),
    bit12Expr.Expression("eyes>illuminati_eye", 0.53, 0.55, false),
    bit12Expr.Expression("eyes>closed_eyes_2", 0.55, 0.6, false),
    bit12Expr.Expression("eyes>crying_eyes", 0.6, 0.64, false),
    bit12Expr.Expression("eyes>alien_eyes", 0.64, 0.66, false),
    bit12Expr.Expression("eyes>woke_eyes", 0.66, 0.68, false),
    bit12Expr.Expression("eyes>crying_eyes_2", 0.68, 0.71, false),
    bit12Expr.Expression("eyes>ethereum_eyes", 0.71, 0.73, true),
    bit12Expr.Expression("eyes>half_closed_eyes", 0.73, 0.755, false),
    bit12Expr.Expression("eyes>monkas_eye", 0.755, 0.78, false),
    bit12Expr.Expression("eyes>mad_eyes", 0.78, 0.8, true),
    bit12Expr.Expression("eyes>ghandi", 0.8, 0.81, false),
    bit12Expr.Expression("eyes>money_eyes", 0.81, 0.845, false),
    bit12Expr.Expression("eyes>red_eyes", 0.845, 0.89, false),
    bit12Expr.Expression("eyes>small_eyes", 0.89, 0.93, false),
    bit12Expr.Expression("eyes>lol_eyes", 0.93, 0.96, false),
    bit12Expr.Expression("eyes>wide_eyes", 0.96, 0.98, true),
    bit12Expr.Expression("eyes>smug_eyes", 0.98, 1.0, false),
  },
}

func EyesTypeExpressor(geneDNA uint32, look *look.PepeLook) {
  ExpressPropType(geneDNA, &look.Head.Eyes.EyeType, &eyeData)
}

var hairData = ExpressorData{
  Ranges: []ExpressorRange{
    bit12Expr.Expression("none", 0, 0.1, true),
    bit12Expr.Expression("hair>cpep_uasf_hat", 0.1, 0.12, false),
    bit12Expr.Expression("hair>nerd_hair", 0.12, 0.15, false),
    bit12Expr.Expression("hair>viking_hat", 0.15, 0.175, false),
    bit12Expr.Expression("hair>spike_hair", 0.175, 0.19, false),
    bit12Expr.Expression("hair>old_bun", 0.19, 0.21, false),
    bit12Expr.Expression("hair>knife_through_head", 0.21, 0.23, false),
    bit12Expr.Expression("hair>nun", 0.23, 0.24, false),
    bit12Expr.Expression("hair>two_braids", 0.24, 0.255, false),
    bit12Expr.Expression("hair>beanie", 0.255, 0.27, false),
    bit12Expr.Expression("hair>mcaffee", 0.27, 0.28, false),
    bit12Expr.Expression("hair>indian_hat", 0.28, 0.3, false),
    bit12Expr.Expression("hair>elvis_hair", 0.3, 0.33, false),
    bit12Expr.Expression("hair>trump_hair", 0.33, 0.355, false),
    bit12Expr.Expression("hair>bun_beard", 0.355, 0.38, false),
    bit12Expr.Expression("hair>thug_life_cap", 0.38, 0.4, true),
    bit12Expr.Expression("hair>frankenstein", 0.4, 0.41, false),
    bit12Expr.Expression("hair>kim_jong_un_hair", 0.41, 0.42, false),
    bit12Expr.Expression("hair>70s_hair", 0.42, 0.45, false),
    bit12Expr.Expression("hair>wizard_hat", 0.45, 0.455, false),
    bit12Expr.Expression("hair>turkish_hat", 0.455, 0.465, false),
    bit12Expr.Expression("hair>terrorist", 0.465, 0.475, false),
    bit12Expr.Expression("hair>mohawk", 0.475, 0.5, false),
    bit12Expr.Expression("hair>minerhat2", 0.5, 0.52, true),
    bit12Expr.Expression("hair>mining_cap", 0.52, 0.53, true),
    bit12Expr.Expression("hair>bitcoin_cap", 0.53, 0.56, false),
    bit12Expr.Expression("hair>uasf_hat", 0.56, 0.575, false),
    bit12Expr.Expression("hair>pharaoh", 0.575, 0.59, false),
    bit12Expr.Expression("hair>egyptian_hat", 0.59, 0.61, false),
    bit12Expr.Expression("hair>hi_hat", 0.61, 0.62, true),
    bit12Expr.Expression("hair>chaplin", 0.62, 0.63, false),
    bit12Expr.Expression("hair>goku_hair", 0.63, 0.65, false),
    bit12Expr.Expression("hair>wig", 0.65, 0.68, false),
    bit12Expr.Expression("hair>mine_hat", 0.68, 0.71, false),
    bit12Expr.Expression("hair>cowboy_hat", 0.71, 0.74, false),
    bit12Expr.Expression("hair>bandana", 0.74, 0.77, false),
    bit12Expr.Expression("hair>nakamoto", 0.77, 0.79, false),
    bit12Expr.Expression("hair>paper_hat", 0.79, 0.81, false),
    bit12Expr.Expression("hair>rollers", 0.81, 0.825, false),
    bit12Expr.Expression("hair>big_bun", 0.825, 0.85, false),
    bit12Expr.Expression("hair>tinfoil_antenna_hat", 0.85, 0.87, false),
    bit12Expr.Expression("hair>hat", 0.87, 0.9, false),
    bit12Expr.Expression("hair>samurai", 0.9, 0.91, false),
    bit12Expr.Expression("hair>sombrero_hat", 0.91, 0.94, false),
    bit12Expr.Expression("hair>modern_hair", 0.94, 0.985, false),
    bit12Expr.Expression("hair>vitalik_hair", 0.985, 1.0, false),
  },
}

func HeadHairTypeExpressor(geneDNA uint32, look *look.PepeLook) {
  ExpressPropType(geneDNA, &look.Head.Hair.HairType, &hairData)
}

var hatGradient = GradientTable{
  {MustHexColor("#cf1d32"), 0.0},
  {MustHexColor("#d290db"), 0.1},
  {MustHexColor("#3c40b5"), 0.2},
  {MustHexColor("#3ab5b5"), 0.3},
  {MustHexColor("#0acc31"), 0.4},
  {MustHexColor("#dee874"), 0.5},
  {MustHexColor("#a82222"), 0.6},
  {MustHexColor("#533b5c"), 0.7},
  {MustHexColor("#404040"), 0.8},
  {MustHexColor("#323835"), 0.9},
  {MustHexColor("#cf1d32"), 1.0},
}

func HeadHatColorExpressor(geneDNA uint32, look *look.PepeLook) {
  bit5Expr.ExpressColor(geneDNA, &look.Head.Hair.HatColor, &hatGradient)
}

var hat2Gradient = GradientTable{
	{MustHexColor("#cf1d32"), 0.0},
	{MustHexColor("#d290db"), 0.1},
	{MustHexColor("#3c40b5"), 0.2},
	{MustHexColor("#3ab5b5"), 0.3},
	{MustHexColor("#0acc31"), 0.4},
	{MustHexColor("#dee874"), 0.5},
	{MustHexColor("#a82222"), 0.6},
	{MustHexColor("#A071B0"), 0.7},
	{MustHexColor("#7A4146"), 0.8},
	{MustHexColor("#22A060"), 0.9},
	{MustHexColor("#cf1d32"), 1.0},
}

func HeadHatColor2Expressor(geneDNA uint32, look *look.PepeLook) {
  bit5Expr.ExpressColor(geneDNA, &look.Head.Hair.HatColor2, &hat2Gradient)
}

var hairGradient = GradientTable{
  {MustHexColor("#e0eb15"), 0.0},
  {MustHexColor("#bd9e20"), 0.2},
  {MustHexColor("#734C33"), 0.3},

  {MustHexColor("#36381f"), 0.33},
  {MustHexColor("#992e03"), 0.35},
  {MustHexColor("#3b2603"), 0.37},
  {MustHexColor("#463D2A"), 0.55},

  {MustHexColor("#353821"), 0.6},

  {MustHexColor("#0d0d0d"), 0.75},
  {MustHexColor("#292e28"), 0.95},

  {MustHexColor("#2c8ebf"), 0.96},
  {MustHexColor("#e60000"), 0.965},
  {MustHexColor("#f2ff03"), 0.97},
  {MustHexColor("#00ffc4"), 0.975},
  {MustHexColor("#0d00ff"), 0.98},
  {MustHexColor("#ff00dd"), 0.99},
  {MustHexColor("#e0eb15"), 1.0},
}

func HeadHairColorExpressor(geneDNA uint32, look *look.PepeLook) {
  bit6Expr.ExpressColor(geneDNA, &look.Head.Hair.HairColor, &hairGradient)
}

var mouthData = ExpressorData{
  Ranges: []ExpressorRange{
    bit12Expr.Expression("mouth>bad_ass_mouth", 0, 0.04, false),
    bit12Expr.Expression("mouth>braces", 0.04, 0.10, false),
    bit12Expr.Expression("mouth>basic_lips", 0.10, 0.25, true),
    bit12Expr.Expression("mouth>cool_smile", 0.25, 0.3, true),
    bit12Expr.Expression("mouth>drink_wine", 0.3, 0.33, false),
    bit12Expr.Expression("mouth>young_lips", 0.33, 0.39, false),
    bit12Expr.Expression("mouth>crying_lips", 0.39, 0.41, false),
    bit12Expr.Expression("mouth>extreme_sad_lips", 0.41, 0.43, false),
    bit12Expr.Expression("mouth>old_sad_mouth", 0.43, 0.45, false),
    bit12Expr.Expression("mouth>fake_smile", 0.45, 0.55, false),
    bit12Expr.Expression("mouth>happy_lips", 0.55, 0.59, true),
    bit12Expr.Expression("mouth>drink_coffee", 0.59, 0.61, false),
    bit12Expr.Expression("mouth>lips_to_kiss", 0.61, 0.63, false),
    bit12Expr.Expression("mouth>rabbit_teeth", 0.63, 0.645, false),
    bit12Expr.Expression("mouth>smug_lips", 0.645, 0.69, false),
    bit12Expr.Expression("mouth>smug_mustache", 0.69, 0.71, false),
    bit12Expr.Expression("mouth>mad_lips_2", 0.71, 0.725, false),
    bit12Expr.Expression("mouth>mad_mouth", 0.725, 0.74, false),
    bit12Expr.Expression("mouth>lol_mouth", 0.74, 0.76, false),
    bit12Expr.Expression("mouth>open_mouth", 0.76, 0.8, false),
    bit12Expr.Expression("mouth>sad_lips_2", 0.8, 0.825, true),
    bit12Expr.Expression("mouth>feels_birthday", 0.825, 0.84, false),
    bit12Expr.Expression("mouth>smiling_lips", 0.84, 0.88, false),
    bit12Expr.Expression("mouth>surprised_mouth", 0.88, 0.9, false),
    bit12Expr.Expression("mouth>tongue_out", 0.9, 0.92, true),
    bit12Expr.Expression("mouth>what_mouth", 0.92, 0.93, false),
    bit12Expr.Expression("mouth>wide_mouth", 0.93, 0.96, true),
    bit12Expr.Expression("mouth>wtf_mouth", 0.96, 1.0, false),
  },
}

func HeadMouthExpressor(geneDNA uint32, look *look.PepeLook) {
  ExpressPropType(geneDNA, &look.Head.Mouth, &mouthData)

}

var neckData = ExpressorData{
  Ranges: []ExpressorRange{
    bit8Expr.Expression("none", 0, 0.75, true),
    bit8Expr.Expression("neck>choker_necklace", 0.65, 0.75, false),
    bit8Expr.Expression("neck>cowboy_neck", 0.75, 0.83, false),
    bit8Expr.Expression("neck>dollar_necklace", 0.83, 0.91, true),
    bit8Expr.Expression("neck>punk_necklace", 0.91, 1.0, false),
  },
}

func BodyNeckExpressor(geneDNA uint32, look *look.PepeLook) {
  ExpressPropType(geneDNA, &look.Body.Neck, &neckData)
}

var shirtData = ExpressorData{
  Ranges: []ExpressorRange{
    bit8Expr.Expression("shirt>basic_shirt", 0, 0.5, true),
    bit8Expr.Expression("shirt>mexican_poncho", 0.5, 0.55, false),
    bit8Expr.Expression("shirt>pepemon", 0.55, 0.58, false),
    bit8Expr.Expression("shirt>dracula", 0.58, 0.61, false),
    bit8Expr.Expression("shirt>vitalik_shirt", 0.61, 0.64, false),
    bit8Expr.Expression("shirt>basic_shirt", 0.64, 0.75, false),
    bit8Expr.Expression("shirt>darth_pepe", 0.75, 0.755, false),
    bit8Expr.Expression("shirt>camp_shirt", 0.755, 0.86, false),
    bit8Expr.Expression("shirt>kim_jong_un_shirt", 0.86, 0.9, false),
    bit8Expr.Expression("shirt>suit", 0.9, 1.0, false),
  },
}

func BodyShirtTypeExpressor(geneDNA uint32, look *look.PepeLook) {
  ExpressPropType(geneDNA, &look.Body.Shirt.ShirtType, &shirtData)
}

var shirtGradient = GradientTable{
  {MustHexColor("#223249"), 0.0},
  {MustHexColor("#960606"), 0.1},
  {MustHexColor("#ff4000"), 0.13},
  {MustHexColor("#56573c"), 0.17},
  {MustHexColor("#00b32e"), 0.2},
  {MustHexColor("#ffbf00"), 0.3},
  {MustHexColor("#00b387"), 0.4},
  {MustHexColor("#91bd9d"), 0.5},
  {MustHexColor("#233b6f"), 0.6},
  {MustHexColor("#bdd214"), 0.7},
  {MustHexColor("#387277"), 0.8},
  {MustHexColor("#2d3030"), 0.9},
  {MustHexColor("#223249"), 1.0},
}

func BodyShirtColorExpressor(geneDNA uint32, look *look.PepeLook) {
  bit10Expr.ExpressColor(geneDNA, &look.Body.Shirt.ShirtColor, &shirtGradient)
}

var glassesData = ExpressorData{
  Ranges: []ExpressorRange{
    bit8Expr.Expression("none", 0, 0.45, true),
    bit8Expr.Expression("glasses>kim_jong_un_glasses", 0.45, 0.47, false),
    bit8Expr.Expression("glasses>future_glasses", 0.47, 0.5, false),
    bit8Expr.Expression("glasses>vr_set", 0.5, 0.52, false),
    bit8Expr.Expression("glasses>explosion_goggles", 0.52, 0.55, false),
    bit8Expr.Expression("glasses>pirate_hat", 0.55, 0.6, false),
    bit8Expr.Expression("glasses>smart_glasses", 0.6, 0.78, false),
    bit8Expr.Expression("glasses>sunglasses_1", 0.78, 0.85, true),
    bit8Expr.Expression("glasses>sunglasses_2", 0.85, 0.95, false),
    bit8Expr.Expression("glasses>thuglife", 0.95, 1.0, false),
  },
}

func GlassesTypeExpressor(geneDNA uint32, look *look.PepeLook) {
  ExpressPropType(geneDNA, &look.Extra.Glasses.GlassesType, &glassesData)
}

// primary color for glasses is a darker than secondary generally.
var glassesPrimaryGradient = GradientTable{
  {MustHexColor("#290e0e"), 0.0},
  {MustHexColor("#1c0f0f"), 0.1},
  {MustHexColor("#090505"), 0.2},
  {MustHexColor("#0f0f52"), 0.3},
  {MustHexColor("#3a5447"), 0.4},
  {MustHexColor("#333333"), 0.5},
  {MustHexColor("#083810"), 0.6},
  {MustHexColor("#236961"), 0.7},
  {MustHexColor("#3941a8"), 0.8},
  {MustHexColor("#6b344e"), 0.9},
  {MustHexColor("#290e0e"), 1.0},
}

func GlassesPrimaryColorExpressor(geneDNA uint32, look *look.PepeLook) {
  bit10Expr.ExpressColor(geneDNA, &look.Extra.Glasses.PrimaryColor, &glassesPrimaryGradient)
}


var glassesSecondaryGradient = GradientTable{
  {MustHexColor("#9e0142"), 0.0},
  {MustHexColor("#d53e4f"), 0.1},
  {MustHexColor("#f79676"), 0.2},
  {MustHexColor("#e37c14"), 0.3},
  {MustHexColor("#fee090"), 0.4},
  {MustHexColor("#ffffbf"), 0.5},
  {MustHexColor("#dedede"), 0.6},
  {MustHexColor("#4beb36"), 0.7},
  {MustHexColor("#4a4da8"), 0.8},
  {MustHexColor("#842e99"), 0.9},
  {MustHexColor("#9e0142"), 1.0},
}

func GlassesSecondaryColorExpressor(geneDNA uint32, look *look.PepeLook) {
  bit10Expr.ExpressColor(geneDNA, &look.Extra.Glasses.SecondaryColor, &glassesSecondaryGradient)
}

var backgroundGradient = GradientTable{
  {MustHexColor("#f3ddca"), 0.0},
  {MustHexColor("#d3e0fc"), 0.1},
  {MustHexColor("#fafad8"), 0.2},
  {MustHexColor("#cbc9e1"), 0.3},
  {MustHexColor("#e8cccc"), 0.4},
  {MustHexColor("#d8fad7"), 0.5},
  {MustHexColor("#fad5f3"), 0.6},
  {MustHexColor("#b5b7e9"), 0.7},
  {MustHexColor("#fda0ed"), 0.8},
  {MustHexColor("#f9aaaa"), 0.9},
  {MustHexColor("#d9d4d0"), 1.0},
}

func BackgroundColorExpressor(geneDNA uint32, look *look.PepeLook) {
  bit10Expr.ExpressColor(geneDNA, &look.BackgroundColor, &backgroundGradient)
}
