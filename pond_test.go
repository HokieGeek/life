package life

import "testing"

func TestPondStatusString(t *testing.T) {
	var status PondStatus

	status = Active
	if len(status.String()) <= 0 {
		t.Error("Unexpectedly retrieved empty string from PondStatus object")
	}

	status = Stable
	if len(status.String()) <= 0 {
		t.Error("Unexpectedly retrieved empty string from PondStatus object")
	}

	status = Dead
	if len(status.String()) <= 0 {
		t.Error("Unexpectedly retrieved empty string from PondStatus object")
	}
}

func TestneighborSelectorString(t *testing.T) {
	var selector neighborsSelector

	selector = NEIGHBORS_ORTHOGONAL
	if len(selector.String()) <= 0 {
		t.Error("Unexpectedly retrieved empty string from pondselector object")
	}

	selector = NEIGHBORS_OBLIQUE
	if len(selector.String()) <= 0 {
		t.Error("Unexpectedly retrieved empty string from pondselector object")
	}

	selector = NEIGHBORS_ALL
	if len(selector.String()) <= 0 {
		t.Error("Unexpectedly retrieved empty string from PondStatus object")
	}
}

func TestlivingTrackerSetTest(t *testing.T) {
	tracker := newLivingTracker()

	// Ok, set the value
	loc := Location{X: 42, Y: 24}
	tracker.Set(loc)

	// Test that it exists
	if !tracker.Test(loc) {
		t.Fatal("Added location unexpectedly tested false")
	}
}

func TestlivingTrackerTestError(t *testing.T) {
	tracker := newLivingTracker()

	if tracker.Test(Location{X: 0, Y: 0}) {
		t.Error("Unexpectedly tested true a location that does not exist in structure")
	}
}

func TestlivingTrackerRemove(t *testing.T) {
	tracker := newLivingTracker()

	// Ok, set the value
	loc := Location{X: 42, Y: 24}
	tracker.Set(loc)

	tracker.Remove(loc)

	// Test that it exists
	if tracker.Test(loc) {
		t.Fatal("Unexpectedly tested true a location that was removed")
	}
}

func TestlivingTrackerRemoveError(t *testing.T) {
	t.Skip("TODO")
}

