import createClient from "openapi-fetch";
import type { paths } from "./api-gen";

export const client = createClient<paths>();
