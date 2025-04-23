'use server'

import { decodeJWT } from "@/lib/jwt"
import { User } from "@/types/user"
import { token } from "./token"

export async function auth(): Promise<User> {
  const auth_token = await token()

  return await decodeJWT(auth_token)
}