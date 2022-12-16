# Go variable 

## Declaring a variable with Initial Value
    1. Using var keyword
            - Using var keywork web must follow it by name and type 
            Syntax: var varName dataType = value
            ```
            var name string = 'Nepal';
            var status bool = true;
            var status1 bool = false;
            var number int = 1000;
            var floatingPoint float32 = 1435.345;
            ```
    2. Using := sign
            - Compiler auto identify the data types
            - It cannot be declared outside the function
            Syntx:  varName := value
            ```
                country := "Nepal"
                status := false
                number := 987655
            ```

## Declaring a variable without Initial Value
    ```
        var name string
        var number int
        var status bool
    ```

## Declaring multiple variable in single line
    ```
        var a, b, c, d int = 1,2,3,4
        a, b := "Hello", "world"
    ```

## For constant
    1. Typed Constant
        - constant using const keyword followed by constant name and data type
    ```
        const name string = "Hello World!"
    ```
    2. Untyped Constant
        - constant using const keywoed followed by constant name
        ```
            const name = "Nepal"
        ```