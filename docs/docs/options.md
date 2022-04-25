# Annotations

TODO

TODO: `.prop` notation vs `= { prop: ... }`.

TODO: Note that nothing is generated unless at least one message is flagged as a model.

## File Options

TODO

### model

Sets `model` for **all** messages in the file. [See `model` below](#model_1).

**Default:** `false`

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

option (gorm.file).model = true;
```

---

### validate

Sets `validate` for **all** messages in the file. [See `validate` below](#validate_1).

**Default:** `false`

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

option (gorm.file).validate = true;
```

---

### crud

Sets `crud` for **all** messages in the file. [See `crud` below](#crud_1).

**Default:** `false`

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

option (gorm.file).crud = true;
```

## Message Options

TODO

### model

TODO

**Default:** `false`

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;
}
```

---

### validate

TODO

**Default:** `false`

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).validate = true;
}
```

---

### crud

TODO

**Default:** `false`

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).crud = true;
}
```

---

### table

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message) = {
    model: true,
    table: "mytable"
  };
}
```

## Field Options

TODO

### column

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string my_field = 1 [
    (gorm.field).column = "my_column"
  ];
}
```

---

### not_null

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string my_field = 1 [
    (gorm.field).not_null = true
  ];
}
```

---

### default

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string my_field = 1 [
    (gorm.field).default = "a default value"
  ];
}
```

---

### unique

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string my_field = 1 [
    (gorm.field).unique = true
  ];
}
```

---

### primary_key

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string uuid = 1 [
    (gorm.field).primary_key = true
  ];
}
```

---

### index

TODO

#### default

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string my_field = 1 [
    (gorm.field).index = {default: true}
  ];
}
```

#### name

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string my_field = 1 [
    (gorm.field).index = {name: "my_index_name"}
  ];
}
```

---

### unique_index

TODO

#### default

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string my_field = 1 [
    (gorm.field).unique_index = {default: true}
  ];
}
```

#### name

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string my_field = 1 [
    (gorm.field).unique_index = {name: "my_index_name"}
  ];
}
```

---

### auto_create_time

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string my_field = 1 [
    (gorm.field).auto_create_time = true
  ];
}
```

---

### auto_update_time

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string my_field = 1 [
    (gorm.field).auto_update_time = true
  ];
}
```

---

### permissions

TODO

#### ignore

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string my_ignored_field = 1 [
    (gorm.field).ignore = true
  ];
}
```

#### deny

TODO

##### create

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string my_field = 1 [
    (gorm.field).deny = {create: true}
  ];
}
```

##### update

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string my_field = 1 [
    (gorm.field).deny = {update: true}
  ];
}
```

##### read

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string my_field = 1 [
    (gorm.field).deny = {read: true}
  ];
}
```

---

### json

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  map<string, string> my_map = 1 [
    (gorm.field).json = true
  ];
}
```
