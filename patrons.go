package oembed

import (
	"log"
	"sort"
	"strings"
)

func init() {
	patrons := []string{"Aurielb", "MoD21k", "Yami Odymel"}
	sort.Strings(patrons)
	log.Print(
		"Support toby3d on Patreon: https://patreon.com/bePatron?c=243288", "\n",
		"The current version of oembed is sponsored by: ",
		strings.Join(patrons[:len(patrons)-1], ", "), " and ", patrons[len(patrons)-1],
	)
}
