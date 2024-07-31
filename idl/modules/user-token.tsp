import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";

using TypeSpec.Http;
using TypeSpec.Rest;

model UserToken {
  id: int32;
  title: string;
  description: string;
  lastUsedAt: int64;
  expiredAt: int64;
}

model CreateUserTokenRes is UserToken {
  tokenString: string;
}

@tag("UserTokenController")
@route("/api/user-token")
interface UserTokens {
  @post
  @useAuth(BasicAuth)
  create(@body userToken: UserToken): CreateUserTokenRes;

  @get
  list(): UserToken[];

  @delete
  @route("{id}")
  @useAuth(BasicAuth)
  delete(@path id: int32): int32;
}
