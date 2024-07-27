import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";

using TypeSpec.Http;
using TypeSpec.Rest;

model AppVariant {
  id: int32;
  title: string;
  description: string;
  imageName: string;
  replicas: int32;
  exposePort: int32;
  matches: AppVariantMatch[];
  appId: int32;
}

model AppVariantMatch {
  header: string;
  value: string;
}

@tag("AppVariantController")
@route("/api/app-variant")
interface AppVariants {
  @post
  @useAuth(BasicAuth)
  create(@body appVariant: AppVariant): int32;

  @get
  list(@query appId: int32): AppVariant[];

  @delete
  @route("{id}")
  @useAuth(BasicAuth)
  delete(@path id: int32): int32;

  @put
  @route("{id}")
  @useAuth(BasicAuth)
  put(@body appVariant: AppVariant, @path id: int32): int32;

  @get
  @route("{id}")
  findById(@path id: int32): AppVariant;
}
