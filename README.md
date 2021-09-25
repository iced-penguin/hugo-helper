# hugo-helper

hugo-helper helps you build Hugo-generated blog.

## Install

```
go get　-u github.com/icedpenguin0504/hugo-helper
```

## How to use

### Step 1: Prepare config file

Create `hugo-helper.yml` in your blog directory.

Example:

```
directory:
  # Articles will be placed in content/section by default
  content: content
  section: posts
taxonomy:
  # You can use these categories when you create new articles
  categories:
    - Programming
    - Books
```

### Step 2: Add new content file

Use `new` command to create a content file.

```
cd path/to/your/blog
hugo-helper new
```

This command creates the file interactively. The file will be placed in the directory specified in config file. (`content/section/`.) You can also use `section` flag to specify the section. For details, run `hugo-helper new -h`.

### Step 3: Write article

Write whatever you like in the file you created in step 2.
