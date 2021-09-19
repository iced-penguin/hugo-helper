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
	Use: "new",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := createNewFile(); err != nil {
			return err
		}
		return nil
	},
}

const (
	sectionKey = "directory.section"
	contentKey = "directory.content"
)

func createNewFile() error {
	prompt := model.NewPrompt()

	baseFilename, category := prompt.Input()
	section := viper.GetString(sectionKey)
	filename := fmt.Sprintf("%s/%s.md", section, baseFilename)

	out, err := exec.Command("hugo", "new", filename).CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to create new file: %v", err)
	}
	fmt.Println(string(out))

	contentDir := viper.GetString(contentKey)
	absFilepath, err := filepath.Abs(fmt.Sprintf("./%s/%s", contentDir, filename))
	if err != nil {
		return fmt.Errorf("failed to get absolute file path: %v", err)
	}

	article := model.NewArticle(absFilepath)

	if err := article.AddCategory(category); err != nil {
		return fmt.Errorf("failed to add category: %v", err)
	}

	return nil
}
