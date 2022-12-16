# Go Array
- Array are used to store multiple values of the same type 
- Array has fixed size

## Declate an Array
1. Using var keyword with length defined
	- Syntax: 	var arrayName = [length]dataType{values}
	Eg:
	```
		var array1 =  [4]string{"a",  "b",  "C",  "d"}
		var array2 = [3]int{1,2,3}
	```
2. Using var keyword without defining length
	- Syntax:		var arrayName = [. . .]dataType{values}
	Eg:
	```
		 var array1 = [...]int{1,2,3,4,5}
	```
	- The lenght  of the array will be the total number of the values
3. Using := sign
	- Syntax: 	arrayName := [length]dataType{values}
	Eg:
	```
		array1 := [4]string{"a","b","b","v"}
		array2 := [...]string{"a","b","b","v"}
		array3 := [...]boot{true,false}
	```

## Access and change the elements of the Array
- We can access and change the element using index.
- Index starts from 0.
Eg:
	```
		array1 := [...]string{"a","b","b","v"}
		fmt.Println(array1[2])
		array1[2] = "50"
	```

## Array Initilization
- We can only delare array, partially initialize, fully initialize the array or initialize particular index of array
	Eg:
	```
		arr1 := [4]int{} //not initialized  
		arr2 := [4]int{6,56} //partially initialized  
		arr3 := [4]int{1,2,6,4,5} //fully initialized
		arr4 := [5]int{1:10,2:40} // index 1 --> 10 and index 2 --> 40
	```	

# Go Slices
- Slices are similar to the array.
- It is more powerfull and flexible
- Store multiple values of same data types
- Length of the slices can grow and shrink
- Syntax:		var sliceName = []dateType{values}
Eg:
	```
		var slice1 =  []string{"a",  "b", "c"}
		slice2 := []int{1,2,43,5,566,76,7}
		fmt.Println(len(slice1)) // return the length of slice
		fmt.Println(cap(slice2)) // return the capasity of slice
	```
## Creating the slice with make function
- Syntax:		sliceName := make([]type, length, capacity)
Eg:
	```
		var slice1 =  make([]int, 5,10)
	```
## Append the slice
- Syntax: 	sliceName = append(sliceName, element1, element2,..)
					sliceName1 = append(sliceName1, sliceName2,...)
Eg:
	```
		var slice1 =  []string{"a",  "b", "c"}
		slice1 = append(slice1, "d","e")
		slice2 := []string{"m","n"}
		slice1 = append(slice1, slice2)
	```