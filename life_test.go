package life

import (
	"testing"
	"time"
)

func TestLifeStatsString(t *testing.T) {
	stats := new(LifeStats)
	stats.Generations++

	if len(stats.String()) <= 0 {
		t.Error("String function unexpectly returned an empty string")
	}
}

func TestLifeCreation(t *testing.T) {
	dims := LifeboardDims{Height: 3, Width: 3}
	strategy, err := NewLife("TestLifeCreation",
		dims,
		NEIGHBORS_ALL,
		Blinkers,
		GetConwayTester(),
		SimultaneousProcessor)
	if err != nil {
		t.Fatalf("Unable to create strategy: %s\n", err)
	}

	expected, _ := newLifeboard(LifeboardDims{3, 3})
	expected.SetValue(LifeboardLocation{X: 0, Y: 1}, 0)
	expected.SetValue(LifeboardLocation{X: 1, Y: 1}, 0)
	expected.SetValue(LifeboardLocation{X: 2, Y: 1}, 0)

	if !strategy.pond.board.Equals(expected) {
		t.Fatalf("Actual board\n%s\ndoes not match expected\n%s\n", strategy.pond.board.String(), expected.String())
	}
}

func TestLifeProcess(t *testing.T) {
	dims := LifeboardDims{Height: 3, Width: 3}
	strategy, err := NewLife("TestLifeProcess",
		dims,
		NEIGHBORS_ALL,
		Blinkers,
		GetConwayTester(),
		SimultaneousProcessor)
	if err != nil {
		t.Fatalf("Unable to create strategy: %s\n", err)
	}

	seededpond, err := strategy.pond.Clone()
	if err != nil {
		t.Fatalf("Unable to clone pond: %s\n", err)
	}

	strategy.process()

	processedpond := strategy.pond

	if seededpond.Equals(processedpond) {
		t.Fatal("pond did not change after processing")
	}

	// Check statistics
	if strategy.Statistics.Generations != 1 {
		t.Errorf("Tracked %d generations when there should only be one\n", strategy.Statistics.Generations)
	}
}

func TestLifeStartStop(t *testing.T) {
	t.Skip("This doesn't work as expected")
	dims := LifeboardDims{Height: 3, Width: 3}
	strategy, err := NewLife("TestLifeStartStop",
		dims,
		NEIGHBORS_ALL,
		Blinkers,
		GetConwayTester(),
		SimultaneousProcessor)
	if err != nil {
		t.Fatalf("Unable to create strategy: %s\n", err)
	}

	seededpond, err := strategy.pond.Clone()
	if err != nil {
		t.Fatalf("Unable to clone pond: %s\n", err)
	}

	updates := make(chan bool)
	strategy.Start(updates)

	// go func() {
	time.Sleep(strategy.UpdateRate * 1)
	strategy.Stop()
	/*}()

	for {
		select {
		case <-updates:
			t.Log(strategy.String())
		}
	}
	*/

	processedpond := strategy.pond

	if seededpond.Equals(processedpond) {
		t.Fatal("pond did not change after processing")
	}

	// Check statistics
	if strategy.Statistics.Generations < 2 {
		t.Errorf("Tracked %d generations when there should be two or more\n", strategy.Statistics.Generations)
	}
}

func TestLifeGetGeneration(t *testing.T) {
	dims := LifeboardDims{Height: 3, Width: 3}
	strategy, err := NewLife("TestLifeString",
		dims,
		NEIGHBORS_ALL,
		Blinkers,
		GetConwayTester(),
		SimultaneousProcessor)
	if err != nil {
		t.Fatalf("Unable to create strategy: %s\n", err)
	}

	expectedNumLiving := 3
	expectedGen := 31
	gen := strategy.GetGeneration(expectedGen)

	if gen.Num != expectedGen {
		t.Error("Retrieved %d generations instead of %d\n", gen.Num, expectedGen)
	}

	if len(gen.Living) != expectedNumLiving {
		t.Fatalf("Retrieved %d living organisms instead of %d\n", len(gen.Living), expectedNumLiving)
	}
	// TODO
}

func TestLifeString(t *testing.T) {
	dims := LifeboardDims{Height: 3, Width: 3}
	strategy, err := NewLife("TestLifeString",
		dims,
		NEIGHBORS_ALL,
		Blinkers,
		GetConwayTester(),
		SimultaneousProcessor)
	if err != nil {
		t.Fatalf("Unable to create strategy: %s\n", err)
	}

	if len(strategy.String()) <= 0 {
		t.Error("String function unexpectly returned an empty string")
	}
}