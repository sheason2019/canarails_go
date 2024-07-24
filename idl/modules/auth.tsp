import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";

using TypeSpec.Http;
using TypeSpec.Rest;

model LoginReq {
  username: string;
  password: string;
}

model LoginRes {
  token: string;
}

model AuthRes {
  id: int32;
  username: string;
}

@tag("AuthController")
@route("/api/auth")
interface Auth {
  @post
  @route("login")
  login(@body loginReq: LoginReq): LoginRes;

  @get
  auth(@header authorization: string): AuthRes;
}
