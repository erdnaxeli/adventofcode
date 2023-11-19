package main

import "testing"

func TestRoomIsValid(t *testing.T) {
	tests := []struct {
		test     string
		expected bool
	}{
		{
			"aaaaa-bbb-z-y-x-123[abxyz]",
			true,
		},
		{
			"a-b-c-d-e-f-g-h-987[abcde]",
			true,
		},
		{
			"not-a-real-room-404[oarel]",
			true,
		},
		{
			"totally-real-room-200[decoy]",
			false,
		},
	}

	for _, test := range tests {
		room := NewRoom(test.test)
		valid := room.IsValid()
		if valid != test.expected {
			t.Errorf(
				"Expected %t but got %t: %v, %s",
				test.expected,
				valid,
				room,
				room.ComputeChecksum(),
			)
		}
	}
}
