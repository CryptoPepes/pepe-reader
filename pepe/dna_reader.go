package pepe

import (
	"cryptopepe.io/cryptopepe-svg/builder/look"
	"math/big"
)

// 2 sets; 1 from mother, 1 from father.
// each side is 256 bits; 4x 64 bits
type PepeDNA [2]*big.Int

type ChromosomeIndex uint8

type Locus struct {
	Start uint
	Len uint
}


var indexedGenes = map[ChromosomeIndex]map[Locus]GeneExpressor{
	0: {
		Locus{0, 10}:  SkinColorExpressor,
		Locus{10, 10}: EyesColorExpressor,
		Locus{20, 12}: EyesTypeExpressor,
		Locus{32, 12}: HeadHairTypeExpressor,
		Locus{44, 5}:  HeadHatColorExpressor,
		Locus{49, 5}:  HeadHatColor2Expressor,
		Locus{54, 6}:  HeadHairColorExpressor,
		Locus{60, 12}: HeadMouthExpressor,
		//128 - 72 = 56 bits unused in this chromosome
	},
	1: {
		Locus{0, 8}:   BodyNeckExpressor,
		Locus{8, 8}:   BodyShirtTypeExpressor,
		Locus{16, 10}: BodyShirtColorExpressor,
		Locus{26, 8}:  GlassesTypeExpressor,
		Locus{34, 10}: GlassesPrimaryColorExpressor,
		Locus{44, 10}: GlassesSecondaryColorExpressor,
		Locus{54, 10}: GlassesSecondaryColorExpressor,
		Locus{64, 10}: BackgroundColorExpressor,
		//128 - 74 = 54 bits unused in this chromosome
	},
}

const chromosomeSize = 128

func (dna *PepeDNA) ParsePepeDNA() *look.PepeLook {
	pepeLook := new(look.PepeLook)


	//read DNA, updating pepeLook
	for chromIndex, locii := range indexedGenes {

		for locus, geneExpressor := range locii {
			// The locii definitions are split up into chromosomes, but the data representation is just one bitstring
			// Calculate the offset of this chromosome, to get the DNA parts that we need
			offset := uint(chromIndex) * chromosomeSize
			// Pre-calculate the bitmask used to isolate the bits of the gene
			mask := (uint64(1) << locus.Len) - 1
			// Get the left side DNA, with the gene aligned to the right
			left := new(big.Int).Rsh(dna[0], offset + (chromosomeSize - (locus.Start + locus.Len)))
			// Mask it, and store as a 32 bit integer (genes can be 32 bits long maximum)
			geneLeft := uint32(left.Uint64() & mask)
			// Same for right side of the DNA
			right := new(big.Int).Rsh(dna[1], offset + (chromosomeSize - (locus.Start + locus.Len)))
			geneRight := uint32(right.Uint64() & mask)

			//express genes
			geneExpressor(geneLeft, pepeLook)
			//recessive results for b will not overwrite those of a, dominant will.
			geneExpressor(geneRight, pepeLook)
		}
	}

	ResolveLookConflicts(pepeLook)

	//Look is fully read from DNA! return!
	return pepeLook
}

var simpleEyes = []string{
	"eyes>colored_eyes",
	"eyes>closed_eyes",
	"eyes>crying_eyes",
	"eyes>half_closed_eyes",
	"eyes>illuminati_eye",
	"eyes>red_eyes",
	"eyes>small_eyes",
	"eyes>smug_eyes",
}

var simpleMouths = []string{
	"mouth>basic_lips",
	"mouth>happy_lips",
	"mouth>smug_lips",
	"mouth>young_lips",
}

func isSimpleEyes(id string) bool {
	for _, eye := range simpleEyes {
		if eye == id {
			return true
		}
	}
	return false
}

func isSimpleMouth(id string) bool {
	for _, mouth := range simpleMouths {
		if mouth == id {
			return true
		}
	}
	return false
}

