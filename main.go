package main

import "fmt"

func main() {

  //declare variable for gas tank size and miles travled.

 var tankSize int
 var milesTravled int

 //ask user to to input their gas tank size.
 
 fmt.Println("Enter the amount of gallon your gas tank holds.")
 
 //scan the line above for the tankSize.

 fmt.Scanln(&tankSize)

 //ask user to enter the amount of miles they have travled.

 fmt.Println("Enter the amount of miles you have travled.")
 
 //scan the line abcve for the amount of miles travled.
 
 fmt.Scanln(&milesTravled)

 //Devide tankSize by milesTravled to find the average miles per gallon.

 var averageMPG = (milesTravled / tankSize)

 //Print a message using the averageMPG to tell the user the average miles per gallon

 fmt.Println("Your average miles per gallon was" , averageMPG ,"mpg") 

 //end.
}