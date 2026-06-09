// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package main

import (
	"os"
	"strings"
	"testing"
)

// conceptsDoc reads CONCEPTS.md once and returns its content.
// Failures here are fatal because all subtests depend on the file.
func readConceptsDoc(t *testing.T) string {
	t.Helper()
	data, err := os.ReadFile("CONCEPTS.md")
	if err != nil {
		t.Fatalf("CONCEPTS.md must exist and be readable: %v", err)
	}
	return string(data)
}

// TestConceptsDoc_FileExists verifies that CONCEPTS.md is present in the
// repository root and is non-empty.
func TestConceptsDoc_FileExists(t *testing.T) {
	info, err := os.Stat("CONCEPTS.md")
	if err != nil {
		t.Fatalf("CONCEPTS.md does not exist: %v", err)
	}
	if info.Size() == 0 {
		t.Fatal("CONCEPTS.md must not be empty")
	}
}

// TestConceptsDoc_Title verifies that the document has the expected top-level
// title as the very first line.
func TestConceptsDoc_Title(t *testing.T) {
	content := readConceptsDoc(t)
	lines := strings.SplitN(content, "\n", 2)
	if len(lines) == 0 {
		t.Fatal("CONCEPTS.md appears to be empty")
	}
	want := "# Terraform Concepts"
	if lines[0] != want {
		t.Errorf("first line: got %q, want %q", lines[0], want)
	}
}

// TestConceptsDoc_RequiredSections verifies that all seven numbered sections
// defined in the original document are present with their correct headings.
func TestConceptsDoc_RequiredSections(t *testing.T) {
	content := readConceptsDoc(t)

	required := []string{
		"## 1. Commands and Operations",
		"## 2. Backends and State Management",
		"## 3. Configuration Loader",
		"## 4. Graph Builder and Execution (DAG)",
		"## 5. Expression Evaluation",
		"## 6. Providers and Plugins",
		"## 7. Terraform Stacks",
	}

	for _, heading := range required {
		t.Run(heading, func(t *testing.T) {
			if !strings.Contains(content, heading) {
				t.Errorf("CONCEPTS.md is missing required section heading: %q", heading)
			}
		})
	}
}

// TestConceptsDoc_SectionCount verifies that the document contains exactly
// seven numbered second-level headings (## N. ...).
func TestConceptsDoc_SectionCount(t *testing.T) {
	content := readConceptsDoc(t)
	count := 0
	for _, line := range strings.Split(content, "\n") {
		// Match lines that look like "## <digit>. ..."
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "## ") && len(trimmed) > 3 {
			rest := trimmed[3:]
			if len(rest) > 0 && rest[0] >= '0' && rest[0] <= '9' {
				count++
			}
		}
	}
	if count != 7 {
		t.Errorf("expected 7 numbered sections, found %d", count)
	}
}

// TestConceptsDoc_CommandsSection verifies that the Commands and Operations
// section references the key packages and concepts it describes.
func TestConceptsDoc_CommandsSection(t *testing.T) {
	content := readConceptsDoc(t)

	terms := []string{
		"command",
		"terraform plan",
		"terraform apply",
		"Operation",
	}
	for _, term := range terms {
		t.Run(term, func(t *testing.T) {
			if !strings.Contains(content, term) {
				t.Errorf("Commands section must mention %q", term)
			}
		})
	}
}

// TestConceptsDoc_BackendsSection verifies that the Backends and State
// Management section references the key types and packages it describes.
func TestConceptsDoc_BackendsSection(t *testing.T) {
	content := readConceptsDoc(t)

	terms := []string{
		"statemgr",
		"local",
		"State Manager",
		"states.State",
	}
	for _, term := range terms {
		t.Run(term, func(t *testing.T) {
			if !strings.Contains(content, term) {
				t.Errorf("Backends section must mention %q", term)
			}
		})
	}
}

// TestConceptsDoc_ConfigLoaderSection verifies that the Configuration Loader
// section references the types and concepts it describes.
func TestConceptsDoc_ConfigLoaderSection(t *testing.T) {
	content := readConceptsDoc(t)

	terms := []string{
		"configload.Loader",
		"configs.Config",
		"hcl.Body",
		"hcl.Expression",
	}
	for _, term := range terms {
		t.Run(term, func(t *testing.T) {
			if !strings.Contains(content, term) {
				t.Errorf("Configuration Loader section must mention %q", term)
			}
		})
	}
}

// TestConceptsDoc_GraphSection verifies that the Graph Builder section
// contains descriptions of the DAG vertices, edges, and the graph walker.
func TestConceptsDoc_GraphSection(t *testing.T) {
	content := readConceptsDoc(t)

	terms := []string{
		"DAG",
		"Graph Builder",
		"Vertices",
		"Edges",
		"Graph Walk",
		"ContextGraphWalker",
		"terraform.EvalContext",
	}
	for _, term := range terms {
		t.Run(term, func(t *testing.T) {
			if !strings.Contains(content, term) {
				t.Errorf("Graph Builder section must mention %q", term)
			}
		})
	}
}

