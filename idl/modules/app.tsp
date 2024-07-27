import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";

using TypeSpec.Http;
using TypeSpec.Rest;

model App {
  id: int32;
  title: string;
  description: string;
  hostnames: string[];
  defaultVariantId: int32;
}

@tag("AppController")
@route("/api/app")
interface Apps {
  @post
  @useAuth(BasicAuth)
  create(@body app: App): int32;

  @get
  list(): App[];

  @delete
  @route("{id}")
  @useAuth(BasicAuth)
  delete(@path id: int32): int32;

  @put
  @route("{id}")
  @useAuth(BasicAuth)
  put(@body app: App, @path id: int32): int32;

  @get
  @route("{id}")
  findById(@path id: int32): App;
}
