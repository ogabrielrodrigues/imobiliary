import { User } from "@/types/user";

declare module "jsonwebtoken" {
  export interface JwtPayload {
    user: User;
  }
}