// TestConceptsDoc_GraphSection_BulletPoints verifies that the Graph Builder
// section uses bullet points for its three sub-concepts.
func TestConceptsDoc_GraphSection_BulletPoints(t *testing.T) {
	content := readConceptsDoc(t)

	bullets := []string{
		"**Vertices (Nodes):**",
		"**Edges:**",
		"**Graph Walk:**",
	}
	for _, bullet := range bullets {
		t.Run(bullet, func(t *testing.T) {
			if !strings.Contains(content, bullet) {
				t.Errorf("Graph Builder section must contain bullet %q", bullet)
			}
		})
	}
}

// TestConceptsDoc_ExpressionSection verifies that the Expression Evaluation
// section references the key types used for dynamic value resolution.
func TestConceptsDoc_ExpressionSection(t *testing.T) {
	content := readConceptsDoc(t)

	terms := []string{
		"hcl.Expression",
		"cty.Value",
		"count",
	}
	for _, term := range terms {
		t.Run(term, func(t *testing.T) {
			if !strings.Contains(content, term) {
				t.Errorf("Expression Evaluation section must mention %q", term)
			}
		})
	}
}

// TestConceptsDoc_ProvidersSection verifies that the Providers and Plugins
// section references the RPC plugin protocol and provider responsibilities.
func TestConceptsDoc_ProvidersSection(t *testing.T) {
	content := readConceptsDoc(t)

	terms := []string{
		"Providers",
		"RPC",
	}
	for _, term := range terms {
		t.Run(term, func(t *testing.T) {
			if !strings.Contains(content, term) {
				t.Errorf("Providers section must mention %q", term)
			}
		})
	}
}

// TestConceptsDoc_StacksSection verifies that the Terraform Stacks section
// references all four key packages that make up the stacks subsystem.
func TestConceptsDoc_StacksSection(t *testing.T) {
	content := readConceptsDoc(t)

	terms := []string{
		"stackconfig",
		"stackplan",
		"stackstate",
		"stackruntime",
	}
	for _, term := range terms {
		t.Run(term, func(t *testing.T) {
			if !strings.Contains(content, term) {
				t.Errorf("Terraform Stacks section must mention package %q", term)
			}
		})
	}
}

// TestConceptsDoc_StacksSection_BulletPoints verifies that the Stacks section
// uses bullet points for its package descriptions.
func TestConceptsDoc_StacksSection_BulletPoints(t *testing.T) {
	content := readConceptsDoc(t)

	bullets := []string{
		"`stackconfig`",
		"`stackplan`",
		"`stackstate`",
		"`stackruntime`",
	}
	for _, bullet := range bullets {
		t.Run(bullet, func(t *testing.T) {
			if !strings.Contains(content, bullet) {
				t.Errorf("Terraform Stacks section must contain backtick-quoted term %q", bullet)
			}
		})
	}
}

// TestConceptsDoc_SectionsOrder verifies that the seven sections appear in the
// correct numerical order within the document.
func TestConceptsDoc_SectionsOrder(t *testing.T) {
	content := readConceptsDoc(t)

	sections := []string{
		"## 1. Commands and Operations",
		"## 2. Backends and State Management",
		"## 3. Configuration Loader",
		"## 4. Graph Builder and Execution (DAG)",
		"## 5. Expression Evaluation",
		"## 6. Providers and Plugins",
		"## 7. Terraform Stacks",
	}

	pos := -1
	for _, section := range sections {
		idx := strings.Index(content, section)
		if idx == -1 {
			t.Errorf("section %q not found in CONCEPTS.md", section)
			continue
		}
		if idx <= pos {
			t.Errorf("section %q appears before its predecessor in CONCEPTS.md", section)
		}
		pos = idx
	}
}

// TestConceptsDoc_NoTrailingWhitespaceOnHeadings verifies that section
// headings do not have trailing whitespace, which can confuse markdown parsers.
func TestConceptsDoc_NoTrailingWhitespaceOnHeadings(t *testing.T) {
	content := readConceptsDoc(t)
	for i, line := range strings.Split(content, "\n") {
		if !strings.HasPrefix(line, "#") {
			continue
		}
		if line != strings.TrimRight(line, " \t") {
			t.Errorf("line %d has a heading with trailing whitespace: %q", i+1, line)
		}
	}
}

// TestConceptsDoc_TitleIsUniqueH1 verifies that the document has exactly one
// H1 heading (lines starting with a single "#" followed by a space).
func TestConceptsDoc_TitleIsUniqueH1(t *testing.T) {
	content := readConceptsDoc(t)
	count := 0
	for _, line := range strings.Split(content, "\n") {
		if strings.HasPrefix(line, "# ") {
			count++
		}
	}
	if count != 1 {
		t.Errorf("expected exactly 1 top-level H1 heading, found %d", count)
	}
}
