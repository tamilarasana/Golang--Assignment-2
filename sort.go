package main
 
import (
    "fmt"
    "sort"
)
 
type Mobile struct {
    Brand string
    Price int
}
 
 
// ByPrice implements sort.Interface for []Mobile based on
// the Price field.
type ByPrice []Mobile
func (a ByPrice) Len() int           { return len(a) }
func (a ByPrice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPrice) Less(i, j int) bool { return a[i].Price < a[j].Price }
 
// ByBrand implements sort.Interface for []Mobile based on
// the Brand field.
type ByBrand []Mobile
func (a ByBrand) Len() int           { return len(a) }
func (a ByBrand) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByBrand) Less(i, j int) bool { return a[i].Brand > a[j].Brand }
 
func main() {
    mobile := []Mobile{
        {"Sony", 952},
        {"Nokia", 468},
        {"Apple", 1219},
        {"Samsung", 1045},
    }
    fmt.Println("\n######## Before Sort #############\n")
    for _, v := range mobile {
        fmt.Println(v.Brand, v.Price)
    }
     
    fmt.Println("\n\n######## Sort By Price [Ascending] ###########\n")
    sort.Sort(ByPrice(mobile))
    for _, v := range mobile {
        fmt.Println(v.Brand, v.Price)
    }   
     
    fmt.Println("\n\n######## Sort By Brand [Descending] ###########\n")
    sort.Sort(ByBrand(mobile))
    for _, v := range mobile {
        fmt.Println(v.Brand, v.Price)
    }
}