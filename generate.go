package main

import (
	"github.com/solo-io/solo-kit/pkg/code-generator/cmd"
	"github.com/solo-io/solo-kit/pkg/code-generator/docgen/options"
	"github.com/solo-io/solo-kit/pkg/utils/log"
	"github.com/solo-io/sqoop/version"
)

func main() {
	err := version.CheckVersions()
	if err != nil {
		log.Fatalf("generate failed!: %v", err)
	}
	log.Printf("starting generate")
	docsOpts := cmd.DocsOptions{
		Output: options.Hugo,
	}
	if err := cmd.Run(".", true, &docsOpts, nil, nil); err != nil {
		log.Fatalf("generate failed!: %v", err)
	}
}
