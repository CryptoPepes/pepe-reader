package bio_gen

import (
  "os"
  "path"
  "runtime"
  "io/ioutil"
  "gopkg.in/yaml.v2"
  "cryptopepe.io/cryptopepe-reader/pepe"
  "bytes"
  "strings"
  "math/big"
)

type BioGenerator struct {

  bioSpec BioYmlSpec
}

func (bioGen *BioGenerator) Load() {


  basePath, envSetBasePath := os.LookupEnv("APP_PATH")
  if !envSetBasePath {
    _, filename, _, ok := runtime.Caller(0)
    if !ok {
      panic("No caller info! Cannot get file path for Bio YML spec loading.")
    }
    basePath = path.Dir(filename)
  }

  partsConfigPath := path.Join(basePath, "bio_config.yml")

  dat, err := ioutil.ReadFile(partsConfigPath)
  if err != nil {
    panic(err)
  }

  if err := yaml.UnmarshalStrict(dat, &bioGen.bioSpec); err != nil {
    panic(err)
  }

}

func (bioGen *BioGenerator) ConvertDnaToBio(dna *pepe.PepeDNA) *Bio {

  bio := new(Bio)

  title := bioGen.bioSpec.Bio.Title
  titleAdj := title.TitleAdj.PickEntry(dna)
  titleNoun := title.TitleNoun.PickEntry(dna)
  suffixAdj := title.SuffixAdj.PickEntry(dna)
  suffixNoun := title.SuffixNoun.PickEntry(dna)

  intro := bioGen.bioSpec.Bio.Intro
  scream := intro.Scream.PickEntry(dna)
  saying := intro.Saying.PickEntry(dna)
  selfDescription := intro.SelfDescription.PickEntry(dna)


  ilike := bioGen.bioSpec.Bio.ILike
  ilikeStart := ilike.Start.PickEntry(dna)
  ilikeSubject := ilike.Subject.PickEntry(dna)
  ilikeWhile := ilike.WhileTopic.PickEntry(dna)


  twist := bioGen.bioSpec.Bio.Twist
  callOut := twist.CallOut.PickEntry(dna)
  question := twist.Question.PickEntry(dna)

  endWish := bioGen.bioSpec.Bio.EndWish
  wishStart := endWish.Start.PickEntry(dna)
  wishEnd := endWish.End.PickEntry(dna)

  additional := bioGen.bioSpec.Bio.Additional
  endScream := additional.EndScream.PickEntry(dna)
  helpRequest := additional.HelpRequest.PickEntry(dna)
  quote := additional.Quote.PickEntry(dna)
  sadEnding := additional.SadEnding.PickEntry(dna)


  dnaBasedChance := func (side int, shift uint, chance float64, yes func(), no func()) {
    dnaSide := dna[side]
    genePart := uint32(new (big.Int).Rsh(dnaSide, shift).Uint64())
    // xor with uniform random 32 bit magic number to not 1:1 correspond to pepe-look results / parts of the bio
    genePart ^= 0x42424242
    if genePart < uint32(float64(0xffffffff) * chance) {
      if yes != nil {
        yes()
      }
    } else {
      if no != nil {
        no()
      }
    }
  }

  w := func (buf *bytes.Buffer, entry *BioPartEntry) {
    buf.WriteString(entry.Content)
  }

  var titleBuf bytes.Buffer
  w(&titleBuf, titleAdj)
  titleBuf.WriteString(" ")
  w(&titleBuf, titleNoun)

  dnaBasedChance(0, 0, 0.7, func(){
    titleBuf.WriteString(" of")
    dnaBasedChance(0, 32, 0.6, func(){
      titleBuf.WriteString(" ")
      w(&titleBuf, suffixAdj)
    }, nil)
    titleBuf.WriteString(" ")
    w(&titleBuf, suffixNoun)
  }, nil)

  bio.Title = titleBuf.String()


  var descBuf bytes.Buffer
  w(&descBuf, scream)
  dnaBasedChance(0, 64, 0.1, func(){
    descBuf.WriteString(" ")
    w(&descBuf, saying)
  }, nil)
  descBuf.WriteString(" ")
  w(&descBuf, selfDescription)



  descBuf.WriteString(" ")
  w(&descBuf, ilikeStart)
  descBuf.WriteString(" ")
  w(&descBuf, ilikeSubject)
  descBuf.WriteString(" ")
  w(&descBuf, ilikeWhile)


  dnaBasedChance(0, 96, 0.1, func(){
    // rare chance for both
    w(&descBuf, callOut)
    descBuf.WriteString(" ")
    w(&descBuf, question)
  }, func(){
    dnaBasedChance(0, 128, 0.4, func(){
      w(&descBuf, callOut)
    }, func(){
      descBuf.WriteString(". ")
      w(&descBuf, question)
    })
  })

  dnaBasedChance(0, 160, 0.8, func(){
    descBuf.WriteString(" ")
    w(&descBuf, wishStart)
    descBuf.WriteString(" ")
    w(&descBuf, wishEnd)
  }, nil)

  descBuf.WriteString(" ")
  dnaBasedChance(1, 0, 0.8, func(){
    dnaBasedChance(1, 32, 0.5, func(){
      dnaBasedChance(1, 64, 0.5, func(){
        dnaBasedChance(1, 96, 0.3, func(){
          w(&descBuf, endScream)
          descBuf.WriteString(". ")
          w(&descBuf, helpRequest)
        }, func(){
          w(&descBuf, sadEnding)
        })
      }, func(){
        w(&descBuf, helpRequest)
      })
    }, func(){
      dnaBasedChance(1, 64, 0.5, func(){
        w(&descBuf, quote)
      }, func(){
        w(&descBuf, sadEnding)
      })
    })
  }, nil)

  bio.Description = descBuf.String()

  applyReplace := func (key string, part *BioPart) {
    bio.Description = strings.Replace(bio.Description, key, part.PickEntry(dna).Content, -1)
  }

  lists := bioGen.bioSpec.Lists
  applyReplace("$shitcoin", &lists.Shitcoin)
  applyReplace("$big_hack", &lists.BigHack)
  applyReplace("$crypto_celeb", &lists.CryptoCeleb)
  applyReplace("$shit_people_say", &lists.ShitPeopleSay)

  return bio
}
