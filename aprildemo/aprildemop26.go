type Day int
var dayName = []string{"Sunday", "Monday"} // a comment
func (d Day) String() string {
   if 0 <= d && int(d) < len(dayName) { return dayName[d] }
   return "NoSuchDay"
}
type Fahrenheit float
func (t Fahrenheit) String() string {
   return fmt.Sprintf("%.1fºF", t)
}
