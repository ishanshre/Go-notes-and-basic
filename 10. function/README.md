# Function in Golang
## Simple Function
- function is a block of code that can be reused repeately.
- Function will be executed only when they are called
- Syntax: 
    ```
        func functionName() {
            // code block
        }
    ```
## Parameters and Arguments
- Parameters are the variables in function that receives the values from the argument
- Argument is the value that are passed to parameters when the function is called
- Syntax:
    ```
        func functionName(param1 dataType, param2 dataType, param3 dataType) {
            // code to be executed
        }

        func main() {
            functionName(argument1, argument2, argument3)
        }
    ```

## Return type function
- Return a value when the function is called
- DataType of the return value must be defined
- Syntax:
    ```
        func functionName(param1 dataType, param2 dataType, param3 dataType) dataType {
            // code to be executed
            return output
        }

        func main() {
            a := functionName(argument1, argument2, argument3)
        }
    ```

## Functions are also recursive
- It means a function calls itself until the counter stops