package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/icedpenguin0504/hugo-helper/model"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check articles",
	Long:  "Check articles before publishing",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := check(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

func check() error {
	// 全記事ファイル取得
	var filenames []string
	content := viper.GetString(contentKey)
	err := filepath.Walk(content, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			filenames = append(filenames, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	var warnings []string
	for _, filename := range filenames {
		// draftチェック
		draft, err := isDraft(filename)
		if err != nil {
			return err
		}
		if draft {
			warnings = append(warnings, fmt.Sprintf("warning: %s is draft", filename))
		}
	}

	for _, warning := range warnings {
		fmt.Println(warning)
	}
	return nil
}

func isDraft(filename string) (bool, error) {
	absFilepath, err := filepath.Abs(filename)
	if err != nil {
		return false, fmt.Errorf("failed to get absolute path: %v", err)
	}
	article := model.NewArticle(absFilepath)
	metadata, err := article.ReadMetadata()
	if err != nil {
		return false, err
	}
	return metadata.Draft, nil
}
