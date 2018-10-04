package pepe

func ExpressPropType(geneDNA uint32, propPtr *string, data *ExpressorData) {

  //get gray-code
  // See: https://en.wikipedia.org/wiki/Gray_code
  //Look-a-like bit-codes are close to each other,
  // which makes it easy to define a range of bit-codes that translate to a property
  // in a consistent way (e.g. do not change on one bit mutation all the time)
  gray := GrayToBinary32(geneDNA)

  for i, expRang := range data.Ranges {
  	//last range is end-inclusive.
    if gray >= expRang.Start && (gray < expRang.End ||
    	(i == len(data.Ranges) - 1 && gray == expRang.End)) {
      //dominant: always overwrite whatever feno-look it may already have.
      //recessive (non-dominant): only write feno-look when it does not have one already.
      if expRang.Dominant || *propPtr == "" {
        *propPtr = expRang.Prop
        return
      }
    }
  }
}

type ExpressorRange struct {
  Prop string
  Start uint32
  End uint32
  Dominant bool
}

type GrayExpressor struct {
  GeneLen uint16
}

var bit5Expr = GrayExpressor{GeneLen: 5}
var bit6Expr = GrayExpressor{GeneLen: 6}
var bit8Expr = GrayExpressor{GeneLen: 8}
var bit10Expr = GrayExpressor{GeneLen: 10}
var bit12Expr = GrayExpressor{GeneLen: 12}

func (grayExpr *GrayExpressor) Expression(prop string, start float64, end float64, dominant bool) ExpressorRange {
  maxGeneUint64 := uint64(1 << grayExpr.GeneLen) - 1
  maxGene := float64(maxGeneUint64)
  start64 := uint64(start * maxGene)
  end64 := uint64(end * maxGene)
  var start32 uint32
  var end32 uint32
  // cap at max value
  if end64 > maxGeneUint64 {
    end32 = uint32(maxGeneUint64)
  } else {
    end32 = uint32(end64)
  }
  // If start is more or equal than end: set start to end - 1
  if start64 >= end64 {
    start32 = end32 - 1
  } else {
    start32 = uint32(start64)
  }

  return ExpressorRange{
    Prop: prop,
    Start: start32,
    End: end32,
    Dominant: dominant,
  }
}

func (grayExpr *GrayExpressor) ExpressColor(geneDNA uint32, propPtr *string, gradient *GradientTable) {
  //decode gray-code
  binary := GrayToBinary32(geneDNA)

  //scale decoded gray-code to 0.0 (incl), 1.0 (excl.) range
  scaled := float64(binary) / float64(uint64(1) << grayExpr.GeneLen)

  color := gradient.GetInterpolatedColorFor(scaled)

  // If bit 1 is set to 1, then the color is dominant, and overwrites any previous value!
  // Use the decoded bitstring, since this has more uniform 0/1 distribution of bit 1.
  if *propPtr == "" || (binary & 1 == 1) {
    *propPtr = color.Hex()
  }
}


type ExpressorData struct {

  // The ranges in the gray-code decoded DNA string to translate to properties.
  Ranges []ExpressorRange

}
