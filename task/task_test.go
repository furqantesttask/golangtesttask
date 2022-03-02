package task

import (
	"reflect"
	"testing"
)

type TestValidityStruct struct {
	stringSequence string
	result bool
}
func TestTestValidity(t *testing.T) {

	stories := []TestValidityStruct {
		{"23-ab-48-caba-56-haha", true},  
		{"123-43534-hello", false},            
		{"34534-ajdh-128379--63465-adhkjahsd", false},     
		{"576-bfndndf-2374898-akjsdnkas-8638-akjdlasjkd", true},   
		{"128039-ansdamsn-56475-aljdkjsld-656734-ajdkashd-", false}, 
	}

	var got, want bool

	for _, story := range stories {
		got = TestValidity(story.stringSequence)
		want = story.result
		if got != want {
			t.Errorf("got %t, wanted %t", got, want)
		}
	}
}
type TestAverageNumberStruct struct {
	text   string
	result float64
}
func TestAverageNumber(t *testing.T) {

	// Hard coded story sequence list for AverageNumber
	stories := []TestAverageNumberStruct {
		{"23-ab-48-caba-56-haha", 42.333333333333336},  
		{"123-43534-hello", 0.0},            
		{"34534-ajdh-128379--63465-adhkjahsd", 0.0},     
		{"576-bfndndf-2374898-akjsdnkas-8638-akjdlasjkd", 20.0},   
		{"128039-ansdamsn-56475-aljdkjsld-656734-ajdkashd-", 0.0},             
	}

	var got, want float64

	for _, story := range stories {
		got = AverageNumber(story.text)
		want = story.result
		if got != want {
			t.Errorf("got %f, wanted %f", got, want)
		}
	}
}

// Unit Test - WholeStory
type WholeStoryStruct struct {
	text, result string
}
func TestWholeStory(t *testing.T) {

	stories := []WholeStoryStruct {
		{"23-ab-48-caba-56-haha", "ab caba haha"},
		{"123-43534-hello", ""},                        
		{"34534-ajdh-128379--63465-adhkjahsd", ""},                 
		{"576-bfndndf-2374898-akjsdnkas-8638-akjdlasjkd", "ab cab haha"},    
		{"128039-ansdamsn-56475-aljdkjsld-656734-ajdkashd-", ""},             
	}

	var got, want string

	for _, story := range stories {
		got = WholeStory(story.text)
		want = story.result
		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}
type StoryStatsResultStruct struct {
	shortestSequence, longestSequence string
	averageWordLength      float64
	list              []string
}
type StoryStatsStruct struct {
	sequence   string
	result StoryStatsResultStruct
}
func TestStoryStats(t *testing.T) {

	stories := []StoryStatsStruct {
		{
			sequence: "23-ab-48-caba-56-haha", 
			result: StoryStatsResultStruct{
				shortestSequence: "ab", 
				longestSequence: "haha", 
				averageWordLength: 3.3333333333333335, 
				list: []string{"caba", "haha"},
			},
		},
		{
			sequence: "23-56-haha", 
			result: StoryStatsResultStruct{
				shortestSequence: "", 
				longestSequence: "", 
				averageWordLength: 0.0, list: nil,
			},
		},
		{
			sequence: "23-ab-48--56-haha", 
			result: StoryStatsResultStruct{
				shortestSequence: "", 
				longestSequence: "", 
				averageWordLength: 0.0, 
				list: nil,
			},
		},
		{
			sequence: "2-ab-48-cab-10-haha", 
			result: StoryStatsResultStruct{
				shortestSequence: "ab", 
				longestSequence: "haha", 
				averageWordLength: 3.0, 
				list: []string{"cab"},
			},
		},
		{
			sequence: "23-ab-48--56-haha-", 
			result: StoryStatsResultStruct{
				shortestSequence: "", 
				longestSequence: "", 
				averageWordLength: 0.0, list: nil,
			},
		},
	}


	var gotshort, gotlong, wantshort, wantlong string
	var gotavg, wantavg float64
	var gotlst, wantlst []string

	for _, story := range stories {
		gotshort, gotlong, gotavg, gotlst = StoryStats(story.sequence)

		wantshort, wantlong, wantavg, wantlst = story.result.shortestSequence, story.result.longestSequence, story.result.averageWordLength, story.result.list

		if gotshort != wantshort {
			t.Errorf("got %q, wanted %q", gotshort, wantshort)
		}
		if gotlong != wantlong {
			t.Errorf("got %q, wanted %q", gotlong, wantlong)
		}
		if float64(gotavg) != float64(wantavg) {
			t.Errorf("got %f, wanted %f", gotavg, wantavg)
		}
		if !reflect.DeepEqual(gotlst, wantlst) {
			t.Errorf("got %q, wanted %q", gotlst, wantlst)
		}
	}
}
