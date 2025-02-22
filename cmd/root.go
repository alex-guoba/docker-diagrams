/*
* Copyright © 2025 alex.guoba <alex.guoba@gmail.com>
 */

package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/alex-guoba/docker-diagrams/pkg/image"
	"github.com/alex-guoba/docker-diagrams/pkg/node"
	"github.com/blushft/go-diagrams/diagram"

	"github.com/spf13/cobra"
)

var composeFile string
var envFiles []string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "docker-diagrams",
	Short: "Generate diagrams from a Docker Compose file.",
	Long: `A command - line tool that can parse Docker Compose Files 
and depict various service nodes and the relationships among them.
`,

	Run: func(cmd *cobra.Command, args []string) {
		var err error
		ctx := context.Background()

		project, err := image.LoadProject(ctx, composeFile, envFiles)
		if err != nil {
			log.Fatal("Load project failed: ", err)
		}

		// Create a new diagram
		projName := project.Name
		fmt.Println("Generating diagram for project: ", projName)
		canvas, err := diagram.New(diagram.Filename("docker-diagram"), diagram.Label(projName),
			diagram.Direction("LR"),
			diagram.WithAttribute("fontsize", "16"),
			diagram.WithAttribute("pad", "16"),
		)
		if err != nil {
			log.Fatal(err)
		}

		// Create a map to store nodes
		nodesMap := make(map[string]*diagram.Node)
		for _, service := range project.Services {
			// ignore service if it is ignored
			if strings.ToLower(service.Labels[image.TagIgnore]) == "true" {
				continue
			}
			if service.Name == "" {
				// A Compose file must declare a services top-level element as a map
				// whose keys are string representations of service names
				continue
			}
			label := service.ContainerName
			if label == "" {
				label = service.Name
			}

			node := node.ImageToNode(service.Name, service.Image, service.Labels[image.TagIcon]).Label(label)
			nodesMap[service.Name] = node
			canvas.Add(node) // Add node to canvas
		}

		// Add group info
		groupMap := make(map[string]*diagram.Group)
		for _, service := range project.Services {
			node, ok := nodesMap[service.Name]
			if !ok {
				continue
			}
			if service.Labels[image.TagGroup] == "" {
				continue
			}
			labels := strings.Split(service.Labels[image.TagGroup], ".")
			if len(labels) == 0 {
				continue
			}
			slices.Reverse(labels)

			var group, subGroup *diagram.Group
			for i, groupLabel := range labels {
				if group, ok = groupMap[groupLabel]; !ok {
					colorIdx := i
					group = diagram.NewGroup(groupLabel, diagram.IndexedBackground(colorIdx)).Label(groupLabel)
					groupMap[groupLabel] = group
				}

				if i == 0 {
					group.Add(node)
				} else {
					group.Group(subGroup)
				}
				subGroup = group
			}
			canvas.Group(group) // add root group to canvas
		}

		// Connect nodes
		for _, service := range project.Services {
			node, ok := nodesMap[service.Name]
			if !ok {
				continue
			}

			for name, _ := range service.DependsOn {
				if dep, ok := nodesMap[name]; ok {
					canvas.Connect(node, dep, diagram.Forward())
				}
			}
		}

		if err := canvas.Render(); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Diagram generated successfully in ./go-diagrams/")
		fmt.Println("Create an ouput image with any graphviz compatible renderer like:")
		fmt.Println("dot -Tpng docker-diagram.dot > diagram.png")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	rootCmd.PersistentFlags().StringVarP(&composeFile, "input", "i", "docker-compose.yml", "Path to the Docker Compose file")
	rootCmd.PersistentFlags().StringSliceVarP(&envFiles, "env", "e", []string{}, "Environment file to load, multiple files can be specified with multiple -f")
}
