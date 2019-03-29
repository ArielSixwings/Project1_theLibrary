package mygraph

/**
 * @brief      Class for edge.
 */
type Edge struct {
    Number     int
    EdgeWeight float32
}

/**
 * @brief      Class for node.
 */
type Node struct {
    Edges []Edge
}

/**
 * @brief      Class for graph. Main
 */
type Graph struct {
    Nodes []Node
    size  int
}
