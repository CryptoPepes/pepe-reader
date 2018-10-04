package bio_gen

type Bio struct {

  Title string

  Description string

}

func (bio *Bio) String() string {
  return bio.Title+`
--------------------------------
`+bio.Description+`

`
}


