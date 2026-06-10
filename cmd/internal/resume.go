package cmds

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/hairyhenderson/gomplate/v5"
	"github.com/spf13/cobra"
	"go.yaml.in/yaml/v3"
)

var Command = &cobra.Command{
	Use:   "resume",
	Short: "Manages resume build and rendering.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if !FileExists("data/resume.yaml") {
			return fmt.Errorf("data/resume.yaml does not exist. Please create it first.")
		}
		return nil
	},
}

var renderCommand = &cobra.Command{
	Use:   "render",
	Short: "Renders the resume to the output directory.",
	RunE:  renderRunFn,
}

var printCommand = &cobra.Command{
	Use:   "print",
	Short: "Prints the resume datafile in the terminal.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fp, err := os.Open("data/resume.yaml")
		if err != nil {
			return err
		}
		defer fp.Close()

		sc := bufio.NewScanner(fp)
		for sc.Scan() {
			fmt.Println(sc.Text())
		}
		if err := sc.Err(); err != nil {
			return err
		}
		return nil
	},
}

var publishCommand = &cobra.Command{
	Use:   "publish",
	Short: "Publishes the resume on the public folder.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !FileExists("templates/resume/resume.pdf") {
			err := renderRunFn(cmd, args)
			if err != nil {
				return err
			}
		}

		src, err := os.Open("templates/resume/resume.pdf")
		if err != nil {
			return fmt.Errorf("failed to open templates/resume/resume.pdf: %w", err)
		}
		defer src.Close()

		// ensure public dir exists
		err = os.MkdirAll("public", 0755)
		if err != nil {
			return fmt.Errorf("failed to create public dir: %w", err)
		}

		dst, err := os.Create("public/resume.pdf")
		if err != nil {
			return fmt.Errorf("failed to create public/resume.pdf: %w", err)
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			return fmt.Errorf("failed to copy resume.pdf: %w", err)
		}
		return nil
	},
}

func renderRunFn(cmd *cobra.Command, args []string) error {
	err := EnsureBinary("tectonic")
	if err != nil {
		return err
	}

	err = renderTemplate()
	if err != nil {
		return fmt.Errorf("failed to render template: %w", err)
	}

	err = renderLatex()
	if err != nil {
		return fmt.Errorf("failed to render latex: %w", err)
	}

	return nil
}

func renderTemplate() error {
	const filename = "templates/resume/resume.tex.tmpl"
	funcs := gomplate.CreateFuncs(context.Background())

	fp, err := os.Create("templates/resume/resume.tex")
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer fp.Close()

	tmpl, err := template.New(filepath.Base(filename)).Funcs(funcs).ParseFiles(filename)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	data, err := getData()
	if err != nil {
		return fmt.Errorf("failed to get data: %w", err)
	}

	err = tmpl.Execute(fp, data)
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}
	return nil
}

func getData() (map[string]any, error) {
	var data map[string]any
	fp, err := os.Open("data/resume.yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to open data file: %w", err)
	}
	defer fp.Close()
	err = yaml.NewDecoder(fp).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode data: %w", err)
	}
	return data, nil
}

func renderLatex() error {
	cmd := exec.Command("tectonic", "templates/resume/resume.tex")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func init() {
	Command.AddCommand(renderCommand)
	Command.AddCommand(printCommand)
	Command.AddCommand(publishCommand)
}
