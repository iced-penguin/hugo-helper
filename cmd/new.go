package cmd

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/icedpenguin0504/hugo-helper/model"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "create new article",
	Long:  "create new article",
	RunE: func(cmd *cobra.Command, args []string) error {
		section := viper.GetString(sectionKey)
		content := viper.GetString(contentKey)
		if err := createNewFile(section, content); err != nil {
			return err
		}
		return nil
	},
}

const (
	sectionKey = "directory.section"
	contentKey = "directory.content"
)

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().StringP("section", "s", "", "section is a directory in which new articles will be placed, under content")
	viper.BindPFlag(sectionKey, newCmd.Flags().Lookup("section"))
}

func createNewFile(section, content string) error {
	prompt := model.NewPrompt()

	baseFilename, category := prompt.Input()
	filename := fmt.Sprintf("%s/%s.md", section, baseFilename)

	out, err := exec.Command("hugo", "new", filename).CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to create new file: %v", err)
	}
	fmt.Println(string(out))

	absFilepath, err := filepath.Abs(fmt.Sprintf("./%s/%s", content, filename))
	if err != nil {
		return fmt.Errorf("failed to get absolute file path: %v", err)
	}

	article := model.NewArticle(absFilepath)

	if err := article.AddCategory(category); err != nil {
		return fmt.Errorf("failed to add category: %v", err)
	}

	return nil
}
