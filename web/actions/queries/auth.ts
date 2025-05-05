'use server'

import { decodeJWT } from "@/lib/jwt"
import { token } from "./token"

export async function auth(): Promise<string | undefined> {
  const auth_token = await token()

  if (!auth_token) {
    return undefined
  }

  return await decodeJWT(auth_token)
}