func TestlivingTrackerGetAll(t *testing.T) {
	tracker := newLivingTracker()

	// Create the expected values
	expectedLocations := make([]Location, 3)
	expectedLocations[0] = Location{X: 42, Y: 42}
	expectedLocations[1] = Location{X: 11, Y: 11}
	expectedLocations[2] = Location{X: 12, Y: 34}

	// Ok, set the values
	for _, l := range expectedLocations {
		tracker.Set(l)
	}

	actualLocations := tracker.GetAll()

	if len(actualLocations) != len(expectedLocations) {
		t.Fatalf("Received %d locations when I expected %d\n", len(actualLocations), len(expectedLocations))
	}

	// Check that all returned locations were expected
	for _, actual := range actualLocations {
		found := false
		for _, expected := range expectedLocations {
			if actual.Equals(&expected) {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("Returned location %s was not in list of test locations\n", actual.String())
			break
		}
	}

	// Check that all expected locations were returned
	for _, expected := range expectedLocations {
		found := false
		for _, actual := range actualLocations {
			if expected.Equals(&actual) {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("Expected location %s was not in list of returned locations\n", expected.String())
			break
		}
	}
}

func TestlivingTrackerCount(t *testing.T) {
	tracker := newLivingTracker()

	// Create the expected values
	expectedLocations := make([]Location, 3)
	expectedLocations[0] = Location{X: 42, Y: 42}
	expectedLocations[1] = Location{X: 11, Y: 11}
	expectedLocations[2] = Location{X: 12, Y: 34}

	// Ok, set the values
	for _, l := range expectedLocations {
		tracker.Set(l)
	}

	// Test the counter
	expectedCount := len(expectedLocations)
	count := tracker.GetCount()
	if count != expectedCount {
		t.Fatalf("Retrieved count of %d instead of expected %d\n", count, expectedCount)
	}

	// Now remove a location and try again
	tracker.Remove(expectedLocations[2])
	expectedCount--
	count = tracker.GetCount()
	if count != expectedCount {
		t.Fatalf("Retrieved count of %d instead of expected %d after remove a location\n", count, expectedCount)
	}
}

func TestpondSettingInitialPatterns(t *testing.T) {
	pond, err := newpond(Dimensions{Height: 3, Width: 3}, NEIGHBORS_ALL)
	if err != nil {
		t.Fatal("Unable to create pond")
	}

	// Create a pattern and call the pond's init function
	initialLiving := make([]Location, 3)
	initialLiving[0] = Location{X: 0, Y: 0}
	initialLiving[1] = Location{X: 1, Y: 1}
	initialLiving[2] = Location{X: 2, Y: 2}

	pond.SetOrganisms(initialLiving)

	// Check each expected value
	for _, loc := range initialLiving {
		if !pond.isOrganismAlive(loc) {
			t.Fatalf("Seed organism is not alive!: %s\n", loc.String())
		}
	}
}

func TestpondNeighborSelection(t *testing.T) {
	t.Skip("This will essentially just retest the board tests.")
}

func TestpondNeighborSelectionError(t *testing.T) {
	t.Skip("Bad location should be the test")
	/*
		fake_selector := 999
		pond := newpond(1, 1, fake_selector)

		neighbors, err := pond.GetNeighbors(Location{X: 0, Y: 0})
		if err != nil {
			t.Error("Did not encounter error when retrieving neighbors using bogus selector")
		}
	*/
}

func TestpondOrganismValue(t *testing.T) {
	expectedVal := 2
	pos := Location{X: 0, Y: 0}
	pond, err := newpond(Dimensions{Height: 1, Width: 1}, NEIGHBORS_ALL)
	if err != nil {
		t.Fatal("Unable to create pond")
	}
	pond.setOrganismValue(pos, expectedVal)

	actualVal := pond.GetOrganismValue(pos)

	if actualVal != expectedVal {
		t.Fatalf("Retrieved actual value %d instead of expected value %d\n", actualVal, expectedVal)
	}
}

func TestpondGetNumLiving(t *testing.T) {
	dims := Dimensions{Height: 3, Width: 3}
	pond, err := newpond(dims, NEIGHBORS_ALL)
	if err != nil {
		t.Fatal("Unable to create pond")
	}

	// Create a pattern and call the pond's init function
	initialLiving := make([]Location, 2)
	initialLiving[0] = Location{X: 0, Y: 0}
	initialLiving[1] = Location{X: 1, Y: 1}
	pond.SetOrganisms(initialLiving)

	numLiving := pond.GetNumLiving()

	if numLiving != len(initialLiving) {
		t.Fatalf("Retrieved actual %d living orgnamisms instead of expected %d\n", numLiving, len(initialLiving))
	}
}

func TestpondNeighborCountCalutation(t *testing.T) {
	dims := Dimensions{Height: 3, Width: 3}
	pond, err := newpond(dims, NEIGHBORS_ALL)
	if err != nil {
		t.Fatal("Unable to create pond")
	}

	// Create a pattern and call the pond's init function
	initialLiving := make([]Location, 2)
	initialLiving[0] = Location{X: 0, Y: 0}
	initialLiving[1] = Location{X: 1, Y: 1}
	pond.SetOrganisms(initialLiving)

	expectedNeighbors := make([]Location, 5)
	expectedNeighbors[0] = Location{X: 0, Y: 0}
	expectedNeighbors[1] = Location{X: 1, Y: 0}
	expectedNeighbors[2] = Location{X: 1, Y: 1}
	expectedNeighbors[3] = Location{X: 0, Y: 2}
	expectedNeighbors[4] = Location{X: 1, Y: 2}

	expectedNeighborCount := len(initialLiving)
	actualNeighborCount, actualNeighbors := pond.calculateNeighborCount(Location{X: 0, Y: 1})

	if actualNeighborCount != expectedNeighborCount {
		t.Fatalf("Retrieved %d neighbor count instead of expected %d\n", actualNeighborCount, expectedNeighborCount)
	}

	if len(actualNeighbors) != len(expectedNeighbors) {
		t.Fatalf("Retrieved %d neighbors instead of expected %d\n", len(actualNeighbors), expectedNeighborCount)
	}

	for _, neighbor := range actualNeighbors {
		found := false
		for _, expected := range expectedNeighbors {
			if expected.Equals(&neighbor) {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("Found unexpected neighbor %s\n", neighbor.String())
		}
	}
}

func TestpondString(t *testing.T) {
	dims := Dimensions{Height: 3, Width: 3}
	pond, err := newpond(dims, NEIGHBORS_ALL)
	if err != nil {
		t.Fatal("Unable to create pond")
	}
	if len(pond.String()) <= 0 {
		t.Error("Unexpectly retrieved empty string from pond string function")
	}
}
