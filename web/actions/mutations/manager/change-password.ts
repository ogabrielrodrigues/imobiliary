'use server'

import { token } from "@/actions/queries/auth/token"
import { env } from "@/lib/env"

export async function changePassword(): Promise<number> {
  const auth_token = await token()

  try {
    const response = await fetch(`${env.SERVER_ADDR}/manager/reset`, {
      method: "PUT",
      headers: {
        "Authorization": `Bearer ${auth_token}`,
      },
    })

    return response.status
  } catch {
    return 500
  }
}