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
  id: int32;
  username: string;
  token: string;
  expireAt: int64;
}

@tag("AuthController")
@route("/api/auth")
interface Auth {
  @post
  @route("login")
  login(
    @body loginReq: LoginReq,
    @header authorization: string,
  ): LoginRes;
}
