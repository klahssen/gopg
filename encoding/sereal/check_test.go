package sereal

import (
	"testing"

	"github.com/klahssen/tester"
)

func TestEncodeMap(t *testing.T) {
	te := tester.NewT(t)
	tests := []struct {
		data map[string]interface{}
		err  error
		res  map[string]interface{}
	}{
		{
			data: map[string]interface{}{"a": "a", "b": 1, "c": map[string]interface{}{"a.a": "a", "a.b": 1}},
			err:  nil,
			res:  map[string]interface{}{"a": "a", "b": 1, "c": map[string]interface{}{"a.a": "a", "a.b": 1}},
		},
		{
			data: map[string]interface{}{"a": "a", "b": 1, "c": nil},
			err:  nil,
			res:  map[string]interface{}{"a": "a", "b": 1},
		},
	}

	for ind, test := range tests {
		b, err := Marshal(test.data)
		te.CheckError(ind, test.err, err)
		if err != nil {
			continue
		}
		dest := map[string]interface{}{}
		err = Unmarshal(b, &dest)
		te.CheckError(ind, test.err, err)
		if err != nil {
			continue
		}
		te.DeepEqual(ind, "unmarshaled", test.res, dest)
	}
}

type testStruct struct {
	A string `sereal:"a"`
	B int    `sereal:"b"`
	C innerStruct
	D interface{} `sereal:"d,omitempty"`
}
type innerStruct struct {
	A string `sereal:"c.a"`
	B int    `sereal:"c.b"`
}

func TestEncodeStruct(t *testing.T) {
	te := tester.NewT(t)
	tests := []struct {
		data testStruct
		err  error
		res  testStruct
	}{
		{
			data: testStruct{
				A: "a",
				B: 1,
				C: innerStruct{
					A: "a",
					B: 1,
				},
				D: int(5),
			},
			err: nil,
			res: testStruct{
				A: "a",
				B: 1,
				C: innerStruct{
					A: "a",
					B: 1,
				},
				D: int(5),
			},
		},
		{
			data: testStruct{
				A: "a",
				B: 1,
				C: innerStruct{
					A: "a",
					B: 1,
				},
				D: nil,
			},
			err: nil,
			res: testStruct{
				A: "a",
				B: 1,
				C: innerStruct{
					A: "a",
					B: 1,
				},
				D: nil,
			},
		},
	}

	for ind, test := range tests {
		b, err := Marshal(test.data)
		te.CheckError(ind, test.err, err)
		if err != nil {
			continue
		}
		dest := testStruct{}
		err = Unmarshal(b, &dest)
		te.CheckError(ind, test.err, err)
		if err != nil {
			continue
		}
		te.DeepEqual(ind, "unmarshaled", test.res, dest)
	}
}
