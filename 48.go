func rotate(matrix [][]int)  {    
    for i := range matrix {
        for j := i; j < len(matrix[0]); j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
    
    for i := range matrix {
        for j := 0; j < len(matrix) / 2; j++ {
            matrix[i][j], matrix[i][len(matrix) - 1 - j] = matrix[i][len(matrix) - 1 - j], matrix[i][j]
        }
    }
}
