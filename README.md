# FinishJSON

## Description
The `FinishJSON` function helps complete a potentially unfinished JSON string by ensuring that it adheres to the correct structure. It adds necessary closing brackets, quotes, or default values (e.g., `null`) if the input JSON string is incomplete.

### Key Features
- **Handles Unmatched Brackets/Quotes**: Automatically closes open `{}`, `[]`, and `""` pairs.
- **Completes Keywords**: Finishes keywords like `true`, `false`, and `null` if they're partially typed.
- **Removes Extraneous Whitespace**: Ignores whitespace while analyzing the JSON string.

### Parameters
- `unfinished` (`string`): The input JSON string that needs to be checked and corrected.

### Returns
- Returns a well-formed JSON string that is either completed based on the detected pattern or returns a default empty JSON object (`{}`) if no initial structure is found.

### How to Install
```bash
go get github.com/josheyr/finishjson
```

### How to Import
```go
import "github.com/josheyr/finishjson/pkg/finishjson"
```

### Example Usage
```go
input := `{"name": "John", "age": 30, "active": true, "skills": ["Go", "Python"], "address":`
result := finishjson.FinishJSON(input)
// result will be '{"name": "John", "age": 30, "active": true, "skills": ["Go", "Python"], "address": null}'
```

This function is particularly useful when working with incomplete JSON strings, such as those truncated by LLMs, during data exchange or in API calls.