package main

import (
       "fmt"
       "io/ioutil"
       "os"
)

type WikiPage struct {		  
     Title string		  
     Body []byte		  
}