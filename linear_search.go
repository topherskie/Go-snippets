// check array data from O to length of array sequentially
package main

import "fmt"

func linear_search(arrLen [7]int, search int) {
  for i:=0; i < len(arrLen); i++ {
    if (search != arrLen[i]) {
      fmt.Println("Data not found!", arrLen[i]);
    } else if (search == arrLen[i]) {
      fmt.Println("Data have match!", arrLen[i]);
    } else {
      fmt.Println("Data Error!");
    }
  }
}


func main() {
  var arr1 = [7]int{2021, 1978, 2006, 1956, 1945, 2008, 2001};

   search := 1945;


  // invoke function
   linear_search(arr1, search);




} // end of main function
