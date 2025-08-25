package gordle

import "testing"

func TestHintReturnsCorrectString(t *testing.T) {
	type testCase struct {
		hint       hint
		wantResult string
	}

	testCases := map[string]testCase{
		"hint string returns no colour if its status is 'Absent'": {
			hint:       hint{character: 'A', status: Absent},
			wantResult: "[A]",
		},
		"hint string returns yellow if its status is 'WrongPosition'": {
			hint:       hint{character: 'A', status: WrongPosition},
			wantResult: TTYYellow + "[A]" + TTYReset,
		},
		"hint string returns green if its status is 'CorrectPosition'": {
			hint:       hint{character: 'A', status: CorrectPosition},
			wantResult: TTYGreen + "[A]" + TTYReset,
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			str := tc.hint.String()
			if tc.wantResult != str {
				t.Errorf("The string representation of this hint is incorrect. Expected %s, got %s", str, tc.wantResult)
			}
		})
	}
}

func TestInvalidHintsPanic(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("This test failed, the expected panic was not caught")
		}
	}()
	h := hint{}
	h.String()
	t.Errorf("This test failed, the expected panic was not caught")
}
