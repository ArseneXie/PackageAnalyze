# AnalyzePackage
 Use go (a.k.a. golang) to establish a tool.
 Analyze a given PLSQL package, show all the dependent packages.


## Features

- the main method is regexp
- ignore PL/SQL comments
- It may not easy to define tables aliais and dependent package name, which both like <XXX>.<YYY>, use a parameter skipexp to do the exception handle.

  
## Run the Program
- parameters
    - srcfile: source file, the PLSQL package.
    - output: show the result
    - skipexp: skip pattern in regular expression
- example `analyzepkg -srcfile="R:/inv_qty_tree_pvt.pck" -output="R:/output2.txt" -skipexp="^(g_[a-z]*|x|mp|c1_rec|loc|lot|mil|mmtt|mln|mms|mpq|moqd|mr|msinv|mtlt|moq)."`