// Resolves conflicts between parts: some parts do not fit with others.
func ResolveLookConflicts(pepeLook *look.PepeLook) {
	if pepeLook.Head.Eyes.EyeType == "eyes>ghandi" {
		pepeLook.Extra.Glasses.GlassesType = "none"
		pepeLook.Head.Hair.HairType = "none"
		pepeLook.Body.Shirt.ShirtType = "none"
		pepeLook.Head.Mouth = "none"
	}
	if pepeLook.Body.Shirt.ShirtType == "shirt>darth_pepe" {
		pepeLook.Extra.Glasses.GlassesType = "none"
		pepeLook.Head.Hair.HairType = "none"
		pepeLook.Body.Neck = "none"
		// Darth vader has special black skin
		pepeLook.Skin.Color = "#000000"
	}
	if pepeLook.Extra.Glasses.GlassesType == "glasses>pirate_hat" {
		pepeLook.Head.Hair.HairType = "none"
	}
	if pepeLook.Head.Hair.HairType == "hair>terrorist" {
		pepeLook.Head.Mouth = "none"
		pepeLook.Extra.Glasses.GlassesType = "none"
	}
	if pepeLook.Head.Hair.HairType == "hair>pharaoh" {
		pepeLook.Extra.Glasses.GlassesType = "none"
		pepeLook.Body.Neck = "none"
		if !isSimpleEyes(pepeLook.Head.Eyes.EyeType) {
			pepeLook.Head.Eyes.EyeType = "eyes>colored_eyes"
		}
	}
	if pepeLook.Head.Hair.HairType == "hair>egyptian_hat" {
		pepeLook.Extra.Glasses.GlassesType = "none"
		if !isSimpleEyes(pepeLook.Head.Eyes.EyeType) {
			pepeLook.Head.Eyes.EyeType = "eyes>colored_eyes"
		}
	}
	if pepeLook.Head.Eyes.EyeType == "eyes>future_robot_eyes" ||
		pepeLook.Head.Eyes.EyeType == "eyes>woke_eyes" {
		pepeLook.Extra.Glasses.GlassesType = "none"
	}
	if pepeLook.Extra.Glasses.GlassesType == "glasses>explosion_goggles" ||
		pepeLook.Extra.Glasses.GlassesType == "glasses>vr_set" {
		pepeLook.Head.Eyes.EyeType = "none"
	}
	if pepeLook.Head.Hair.HairType == "hair>frankenstein" ||
		pepeLook.Head.Hair.HairType == "hair>knife_through_head" ||
		pepeLook.Head.Hair.HairType == "hair>rollers" {
		pepeLook.Extra.Glasses.GlassesType = "none"
	}
	if pepeLook.Head.Eyes.EyeType == "eyes>monkas_eye" {
		// monkaS has sweat as "hair/hat"
		pepeLook.Head.Hair.HairType = "none"
	}
	if pepeLook.Head.Hair.HairType == "hair>chaplin" ||
		pepeLook.Head.Hair.HairType == "hair>bun_beard" ||
		pepeLook.Head.Hair.HairType == "hair>mcafee" {
		if !isSimpleMouth(pepeLook.Head.Mouth) {
			pepeLook.Head.Mouth = "mouth>basic_lips"
		}
	}//
	if pepeLook.Body.Shirt.ShirtType == "shirt>pepemon" {
		pepeLook.Head.Hair.HairType = "none"
		pepeLook.Body.Neck = "none"
		// Special pika-yellow skin
		pepeLook.Skin.Color = "#fef135"
		// Enforce simple lips, red dot next to mouth has to fit
		if !isSimpleMouth(pepeLook.Head.Mouth) {
			pepeLook.Head.Mouth = "mouth>basic_lips"
		}
	}
	if pepeLook.Head.Mouth == "mouth>feels_birthday" {
		pepeLook.Head.Hair.HairType = "none"
	}
	if pepeLook.Head.Mouth == "mouth>pacman" {
		pepeLook.Extra.Glasses.GlassesType = "none"
		pepeLook.Head.Eyes.EyeType = "none"
		pepeLook.Head.Hair.HairType = "none"
	}
	if pepeLook.Head.Hair.HairType == "hair>nun" ||
		pepeLook.Head.Hair.HairType == "hair>samurai"{
		pepeLook.Extra.Glasses.GlassesType = "none"
		pepeLook.Body.Shirt.ShirtType = "none"
	}
}
