'use server'

import { token } from "@/actions/queries/auth/token"
import { env } from "@/lib/env"
import { Property } from "@/types/property"

export async function getProperty(id: string): Promise<{ property: Property | undefined, status: number }> {
  const auth_token = await token()

  try {
    const response = await fetch(`${env.SERVER_ADDR}/properties/${id}`, {
      method: "GET",
      headers: {
        "Authorization": `Bearer ${auth_token}`
      }
    })

    if (response.status !== 200) {
      return { property: undefined, status: response.status }
    }

    const property = await response.json() as Property

    return { property, status: response.status }
  } catch {
    return { property: undefined, status: 500 }
  }
}