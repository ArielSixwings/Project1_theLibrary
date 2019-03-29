package mygraph

/**
 * @brief      Class for letter ring.
 */
type LetterRing struct {
    LetterGallery []LiteraryWork
    Letter        int
    LetterLabel   rune
}

/**
 * @brief      Class for literary work.
 */
type LiteraryWork struct {
    Author     string
    AuthorKey  int
    PrimaryKey int
    Letter     int
    Title      string
}

/**
 * @brief      Class for outer ring.
 */
type OuterRing struct {
    Gallery []LiteraryWork
    Label   string
}

/**
 * @brief      Class for iner ring.
 */
type InerRing struct {
    Right OuterRing
    Left  OuterRing
    Label string
}

/**
 * @brief      Class for mandala graph.
 */
type MandalaGraph struct {
    Field     []InerRing
    Dictonary []LetterRing
    size      int
}

func (this *MandalaGraph) BuildMandala() bool {

    this.Field = make([]InerRing, 8)
    this.Dictonary = make([]LetterRing, 26)
    this.size = 8

    Subjects := []strings{
        "Artes",
        "Biologia",
        "Computação",
        "Física",
        "História",
        "Matemática",
        "Química",
        "Saude",
    }
    auxOutherRing = make([]OuterRing, 8)
    SecondarySubjects := []strings{
        "Análise Numérica",
        "Arte Digital",
        "História da Arte",
        "Psicologia",
        "Paleolontogia",
        "Nutrição",
        "Estudos Moleculares",
        "Filosofia",
    }
    for i := 0; i < 8; i++ {
        auxOuterRing.Label = SecondarySubjects[i]
        this.Field[i].Label = Subjects[i]
        this.Field[i].Right = auxOuterRing[i]
        if i < 7 {
            this.Field[i].Left = auxOuterRing[i+1]
        } else {
            this.Field[i].Left = auxOuterRing[0]
        }
    }
    for i := 0; i < 26; i++ {
        this.Dictonary[i].Letter = i
        this.Dictonary[i].LetterLabel = rune(i + 65)
        //this.Dictonary[i].LetterGallery = ReadLetter(i) // TO DO função que lé todas as obras de uma letra de um arquivo feito para a letra
    }
}
