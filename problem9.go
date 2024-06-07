package main

import (
	"fmt"
)

func arrayMultiplier (matrix [3][3]int, vector [3]int) ([3]int){
    // helper for tree method
    // multiplies a 3x3 array by a 3 element vector and returns result
    var vectorRes [3]int;
    vectorRes[0] = matrix[0][0]*vector[0]+matrix[0][1]*vector[1]+matrix[0][2]*vector[2];
    vectorRes[1] = matrix[1][0]*vector[0]+matrix[1][1]*vector[1]+matrix[1][2]*vector[2];
    vectorRes[2] = matrix[2][0]*vector[0]+matrix[2][1]*vector[1]+matrix[2][2]*vector[2];

    return vectorRes;
}

func treeMethod() (int) {

    // will try to solve [problem 9](https://projecteuler.net/problem=9)
    // using [the tree](https://en.wikipedia.org/wiki/Tree_of_primitive_Pythagorean_triples)

    var a, b, c int 
    var  s = a + b + c;

    a, b, c = 4, 3, 5;

    // Matrix A
    pythMatr := [3][3]int{
        {1, -2, 2},
        {2, -1, 2},
        {2, -2, 3},
    }

    // Matrix B
    // pythMatr := [3][3]int{
    //     {1, 2, 2},
    //     {2, 1, 2},
    //     {2, 2, 3},
    // }

    // Matrix C
    // pythMatr := [3][3]int{
    //     {-1, 2, 2},
    //     {-2, 1, 2},
    //     {-2, 2, 3},
    // }

    // column vector triplet
    triplet := [3]int{a, b, c,} 


    for s <= 1000 {
        triplet = arrayMultiplier(pythMatr,triplet)
        a = triplet[0]
        b = triplet[1]
        c = triplet[2]
        s = a+b+c

        fmt.Print(a, b, c, "\n")
        fmt.Print(s, "\n")
    }

    return s

}

func dicksons() {

    // [Dicksons & Height excess](https://en.wikipedia.org/wiki/Formulas_for_generating_Pythagorean_triples)

    var tripletSlice [][3]int;

    for r := 6; r < 200; r++ {
        
        if ((r*r%2)==0){  

            for s := 1; s < (r*r/2)+1; s++ {
                for t := 1; t < (r*r/2)+1; t++ {
                    if (s*t == r*r/2){
                        tripletSlice = append(tripletSlice, [3]int{r, s, t})
                        // abc calc
                        abcSum := (r+s)+(r+t)+(r+s+t);
                        
                        if (abcSum==1000){
                            abcProd := (r+s)*(r+t)*(r+s+t);
                            fmt.Println("r, s, t: ", r,s,t,)
                            fmt.Println("a: ", r+s, ", b: ", r+t, " ,c: ", r+s+t)
                            fmt.Println("a, b, c Sum:", abcSum)
                            fmt.Println("a, b, c Product: ", abcProd)
                            break
                        } 


                    } 
                }         
            }

        }

    }
    // fmt.Println(tripletSlice)


}

func main() {
    fmt.Println("Hello, World!")

    // uncomment to run the tree method
    // treeMethod()

    // uncomment to run dicksons
    dicksons()
}
