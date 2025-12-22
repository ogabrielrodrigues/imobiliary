'use server'

import jwt, { JwtPayload } from "jsonwebtoken";
import { env } from "./env";

export async function decodeJWT(token: string): Promise<string | undefined> {
  const payload = jwt.verify(token, env.JWT_SECRET!) as JwtPayload

  return payload.sub
}