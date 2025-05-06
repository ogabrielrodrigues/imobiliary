'use server'

import { token } from "@/actions/queries/auth/token"
import { env } from "@/lib/env"
import { Owner } from "@/types/owner"

export async function getOwner(id: string): Promise<{ owner: Owner | undefined, status: number, }> {
  const auth_token = await token()

  try {
    const response = await fetch(`${env.SERVER_ADDR}/owners/${id}`, {
      method: "GET",
      headers: {
        "Authorization": `Bearer ${auth_token}`
      }
    })

    if (response.status != 200) {
      return { owner: undefined, status: response.status }
    }

    const owner = await response.json() as Owner

    return { owner, status: response.status }
  } catch {
    return { owner: undefined, status: 500 }
  }
}