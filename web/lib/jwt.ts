'use server'

import { User } from "@/types/user";
import jwt, { JwtPayload } from "jsonwebtoken"
import { env } from "./env";

export async function decodeJWT(token: string): Promise<User> {
  const payload = jwt.verify(token, env.SECRET_KEY!) as JwtPayload

  return payload.user as User
}