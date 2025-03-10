# JSON-RPC interface

Delve exposes a [JSON-RPC](https://www.jsonrpc.org/specification_v1) API interface. 

Note that this JSON-RPC interface is served over a streaming socket, *not* over HTTP.

# API versions

Delve currently only supports v2 of its API. Support for v1 was dropped in version v1.24.0 of Delve.

# API version 2 documentation

All the methods of the type `service/rpc2.RPCServer` can be called using JSON-RPC, the documentation for these calls is [available on godoc](https://pkg.go.dev/github.com/fansqz/delve/service/rpc2#RPCServer).

Note that all exposed methods take one single input parameter (usually called `args`) of a struct type and also return a result of a struct type. Also note that the method name should be prefixed with `RPCServer.` in JSON-RPC.

# Example

Your client wants to set a breakpoint on the function `main.main`.
The first step will be calling the method `FindLocation` with `Scope = api.EvalScope{ GoroutineID: -1, Frame: 0}` and `Loc = "main.main"`. The JSON-RPC request packet should look like this:

```
{"method":"RPCServer.FindLocation","params":[{"Scope":{"GoroutineID":-1,"Frame":0},"Loc":"main.main"}],"id":2}
```

the response packet will look like this:

```
{"id":2,"result":{"Locations":[{"pc":4199019,"file":"/home/a/temp/callme/callme.go","line":31,"function":{"name":"main.main","value":4198992,"type":84,"goType":0}}]},"error":null}
```

Now your client should call the method `CreateBreakpoint` and specify `4199019` (the `pc` field in the response object) as the target address:

```
{"method":"RPCServer.CreateBreakpoint","params":[{"Breakpoint":{"addr":4199019}}],"id":3}
```

if this request is successful your client will receive the following response:

```
{"id":3,"result":{"Breakpoint":{"id":1,"name":"","addr":4199019,"file":"/home/a/temp/callme/callme.go","line":31,"functionName":"main.main","Cond":"","continue":false,"goroutine":false,"stacktrace":0,"LoadArgs":null,"LoadLocals":null,"hitCount":{},"totalHitCount":0}},"error":null}
```
