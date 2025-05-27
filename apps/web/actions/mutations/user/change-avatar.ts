'use server'

import { token } from "@/actions/queries/auth/token"
import { env } from "@/lib/env"

export async function changeAvatar(data: FormData): Promise<number> {
  const auth_token = await token()

  try {
    const response = await fetch(`${env.SERVER_ADDR}/users/avatar`, {
      method: "PUT",
      body: data,
      headers: {
        "Authorization": `Bearer ${auth_token}`
      }
    })

    return response.status
  } catch {
    return 500
  }
}