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

model OptionalAppVariant is OptionalProperties<AppVariant>;

@tag("AppVariantController")
@route("/api/app-variant")
interface AppVariants {
  @post
  @useAuth(BasicAuth)
  create(@body appVariant: OptionalAppVariant): int32;

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

  @patch
  @useAuth(BasicAuth)
  batchPatch(
    @body appVariant: OptionalAppVariant,
    @query appId: int32,
    @query title?: string,
    @query id?: int32,
  ): int32;

  @delete
  @useAuth(BasicAuth)
  batchDelete(@query title: string): int32;

  @get
  @route("{id}")
  findById(@path id: int32): AppVariant;
}
