# Go For Loops
- For loop iterate through a code block
- For loops contains three statements: counter value , condition and counter value increase or decrease
- Syntax:
    ```
        for counter value;condition;counter value inc or dec {
            // code block to iterate
        }
    ```

## Go for loop for array, slice and maps
- We use range keyword 
- Syntax:
    ```
        for index, value := range varName {
            // code block to execute
        }
        for _, value := range varName {
            // code block to execute
        }
        for index, _ := range varName {
            // code block to execute
        }
    ```
- If we only want either index or value, replace the one you don't need with "_"

