# How to use pongo2 based templating ?

[pongo2](https://github.com/flosch/pongo2) implements the [Django template language v1.7](https://django.readthedocs.io/en/1.7.x/topics/templates.html#id1). However, the library also includes features such as `macro`, `set` as found in [Jinja](https://jinja.palletsprojects.com/en/3.1.x/templates/). 

## Available patterns for re-use

Organising a template into its constituent parts is an interesting problem to solve when authoring templates. We consider three primary mechanisms that pongo2 provides in this regard.

1. `include`
2. `block` / `extends`
3. `macro`

### `include`

This is simple composition; where rendering of certain parts of the document is delegated to other files.

```
# include_main.j2
{% for name in names -%}
- {% include "include_snippet.j2" %}
{% endfor %}

# include_snippet.j2
name is {{ name }}

# output
- name is apple
- name is orange
```

The include mechanism allows:
- passing parameters. See [examples](https://github.com/flosch/pongo2/blob/master/template_tests/includes.tpl).
- include target to be dynamic such as a variable or an expression


### `block` / `extends`

In block based inheritance, a template inherits the contents of its parent and further customises the blocks defined there. A template can extend utmost one parent. 

Perhaps, one way to easily remember the semantics of `block` / `extends` is to think of it analogously to class based inheritance in Java. A subclass may override (redefine) methods in the parent class. Those methods it does not override, it inherits from the parent. Further, the subclass may define new methods for its children to override.

In the same way, the inheriting template may leave certain blocks with their original contents, while replace others with new definitions (which themselves may introduce new blocks). When redefining a block, `block.Super` can be used to include the original contents of the block.


```
# block_base.j2
{% for name in names -%}
- {% block fruit_item %}default{% endblock %}
{% endfor %}

# block_main.j2
{% extends "block_base.j2" %}
{% block fruit_item -%} name is {{ name }}{% endblock %}

# output
- name is apple
- name is orange
```

### `macro`

A macro is similar to an include in that it allows composing a template from sub components. However, unlike `include`, a macro invocation returns a value which may then be further processed before rendering. 

Further, an `include` is necessarily file oriented where the entire file contents is included. On the other hand, more than one macro may be defined in a file and used in the same or different file (through `import`/`export`).

```
# macro_main.j2
{% import "macro_snippet.j2" item -%}
{% for name in names -%}
- {{ item(name) }}
{% endfor %}

# macro_snippet.j2
{% macro item(name, details) export %}{{name}}{% endmacro %}
```

## How to decide which pattern to use ?

- `block`/`extends`: Inheritance, models "Is-a" relationship
- `include`, `macro`: Composition, models "Has-a" relationship

## Integrating `pongo2` in a Go application

```go
package main

import (
	"bytes"
	"fmt"

	pongo2 "github.com/flosch/pongo2/v6"
)

var templateSet = pongo2.NewSet("main", pongo2.MustNewLocalFileSystemLoader("templates/"))

func render() error {

    t, err := templateSet.FromFile("main.j2")

    if err != nil {
        return err
    }

	var b bytes.Buffer

	context := pongo2.Context{
        // values
		"x":     "y",
        // functions
        "f": func() {
            return "some function"
        },
	}
	err := t.ExecuteWriterUnbuffered(context, &b)
	if err != nil {
		return err
	}
	fmt.Print(b.String())
	return nil
}
```

## References

- [pongo2's tests](https://github.com/flosch/pongo2/blob/master/template_tests/) are good examples of what's possible
