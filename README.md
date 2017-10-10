## Controller and Model logic control flow

The controller and model code is intended to be as thin as possible, with all of the heavy lifting done by the model base class that everything shares.

`models/base.go` is the file that is of interest here.  It contains all interface and struct definitions, and it exports the CRUD symbols that are called by the controllers:

- Create
- Read
- Update
- Delete
- List

There are also some other exported symbols that are used internally within the `models` package.

When a function is called by a controller, it will create a new model type and a context that it will use to perform tasks.

For instance, the model factory will return an instance of a model struct and then call that model's `GetCtx` method that will take care of returning the correct context for the pending CRUD operation.

(This returned context object is dependent upon the GOA context that is passed as a function argument into each generated controller.)

At this point, the called CRUD symbol will either delegate to the model for any more information it needs or handle the request on its own.

At present, each model **must** define the following methods:

- AllocateID
- GetCtx
- GetModelInstance
- GetModelCollection
- SetModel

These methods satisfy the `Model` interface defined in `base.go`.

