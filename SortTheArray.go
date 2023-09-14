package main
import "fmt"

func sort_array(arr []int, n int, min int) []int {
    temp := 0
    for i := 0; i < n; i++ {
        min = i
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[i] {
                min = j
            }
        }
        temp = arr[i]
        arr[i] = arr[min]
        arr[min] = temp
    }
    return arr
}

func main() {
    fmt.Printf("Enter size of your array: ")
    var size int
    fmt.Scanln(&size)
    var arr = make([]int, size)
    for i := 0; i < size; i++ {
        fmt.Scan(&arr[i])
    }
    fmt.Println("Sorted Array : ")
    var min int = 0
    array := sort_array(arr, size, min)

    fmt.Println(array)
}