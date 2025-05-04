package service

import (
	"testing"
)

func TestSortByDomain(t *testing.T) {
	tests := []struct {
		name         string
		domainCounts map[string]int
		want         []DomainCount
	}{
		{
			name: "Basic sorting",
			domainCounts: map[string]int{
				"zebra.com":  1,
				"apple.com":  2,
				"banana.com": 1,
			},
			want: []DomainCount{
				{Domain: "apple.com", Count: 2},
				{Domain: "banana.com", Count: 1},
				{Domain: "zebra.com", Count: 1},
			},
		},
		{
			name: "Empty map",
			domainCounts: map[string]int{},
			want: []DomainCount{},
		},
		{
			name: "Single domain",
			domainCounts: map[string]int{
				"example.com": 5,
			},
			want: []DomainCount{
				{Domain: "example.com", Count: 5},
			},
		},
		{
			name: "Case sensitive sorting",
			domainCounts: map[string]int{
				"Zebra.com": 1,
				"apple.com": 2,
				"Banana.com": 1,
			},
			want: []DomainCount{
				{Domain: "Banana.com", Count: 1},
				{Domain: "Zebra.com", Count: 1},
				{Domain: "apple.com", Count: 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SortByDomain(tt.domainCounts)
			if len(got) != len(tt.want) {
				t.Errorf("SortByDomain() length = %v, want %v", len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i].Domain != tt.want[i].Domain || got[i].Count != tt.want[i].Count {
					t.Errorf("SortByDomain()[%d] = %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestSortByCount(t *testing.T) {
	tests := []struct {
		name         string
		domainCounts map[string]int
		want         []DomainCount
	}{
		{
			name: "Basic sorting by count",
			domainCounts: map[string]int{
				"zebra.com":  1,
				"apple.com":  3,
				"banana.com": 2,
			},
			want: []DomainCount{
				{Domain: "apple.com", Count: 3},
				{Domain: "banana.com", Count: 2},
				{Domain: "zebra.com", Count: 1},
			},
		},
		{
			name: "Equal counts sort by domain",
			domainCounts: map[string]int{
				"zebra.com":  2,
				"apple.com":  2,
				"banana.com": 2,
			},
			want: []DomainCount{
				{Domain: "apple.com", Count: 2},
				{Domain: "banana.com", Count: 2},
				{Domain: "zebra.com", Count: 2},
			},
		},
		{
			name: "Empty map",
			domainCounts: map[string]int{},
			want: []DomainCount{},
		},
		{
			name: "Single domain",
			domainCounts: map[string]int{
				"example.com": 5,
			},
			want: []DomainCount{
				{Domain: "example.com", Count: 5},
			},
		},
		{
			name: "Large numbers",
			domainCounts: map[string]int{
				"zebra.com":  1000,
				"apple.com":  5000,
				"banana.com": 2000,
			},
			want: []DomainCount{
				{Domain: "apple.com", Count: 5000},
				{Domain: "banana.com", Count: 2000},
				{Domain: "zebra.com", Count: 1000},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SortByCount(tt.domainCounts)
			if len(got) != len(tt.want) {
				t.Errorf("SortByCount() length = %v, want %v", len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i].Domain != tt.want[i].Domain || got[i].Count != tt.want[i].Count {
					t.Errorf("SortByCount()[%d] = %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
} 