# hugo-helper

hugo-helper helps you build Hugo-generated blog.

## Install

```
go get github.com/icedpenguin0504/hugo-helper
```

## How to use

### Step 1: Prepare config file

Create `hugo-helper-config.yml` in your blog directory.

Example:

```
directory:
  content: content
  section: posts
```

### Step 2: Add new content file

Use `new` command to create a content file.

```
cd path/to/your/blog
hugo-helper new
```

This command creates the file interactively. The file will be placed in the directory specified in config file. (`content/section/`.) You can also use `section` flag to specify the section.

### Step 3: Write article

Write whatever you like in the file you created in step 2.

## Command details

```
$ hugo-hekper new -h
create new article

Usage:
  hugo-helper new [flags]

Flags:
  -h, --help             help for new
  -s, --section string   section is a directory in which new articles will be placed, under content
```
