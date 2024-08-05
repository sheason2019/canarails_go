import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";

using TypeSpec.Http;
using TypeSpec.Rest;

model Version {
  buildTime: string;
  gitHash: string;
}

@tag("VersionController")
@route("/api/version")
interface Versions {
  @get
  getVersion(): Version;
}
