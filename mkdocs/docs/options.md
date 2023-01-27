# Options

Options control what `protoc-gen-gorm` does. You set them in your `.proto` files, as regular  
[Protocol Buffer Options](https://developers.google.com/protocol-buffers/docs/proto3#options).

The plugin does nothing by default, you'll have to flag some of your messages to be models first, e.g. set [`model`](#model_1) to `true`.

## File Options

File options apply to all message types within the `.proto` file.

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

Implies `model = true` when set to `true`.

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

Implies `model = true` when set to `true`.

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

option (gorm.file).crud = true;
```

## Message Options

Message options control generation of model and supporting code for your message types.

### model

Marks a message as a _model_ so `protoc-gen-gorm` generates a Go struct and converter methods for use with GORM v2.

The struct type name is the message's name with "Model" appended.

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

Generates:

```go
package mypackage

type MyMessageModel struct{ /* ... */ }

func (m *MyMessageModel) ToProto() (*MyMessage, error) { /* ... */ }
func (p *MyMessage) ToModel() (*MyMessageModel, error) { /* ... */ }
```

---

### validate

**TODO**

**Default:** `false`

Implies `model = true` when set to `true`.

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

Generates supporting types and methods to implement [CRUD](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete) for your model.

**Default:** `false`

Implies `model = true` when set to `true`.

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).crud = true;
}
```

Generates:

```go
package mypackage

type MyMessageWithDB struct{ /* ... */ }
type CrudGetOption
type CrudListOption

// Attach a GORM DB handle to your message.
func (p *MyMessage) WithDB(db *gorm.DB) MyMessageWithDB

// CRUD support without need to convert to model type and back.
func (c MyMessageWithDB) Create(context.Context) (*MyMessage, error)
func (c MyMessageWithDB) Get(context.Context, opts ...MyMessageGetOption) (*MyMessage, error)
func (c MyMessageWithDB) List(context.Context, opts ...MyMessageListOption) ([]*MyMessage, error)
func (c MyMessageWithDB) Update(context.Context) (*MyMessage, error)
func (c MyMessageWithDB) Patch(context.Context, mask *fieldmaskpb.FieldMask) error
func (c MyMessageWithDB) Delete(context.Context) error 
```

---

### table

Set the table name for models of this type.

**Default:** Unset, uses the [GORM default](https://gorm.io/docs/conventions.html#Pluralized-Table-Name).

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

The generated struct now implements [GORM's Tabler interface](https://pkg.go.dev/gorm.io/gorm/schema#Tabler):

```go
package mypackage

type MyMessageModel struct {
	// ...
}

func (m *MyMessageModel) TableName() string {
	return "mytable"
}
```

## Field Options

Field options refine how your generated model works with GORM through struct field tags and supporting code.

### column

Sets the [database column name](https://gorm.io/docs/conventions.html#Column-Name).

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

Equivalent GORM struct field tag:

```go
package mypackage

type MyMessageModel struct {
	MyField string `gorm:"column:my_column"`
}
```

---

### not_null

Specifies the field's column as "NOT NULL". See "not null" under [GORM: Field Tags](https://gorm.io/docs/models.html#Fields-Tags).

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

Equivalent GORM struct field tag:

```go
package mypackage

type MyMessageModel struct {
	MyField string `gorm:"not null"`
}
```

---

### default

Sets the default value for the field's column. See "default" under [GORM: Field Tags](https://gorm.io/docs/models.html#Fields-Tags).

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

Equivalent GORM struct field tag:

```go
package mypackage

type MyMessageModel struct {
	MyField string `gorm:"default:\"a default value\""`
}
```

---

### unique

Flags the field's column to be indexed with a unique indep. See [GORM: Indexes](https://gorm.io/docs/indexes.html#uniqueIndex).

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

Equivalent GORM struct field tag:

```go
package mypackage

type MyMessageModel struct {
	MyField string `gorm:"uniqueIndex"`
}
```

---

### primary_key

Makes the field a primary key.

Also see:

- [GORM: ID as Primary Key](https://gorm.io/docs/conventions.html#ID-as-Primary-Key)
- [GORM: Composite Primary Key](https://gorm.io/docs/composite_primary_key.html#content-inner)

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

Equivalent GORM struct field tag:

```go
package mypackage

type MyMessageModel struct {
	Uuid string `gorm:"primaryKey"`
}
```

---

### index

Adds an index to a field. [Composite](https://gorm.io/docs/indexes.html#Composite-Indexes) and [multiple indexes](https://gorm.io/docs/indexes.html#Multiple-indexes) are possible.

#### default

Use defaults for the index, e.g. name, type, etc.

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

Equivalent GORM struct field tag:

```go
package mypackage

type MyMessageModel struct {
	MyField string `gorm:"index"`
}
```

#### name

Gives the index a custom name.

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

Equivalent GORM struct field tag:

```go
package mypackage

type MyMessageModel struct {
	MyField string `gorm:"index:my_index_name"`
}
```

---

### unique_index

Same as [`index`](#index) above except that the index is unique.

#### default

Use defaults for the unique index, e.g. name, type, etc.

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

Equivalent GORM struct field tag:

```go
package mypackage

type MyMessageModel struct {
	MyField string `gorm:"uniqueIndex"`
}
```

#### name

Gives the unique index a custom name.

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

Equivalent GORM struct field tag:

```go
package mypackage

type MyMessageModel struct {
	MyField string `gorm:"uniqueIndex:my_index_name"`
}
```

---

### auto_create_time

Instructs GORM to [track creation time](https://gorm.io/docs/models.html#Creating-x2F-Updating-Time-x2F-Unix-Milli-x2F-Nano-Seconds-Tracking) in the flagged field.

**Example:**

```protobuf
syntax = "proto3";
import "google/protobuf/timestamp.proto";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  google.protobuf.Timestamp my_time = 1 [
    (gorm.field).auto_create_time = true
  ];
}
```

Equivalent GORM struct field tag:

```go
package mypackage

import (
	"database/sql"
)

type MyMessageModel struct {
	MyTime sql.NullTime `gorm:"autoCreateTime"`
}
```

---

### auto_update_time

Instructs GORM to [track update time](https://gorm.io/docs/models.html#Creating-x2F-Updating-Time-x2F-Unix-Milli-x2F-Nano-Seconds-Tracking) in the flagged field.

**Example:**

```protobuf
syntax = "proto3";
import "google/protobuf/timestamp.proto";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  google.protobuf.Timestamp my_time = 1 [
    (gorm.field).auto_update_time = true
  ];
}
```

Equivalent GORM struct field tag:

```go
package mypackage

import (
	"database/sql"
)

type MyMessageModel struct {
	MyTime sql.NullTime `gorm:"autoUpdateTime"`
}
```

---

### permissions

Sets the [field level permissions](https://gorm.io/docs/models.html#Field-Level-Permission) to turn columns into read-only, write-only, create-only, update-only or to ignore a column entirely.

#### ignore

Ignores the column entirely.

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

Equivalent GORM struct field tag:

```go
package mypackage

type MyMessageModel struct {
	MyField string `gorm:"-"`
}
```

#### deny

Restricts access to a field. Multiple "denys" can be combined to the desired effect.

##### create

Prevent creation, still allows reads and updates.

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

Equivalent GORM struct field tag:

```go
package mypackage

type MyMessageModel struct {
	MyField string `gorm:"<-:update"`
}
```

##### update

Prevent updates, still allows creation and reads.

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

Equivalent GORM struct field tag:

```go
package mypackage

type MyMessageModel struct {
	MyField string `gorm:"<-:create"`
}
```

##### read

Prevent reads, still allows writes.

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

Equivalent GORM struct field tag:

```go
package mypackage

type MyMessageModel struct {
	MyField string `gorm:"->:false;<-"`
}
```

---

### json

Encode and decode the field as JSON strings.

The converter methods, `MyMessageModel.ToProto()` and `MyMessage.ToModel()` in this case, call `json.Unmarshal` and `json.Marshal` respectively to decode the field's contents.

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
