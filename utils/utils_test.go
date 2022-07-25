package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMapToList(t *testing.T) {
	input := map[string]string{
		"b": "b-value",
		"a": "a-value",
		"c": "c-value",
	}
	list := MapToList(input)
	assert.Len(t, list, 3)

	// should be sorted
	assert.Equal(t, []string{
		"a=a-value", "b=b-value", "c=c-value",
	}, list)
}

func TestNormalizeJSON(t *testing.T) {
	testCases := []struct {
		name         string
		input        interface{}
		expected     string
		expectedHash string
	}{
		{
			name: "unordered map",
			input: map[string]interface{}{
				"k": 7,
				"z": "91",
				"f": 123.45,
				"b": "ninety nine",
			},
			expected:     `{"b":"ninety nine","f":123.45,"k":7,"z":"91"}`,
			expectedHash: "8faf96ab488c4d7942c23b2b061533664051de567328bc0f97e4c36044d11727",
		},
		{
			name:         "ordered raw object",
			input:        `{"b":"ninety nine","f":123.45,"k":7,"z":"91"}`,
			expected:     `{"b":"ninety nine","f":123.45,"k":7,"z":"91"}`,
			expectedHash: "8faf96ab488c4d7942c23b2b061533664051de567328bc0f97e4c36044d11727",
		},
		{
			name:         "unordered raw object",
			input:        []byte(`{"z":"91","b":"ninety nine","f":123.45,"k":7}`),
			expected:     `{"b":"ninety nine","f":123.45,"k":7,"z":"91"}`,
			expectedHash: "8faf96ab488c4d7942c23b2b061533664051de567328bc0f97e4c36044d11727",
		},
		{
			name: "unordered struct with map field",
			input: struct {
				Zoo string                 `json:"zoo"`
				Boo string                 `json:"boo"`
				Foo string                 `json:"foo"`
				Xoo map[string]interface{} `json:"xoo"`
			}{
				Zoo: "zoo",
				Boo: "boo",
				Foo: "foo",
				Xoo: map[string]interface{}{
					"f": 11,
					"k": 18,
					"b": 2,
				},
			},
			expected:     `{"boo":"boo","foo":"foo","xoo":{"b":2,"f":11,"k":18},"zoo":"zoo"}`,
			expectedHash: "f7ae200915fea285b99b2ee5647adf739ed94bbfdfded2b6ac3af174e835acf3",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			r := require.New(t)
			result := NormalizeJSON(testCase.input)
			hash := HashNormalizedJSON(testCase.input)
			r.Equal(testCase.expected, string(result))
			r.Equal(testCase.expectedHash, hash)
		})
	}
}
