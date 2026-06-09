# Terraform Concepts

This document highlights the most important architectural concepts and subsystems in the Terraform Core codebase.

## 1. Commands and Operations
Execution begins in the CLI (the `command` package). When a user runs a command like `terraform plan` or `terraform apply`, it reads command-line arguments and constructs an `Operation`. An operation represents an action to be taken, a workspace, variables, and path information. The operation is then passed to a backend to execute.

## 2. Backends and State Management
A **backend** determines where Terraform stores its state snapshots. Most backends simply implement state storage (`statemgr`), while the default `local` backend executes operations. The **State Manager** is responsible for reading and writing snapshots of the Terraform state (as a `states.State` object) to local disk or remote network services.

## 3. Configuration Loader
The `configload.Loader` parses configuration files and recursively handles child modules to construct a full `configs.Config` object representing the infrastructure. Because Terraform's DSL relies on dynamic evaluation, parts of the configuration that depend on un-evaluated data remain as `hcl.Body` and `hcl.Expression` nodes until graph evaluation.

## 4. Graph Builder and Execution (DAG)
Terraform builds a Directed Acyclic Graph (DAG) representing operations using a **Graph Builder**.
- **Vertices (Nodes):** Represent specific configuration objects, such as `resource` blocks.
- **Edges:** Represent "happens after" dependencies between resources.
- **Graph Walk:** `ContextGraphWalker` traverses the graph. It evaluates vertices concurrently, respecting edges. Each vertex executes its logic (e.g., retrieving instance state, diffing configuration, or interacting with providers) using a `terraform.EvalContext`.

## 5. Expression Evaluation
While walking the graph, Terraform evaluates HCL expressions (`hcl.Expression`) into specific dynamic values (`cty.Value`). This resolves interpolations, variables, and function calls dynamically as the dependent data becomes available during graph execution. Sometimes, a node uses dynamic expansion (like the `count` meta-argument) to construct sub-graphs dynamically.

## 6. Providers and Plugins
Terraform offloads interaction with specific external APIs to **Providers**. Core communicates with providers through an RPC API plugin wire protocol. Providers handle creating instance diffs and executing operations (e.g., creation or updating of infrastructure).

## 7. Terraform Stacks
Terraform Stacks provides an orchestration layer built on top of multiple trees of Terraform modules. It contains analog packages to Core concepts for stacks features:
- `stackconfig` for parsing stacks language.
- `stackplan` and `stackstate` for maintaining stacks versions of plan and state models.
- `stackruntime` for evaluating, creating plans, and applying operations across multiple components.
