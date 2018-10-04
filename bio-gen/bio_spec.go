package bio_gen

import (
  "cryptopepe.io/cryptopepe-reader/pepe"
  "math/big"
)

type BioYmlSpec struct {

  Bio BioData `yaml:"bio"`
  Lists Lists `yaml:"lists"`

}

type BioData struct {

  Title Title `yaml:"title"`
  Intro Intro `yaml:"intro"`
  ILike ILike `yaml:"ilike"`
  Twist Twist `yaml:"twist"`
  EndWish EndWish `yaml:"end_wish"`
  Additional Additional `yaml:"additional"`
}

type Title struct {

  TitleAdj BioPart `yaml:"title_adj"`
  TitleNoun BioPart `yaml:"title_noun"`
  SuffixAdj BioPart `yaml:"suffix_adj"`
  SuffixNoun BioPart `yaml:"suffix_noun"`

}

type Intro struct {

  Scream BioPart `yaml:"scream"`
  Saying BioPart `yaml:"saying"`
  SelfDescription BioPart `yaml:"self_description"`

}

type ILike struct {

  Start BioPart `yaml:"start"`
  Subject BioPart `yaml:"subject"`
  WhileTopic BioPart `yaml:"while_topic"`

}

type Twist struct {

  CallOut BioPart `yaml:"call_out"`
  Question BioPart `yaml:"question"`

}

type EndWish struct {

  Start BioPart `yaml:"start"`
  End BioPart `yaml:"end"`

}

type Additional struct {

  Quote BioPart `yaml:"quote"`
  EndScream BioPart `yaml:"end_scream"`
  SadEnding BioPart `yaml:"sad_ending"`
  HelpRequest BioPart `yaml:"help_request"`

}

type Lists struct {

  CryptoCeleb BioPart `yaml:"crypto_celeb"`
  BigHack BioPart `yaml:"big_hack"`
  Shitcoin BioPart `yaml:"shitcoin"`
  ShitPeopleSay BioPart `yaml:"shit_people_say"`

}

type BioPart struct {

  Entries []*BioPartEntry

  DnaSpec DnaSpec

}

type DnaSpec struct {

  // The position and length of the gene, range [0, 255],
  Position uint32 `yaml:"pos"`
  Length uint32 `yaml:"length"`

}

type BioPartEntry struct {

  Content string `yaml:"content"`

  Rarity float64 `yaml:"rarity"`
  Dominant bool `yaml:"dominant"`

  // Ignore these fields, they are changed dynamically on loading time.
  // These are the limits (in normal system, not gray) of the allele.
  Start uint32 `yaml:"-"`
  End uint32 `yaml:"-"`

}

type YmlRawBioPart struct {
  DnaSpec DnaSpec `yaml:"dna"`

  Entries []*BioPartEntry `yaml:"entries"`
}

// Implements the Unmarshaller interface of the yaml pkg.
func (part *BioPart) UnmarshalYAML(unmarshal func(interface{}) error) error {

  // Create temporary struct, to prevent unmarshaller from looping back to this function
  tempBioPart := new(YmlRawBioPart)

  if err := unmarshal(tempBioPart); err != nil {
    return err
  }

  part.DnaSpec = tempBioPart.DnaSpec
  part.Entries = tempBioPart.Entries

  maxGeneUint64 := uint64(1 << part.DnaSpec.Length) - 1
  maxGeneFloat64 := float64(maxGeneUint64)

  // Make rarities relative to summed rarity
  raritySum := 0.0
  for _, entry := range part.Entries {
    raritySum += entry.Rarity
  }

  lastEnd := uint32(0)
  for _, entry := range part.Entries {
    geneSpan := uint32((entry.Rarity / raritySum) * maxGeneFloat64)
    entry.Start = lastEnd
    entry.End = entry.Start + geneSpan
    lastEnd = entry.End
  }

  // Set the end limit of the last entry to the max gene;
  //  floating point errors may cause a miss in the last (few) possible dna values.
  part.Entries[len(part.Entries) - 1].End = uint32(maxGeneUint64)

  return nil
}

func (part *BioPart) pickEntryWithRawDNA(geneDNA uint32, dst **BioPartEntry) {
  gray := pepe.GrayToBinary32(geneDNA)

  for i, entry := range part.Entries {
    //last range is end-inclusive.
    if gray >= entry.Start && (gray < entry.End ||
      (i == len(part.Entries) - 1 && gray == entry.End)) {
      //dominant: always overwrite whatever feno-look it may already have.
      //recessive (non-dominant): only write feno-look when it does not have one already.
      if entry.Dominant || dst == nil {
        *dst = entry
        return
      }
    }
  }
}

func (part *BioPart) PickEntry(dna *pepe.PepeDNA) *BioPartEntry {

  // For bot sides of the DNA:
  // 1) Get Gene (masked 32 bit int) from the 256 bit integer

  // Pre-calculate the bitmask used to isolate the bits of the gene
  mask := (uint64(1) << part.DnaSpec.Length) - 1
  // Get the left side DNA, with the gene aligned to the right
  left := new(big.Int).Rsh(dna[0], uint(256 - (part.DnaSpec.Position + part.DnaSpec.Length)))
  // Mask it, and store as a 32 bit integer (genes can be 32 bits long maximum)
  geneLeft := uint32(left.Uint64() & mask)
  // Same for right side of the DNA
  right := new(big.Int).Rsh(dna[1], uint(256 - (part.DnaSpec.Position + part.DnaSpec.Length)))
  geneRight := uint32(right.Uint64() & mask)


  var res *BioPartEntry
  part.pickEntryWithRawDNA(geneLeft, &res)
  // second time will only apply if first is not already dominant
  part.pickEntryWithRawDNA(geneRight, &res)

  return res
}
