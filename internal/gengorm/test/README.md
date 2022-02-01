# About

The acceptance test suite.

A test case consists of:

- A `.proto` file that makes use of a feature
- The generated code (`make gen`)
- A test that asserts compilation, types, and so on

## Including `options.proto`

The `.proto` files in this test suite reference `gorm/options.proto` just like a user would. We just symlink them into the local tree for convenient feature development without having to copy files around.
