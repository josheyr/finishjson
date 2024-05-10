package finishjson

import "testing"

func TestFinishJson(t *testing.T) {
	tests := []struct {
		name           string
		unfinishedJson string
		expected       string
	}{
		{
			name:           "Complete JSON",
			unfinishedJson: `{"key": "value", "key2": "value2"}`,
			expected:       `{"key": "value", "key2": "value2"}`,
		},
		{
			name:           "Truncated JSON at the end",
			unfinishedJson: `{"key": "value", "key2":`,
			expected:       `{"key": "value", "key2": null}`, // Adjust based on the function's expected behavior
		},
		{
			name:           "Truncated true value",
			unfinishedJson: `{"key": tr`,
			expected:       `{"key": true}`, // Adjust based on the function's expected behavior
		},
		{
			name:           "Truncated false value",
			unfinishedJson: `{"key": f`,
			expected:       `{"key": false}`, // Adjust based on the function's expected behavior
		},
		{
			name:           "Truncated escaped quote",
			unfinishedJson: `{"key": "\""`,
			expected:       `{"key": "\""}`, // Adjust based on the function's expected behavior
		},

		{
			name:           "Truncated JSON with nested objects",
			unfinishedJson: `{"outer": {"inner1": "value1", "inner2": "value2", "inner3": {`,
			expected:       `{"outer": {"inner1": "value1", "inner2": "value2", "inner3": {}}}`, // Modify as per your expected behavior
		},
		{
			name:           "Truncated JSON with arrays",
			unfinishedJson: `{"array": [1, 2, 3,`,
			expected:       `{"array": [1, 2, 3,null]}`, // Adjust as per expected correction logic
		},
		{
			name:           "Empty JSON",
			unfinishedJson: ``,
			expected:       `{}`, // Example of expected behavior for empty input
		},
		{
			name:           "Malformed JSON structure",
			unfinishedJson: `{"key": tr`,
			expected:       `{"key": true}`, // Example of expected behavior for malformed JSON
		},
		{
			name:           "Truncated numeric value",
			unfinishedJson: `{"key": 12`,
			expected:       `{"key": 12}`, // Adjust based on the function's expected behavior
		},
		{
			name:           "Truncated array element",
			unfinishedJson: `{"array": [1, 2, `,
			expected:       `{"array": [1, 2, null]}`, // Example of expected behavior for an unfinished array element
		},
		{
			name:           "Incomplete Unicode escape sequence",
			unfinishedJson: `{"key": "\u123`,
			expected:       `{"key": "\u123"}`, // Completion logic should finish a Unicode escape sequence
		},
		{
			name:           "Empty object key",
			unfinishedJson: `{"": "value"`,
			expected:       `{"": "value"}`, // Handle missing key gracefully if relevant
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FinishJSON(test.unfinishedJson)
			if result != test.expected {
				t.Errorf("Got %q, expected %q", result, test.expected)
			}
		})
	}
}
