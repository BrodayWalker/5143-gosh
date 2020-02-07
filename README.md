# 5143-gosh
This is a repository for ``gosh``, the Go shell implementation for Project 1 in CMPS 5143 - Advanced Operating Systems. The project team members are Jeremy Glebe, Corbin Matamoros, and Broday Walker.

## How gosh Works
Currently, the ``gosh`` folder contains all ``.go`` files. ``gosh.go`` contains the ``main`` function, which is used to accept instructions and call the appropriate ``<command>.go`` file. <br>

## Naming Conventions
``gosh`` uses the standard Go naming conventions. Each supported command is contained in an appropriately named ``.go`` file. For example, the command ``mkdir`` is implemented in the ``mkdir.go`` file. <br>
 All exported commands should be capitalized and commented appropriately. <br>

Each ``.go`` file that contains executable code will have the ``package main`` declaration at the top of the file. <br>
#### Example ``.go`` File
```go
package main

import (
	"fmt"
)

// NamePrinter This function prints a name.
func  NamePrinter() {
	name := "Frank Sinatra"
	fmt.Printf("Hello, %s.\n", name)
}
```

The function ``NamePrinter`` will be visible and usable in other ``.go`` files in the folder without having to import anything.
## Building gosh
Build ``gosh`` from the ``gosh`` folder  with ``go build .`` The result will be an executable called ``gosh.exe`` .
## Running gosh
Run ``gosh`` with ``./gosh.exe`` .
