package mygraph

/**
 * @brief      Class for literary work.
 */
type LiteraryWork struct {
    Author     string
    AuthorKey  int
    PrimaryKey int
    Letter     int
    Name       string
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
    Field []InerRing
    size  int
}
