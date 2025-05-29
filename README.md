fibonacci-service
is just a project for exploring full-stack development with React and golang

TODOs
- Setup github workflow/s
- Setup dependabot
- Setup auto-merge for dependabot merges (for minor updates and patches only)
  - test dependabot automerge
- Develop front-end: visualiser of fibonacci sequences using d3.js
- Re-organise the backend
  - interfaces
    fibonacci service type definitions
  - implementations
    - data folder
    - Go map implementation
    - experiment with in-memory databases e.g. redis, or https://github.com/hashicorp/go-memdb