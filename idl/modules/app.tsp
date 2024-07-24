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
}

@tag("AppController")
@route("/api/app")
interface Apps {
  @post
  create(@header authorization: string, @body app: App): int32;

  @get
  list(): App[];

  @delete
  delete(@header authorization: string, id: int32): int32;

  @put
  put(@header authorization: string, @body app: App): int32;

  @get
  @route("{id}")
  findById(@path id: int32): App;
}
