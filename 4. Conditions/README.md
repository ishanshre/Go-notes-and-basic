# GO Conditions

- Less than: <
- Greater than: >
- Less than or equal: <=
- Greater than or equal: >=
- Equal to: ==
- Not Equal to: !=
- Logical AND: &&
- Logical OR: ||
- Logical NOT: !


## GO if statements
- If statement to execute a block of code when the condition is true
- Syntax:
    ```
        if condition {
            // code to be executed when the condition is true
        }
    ```


## GO if else statements
- Execute if block when condition true else execute else block code
- Syntax:
    ```
        if condition {
            // code to be executed when the condition is true
        } else {
            // code to be executed when the condition is false
        }
    ```

## Go else if statements
- Execute if block statement when its condition else check second condition else execute its statement
- Syntax:
- Syntax:
    ```
        if condition {
            // code to be executed when the condition is true
        } else if condition2 {
            // code to be executed when the condition2 is true
        } else if condition3 {
            // code to be executed when the condotion3 is true
        } else {
            // code to be executed when the condition is false
        }
    ```

## Go nested if 
- If statement inside another if statement 
- Syntax:
    ```
        if condition {
            if condition2 {
                // code to be executed when the condition is true
            }
        }
    ```
- Syntax:
    ```
        if condition {
            if condition2 {
                // code to be executed when the condition2 is true
            }
        }
        else {
            if condition3 {
                // code to be executed when the condition3 is true
            }
        }
    ```

