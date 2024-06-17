package pokeapi

import "testing"

func TestPokeApiCache(t *testing.T) {
	key := "location-areas"
	value := []byte("{\"count\":23}")
	expected := LocationAreaResponse{Count: 23}

	client := NewClient()
	client.cache.Set(key, value)
	res, err := client.FetchLocationAreas(&key)
	if err != nil {
		t.Error("expected no error")
	}
	if res.Count != expected.Count {
		t.Errorf("expected %q, got %q", expected.Count, res.Count)
	}
}

func TestPokeApiFetchLocationAreas(t *testing.T) {
	key := "location-areas"
	expected := LocationAreaResponse{Count: 23}

	client := NewClient()
	res, err := client.FetchLocationAreas(&key)
	if err != nil {
		t.Error("expected no error")
	}
	if res.Count != expected.Count {
		t.Errorf("expected %q, got %q", expected.Count, res.Count)
	}
}
